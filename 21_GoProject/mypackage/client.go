// client
package mypackage

import "net"

type Client struct {
	Name          string
	Addr          string
	ClientMessage chan string
	conn          net.Conn
}

func (this *Client) ListenClientMessage() {
	go func() {
		for {
			msg := <-this.ClientMessage
			this.conn.Write([]byte(msg + "\n"))
		}
	}()
}
