package command

import "fmt"

//通过演示遥控器操作电视机的业务实现来展示命令模式

type TV struct {
}

func (tv *TV) Open() {
	fmt.Println("open the tv")
}

func (tv *TV) Close() {
	fmt.Println("close the tv")
}

type Command interface {
	Press()
}

type OpenCommand struct {
	tv TV
}

func (o *OpenCommand) Press() {
	o.tv.Open()
}

type CloseCommand struct {
	tv TV
}

func (o *CloseCommand) Press() {
	o.tv.Close()
}

type Invoker struct {
	cmd Command
}

func (i *Invoker) SetCommand(cmd Command) {
	i.cmd = cmd
}
func (i *Invoker) Do() {
	i.cmd.Press()
}
