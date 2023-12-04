package state

type CapeMario struct {
	sateMachine *MarioSateMachine
}

func (s *CapeMario) GetState() State {
	return Cape
}
func (s *CapeMario) ObtainMushRoom() {
	s.sateMachine.score += 100
}
func (s *CapeMario) ObtainCape() {
	s.sateMachine.score += 200
}
func (s *CapeMario) ObtainFireFlower() {
	s.sateMachine.currentState = &FireMario{sateMachine: s.sateMachine}
	s.sateMachine.score += 300
}
func (s *CapeMario) MeetMonster() {
	s.sateMachine.currentState = &SuperMario{sateMachine: s.sateMachine}

}
