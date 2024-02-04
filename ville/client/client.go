package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	var message string
	client, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Println(err)
	}
	fmt.Println("client has connnected to the server")
	go func() {
		Reader(client)
	}()
	for {
		fmt.Fscan(os.Stdin, &message)
		client.Write([]byte(message + "\n"))
	}

}
func Reader(n net.Conn) {
	reader := bufio.NewReader(n)
	for {
		s, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("its from client")
			log.Println(err)
			n.Close()
			break
		}
		fmt.Println(strings.TrimSpace(s))
	}
}
