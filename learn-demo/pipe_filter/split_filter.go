package pipe_filter

import (
	"errors"
	"strings"
)

type SplitFilter struct {
	delimiter string
}

var SplitFilterWrongFormatError = errors.New("splitFilter input param is error")

func NewSplitFilter(delimiter string) *SplitFilter {
	return &SplitFilter{delimiter: delimiter}
}

func (sf *SplitFilter) Process(data Request) (Response, error) {
	str, ok := data.(string)
	if !ok {
		return nil, SplitFilterWrongFormatError
	}
	parts := strings.Split(str, sf.delimiter)
	return parts, nil
}
