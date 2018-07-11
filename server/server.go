package server

import (
	"bytes"
	"io"
	"net"
	"time"

	"github.com/dkeng/koala/errors"
	"github.com/dkeng/koala/packet"
	"github.com/dkeng/koala/server/client"
	"github.com/dkeng/koala/server/event"
	"github.com/dkeng/pkg/convert"
)

// KoalaServer 考拉服务器
type KoalaServer struct {
	// 地址
	Addr string
	// 客户端
	clients  map[string]*client.Client
	listener net.Listener
	// 日志事件
	LogEvent event.LogEvent
	// 接受事件
	AcceptEvent event.AcceptEvent
	// 收到客户端消息
	ReceiveEvent event.ReceiveEvent
	// 客户端连接关闭事件
	ClientConnCloseEvent event.ClientConnCloseEvent
	// 缓冲区大小 默认1024
	BufferLength int
	// 服务器状态
	watch bool
}

// accept 接受
func (k *KoalaServer) accept(conn net.Conn) {
	clientID := conn.RemoteAddr().String()
	// 判断客户端是否存在
	if _, ok := k.clients[clientID]; ok {
		k.LogEvent("error", errors.ErrClientExist)
	} else {
		k.clients[clientID] = client.New(clientID, conn)
		k.AcceptEvent(clientID)
		go k.receive(clientID)
	}
}

// receive 收到客户端消息
func (k *KoalaServer) receive(clientID string) {
	// 关闭
	defer func(cid string) {
		k.CloseClient(cid)
	}(clientID)
	// 缓冲区
	buffer := new(bytes.Buffer)
	// 数据包长度
	dataLenght := 0
	// 是否是第一次
	isFirst := true
	conn := k.clients[clientID].Connection
	for {
		// 临时缓冲区
		readBuffer := make([]byte, k.BufferLength)
		n, err := conn.Read(readBuffer)
		if err != nil {
			if err == io.EOF {
				k.CloseClient(clientID)
				k.ClientConnCloseEvent(clientID)
				break
			}
			k.LogEvent("error", err)
			break
		}
		if n < 5 {
			k.LogEvent("error", errors.ErrReadDataLenght)
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

// task 任务
func (k *KoalaServer) task(clientID string, p *packet.Packet) {
	switch p.GetHead().Type {
	case packet.TypeClose:
		// 关闭包
		k.CloseClient(clientID)
	case packet.TypeQrs:
		// 心跳包
		k.clients[clientID].LastHeartbeatTime = time.Now()
	case packet.TypeData:
		// 数据包
		k.ReceiveEvent(p.GetData())
	}
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
		clients:      make(map[string]*client.Client, 0),
	}

	return koalaServer, nil
}

// ListenAndServe 监听和服务
func (k *KoalaServer) ListenAndServe() {
	k.watch = true
	for k.watch {
		conn, err := k.listener.Accept()
		if err != nil {
			k.LogEvent("error", err)
			continue
		}
		go k.accept(conn)
	}
}

// CloseClient 关闭客户端
func (k *KoalaServer) CloseClient(clientID string) error {
	if _, ok := k.clients[clientID]; ok {
		err := k.clients[clientID].Connection.Close()
		delete(k.clients, clientID)
		if err != nil {
			return err
		}
	} else {
		return errors.ErrClientNotFound
	}
	return nil
}

// GetClientCount 获取客户端数量
func (k *KoalaServer) GetClientCount() int {
	return len(k.clients)
}

// SendClient 发送消息给客户端
func (k *KoalaServer) SendClient(clientID string, p *packet.Packet) error {
	if _, ok := k.clients[clientID]; ok {
		k.clients[clientID].Send(p)
	} else {
		return errors.ErrClientNotFound
	}
	return nil
}

// Close 关闭
func (k *KoalaServer) Close() error {
	// 停止服务器
	k.watch = false
	// 关闭所有客户端
	for key := range k.clients {
		k.CloseClient(key)
	}
	// 关闭监听
	err := k.listener.Close()
	if err != nil {
		return err
	}
	return nil
}

// Heartbeat 心跳
func (k *KoalaServer) Heartbeat() {
	// 心跳返回
	span := 5 * time.Minute
	// 延时
	sleep := 5 * time.Minute
	// 默认心跳包
	defaultQrs := packet.NewQrs()

	go func(sp time.Duration, sleep time.Duration, qrs *packet.Packet) {
		for k.watch {
			// 延时
			time.Sleep(sleep)
			// 发送心跳包
			for clientID, val := range k.clients {
				end := time.Now()
				span := end.Sub(val.LastHeartbeatTime)
				if span > sp {
					k.SendClient(clientID, qrs)
				}
			}
		}
	}(span, sleep, defaultQrs)
}
