package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//create new record
func createEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var emp Employee
	_ = json.NewDecoder(r.Body).Decode(&emp)
	fmt.Printf("%+v", &emp)
	DB.Create(&emp)
	json.NewEncoder(w).Encode(emp)
	fmt.Println("DONE")
}

//Get all the employee
func getEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var employees []Employee
	DB.Find(&employees)
	json.NewEncoder(w).Encode(employees)
	w.WriteHeader(http.StatusOK)
	//w.Write([]byte(employees))
}

func getEmployeeById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var emp Employee
	DB.First(&emp, mux.Vars(r)["id"])
	json.NewEncoder(w).Encode(emp)
}

func updateEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var emp Employee
	DB.First(&emp, mux.Vars(r)["id"])
	json.NewDecoder(r.Body).Decode(&emp)
	DB.Save(&emp)
	json.NewEncoder(w).Encode(emp)
}
func removeEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var emp Employee
	DB.Delete(&emp, mux.Vars(r)["id"])
	json.NewEncoder(w).Encode("Employee removed")
}
