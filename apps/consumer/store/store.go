package store

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Store interface {
	TestConnection()
	Connect()
}

type StoreImpl struct {
	balancesCollection *mongo.Collection 
	database *mongo.Database
}

func (store *StoreImpl) TestConnection() {
	fmt.Println("Test connection")
}

func (store *StoreImpl) Connect() {
	fmt.Println("Connect")
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_URI"))
	context := context.TODO()
	client, err := mongo.Connect(context, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context, nil)
	if err != nil {
		log.Fatal(err)
	}
	database := client.Database("try-golang")
	balancesCollection := database.Collection("balances")
	store.database = database
	store.balancesCollection = balancesCollection
	fmt.Println("Connected to MongoDB!")
}

func StoreProvider() Store {
	return &StoreImpl{
		balancesCollection: nil,
		database: nil,
	}
}