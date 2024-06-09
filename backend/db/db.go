package db

import (
	"context"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

type Database struct {
	db *mongo.Client
}

func ConnectDB() (*Database, error) {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	uri := os.Getenv("MONGODB_URI")
	log.Printf("Connecting to MongoDB at %s", uri)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	return &Database{client}, nil
}

func (d *Database) GetClient() *mongo.Client { return d.db }

func (d *Database) Close() {
	if err := d.db.Disconnect(context.TODO()); err != nil {
		log.Fatal(err)
	}
}
