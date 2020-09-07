package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	readFromFileByIO()
}

func readFromFile() {
	fileObj, err := os.Open("crawl.go")
	if err != nil {
		fmt.Printf("open file err:%v \n", err)
	}
	defer fileObj.Close()
	var temp = make([]byte, 1000)
	n, err := fileObj.Read(temp)
	if err != nil {
		fmt.Printf("read from file err:%v", err)
	}
	fmt.Printf("read byte %d \n", n)
	fmt.Printf("read from file %s", string(temp[:]))
}

func readFromFileByBuff() {
	fileObj, err := os.Open("crawl.go")
	if err != nil {
		fmt.Printf("open file err:%v \n", err)
	}
	defer fileObj.Close()

	reader := bufio.NewReader(fileObj)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Printf("%v", err)
			return
		}
		if err != nil {
			fmt.Printf("%v", err)
			return
		}
		fmt.Print(line)
	}

}

func readFromFileByIO() {
	ret, err := ioutil.ReadFile("./crawl.go")
	if err != nil {
		fmt.Printf("%v", err)
	}
	fmt.Println(ret)
}
