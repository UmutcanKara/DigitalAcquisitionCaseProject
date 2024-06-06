package auth

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CLR interface {
	Database(name string, opts ...*options.DatabaseOptions) *mongo.Database
	ListDatabases(ctx context.Context, filter interface{}, opts ...*options.ListDatabasesOptions) (mongo.ListDatabasesResult, error)
}

type repository struct {
	client CLR
}

func NewRepository(c *mongo.Client) Repository { return &repository{c} }

func (r *repository) login(username, password string) (bson.M, error) {
	result := bson.M{username: "", password: ""}
	coll := r.client.Database("DigitalAcquisitionCaseProject").Collection("users")
	err := coll.FindOne(context.TODO(), bson.M{"username": username}).Decode(&result)
	if err != nil {
		return bson.M{}, err
	}
	if errors.Is(err, mongo.ErrNoDocuments) {
		return bson.M{}, errors.New("user not found")
	}

	return result, nil
}

func (r *repository) register(username, password, hometown string) error {
	insertUser := bson.M{"username": username, "password": password, "hometown": hometown}
	coll := r.client.Database("DigitalAcquisitionCaseProject").Collection("users")
	_, err := coll.InsertOne(context.TODO(), insertUser)
	if err != nil {
		return err
	}
	return nil
}

// mongodb+srv://testuser:testuser@users.n6ag96y.mongodb.net/?retryWrites=true&w=majority&appName=users
// mongodb+srv://testuser:testuser@users.n6ag96y.mongodb.net/?retryWrites=true&w=majority&appName=users
