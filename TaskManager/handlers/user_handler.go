package handlers

import (
	"TaskManager/dto"
	"TaskManager/services"
	"encoding/json"
	"fmt"
	"net/http"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if r.Method != http.MethodPost {
		http.Error(w, "Method not Allowed", http.StatusBadRequest)
		return
	}
	var registerUser dto.RegisterRequest
	err := json.NewDecoder(r.Body).Decode(&registerUser)
	savedUser, err := services.CreateUser(ctx, registerUser)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	fmt.Println(savedUser)
	json.NewEncoder(w).Encode(savedUser)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	var loginRequest dto.LoginRequest
	json.NewDecoder(r.Body).Decode(&loginRequest)
	token, err := services.LoginUser(ctx, loginRequest)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}
	fmt.Printf("Email: '%s'\n", loginRequest.Email)
	fmt.Printf("Password: '%s'\n", loginRequest.Password)
	fmt.Println("Password length:", len(loginRequest.Password))
	response := dto.Loginresponse{
		Token: token,
	}
	json.NewEncoder(w).Encode(response)
}
