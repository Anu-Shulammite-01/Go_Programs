package mongodb

import (
	model "TemplateUserDetailsTask/Model"
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// MongoDB
type MongoDB struct {
	client *mongo.Client
}

func (db *MongoDB) CreateTemplate(key string, value model.Template) {
	collection := db.client.Database("UserInfo").Collection("Details")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res,err := collection.InsertOne(ctx,map[string]interface{}{"Key":key,"Value":value})
	if err != nil {
		panic(err)
	}
	fmt.Println("Added record to MongoDB with ID:",res.InsertedID)
}


func (db *MongoDB) UpdateTemplate(oldKey string, newKey string,value model.Template) {
	collection := db.client.Database("UserInfo").Collection("Details")
	filter := bson.D{{Key: "Key", Value: oldKey}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "Key",Value: newKey}, {Key: "Value", Value: value}}}}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res,err := collection.UpdateOne(ctx,filter,update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Updated records in MongoDB:",res.ModifiedCount)

}

func (db *MongoDB) DeleteTemplate(key string) {
	collection := db.client.Database("UserInfo").Collection("Details")
	filter := bson.D{{Key: "Key",Value: key}}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result,err:= collection.DeleteOne(ctx,filter)
	if err!=nil {
		log.Fatal(err)
	}
	fmt.Printf("Deleted %v documents from UserInfo\n", result.DeletedCount)
}

func (db *MongoDB) Refresh() error {
	//Implement refresh function for Mongo DB

	return nil
}

func (db *MongoDB) Test(string)([]string,error) {
	collection := db.client.Database("UserInfo").Collection("Details")
	var results []string
	cursor,err:=collection.Find(context.TODO(),bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())
	if err:=cursor.All(context.TODO(),&results); err!=nil{
		log.Fatal(err)
	}
	return results,nil
}
