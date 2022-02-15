package error

import (
	"errors"
	"testing"
)

var smallerError = errors.New("num cannot smaller than 2")
var biggerError = errors.New("num cannot bigger than 100")

func fib(num int) ([]int, error) {
	if num < 2 {
		return nil, smallerError
	}
	if num > 100 {
		return nil, biggerError
	}
	fib := []int{1, 1}

	for i := 2; i < num; i++ {
		fib = append(fib, fib[i-1]+fib[i-2])
	}
	return fib, nil
}

func TestError(t *testing.T) {
	var (
		fibList []int
		ok      error
	)
	if fibList, ok = fib(5); ok != nil {
		t.Error(ok)
		return
	}
	t.Log(fibList)

}

func TestPanicVsExit(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Error("recovered panic",err)
		}
		t.Log("end")
	}()
	t.Log("start")
	panic(errors.New("Something wrong"))
	//os.Exit(-1)
}
