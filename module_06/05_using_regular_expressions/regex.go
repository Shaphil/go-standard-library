package main

import (
	"fmt"
	"regexp"
)

func main() {
	sampleString := "This is a song. There are many strings like it, but this one is mine"
	// regex, _ := regexp.Compile(`s([a-z]+)g`)
	// regex, _ := regexp.Compile(`T([a-z]+)r`)
	regex, _ := regexp.Compile(`s(\w[a-z]+)g\b`)

	fmt.Println(regex.MatchString(sampleString))
	fmt.Println(regex.FindAllString(sampleString, -1))
	fmt.Println(regex.FindStringIndex(sampleString))

	newText := regex.ReplaceAllString(sampleString, "apple")
	fmt.Println(newText)
}
