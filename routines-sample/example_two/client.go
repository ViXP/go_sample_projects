package example_two

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func RunClient() {
	fmt.Println("Starting client...")

	connection, err := net.Dial("tcp", "127.0.0.1:3330")

	if err != nil {
		log.Fatal(err)
	}

	defer connection.Close()

	mustCopy(os.Stdout, connection)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
