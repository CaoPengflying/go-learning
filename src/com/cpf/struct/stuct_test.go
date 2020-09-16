package _struct

import (
	"fmt"
	"testing"
	"unsafe"
)

type Employee struct {
	Id   int
	Name string
	Age  int
}

func TestCreateEmployee(t *testing.T) {
	e := Employee{1, "bob", 20}
	t.Log(e)
	t.Log(e.Id)
	e1 := Employee{Name: "mike", Age: 18}
	t.Log(e1)
	e2 := new(Employee)
	e2.Age = 16
	e2.Name = "j"
	e2.Id = 3
	t.Log(e2)

	t.Logf("e is %T", &e)
	t.Logf("e2 is %T", e2)

}
// go struct定义的行为 使用指针则不需要复制参数
func (e Employee) String() string {
	fmt.Printf("address is %x \n", unsafe.Pointer(&e))
	return fmt.Sprintf("Id:%d-Name:%s-Age:%d", e.Id, e.Name, e.Age)
}

func (e *Employee) PointString() string {
	fmt.Printf("address is %x\n", unsafe.Pointer(e))
	return fmt.Sprintf("Id:%d-Name:%s-Age:%d", e.Id, e.Name, e.Age)
}

func TestStructOperations(t *testing.T) {
	e := Employee{1, "cpf", 20}
	fmt.Printf("address is %x\n", unsafe.Pointer(&e))
	t.Log(e.String())
	t.Logf(e.PointString())
}
