// server
package mypackage

import (
	"fmt"
	"io"
	"net"
	"runtime"
	"sync"
	"time"
)

// var SynMsg = make(chan string)

type Server struct {
	Id   string
	Port int

	// 在线用户列表
	OnlineMap map[string]User
	mapLock   sync.RWMutex

	// 消息广播的channel
	BroadcastMessage chan string
}

// 广播消息
func (this *Server) BroadcastMessageSend(User User, msg string) {
	sendMsg := "[" + User.Addr + "]" + ":" + msg
	this.BroadcastMessage <- sendMsg
}

// handler function
func (this *Server) Handler(conn net.Conn) {
	// 当前连接的业务
	UserAddr := conn.RemoteAddr().String()
	fmt.Println("连接建立成功：", UserAddr)

	// conn close
	defer conn.Close()

	// 客户端Online
	User := User{
		Name:        UserAddr,
		Addr:        UserAddr,
		UserMessage: make(chan string),
		conn:        conn,
		server:      this,
	}
	User.OnlineUser()

	// 保证当前客户不会收到广播消息
	// onlineEndMsg := <-SynMsg
	// fmt.Println(onlineEndMsg)

	// 启动客户端广播监听
	User.ListenUserMessage()

	alive := make(chan bool)

	// 创建向用户消息处理的进程
	go func() {
		buf := make([]byte, 1024)
		for {
			n, err := conn.Read(buf)
			if n == 0 {
				// this.BroadcastMessageSend(User, "已下线！")
				User.OutlineUser()
				fmt.Println("断开与[", User.Addr, "]的连接")
				return
			}

			if err != nil && err != io.EOF {
				fmt.Println("conn.Read err:", err)
				User.OutlineUser()
				fmt.Println("断开与[", User.Addr, "]的连接")
				return
			}

			msg := string(buf[:n])
			User.DoUserMsg(msg)

			alive <- true
		}
	}()

	// 阻塞当前Handler
	for {
		select {
		// case msg := <-SynMsg:
		// 	if msg == "err" || msg == "end" {
		// 		// 结束关闭连接
		// 		conn.Close()
		// 		return
		// 	}
		// 阻塞for循环，当alive接收到消息时结束阻塞
		case <-alive:
			// 什么都不做，进入下一个for循环

		// 每次进入for循环计时器都会重置
		// 当计时器时间到点时，触发阻塞
		case <-time.After(300 * time.Second):
			// 向客户端发送踢出消息
			User.SendMsg("等待超时，你已被踢出！！！")

			// 销毁客户端资源
			close(User.UserMessage)

			// 关闭连接
			conn.Close()

			// 退出当前handler routine
			runtime.Goexit() //或者直接：return
		}
	}
}

// 监听广播消息
func (this *Server) ListenBroadcastMessage() {
	for {
		msg := <-this.BroadcastMessage

		// 将消息发送给在线的客户端
		this.mapLock.Lock()
		for _, User := range this.OnlineMap {
			User.UserMessage <- msg
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
	// channel close
	defer close(this.BroadcastMessage)

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
