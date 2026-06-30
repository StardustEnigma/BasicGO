package repository

import (
	"TaskManager/db"
	"TaskManager/models"
	"context"
	"database/sql"
	"fmt"
)

func CreateTask(ctx context.Context, task models.Task)(models.Task ,error) {
	query := `INSERT INTO tasks
			(title,status,created_at,user_id)
			VALUES($1,$2,$3,$4)
			RETURNING task_id, title, status, created_at, completed_at,user_id
			`
	var newTask models.Task
	var nullCompleteAt sql.NullTime
	 err := db.DB.QueryRowContext(
		ctx,
		query,
		task.Title,
		task.TaskStatus,
		task.CreatedAt,
		task.UserId,
	).Scan(
		&newTask.TaskId,
		&newTask.Title,
		&newTask.TaskStatus,
		&newTask.CreatedAt,
		&nullCompleteAt,
		&newTask.UserId,
	)
	if err != nil {
		return models.Task{},err
	}
	if nullCompleteAt.Valid {
		newTask.CompletedAt=&nullCompleteAt.Time
	}
	return newTask,nil
}

func GetAllTask(ctx context.Context,userId int,offset int,limit int) ([]models.Task, error) {
	query := `SELECT
				task_id,
				title,
				status,
				created_at,
				completed_at,
				user_id
			FROM tasks
			WHERE user_id = $1
			LIMIT $2
			OFFSET $3`
	rows, err := db.DB.QueryContext(
		ctx,
		query,
		userId,
		limit,
		offset,
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
			&task.UserId,
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
	
	result,err := db.DB.ExecContext(
		ctx,
		query,
		id,
	)

	rowsAffected,err:=result.RowsAffected()
	if err != nil{
		return err
	}
	if rowsAffected==0 {
		return fmt.Errorf("Id : %d not present",id)
	}
	return nil
}

func UpdateTask(ctx context.Context,id int,title string)(models.Task,error){
	query := `UPDATE tasks
				SET title = $2
				WHERE task_id = $1
				RETURNING title, task_id, status, created_at, completed_at`
	var task models.Task
	var nullCompleteAt sql.NullTime
	err:=db.DB.QueryRowContext(
		ctx,
		query,
		id,
		title,
	).Scan(
		&task.Title,
		&task.TaskId,
		&task.TaskStatus,
		&task.CreatedAt,
		&nullCompleteAt,
	)
	if err != nil {
		return models.Task{},err
	}
	if nullCompleteAt.Valid {
		task.CompletedAt=&nullCompleteAt.Time
	}
	return task,nil
	
}
func CompleteTask(ctx context.Context,id int)(models.Task,error){
	query := `UPDATE tasks
				SET completed_at=NOW(),
				 status=$2
				WHERE task_id = $1
				RETURNING task_id, title, status, created_at, completed_at`
	var task models.Task
	var nullCompleteAt sql.NullTime
	err := db.DB.QueryRowContext(
		ctx,
		query,
		id,
		models.StatusCompleted,
	).Scan(
		&task.TaskId,
		&task.Title,
		&task.TaskStatus,
		&task.CreatedAt,
		&nullCompleteAt,
	)
	if err !=nil {
		if err == sql.ErrNoRows {
			return models.Task{},fmt.Errorf("Task with %d not found",id)
		}
		return models.Task{},err
	}
	if nullCompleteAt.Valid {
		task.CompletedAt=&nullCompleteAt.Time
	}
	return task,nil

}

func ProcessTask(ctx context.Context,id int)(error){
	query :=`UPDATE tasks
			SET status = $2
			where task_id = $1`
	
	result,err :=db.DB.ExecContext(
		ctx,
		query,
		id,
		models.StatusProgress,
	)
	if err != nil {
		return err
	}
	rowsAffected,err :=result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected==0 {
		return fmt.Errorf("cannot find %d id ",id)
	}

	return nil
}

func CheckTaskExist(ctx context.Context,id int)(bool,error){
 query := `SELECT 1 
 			FROM tasks
			WHERE task_id=$1`
	var dummy int

	err := db.DB.QueryRowContext(ctx,query,id).Scan(dummy)

	if err != nil {
		if err==sql.ErrNoRows{
			return false,nil
		}
		return false,err
	}

	return true,nil
}

func GetStatusTasks(ctx context.Context,limit int,offset int,status string,userId int)([]models.Task,error){
	query := `SELECT 
				task_id,
				title,
				status,
				created_at,
				completed_at,
				user_id
				FROM tasks
				WHERE user_id = $4 AND status = $3
				LIMIT $1
				OFFSET $2
`
	 rows,err:= db.DB.QueryContext(
		ctx,
		query,
		limit,
		offset,
		status,
		userId,

	)
	if err != nil {
		return nil,err
	}
	defer rows.Close()
	var tasks []models.Task
	for rows.Next(){
		var task models.Task
		var nullCompleteAt sql.NullTime

		err:=rows.Scan(
			&task.TaskId,
			&task.Title,
			&task.TaskStatus,
			&task.CreatedAt,
			&nullCompleteAt,
			&task.UserId,
		)
		if err != nil {
			return []models.Task{},err
		}
		if nullCompleteAt.Valid{
			task.CompletedAt=&nullCompleteAt.Time
		}
		tasks = append(tasks, task)
	}
	return tasks,nil
}