package pipe_filter

import (
	"errors"
	"strconv"
)

type ToIntFilter struct {
}

var ToIntFilterWrongFormatError = errors.New("to int filter input param is error")

func NewToIntFilter() *ToIntFilter {
	return &ToIntFilter{}
}

func (tif *ToIntFilter) Process(data Request) (Response, error) {
	strArr, ok := data.([]string)
	if !ok {
		return nil, ToIntFilterWrongFormatError
	}
	res := []int{}

	for _, str := range strArr {
		num, err := strconv.Atoi(str)
		if err != nil {
			return nil, err
		}
		res = append(res, num)
	}
	return res, nil
}
