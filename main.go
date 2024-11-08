package main

import "fmt"

func isLeapYear(year int) bool {
	if year % 400 == 0 {
		return true
	} else if year % 4 == 0 && year % 100 != 0 {
		return true
	}
	return false
}

func main() {
	var year int
	fmt.Scan(&year)
	fmt.Println(isLeapYear(year))
}