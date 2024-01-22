package mongodb

import (
	inmemory "TemplateUserDetailsTask/Database/In-Memory"
	model "TemplateUserDetailsTask/Model"
	"bytes"
	"context"
	"errors"
	"fmt"
	"text/template"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDB struct {
	Client *mongo.Client
}

func NewMongoDB(client *MongoDB) *MongoDB {
	return &MongoDB{
		Client:      client.Client,
	}
}

func (db *MongoDB) CreateTemplate(data model.Data) error {
	// Check if data is empty
	if data.Name == "" {
		return errors.New("name cannot be empty")
	}
	//check if user data.Name already exist or not
	

	// Create a new template 
	tmpl := data.Description.Value
	t, err := template.New("template").Parse(tmpl)
	if err != nil {
		return fmt.Errorf("failed to parse template: %v", err)
	}

	// Execute the template with the supplied data
	var tpl bytes.Buffer
	err = t.Execute(&tpl, data)
	if err != nil {
		return fmt.Errorf("failed to execute template: %v", err)
	}

	data.Description.Value = tpl.String()

	collection := db.Client.Database("UserInfo").Collection("Details")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	
	// Check if user already exists or not
	count, err := collection.CountDocuments(ctx, bson.D{{Key: "name", Value: data.Name}})
	if err != nil {
		return fmt.Errorf("failed to check if user exists: %v", err)
	}
	if count > 0 {
		return errors.New("user already exists")
	}

	// Insert the new user
	res, err := collection.InsertOne(ctx, data)
	if err != nil {	
		return fmt.Errorf("failed to create new user: %v", err)
	}
	fmt.Println("Created new Template with ID: ", res.InsertedID)
	return nil
}



func (db *MongoDB) UpdateTemplate(data model.Data)error {
	// Create a new template
	tmpl := data.Description.Value
	t, err := template.New("template").Parse(tmpl)
	if err != nil {
		return fmt.Errorf("failed to parse template: %v", err)
	}

	// Execute the template with the supplied data
	var tpl bytes.Buffer
	err = t.Execute(&tpl, data)
	if err != nil {
		return fmt.Errorf("failed to execute template: %v", err)
	}

	data.Description.Value = tpl.String()
	
	collection := db.Client.Database("UserInfo").Collection("Details")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
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
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
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

func (db *MongoDB) RefreshData(inMemoryDB *inmemory.InMemoryDB, data string)error{
		// Context
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		
		collection := db.Client.Database("UserInfo").Collection("Details")
		filter := bson.D{{Key:"name",Value: data}}

		// Retrieve the document from MongoDB
		var result model.Data
		err := collection.FindOne(ctx, filter).Decode(&result)
		if err != nil {
			return fmt.Errorf("error getting data from MongoDB: %s", err)
		}
	
		// Write the value to the in-memory database
		inMemoryDB.User[result.Name] = result.Description
	
		fmt.Printf("Refreshed in-memory database; Key: %s, Value: %s\n", result.Name, result.Description.Key+":"+result.Description.Value)
	
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
