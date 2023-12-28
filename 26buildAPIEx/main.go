package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// model for course - file
type Course struct{
	CourseId string `json:"ID"`
	CourseName string `json:"course"`
	Price int `json:"price"`
	Author *Author `json:"author"`
}
type Author struct{
	Fullname string `json:"fullname"`
	Website string `json:"website"`
}

//fake DB
var coursesDB []Course

//Helpers
func (c *Course) isEmpty() bool{
	//return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
	
}

func main(){
	fmt.Println("API")
	r:=mux.NewRouter()

	// seeding
	coursesDB = append(coursesDB, Course{CourseId: "2",CourseName: "ReactJS",Price: 299, Author: &Author{Fullname:"Anu Shulammite",Website:"lco.dev"}})
	coursesDB = append(coursesDB, Course{CourseId: "4",CourseName: "GoLang",Price: 399, Author: &Author{Fullname:"Jane Williams",Website:"go.dev"}})

	//routing
	r.HandleFunc("/",serveHome).Methods("GET")
	r.HandleFunc("/courses",getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}",getOneCourse).Methods("GET")
	r.HandleFunc("/course",createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}",updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}",deleteOneCourse).Methods("DELETE")

	//listen to a port
	log.Fatal(http.ListenAndServe(":4000",r))

}

//controllers - file


//serve home route
func serveHome(w http.ResponseWriter,r *http.Request){
	w.Write([]byte("<h1>Welcome to API</h1>"))
}

func getAllCourses(w http.ResponseWriter,r *http.Request){
	w.Write([]byte("<h1>Get all Courses</h1>"))
	w.Header().Set("Cotent-Type","/application.json")
	json.NewEncoder(w).Encode(coursesDB)
}

func getOneCourse(w http.ResponseWriter,r *http.Request){
	fmt.Println("Get one course")
	w.Header().Set("Cotent-Type","/application.json")

	//get id from request
	params := mux.Vars(r)

	for _,course:=range coursesDB{
		if course.CourseId == params["id"]{
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with the given ID")
	return
}

func createOneCourse(w http.ResponseWriter,r *http.Request){
	fmt.Println("Create one course")
	w.Header().Set("Content-Type","application/json")

	//What if: body is empty
	if r.Body == nil{
		json.NewEncoder(w).Encode("Please send some data")
	}

	//What if - {}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.isEmpty(){
		json.NewEncoder(w).Encode("Fields are missing in your JSON")
		return
	}
	
	// generate uniqueid and convert it to string
	//append course into coursesDB

	source:=rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)
	course.CourseId=strconv.Itoa(rng.Intn(100))
	coursesDB=append(coursesDB,course)
	json.NewEncoder(w).Encode(course)
	return
}

func updateOneCourse(w http.ResponseWriter,r *http.Request){
	fmt.Println("Update one course")
	w.Header().Set("Content-Type","application/json")

	//grab id from req
	params := mux.Vars(r)
	
	//loop, id,remove,add with my id
	for i,v := range coursesDB {
		if v.CourseId == params["id"]{
			coursesDB = append(coursesDB[:i],coursesDB[:i+1]... )
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			v.CourseId = params["id"]
			coursesDB = append(coursesDB, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	//send a response when id is not found
	json.NewEncoder(w).Encode("No course with the given ID was found.")
}

func deleteOneCourse(w http.ResponseWriter,r *http.Request){
	fmt.Println("Delete one course")
	w.Header().Set("Content-Type","application/json")

	//Grabs the id parameter from the URL
	params := mux.Vars(r)  
	for i,v := range coursesDB{
		if v.CourseId == params["id"]{
			coursesDB = append(coursesDB[:i], coursesDB[i+1:]...)  //Remove element at index i
			json.NewEncoder(w).Encode(v)
			break
		}
	}

}