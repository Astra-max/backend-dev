package server


import (
	"net/http"
	"log"
	"encoding/json"
)


func Json(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Println("Failed to encode to json", err)
	}
}