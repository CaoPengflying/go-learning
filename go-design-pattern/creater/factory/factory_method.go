package factory

type IRuleConfigParserFactory interface {
	CreateParser() IRuleConfigParser
}

type JsonParserFactory struct{}

func (JsonParserFactory) CreateParser() IRuleConfigParser {
	return JsonConfigParser{}
}

type YamlParserFactory struct{}

func (YamlParserFactory) CreateParser() IRuleConfigParser {
	return YamlConfigParser{}
}
