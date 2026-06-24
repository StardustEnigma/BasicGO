package services

import (
	"TaskManager/dto"
	"TaskManager/models"
	"TaskManager/queue"
	"TaskManager/repository"
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"
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

func CreateUser(ctx context.Context,registerUser dto.RegisterRequest)(models.User,error){
	hashedPassword, err:= bcrypt.GenerateFromPassword([]byte(registerUser.Password),bcrypt.DefaultCost)
	if err != nil {
		return models.User{},err
	}
	var user models.User
	user.Email=registerUser.Email
	registerUser.Password=string(hashedPassword)
	user.Password=registerUser.Password
	user.CreatedAt=time.Now().UTC()
	savedUser,err:=repository.CreateUser(ctx,user)
	if err != nil {
		return models.User{},err
	}

	return savedUser,nil
}