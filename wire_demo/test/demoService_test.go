// @Description test
// @Author caopengfei 2021/6/8 20:52

package test

import "testing"

func TestDemo(t *testing.T) {
	service := InitDemoService()
	service.Demo()
}

