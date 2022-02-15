// @Description err_group
// @Author caopengfei 2021/4/16 18:53
package err_group

import (
	"context"
	"errors"
	"fmt"
	"github.com/mdlayher/schedgroup"
	"github.com/neilotoole/errgroup"
	"github.com/vardius/gollback"
	"testing"
	"time"
)

func TestErrGroup(t *testing.T) {
	var g errgroup.Group

	g.Go(func() error {
		time.Sleep(5 * time.Second)
		fmt.Println("exec #1")
		return nil
	})

	g.Go(func() error {
		time.Sleep(10 * time.Second)
		fmt.Println("exec #2")
		return errors.New("2 error")
	})

	g.Go(func() error {
		time.Sleep(15 * time.Second)
		fmt.Println("exec #3")
		return errors.New("3 error")
	})
	if err := g.Wait(); err == nil {
		fmt.Println("success", err)
	} else {
		fmt.Println("failed", err)
	}
}

func TestErrGroupList(t *testing.T) {
	var g errgroup.Group
	result := make([]error, 3)

	g.Go(func() error {
		time.Sleep(5 * time.Second)
		fmt.Println("exec #1")
		result[0] = nil
		return nil
	})

	g.Go(func() error {
		time.Sleep(10 * time.Second)
		fmt.Println("exec #2")
		result[1] = errors.New("2 error")
		return result[1]
	})

	g.Go(func() error {
		time.Sleep(15 * time.Second)
		fmt.Println("exec #3")
		result[2] = nil
		return nil
	})

	if err := g.Wait(); err == nil {
		fmt.Println("success", result)
	} else {
		fmt.Println("failed", result)
	}
}

func TestNewErrGroup(t *testing.T) {
	g, _ := errgroup.WithContextN(context.Background(), 2, 10)
	g.Go(func() error {
		time.Sleep(5 * time.Second)
		fmt.Println("exec #1")
		return nil
	})

	g.Go(func() error {
		time.Sleep(10 * time.Second)
		fmt.Println("exec #2")
		return errors.New("2 error")
	})

	g.Go(func() error {
		time.Sleep(15 * time.Second)
		fmt.Println("exec #3")
		return nil
	})
	if err := g.Wait(); err == nil {
		fmt.Println("success")
	} else {
		fmt.Println("failed")
	}

}

func TestGollback(t *testing.T) {
	rs, errs := gollback.All(context.Background(), func(ctx context.Context) (interface{}, error) {
		time.Sleep(3 * time.Second)
		return 1, nil
	}, func(ctx context.Context) (interface{}, error) {
		time.Sleep(5 * time.Second)
		return 2, errors.New("error")
	}, func(ctx context.Context) (interface{}, error) {
		time.Sleep(8 * time.Second)
		return 3, nil
	})

	t.Log(rs)
	t.Log(errs)
}

func TestGollbackRace(t *testing.T) {
	rs, errs := gollback.Race(context.Background(), func(ctx context.Context) (interface{}, error) {
		time.Sleep(3 * time.Second)
		return 1, nil
	}, func(ctx context.Context) (interface{}, error) {
		time.Sleep(5 * time.Second)
		return 2, errors.New("error")
	})

	t.Log(rs)
	t.Log(errs)
}

func TestGollbackRetry(t *testing.T) {
	res, err := gollback.Retry(context.Background(), 2, func(ctx context.Context) (interface{}, error) {
		return nil, errors.New("err")
	})
	t.Log(res)
	t.Log(err)
}

func TestSchedGroup(t *testing.T) {
	g := schedgroup.New(context.Background())
	g.Delay(time.Second, func() {
		fmt.Println("1s延迟执行")
	})
	g.Schedule(time.Now(), func() {
		fmt.Println("指定时间执行")
	})
	err := g.Wait()
	t.Log(err)
}
