package main

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

//set the headers
func setHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

//send the response
func sendResponse(w http.ResponseWriter, data interface{}, message string) {
	json.NewEncoder(w).Encode(Response{
		Data:    data,
		Message: message,
	})
}

func notFound(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"message": "Employee Not Found"}`))
}
