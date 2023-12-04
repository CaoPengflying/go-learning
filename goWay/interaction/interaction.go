package interaction

type Interaction interface {
	Release(req *ReleaseReq, resp *ReleaseReply) (err error)
	Stop(req *StopReq, resp *StopReply) (err error)
}
