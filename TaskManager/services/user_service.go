package services

import (
	"TaskManager/dto"
	"TaskManager/models"
	"TaskManager/repository"
	"TaskManager/utils"
	"context"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func CreateUser(ctx context.Context, registerUser dto.RegisterRequest) (models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, err
	}
	var user models.User
	user.Email = registerUser.Email
	registerUser.Password = string(hashedPassword)
	user.Password = registerUser.Password
	user.CreatedAt = time.Now().UTC()
	savedUser, err := repository.CreateUser(ctx, user)
	if err != nil {
		return models.User{}, err
	}

	return savedUser, nil
}
func LoginUser(ctx context.Context,request dto.LoginRequest)(string,error){
	user,err := repository.GetUserFromLogin(ctx,request.Email)
	if err != nil {
		return "",err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(request.Password))
	if err != nil {
		return "",errors.New("Invalid Password")
	}
	token,err := utils.GenerateTokens(user)
	if err != nil {
		return "",err
	}
	

	return token,nil
}