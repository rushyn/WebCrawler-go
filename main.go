package main

import (
	"fmt"
	"log"
	"os"
	"sort"
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
	
	printReport(cfg.pages, os.Args[1])
	
}


type page struct{
	link	string
	count	int
}

func printReport(pages map[string]int, baseURL string) {
	report := []page{}
	for key := range pages {
		report = append(report, page{
			link:  key,
			count: pages[key],
		})
	}


	sort.Slice(report, func(i, j int) bool {
		if report[i].count > report[j].count {
			return true
		}
		
		if z, err := strconv.Atoi(report[i].link); err == nil {
			if y, err := strconv.Atoi(report[j].link); err == nil {
				if report[i].count == report[j].count {
					return y < z
				}
				return false
			}
			if report[i].count == report[j].count {
				return true
			}
		}

		if report[i].count == report[j].count {
			return report[j].link > report[i].link
		}
		return false
	})
	



	fmt.Println("=============================")
	fmt.Printf("REPORT for %s\n", baseURL)
	fmt.Println("=============================")
	for _, item := range report {
		fmt.Printf("Found %d internal links to %s \n", item.count, item.link)
	}
	
}