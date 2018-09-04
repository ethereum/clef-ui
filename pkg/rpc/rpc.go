package rpc

import (
	"fmt"
	"github.com/kyokan/clef-ui/internal/ui"
)

type ClefService struct {
	ui 		ui.ClefUI
}

type Meta struct {
	Transport 		string `json:"scheme"`
	Remote 			string `json:"remote"`
	Local 			string `json:"local"`
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
	Address			string
}

type ApproveListingReply struct {
	Approved		bool
	Password		string
}

func (c *ClefService) OnSignerStartup(params []*OnSignerStartupParam, _ *struct{}) error {
	//fmt.Println(params[0].Info.Extapi_http)
	fmt.Println("hihihihihihihi")
	r := ui.RpcRequest{
		Method: ui.OnSignerStartup,
	}

	c.ui.IncomingRequest <- r
	return nil
}

func (c *ClefService) ApproveListing(args []*ApproveListingArgs, reply *ApproveListingReply) error {
	fmt.Println("hihihihihihihi")
	return nil
}

type ApproveSignDataParam struct {
	Address 		string `json:"address"`
	Raw_data		string `json:"raw_data"`
	Message			string `json:"message"`
	Hash 			string `json:"hash"`
	Meta 			Meta `json:"meta"`
}

func (c *ClefService) ApproveSignData(params []*ApproveSignDataParam, reply *ApproveListingReply) error {
	//c.ui.Channel <- map[string]string{
	//c.ui.Channel <- map[string]string{
	p := map[string]string{
		"address": params[0].Address,
		"hash": params[0].Hash,
		"message": params[0].Message,
		"raw_data": params[0].Raw_data,
		"transport": params[0].Meta.Transport,
		"remote": params[0].Meta.Remote,
		"local": params[0].Meta.Local,
	}

	ch := make(chan map[string]string)
	r := ui.RpcRequest{
		Params: p,
		Response: ch,
		Method: ui.ApproveSignData,
	}

	c.ui.IncomingRequest <- r

	res := <-ch
	reply.Approved = res["approved"] == "true"
	reply.Password = res["password"]

	c.ui.IncomingRequest <- ui.RpcRequest{ Method: "" }

	return nil
}


//map[jsonrpc:2.0 id:2 method:ClefService.ApproveSignData params:[map[hash:0xe35ba1e4664bb69c56eb414044a09c5f673aae2d54f29aafdd5978db1a643283 meta:map[scheme:HTTP/1.1 remote:127.0.0.1:58801 local:localhost:8550] address:0x9ab95a4082e537e5e6f82fcff41996077143029c raw_data:0xaabbccdd message:Ethereum Signed Message:
//4����]]]
