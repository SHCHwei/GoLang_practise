package database



import (
	"context"
	"fmt"
	"log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	ctx         context.Context
	err         error
)

type MongoDB struct{
    Collection *mongo.Collection
    Client *mongo.Client
    Ctx context.Context
}

func(m *MongoDB) connect() {

    conn := options.Client().ApplyURI("mongodb://teacher:teacher01@localhost:27017")
    m.Client, err = mongo.Connect(ctx, conn)

    if err != nil{
        log.Fatal("error while connecting with mongo", err)
    }

    err = m.Client.Ping(ctx, readpref.Primary())

    if err != nil {
        log.Fatal("error while trying to ping mongo", err)
    }

    fmt.Println("mongo connection established")

    m.Collection = m.Client.Database("stock").Collection("stocks")
    m.Ctx = ctx

}


func(m *MongoDB) disConnect(){

    m.Client.Disconnect(m.Ctx)
    log.Println("disConnect func.")
}