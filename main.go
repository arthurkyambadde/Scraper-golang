package main

import (
	gethttprequest "webscrapper/getHttpRequest"
	readFile "webscrapper/readFile"
	regexfunc "webscrapper/regexfunc"
)

func main() {
	companies := readFile.Readfile("dummy.txt")
	for i := 0; i < len(companies); i++ {
		body := gethttprequest.Gethttprequest(companies[i])

		regexfunc.GetFacebookURL(body)

	}
}
