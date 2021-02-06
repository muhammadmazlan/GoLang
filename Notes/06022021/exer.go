package main

import "fmt"

func main() {
	const (
		a string  = "Hello"
		b int     = 102
		c bool    = true
		d float64 = 1034.67890
	)
	fmt.Printf("Type = %T"+", "+"value = %s\n", a, a)   // to print string format
	fmt.Printf("Type = %T"+", "+"value = %d\n", b, b)   // to print integer format
	fmt.Printf("Type = %T"+", "+"value = %t\n", c, c)   // to print boolean format
	fmt.Printf("Type = %T"+", "+"value = %.3f\n", d, d) // to print round to 3 decimal format
	/*
		fmt.Printf("%s\n", a)    // to print string format
		fmt.Printf("%i"+"\n", b) // to print integer format
		fmt.Printf("%b"+"\n", c) // to print boolean format
		fmt.Printf("%.3f", d)    // to print round to 3 decimal format
	*/
}
