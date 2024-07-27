package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) > 1 {
		searchTerm := os.Args[1]

		file, _ := os.Open("log.txt")
		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			line := scanner.Text()
			result := strings.HasPrefix(line, searchTerm)
			if result {
				fmt.Println(line)
			}
		}
	} else {
		fmt.Println("You must enter a search term")
	}
}

// to redirect errors into their own log file,
// `go run logparser.go ERR > errors.txt`
