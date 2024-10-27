package main

import (
	"fmt"
	"log"
	"net/url"
	"strings"
)



func normalizeURL(link string) (string, error) {

	u, err := url.Parse(link)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(u)

	if len(u.Path) > 0 {
		if u.Path[len(u.Path)-1:] == "/" {
			u.Path = u.Path[:len(u.Path)-1]
		}
	}
	
	return strings.ToLower(u.Host + u.Path) , err
	
}