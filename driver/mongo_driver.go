package driver

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoDB struct {
	Client *mongo.Client
}

var Mongo = &MongoDB{}

func ConnectMongoDB(user, password string) *MongoDB { 
	connStr := fmt.Sprintf("mongodb://%s:%s@ds129098.mlab.com:29098/go", user, password)
	client, err := mongo.NewClient(options.Client().ApplyURI(connStr))
	if err != nil {
		panic(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10 * time.Second)
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		fmt.Println("Fail To Connect To Mongo")
		panic(err)
	}

	fmt.Println("MongoDB connected.")
	Mongo.Client = client
	return Mongo
}
