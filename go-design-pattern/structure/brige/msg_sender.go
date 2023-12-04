package brige

import "log"

type IMsgSender interface {
	Send(msg string)
}

type TelephoneMsgSender struct {
}

func (t *TelephoneMsgSender) Send(msg string) {
	log.Print("telephone msg send")
	log.Println(msg)
}
// NewEmailMsgSender NewEmailMsgSender
func NewTelephoneMsgSender() *TelephoneMsgSender {
	return &TelephoneMsgSender{}
}


type INotification interface {
	Notify(msg string) error
}

type ErrorNotification struct {
	iMsgSender IMsgSender
}

func (e *ErrorNotification) Notify(msg string) error {
	e.iMsgSender.Send(msg)
	return nil
}
