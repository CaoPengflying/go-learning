package extension

import "testing"

type Pet struct {
}

func (p *Pet) speak() {
	println("...")
}

func (p *Pet) speakTo(host string) {
	p.speak()
	println("  ", host)
}
// 使用组合的方式调用方法
type Dog struct {
	p *Pet
}

func (d *Dog) speak() {
	d.p.speak()
}

func (d *Dog) speakTo(host string) {
	d.p.speakTo(host)
}
// 匿名的方式去实现继承，并且子类不能替换父类
type Cat struct {
	Pet
}
func (c *Cat) speak() {
	println("喵...")
}





func TestExtension(t *testing.T) {
	dog := Dog{}
	dog.speak()
	dog.speakTo("dog")

	cat := new(Cat)
	cat.speakTo("cat")
	cat.speak()

	//var p Pet = new(Dog)
}
