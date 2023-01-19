package main

import "fmt"

func main() {
	x := 10

	if x > 5 {
		fmt.Println(x)
		x := 5
		fmt.Println(x)
	}
	fmt.Println(x)
	/*
	   10
	   5
	   10
	*/

	if x > 5 {
		x, y := 5, 20
		fmt.Println(x, y) // 5 20
	}
	fmt.Println(x) // 10

	// fmt := "oops"
	// fmt.Println("x")
	// fmt.Println undefined (type string has no field or method Println)

	// REALLY???!!!
	true := 10
	fmt.Println(true) // true
}
