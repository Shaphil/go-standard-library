package main

import (
	"fmt"
	"strings"
)

func main() {
	string1 := "i like turtles"
	string2 := "I like turtles"

	fmt.Println(CompareCaseIns(string1, string2))
}

func CompareCaseIns(a, b string) bool {
	if len(a) == len(b) {
		if strings.EqualFold(a, b) {
			return true
		}
	}
	return false
}
