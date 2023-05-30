package main

import "fmt"

func minutesToHours(minutes int) float64 {
	hours := float64(minutes) / 60
	roundedHours := float64(int(hours*2)) / 2

	return roundedHours
}

func main() {
	minutes := 60
	hours := minutesToHours(minutes)
	fmt.Printf("%d minutes is equal to %.1f hour(s)\n", minutes, hours)

	minutes = 75
	hours = minutesToHours(minutes)
	fmt.Printf("%d minutes is equal to %.1f hour(s)\n", minutes, hours)

	minutes = 120
	hours = minutesToHours(minutes)
	fmt.Printf("%d minutes is equal to %.1f hour(s)\n", minutes, hours)

	minutes = 179
	hours = minutesToHours(minutes)
	fmt.Printf("%d minutes is equal to %.1f hour(s)\n", minutes, hours)

	minutes = 180
	hours = minutesToHours(minutes)
	fmt.Printf("%d minutes is equal to %.1f hour(s)\n", minutes, hours)
}
