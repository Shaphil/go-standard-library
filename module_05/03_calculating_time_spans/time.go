package main

import (
	"fmt"
	"time"
)

func main() {
	first := time.Now()
	fmt.Printf("It is currently\t %v\n", first.Format("15:04:05"))
	time.Sleep(time.Second)

	second := time.Now()
	fmt.Printf("It is now\t %v\n", second.Format("15:04:05"))
	fmt.Println()

	// Reference time
	// Mon Jan 2 15:04:05 MST 2006
	today := time.Now()
	dateFormat := "Monday, Jan 2 2006"
	fmt.Printf("It is currently %v\n", today.Format(dateFormat))
	fmt.Println()

	startDate := time.Date(2018, 07, 04, 9, 00, 00, 0, time.UTC)
	elapsed := time.Since(startDate)
	fmt.Printf("%v has passed since %v\n", elapsed, startDate.Format(dateFormat))
	fmt.Printf("Hours: %.0f, Minutes: %.0f, Seconds: %.0f\n", elapsed.Hours(), elapsed.Minutes(), elapsed.Seconds())
	fmt.Println()

	future := today.AddDate(0, 6, 0)
	fmt.Printf("In six months it will be %v\n", future.Format((dateFormat)))

	past := today.AddDate(0, -6, 0)
	fmt.Printf("Six months ago it was\t %v\n", past.Format((dateFormat)))
	fmt.Println()

	nextHour := today.Add(6 * time.Hour)
	fmt.Printf("In six hours it will be %v\n", nextHour.Format("15:04"))
	fmt.Println()

	bedtime := time.Date(2024, 7, 27, 23, 0, 0, 0, time.Local)
	fmt.Printf("There is %0.f hours left until bed time\n", time.Until(bedtime).Hours())
}
