package repository

import (
	"TaskManager/db"
	"TaskManager/models"
	"context"
)

func CreateTask(ctx context.Context,task models.Task) error{
	query :=`INSERT INTO tasks
			(title,status,created_at)
			VALUES($1,$2,$3)
			`
	_,err :=db.DB.ExecContext(
		ctx,
		query,
		task.Title,
		task.TaskStatus,
		task.CreatedAt,
	)
	return err
}