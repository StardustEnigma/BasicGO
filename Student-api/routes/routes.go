package routes

import (
	"Student-api/handlers"
	"net/http"
)

func RegisterRoutes(){
	http.HandleFunc("/",handlers.Hello)
	http.HandleFunc("/students",handlers.Students)
	http.HandleFunc("/createStudent",handlers.CreateStudent)
	http.HandleFunc("/student",handlers.GetStudent)
	http.HandleFunc("/deleteStudent",handlers.DeleteStudent)
	http.HandleFunc("/updateStudent",handlers.UpdateStudent)
}