package db

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect() *mongo.Database {
	dbUri, ok := os.LookupEnv("DB_URI")

	if !ok {
		log.Fatalln("DB_URI env not found!!")
	}

	dbName, ok := os.LookupEnv("DATABASE_NAME")

	if !ok {
		log.Fatalln("missing env DATABASE_NAME")
	}

	mongodbOptions := options.Client().ApplyURI(dbUri)

	client, err := mongo.Connect(context.Background(), mongodbOptions)

	if err != nil {
		log.Fatalln(err)
	}

	log.Println("connected to mongodb cluster")

	db := client.Database(dbName)

	if err := db.CreateCollection(context.Background(), "users"); err != nil {
		log.Fatalln(err)
	}

	return db

}
