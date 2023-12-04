// @Description channel
// @Author caopengfei 2022/4/20 10:36

package channel

import (
	"log"
	"testing"
	"time"
)

type AcceptMsg struct {
	msg string
}

type Server struct {
	netServer NetServer

	acceptChan chan *AcceptMsg
}

type NetServer struct {
	acceptChan chan *AcceptMsg
}

func (s *Server) run() {
	for {
		select {
		case m := <-s.acceptChan:
			log.Println(m)
		}
	}
}

func TestFun(t *testing.T) {
	s := Server{}
	acceptChan := make(chan *AcceptMsg, 100)
	s.acceptChan = acceptChan
	netServer := NetServer{
		s.acceptChan,
	}
	s.netServer = netServer

	go func() {
		s.run()
	}()

	netServer.acceptChan <- &AcceptMsg{msg: "test"}
	time.Sleep(1000 * 10)
}
