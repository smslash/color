package main

import (
	"os"

	"color/logic"
)

func main() {
	if len(os.Args) != 3 && len(os.Args) != 4 && len(os.Args) != 5 {
		logic.Error(1)
		return
	}
	if len(os.Args) == 4 {
		if len(os.Args[1]) <= 8 || os.Args[1][:8] != "--color=" {
			logic.Error(1)
			return
		}
		if !logic.ColorCheck(os.Args[1]) {
			return
		}

		logic.Default(os.Args[3])
	} else if len(os.Args) == 3 {
		logic.Default(os.Args[2])
	} else if len(os.Args) == 5 {
		if !logic.ColorCheck(os.Args[1]) {
			return
		}
		if !logic.Characters(os.Args[3]) {
			logic.Error(2)
			return
		}
		if os.Args[4] != "standard" && os.Args[4] != "shadow" && os.Args[4] != "thinkertoy" {
			logic.Error(3)
			return
		}
		logic.Check(os.Args)
	}
}
