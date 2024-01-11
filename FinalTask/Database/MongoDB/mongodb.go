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

func (db *MongoDB) CreateTemplate(data model.Data) {
	collection := db.client.Database("UserInfo").Collection("Details")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	//check if user already exist or not
	count, _ := collection.CountDocuments(ctx, bson.D{{Key: "Name", Value: data.Name}})
	if count == 0 {
		res, err := collection.InsertOne(ctx, data)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Created new Template with ID: ", res.InsertedID)
	}else{
		log.Printf("The User %s is already exists.", data.Name)
	}
}


func (db *MongoDB) UpdateTemplate(oldData model.Data,newData model.Data) {
	collection := db.client.Database("UserInfo").Collection("Details")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	filter := bson.D{{Key:"Name",Value: oldData.Name}}
	update := bson.D{
		{Key: "$set", Value: newData},
	}
	res,err := collection.UpdateOne(ctx,filter,update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Updated records in MongoDB:",res.ModifiedCount)
}

func (db *MongoDB) DeleteTemplate(data model.Data) {
	collection := db.client.Database("UserInfo").Collection("Details")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	filter := bson.D{{Key: "Name",Value: data.Name}}
	result,err:= collection.DeleteOne(ctx,filter)
	if err!=nil {
		log.Fatal(err)
	}
	fmt.Printf("Deleted %v documents from UserInfo\n", result.DeletedCount)
}

func (db *MongoDB) RefreshData() error {
	//Implement refresh function for Mongo DB

	return nil
}

func (db *MongoDB) TestData()([]string,error) {
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
