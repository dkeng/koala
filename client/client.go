package client

import (
	"net"
	"time"

	"github.com/dkeng/koala-mini-server/packet"
)

// Client 客户端
type Client struct {
	// ID
	ID string
	// 连接
	Connection net.Conn
	// 最后心跳时间
	LastHeartbeatTime time.Time
}

// New 创建客户端
func New(id string, conn *net.Conn) *Client {
	return &Client{
		ID:                id,
		Connection:        *conn,
		LastHeartbeatTime: time.Now(),
	}
}

// Send 发送
func (c *Client) Send(p *packet.Packet) {
	c.Connection.Write(p.Bytes())
}

// Sends 发送多个
func (c *Client) Sends(ps []*packet.Packet) {
	for i := 0; i < len(ps); i++ {
		c.Send(ps[i])
	}
}
