package utils

import (
	"TaskManager/models"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserId int `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

func GenerateTokens(user models.User)(string,error){
	claims := Claims{
		UserId: user.UserId,
		Email: user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(
				time.Now().Add(24*time.Hour),
			),
			IssuedAt: jwt.NewNumericDate(
				time.Now(),
			),
		},
	}
	token := jwt.NewWithClaims(
		jwt.SigningMethodES256,
		claims,
	)
	tokenString,err :=token.SignedString(
		"22233",
	)
	if err != nil{
		return "",err
	}
	return tokenString,nil
}

func ValidateToken(tokenString string)(*Claims,error){
	claims := &Claims{}
	token, err :=jwt.ParseWithClaims(
		tokenString,
		claims,
		func(token *jwt.Token) (interface{}, error) {
			return "22233",nil
		},
	)
	if err != nil {
		return nil,err
	}
	if !token.Valid {
		return nil,
		errors.New("Invalid token")
	}
	return claims,nil
}