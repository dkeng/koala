package packet

import (
	"bytes"
	"encoding/binary"
	"errors"

	"github.com/dkeng/pkg/convert"
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

// GetHead 获取head
func (p *Packet) GetHead() *Head {
	return p.head
}

// GetData 获取data
func (p *Packet) GetData() []byte {
	return p.data
}

// NewSend 创建发送包
func NewSend(data []byte) (*Packet, error) {
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

// NewReceive 创建接受包
func NewReceive(data []byte) (*Packet, error) {
	lenght := len(data)
	if lenght == 0 {
		return nil, errors.New("data is null")
	}
	if lenght < 5 {
		return nil, errors.New("data is error")
	}
	dataLenght := convert.BytesToInt(data[1:5])
	return &Packet{
		head: &Head{
			Type:       data[0],
			DataLenght: dataLenght,
		},
		data: data[5 : 5+dataLenght],
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
