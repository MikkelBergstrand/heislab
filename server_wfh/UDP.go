package main

import (
	"bufio"
	"fmt"
	"net"
)

func udp_client(l_addr string, port string) {
	p := make([]byte, 2048)
	update_address := l_addr + ":" + port
	conn, err := net.Dial("udp", update_address)
	if err != nil {
		fmt.Printf("Some error %v", err)
		return
	}
	fmt.Fprintf(conn, "Hi UDP Server, How are you doing?")
	_, err = bufio.NewReader(conn).Read(p)
	if err == nil {
		fmt.Printf("%s\n", p)
	} else {
		fmt.Printf("Some error %v\n", err)
	}
	conn.Close()

}
func sendResponse(conn *net.UDPConn, addr *net.UDPAddr) {
	_, err := conn.WriteToUDP([]byte("From server: Hello I got your message "), addr)
	if err != nil {
		fmt.Printf("Couldn't send response %v", err)
	}
}

func udp_server(l_addr string) {
	p := make([]byte, 2048)
	addr := net.UDPAddr{
		Port: 1234,
		IP:   net.ParseIP(l_addr),
	}
	ser, err := net.ListenUDP("udp", &addr)
	if err != nil {
		fmt.Printf("Some error %v\n", err)
		return
	}
	for {
		_, remoteaddr, err := ser.ReadFromUDP(p)
		fmt.Printf("Read a message from %v %s \n", remoteaddr, p)
		if err != nil {
			fmt.Printf("Some error  %v", err)
			continue
		}
		go sendResponse(ser, remoteaddr)
	}
}

func read_UDP(network, laddr string) {
	net.ListenUDP(network, laddr)
}

func write_UDP(raddr, msg string, msg_size uint64) {
	net.WriteToUDP(msg, raddr)
}

func main() {
	network = "udp"
	laddr := "127.0.0.1"
	raddr := "0"
	msg = "HEI, JEG ER KLAR FOR ALT!\n"
	msg_size = len(msg)
	conn, err = net.Dial(network, laddr)
	go read_UDP(network, laddr)
	go write_UDP(raddr, msg, msg_size)

}
