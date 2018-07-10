package packet

const (
	// DefaultHeadLenght 默认头长度
	DefaultHeadLenght = 5
)

// Head 包头
type Head struct {
	// 数据类型
	Type byte
	// DataLenght 数据长度
	DataLenght int
}
