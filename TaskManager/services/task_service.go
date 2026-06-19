package services

import (
	"TaskManager/models"
	"TaskManager/storage"
	"errors"
	"strconv"
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

func GetAllTask()[]models.Task{
	return storage.TaskData
}

func GetTask(id string)(models.Task,error){
	taskid,_:=strconv.Atoi(id)
	var taskAsk models.Task
	for _, task := range storage.TaskData{
		if taskid == task.TaskId{
			taskAsk=task
		}
	}
	return taskAsk,nil
}

func DeleteTask(id string)(bool){
	taskid,_:=strconv.Atoi(id)

	for index,task := range storage.TaskData{
		if taskid==task.TaskId{
			storage.TaskData=append(
			storage.TaskData[:index],
			storage.TaskData[index+1:]...
		)	
		}
	}
	return true
}

func UpdateTask(id string,title string)(models.Task,error){
	taskid,_:=strconv.Atoi(id)
	var i int
	for index,task :=range storage.TaskData{
		if taskid==task.TaskId{
			i=index
			storage.TaskData[index].Title=title
		}
	}
	return storage.TaskData[i],nil
}

func CompleteTask(id string)(models.Task,error){
	taskid,_:=strconv.Atoi(id)
	var i int
	for index,task := range storage.TaskData{
		if taskid==task.TaskId {
			i=index
			storage.TaskData[index].CompletedAt=time.Now().UTC()
			storage.TaskData[index].TaskStatus=models.StatusCompleted
		}
	}
	return storage.TaskData[i],nil
}