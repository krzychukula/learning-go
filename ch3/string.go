package main

import (
	"fmt"
)

func main() {
	var s string = "Hello, ę"
	var bs []byte = []byte(s)
	var rs []rune = []rune(s)

	fmt.Println(bs)
	fmt.Println(rs)
	/*
		[72 101 108 108 111 44 32 196 153]
		[72 101 108 108 111 44 32 281]
	*/
}
