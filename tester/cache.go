package main

import (
	"goprojects/memcli"
	"fmt"
)

func main() {
	cli, e := memcli.New("127.0.0.1:11211")
	if e != nil {
		panic(e)
	}
	bytes, i := cli.Get("name")
	if i != nil {
		panic(i)
	}
	fmt.Println(bytes)
}
