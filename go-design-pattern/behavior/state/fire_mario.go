package state

type FireMario struct {
	sateMachine *MarioSateMachine
}

func (s *FireMario) GetState() State {
	return Fire
}
func (s *FireMario) ObtainMushRoom() {
	s.sateMachine.score += 100
}
func (s *FireMario) ObtainCape() {
	s.sateMachine.score += 200
}
func (s *FireMario) ObtainFireFlower() {
	s.sateMachine.score += 300
}
func (s *FireMario) MeetMonster() {
	s.sateMachine.currentState = &CapeMario{sateMachine: s.sateMachine}
}