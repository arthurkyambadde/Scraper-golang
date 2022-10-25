package gethttprequest

import (
	"fmt"

	"github.com/gocolly/colly"
)

func Gethttprequest() {
	fmt.Println("worked")
	collyLib := colly.NewCollector()

	// On every a element which has href attribute call callback
	collyLib.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")

		// Print link
		fmt.Printf("Link found: %q -> %s\n", e.Text, link)

		// Visit link found on page
		collyLib.Visit(e.Request.AbsoluteURL(link))
	})

	// Before making a request print "Visiting ..."
	collyLib.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping on through the link
	collyLib.Visit("https://www.google.com/search?q=stanbic")
}
