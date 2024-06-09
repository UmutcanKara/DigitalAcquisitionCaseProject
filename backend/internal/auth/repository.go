package auth

import (
	"backend/util"
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

func (r *repository) login(username, password string) error {
	//result := bson.M{username: "", password: ""}
	loginReq := LoginReq{}
	coll := r.client.Database("DigitalAcquisitionCaseProject").Collection("users")
	err := coll.FindOne(context.TODO(), bson.M{"username": username}).Decode(&loginReq)
	if err != nil {
		return err
	}
	if errors.Is(err, mongo.ErrNoDocuments) {
		return errors.New("user not found")
	}
	if err = util.CheckPassword(password, loginReq.Password); err != nil {
		return err
	}
	return nil
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

func (r *repository) getUser(username string) (bson.M, error) {
	result := bson.M{}
	coll := r.client.Database("DigitalAcquisitionCaseProject").Collection("users")
	err := coll.FindOne(context.TODO(), bson.M{"username": username}).Decode(&result)
	if err != nil {
		return bson.M{}, err
	}

	return result, nil
}
