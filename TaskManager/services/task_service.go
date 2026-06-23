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
	task,err := repository.CreateTask(
		ctx,
		task,
	)
	if err !=nil {
		return models.Task{},err
	}
	return  task,nil
}

func GetAllTask(ctx context.Context)([]models.Task,error){
	tasks,err:=repository.GetAllTask(ctx)
	if err != nil {
		return []models.Task{},err
	}
	return tasks,err
}

func GetTask(id string,ctx context.Context)(models.Task,error){
	taskid,_:=strconv.Atoi(id)
	
	taskAsk,err :=repository.GetTask(ctx,taskid)
	if err != nil {
		return models.Task{},err
	}
	return taskAsk,nil
}

func DeleteTask(id string,ctx context.Context)(error){
	taskid,_:=strconv.Atoi(id)

	err := repository.DeleteTask(ctx,taskid)
	if err != nil {
		return err
	}
	return nil
}

func UpdateTask(id string,title string,ctx context.Context)(models.Task,error){
	taskid,_:=strconv.Atoi(id)
	task,err:= repository.UpdateTask(ctx,taskid,title)
	if err !=nil {
		return models.Task{},err
	}
	return task,nil
}

func CompleteTask(id string,ctx context.Context)(models.Task,error){
	taskid,_:=strconv.Atoi(id)
	task,err:=repository.CompleteTask(ctx,taskid)
	if err != nil {
		return models.Task{},err
	}
	return task,nil
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
			currentTime := time.Now().UTC()
			storage.TaskData[i].CompletedAt=&currentTime
			storage.TaskData[i].TaskStatus=models.StatusCompleted
			storage.Mu.Unlock()
		}
	}
}