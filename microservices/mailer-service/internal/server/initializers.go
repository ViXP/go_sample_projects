package server

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

const (
	serverPort = 80
	rpcPort    = 5001
)

func StartHTTPServer() error {
	fmt.Printf("Starting the Mailer REST server on the port %v...\n", serverPort)
	return http.ListenAndServe(fmt.Sprintf(":%v", serverPort), Routes())
}

func StartRPCServer() error {
	fmt.Printf("Starting the Mailer RPC server on the port %v...\n", rpcPort)
	listen, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%v", rpcPort))

	if err != nil {
		log.Panicf("Can't listen on 0.0.0.0:%v", rpcPort)
		return err
	}

	defer listen.Close()

	err = rpc.Register(&RPCProcedures{})

	if err != nil {
		log.Panic("Can't register RPC procedures")
		return err
	}

	for {
		rpcConn, err := listen.Accept()
		if err != nil {
			continue
		}

		go rpc.ServeConn(rpcConn)
	}
}
