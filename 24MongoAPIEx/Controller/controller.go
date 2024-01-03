package controller

import (
	model "MongoAPIExample/Model"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://Anu_Shulammite:Password@cluster0.kwjddrn.mongodb.net/"
const dbName = "netflix"
const colName = "watchlist"

var collection *mongo.Collection //reference of mongodb collection

//connect with mongoDB
func init(){
	clientOptions := options.Client().ApplyURI(connectionString)
	//connect
	client,err:=mongo.Connect(context.TODO(),clientOptions)
	if err!=nil{
		log.Fatalln(err)
	}
	collection = client.Database(dbName).Collection(colName) //reached inside the collection, it is a reference
	fmt.Println("Collection reference is ready")
}

//insert 1 record
func insertOneMovie(movie model.Netflix){
	inserted,err := collection.InsertOne(context.Background(),movie)
	if err!= nil{
		log.Fatalln(err)
	}
	fmt.Printf("%v is inserted into db with ID:%v\n",movie,inserted.InsertedID)
}

//updated 1 record
func updateOneMovie(movieId string){
	id,_ := primitive.ObjectIDFromHex(movieId)
	filter:=bson.M{"_id":id}
	update:=bson.M{"$set":bson.M{"watched":true}}
	updated,err := collection.UpdateOne(context.Background(),filter,update)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Modified count",updated.ModifiedCount)
}

//delete 1 record
func deleteOneMovie(movieId string){
	id,_ := primitive.ObjectIDFromHex(movieId)
	filter:=bson.M{"_id":id}
	deleted,err:=collection.DeleteOne(context.Background(),filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted ", deleted.DeletedCount," document(s)")
}

//Delete all records
func deleteAllMovie() int64{
	deleteres,err:=collection.DeleteMany(context.Background(),bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted ", deleteres.DeletedCount," documents in DB")
	return deleteres.DeletedCount
}

//get all movies
func getAllMovies() []primitive.M{
	cur,err:=collection.Find(context.Background(),bson.D{{}})
	if err !=nil{
		log.Fatal(err)
	}
	defer cur.Close(context.TODO())

	//Method1
	var movies[]primitive.M
	for cur.Next(context.Background()){
		var movie bson.M
		if err=cur.Decode(&movie);err!=nil{
			log.Fatal(err)
		}
		movies = append(movies,movie)
	}
	return movies

	//Method2
	// var movies []bson.M
	// if err:= cur.All(context.TODO(),&movies);err!=nil{
	// 	log.Fatal(err)
	// }
	// for _,movie:=range movies{
	// 	fmt.Println(movie)
	// }	
}

// Actual controller file

func GetMyAllMovies(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/x-www-form-urlencode")
	allmovies := getAllMovies()
	json.NewEncoder(w).Encode(allmovies)
}

func CreateMovie(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods","POST")

	var movie model.Netflix
	_=json.NewDecoder(r.Body).Decode(&movie)
	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)
}

func MarkAsWatched(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods","PUT")

	params:=mux.Vars(r)
	updateOneMovie(params["id"])
	json.NewEncoder(w).Encode("Marked as watched!")
}

func DeleteAMovie(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods","DELETE")

	params:=mux.Vars(r)
	deleteOneMovie(params["id"])
	json.NewEncoder(w).Encode("Removed from Netflix Library.")
}

func DeleteAlMovies(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods","DELETE")

	count:= deleteAllMovie()
	json.NewEncoder(w).Encode(count)
}
