// ZK-Client
package main

import (
	"flag"
	. "project/mypackage"
)

var RemoteIP string
var RemotePort int

// 帮助命令：**.exe -h
// 启动命令：**.exe -ip [127.0.0.1] -port [8888]
func init() {
	flag.StringVar(&RemoteIP, "ip", "127.0.0.1", "服务器IP")
	flag.IntVar(&RemotePort, "port", 8888, "服务器Port")
}

func main() {
	// 解析启动命令
	flag.Parse()

	client := Client{
		ServerIP:   RemoteIP,
		ServerPort: RemotePort,
		Name:       "ZK",
		Flag:       999,
	}

	client.CreateConn()

	go client.Run()

	select {}
}
