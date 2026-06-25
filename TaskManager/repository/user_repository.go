package repository

import (
	"TaskManager/db"
	"TaskManager/models"
	"context"

)

func CreateUser(ctx context.Context,user models.User)(models.User,error){
	query := `INSERT INTO users 
				(email,password,created_at)
				VALUES ($1,$2,$3)
				RETURNING user_id,password,email,created_at`
	var newUser models.User
	err := db.DB.QueryRowContext(
		ctx,
		query,
		user.Email,
		user.Password,
		user.CreatedAt,
		).Scan(
			&newUser.UserId,
			&newUser.Email,
			&newUser.Password,
			&newUser.CreatedAt,
		)
		if err != nil {
			return models.User{},err
		}
		return newUser,nil
}

func GetUserFromLogin(ctx context.Context,email string)(models.User,error){
	query := `SELECT 
				user_id,
				email,
				password,
				created_at
				FROM users
				WHERE email = $1`
	var user models.User
	err := db.DB.QueryRowContext(
		ctx,
		query,
		email,
	).Scan(
		&user.UserId,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)
	if err != nil{
		return models.User{},err
	}
	return user,nil

}