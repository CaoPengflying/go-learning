package interaction

type Interaction1 struct {
	Releases []Release
}

func GetInteraction1() *Interaction1 {
	return &Interaction1{}
}

func (i *Interaction1) Release(req *ReleaseReq, resp *ReleaseReply) (err error) {
	for _, release := range i.Releases {
		return release.Release(req, resp)
	}
	return
}

func (i *Interaction1) Stop(req *StopReq, resp *StopReply) (err error) {
	return
}
