package readfile

import (
	"fmt"
	"io/ioutil"
	"log"
)

func readfile(fileName string) {

	fileContent, err := ioutil.ReadFile(fileName)

	if err != nil {
		log.Fatalf("%q has not been read", err)
	}

	fileData := string(fileContent)

	fmt.Println(fileData)
}
