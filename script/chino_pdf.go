package script

import (
	"fmt"
	"io/ioutil"
)

func ChinoPdf() string {
	filename := "data/chino.txt"
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
	}

	return string(content)
}
