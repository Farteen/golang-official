package netlearning

import (
	"bufio"
	"fmt"
	"net"
)

func DialFunc() string {
	conn, err := net.Dial("tcp", "github.com:443")
	if err != nil {
		return ""
	}
	
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	status, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		return ""
	}
	return status
}
