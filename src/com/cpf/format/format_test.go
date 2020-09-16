package format

import (
	"fmt"
	"testing"
)

type Point struct {
	X, Y int
}

func TestFormat(t *testing.T) {
	p := Point{1, 2}
	fmt.Printf("%v\n", p)

	fmt.Printf("%+v\n", p)

	fmt.Printf("%#v\n", p)

	fmt.Printf("%T\n",p)

	fmt.Printf("%t\n",true)

	fmt.Printf("%b\n",14)

	fmt.Printf("%p\n",&p)
	fmt.Printf("%x\n",10)


}
