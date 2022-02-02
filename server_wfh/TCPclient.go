package main

import (
	"fmt"
	"net"
	"time"
)

func writeTo(conn *net.TCPConn) {
	for {
		msg := "Ping Ping <3\000"
		conn.Write([]byte(msg))
		time.Sleep(1000 * time.Millisecond)
	}
}

func readFrom(conn *net.TCPConn) {
	for {
		buf := make([]byte, 2048)
		conn.Read(buf)
		fmt.Println(string(buf))
	}
}

func main() {
	//laddr_str := "10.22.72.55:0"
	raddr_str := "10.100.23.240:33546"
	//laddr, _ := net.ResolveTCPAddr("tcp", laddr_str)
	raddr, err := net.ResolveTCPAddr("tcp", raddr_str)
	if err != nil {
		fmt.Print("err 1")
		panic(err)
	}

	conn, err := net.DialTCP("tcp", nil, raddr)
	if err != nil {
		fmt.Print("err 2")
		panic(err)
	}

	defer conn.Close()

	go readFrom(conn)
	go writeTo(conn)

	for {
	}

}
