package state

import "testing"

func TestNewMarioStateMachine(t *testing.T) {
	machine := NewMarioStateMachine()
	machine.Print()

	machine.ObtainMushRoom()
	machine.Print()

	machine.MeetMonster()
	machine.Print()

	machine.ObtainMushRoom()
	machine.Print()


	machine.ObtainCape()
	machine.Print()


	machine.ObtainFireFlower()
	machine.Print()

	machine.MeetMonster()
	machine.Print()

	machine.ObtainFireFlower()
	machine.Print()

	machine.ObtainCape()
	machine.Print()

	machine.MeetMonster()
	machine.Print()



}

