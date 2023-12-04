package Interpreter

import "testing"

func TestNewAlterRuleInterpreter(t *testing.T) {
	expression := NewAlterRuleInterpreter("value1 > 100&&value2 < 24")
	stats := map[string]int{}
	stats["value1"] = 200
	stats["value2"] = 20
	t.Log(expression.Interpret(stats))
}
