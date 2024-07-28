package main

import (
	"fmt"
	"reflect"
)

type employee struct {
	personId  int
	firstName string
	lastName  string
}

func main() {
	newEmployee := employee{0, "Lloyd", "Christmas"}

	// name
	fmt.Printf("The name of the type is %v\n", reflect.TypeOf(newEmployee).Name())
	// type
	fmt.Printf("The type is %v\n", reflect.TypeOf(newEmployee))
	// kind
	fmt.Printf("The kind is %v\n", reflect.TypeOf(newEmployee).Kind())
	// value
	fmt.Printf("The value is %v\n", reflect.ValueOf(newEmployee))

	fmt.Println()

	employees := make([]employee, 3)
	employees = append(employees, employee{0, "Floyd", "Rose"})
	employees = append(employees, employee{1, "Lloyd", "Marshall"})
	employees = append(employees, employee{2, "Sea", "Horse"})

	// name
	fmt.Printf("The name of the type is %v\n", reflect.TypeOf(employees).Name())
	// type
	fmt.Printf("The type is %v\n", reflect.TypeOf(employees))
	// kind
	fmt.Printf("The kind is %v\n", reflect.TypeOf(employees).Kind())
	// value
	fmt.Printf("The value is %v\n", reflect.ValueOf(employees))
	fmt.Println()

	eType := reflect.TypeOf(employees)
	newEmployeeList := reflect.MakeSlice(eType, 0, 0)
	newEmployeeList = reflect.Append(
		newEmployeeList,
		reflect.ValueOf(employee{3, "Mary", "Watson"}),
	)

	fmt.Printf("First list of employees: %v\n", employees)
	fmt.Printf("List created by reflection: %v\n", newEmployeeList)
}
