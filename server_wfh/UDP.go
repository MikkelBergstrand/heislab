package main

import (
	"fmt"
	"net"
	"time"
)

func handleTheQuit(q chan string) {
	var num_secs int64 = 10
	for i := 1; i <= int(num_secs); i++ {
		time.Sleep(1000 * time.Millisecond)
	}
	q <- "Quit"
}

func readingUDP(ln *net.UDPConn, r chan string) {
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

//func sendtoClient(ln *net.UDPConn)

func main() {

	reading := make(chan string)
	writing := make(chan string)
	quit := make(chan string)

	address_str := "10.22.72.55:30000"
	address, _ := net.ResolveUDPAddr("udp", address_str)
	ln, err := net.ListenUDP("udp", address)
	if err != nil {
		panic(err)
	}
	defer ln.Close()
	conn, _ := net.DialUDP("udp", nil, address)
	defer conn.Close()

	go readingUDP(ln, reading)
	go sendtoServer(conn, writing)
	go writeToChan(writing)

	go handleTheQuit(quit)

mainloop:
	for {
		select {
		case msg := <-reading:
			fmt.Println("\nReceived: ", msg)
		case <-quit:
			break mainloop
		default:

		}
	}

}
