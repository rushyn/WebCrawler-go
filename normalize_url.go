package main

import (
	"fmt"
	"log"
	"net/url"
)



func normalizeURL(link string) (string, error) {

	u, err := url.Parse(link)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(u)

	if u.Path[len(u.Path)-1:] == "/" {
		u.Path = u.Path[:len(u.Path)-1]
	}
	
	return u.Host + u.Path , err
	
}