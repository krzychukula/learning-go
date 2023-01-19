package main

import "fmt"

func main() {
	evenVals := []int{2, 4, 6, 8, 10, 12}
	for i, v := range evenVals {
		fmt.Println(i, v)
	}
	/*
		0 2
		1 4
		2 6
		3 8
		4 10
		5 12
	*/
	for _, v := range evenVals {
		fmt.Println(v)
	}
	uniqueNames := map[string]bool{
		"Fred":  true,
		"Bob":   false,
		"Alice": true,
	}
	for k := range uniqueNames {
		// you can't depend on the order when iterating over a map
		fmt.Println(k)
	}

	for i, r := range "abcąbć" {
		// only use value in strings!!!
		fmt.Println(i, r, string(r))
	}
	/*
	   0 97 a
	   1 98 b
	   2 99 c
	   3 261 ą
	   5 98 b
	   6 263 ć
	*/
}
