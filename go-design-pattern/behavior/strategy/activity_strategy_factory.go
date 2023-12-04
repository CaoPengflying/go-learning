package strategy

import "errors"

var strategyMap = map[string]IActivityStrategy{
	"gdzk": &GdzkStrategy{},
	"mysy": &MysyStrategy{},
}

type ActivityStrategyFactory struct {
}

func (factory *ActivityStrategyFactory) GetStrategy(name string) (IActivityStrategy, error) {
	if len(name) < 1 {
		return nil, errors.New("name can't be empty")
	}
	if strategyMap[name] == nil {
		return nil, errors.New("not found strategy")
	}
	return strategyMap[name], nil
}
