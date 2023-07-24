package logic

import (
	"fmt"
	"os"
)

func Check(s []string) {
	stringValue := s[3]
	bannerName := s[4]

	data, err := os.ReadFile("./banners/" + bannerName + ".txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fileHash := MD5(string(data))
	if !CheckHash(fileHash, bannerName) {
		fmt.Println("Error: hash code of", bannerName+".txt", "has been changed")
		return
	}

	result := Process(stringValue, data)

	fmt.Print(result)
}
