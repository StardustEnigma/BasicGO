package routes

import (
	"TaskManager/handlers"
	"net/http"
)

func ResgisterRoutes(){
	http.HandleFunc("/CreateTask",handlers.CreateTask)
	http.HandleFunc("/Tasks",handlers.PrintAllTasks)
	http.HandleFunc("/Task",handlers.GetTask)
	http.HandleFunc("/Delete",handlers.DeleteTask)
	http.HandleFunc("/updateTask",handlers.UpdateTask)
	http.HandleFunc("/completeTask",handlers.MarkCompleted)
}
