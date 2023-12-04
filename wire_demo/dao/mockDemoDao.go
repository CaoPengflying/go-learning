// @Description dao
// @Author caopengfei 2021/6/3 17:42

package dao

type MockDemoDao struct {
}

func NewMockDemoDao() *MockDemoDao {
	return &MockDemoDao{}
}

func (mock *MockDemoDao) Demo() string {
	return "demo"
}
