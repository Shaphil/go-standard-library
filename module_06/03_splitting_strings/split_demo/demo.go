package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, _ := os.Open("customerlist.csv")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		items := strings.Split(line, ",")

		fmt.Println("-- new record ---")
		for _, item := range items {
			fmt.Println(item)
		}
		fmt.Println()
	}
}
