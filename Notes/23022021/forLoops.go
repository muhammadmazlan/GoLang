package main

import "fmt"

func main() {
	// inifinite loops need to have exits mechanism
	// i := 0
	// for i < 3 {
	// 	fmt.Println("hello")
	// 	i++
	// }

	// for i := 0; i < 3; i++ {
	// 	fmt.Println("Hello2")
	// }

	// for i := 0; i < 20; i++ {
	// 	if i == 6{
	// 		break					//for exit
	// 	}
	// 	fmt.Println("Test")
	// }

	// for i := 0; i < 10; i++ {
	// 	if i == 6 {
	// 		continue				//for skip
	// 	}
	// 	fmt.Println("Level")
	// }

	for i := 0; i < 5; i++ {
		fmt.Println("i", i)
		for j := 0; j < 3; j++ {
			fmt.Println("j", j)
		}
	}
}
