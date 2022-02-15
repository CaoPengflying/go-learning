package performance

import "testing"

func TestCreateRequest(t *testing.T) {
	str := createRequest()
	t.Log(str)
}

func TestProcessRequest(t *testing.T) {
	reqs := []string{}
	reqs = append(reqs, createRequest())
	reps := processRequest(reqs)
	t.Log(reps[0])
}

func BenchmarkProcessRequest(b *testing.B) {
	b.ResetTimer()
	var reqs []string
	reqs = append(reqs, createRequest())
	for i := 0; i < b.N; i++ {
		_ = processRequest(reqs)
	}
	b.StopTimer()
}

//func BenchmarkProcessRequestEasyJson(b *testing.B) {
//	b.ResetTimer()
//	var reqs []string
//	reqs = append(reqs, createRequest())
//	for i := 0; i < b.N; i++ {
//		_ = processRequest(reqs)
//	}
//	b.StopTimer()
//}
