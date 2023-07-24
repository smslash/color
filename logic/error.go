package logic

import "fmt"

func Error(a int) {
	if a == 1 {
		fmt.Println("Usage: go run . --color=<color> [LETTERS] [STRING] [BANNER]")
		fmt.Println("EX: go run . --color=\"Blue\" so something standard")
		return
	} else if a == 2 {
		fmt.Println("Error: input only ASCII Table characters between 32-126 numbers")
		return
	} else if a == 3 {
		fmt.Println("Error: only standard, shadow, and thinkertoy banners are available")
		return
	}
}
