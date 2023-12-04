package main

import "go-learning/goWay/interaction"

type Interaction interface {
	Release(req *interaction.ReleaseReq,
		resp *interaction.ReleaseReply) (err error)
	Stop(req *interaction.StopReq,
		resp *interaction.StopReply) (err error)
}

type Interaction1 struct {
}

func (i *Interaction1) Release(req *interaction.ReleaseReq, resp *interaction.ReleaseReply) (err error) {
	return
}

func (i *Interaction1) Stop(req *interaction.StopReq, resp *interaction.StopReply) (err error) {
	return
}

type Interaction2 struct {
}

func (i *Interaction2) Release(req *interaction.ReleaseReq, resp *interaction.ReleaseReply) (err error) {
	return
}

func (i *Interaction2) Stop(req *interaction.StopReq, resp *interaction.StopReply) (err error) {
	return
}

func factory(interactionType int) Interaction {
	switch interactionType {
	case 1:
		return &Interaction1{}
	case 2:
		return &Interaction2{}
	}
	return nil
}
