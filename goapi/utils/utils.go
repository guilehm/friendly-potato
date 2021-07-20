package utils

import (
	"encoding/json"
	"net/http"
	"os"
)

var SECRET_KEY = os.Getenv("JWT_SECRET_KEY")

func HandleApiErrors(w http.ResponseWriter, status int, message string) {
	if message == "" {
		message = http.StatusText(status)
	}

	jsonResponse, _ := json.Marshal(struct {
		Error string `json:"error"`
	}{message})
	w.WriteHeader(status)
	w.Write(jsonResponse)
}
