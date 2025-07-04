package main

import (
	"flag"
	"fmt"
	. "project/mypackage"
)

var ServerIP string
var ServerPort int

// 帮助命令：**.exe -h
// 启动命令：**.exe -ip [127.0.0.1] -port [8888]
func init() {
	flag.StringVar(&ServerIP, "ip", "127.0.0.1", "服务器IP")
	flag.IntVar(&ServerPort, "port", 8888, "服务器Port")
}

func main() {
	// 解析启动命令
	flag.Parse()

	fmt.Println("main start...")

	server := Server{
		Id:               ServerIP,
		Port:             ServerPort,
		OnlineMap:        make(map[string]User),
		BroadcastMessage: make(chan string),
	}
	server.Start()
}
