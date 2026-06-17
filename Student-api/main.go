package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"net/http"
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
func getStudent(w http.ResponseWriter, r *http.Request){
	idstr := r.URL.Query().Get("id")
	id,_:=strconv.Atoi(idstr)
	var idStudent Student
	for _,student:=range studentsData{
		if student.Id==id {
			idStudent=student
		}
	}
	json.NewEncoder(w).Encode(idStudent)
	
}
func deleteStudent(w http.ResponseWriter, r *http.Request){
	idStr :=r.URL.Query().Get("id")
	id,_:=strconv.Atoi(idStr)
	
	for index,student := range studentsData{
		if student.Id==id {
			studentsData=append( 
			studentsData[:index],
			studentsData[index+1:]...,)
		}
		
	}
	fmt.Fprintln(w,"student deleted")
	
}
func updateStudent(w http.ResponseWriter,r *http.Request){
	idStr:=r.URL.Query().Get("id")
	id,_:=strconv.Atoi(idStr)
	var updatedStudent Student
	json.NewDecoder(r.Body).Decode(&updatedStudent)
	for index,student:=range studentsData{
		if student.Id==id {
			studentsData[index].Marks=updatedStudent.Marks
			studentsData[index].Name=updatedStudent.Name
			studentsData[index].Status=updatedStudent.Status

			json.NewEncoder(w).Encode(studentsData[index])
			return
		}
	}
	
}
func main(){
	http.HandleFunc("/",hello)
	http.HandleFunc("/students",students)
	fmt.Println("server is runing at port: 8080")
	http.HandleFunc("/createStudent",createStudent)
	http.HandleFunc("/student",getStudent)
	http.HandleFunc("/deleteStudent",deleteStudent)
	http.HandleFunc("/updateStudent",updateStudent)
	http.ListenAndServe(":8080",nil)
	
	
}