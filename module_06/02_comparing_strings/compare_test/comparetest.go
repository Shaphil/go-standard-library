package main

import (
	"fmt"
	"log"
	"strings"
	"time"
)

func main() {
	string1 := `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Proin interdum tincidunt lacus, 
	sed tempor lorem consectetur et. Pellentesque et egestas sem, at cursus massa. Nunc feugiat elit sit 
	amet ipsum commodo luctus. Proin auctor dignissim pharetra. Integer iaculis quam a tellus auctor, 
	vitae auctor nisl viverra. Nullam consequat maximus porttitor. Pellentesque tortor enim, molestie 
	at varius non, tempor non nibh. Suspendisse tempus erat lorem, vel faucibus magna blandit vel. Sed 
	pellentesque ligula augue, vitae fermentum eros blandit et. Cras dignissim in massa ut varius. 
	Vestibulum commodo nunc sit amet pellentesque dignissim.

	Donec imperdiet blandit lobortis. Suspendisse fringilla nunc quis venenatis tempor. Nunc tempor sed 
	erat sed convallis. Pellentesque aliquet elit lectus, quis vulputate arcu pharetra sed. Etiam laoreet 
	aliquet arcu cursus vehicula. Maecenas odio odio, elementum faucibus sollicitudin vitae, pellentesque 
	ac purus. Donec venenatis faucibus lorem, et finibus lacus tincidunt vitae. Quisque laoreet metus 
	sapien, vitae euismod mauris lobortis malesuada. Integer sit amet elementum turpis. Maecenas ex 
	mauris, dapibus eu placerat vitae, rutrum convallis enim. Nulla vitae orci ultricies, sagittis turpis 
	et, lacinia dui. Praesent egestas urna turpis, sit amet feugiat mauris tristique eu. Quisque id tempor 
	libero. Donec ullamcorper dapibus lorem, vel consequat risus congue a.

	Nullam dignissim ut lectus vitae tempor. Pellentesque ut odio fringilla, volutpat mi et, vulputate tellus. 
	Fusce eget diam non odio tincidunt viverra eu vel augue. Vestibulum ante ipsum primis in faucibus orci 
	luctus et ultrices posuere cubilia curae; Nullam sed eleifend purus, vitae aliquam orci. Cras fringilla 
	justo eget tempus bibendum. Phasellus venenatis, odio nec pulvinar commodo, quam neque lacinia turpis, 
	ut rutrum tortor massa eu nulla. Vivamus tincidunt ut lectus a gravida. Donec varius mi quis enim interdum 
	ultrices. Sed aliquam consectetur nisi vitae viverra. Praesent nec ligula egestas, porta lectus sed, 
	consectetur augue.`

	string2 := `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Proin interdum tincidunt lacus, 
	sed tempor lorem consectetur et. Pellentesque et egestas sem, at cursus massa. Nunc feugiat elit sit 
	amet ipsum commodo luctus. Proin auctor dignissim pharetra. Integer iaculis quam a tellus auctor, 
	vitae auctor nisl viverra. Nullam consequat maximus porttitor. Pellentesque tortor enim, molestie 
	at varius non, tempor non nibh. Suspendisse tempus erat lorem, vel faucibus magna blandit vel. Sed 
	pellentesque ligula augue, vitae fermentum eros blandit et. Cras dignissim in massa ut varius. 
	Vestibulum commodo nunc sit amet pellentesque dignissim.

	Donec imperdiet blandit lobortis. Suspendisse fringilla nunc quis venenatis tempor. Nunc tempor sed 
	erat sed convallis. Pellentesque aliquet elit lectus, quis vulputate arcu pharetra sed. Etiam laoreet 
	aliquet arcu cursus vehicula. Maecenas odio odio, elementum faucibus sollicitudin vitae, pellentesque 
	ac purus. Donec venenatis faucibus lorem, et finibus lacus tincidunt vitae. Quisque laoreet metus 
	sapien, vitae euismod mauris lobortis malesuada. Integer sit amet elementum turpis. Maecenas ex 
	mauris, dapibus eu placerat vitae, rutrum convallis enim. Nulla vitae orci ultricies, sagittis turpis 
	et, lacinia dui. Praesent egestas urna turpis, sit amet feugiat mauris tristique eu. Quisque id tempor 
	libero. Donec ullamcorper dapibus lorem, vel consequat risus congue b.

	Nullam dignissim ut lectus vitae tempor. Pellentesque ut odio fringilla, volutpat mi et, vulputate tellus. 
	Fusce eget diam non odio tincidunt viverra eu vel augue. Vestibulum ante ipsum primis in faucibus orci 
	luctus et ultrices posuere cubilia curae; Nullam sed eleifend purus, vitae aliquam orci. Cras fringilla 
	justo eget tempus bibendum. Phasellus venenatis, odio nec pulvinar commodo, quam neque lacinia turpis, 
	ut rutrum tortor massa eu nulla. Vivamus tincidunt ut lectus a gravida. Donec varius mi quis enim interdum 
	ultrices. Sed aliquam consectetur nisi vitae viverra. Praesent nec ligula egestas, porta lectus sed, 
	consectetur augue.`

	// basic compare
	start1 := time.Now()
	if string1 == string2 {
		fmt.Println("Strings are equal")
	} else {
		fmt.Println("Strings are not equal")
	}
	timeTrack(start1, "Basic Compare")

	// using strings compare method
	start2 := time.Now()
	if strings.Compare(string1, string2) == 0 {
		fmt.Println("Strings are equal")
	} else {
		fmt.Println("Strings are not equal")
	}
	timeTrack(start2, "strings.Compare()")
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %d µs", name, elapsed.Microseconds())
}
