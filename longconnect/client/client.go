package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func createSocket() {
	tcpAdd, err := net.ResolveTCPAddr("tcp", "127.0.0.1:9100")
	if err != nil {
		fmt.Println("net.ResolveTCPAddr error", err)
		return
	}
	conn, err := net.DialTCP("tcp", nil, tcpAdd)
	if err != nil {
		fmt.Println("net.DailTCP error:", err)
		return
	}
	defer func() {
		err = conn.Close()
		if err != nil {
			fmt.Println("conn.Close error:", err)
		}
	}()
	fmt.Println("connected")
	go onMessageRectived(conn) //读取服务端广播的信息

	for {
		// 自己发送的信息
		var data string
		fmt.Scan(&data)
		if data == "quit" {
			break
		}
		b := []byte(data + "\n")
		conn.Write(b)
	}

}

// 获取服务端发送来的信息
func onMessageRectived(conn *net.TCPConn) {
	reader := bufio.NewReader(conn)
	for {
		// var data string
		msg, err := reader.ReadString('\n') //读取直到输入中第一次发生 ‘\n’
		fmt.Println(msg)
		if err != nil {
			fmt.Println("err:", err)
			os.Exit(1) //服务端错误的时候，就将整个客户端关掉
		}
	}
}

func main() {
	fmt.Println("开启客户端")
	createSocket()
}
