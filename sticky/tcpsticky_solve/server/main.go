package main

import (
	"bufio"
	"fmt"
	"io"
	"net"

	"github.com/julianlee107/learning/sticky/tcpsticky_solve/proto"
)

func process(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)

	for {
		msg, err := proto.Decode(reader)

		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Printf("read from client failed: %v \n", err)
			break
		}
		fmt.Println("receive data from client:", msg)
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
