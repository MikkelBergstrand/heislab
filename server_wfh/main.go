package main

import "net"

func main() {
	rAddr, _ := net.ResolveUDPAddr("udp", "10.100.23.240:20010")
	conn, err := net.DialUDP("udp", nil, rAddr)

	if err != nil {
		print("Taper")
	}

	defer conn.Close()
	conn.Write([]byte("EG ER KLAR FOR ALT!"))

}

//send, listen, dial,
//
