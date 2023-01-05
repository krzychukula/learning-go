package main

import (
	"fmt"
)

const x int64 = 10
const (
	idKey   = "id"
	nameKey = "name"
)
const z = 20 * 10

func main() {
	const y = "hello"

	fmt.Println(x)
	fmt.Println(y)

	x = x + 1
	y = "bye"
	//./const.go:20:2: cannot assign to x (constant 10 of type int64)
	//./const.go:21:2: cannot assign to y (untyped string constant "hello")

	fmt.Println(x)
	fmt.Println(y)
}
