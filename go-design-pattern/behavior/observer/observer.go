package observer

type ISubject interface {
	Register(observer IObserver)
	Remove(observer IObserver)
	Notify(msg string)
}

type IObserver interface {
	Update(msg string)
}

type Subject struct {
	observers []IObserver
}

func (s *Subject) Register(observer IObserver) {
	s.observers = append(s.observers, observer)
	s.Notify("one observer regist")
}

func (s *Subject) Remove(observer IObserver) {
	for i, iObserver := range s.observers {
		if iObserver == observer {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
		}
	}
}
func (s *Subject) Notify(msg string) {
	for _, observer := range s.observers {
		observer.Update(msg)
	}
}

type Observer1 struct {
}

func (o *Observer1) Update(msg string) {
	println("observer1 : " + msg)
}

type Observer2 struct {
}

func (o *Observer2) Update(msg string) {
	println("observer2 :" + msg)
}
