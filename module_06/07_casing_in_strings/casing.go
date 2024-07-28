package main

import (
	"fmt"
	"strings"

	utils "shaphil.me/go-standard-library"
)

func main() {
	sampleString := "Never trust a programmer who carries a screwdriver"
	fmt.Println("Before:", sampleString)

	lower := strings.ToLower(sampleString)
	fmt.Println("lower case:", lower)

	upper := strings.ToUpper(sampleString)
	fmt.Println("UPPER CASE:", upper)

	titleCase := utils.ToTitle(sampleString)
	fmt.Println("Title Case:", titleCase)

	fmt.Println("Proper Title Case:", properTitle(sampleString))
}

func properTitle(input string) string {
	words := strings.Fields(input)
	smallwords := " a an on the to "

	for index, word := range words {
		if strings.Contains(smallwords, " "+word+" ") {
			words[index] = word
		} else {
			words[index] = strings.Title(word)
		}
	}

	return strings.Join(words, " ")
}
