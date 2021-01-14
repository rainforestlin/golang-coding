package main

import (
	"fmt"
	"net"
	"strconv"

	"github.com/julianlee107/learning/sticky/tcpsticky_solve/proto"
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
		data, err := proto.Encode(msg)
		if err != nil {
			fmt.Println("encode msg failed,err:", err)
			return
		}
		conn.Write(data)
	}
}
