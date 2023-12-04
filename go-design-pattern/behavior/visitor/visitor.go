package visitor

import "fmt"

type IVisitor interface {
	Visit()
}

type ProdVisitor struct {
}

func (p *ProdVisitor) Visit() {
	fmt.Println("prod")
}

type TestVisitor struct {
}

func (t *TestVisitor) Visit() {
	fmt.Print("test")
}

type Element struct {

}

func (e *Element)Accept(visitor IVisitor)  {
	visitor.Visit()
}


