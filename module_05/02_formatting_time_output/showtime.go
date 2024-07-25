package main

import (
	"fmt"
	"time"
)

const (
	layoutISO = "2006-01-02"
	layoutUS  = "January 2, 2006"
	layoutEU  = "2 January, 2006"
)

func main() {
	t := time.Now()

	Year := t.Year()
	Month := t.Month()
	Day := t.Day()

	fmt.Printf("Today is %d/%d/%d\n\n", Month, Day, Year)

	// date/time printed as ANSII C format
	fmt.Println("ANSII C time:", t.Format(time.ANSIC))
	// fmt.Println()

	// date/time printed as RFC3339 format
	fmt.Println("RFC3339 time:", t.Format(time.RFC3339))
	// fmt.Println()

	// date/time printed as UNIX date
	fmt.Println("Unix date/time:", t.Format(time.UnixDate))
	// fmt.Println()

	// time printed as Kitchen format
	fmt.Println("Kitchen time:", t.Format(time.Kitchen))
	fmt.Println()

	// custom time format
	// Reference time: `Mon Jan 2 15:04:05 MST 2006`
	format := "Mon Jan 2 15:04:05 MST 2006"
	fmt.Println(t.Format(format))
	fmt.Println(t.Format("Monday, January 2 in the year 2006"))
	fmt.Println()

	startDate := time.Date(2024, 07, 25, 21, 00, 00, 00, time.UTC)
	fmt.Println(startDate)

	longFormat := "Monday, January 2 2006 at 15:04:05"
	fmt.Println(startDate.Format(longFormat))
	fmt.Println()

	fmt.Println(startDate.Format(layoutISO))
	fmt.Println(startDate.Format(layoutEU))
	fmt.Println(startDate.Format(layoutUS))
}
