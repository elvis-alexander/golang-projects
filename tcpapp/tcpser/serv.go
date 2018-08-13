package main

import (
	"fmt"
	"net"
)

func echoServer(conn net.Conn) {
	for {

		var bytes = make([]byte, 1024)
		numBytes, err := conn.Read(bytes)
		if err != nil {
			return
		}
		var in = bytes[:numBytes]
		fmt.Printf("Server got: %s\n", string(in))
		_, e := conn.Write(in)
		if e != nil {
			panic(e)
		} else {
			fmt.Println("wrote to client")
		}
	}

}

func main() {
	/**/
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		panic(err)
	}
	for {
		fmt.Println("accepting")
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go echoServer(conn)
	}
}
