package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("What is your name? ")
	name, _ := reader.ReadString('\n')
	fmt.Printf("Hello %v", name)

	fmt.Printf("We are using Go %v running in %v\n", runtime.Version(), runtime.GOOS)
}
