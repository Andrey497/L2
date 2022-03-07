package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func (h *Handler) GetEventForDay(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		log.Print("error method")
		http.Error(w, "error method", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	now := time.Now().Format("2006-01-02")
	lastDate := time.Now().Format("2006-01-02")
	events, err := h.repository.GetEvent(lastDate, now)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	response, err := json.Marshal(events)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
