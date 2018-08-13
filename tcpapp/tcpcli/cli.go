package main

import (
	"fmt"
	"net"
	"time"
)

func readStuff(conn net.Conn) {
	for {
		var bytes = make([]byte, 1024)
		fmt.Println("read blocking???")
		n, err := conn.Read(bytes)
		if err != nil {
			panic(err)
		}
		fmt.Println("Got", string(bytes[:n]))
	}
}

func main() {
	conn, e := net.Dial("tcp", "127.0.0.1:8080")
	if e != nil {
		panic(e)
	}
	defer conn.Close()
	go readStuff(conn)
	for {
		_, err := conn.Write([]byte("Hello Server"))
		if err != nil {
			panic(err)
		}
		/*ping server every 3 seconds*/
		time.Sleep(time.Second * 1)
	}
}
