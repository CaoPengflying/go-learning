package factory

import "errors"

type IRuleConfigParser interface {
	Parse(data []byte)
}

type JsonConfigParser struct {
}

type YamlConfigParser struct {
}

func (j JsonConfigParser) Parse(data []byte) {
	println("json config parser")
}

func (y YamlConfigParser) Parse(data []byte) {
	println("yaml config parser")
}

func GetConfigParserByName(name string) (IRuleConfigParser, error) {
	switch name {
	case "Json":
		return JsonConfigParser{}, nil
	case "Yaml":
		return YamlConfigParser{}, nil
	default:
		return nil, errors.New("not found the name" + name)
	}
}
