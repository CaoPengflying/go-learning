package visitor

import "testing"

func TestElement_Accept(t *testing.T) {
	prod := &ProdVisitor{}

	test := &TestVisitor{}

	element := &Element{}

	element.Accept(prod)

	element.Accept(test)
}

