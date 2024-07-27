package main

import "fmt"

func main() {
	ourString := "\x47\x6f\x20\x69\x73\x20\x41\x77\x65\x73\x6f\x6d\x65\x21"
	fmt.Println(ourString)

	// print all the hexadeicmal values
	for i := 0; i < len(ourString); i++ {
		fmt.Printf("%x ", ourString[i])
	}
	fmt.Println()

	// print all the items as quoted byte sequence
	for i := 0; i < len(ourString); i++ {
		fmt.Printf("%q ", ourString[i])
	}
	fmt.Println()

	goBlogString := "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"
	fmt.Println(goBlogString)

	for i := 0; i < len(goBlogString); i++ {
		fmt.Printf("%q\n", goBlogString[i])
	}
	fmt.Println()

	newString := "This is a string!"
	bytes := []byte(newString)
	fmt.Println(newString[3], newString[3:4], bytes[3:4])
	fmt.Println(newString[2:3], bytes[2:3])
	fmt.Println(newString[0:4])
}
