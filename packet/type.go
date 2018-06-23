package packet

// Type 包类型
type Type int

const (
	// TypeQrs 心跳包
	TypeQrs Type = 0
	// TypeData 数据包
	TypeData Type = 1
	// TypeClose 关闭包
	TypeClose Type = 2
)
