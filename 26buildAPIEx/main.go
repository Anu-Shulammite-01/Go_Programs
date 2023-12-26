package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// model for course - file
type Course struct{
	CourseId string `json:"ID"`
	CourseName string `json:"course"`
	Price string `json:"price"`
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
	return c.CourseId == "" && c.CourseName == ""
	
}

func main(){

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
	var newCourse Course
	_ = json.NewDecoder(r.Body).Decode(&course)
}