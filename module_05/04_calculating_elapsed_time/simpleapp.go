package main

import (
	"bufio"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()

	args := os.Args
	// open the customer list
	customers, err := os.Open(string(args[1]))
	check(err)
	defer customers.Close()
	writeTime(start, "Opened Customer List")

	// create an output file
	outfile, err := os.Create("outfile.csv")
	check(err)
	defer outfile.Close()
	writeTime(start, "Created output file")

	// calculate time elaspsed for processing the file
	// elapsed := time.Since(start)
	// fmt.Println(elapsed)

	// scan the customer list
	scanner := bufio.NewScanner(customers)
	for scanner.Scan() {
		names := strings.Split(scanner.Text(), ",")
		outfile.WriteString((names[1] + "," + names[2] + "\n"))
	}
	check(scanner.Err())
	writeTime(start, "Wrote data to the output file")
	defer writeTime(start, "Application Finished Execution")
}

func writeTime(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
