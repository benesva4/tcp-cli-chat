package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

var address string

func main() {
	fmt.Println("\n>LAUNCHING THE MESSAGE RECEIVER APPLICATION.")
	acceptCaller()
	client()
}

func acceptCaller() {
	l, err := net.Listen("tcp", "localhost:9000")
	errorCheck(err)
	c, err := l.Accept()
	address = c.RemoteAddr().String()
	l.Close()
}


func errorCheck(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func client() {
	conn, err := net.Dial("tcp", address)
	fmt.Println(">Connection established:")
	fmt.Println("")
	errorCheck(err)
	defer conn.Close()

	for {
		fmt.Print(">>: ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		message := scanner.Text()
		conn.Write([]byte(message))
	}
}
