package main

import "fmt"

func main() {
	// arr := [4]int{1, 2, 3, 4}
	// for idx, val := range arr {
	// 	fmt.Println(idx)
	// 	fmt.Println(val)
	// }

	// arr := [4]int{1, 2, 3, 4}
	// for idx := range arr {
	// 	fmt.Println(idx)
	// }

	// myObj := map[string]string{"a": "test", "b": "hello"}
	// // for k, v := range myObj {
	// for _, v := range myObj {
	// 	// fmt.Println(k)
	// 	fmt.Println(v)
	// }

	myObj := map[string]string{"a": "test", "b": "hello"}
	for v := range myObj {
		fmt.Println(v)
	}
}
