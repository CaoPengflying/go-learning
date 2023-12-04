package command

import "testing"

func TestCommand(t *testing.T) {
	var tv TV
	openCommand := OpenCommand{tv: tv}
	invoker := Invoker{&openCommand}
	invoker.Do()

	invoker.SetCommand(&CloseCommand{
		tv: tv,
	})
	invoker.Do()
}

