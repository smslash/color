package logic

import (
	"fmt"
	"io/ioutil"
)

func Default(s string) {
	data, err := ioutil.ReadFile("./banners/standard.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fileHash := MD5(string(data))
	if !CheckHash(fileHash, "standard") {
		fmt.Println("Error: hash code of standard.txt has been changed")
		return
	}

	fmt.Print(Process(s, data))
}
