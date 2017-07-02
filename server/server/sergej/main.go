package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	fmt.Println("\n>LAUNCHING THE SERVER APPLICATION.")
	server()

}

func errorCheck(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func callHome() {
	caller, err := net.Dial("tcp", "94.230.146.233:9000")
	defer caller.Close()

}

func server() {
	listener, err := net.Listen("tcp", adress)
	errorCheck(err)
	defer listener.Close()

	fmt.Println(">Listening on", adress)

	for {
		conn, err := listener.Accept()
		errorCheck(err)
		network := conn.RemoteAddr().Network()
		address := conn.RemoteAddr().String()
		fmt.Println(">New connection with", address, "over", network, "protocol")
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
		fmt.Println("<<", message)
	}
}
