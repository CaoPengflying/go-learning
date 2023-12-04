package interaction

type Server struct {
}

func (s *Server) Release(req *ReleaseReq, resp *ReleaseReply) (err error) {
	release := GetRelease(1)
	release.Release(req, resp)
	return
}

func (s *Server) Stop(req *StopReq, resp *StopReply) (err error) {
	stop := GetStop(1)
	stop.Stop(req, resp)
	return
}
