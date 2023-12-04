// @Description service
// @Author caopengfei 2021/6/8 20:35

package service

type IUserService interface {
	GetUserById(id int) string
}

type UserServiceImpl struct {
}

func (userService *UserServiceImpl) GetUserById(id int) string {
	return "user"
}

func NewUserServiceImpl() *UserServiceImpl {
	return &UserServiceImpl{}
}
