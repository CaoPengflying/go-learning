package pipe_filter

import "testing"

func TestStraightPipeline_Process(t *testing.T) {
	spliter := NewSplitFilter(",")
	converter := NewToIntFilter()
	sumer := NewSumFilter()

	sp := NewStraightPipeline("p1",spliter,converter,sumer)
	res ,err := sp.Process("1,2,3,4")
	if err != nil {
		t.Error(err)
	}
	t.Log(res)
}

