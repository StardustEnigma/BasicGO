package services

import (
	"TaskManager/models"
	"TaskManager/storage"
	"errors"
	"time"
)
func CreateTask( task models.Task)(models.Task,error){
	if task.Title=="" {
		return models.Task{},errors.New("title is required")
	}
	task.TaskId=len(storage.TaskData)+1
	task.CreatedAt=time.Now().UTC()
	task.TaskStatus=models.StatusPending
	storage.TaskData=append(storage.TaskData, task)

	return  task,nil
}