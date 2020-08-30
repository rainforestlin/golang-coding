package version_a

import (
	"github.com/julianlee107/learning/crawl/base"
	"os"
)

func Crawl() {
	workList := make(chan []string)

	go func() { workList <- os.Args[1:] }()

	seen := make(map[string]bool)

	for list := range workList {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				go func(link string) {
					workList <- base.Crawl(link)
				}(link)
			}
		}
	}
}
