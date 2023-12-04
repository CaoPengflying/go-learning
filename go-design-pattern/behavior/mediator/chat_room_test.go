package mediator

import "testing"

func TestChatRoom_ShowMessage(t *testing.T) {
	user := new(User)
	user.SetName("cpf")

	user2 := new(User)
	user2.SetName("zzc")

	user.SendMessage("早上好呀！")
	user2.SendMessage("早上好")
}

