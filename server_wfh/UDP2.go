package main

import (
	"fmt"
	"net"
	"time"
)

func readingUDP(ln net.PacketConn, r chan string) {
	for {
		buf := make([]byte, 2048)
		n, addr, _ := ln.ReadFrom(buf)
		message := string(buf[:n])
		message = "'" + message + "'" + " from " + addr.String()
		r <- message
	}
}

func writeToChan(w chan string) {
	for {
		msg := "Hi UDP! Hope this reaches you!"
		w <- msg
		time.Sleep(1000 * time.Millisecond)
	}
}

func sendtoServer(conn net.Conn, w chan string) {
	for {
		msg := <-w
		if msg != "" {
			conn.Write([]byte(msg))
		}
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	local_addr_str := "127.0.0.1:1234"
	remote_addr_str := "127.0.0.1:5555"
	local_addr, _ := net.ResolveUDPAddr("udp", local_addr_str)
	remote_addr, _ := net.ResolveUDPAddr("udp", remote_addr_str)
	ln, _ := net.ListenUDP("udp", local_addr)
	defer ln.Close()
	conn, _ := net.DialUDP("udp", remote_addr, local_addr)
	defer conn.Close()

	read_ch := make(chan string)
	write_ch := make(chan string)

	go writeToChan(write_ch)

	go sendtoServer(conn, write_ch)
	go readingUDP(ln, read_ch)

	for {
		select {
		case msg := <-read_ch:
			fmt.Println(msg)
		default:
		}

	}

}
