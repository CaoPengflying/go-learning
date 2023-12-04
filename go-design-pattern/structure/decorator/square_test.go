package decorator

import "testing"

func TestDraw(t *testing.T) {
	sq := Square{}
	cs := NewColorSquare(sq, "yellow")
	cs.draw()
}
