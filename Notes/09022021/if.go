// ==, !=, &&, ||

package main

import (
	"fmt"	"time"
)

/*
func main(){
	if i := 0; i % 2 == 0 {
		fmt.Println("condition Met")
	}else if i == 3{
		fmt.Println("not close")
	}
	else {
		fmt.Println("Not Met")
	}
}
*/

func main(){
	day := time.Now().Day()
	switch day{
	case time.Saturday:
		fmt.Println("Yay it's weekend")
		fallthrough
	default:
		fmt.Println("No it's weekday")
	}
}