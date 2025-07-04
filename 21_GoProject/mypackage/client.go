// Client
package mypackage

import (
	"fmt"
	"io"
	"net"
	"os"
)

type Client struct {
	ServerIP   string
	ServerPort int
	Name       string
	conn       net.Conn
	Flag       int
}

// 连接server
func (this *Client) CreateConn() {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", this.ServerIP, this.ServerPort))
	if err != nil {
		fmt.Println("net.Dial err:", err)
		return
	}

	this.conn = conn
	fmt.Println("创建连接成功！")
}

// 客户端菜单
func (this *Client) menu() bool {
	fmt.Println("1.公聊模式")
	fmt.Println("2.私聊模式")
	fmt.Println("3.更改用户名")
	fmt.Println("0.退出登录")
	fmt.Printf("请选择功能：")

	var flag int
	fmt.Scan(&flag)
	if flag < 0 || flag > 3 {
		fmt.Println("请输入正确数字")
		return false
	} else {
		this.Flag = flag
		return true
	}
}

// 公聊模式
func (this *Client) PublicChat() {
	fmt.Println("请输入聊天内容，输入exit退出！")
	var chatMsg string
	for {
		fmt.Scan(&chatMsg)

		if len(chatMsg) == 0 {
			continue
		}

		if chatMsg == "exit" {
			return
		}

		_, err := this.conn.Write([]byte(chatMsg))
		if err != nil {
			fmt.Println("conn.Write err", err)
			return
		}
		chatMsg = ""
	}

}

// 私聊模式
func (this *Client) PrivateChat() {
	// 查询在线用户
	selectMsg := "who"
	_, err := this.conn.Write([]byte(selectMsg))
	if err != nil {
		fmt.Println("conn.Write err", err)
		return
	}

	var remoteName string
	for {
		fmt.Println("请输入目标用户名：")
		fmt.Scan(&remoteName)
		fmt.Println("请输入聊天内容,输入exit退出")
		var remoteMsg string
		for {
			fmt.Scan(&remoteMsg)

			if len(remoteMsg) == 0 {
				continue
			}

			if remoteMsg == "exit" {
				return
			}

			remoteMsg = "to|" + remoteName + "|" + remoteMsg + "\n"
			_, err := this.conn.Write([]byte(remoteMsg))
			if err != nil {
				fmt.Println("conn.Write err", err)
				return
			}
			remoteMsg = ""
		}
	}
}

// 修改用户名
func (this *Client) ChangeName() {
	fmt.Printf("请输入用户名：")
	fmt.Scan(&this.Name)

	sendMsg := "rename|" + this.Name
	_, err := this.conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println("conn.Write err", err)
	}
}

// 处理服务器返回的消息
func (this *Client) DoResponse() {
	// 循环打印消息到终端，否则永久阻塞
	io.Copy(os.Stdout, this.conn)

	// 以上语句相当于↓
	// for {
	// 	buf := make([]byte, 1024)
	// 	this.conn.Read(buf[:])
	// 	fmt.Println(buf)
	// }
}

// 启动客户端服务
func (this *Client) Run() {
	// 结束关闭连接
	defer this.conn.Close()

	for this.Flag != 0 {
		for this.menu() == false {
		}

		switch this.Flag {
		case 1:
			fmt.Println("开启公聊模式")
			this.PublicChat()
			break
		case 2:
			fmt.Println("开启私聊模式")
			this.PrivateChat()
			break
		case 3:
			fmt.Println("更改用户名")
			this.ChangeName()
			break
		case 0:
			fmt.Println("退出登录")
			break
		default:
			continue
		}
	}
}
