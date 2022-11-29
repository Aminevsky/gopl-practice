package main

import (
	"time"
)

type Employee struct {
	ID       int
	Name     string
	Address  string
	DoB      time.Time
	Position string
	Salary   int
	ManageID int
}

func main() {
	var dilbert Employee

	id := dilbert.ID
	EmployeeByID(id).Salary = 0
}

func EmployeeByID(id int) Employee {
	return Employee{
		ID: 1,
		Name: "Dilbert",
		Address: "Main Street",
		DoB: time.Now(),
		Position: "Senior Engineer",
		Salary: 10000,
		ManageID: 10,
	}
}
