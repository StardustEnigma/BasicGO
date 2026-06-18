package handlers

import (
	"encoding/json"
	"fmt"
	"strconv"
	"net/http"
	"Student-api/models"
	"Student-api/storage"
)


func Hello(w http.ResponseWriter,r *http.Request){
	fmt.Fprintln(w,"Method :",r.Method)
}

func Students(w http.ResponseWriter,r *http.Request){
	json.NewEncoder(w).Encode(storage.StudentsData)
}

func CreateStudent(w http.ResponseWriter, r *http.Request){
	var student models.Student
	json.NewDecoder(r.Body).Decode(&student)
	fmt.Println(student)
	storage.StudentsData = append(storage.StudentsData,student)
	json.NewEncoder(w).Encode(student)
}

func GetStudent(w http.ResponseWriter, r *http.Request){
	idstr := r.URL.Query().Get("id")
	id,_:=strconv.Atoi(idstr)
	var idStudent models.Student
	for _,student:=range storage.StudentsData{
		if student.Id==id {
			idStudent=student
		}
	}
	json.NewEncoder(w).Encode(idStudent)
	
}

func DeleteStudent(w http.ResponseWriter, r *http.Request){
	idStr :=r.URL.Query().Get("id")
	id,_:=strconv.Atoi(idStr)
	
	for index,student := range storage.StudentsData{
		if student.Id==id {
			storage.StudentsData=append( 
			storage.StudentsData[:index],
			storage.StudentsData[index+1:]...,)
		}
		
	}
	fmt.Fprintln(w,"student deleted")
	
}

func UpdateStudent(w http.ResponseWriter,r *http.Request){
	idStr:=r.URL.Query().Get("id")
	id,_:=strconv.Atoi(idStr)
	var updatedStudent models.Student
	json.NewDecoder(r.Body).Decode(&updatedStudent)
	for index,student:=range storage.StudentsData{
		if student.Id==id {
			storage.StudentsData[index].Marks=updatedStudent.Marks
			storage.StudentsData[index].Name=updatedStudent.Name
			storage.StudentsData[index].Status=updatedStudent.Status

			json.NewEncoder(w).Encode(storage.StudentsData[index])
			return
		}
	}
	
}
