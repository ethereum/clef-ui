package rpc

type ClefService struct {
}

type OnSignerStartupInfo struct {
	Extapi_http    string
	Extapi_ipc     string
	Extapi_version string
	Intapi_version string
}

type OnSignerStartupParam struct {
	Info *OnSignerStartupInfo
}

type OnSignerStartupReply struct {
}

type ApproveListingArgs struct {
}

type ApproveListingReply struct {
}

func (c *ClefService) OnSignerStartup(info []*OnSignerStartupInfo, _ *struct{}) error {
	return nil
}

func (c *ClefService) ApproveListing(args *ApproveListingArgs, reply *ApproveListingReply) error {
	return nil
}
