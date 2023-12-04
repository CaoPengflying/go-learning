package proxy

import "log"

type UserControllerProxy struct {
	userController *UserController
}

func (u *UserControllerProxy) login(loginName string, pwd string) bool {
	log.Println("proxy pre")
	result := u.userController.login(loginName, pwd)
	log.Println("proxy latter")
	return result
}
