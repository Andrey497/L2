package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

func (h *Handler) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		log.Print("error method")
		http.Error(w, "error method", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	var id int
	err := json.NewDecoder(r.Body).Decode(&id)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = h.repository.DeleteEvent(id)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Ok"))
}
