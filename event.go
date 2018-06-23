package server

// LogEvent log 事件
type LogEvent func(interface{})

// AcceptEvent 接受事件
type AcceptEvent func(interface{})

// ReceiveEvent 收到事件
type ReceiveEvent func(interface{})

// ClientConnCloseEvent 客户端连接关闭事件
type ClientConnCloseEvent func(interface{})
