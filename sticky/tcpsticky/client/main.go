package main

import (
	"fmt"
	"net"
	"strconv"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:3000")

	if err != nil {
		fmt.Println("dial failed,err:", err)
		return
	}

	defer conn.Close()

	for i := 0; i < 20; i++ {
		msg := `Hello,hello,how are you ` + strconv.Itoa(i)
		conn.Write([]byte(msg))
	}
}
