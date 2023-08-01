package store

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"try-golang/utils/protos"
)

type Store interface {
	TestConnection()
	Connect()
	SavePurchase(purchase_event *protos.PurchaseCloudEvent) bool
}

type StoreImpl struct {
	balancesCollection *mongo.Collection
	database           *mongo.Database
	isConnected        bool
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
	store.isConnected = true
	fmt.Println("Connected to MongoDB!")
}

func (store *StoreImpl) SavePurchase(purchaseEvent *protos.PurchaseCloudEvent) bool {
	fmt.Println("Save purchase with amount: ", purchaseEvent.Data.Amount)
	fmt.Println("Connected: ", store.isConnected)
	opts := options.Update().SetUpsert(true)
	filter := bson.M{"customerId": purchaseEvent.Data.CustomerId}
	update := bson.M{
		"$inc": bson.M{"amount": int64(purchaseEvent.Data.Amount)},
		"$set": bson.M{"updatedAt": time.Now().UTC()},
	}
	result, err := store.balancesCollection.UpdateOne(context.Background(), filter, update, opts)
	if err != nil {
		fmt.Println("Error updating document:", err)
		return false
	}
	fmt.Println("Matched count:", result.MatchedCount)
	fmt.Println("Modified count:", result.ModifiedCount)
	return true
}

func StoreProvider() Store {
	return &StoreImpl{
		balancesCollection: nil,
		database:           nil,
	}
}
