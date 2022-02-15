// @Description once
// @Author caopengfei 2021/2/23 15:10
package main

import (
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
)

type Conn struct {
	flag int
}

func (c *Conn) DoSomething() {
	fmt.Println("conn's flag", c.flag)
}
func main() {
	m := new(MuOnce)
	fmt.Println(m.strings())
	fmt.Println(m.strings())
}

func GetConn() (*Conn, error) {
	fmt.Println("获取连接")
	return nil, errors.New("获取失败")
}

type once struct {
	done uint32
	sync.Mutex
}

func (o *once) Do(f func() error) error {
	if atomic.LoadUint32(&o.done) == 1 {
		return nil
	}
	return o.slowDo(f)
}

func (o *once) slowDo(f func() error) error {
	o.Lock()
	defer o.Unlock()

	var err error
	if o.done == 0 {
		err = f()
		if err == nil {
			atomic.StoreUint32(&o.done, 1)
		}
	}
	return err
}

type Once struct {
	sync.Once
}

func (o *Once) Done() bool {
	return atomic.LoadUint32((*uint32)(unsafe.Pointer(&o.Once))) == 1
}

type MuOnce struct {
	sync.Once
	sync.RWMutex
	mtime time.Time
	vals  []string
}

func (m *MuOnce) refresh() {
	m.Lock()
	defer m.Unlock()

	m.Once = sync.Once{}
	m.mtime = time.Now()
	m.vals = []string{m.mtime.String()}
}

//超过时间就重新初始化
func (m *MuOnce) strings() []string {
	now := time.Now()
	m.Lock()
	if now.After(m.mtime) {
		defer m.Do(m.refresh)
	}
	vals := m.vals
	m.Unlock()
	return vals
}
