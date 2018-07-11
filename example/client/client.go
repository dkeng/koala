package main

import (
	"fmt"
	"net"
	"os"

	"github.com/dkeng/koala/packet"
)

const (
	str = "-abcdefghijklmnopqistuvwxyz|abcdefghijklmnopqistuvwxyz|abcdefghijklmnopqistuvwxyz|abcdefghijklmnopqistuvwxyz|abcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyz|abcdefghijklmnopqistuvwxyz|abcdefghijklmnopqistuvwxyz|abcdefghijklmnopqistuvwxyz|abcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyz|abcdefghijklmnopqistuvwxyz|abcdefghijklmnopqistuvwxyz|abcdefghijklmnopqistuvwxyz|abcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyz|abcdefghijklmnopqistuvwxyz|abcdefghijklmnopqistuvwxyz|abcdefghijklmnopqistuvwxyz|abcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyz|abcdefghijklmnopqistuvwxyz|abcdefghijklmnopqistuvwxyz|abcdefghijklmnopqistuvwxyz|abcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyz|abcdefghijklmnopqistuvwxyz|abcdefghijklmnopqistuvwxyz|abcdefghijklmnopqistuvwxyz|abcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyz|abcdefghijklmnopqistuvwxyz|abcdefghijklmnopqistuvwxyz|abcdefghijklmnopqistuvwxyz|abcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyz|abcdefghijklmnopqistuvwxyz|abcdefghijklmnopqistuvwxyz|abcdefghijklmnopqistuvwxyz|abcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyz|abcdefghijklmnopqistuvwxyz|abcdefghijklmnopqistuvwxyz|abcdefghijklmnopqistuvwxyz|abcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyz|abcdefghijklmnopqistuvwxyz|abcdefghijklmnopqistuvwxyz|abcdefghijklmnopqistuvwxyz|abcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyz|abcdefghijklmnopqistuvwxyz|abcdefghijklmnopqistuvwxyz|abcdefghijklmnopqistuvwxyz|abcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyz|abcdefghijklmnopqistuvwxyz|abcdefghijklmnopqistuvwxyz|abcdefghijklmnopqistuvwxyz|abcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyz|abcdefghijklmnopqistuvwxyz|abcdefghijklmnopqistuvwxyz|abcdefghijklmnopqistuvwxyz|abcdefghijklmnopqistuvwxyzabcdefghijklmnopqistuvwxyz|abcdefghijklmnopqistuvwxyz|abcdefghijklmnopqistuvwxyz|abcdefghijklmnopqistuvwxyz|abcdefghijklmnopqistuvwxyz-"
)

var (
	conns []*net.TCPConn
	cchan = make(chan int)
)

func main() {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:8989")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for i := 0; i < 1024; i++ {
		conn, err := net.DialTCP("tcp", nil, tcpAddr)
		if err != nil {
			fmt.Println(err)
		} else {
			conns = append(conns, conn)
		}
	}
	p, err := packet.NewSend([]byte(str))
	if err != nil {
		fmt.Println(err)
		return
	}
	for i, conn := range conns {
		conn.Write(p.Bytes())
		fmt.Println(i)
	}
	<-cchan
}
