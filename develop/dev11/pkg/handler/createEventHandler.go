package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"rest/pkg/models"
)

func (h *Handler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		log.Print("error method")
		http.Error(w, "error method", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	var event models.Event
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if !event.Valid() {
		log.Print("no valid event")
		http.Error(w, "no valid event", http.StatusInternalServerError)
		return
	}
	id, err := h.repository.CreateEvent(event)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	response, err := json.Marshal(map[string]int{
		"id": id,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
