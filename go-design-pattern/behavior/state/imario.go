package state

import "fmt"

type IMario interface {
	GetState() State
	ObtainMushRoom()
	ObtainCape()
	ObtainFireFlower()
	MeetMonster()
}

//状态枚举
type State int

const (
	Small State = iota
	Super
	Cape
	Fire
)

type MarioSateMachine struct {
	currentState IMario
	score        int
}

func (s *MarioSateMachine) GetCurrentState() State {
	return s.currentState.GetState()
}
func (s *MarioSateMachine) ObtainMushRoom() {
	s.currentState.ObtainMushRoom()
}
func (s *MarioSateMachine) ObtainCape() {
	s.currentState.ObtainCape()
}
func (s *MarioSateMachine) ObtainFireFlower() {
	s.currentState.ObtainFireFlower()
}
func (s *MarioSateMachine) MeetMonster() {
	s.currentState.MeetMonster()
}
func (s *MarioSateMachine) Print() {
	fmt.Println("state:", s.currentState.GetState())
	fmt.Println("score:", s.score)
}

func NewMarioStateMachine() *MarioSateMachine {
	machine := &MarioSateMachine{

	}
	machine.score = 0
	machine.currentState = &SmallMario{
		sateMachine: machine,
	}
	return machine
}
