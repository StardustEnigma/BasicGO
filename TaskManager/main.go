package main

import (
	"TaskManager/routes"
	"TaskManager/workers"
	"fmt"
	"net/http"
)

func main(){
	router :=routes.RegisterRoutes()
	fmt.Println("Server is running at port : 8080")
	for i := 0; i <=3; i++ {
		go workers.Worker(i)
	}
	http.ListenAndServe(":8080",router)
}