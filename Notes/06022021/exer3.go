package main

import "fmt"

func main() {
	fmt.Println(square(2))
	fmt.Println(square(5))
	fmt.Println(square(10))
	fmt.Println(square(20))
}

func square(a int) int {
	return (a * a)
}
