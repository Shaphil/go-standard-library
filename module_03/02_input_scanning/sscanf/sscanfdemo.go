package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("family.csv")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var name string
		var age int

		n, err := fmt.Sscanf(scanner.Text(), "%s is %d years old", &name, &age)
		if err != nil {
			panic(err)
		}
		if n == 2 {
			fmt.Printf("%s, %d\n", name, age)
		}
	}
}
