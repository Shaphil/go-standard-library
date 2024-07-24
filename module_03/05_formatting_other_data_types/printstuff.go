package main

import "fmt"

// Learn more at - https://pkg.go.dev/fmt

type point struct {
	x, y int
}

type Person struct {
	firstName string
	lastName  string
	age       int
}

func main() {
	p := point{2, 3}
	fmt.Printf("%v\n", p)
	fmt.Println()

	person := Person{
		firstName: "Shaphil",
		lastName:  "Mahmud",
		age:       36,
	}
	fmt.Printf("%v\n", person)
	fmt.Println()

	// decimals, binaries, octals and hexadecimals
	number := 12345678
	fmt.Printf("%d\n", number) // decimal
	fmt.Printf("%b\n", number) // binary
	fmt.Printf("%o\n", number) // octal
	fmt.Printf("%x\n", number) // hexadecimal
	fmt.Printf("%X\n", number) // HEXADECIMAL
	fmt.Println()

	// characters
	charAsciiVal := 65
	fmt.Printf("%c", charAsciiVal)

}
