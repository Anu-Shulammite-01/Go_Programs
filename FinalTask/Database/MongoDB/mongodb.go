package mongodb

import (
	model "TemplateUserDetailsTask/Model"
	"bytes"
	"context"
	"fmt"
	"log"
	"text/template"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDB struct {
	Client *mongo.Client
	UpdateChan chan model.Data
	DeleteChan chan string
}

func NewMongoDB(client *MongoDB) *MongoDB {
	return &MongoDB{
		Client:      client.Client,
		UpdateChan: make(chan model.Data,100),
		DeleteChan: make(chan string,100),
	}
}

func (db *MongoDB) CreateTemplate(data model.Data)error {
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
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()
	filter := bson.D{{Key:"name",Value: data.Name}}
	update := bson.D{
		{Key: "$set", Value: bson.D{{Key: "description", Value: data.Description}}},
	}
	_ = collection.FindOneAndUpdate(ctx,filter,update)
	fmt.Println("Updated records in MongoDB")
	db.UpdateChan <- data
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
		db.DeleteChan <- data
	} else {
		fmt.Printf("No Record found for deletion.\n")
	}
	return nil
}

func (db *MongoDB) RefreshData(appState *model.AppState) {
	go func() {
		for {
			select {
			case data1 := <-db.UpdateChan:
				appState.Templates[data1.Name] = data1.Description
				fmt.Printf("Updated appState; Key: %s, Template: %+v\n", data1.Name, data1.Description)
			case data2 := <-db.DeleteChan:
				delete(appState.Templates, data2)
				fmt.Printf("Deleted from appState; Key: %s\n", data2)
			}
		}
	}()
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
