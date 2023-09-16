package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gui-laranjeira/todo-api/models"
)

func Create(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo

	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Printf("Error while decoding json: %v", err)
		return
	}

	id, err := models.Insert(todo)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Printf("Error while inserting todo: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(fmt.Sprintf("Task with ID '%v' created!", id))
}
