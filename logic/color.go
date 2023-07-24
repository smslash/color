package logic

import (
	"fmt"
)

func ColorCheck(s string) bool {
	s = s[8:]
	if s != "black" && s != "red" && s != "green" && s != "yellow" && s != "blue" && s != "purple" && s != "aqua" && s != "cyan" && s != "orange" && s != "violet" && s != "light-green" && s != "gray" && s != "bold" && s != "white" {
		fmt.Print("We have only this colors:\n- black\n- red\n- green\n- yellow\n- blue\n- purple\n- aqua\n- cyan\n- orange\n- violet\n- light-green\n- gray\n- bold\n- white\n")
		return false
	}
	return true
}
