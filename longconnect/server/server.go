package main

import (
	"bufio"
	"errors"
	"fmt"
	"net"
)

var connSlice []*net.TCPConn

// 长连接
func createTcp() {
	tcpAddr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:9100")
	if err != nil {
		fmt.Println("net.ResolveTcpAddr error:", err)
		return
	}
	tcpListener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		fmt.Println("net.ListenTCP error:", err)
		return
	}
	defer func() {
		err = tcpListener.Close()
	}()
	for {
		tcpConn, err := tcpListener.AcceptTCP() // 阻塞，当有连接时，才会运行
		if err != nil {
			fmt.Println("tcpListener error:", err)
			continue
		}
		fmt.Println("a client connected:", tcpConn.RemoteAddr().String())
		boradcastMessage(tcpConn.RemoteAddr().String() + "entered \n")
		connSlice = append(connSlice, tcpConn)
		// 监听被访问时，开启一个协程处理
		go tcpPipe(tcpConn)
	}

}

// 广播数据
func boradcastMessage(message string) error {
	b := []byte(message)
	for i := 0; i < len(connSlice); i++ {
		fmt.Println(connSlice[i])
		_, err := connSlice[i].Write(b)
		if err != nil {
			fmt.Println("send ", connSlice[i].RemoteAddr().String(), " error:", err)
			continue
		}

	}
	return nil
}

// 移除已经关闭的客户端
func deleteConn(conn *net.TCPConn) error {
	if conn == nil {
		fmt.Println("conn is nil")
		return errors.New("conn is nil")
	}
	for i := 0; i < len(connSlice); i++ {
		if conn == connSlice[i] {
			connSlice = append(connSlice[:i], connSlice[i+1:]...)
			break
		}

	}
	return nil
}

func tcpPipe(conn *net.TCPConn) {
	ipStr := conn.RemoteAddr().String()
	fmt.Println("ip:", ipStr)
	defer func() {
		fmt.Println("disconnected:", ipStr)
		err := conn.Close()
		if err != nil {
			fmt.Println("tcpPipe.conn.CLose err:", err)
		}
		err = deleteConn(conn)
		if err != nil {
			fmt.Println("tcpPipe.deleteConn error:", err)
		}
		err = boradcastMessage(ipStr + "left the room")

	}()
	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if message == "\n" {
			return
		}
		message = ipStr + " said " + message
		if err != nil {
			fmt.Println("tcpPipe error:", err)
			return
		}
		fmt.Println(ipStr, " said ", message)
		err = boradcastMessage(message)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func main() {
	fmt.Println("服务端")
	createTcp()
}
