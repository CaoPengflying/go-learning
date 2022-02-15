// @Description pool
// @Author caopengfei 2021/2/26 16:04
package main

import (
	"bytes"
	"fmt"
	"sync"
	"time"
)

type Conn struct {
}

func main() {
	go func() {
		buffer := GetBuffer()
		buffer.WriteString("hello")
		PutBuffer(buffer)
	}()
	time.Sleep(time.Millisecond)
	go func() {
		for  {
			b := GetBuffer()
			fmt.Println(b.String())
		}
	}()
	select {

	}
}

func NewConn() interface{} {
	return Conn{}
}

var buffers = sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}

func GetBuffer() *bytes.Buffer {
	return buffers.Get().(*bytes.Buffer)
}
func PutBuffer(buf *bytes.Buffer) {
	buf.Reset()
	buffers.Put(buf)
}
