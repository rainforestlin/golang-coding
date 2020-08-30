package base

import (
	"fmt"
	"github.com/julianlee107/learning/crawl/links"
	"log"
)

func Crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Println(err)
	}
	return list
}
