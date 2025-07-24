package handlers

import (
	"encoding/json"
	"net/http"
)

func HealthHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		response := map[string]string{"message": "Server is Okay"}
		json.NewEncoder(w).Encode(response)
	}
}
