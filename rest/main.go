package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Employee struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Company string `json:"company"`
	Health  string `json:"health"`
}

var employees = []Employee{
	{ID: 1, Name: "Rohit", Age: 30, Company: "Lupin", Health: "Fit"},
	{ID: 2, Name: "Aarti", Age: 28, Company: "JK Tyres", Health: "Needs checkup"},
}

// GET /employees
func getEmployees(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employees)
}

// POST /employees
func addEmployee(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var newEmp Employee
	err := json.NewDecoder(r.Body).Decode(&newEmp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newEmp.ID = len(employees) + 1
	employees = append(employees, newEmp)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newEmp)
}

func main() {
	
	http.HandleFunc("/employees", getEmployees)    // GET
	http.HandleFunc("/employees/add", addEmployee) // POST

	fmt.Println("Server started on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
