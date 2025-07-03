package main

import (
	"fmt"
	. "project/mypackage"
)

func main() {
	fmt.Println("main start...")

	server := Server{
		Id:               "127.0.0.1",
		Port:             8888,
		OnlineMap:        make(map[string]Client),
		BroadcastMessage: make(chan string),
	}
	server.Start()
}
