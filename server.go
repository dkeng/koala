package server

import (
	"net"

	"github.com/dkeng/koala-mini-server/client"
)

// KoalaServer 考拉服务器
type KoalaServer struct {
	// 地址
	Addr string
	// 客户端
	Clients  map[string]*client.Client
	listener net.Listener

	Log             LogEvent
	accept          AcceptEvent
	receive         ReceiveEvent
	clientConnClose ClientConnCloseEvent
	BufferLength    int
}

// New 创建服务器
func New(addr string) (*KoalaServer, error) {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}
	return &KoalaServer{
		listener:     l,
		BufferLength: 1024,
		Clients:      make(map[string]*client.Client, 0),
	}, nil
}

// ListenAndServe 监听和服务
func (k *KoalaServer) ListenAndServe() {
	for {
		c, err := k.listener.Accept()
		if err != nil {
			continue
		}
		id := c.LocalAddr().String()
		k.Clients[id] = client.New(id, &c)
	}
}

// Close 关闭
func (k *KoalaServer) Close() {

}

// Heartbeat 心跳
func (k *KoalaServer) Heartbeat() {

}
