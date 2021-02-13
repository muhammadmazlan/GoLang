package main

import "fmt"

func main() {
	i := "mamat"
	fmt.Println(whoAmI(i))

	j := 2
	fmt.Println(whoAmI(j))

	k := 35.6
	fmt.Println(whoAmI(k))
}

func whoAmI(t interface{}) string {
	switch t.(type) { //assertion
	case string:
		return "The String"
	case int:
		return "Integer"
	default:
		return "not known"
	}
}
