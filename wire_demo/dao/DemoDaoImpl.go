// @Description dao
// @Author caopengfei 2021/6/8 20:55

package dao

type DemoDaoImpl struct {
}

func NewDemoDaoImpl() *DemoDaoImpl {
	return &DemoDaoImpl{}
}

func (dao *DemoDaoImpl) Demo() string {
	return "demoImpl"
}

