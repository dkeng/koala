package main

import (
	"fmt"
	"os"

	server "github.com/dkeng/koala/server"
)

var (
	koalaServer *server.KoalaServer
)

func main() {
	koalaServer, err := server.New(":8989")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	koalaServer.AcceptEvent = func(value interface{}) {
		fmt.Printf("AcceptEvent:\n%v\n", value)
	}
	koalaServer.LogEvent = func(tp string, value interface{}) {
		fmt.Printf("Log-%s:\n%v\n", tp, value)
	}
	koalaServer.ClientConnCloseEvent = func(value interface{}) {
		fmt.Printf("ClientConnCloseEvent:\n%v\n", value)
	}
	koalaServer.ReceiveEvent = func(value []byte) {
		fmt.Printf("ReceiveEvent:\n%v\n", string(value))
		fmt.Printf("客户端数量:\n%v\n", koalaServer.GetClientCount())
	}
	koalaServer.ListenAndServe()
}
