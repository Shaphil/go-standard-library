package main

import "fmt"

func main() {
	printf_demo()
	fmt.Println("-----------------------------------------------")
	format_specifiers()
	fmt.Println("-----------------------------------------------")
	formatting_output()
}

func printf_demo() {
	var age = 42
	var out, _ = fmt.Print("Jeremy is ", age, " years old\n")
	print("Bytes written: ", out, "\n")

	var name = "Jeremy"
	// `%v` tells `printf` to infer the value from the variable
	fmt.Printf("Name is %v and age is %v years\n", name, age)
	// using the correct format specifier for the corresponding types
	fmt.Printf("Name is %s and age is %d years\n", name, age)
}

func format_specifiers() {
	var name string = "Jeremy"
	var age int = 10

	/*
		%v - formats the value in a default format
		%s - formats string values
		%d - formats decimal integers
		%g/%f - formats the floating-point numbers
		%b - formats base 22 numbers
		%o - formats basr 88 numbers
		%t - formats booleans
	*/

	fmt.Printf("My name is %s and I am %d years old\n", name, age)
}

func formatting_output() {
	var pi float32 = 3.1415926535897932
	fmt.Printf("Pi: %g\n", pi)
	fmt.Println()

	fmt.Println("Formatting numeric tabular data (right justified)")
	fmt.Printf("|%7.2f|%7.2f|%7.2f|\n", 23.3774, 577.45, 1234.56)
	fmt.Printf("|%7.2f|%7.2f|%7.2f|\n", 98.999, 12.3456, 12.01)
	fmt.Println()

	fmt.Println("Formatting numeric tabular data (left justified)")
	fmt.Printf("|%-7.2f|%-7.2f|%-7.2f|\n", 23.3774, 577.45, 1234.56)
	fmt.Printf("|%-7.2f|%-7.2f|%-7.2f|\n", 98.999, 12.3456, 12.01)
	fmt.Println()

	fmt.Println("Formatting string tabular data (left justified)")
	fmt.Printf("|%-7s|%-7s|%-7s|\n", "foo", "bar", "go")
	fmt.Printf("|%-7s|%-7s|%-7q|\n", "a", "ab", "abc") // `%q` is used for "quoted" string
	fmt.Println()

	// `fmt.Sprintf()` will return the formatted string
	output := fmt.Sprintf("|%-7s|%-7s|%-7s|\n", "foo", "bar", "go")
	print(output)
}
