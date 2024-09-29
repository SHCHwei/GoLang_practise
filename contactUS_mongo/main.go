package main

import (

	"context"
	"fmt"
	"log"
    "contentAPI/controllers"
    "contentAPI/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	uc          controllers.ContactController
	ctx         context.Context
	mongoclient *mongo.Client
	err         error
)


func init(){
    ctx = context.TODO() // 返回一個非nil的Context，當還未清楚要如何使用或尚不可用時可以使用TODO方式宣告。

	mongoconn := options.Client().ApplyURI("mongodb://localhost:27017")
	mongoclient, err = mongo.Connect(ctx, mongoconn)
	if err != nil {
		log.Fatal("error while connecting with mongo", err)
	}
	// 測試是否連上，可以使用Ping
	err = mongoclient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal("error while trying to ping mongo", err)
	}

	fmt.Println("mongo connection established")

    contactC := mongoclient.Database("test").Collection("contact")
	us := services.NewContactService(contactC, ctx)
    uc = controllers.New(us)
}


func main(){


    defer mongoclient.Disconnect(ctx)


    server := gin.Default()

    api := server.Group("/api")
    uc.RegisterContactRoutes(api)
    log.Fatal(server.Run(":8099"))
}