package main

import (
	"fmt"
	"log"
	"net"
)

var adress = "localhost:9000"

func main() {
	fmt.Println("\n>LAUNCHING THE SERVER APPLICATION.")
	server()

}

func errorCheck(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

//napsat nějakej ping na clienta, abych získal adresu serveru.

func server() {
	listener, err := net.Listen("tcp", adress)
	errorCheck(err)
	defer listener.Close()

	fmt.Println(">Listening on", adress)

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
