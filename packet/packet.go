package packet

import (
	"bytes"
	"encoding/binary"
	"errors"
)

// Packet 包
type Packet struct {
	// Head 头
	head *Head
	// Data 数据
	data []byte
}

// Bytes 获取数据
func (p *Packet) Bytes() []byte {
	buffer := &bytes.Buffer{}
	// 类型
	buffer.WriteByte(p.head.Type)
	// 长度
	binary.Write(buffer, binary.BigEndian, int32(p.head.DataLenght))
	// 数据
	if p.data != nil {
		buffer.Write(p.data)
	}
	return buffer.Bytes()
}

// New 创建包
func New(data []byte) (*Packet, error) {
	lenght := len(data)
	if lenght == 0 {
		return nil, errors.New("data is null")
	}
	return &Packet{
		head: &Head{
			Type:       TypeData,
			DataLenght: lenght,
		},
		data: data,
	}, nil
}

// NewQrs 创建一个心跳包
func NewQrs() *Packet {
	return &Packet{
		head: &Head{
			Type:       TypeQrs,
			DataLenght: 0,
		},
		data: nil,
	}
}

// NewClose 创建一个关闭包
func NewClose() *Packet {
	return &Packet{
		head: &Head{
			Type:       TypeClose,
			DataLenght: 0,
		},
		data: nil,
	}
}
