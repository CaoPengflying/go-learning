package pipe_filter

type StraightPipeline struct {
	Name    string
	Filters *[]Filter
}

func NewStraightPipeline(name string, filters ...Filter) *StraightPipeline {
	return &StraightPipeline{
		Name:    name,
		Filters: &filters,
	}

}

func (f *StraightPipeline) Process(data Request) (Response, error) {
	var res interface{}
	var err error
	for _, filter := range *f.Filters {
		res, err = filter.Process(data)
		if err != nil {
			return nil, err
		}
		data = res
	}
	return res, err
}
