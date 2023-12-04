package template

import (
	"errors"
	"fmt"
)

type ISMS interface {
	send(content, phone string) error
}

type sms struct {
	ISMS
}

func (s *sms) Valid(content string) error {
	if len(content) > 50 {
		return errors.New("content is too long")
	}
	if len(content) < 1 {
		return errors.New("content can't be empty")
	}
	return nil
}

func (s *sms) Send(content, phone string) error {
	if err := s.Valid(content); err != nil {
		return err
	}

	return s.send(content, phone)
}

type TelecomSms struct {
	*sms
}

func (t *TelecomSms) send(content, phone string) error {
	fmt.Println("send by telecomSms")
	return nil
}

func NewTelecomSms() *TelecomSms {
	tel := &TelecomSms{}
	tel.sms = &sms{ISMS: tel}
	return tel
}
