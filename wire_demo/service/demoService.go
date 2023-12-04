// @Description service
// @Author caopengfei 2021/6/3 17:34

package service

import "wire_demo/dao"

type DemoService interface {
	Demo() string
}

type DemoServiceImpl struct {
	dao dao.DemoDao
}

func NewDemoServiceImpl(dao dao.DemoDao) *DemoServiceImpl {
	return &DemoServiceImpl{
		dao: dao,
	}
}

func (demo *DemoServiceImpl) Demo() string {
	return demo.dao.Demo()
}
