package observer

import "testing"

func TestSubject_Notify(t *testing.T) {
	sub := &Subject{}
	ob := &Observer1{}
	sub.Register(ob)
	sub.Register(&Observer2{})
	sub.Notify("hello ")
	sub.Remove(ob)
	sub.Notify("remove observer1")
}

