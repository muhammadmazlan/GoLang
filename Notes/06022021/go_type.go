package main

import "fmt"

// func main() {
// 	name := "محمد"
// 	fmt.Println(name)
// 	fmt.Println([]byte(name))
// }

// func main() {
// 	a := "hello" //same as ` var a string = "hello" `
// 	fmt.Println(a)
// }

func main() {
	const (
		a float64 = 100.6766
		b float64 = 300.1234
	)
	fmt.Printf("%v\n", a)
	fmt.Printf("%.3f", b)
}
