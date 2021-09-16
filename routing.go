package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRiuting() {
	r := mux.NewRouter()
	r.HandleFunc("/api/create/employee", createEmployee).Methods("POST")
	r.HandleFunc("/api/employees", getEmployees).Methods("GET")
	r.HandleFunc("/api/employee/{id}", getEmployeeById).Methods("GET")
	r.HandleFunc("/api/update/employee/{id}", updateEmployee).Methods("PUT")
	r.HandleFunc("/api/remove/employee/{id}", removeEmployee).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", r))
}
