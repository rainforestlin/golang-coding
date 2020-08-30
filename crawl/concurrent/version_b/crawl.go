package version_b

//用于解决version_a 中因为同时创建太多的网络连接，超过了程序能打开文件数的限制，导致类似于DNS查询和net.Dial的连接失败
//使用容量为n的缓冲通道来建立并发原语，即技术信号量
import (
	"fmt"
	"github.com/julianlee107/learning/crawl/links"
	"log"
	"os"
)

// 令牌是一个计数信号量
// 确保并发请求限制在我们定义的数量内
var tokens = make(chan struct{}, 20)

func Crawl() {
	workList := make(chan []string)

	go func() { workList <- os.Args[1:] }()

	seen := make(map[string]bool)

	for list := range workList {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
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
