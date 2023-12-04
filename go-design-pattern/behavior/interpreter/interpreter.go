package Interpreter

import (
	"errors"
	"strconv"
	"strings"
)

// 解释器模式

type Expression interface {
	Interpret(stats map[string]int) bool
}

type GreaterExpression struct {
	key   string
	value int
}

func (g *GreaterExpression) Interpret(stats map[string]int) bool {
	statValue, exists := stats[g.key]
	if !exists {
		return exists
	}
	return g.value < statValue
}

func NewGreaterExpression(strExpression string) (*GreaterExpression, error) {
	elements := strings.Split(strExpression, " ")
	if len(elements) != 3 || !(elements[1] == ">") {
		err := errors.New("param is error")
		return nil, err
	}
	val, err := strconv.Atoi(elements[2])
	if err != nil {
		return nil, err
	}
	return &GreaterExpression{
		key:   elements[0],
		value: val,
	}, nil
}

type AndExpression struct {
	ExpressionList []Expression
}

func (a *AndExpression) Interpret(stats map[string]int) bool {
	for _, expression := range a.ExpressionList {
		if expression.Interpret(stats) {
			return true
		}
	}
	return false
}
func NewAndExpression(strAndExpression string) *AndExpression {
	strExpressions := strings.Split(strAndExpression, "&&")
	and := AndExpression{}
	for _, str := range strExpressions {
		if strings.Contains(str, ">") {
			greater, _ := NewGreaterExpression(str)
			and.ExpressionList = append(and.ExpressionList, greater)
		}
	}
	return &and
}


type AlterRuleInterpreter struct {
	expression Expression
}

func (a *AlterRuleInterpreter) Interpret(stats map[string]int) bool {
	return a.expression.Interpret(stats)
}

func NewAlterRuleInterpreter(ruleExpression string) *AlterRuleInterpreter {
	return &AlterRuleInterpreter{
		expression: NewAndExpression(ruleExpression),
	}
}
