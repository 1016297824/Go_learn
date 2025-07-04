// Client
package mypackage

import (
	"fmt"
	"net"
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

func (this *Client) Run() {
	// 启动客户端服务
	for this.Flag != 0 {
		for this.menu() == false {
		}

		switch this.Flag {
		case 1:
			fmt.Println("开启公聊模式")
		case 2:
			fmt.Println("开启私聊模式")
		case 3:
			fmt.Println("更改用户名")
		case 0:
			fmt.Println("退出登录")
		default:
			continue
		}
	}
}
