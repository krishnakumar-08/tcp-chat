package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
)

func main() {
	server, err := net.Listen("tcp", "localhost:8080")
	var db sync.Map
	if err != nil {
		log.Println(err)
	}
	fmt.Println("server has started")
	for {
		accept, err := server.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		register(&db, accept)
		go handleconn(accept, &db)
	}
}
func handleconn(n net.Conn, db *sync.Map) {
	var mess string
	reader := bufio.NewReader(n)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			log.Println(err)
			break
		}
		if strings.TrimSpace(strings.ToLower(message)) == "exit" {
			mess = fmt.Sprint((n.RemoteAddr().String() + "is closed"))
			n.Close()
			break
		} else {
			mess = fmt.Sprint(n.RemoteAddr().String() + ": " + strings.TrimSpace(message))
		}
		db.Range(func(key, value any) bool {
			if key != n.RemoteAddr().String() {
				fmt.Println(n.RemoteAddr().String(), " != ", key)
				fmt.Println(n.RemoteAddr().String(), " != ", value.(net.Conn).RemoteAddr().String())
				value.(net.Conn).Write([]byte(mess + "\n"))
			}
			return true
		})
	}

}
func register(m *sync.Map, n net.Conn) {
	m.Store(n.RemoteAddr().String(), n)
	fmt.Println("registered ", n.RemoteAddr().String())
}
