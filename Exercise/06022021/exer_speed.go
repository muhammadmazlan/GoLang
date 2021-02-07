package main

import "fmt"

func main() {
	var distanceA, speedA float64 = 165, 3.5
	var distanceB, speedB float64 = 60, 2

	fmt.Printf("The car #1 speed is %.2f km/h \n", speed(distanceA, speedA))
	fmt.Printf("The car #2 speed is %.2f km/h \n", speed(distanceB, speedB))
}

func speed(distance, speed float64) float64 {
	return distance / speed
}
