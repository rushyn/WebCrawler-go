package main

import (
	"fmt"
	"log"
	"strings"
)


func (cfg *config) crawlPage(rawCurrentURL string){
	defer cfg.wg.Done()

	if cfg.stop() {
		fmt.Println("stop!!!")
		return
	}

	cfg.concurrencyControl <- struct{}{}	

	if !strings.Contains(rawCurrentURL, cfg.baseURL) {
		log.Printf("base url is not pressent in current url, base url %s, url %s will not be cralled.", cfg.baseURL, rawCurrentURL)
		<- cfg.concurrencyControl
		return
	}
	normalizedCurrentUrl, err := normalizeURL(rawCurrentURL)
	if err != nil {
		log.Println(err)
		<- cfg.concurrencyControl
		return
	}

	if !cfg.addPageVisit(normalizedCurrentUrl) {
		<- cfg.concurrencyControl
		return
	}

	log.Printf("Getting html for url: %s\n", normalizedCurrentUrl)
	html, err := getHTML(normalizedCurrentUrl)
	if err != nil {
		log.Println(err)
		<- cfg.concurrencyControl
		return
	}
	urls, err := getURLsFromHTML(html, normalizedCurrentUrl)
	if err != nil {
		log.Println(err)
	}

	for _, url := range urls {
		cfg.wg.Add(1)
		go cfg.crawlPage(url)
	}

	<- cfg.concurrencyControl
	
}


func (cfg *config) addPageVisit(normalizedURL string) (isFirst bool) {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()
	if _, exist := cfg.pages[normalizedURL]; exist {
		cfg.pages[normalizedURL] ++
		log.Printf("page alrady cralled %s and has been seen %d\n", normalizedURL, cfg.pages[normalizedURL])
		return false
	}
	cfg.pages[normalizedURL] = 1
	return true
	
}



func (cfg *config) stop() (tooLong bool) {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()
	return len(cfg.pages) >= cfg.maxPages
	
}