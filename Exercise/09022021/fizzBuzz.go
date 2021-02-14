package main

import "fmt"

func main() {
	a, b, c, d := -3, 5, 15, 26

	fmt.Println(isFizzBuzz(a))
	fmt.Println(isFizzBuzz(b))
	fmt.Println(isFizzBuzz(c))
	fmt.Printf("%s", isFizzBuzz(d))

}

func isFizzBuzz(i int) string {
	switch {
	case i%3 == 0:
		return "fizz"
	case i%5 == 0:
		return "buzz"
	case (i%3 == 0 && i%5 == 0):
		return "fizzbuzz"
	default:
		return "no go"
	}
}
