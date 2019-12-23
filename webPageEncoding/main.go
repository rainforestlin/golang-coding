package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"net/http"
)

// 使用http对网页进行解析时，当中文字符不为utf-8格式时，中文显示乱码
// when I was using http to parse html, can't display simplified chinese aright while charset is gbk
func main() {
	resp, err := http.Get("https://www.lagou.com/")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		encodeBody := determineEncoding(resp.Body)
		utf8Reader := transform.NewReader(resp.Body, encodeBody.NewDecoder())
		all, err := ioutil.ReadAll(utf8Reader)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s\n", all)

	}

}

func determineEncoding(r io.Reader) encoding.Encoding {
	//读取前面1024个字节
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		panic(err)
	}
	//判断编码
	encode, name, _ := charset.DetermineEncoding(bytes, "")
	// 为编码类型
	fmt.Print(name)
	return encode
}
