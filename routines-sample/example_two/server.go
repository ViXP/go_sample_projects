package example_two

import (
	"fmt"
	"io"
	"net"
	"time"
)

func RunServer() {
	fmt.Println("Starting server...")
	listener, err := net.Listen("tcp", "127.0.0.1:3330")

	if err != nil {
		panic(err)
	}

	for {
		connection, err := listener.Accept()

		if err != nil {
			continue
		}

		go handleConn(connection)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, "server's response\n")
		if err != nil {
			panic("connection is not writable!")
		}
		time.Sleep(time.Second)
	}
}
