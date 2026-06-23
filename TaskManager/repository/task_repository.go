package repository

import (
	"TaskManager/db"
	"TaskManager/models"

	"context"
	"database/sql"
)

func CreateTask(ctx context.Context, task models.Task) error {
	query := `INSERT INTO tasks
			(title,status,created_at)
			VALUES($1,$2,$3)
			`
	_, err := db.DB.ExecContext(
		ctx,
		query,
		task.Title,
		task.TaskStatus,
		task.CreatedAt,
	)
	return err
}

func GetAllTask(ctx context.Context) ([]models.Task, error) {
	query := `SELECT
				task_id,
				title,
				status,
				created_at,
				completed_at
			FROM tasks`
	rows, err := db.DB.QueryContext(
		ctx,
		query,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []models.Task

	for rows.Next() {
		var task models.Task
		var nullCompleteAt sql.NullTime
		err := rows.Scan(
			&task.TaskId,
			&task.Title,
			&task.TaskStatus,
			&task.CreatedAt,
			&nullCompleteAt,
		)
		if err != nil {
			return nil, err
		}
		if nullCompleteAt.Valid {
			task.CompletedAt = &nullCompleteAt.Time
		}
		tasks = append(tasks, task)
	}
	if err:=rows.Err(); err !=nil{
		return nil,err
	}
	return tasks, nil
}

func GetTask(ctx context.Context, id int) (models.Task, error) {
	query := `SELECT * 
			FROM tasks
			WHERE task_id = $1`
	var task models.Task
	var nullCompleteAt sql.NullTime
	row := db.DB.QueryRowContext(
		ctx,
		query,
		id,
	)
	err := row.Scan(
		&task.TaskId,
		&task.Title,
		&task.TaskStatus,
		&task.CreatedAt,
		&nullCompleteAt,
	)
	if err != nil {
		return models.Task{}, err
	}
	if nullCompleteAt.Valid {
		task.CompletedAt = &nullCompleteAt.Time
	}
	return task, nil
}

func DeleteTask(ctx context.Context,id int)(error){
	query := `DELETE FROM
				tasks 
				WHERE task_id = $1`
	
	_,err := db.DB.ExecContext(
		ctx,
		query,
		id,
	)
	return err
}
