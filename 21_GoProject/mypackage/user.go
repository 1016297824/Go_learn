// User
package mypackage

import (
	"net"
	"strings"
)

type User struct {
	Name        string
	Addr        string
	UserMessage chan string
	conn        net.Conn
	server      *Server
}

func (this *User) ListenUserMessage() {
	go func() {
		for {
			msg := <-this.UserMessage
			this.conn.Write([]byte(msg + "\n"))
		}
	}()
}

// 用户上线处理
func (this *User) OnlineUser() {
	// 将客户加入OnlineMap中
	this.server.mapLock.Lock()
	this.server.OnlineMap[this.Addr] = *this
	this.server.mapLock.Unlock()

	// 广播当前客户端上线消息
	this.server.BroadcastMessageSend(*this, "已上线！")
}

// 用户下线处理
func (this *User) OutlineUser() {
	// 将客户从OnlineMap中移除
	this.server.mapLock.Lock()
	delete(this.server.OnlineMap, this.Addr)
	this.server.mapLock.Unlock()

	// 广播当前客户端下线消息
	this.server.BroadcastMessageSend(*this, "已下线！")
	this.conn.Close()
}

// 向用户发送消息
func (this *User) SendMsg(msg string) {
	this.conn.Write([]byte(msg))
}

// 用户消息处理
func (this *User) DoUserMsg(msg string) {
	if msg == "who" {
		sendMsg := "在线的用户：\n"
		this.server.mapLock.Lock()
		for _, value := range this.server.OnlineMap {
			if value.Addr != this.Addr {
				sendMsg = sendMsg + value.Name + "\n"
			}
		}
		this.server.mapLock.Unlock()
		this.SendMsg(sendMsg)
	} else if len(msg) >= 7 && msg[:7] == "rename|" {
		newName := msg[7:]
		// 判断名字是否存在
		exist := false
		for _, value := range this.server.OnlineMap {
			if value.Name == newName {
				exist = true
			}
		}

		if exist {
			this.SendMsg("用户名已存在")
		} else {
			this.Name = newName
			this.server.mapLock.Lock()
			this.server.OnlineMap[this.Addr] = *this
			this.server.mapLock.Unlock()
			this.SendMsg("用户名修改成功")
		}
	} else if len(msg) >= 2 && msg[:2] == "to" {
		// 获取目标客户名
		remotName := strings.Split(msg, "|")[1]
		if remotName == "" {
			this.SendMsg("请输入目标用户名")
			return
		}

		// 判断目标客户端是否存在
		exist := false
		remotUser := User{}
		for _, value := range this.server.OnlineMap {
			if remotName == value.Name {
				exist = true
				remotUser = value
			}
		}
		if exist {
			remotUser.SendMsg(strings.Split(msg, "|")[2])
		}
	} else {
		// 默认广播消息
		this.server.BroadcastMessageSend(*this, msg)
	}
}
