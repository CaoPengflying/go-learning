package facade

import "log"

type IUserService interface {
	Register(name string)
}

type UserServiceImpl struct {
}

func (u UserServiceImpl) Register(name string) {
	log.Print("user register")
}
