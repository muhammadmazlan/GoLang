package main

import "fmt"

func main() {
	var sum int
	for i := 0; i <= 100; i++ {
		sum += i // sum = sum + i
	}
	fmt.Println("SUM", sum)
}
