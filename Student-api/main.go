package main

import (
	"fmt"
	"net/http"
	"Student-api/routes"
)

func main(){
	routes.RegisterRoutes()
	fmt.Println("server is runing at port: 8080")
	
	http.ListenAndServe(":8080",nil)
	
	
}