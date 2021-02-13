package main

import "fmt"

func main() {
	isEvenIf(9)
	isEvenIf(2)
	isEventSwitch(4)
	isEventSwitch(7)
}

func isEvenIf(n int) {
	if n%2 == 0 {
		fmt.Println(n, "is Even number")
	} else {
		fmt.Println(n, "is Odd number")
	}
}

func isEventSwitch(n int) {
	switch {
	case n%2 == 0:
		fmt.Println(n, "is Even number")
		break
	default:
		fmt.Println(n, "is Odd number")
	}
}
