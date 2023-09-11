package services

import (
	"context"
	"contentAPI/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"time"
)



type ContactServiceImpl struct {
	ContactCollection *mongo.Collection
	ctx            context.Context
}

//初始化"服務"
func NewContactService(contactCollection *mongo.Collection, ctx context.Context) ContactService {
    return &ContactServiceImpl{
        ContactCollection: contactCollection,
        ctx: ctx,
    }

}

func (u *ContactServiceImpl) CreateContact(data *models.ContactOne) error {
    _,err := u.ContactCollection.InsertOne(u.ctx, data)
    return err
}



func (u *ContactServiceImpl) SearchContact(StartDate time.Time, EndDate time.Time) ([]models.Contacts, error) {

    var results []models.Contacts

    log.Println("doing search time")

    filter := bson.D{
       {"$and",
          bson.A{
             bson.D{{"time", bson.D{{"$gte", StartDate}}}},
             bson.D{{"time", bson.D{{"$lte", EndDate}}}},
          },
       },
    }

    cursor, _ := u.ContactCollection.Find(u.ctx, filter)

    err := cursor.All(u.ctx, &results)

    return results, err
}

