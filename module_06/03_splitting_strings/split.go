package main

import (
	"fmt"
	"strings"
)

func main() {
	split_into_words()
	fmt.Println()
	split_into_sentences()
	fmt.Println()
	split_after()
	fmt.Println()
	split_after_n()
}

func split_into_words() {
	ourString := "This is a string"
	stringCollection := strings.Split(ourString, " ")

	for _, item := range stringCollection {
		fmt.Println(item)
	}
}

func split_into_sentences() {
	ourString := "This is a string|This is another one|I like turtles"
	stringCollection := strings.Split(ourString, "|")

	for _, item := range stringCollection {
		fmt.Println(item)
	}
}

func split_after() {
	ourString := "This is a string|This is another one|I like turtles"
	stringCollection := strings.SplitAfter(ourString, "|")

	for _, item := range stringCollection {
		fmt.Println(item)
	}
}

func split_after_n() {
	ourString := "This is a string"
	stringCollection := strings.SplitAfterN(ourString, " ", 2)

	for _, item := range stringCollection {
		fmt.Println(item)
	}
}
