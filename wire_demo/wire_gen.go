// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package wire_demo

// Injectors from ship.go:

func InitShip() *Ship {
	pulp := NewPulp()
	ship := NewShip(pulp)
	return ship
}

// ship.go:

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
	return &Pulp{}
}