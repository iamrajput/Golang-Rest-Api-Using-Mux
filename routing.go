package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRiuting() {
	r := mux.NewRouter()
	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/create/employee", createEmployee).Methods(http.MethodPost)
	api.HandleFunc("/employees", getEmployees).Methods(http.MethodGet)
	api.HandleFunc("/employee/{id}", getEmployeeById).Methods(http.MethodGet)
	api.HandleFunc("/update/employee/{id}", updateEmployee).Methods(http.MethodPut)
	api.HandleFunc("/remove/employee/{id}", removeEmployee).Methods(http.MethodDelete)
	log.Fatal(http.ListenAndServe(":8080", r))
}
