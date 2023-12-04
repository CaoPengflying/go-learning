package proxy

import (
	"log"
)

type UserController struct {
}

func (u *UserController) login(loginName string, pwd string) bool {
	log.Printf("%s's pwd is right %s's login success", loginName, pwd)
	return true
}
