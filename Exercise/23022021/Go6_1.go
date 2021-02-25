package main

import "fmt"

func main() {
	for i := 0; i <= 100; i++ {
		fmt.Println(i, isFizzBuzz(i))
	}
}

func isFizzBuzz(i int) string {
	switch {
	case (i%3 == 0 && i%5 == 0):
		return "fizzbuzz"
	case i%3 == 0:
		return "fizz"
	case i%5 == 0:
		return "buzz"
	default:
		return "no go"
	}
}
