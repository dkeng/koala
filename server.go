package server

import (
	"bytes"
	"io"
	"net"
	"time"

	"github.com/dkeng/koala-mini-server/client"
	"github.com/dkeng/koala-mini-server/packet"
	"github.com/dkeng/pkg/convert"
)

// KoalaServer 考拉服务器
type KoalaServer struct {
	// 地址
	Addr string
	// 客户端
	Clients  map[string]*client.Client
	listener net.Listener

	LogEvent             LogEvent
	AcceptEvent          AcceptEvent
	ReceiveEvent         ReceiveEvent
	ClientConnCloseEvent ClientConnCloseEvent
	BufferLength         int
}

// New 创建服务器
func New(addr string) (*KoalaServer, error) {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}

	koalaServer := &KoalaServer{
		listener:     l,
		BufferLength: 1024,
		Clients:      make(map[string]*client.Client, 0),
	}

	return koalaServer, nil
}

// ListenAndServe 监听和服务
func (k *KoalaServer) ListenAndServe() {
	for {
		conn, err := k.listener.Accept()
		if err != nil {
			continue
		}
		go k.accept(conn)
	}
}

// accept 接受
func (k *KoalaServer) accept(conn net.Conn) {
	id := conn.RemoteAddr().String()
	k.Clients[id] = client.New(id, conn)
	k.AcceptEvent(id)
	go k.receive(id)
}

// receive 收到客户端消息
func (k *KoalaServer) receive(clientID string) {
	// 缓冲区
	buffer := new(bytes.Buffer)
	// 数据包长度
	dataLenght := 0
	// 是否是第一次
	isFirst := true
	conn := k.Clients[clientID].Connection
	for {
		// 临时缓冲区
		readBuffer := make([]byte, k.BufferLength)
		n, err := conn.Read(readBuffer)
		if err != nil {
			if err == io.EOF {
				delete(k.Clients, clientID)
				k.ClientConnCloseEvent(clientID)
				break
			}
			k.LogEvent("error", err)
			break
		}
		if n < 5 {
			k.LogEvent("error", "数据错误，长度小于5")
			break
		}
		if isFirst {
			dataLenght = packet.DefaultHeadLenght + convert.BytesToInt(readBuffer[1:packet.DefaultHeadLenght])
			isFirst = false
		}
		for i, v := range readBuffer {
			if i > n {
				break
			}
			buffer.WriteByte(v)

			if buffer.Len() == dataLenght {
				p, err := packet.NewReceive(buffer.Bytes())
				if err != nil {
					k.LogEvent("error", err)
				} else {
					k.task(clientID, p)
				}
				// 初始化 新包的长度
				dataLenghtStart := i + 2
				dataLenghtEnd := dataLenghtStart + packet.DefaultHeadLenght - 1
				dataLenght = packet.DefaultHeadLenght + convert.BytesToInt(readBuffer[dataLenghtStart:dataLenghtEnd])
				// 重置缓冲区
				buffer.Reset()

			}
		}
	}
}

// Close 关闭
func (k *KoalaServer) Close() {

}

// task 任务
func (k *KoalaServer) task(clientID string, p *packet.Packet) {
	switch p.GetHead().Type {
	case packet.TypeClose:
		// 关闭包
		delete(k.Clients, clientID)
	case packet.TypeQrs:
		// 心跳包
		k.Clients[clientID].LastHeartbeatTime = time.Now()
	case packet.TypeData:
		k.ReceiveEvent(p.GetData())
	}
}

// Heartbeat 心跳
func (k *KoalaServer) Heartbeat() {

}
