// server
package mypackage

import (
	"fmt"
	"io"
	"net"
	"sync"
)

// var SynMsg chan string = make(chan string)

type Server struct {
	Id   string
	Port int

	// 在线用户列表
	OnlineMap map[string]Client
	mapLock   sync.RWMutex

	// 消息广播的channel
	BroadcastMessage chan string
}

// 广播消息
func (this *Server) BroadcastMessageSend(client Client, msg string) {
	sendMsg := "[" + client.Addr + "]" + ":" + msg
	this.BroadcastMessage <- sendMsg
}

// handler function
func (this *Server) Handler(conn net.Conn) {
	// 当前连接的业务
	clientAddr := conn.RemoteAddr().String()
	fmt.Println("连接建立成功：", clientAddr)

	// 客户端Online
	client := Client{
		Name:          clientAddr,
		Addr:          clientAddr,
		ClientMessage: make(chan string),
		conn:          conn,
		server:        this,
	}
	client.OnlineClient()

	// 保证当前客户不会收到广播消息
	// onlineEndMsg := <-SynMsg
	// fmt.Println(onlineEndMsg)

	// 启动客户端广播监听
	client.ListenClientMessage()

	// 创建向用户消息处理的进程
	go func() {
		buf := make([]byte, 1024)
		for {
			n, err := conn.Read(buf)
			if n == 0 {
				// this.BroadcastMessageSend(client, "已下线！")
				client.OutlineClient()
				fmt.Println("断开与[", client.Name, "]的连接")
				return
			}

			if err != nil && err != io.EOF {
				fmt.Println("conn.Read err:", err)
				client.OutlineClient()
				fmt.Println("断开与[", client.Name, "]的连接")
				return
			}

			msg := string(buf[:n])
			client.DoClientMsg(msg)
		}
	}()

	// 阻塞当前Handler
	select {
	// case msg := <-SynMsg:
	// 	if msg == "err" || msg == "end" {
	// 		// 结束关闭连接
	// 		conn.Close()
	// 		return
	// 	}
	}
}

// 监听广播消息
func (this *Server) ListenBroadcastMessage() {
	for {
		msg := <-this.BroadcastMessage

		// 将消息发送给在线的客户端
		this.mapLock.Lock()
		for _, client := range this.OnlineMap {
			client.ClientMessage <- msg
		}
		this.mapLock.Unlock()

		// 对比是否包含 <"已上线">
		// if strings.Contains(msg, "已上线") {
		// 	SynMsg <- "客户端上线广播完成"
		// }
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
