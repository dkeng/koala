package packet

import (
	"fmt"
)

// Packet 包
type Packet struct {
	// Head 头
	Head Head
	// Type 类型
	Type Type
	// Data 数据
	Data []byte
	// DataLenght 数据长度
	DataLenght int
}

// New 创建包
func New([]byte) *Packet {
	return &Packet{}
}

// News 创建多个包
func News([]byte) []*Packet {
	return nil
}

// NewQrs 创建一个心跳包
func NewQrs(lenght int) (*Packet, error) {
	if lenght < 2 {
		return nil, fmt.Errorf("qrs packet lenght:%d,error", lenght)
	}
	data := make([]byte, lenght-2)
	return &Packet{
		Head:       HeadClosure,
		Type:       TypeQrs,
		Data:       data,
		DataLenght: 0,
	}, nil
}

// NewClose 创建一个关闭包
func NewClose(lenght int) (*Packet, error) {
	if lenght < 2 {
		return nil, fmt.Errorf("close packet lenght:%d,error", lenght)
	}
	data := make([]byte, lenght-2)
	return &Packet{
		Head:       HeadClosure,
		Type:       TypeClose,
		Data:       data,
		DataLenght: 0,
	}, nil
}

// Bytes 字节
func (p *Packet) Bytes() []byte {
	return nil
}
