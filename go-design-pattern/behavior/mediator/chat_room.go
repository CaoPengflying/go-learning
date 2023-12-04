package mediator

import "fmt"

type User struct {
	name string
}

func (u *User) SetName(name string) {
	u.name = name
}
func (u *User) GetName() string {
	return u.name
}

func (u *User) SendMessage(message string) {
	chatRoom := new(ChatRoom)
	chatRoom.ShowMessage(u, message)
}

type ChatRoom struct {
}

func (c *ChatRoom) ShowMessage(user *User, message string) {
	fmt.Println(user.GetName() + ":" + message)
}
