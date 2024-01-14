package mongodb

import (
	model "TemplateUserDetailsTask/Model"
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// MongoDB
type MongoDB struct {
	Client *mongo.Client
}

func (db *MongoDB) CreateTemplate(data model.Data)error {
	collection := db.Client.Database("UserInfo").Collection("Details")
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()
	//check if user already exist or not
	count, err := collection.CountDocuments(ctx, bson.D{{Key: "Name", Value: data.Name}})
	if err != nil {
		return fmt.Errorf("failed to check if user exists: %v", err)
	}
	if count == 0 {
		res, err := collection.InsertOne(ctx, data)
		if err != nil {
			return fmt.Errorf("failed to create new user: %v", err)
		}
		fmt.Println("Created new Template with ID: ", res.InsertedID)
	}else{
		log.Printf("The User %s is already exists.", data.Name)
	}
	return nil
}


func (db *MongoDB) UpdateTemplate(data model.Data)error {
	collection := db.Client.Database("UserInfo").Collection("Details")
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()
	filter := bson.D{{Key:"name",Value: data.Name}}
	update := bson.D{
		{Key: "$set", Value: bson.D{{Key: "description", Value: data.Description}}},
	}
	_ = collection.FindOneAndUpdate(ctx,filter,update)
	fmt.Println("Updated records in MongoDB")
	return nil
}

func (db *MongoDB) DeleteTemplate(data string)error {
	collection := db.Client.Database("UserInfo").Collection("Details")
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()
	result,err := collection.DeleteOne(ctx,bson.D{{Key: "name",Value : data}})
	if err != nil {
		return fmt.Errorf("failed to delete the record: %v", err)
	}
	if result.DeletedCount > 0 {
		fmt.Printf("Successfully deleted the record of %v\n", data)
	} else {
		fmt.Printf("No Record found for deletion.\n")
	}
	return nil
}

func (db *MongoDB) RefreshData(appState *model.AppState)error{
	// Specify the context and the collection
	ctx := context.Background()
	collection := db.Client.Database("UserInfo").Collection("Details")
	// Create a cursor for the Find operation
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return fmt.Errorf("failed to get documents from MongoDB: %v", err)
	}
	defer cursor.Close(ctx)

	var wg sync.WaitGroup
	// For each document, fetch the associated value and update your application's state
	for cursor.Next(ctx) {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// Decode the document into a Template object
			var template model.Data
			if err := cursor.Decode(&template); err != nil {
				fmt.Printf("failed to decode document: %s\n", err)
				return
			}

			// Update the application's state with the new template
			appState.Templates[template.Name] = template.Description
			fmt.Printf("From MongoDB ; Key: %s, Template: %+v\n", template.Name, template.Description)
		}()
	}

	wg.Wait()

	if err := cursor.Err(); err != nil {
		return fmt.Errorf("cursor encountered an error: %v", err)
	}

	return nil
}

func (db *MongoDB) TestData()([]string,error) {
	collection := db.Client.Database("UserInfo").Collection("Details")
	var results []string
	cursor,err:=collection.Find(context.TODO(),bson.D{})
	if err != nil {
		return nil,err
	}
	for cursor.Next(context.TODO()) {
		var elem struct {
			Name string `json:"name"`
			Description model.Template `json:"description"`
		}
		err = cursor.Decode(&elem)
		if err != nil {
			return nil,err
		}
		results=append(results,elem.Name+" : "+elem.Description.Key+" = "+elem.Description.Value)
	}
	return results,nil
}
