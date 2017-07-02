package main

import (
	"fmt"
	"log"
	"net"
)

var address = "192.168.1.9:8999"

func main() {
	fmt.Println("\n>LAUNCHING THE SERVER APPLICATION.")
	giveAdress()
	server()
}

func errorCheck(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func giveAdress() {
	conn, err := net.Dial("tcp", "192.168.1.9:9000")
	errorCheck(err)
	conn.Close()
}

func server() {
	listener, err := net.Listen("tcp", address)
	errorCheck(err)
	defer listener.Close()

	fmt.Println(">Listening on", address)

	for {
		conn, err := listener.Accept()
		errorCheck(err)
		net := conn.RemoteAddr().Network()
		address := conn.RemoteAddr().String()
		fmt.Println(">New connection with", address, "over", net, "protocol")
		fmt.Println("")
		go listenConnection(conn)
	}
}

func listenConnection(conn net.Conn) {
	for {
		buffer := make([]byte, 1024)
		dataSize, err := conn.Read(buffer)
		errorCheck(err)
		data := buffer[:dataSize]
		message := string(data)
		fmt.Println("<<:", message)
	}
}
