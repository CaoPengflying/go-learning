// @Description controller
// @Author caopengfei 2021/6/3 17:33

package controller

import (
	"wire_demo/service"
)

type DemoController struct {
	demoService service.DemoService
	userService service.IUserService
}

func (demo *DemoController) Demo() string {
	return demo.demoService.Demo() + demo.userService.GetUserById(1)
}

func NewDemoController(demoService service.DemoService, userService service.IUserService) *DemoController {

	return &DemoController{demoService: demoService,userService: userService}
}
