package main

import (
	"fmt"
	"net"
)

func main() {
	//laddr_str := "10.22.72.55:0"
	raddr_str := "10.100.23.240:34933"
	//laddr, _ := net.ResolveTCPAddr("tcp", laddr_str)
	raddr, _ := net.ResolveTCPAddr("tcp", raddr_str)
	conn, _ := net.DialTCP("tcp", nil, raddr)
	for {
		buf := make([]byte, 2048)
		conn.Read(buf)
		fmt.Println(string(buf))
	}
}
