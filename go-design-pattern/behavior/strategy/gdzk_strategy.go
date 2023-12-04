package strategy

import "fmt"

type GdzkStrategy struct {

}

func (g *GdzkStrategy)execute() error {
	fmt.Println("固定打折活动")
	return nil
}