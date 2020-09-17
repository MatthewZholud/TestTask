package Handlers

import (
	"net/http"
)

const message = "Hi, Matthew"


func GetGroups(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(message))
	w.WriteHeader(http.StatusOK)
}
