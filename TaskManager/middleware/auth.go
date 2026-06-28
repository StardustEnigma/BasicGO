package middleware

import (
	"TaskManager/utils"
	"context"
	"net/http"
	"strings"
)

func AuthMiddleware(next http.Handler)http.Handler{
		return http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				authHeader := r.Header.Get("Authorization")
				if authHeader == "" {
					http.Error(w,"Missing Authorization Header",http.StatusUnauthorized)
					return 
				}
				tokenString := strings.TrimPrefix(
					authHeader,
					"Bearer ",
				)
				if tokenString==authHeader {
					http.Error(w,"Invalid authorization Header",http.StatusUnauthorized)
					return 
				}
				claims,err :=utils.ValidateToken(
					tokenString,
				)
				type contextKey string
				const UserIdKey contextKey = "userId"
				if err != nil {
					http.Error(w,"Invalid Token",http.StatusUnauthorized)
					return
				}
				ctx := context.WithValue(
					r.Context(),
					UserIdKey,
					claims.UserId,
				)
				r = r.WithContext(
					ctx,
				)
				next.ServeHTTP(
					w,
					r,
				)
			},
		)
}