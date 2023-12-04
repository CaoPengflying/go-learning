package factory

//抽象工厂模式，通过一个工厂生产多个不同类型的对象

type ISystemConfigParser interface {
	ParseSystem(data []byte)
}

type JsonSystemParser struct {
}

func (j JsonSystemParser) ParseSystem(data []byte) {
	println("json system parser")
}

type IAbstractConfigParserFactory interface {
	CreateRuleParser() IRuleConfigParser
	CreateSystemParser() ISystemConfigParser
}

type JsonConfigParserAbstractFactory struct {
}

func (x *JsonConfigParserAbstractFactory) CreateRuleParser() IRuleConfigParser {
	return JsonConfigParser{}
}

func (x *JsonConfigParserAbstractFactory) CreateSystemParser() ISystemConfigParser {
	return JsonSystemParser{}
}
