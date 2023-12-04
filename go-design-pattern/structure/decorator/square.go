package decorator

import "log"
//装饰器模式
type Square struct {
}

func (s *Square) draw() {
	log.Print("draw a square")
}

type ColorSquare struct {
	square Square
	color  string
}

func (c ColorSquare) draw() {
	c.square.draw()
	log.Print(c.color)
}

func NewColorSquare(square Square, color string) *ColorSquare {
	return &ColorSquare{square: square, color: color}
}
