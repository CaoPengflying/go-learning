package interaction

type ReleaseReq struct {
}

type ReleaseReply struct {
}

type Release interface {
	Release(req *ReleaseReq, resp *ReleaseReply) (err error)
}

func GetRelease(businessType int) Release {
	switch businessType {
	case 1:
		return GetInteraction1()
	case 2:
		return GetInteraction2()
	}
	return nil
}
