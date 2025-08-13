package auth

import (
	"encoding/json"
	"net/http"
)

func Salom(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := map[string]string{
		"message": "salom",
	}

	json.NewEncoder(w).Encode(response)
}