package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	writeToFileByIO()
}

func writeToFileByWrite() {
	file, err := os.OpenFile("main.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("write to file err%v", err)
	}
	defer file.Close()
	str := "Julianlee"
	file.Write([]byte(str))
	file.WriteString(str)
}

func writeToFileByBuff() {
	file, err := os.OpenFile("main.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("write to file err%v", err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString("hello world \n")
	}
	// 将缓存内的数据写入文件
	writer.Flush()
}

func writeToFileByIO() {
	str := "hello IO"
	err := ioutil.WriteFile("main.txt", []byte(str), 0644)
	if err != nil {
		return
	}
}
