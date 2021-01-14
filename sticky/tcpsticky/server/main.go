package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)

	var buf [1024]byte

	for {
		n, err := reader.Read(buf[:])

		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Printf("read from client failed: %v \n", err)
			break
		}

		recvStr := string(buf[:n])
		fmt.Println("receive data from client:", recvStr)
	}
}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:3000")
	if err != nil {
		fmt.Println("listen failed,err:", err)
		return
	}

	defer listen.Close()

	for {
		conn, err := listen.Accept()

		if err != nil {
			fmt.Println("accept failed,err:", err)
			continue
		}
		go process(conn)
	}
}
