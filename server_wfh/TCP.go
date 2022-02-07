package main

import (
	"net"
	"time"
)

func readTCP(conn net.Conn, r chan string) {
	for {
		buf := make([]byte, 2048)
		n, _ := conn.Read(buf)
		message := string(buf[:n])
		message = "'" + message + "'" + " from " + conn.RemoteAddr().String()
		r <- message
	}
}
func writeTCP(conn net.Conn, w chan string) {
	msg := <-w
	if msg != "" {
		conn.Write([]byte(msg))
	}
	time.Sleep(100 * time.Millisecond)
}

func main() {
	server_addr_str := ""
	local_addr_str := ""
	server_addr, _ := net.ResolveTCPAddr("tcp", server_addr_str)
	local_addr, _ := net.ResolveTCPAddr("tcp", local_addr_str)

	conn, _ := net.DialTCP("tcp", local_addr, server_addr)
	defer conn.Close()
	ln, _ := net.ListenTCP("tcp", local_addr)
	defer ln.Close()

	read_ch := make(chan string)
	write_ch := make(chan string)
	go readTCP(conn, read_ch)
	go writeTCP(conn, write_ch)

	write_ch <- "Connect to: #.#.#.#:#\000"
	for {
		c, _ := ln.Accept()
		defer c.Close()
	}
}
