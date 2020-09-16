package pipe_filter

import "errors"

type SumFilter struct {
}

func NewSumFilter() *SumFilter {
	return &SumFilter{}
}

var SumFilterFormatterError = errors.New("sum filter param  is  error")

func (sf *SumFilter) Process(data Request) (Response, error) {
	intArr,ok := data.([]int)
	if !ok {
		return nil,SumFilterFormatterError
	}
	res := 0
	for _, num := range intArr {
		res += num
	}
	return res,nil
}
