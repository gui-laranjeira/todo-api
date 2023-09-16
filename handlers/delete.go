package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/gui-laranjeira/todo-api/models"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Printf("Error while converting id to int: %v", err)
		return
	}

	rowsAffected, err := models.Delete(int64(id))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Printf("Error while removing todo: %v", err)
		return
	}

	if rowsAffected == 0 {
		http.Error(w, "No rows affected", http.StatusNotFound)
		log.Print("No rows affected")
		return
	}

	if rowsAffected > 1 {
		http.Error(w, "More than one row affected", http.StatusInternalServerError)
		log.Print("More than one row affected")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(fmt.Sprintf("Task with ID '%v' deleted!", id))

}
