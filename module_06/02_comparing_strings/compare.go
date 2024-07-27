package main

import (
	"fmt"
	"strings"
)

func main() {
	string1 := "this is a string"
	string2 := "This is a string"

	if string1 == string2 {
		fmt.Println("The strings are identical")
	} else {
		fmt.Println("The strings do not match")
	}

	if strings.Compare(string1, string2) == 0 {
		fmt.Println("The strings are identical")
	} else {
		fmt.Println("The strings do not match")
	}
	fmt.Println()

	stooges := []string{"Larry", "Curly", "Moe"}
	for _, stooge := range stooges {
		fmt.Println(strings.Compare("Larry", stooge))
	}
}
