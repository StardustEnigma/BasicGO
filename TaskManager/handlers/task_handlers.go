package handlers

import (
	"TaskManager/dto"
	"TaskManager/models"
	"TaskManager/services"
	"encoding/json"
	"fmt"
	"net/http"
	
)


func CreateTask(w http.ResponseWriter,r *http.Request){
if r.Method !=http.MethodPost{
	http.Error(w,"Method Not Allowed",http.StatusMethodNotAllowed)
	return
}
var task models.Task
err :=json.NewDecoder(r.Body).Decode(&task)
if err !=nil {
	http.Error(w,"Invalid Json",http.StatusBadRequest)
	return
}
Createdtask,err:=services.CreateTask(task)
if err != nil {
	http.Error(w,"Bad Request",http.StatusBadRequest)
	return
}

fmt.Println(Createdtask)
w.Header().Set("Content-Type","application/json")

w.WriteHeader(http.StatusCreated)

json.NewEncoder(w).Encode(Createdtask)
}

func PrintAllTasks(w http.ResponseWriter,r *http.Request){
	if r.Method != http.MethodGet{
		http.Error(w,"Method Not allowed",http.StatusMethodNotAllowed)
		return
	}
	taskData := services.GetAllTask()
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(taskData)
}

func GetTask(w http.ResponseWriter,r *http.Request){
	if r.Method !=http.MethodGet {
		http.Error(w,"Method not Allowed",http.StatusMethodNotAllowed)
		return
	}
	idstr :=r.URL.Query().Get("id")
	if idstr=="" {
		http.Error(w,"Id is required",http.StatusBadRequest)
		return
	}
	task, err:=services.GetTask(idstr)
	if err != nil {
	http.Error(w,"Bad Request",http.StatusBadRequest)
	return
}
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(task)
}

func DeleteTask(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodDelete {
		http.Error(w,"Method Not Allowed",http.StatusMethodNotAllowed)
		return
	}
	idstr := r.URL.Query().Get("id")

	if idstr=="" {
		http.Error(w,"Id is required",http.StatusBadRequest)
		return
	}
	result := services.DeleteTask(idstr)

	if result==false {
		http.Error(w,"Cannot find Task to delete",http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type","application/json")
	fmt.Fprint(w,"Task deleted successfully !!")
}

func UpdateTask(w http.ResponseWriter,r *http.Request){
	if r.Method != http.MethodPut{
		http.Error(w,"Method Not Allowed",http.StatusMethodNotAllowed)
		return
	}
	idstr := r.URL.Query().Get("id")

	if idstr==""{
		http.Error(w,"Id is required",http.StatusBadRequest)
		return
	}
	var request dto.UpdateTitle
	err :=json.NewDecoder(r.Body).Decode(&request)
	if err !=nil {
		http.Error(w,"Invalid Json" ,http.StatusBadRequest)
		return
	}
	task,err:=services.UpdateTask(idstr,request.Title)
	if err != nil {
		http.Error(w,"Bad Request",http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(task)

}

func MarkCompleted(w http.ResponseWriter,r *http.Request){
	if r.Method != http.MethodPost {
		http.Error(w,"Method Not Allowed",http.StatusMethodNotAllowed)
		return
	}
	idstr :=r.URL.Query().Get("id")

	if idstr=="" {
		http.Error(w,"Id is required",http.StatusBadRequest)
		return
	}
	var response dto.TaskComplete
	task,err:=services.CompleteTask(idstr)
	if err!=nil {
		http.Error(w,"Bad Request",http.StatusBadRequest)
		return
	}
	response.Title=task.Title
	response.TaskId=task.TaskId
	response.CreatedAt=task.CreatedAt
	response.CompletedAt=task.CompletedAt
	
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(response)
}

func StartTask(w http.ResponseWriter,r *http.Request){
	if r.Method != http.MethodPost {
		http.Error(w,"Method Not allowed",http.StatusMethodNotAllowed)
		return
	}
	idstr:=r.URL.Query().Get("id")

	if idstr =="" {
		http.Error(w,"Id is required",http.StatusBadRequest)
		return
	}
	services.StartTask(idstr)

	fmt.Fprintln(w,"Task Started")
	fmt.Println("Task Started")
	
}