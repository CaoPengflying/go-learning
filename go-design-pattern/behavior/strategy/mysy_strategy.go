package strategy

import "fmt"

type MysyStrategy struct {
}

func (m *MysyStrategy) execute() error {
	fmt.Println("参加满元送元互动")
	return nil
}
