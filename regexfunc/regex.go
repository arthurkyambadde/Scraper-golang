package regex

import (
	"fmt"
	"regexp"
)

func GetFacebookURL(link string) {

	r, err := regexp.Compile(link)

	if err != nil {
		// log.Fatal(err)
	}
	fmt.Println(r.FindString(`facebook`))

}
