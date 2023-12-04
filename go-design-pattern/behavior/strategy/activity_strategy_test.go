package strategy

import "testing"

func TestExecute(t *testing.T) {
	factory := ActivityStrategyFactory{}
	activityStrategy, _ := factory.GetStrategy("gdzk")
	activityStrategy.execute()
}
