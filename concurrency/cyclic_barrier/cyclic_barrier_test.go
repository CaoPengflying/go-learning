// @Description cyclic_barrier
// @Author caopengfei 2021/4/16 15:08
package cyclic_barrier

import (
	"context"
	"fmt"
	"github.com/marusama/cyclicbarrier"
	"golang.org/x/sync/semaphore"
	"sort"
	"sync"
	"testing"
	"time"
)

func TestCyclicBarrier(t *testing.T) {
	cb := cyclicbarrier.New(10)
	for i := 0; i < 11; i++ {
		go func(i int) {
			err := cb.Await(context.Background())
			if err != nil {
				t.Error(err)
			}
			fmt.Println(i)
		}(i)
	}

}

func TestCyclicBarrier1(t *testing.T) {
	count := 0
	b := cyclicbarrier.NewWithAction(10, func() error {
		count++
		return nil
	})
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 5; j++ {
				time.Sleep(100 * time.Millisecond)
				err := b.Await(context.Background())

				if err != nil {
					panic(err)
				}
			}
		}()
	}
	wg.Wait()
	t.Log(count)
}

type H2O struct {
	semaH *semaphore.Weighted
	semaO *semaphore.Weighted
	b     cyclicbarrier.CyclicBarrier
}

func New() *H2O {
	return &H2O{
		semaH: semaphore.NewWeighted(2),
		semaO: semaphore.NewWeighted(1),
		b:     cyclicbarrier.New(3),
	}
}

//氢原子处理
func (h2o *H2O) HandleH(printH func()) {
	_ = h2o.semaH.Acquire(context.Background(), 1)
	printH()
	_ = h2o.b.Await(context.Background())
	h2o.semaH.Release(1)
}

//氧原子处理
func (h2o *H2O) HandleO(printO func()) {
	h2o.semaO.Acquire(context.Background(), 1)
	printO()
	h2o.b.Await(context.Background())
	h2o.semaO.Release(1)
}

func TestWaterFactory(t *testing.T) {
	h2o := New()

	n := 100

	ch := make(chan string, 3*n)

	handleH := func() {
		ch <- "H"
	}

	handleO := func() {
		ch <- "O"
	}
	var wg sync.WaitGroup

	wg.Add(3 * n)

	for i := 0; i < 2*n; i++ {
		go func() {
			defer wg.Done()
			time.Sleep(100 * time.Millisecond)
			h2o.HandleH(handleH)
		}()
	}

	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			time.Sleep(100 * time.Millisecond)
			h2o.HandleO(handleO)
		}()
	}
	wg.Wait()

	if len(ch) != 3*n {
		t.Error("expect not fit")
	}
	var s = make([]string, 3)
	for i := 0; i < n; i++ {
		s[0] = <-ch
		s[1] = <-ch
		s[2] = <-ch
		sort.Strings(s)
		water := s[0] + s[1] + s[2]
		t.Log(water)
	}
}
