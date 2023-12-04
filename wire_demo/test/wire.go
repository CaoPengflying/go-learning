// @Description test
// @Author caopengfei 2021/6/8 20:53
//+build wireinject

package test

import (
	"github.com/google/wire"
	"wire_demo/dao"
	"wire_demo/service"
)

var Set2 = wire.NewSet(dao.NewMockDemoDao, wire.Bind(new(dao.DemoDao), new(*dao.MockDemoDao)))

func InitDemoService() *service.DemoServiceImpl {
	wire.Build(service.NewDemoServiceImpl, Set2)
	return nil
}
