package main

import (
	"fmt"
	"reflect"
)

func main() {
	// type person struct {
	// 	id        int
	// 	firstName string
	// 	lastName  string
	// }

	// newPerson := person{0, "Fred", "Flintstone"}
	// fmt.Printf("Our personis %s %s with and id of %d\n",
	// 	newPerson.firstName, newPerson.lastName, newPerson.id)

	// fmt.Printf("The type of `newPerson` is %v\n", reflect.TypeOf(newPerson))
	// fmt.Printf("The value of `newPerson` is %v\n", reflect.ValueOf(newPerson))
	// // fmt.Printf("The value of `newPerson` is %v\n", newPerson)	// same as above!!!
	// fmt.Printf("The kind of type it is: %v\n", reflect.ValueOf(newPerson).Kind())

	type employee struct {
		personId  int
		firstName string
		lastName  string
	}

	type customer struct {
		customerId int
		firstName  string
		lastName   string
		company    string
	}

	newEmployee := employee{0, "Fred", "Flintstone"}
	newCustomer := customer{234, "Barney", "Rubble", "Slate Rock and Gravel"}

	addPerson(newEmployee)
	addPerson(newCustomer)
}

func addPerson(p interface{}) bool {
	if reflect.ValueOf(p).Kind() == reflect.Struct {
		v := reflect.ValueOf(p)
		switch reflect.TypeOf(p).Name() {
		case "employee":
			empSqlString := "INSERT INTO employees (personId, firstName, lastName) VALUES (?, ?, ?)"
			fmt.Printf("SQL: %s\n", empSqlString)
			fmt.Printf("Added %v\n", v.Field(1))
		case "customer":
			customerSqlString := "INSERT INTO customers (customerId, firstName, lastName, company) VALUES (?, ?, ?, ?)"
			fmt.Printf("SQL: %s\n", customerSqlString)
			fmt.Printf("Added %v\n", v.Field(1))
		}

		return true
	} else {
		return false
	}
}
