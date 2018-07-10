package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"

	"github.com/dkeng/koala-mini-server/packet"
)

const (
	str = "-abcdefghijklmnopqistuvwxyz|abcdefghijklmnopqistuvwxyz|abcdefghijklmnopqistuvwxyz|abcdefghijklmnopqistuvwxyz|abcdefghijklmnopqistuvwxyz-"
)

func main() {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:8989")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for i := 0; i < 3; i++ {
		p, err := packet.NewSend([]byte(str))
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(i)
		conn.Write(p.Bytes())
	}
	fmt.Println("发送成功")
	result, err := ioutil.ReadAll(conn)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(result))
}
