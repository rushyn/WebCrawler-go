package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
)

type config struct {
	maxPages		   int
	pages              map[string]int
	baseURL            string
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
}


func main(){

	if len(os.Args) < 2 {
		fmt.Println("no website provided")
		os.Exit(1)
	}
	if len(os.Args) < 3 {
		fmt.Println("purrelle workers number not suppliesd")
		os.Exit(1)
	}
	if len(os.Args) < 4 {
		fmt.Println("max pages to crall not supplied")
		os.Exit(1)
	}
	if len(os.Args) > 5 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	maxWorkers, err := strconv.Atoi(os.Args[2])
    if err != nil {
        // ... handle error
        panic(err)
    }

	maxPages, err := strconv.Atoi(os.Args[3])
    if err != nil {
        // ... handle error
        panic(err)
    }




	url, err := normalizeURL(os.Args[1])
	if err != nil {
		log.Println(err)
		return
	}

	cfg := config{
		maxPages: 			maxPages,
		pages:              map[string]int{},
		baseURL:            url,
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, maxWorkers),
		wg:                 &sync.WaitGroup{},
	}

	cfg.wg.Add(1)
	go cfg.crawlPage(cfg.baseURL)


	cfg.wg.Wait()
	i := 0
	for key := range cfg.pages {
		fmt.Printf("%d %s %d\n", i, key, cfg.pages[key])
		i++
	}
	
}