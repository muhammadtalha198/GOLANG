package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Model for courses - file
type Course struct {
	CourseId    string  `json:"course_id"`
	CourseName  string  `json:"course_name"`
	CoursePrice string  `json:"course_price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"full_name"`
	Website  string `json:"website"`
}

// fake db
var courses []Course

//middleware helper file

func (c *Course) isEmmpty() bool {
	return c.CourseId == "" && c.CourseName == ""
}

func main() {
	fmt.Println("start making a Api in go.")

}

//controllers -fie

//serveHome route

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<hi> Welcome to the mai page </h1>"))
}

func getAllCoursese(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Get all the courses")
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {

	fmt.Println("get one courses")
	w.Header().Set("Content-Type", "application/json")

	//grab an id from request

	params := mux.Vars(r)

	//loop througnjh the courses

	for _, course := range courses {

		if course.CourseId == params["id"]{
			json.NewEncoder(w).Encode(course)
			return 
		}
	}

	json.NewEncoder(w).Encode("No course found.")

}
