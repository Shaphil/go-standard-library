package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// simpleSearch()
	// fmt.Println()
	// searchContains()
	// fmt.Println()
	// searchContainsPrefix()
	// fmt.Println()
	searchContainsSuffix()
}

func searchContainsSuffix() {
	sampleString := "I really enjoy the Go language. It's one of my favorites"

	if len(os.Args) > 1 {
		searchTerm := os.Args[1]
		result := strings.HasSuffix(sampleString, searchTerm)

		if result {
			fmt.Printf("The sample string ends with `%s`\n", searchTerm)
		} else {
			fmt.Printf("The sample string does NOT end with `%s`\n", searchTerm)
		}
	} else {
		fmt.Println("You must enter a search term")
	}
}

func searchContainsPrefix() {
	sampleString := "I really enjoy the Go language. It's one of my favorites"

	if len(os.Args) > 1 {
		searchTerm := os.Args[1]
		result := strings.HasPrefix(sampleString, searchTerm)

		if result {
			fmt.Printf("The sample string starts with `%s`\n", searchTerm)
		} else {
			fmt.Printf("The sample string does NOT start with `%s`\n", searchTerm)
		}
	} else {
		fmt.Println("You must enter a search term")
	}
}

func searchContains() {
	sampleString := "I really enjoy the Go language. It's one of my favorites"

	if len(os.Args) > 1 {
		searchTerm := os.Args[1]
		result := strings.Contains(sampleString, searchTerm)

		if result {
			fmt.Printf("The sample string contains `%s`\n", searchTerm)
		} else {
			fmt.Printf("The sample string does NOT contain `%s`\n", searchTerm)
		}
	} else {
		fmt.Println("You must enter a search term")
	}
}

func simpleSearch() {
	sampleString := "I really enjoy the Go language. It's one of my favorites"
	searchTerm := "Go"

	result := strings.Contains(sampleString, searchTerm)
	fmt.Printf("The sample text includes %s: %t\n", searchTerm, result)
}
