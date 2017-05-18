package main

import (
	"fmt"
	"time"
)

func main() {
	date := time.Date(2017, 1, 20, 00, 00, 00, 00, time.Local)
	fmt.Printf("Date is: %s\n", date)

	now := time.Now()
	fmt.Printf("Time is: %s\n", now)

	fmt.Println("Date is:", now.Day())
	fmt.Println("Day is:", now.Weekday())
	fmt.Println("Month is:", now.Month())
	fmt.Println("Year is:", now.Year())

	tomorrow := now.AddDate(0, 0, 1)
	fmt.Printf("Tomorrow is: %v, %v %v %v\n", tomorrow.Weekday(), tomorrow.Day(), tomorrow.Month(), tomorrow.Year())

	/*
		Format is always specified in terms of a fixed mask -:
		Weekday -: Monday
		Month 	-: 1 / January
		Day   	-: 2
		Hours	-: 3
		Minutes	-: 4
		Seconds	-: 5
		Year	-: 6
		Time Zone -: -7
	*/

	dateFormat := "Monday, January 2, 2006"
	fmt.Println("Today is:", now.Format(dateFormat))

	shortDateFormat := "02/01/2006"
	fmt.Println("Today is:", now.Format(shortDateFormat))
}
