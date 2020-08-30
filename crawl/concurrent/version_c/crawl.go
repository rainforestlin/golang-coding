package version_c

import (
	"fmt"
	"github.com/julianlee107/learning/crawl/links"
	"log"
	"os"
)

//解决程序永远不会结束的问题，即使所有可达连接都已经发现。
//原因大概是因为没有那个goroutine去关闭通道，导致一直循环

// 令牌是一个计数信号量
// 确保并发请求限制在我们定义的数量内
var tokens = make(chan struct{}, 5)

func Crawl() {
	workList := make(chan []string)
	// 等待发送到任务列表的数量
	var n int
	n++
	go func() { workList <- os.Args[1:] }()
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-workList
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string) {
					workList <- crawl(link)
				}(link)
			}
		}
	}
}

func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{}
	list, err := links.Extract(url)
	if err != nil {
		log.Println(err)
	}
	return list
}
