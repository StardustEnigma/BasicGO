package services

import (
	"TaskManager/models"
	"TaskManager/queue"
	"TaskManager/repository"
	"context"
	"errors"
	"fmt"
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

func GetAllTask(ctx context.Context,userId int,pagestr string,limitstr string)([]models.Task,error){
	page,err := strconv.Atoi(pagestr)
	if err!=nil || page <=0 {
		page=1;
	}
	limit,err := strconv.Atoi(limitstr)
	if err != nil || limit <=0 {
		limit=10
	}
	offset := (page-1)*limit
	if err != nil {
		return []models.Task{},err
	}
	tasks,err:=repository.GetAllTask(ctx,userId,offset,limit)
	if err != nil {
		return []models.Task{},err
	}
	return tasks,nil
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

func StartTask(id string,ctx context.Context)error{
	taskid,err:=strconv.Atoi(id)
	if err != nil {
		return fmt.Errorf("invalid task id")
	}
	taskExists,err :=repository.CheckTaskExist(ctx,taskid)

	if err != nil {
		return fmt.Errorf("database error: %v", err)
	}
	if !taskExists {
		return fmt.Errorf("task with id %d does not exist", taskid)
	}
	queue.TaskQueue <- taskid

	return nil
}

func ProcessTask(taskid int,ctx context.Context) error{
	
	err := repository.ProcessTask(ctx,taskid)
	if err != nil {
		return err
	}
	time.Sleep(5*time.Second)

	_,err =repository.CompleteTask(ctx,taskid)
	if err != nil {
		return err
	}
	return nil
}

func GetAllStatusTasks(ctx context.Context,pageStr string,limitStr string,status string,userId int)([]models.Task,error){
	page,err := strconv.Atoi(pageStr)
	if err != nil || page <=0{
		page=1
	}
	limit,err := strconv.Atoi(limitStr)
	if err != nil || limit<=0{
		limit=10
	}
	offset := (page-1)*limit
	tasks,err := repository.GetStatusTasks(ctx,limit,offset,status,userId)
	if err != nil {
		return []models.Task{},err
	}
	return tasks,nil
}

