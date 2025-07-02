// server
package mypackage

import (
	"fmt"
	"net"
	"sync"
)

type Server struct {
	Id   string
	Port int

	// 在线用户列表
	OnlineMap map[string]Client
	mapLock   sync.RWMutex

	// 消息广播的channel
	ServerMessage chan string
}

// 广播消息
func (this *Server) BroadcastMessageSend(client Client, msg string) {
	sendMsg := "[" + client.Addr + "]" + ":" + msg
	this.ServerMessage <- sendMsg
}

// handler function
func (this *Server) Handler(conn net.Conn) {
	// 当前连接的业务
	fmt.Println("链接建立成功：", conn.RemoteAddr().String())

	// 创建当前客户端
	clientAddr := conn.RemoteAddr().String()
	client := Client{
		Name:          clientAddr,
		Addr:          clientAddr,
		ClientMessage: make(chan string),
		conn:          conn,
	}
	// 启动客户端广播监听
	client.ListenClientMessage()

	// 广播当前客户端上线消息
	this.BroadcastMessageSend(client, "已上线")

	// 保证当前客户不会收到广播消息
	endMsg := <-this.ServerMessage
	fmt.Println(endMsg)

	// 将客户加入OnlineMap中
	this.mapLock.Lock()
	this.OnlineMap[client.Name] = client
	this.mapLock.Unlock()

	// 阻塞当前Handler
	select {}
}

// 监听广播消息
func (this *Server) ListenBroadcastMessage() {
	for {
		msg := <-this.ServerMessage

		// 将消息发送给在线的客户端
		this.mapLock.Lock()
		for _, client := range this.OnlineMap {
			client.ClientMessage <- msg
		}
		this.mapLock.Unlock()

		this.ServerMessage <- "广播完成"
	}
}

// Server Start
func (this *Server) Start() {
	// socket listen
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Id, this.Port))
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}

	// listen close
	defer listener.Close()

	// 启用广播消息监听
	go this.ListenBroadcastMessage()

	for {
		// accept
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept err:", err)
			continue
		}

		// do handler
		go this.Handler(conn)
	}
}
