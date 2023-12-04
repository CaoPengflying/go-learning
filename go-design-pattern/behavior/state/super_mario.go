package state

type SuperMario struct {
	sateMachine *MarioSateMachine
}

func (s *SuperMario) GetState() State {
	return Super
}
func (s *SuperMario) ObtainMushRoom() {
	s.sateMachine.score += 100
}
func (s *SuperMario) ObtainCape() {
	s.sateMachine.currentState = &CapeMario{sateMachine: s.sateMachine}
	s.sateMachine.score += 200
}
func (s *SuperMario) ObtainFireFlower() {
	s.sateMachine.currentState = &FireMario{sateMachine: s.sateMachine}
	s.sateMachine.score += 300
}
func (s *SuperMario) MeetMonster() {
	s.sateMachine.currentState = &SmallMario{sateMachine: s.sateMachine}
}
