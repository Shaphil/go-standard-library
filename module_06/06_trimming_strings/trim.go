package main

import (
	"fmt"
	"strings"
)

func main() {
	sampleString := "    This is our text    "
	fmt.Printf("Example (sampleString):\t\t\t %q\n", sampleString)

	newString := strings.TrimSpace(sampleString)
	fmt.Printf("Example `TrimSpace(sampleString)`:\t %q\n", newString)

	newString = strings.TrimLeft(sampleString, " ")
	fmt.Printf("Example `TrimLeft(sampleString)`:\t %q\n", newString)

	newString = strings.TrimRight(sampleString, " ")
	fmt.Printf("Example `TrimRight(sampleString)`:\t %q\n", newString)

	url := "https://shaphil.me"
	domainName := strings.TrimPrefix(url, "https://")
	fmt.Println("Example `TrimPrefix(url)`:\t\t", domainName)

	fileName := "important.doc"
	newFileName := strings.TrimSuffix(fileName, ".doc")
	fmt.Println("Example `TrimSuffix(fileName)`:\t\t", newFileName)
}
