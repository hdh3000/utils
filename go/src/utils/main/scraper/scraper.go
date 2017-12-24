package main

import (
	"flag"
	"log"
	"os"
	"utils/scraper"
)

var fUrl = flag.String("url", "", "The url to scrape")
var fSelector = flag.String("s", "", "The default css selector to grab text from")

func main() {
	flag.Parse()
	s := scraper.NewScraper()
	contents, err := s.GetContents(*fUrl, *fSelector)
	if err != nil {
		log.Fatalf("failed to scrape %s", err)
	}

	os.Stdout.WriteString(contents + "\n")
}
