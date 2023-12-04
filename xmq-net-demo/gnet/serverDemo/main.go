package main

import (
	"context"
	"fmt"
	"github.com/cloudwego/netpoll"
)

func main() {
	network, address := "tcp", "127.0.0.1:8888"

	// 创建 listener
	listener, err := netpoll.CreateListener(network, address)
	if err != nil {
		panic("create netpoll listener fail")
	}

	// handle: 连接读数据和处理逻辑
	var onRequest netpoll.OnRequest = handler

	// 创建 EventLoop
	eventLoop, err := netpoll.NewEventLoop(onRequest)
	if err != nil {
		panic("create netpoll event-loop fail")
	}

	// 运行 Server
	err = eventLoop.Serve(listener)
	if err != nil {
		panic("netpoll server exit")
	}
}

// 读事件处理
func handler(ctx context.Context, connection netpoll.Connection) error {
	reader := connection.Reader()
	bts, err := reader.Next(reader.Len())
	if err != nil {
		return err
	}
	fmt.Println(string(bts))

	connection.Write([]byte("success 1"))
	return connection.Writer().Flush()
}
