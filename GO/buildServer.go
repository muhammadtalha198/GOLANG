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

// Model for courses - file
type Course struct {
	CourseId    string  `json:"course_id"`
	CourseName  string  `json:"course_name"`
	CoursePrice int     `json:"course_price"`
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
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}

func main() {

	fmt.Println("start making a Api in go.")
	r := mux.NewRouter()

	//Seeding
	courses = append(courses, Course{CourseId: "2", CourseName: "react js",
		CoursePrice: 266, Author: &Author{Fullname: "Muhammad", Website: "local.dev"}})

	courses = append(courses, Course{CourseId: "1", CourseName: "Power js",
		CoursePrice: 276, Author: &Author{Fullname: "Muhammad", Website: "go.dev"}})

	courses = append(courses, Course{CourseId: "12", CourseName: "Lifter js",
		CoursePrice: 286, Author: &Author{Fullname: "Muhammad", Website: "heap.dev"}})

	//routing

	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCoursese).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")
	//Listen to a port
	log.Fatal(http.ListenAndServe(":8000", r))

}

//controllers -fie
//serveHome route

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<hi> Welcome to the main page </h1>"))
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

		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode("No course found.")
	return

}

func createOneCourse(w http.ResponseWriter, r *http.Request) {

	fmt.Println("create one courses")
	w.Header().Set("Content-Type", "application/json")

	//what if : body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("please send some data.")
		return
	}

	var course Course

	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.isEmmpty() {
		json.NewEncoder(w).Encode("no data inside json.")
		return
	}

	// Create a new random generator
	// randGen := rand.New(rand.NewSource(time.Now().UnixNano()))
	rand.NewSource(time.Now().UnixNano())

	course.CourseId = strconv.Itoa(rand.Intn(100))

	// Use randGen for generating random numbers
	// course.CourseId = fmt.Sprintf("%d", randGen.Intn(1000000))

	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {

	fmt.Println("update one course")
	w.Header().Set("Content-Type", "application/json")

	//what if : body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("please send some data.")
		return
	}

	//get and grab id from req
	params := mux.Vars(r)

	for index, course := range courses {
		if course.CourseId == params["id"] { // Use '==' for comparison
			courses = append(courses[:index], courses[index+1:]...)

			// fmt.Println()

			var course Course

			_ = json.NewDecoder(r.Body).Decode(&course)
			// if course.isEmmpty() {
			// 	json.NewEncoder(w).Encode("no data inside json.")
			// }

			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return

		}
	}

	//Todo: send a response when id is not found
	json.NewEncoder(w).Encode("No course found with given ID")
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {

	fmt.Println("delete one course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	// loop , id , remove (index, index+1)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			break
		}

	}
	json.NewEncoder(w).Encode("Course deleted successfully")
}
