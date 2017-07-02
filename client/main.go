package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

var address string
var addressNew string

func main() {
	fmt.Println("\n>LAUNCHING THE MESSAGE APPLICATION.")
	acceptCaller()
	client()
}

func acceptCaller() {
	l, err := net.Listen("tcp", "192.168.1.9:9000")
	errorCheck(err)
	c, err := l.Accept()
	address = c.RemoteAddr().String()
	addressSlice := strings.Split(address, ":")
	addressString := addressSlice[0]
	fmt.Println(addressString)
	addressNew = addressString + ":8999"
	fmt.Println(addressNew)
	l.Close()
}

func errorCheck(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func client() {
	conn, err := net.Dial("tcp", addressNew)
	fmt.Println("CONNECTION ESTABLISHED \a")
	fmt.Println("You can write your messages now:")
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
