package main

import (
	"fmt"
	"reflect"
)

func main() {
	// res := add(1, 2)
	// fmt.Println(res)

	// num1, num2 := nums(1, 2)
	// fmt.Println(num1, num2)

	// firstSeq := nextSeq()
	// fmt.Println(firstSeq())
	// fmt.Println(firstSeq())
	// fmt.Println(firstSeq())

	// var i int = 5
	// var f float64 = float64(i)   // type casting
	// fmt.Println(i, f)

	// var i int = 5
	// var j int = 10
	// var avg float64 = float64(i) / float64(j)
	// fmt.Println(avg)

	var i = "hello"
	var j = 3
	var k = 2.5
	fmt.Println(reflect.TypeOf(i))
	fmt.Println(reflect.TypeOf(j))
	fmt.Println(reflect.TypeOf(k))
}

// // func add(a int, b int) int {
// // 	return a + b
// // }

// func nums(a, b int) (c, d int) {
// 	return a, b
// }
//closure
// func nums() func() int { //anonymous func
// 	return func() int {
// 		return 1
// 	}
// }

// func nextSeq() func() int {
// 	i := 0
// 	return func() int {
// 		i++
// 		return i
// 	}
// }
