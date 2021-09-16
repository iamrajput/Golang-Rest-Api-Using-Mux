package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//create new record
func createEmployee(w http.ResponseWriter, r *http.Request) {
	setHeaders(w)
	var emp Employee
	_ = json.NewDecoder(r.Body).Decode(&emp)
	fmt.Printf("%+v", &emp)
	DB.Create(&emp)
	sendResponse(w, emp, "Employee Created Successfully")
}

//Get all the employee
func getEmployees(w http.ResponseWriter, r *http.Request) {
	setHeaders(w)
	var employees []Employee
	DB.Find(&employees)
	sendResponse(w, employees, "Employee List")
}

func getEmployeeById(w http.ResponseWriter, r *http.Request) {
	setHeaders(w)
	var emp Employee
	result := DB.First(&emp, mux.Vars(r)["id"])
	if result.Error != nil {
		notFound(w)
		return
	}
	sendResponse(w, emp, "Employee details")
}

func updateEmployee(w http.ResponseWriter, r *http.Request) {
	setHeaders(w)
	var emp Employee
	result := DB.First(&emp, mux.Vars(r)["id"])
	if result.Error != nil {
		notFound(w)
		return
	}
	json.NewDecoder(r.Body).Decode(&emp)
	DB.Save(&emp)
	sendResponse(w, emp, "Employee Updated Successfully")
}

func removeEmployee(w http.ResponseWriter, r *http.Request) {
	setHeaders(w)
	var emp Employee
	result := DB.First(&emp, mux.Vars(r)["id"])
	if result.Error != nil {
		notFound(w)
		return
	}
	DB.Delete(&emp, mux.Vars(r)["id"])
	w.Write([]byte(`{"message": "Employee removed Successfully"}`))
}
