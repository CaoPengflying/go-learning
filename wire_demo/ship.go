//+build wireinject
// @Description wire_demo
// @Author caopengfei 2021/6/3 16:58

package wire_demo

import (
	"github.com/google/wire"
)

type Ship struct {
	Pulp *Pulp
}
func NewShip(pulp *Pulp) *Ship {
	return &Ship{
		Pulp: pulp,
	}
}
type Pulp struct {
	Count int
}

func NewPulp() *Pulp {
	return &Pulp{
	}
}


func InitShip() *Ship {
	wire.Build(
		NewPulp,
		NewShip,
	)
	return &Ship{}
}

