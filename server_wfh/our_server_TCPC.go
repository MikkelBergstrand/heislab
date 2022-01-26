package main

addr = new InternetAddress(serverIP, serverPort) 
sock = new Socket(tcp) // TCP, aka SOCK_STREAM
sock.connect(addr)
// use sock.recv() and sock.send()

//server_ip = 30000

