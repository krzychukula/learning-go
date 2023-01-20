package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		func(j int) {
			fmt.Println("printing", j)
		}(i)
	}
}

/*
printing 0
printing 1
printing 2
printing 3
printing 4
*/
