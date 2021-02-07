package main

import "fmt"

func main() {
	fmt.Println(sqRoot(25))
}

func sqRoot(a int) float64 {
	return float64(a ^ 2)
}
