package interaction

type Interaction2 struct {
}

func GetInteraction2() *Interaction2 {
	return &Interaction2{}
}

func (i *Interaction2) Release(req *ReleaseReq, resp *ReleaseReply) (err error) {
	return
}
