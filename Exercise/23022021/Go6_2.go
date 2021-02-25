package main

import "fmt"

func main() {
	var mux int
	for i := 0; i <= 50; i++ {
		mux = i * 2
		if mux%2 != 0 {
			defer fmt.Println(mux)
		} else {
			fmt.Println(mux)
		}
	}
}
