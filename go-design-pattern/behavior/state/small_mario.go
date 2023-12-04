package state

type SmallMario struct {
	sateMachine *MarioSateMachine
}

func (s *SmallMario) GetState() State {
	return Small
}
func (s *SmallMario) ObtainMushRoom() {
	s.sateMachine.currentState = &SuperMario{sateMachine: s.sateMachine}
	s.sateMachine.score += 100
}
func (s *SmallMario) ObtainCape() {
	s.sateMachine.currentState = &CapeMario{sateMachine: s.sateMachine}
	s.sateMachine.score += 200
}
func (s *SmallMario) ObtainFireFlower() {
	s.sateMachine.currentState = &FireMario{sateMachine: s.sateMachine}
	s.sateMachine.score += 300
}
func (s *SmallMario) MeetMonster() {
	panic("game over")
}
