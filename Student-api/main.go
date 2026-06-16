package main

import (
	"fmt"
	"net/http"
	"encoding/json"
)
type Student struct{
	Id int `json:"id"`
	Name string `json:"name"`
	Marks int `json:"marks"`
	Status string `json:"status"`
}
var studentsData = []Student{
		{1,"Atharva",86,"Passes"},
		{2, "Demo", 70, "Passed"},
		{3, "Check", 35, "Failed"},
	}
func hello(w http.ResponseWriter,r *http.Request){
	fmt.Fprintln(w,"Method :",r.Method)
}
func students(w http.ResponseWriter,r *http.Request){
	json.NewEncoder(w).Encode(studentsData)
}
func createStudent(w http.ResponseWriter, r *http.Request){
	var student Student
	json.NewDecoder(r.Body).Decode(&student)
	fmt.Println(student)
	studentsData = append(studentsData,student)
	json.NewEncoder(w).Encode(student)
}
func main(){
	http.HandleFunc("/",hello)
	http.HandleFunc("/students",students)
	fmt.Println("server is runing at port: 8080")
	http.HandleFunc("/createStudent",createStudent)
	http.ListenAndServe(":8080",nil)

	
}