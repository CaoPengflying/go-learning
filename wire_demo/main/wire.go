//+build wireinject
// @Description wire_demo
// @Author caopengfei 2021/6/3 21:19

package main

import (
	"github.com/google/wire"
	"wire_demo/controller"
	"wire_demo/dao"
	"wire_demo/service"
)

var Set1 = wire.NewSet(service.NewDemoServiceImpl, wire.Bind(new(service.DemoService), new(*service.DemoServiceImpl)))
var Set2 = wire.NewSet(dao.NewDemoDaoImpl, wire.Bind(new(dao.DemoDao), new(*dao.DemoDaoImpl)))
var Set3 = wire.NewSet(service.NewUserServiceImpl, wire.Bind(new(service.IUserService), new(*service.UserServiceImpl)))

func InitController() *controller.DemoController {
	wire.Build(controller.NewDemoController, Set1, Set2, Set3)
	return nil
}
