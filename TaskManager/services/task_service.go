package services

import (
	"TaskManager/models"
	"TaskManager/queue"
	"TaskManager/repository"
	"TaskManager/storage"
	"context"
	"errors"
	"strconv"
	"time"
)
func CreateTask( task models.Task,ctx context.Context)(models.Task,error){
	if task.Title=="" {
		return models.Task{},errors.New("title is required")
	}
	task.CreatedAt=time.Now().UTC()
	task.TaskStatus=models.StatusPending
	err := repository.CreateTask(
		ctx,
		task,
	)
	if err !=nil {
		return models.Task{},err
	}
	return  task,nil
}

func GetAllTask(ctx context.Context)[]models.Task{
	return storage.TaskData
}

func GetTask(id string,ctx context.Context)(models.Task,error){
	taskid,_:=strconv.Atoi(id)
	var taskAsk models.Task
	for _, task := range storage.TaskData{
		if taskid == task.TaskId{
			taskAsk=task
		}
	}
	return taskAsk,nil
}

func DeleteTask(id string,ctx context.Context)(bool){
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

func UpdateTask(id string,title string,ctx context.Context)(models.Task,error){
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

func CompleteTask(id string,ctx context.Context)(models.Task,error){
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

func StartTask(id string,ctx context.Context){
	taskid,_:=strconv.Atoi(id)
	
	queue.TaskQueue <- taskid
}

func ProcessTask(taskid int){
	for i,task := range storage.TaskData{
		if taskid == task.TaskId{
			storage.Mu.Lock()
			storage.TaskData[i].TaskStatus=models.StatusProgress
			storage.Mu.Unlock()

			time.Sleep(5*time.Second)

			storage.Mu.Lock()
			storage.TaskData[i].CompletedAt=time.Now().UTC()
			storage.TaskData[i].TaskStatus=models.StatusCompleted
			storage.Mu.Unlock()
		}
	}
}