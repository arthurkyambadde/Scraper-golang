package gethttprequest

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func Gethttprequest(company string) string {
	bodyHtml := ""
	link := fmt.Sprintf("https://www.google.com/search?q=%s", company)
	resp, err := http.Get(link)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	log.Println("Response status:", resp.Status)

	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 15; i++ {
		bodyHtml = bodyHtml + scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	log.Println(bodyHtml)

	return bodyHtml
}
