package main

import (
	"fmt"
)

func main() {
	x := make([]int, 0, 5)
	x = append(x, 1, 2, 3, 4)
	y := x[:2]
	z := x[2:]

	fmt.Println("cap x y z", cap(x), cap(y), cap(z))

	y = append(y, 30, 40, 50)
	x = append(x, 60)
	z = append(z, 70)

	fmt.Println("x", x)
	fmt.Println("y", y)
	fmt.Println("z", z)
	/*
	   cap x y z 5 5 3
	   x [1 2 30 40 70]
	   y [1 2 30 40 70]
	   z [30 40 70]
	*/
}
