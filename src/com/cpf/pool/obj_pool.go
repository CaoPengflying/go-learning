package pool

import (
	"errors"
	"time"
)

type ReusableObj struct {
}

type ObjPool struct {
	bufChan chan *ReusableObj
}

func NewObjPool(numOfPool int) *ObjPool {
	objPool := ObjPool{}
	objPool.bufChan = make(chan *ReusableObj, numOfPool)
	for i := 0; i < numOfPool; i++ {
		objPool.bufChan <- &ReusableObj{}
	}
	return &objPool
}

func (p *ObjPool) GetObj(timeout time.Duration) (*ReusableObj, error) {
	select {
	case res := <-p.bufChan:
		return res, nil
	case <-time.After(timeout):
		return nil, errors.New("time out")
	}
}

func (p *ObjPool) ReleaseObj(obj *ReusableObj) error {
	select {
	case p.bufChan <- obj:
		return nil
	default:
		return errors.New("overflow")
	}
}
