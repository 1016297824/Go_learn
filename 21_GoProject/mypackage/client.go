// client
package mypackage

import "net"

type Client struct {
	Name          string
	Addr          string
	ClientMessage chan string
	conn          net.Conn
	server        *Server
}

func (this *Client) ListenClientMessage() {
	go func() {
		for {
			msg := <-this.ClientMessage
			this.conn.Write([]byte(msg + "\n"))
		}
	}()
}

// 用户上线处理
func (this *Client) OnlineClient() {

	// 将客户加入OnlineMap中
	this.server.mapLock.Lock()
	this.server.OnlineMap[this.Name] = *this
	this.server.mapLock.Unlock()

	// 广播当前客户端上线消息
	this.server.BroadcastMessageSend(*this, "已上线！")
}

// 用户下线处理
func (this *Client) OutlineClient() {
	// 将客户从OnlineMap中移除
	this.server.mapLock.Lock()
	delete(this.server.OnlineMap, this.Name)
	this.server.mapLock.Unlock()

	// 广播当前客户端下线消息
	this.server.BroadcastMessageSend(*this, "已下线！")
	this.conn.Close()
}

// 用户消息处理
func (this *Client) DoClientMsg(msg string) {
	// 默认广播消息
	this.server.BroadcastMessageSend(*this, msg)
}
