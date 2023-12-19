package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" 
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Database interface
type Database interface {
    Add(string)
    Update(string,string)
}

// MongoDB struct
type MongoDB struct{
    Client *mongo.Client
}

func (m MongoDB) Add(data string) {
    collection := m.Client.Database("test").Collection("documents")
    res, err := collection.InsertOne(context.TODO(), map[string]string{"Book": data})
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Added record to MongoDB with ID:", res.InsertedID)
}

func (m MongoDB) Update(olddata string, newdata string) {
    collection := m.Client.Database("test").Collection("documents")
	filter := bson.M{"Book": olddata}
	update := bson.M{"$set": bson.M{"Book": newdata}}
	res, err := collection.UpdateOne(context.TODO(), filter, update)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Updated records in MongoDB:", res.ModifiedCount)
}

// SQL struct
type SQL struct{
    DB *sql.DB
}

func (s SQL) Add(data string) {
    res, err := s.DB.Exec("INSERT INTO BOOKLIST (book) VALUES (?)", data)
    if err != nil {
        log.Fatal(err)
    }
    id, err := res.LastInsertId()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Added record to SQL with ID:", id)
}

func (s SQL) Update(olddata string, newdata string) {
    res, err := s.DB.Exec("UPDATE BOOKLIST SET book = ? WHERE book = ?", newdata, olddata)
    if err != nil {
        log.Fatal(err)
    }
    rows, err := res.RowsAffected()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Updated records in SQL:", rows)
}

func main() {
    var db Database

    // Connect to MongoDB Atlas
    clientOptions := options.Client().ApplyURI("mongodb+srv://Anu_Shulammite:<Password>@cluster0.kwjddrn.mongodb.net/")
    client, err := mongo.Connect(context.TODO(), clientOptions)
    if err != nil {
        log.Fatal(err)
    }

	// Clear the collection
    err = client.Database("test").Collection("documents").Drop(context.TODO())
    if err != nil {
        log.Fatal(err)
    }
	
    db = MongoDB{Client: client}
    db.Add("Silence")
	db.Add("Wings Of Fire")
    db.Update("Silence","When Breath becomes Air")

    // Connect to MySQL
    sqlDB, err := sql.Open("mysql", "root:<Password>@tcp(127.0.0.1:3306)/go_training")
    if err != nil {
        log.Fatal(err)
    }

	// Clear the table
    _, err = sqlDB.Exec("TRUNCATE TABLE BOOKLIST")
    if err != nil {
        log.Fatal(err)
    }

    db = SQL{DB: sqlDB}
    db.Add("When Breath becomes Air")
	db.Add("Wings of Fire")
    db.Update("When Breath becomes Air","Silence")
}

