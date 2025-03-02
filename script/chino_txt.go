package script

import (
	"io/ioutil"
	"log"
)

func ChinoTxt() string {
	filename := "data/script.txt"
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	return string(content)
}
