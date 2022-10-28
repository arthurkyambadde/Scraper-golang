package readfsile

import (
	"bufio"
	"log"
	"os"
)

var fileContent []string

func Readfile(fileName string) []string {
	fileToBeRead, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("%q has not been read", err)
	}
	defer fileToBeRead.Close()

	scanner := bufio.NewScanner(fileToBeRead)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		fileContent = append(fileContent, scanner.Text())
	}

	return fileContent
}
