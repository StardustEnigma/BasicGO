package handlers

import (
	"TaskManager/dto"
	"TaskManager/services"
	"encoding/json"
	"fmt"
	"net/http"
)

func RegisterUser(w http.ResponseWriter, r *http.Request){
	ctx := r.Context()

	if r.Method != http.MethodPost {
		http.Error(w,"Method not Allowed",http.StatusBadRequest)
		return
	}
	var registerUser dto.RegisterRequest
	err := json.NewDecoder(r.Body).Decode(&registerUser)
	savedUser,err:=services.CreateUser(ctx,registerUser)
	if err != nil {
		http.Error(w,"Bad Request",http.StatusBadRequest)
	}

	fmt.Println(savedUser)
	json.NewEncoder(w).Encode(savedUser)
}

func LoginUser(w http.ResponseWriter,r *http.Request){
	ctx:=r.Context()

	if r.Method != http.MethodPost {
		http.Error(w,"Method Not Allowed",http.StatusMethodNotAllowed)
		return
	}
	var loginRequest dto.LoginRequest

	token,err:=services.LoginUser(ctx,loginRequest)
	if err != nil {
		http.Error(w,"Bad Request",http.StatusBadRequest)
	}
	response :=dto.Loginresponse{
		Token: token,
	}
	json.NewEncoder(w).Encode(response)
}