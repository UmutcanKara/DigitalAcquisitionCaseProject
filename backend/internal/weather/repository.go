package weather

import (
	"context"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"strings"
)

type CLR interface {
	Database(name string, opts ...*options.DatabaseOptions) *mongo.Database
	ListDatabases(ctx context.Context, filter interface{}, opts ...*options.ListDatabasesOptions) (mongo.ListDatabasesResult, error)
}

type repository struct {
	client CLR
}

func NewRepository(c *mongo.Client) Repository {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	env := make(map[string]string)
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		env[pair[0]] = pair[1]
	}
	return &repository{c}
}

func (r *repository) findWeather(hometown string) (RepositoryBson, error) {
	result := RepositoryBson{}
	coll := r.client.Database("DigitalAcquisitionCaseProject").Collection("weather")
	err := coll.FindOne(context.TODO(), bson.M{"town": hometown}).Decode(&result)
	if err != nil {
		return RepositoryBson{}, err
	}
	return result, nil
}
func (r *repository) updateWeather(hometown string, weather HistoryWeatherResponse) error {
	updateWeather := bson.M{"town": hometown, "weather": weather}
	coll := r.client.Database("DigitalAcquisitionCaseProject").Collection("weather")
	_, err := coll.UpdateOne(context.TODO(), bson.M{"town": hometown}, updateWeather)
	if err != nil {
		return err
	}
	return nil
}
func (r *repository) addWeather(hometown string, weather HistoryWeatherResponse) error {
	insertWeather := bson.M{"town": hometown, "weather": weather}
	coll := r.client.Database("DigitalAcquisitionCaseProject").Collection("weather")
	_, err := coll.InsertOne(context.TODO(), insertWeather)
	if err != nil {
		return err
	}
	return nil
}
