package main

import (
	"fmt"

	"shaphil.me/go-standard-library/module_07/02_creating_custom_types/media"
)

func main() {
	fmt.Println("My Favorite Movie")

	var myFav = &media.Movie{}
	myFav.NewMovie("Top Gun", media.PG, 32.8)

	fmt.Printf("My favorite movie is %s\n", myFav.GetTitle())
	fmt.Printf("It was rated %v\n", myFav.GetRating())
	fmt.Printf("It made %.1f in the box office\n", myFav.GetBoxOffice())

	myFav.SetTitle("Dumb and Dumber")
	fmt.Printf("My favorite movie is now %s\n", myFav.GetTitle())
}
