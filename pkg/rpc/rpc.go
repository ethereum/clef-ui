package rpc

import "fmt"

type ClefService struct {
}

type OnSignerStartupInfo struct {
	Extapi_http    string `json:"extapi_http"`
	Extapi_ipc     string `json:"extapi_ipc"`
	Extapi_version string `json:"extapi_version"`
	Intapi_version string `json:"intapi_version"`
}

type OnSignerStartupParam struct {
	Info *OnSignerStartupInfo `json:"info"`
}

type OnSignerStartupReply struct {
}

type ApproveListingArgs struct {
}

type ApproveListingReply struct {
}

func (c *ClefService) OnSignerStartup(info []*OnSignerStartupParam, _ *struct{}) error {
	fmt.Println(info[0].Info.Extapi_http)
	return nil
}

func (c *ClefService) ApproveListing(args *ApproveListingArgs, reply *ApproveListingReply) error {
	return nil
}
