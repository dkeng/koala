package server

// LogEvent log 事件
type LogEvent func(string, interface{})

// AcceptEvent 接受事件
type AcceptEvent func(interface{})

// ReceiveEvent 收到事件
type ReceiveEvent func([]byte)

// ClientConnCloseEvent 客户端连接关闭事件
type ClientConnCloseEvent func(interface{})
