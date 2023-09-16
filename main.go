package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/gui-laranjeira/todo-api/configs"
	"github.com/gui-laranjeira/todo-api/handlers"
)

func main() {
	err := configs.Load()
	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()
	r.Post("/", handlers.Create)
	r.Put("/{id}", handlers.Update)
	r.Get("/{id}", handlers.Get)
	r.Get("/", handlers.List)
	r.Delete("/{id}", handlers.Delete)
	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetPort()), r)
	fmt.Println("Server running on port", configs.GetPort())
}
