package scraper

import (
	"github.com/PuerkitoBio/goquery"
	"fmt"
)

// Scraper takes urls, and selectors and returns information.
type Scraper interface{
	// GetContents returns the concatenated contents of all matching nodes
	GetContents(string, string) (string, error)
}

func NewScraper() Scraper {
	return &scraper{}
}

type scraper struct {
}

func (s *scraper) GetContents(url, selector string) (string, error) {
	d, err := goquery.NewDocument(url)
	if err != nil {
		return "", fmt.Errorf("failed to get document at %s...\n%s", url, err)
	}

	return d.Find(selector).Text(), nil
}
