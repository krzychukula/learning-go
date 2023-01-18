package main

import (
	"fmt"
)

func main() {

	x := []int{1, 2, 3, 4}
	y := make([]int, 4)

	num := copy(y, x)
	fmt.Println(y, num)
	// [1 2 3 4] 4

	// Copy from the middle

	z := make([]int, 2)
	num = copy(z, x[2:]) // the only good use of sicing
	fmt.Println(z, num)
	// [3 4] 2
}
