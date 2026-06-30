package routes

import (
	"TaskManager/handlers"
	middlerware "TaskManager/middleware"

	"github.com/go-chi/chi/v5"
)

func RegisterRoutes() *chi.Mux{
	 r := chi.NewRouter()
	 r.Use(middlerware.Logger)

	r.Post("/register",handlers.RegisterUser)
	r.Post("/login",handlers.LoginUser)
	r.Group(
		func(r chi.Router) {
			r.Use(middlerware.AuthMiddleware)
			r.Post("/tasks",handlers.CreateTask)
			r.Get("/tasks",handlers.PrintAllTasks)
			r.Get("/tasks/{id}",handlers.GetTask)
			r.Delete("/tasks/{id}",handlers.DeleteTask)
			r.Put("/tasks/{id}",handlers.UpdateTask)
			r.Post("/tasks/{id}/complete",handlers.MarkCompleted)
			r.Post("/tasks/{id}/start",handlers.StartTask)
			r.Get("/tasks",handlers.GetTasksbyStatus)
		},
	)
	return r

}
