package auth

import (
	"net/http"
)

func AllAcces(w http.ResponseWriter, r *http.Request) {

	// Cors Header HTTP

	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Headers", "Authorization")

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"Hola a un punto publico! Necesitas acceder para ver esto."}`))
}
