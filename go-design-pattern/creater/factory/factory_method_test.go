package factory

import (
	"reflect"
	"testing"
)

func TestJsonParserFactory_CreateParser(t *testing.T) {
	jsonFactory := JsonParserFactory{}
	jsonParser := jsonFactory.CreateParser()
	jsonParser.Parse([]byte{})

	yamlFactory := YamlParserFactory{}
	yamlParser := yamlFactory.CreateParser()
	yamlParser.Parse([]byte{})

}

func Test_jsonConfigParserFactory_CreateRuleParser(t *testing.T) {
	tests := []struct {
		name string
		want IRuleConfigParser
	}{
		{
			name: "json",
			want: JsonConfigParser{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := JsonParserFactory{}
			if got := j.CreateParser(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateRuleParser() = %v, want %v", got, tt.want)
			}
		})
	}
}