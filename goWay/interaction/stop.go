package interaction

type StopReq struct {
}

type StopReply struct {
}

type Stop interface {
	Stop(req *StopReq, resp *StopReply) (err error)
}

func GetStop(businessType int) Stop {
	switch businessType {
	case 1:
		return GetInteraction1()
	}
	return nil
}
