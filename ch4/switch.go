package main

import "fmt"

func main() {
	words := []string{"a", "cow", "smile", "gopher", "octopus", "anthropologist"}

	for _, word := range words {
		switch size := len(word); size {
		case 1, 2, 3, 4:
			fmt.Println(word, "is a short word")
		case 5:
			wordLen := len(word)
			fmt.Println(word, "is exactly right len:", wordLen)
		case 6, 7, 8, 9:
		default:
			fmt.Println(word, "is a long word")
		}
	}
	/*
	   a is a short word
	   cow is a short word
	   smile is exactly right len: 5
	   anthropologist is a long word
	*/

	fmt.Println("--- Blank Switches")

	for _, word := range words {
		switch wordLen := len(word); {
		case wordLen < 5:
			fmt.Println(word, "is a short word")
		case wordLen > 10:
			fmt.Println(word, "is a long word")
		default:
			fmt.Println(word, "is exactly right len:", wordLen)
		}
	}
	/*
	   a is a short word
	   cow is a short word
	   smile is exactly right len: 5
	   gopher is exactly right len: 6
	   octopus is exactly right len: 7
	   anthropologist is a long word
	*/
}
