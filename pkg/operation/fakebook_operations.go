package operation

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type FakeNewsEntity struct {
	FakeNewsId  primitive.ObjectID `bson:"_id" json:"id"`
	UserId      string             `bson:"userId" json:"userId"`
	FakeNewsUrl string             `json:"fakeNewsUrl" bson:"fakeNewsUrl"`
	Rating      int8               `json:"rating" bson:"rating"`
}

type FakeNewsRepository interface {
	InsertFakeNews(userId, fakeNewsUrl string) (string, error)
}

type fakeNewsRepository struct {
	MongoClient *mongo.Client
}

func NewFakeNewsRepository() *fakeNewsRepository {
	return &fakeNewsRepository{}
}

func (repo *fakeNewsRepository) InsertFakeNews(userId, fakeNewsUrl string) (string, error) {
	var err error
	ctx := context.Background()
	mongoClient := repo.MongoClient
	//Create mongo collection
	collection := mongoClient.Database("fakenews_database").Collection("fakenewsrecords")
	//Create record for insertion
	id := primitive.NewObjectID()
	fakeNewsRecord := FakeNewsEntity{FakeNewsId: id, UserId: userId, FakeNewsUrl: fakeNewsUrl, Rating: 0}
	//Insert record
	result, err := collection.InsertOne(ctx, fakeNewsRecord)
	if err != nil || result == nil {
		log.Println("Error inserting to database", err)
		return "", err
	}
	return id.Hex(), err
}
