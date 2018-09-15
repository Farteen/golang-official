package netlearning

import (
	"bufio"
	"fmt"
	"net"
)

func DialFunc() string {
	conn, err := net.Dial("tcp", "www.baidu.com:80")
	if err != nil {
		fmt.Println("dail error", err)
		return ""
	}
	
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	status, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println("read error", err)
		return ""
	}
	return status
}


func Server(ch chan string) {
	ln, err := net.Listen("tcp", ":10001")
	if err != nil {
		ch<-"wrong"
	}

	for {
		_, err := ln.Accept()
		if err != nil {
			ch<-"wrong"
		}
	}
}

// func HandleConnection(conn Conn) {
// 	fmt.Println(conn)
// }