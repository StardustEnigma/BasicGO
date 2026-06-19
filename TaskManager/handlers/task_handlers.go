package handlers

import (
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
Createdtask,_:=services.CreateTask(task)
fmt.Println(Createdtask)
w.Header().Set("Content-Type","application/json")

w.WriteHeader(http.StatusCreated)

json.NewEncoder(w).Encode(Createdtask)
}


