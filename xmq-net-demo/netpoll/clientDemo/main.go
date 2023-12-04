package main

import (
	"context"
	"fmt"
	"time"

	"github.com/cloudwego/netpoll"
)

func main() {
	network, address := "tcp", "127.0.0.1:9000"

	// 通过 dialer 创建连接
	dialer := netpoll.NewDialer()
	conn, err := dialer.DialConnection(network, address, 50*time.Millisecond)
	if err != nil {
		panic("dialer netpoll connection fail")
	}

	// 设置读事件回调
	conn.SetOnRequest(callback())

	// conn write & flush message
	conn.Writer().WriteBinary([]byte("hello world"))
	conn.Writer().Flush()

	time.Sleep(1000 * time.Microsecond)
}

func callback() netpoll.OnRequest {
	return func(ctx context.Context, connection netpoll.Connection) error {
		reader := connection.Reader()
		bts, err := reader.Next(reader.Len())
		if err != nil {
			return err
		}
		fmt.Println(string(bts))
		fmt.Println("callback")
		return nil
	}
}
