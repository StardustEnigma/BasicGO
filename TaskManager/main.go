package main

import(
	"fmt"
	"net/http"
	"TaskManager/routes"
)

func main(){
	routes.ResgisterRoutes()
	fmt.Println("Server is running at port : 8080")

	http.ListenAndServe(":8080",nil)
}