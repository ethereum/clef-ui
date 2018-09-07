package rpc

import (
	"fmt"
	"github.com/kyokan/clef-ui/internal/params"
	"github.com/kyokan/clef-ui/internal/ui"
)

type ClefService struct {
	ui 		ui.ClefUI
}

type ClefResponse struct {
	Approved		bool
	Password		string
}

type AccountListResponse struct {
	Accounts 		[]params.ApproveListingAccount
}

func (c *ClefService) OnSignerStartup(params []*params.OnSignerStartupParam, _ *struct{}) error {
	//fmt.Println(params[0].Info.Extapi_http)
	fmt.Println("hihihihihihihi")
	//r := ui.RpcRequest{
	//	Method: ui.OnSignerStartup,
	//}

	//c.ui.IncomingRequest <- r
	return nil
}

func (c *ClefService) ApproveImport(p []*params.ApproveImportParams, reply *params.ApproveImportResponse) error {
	ch := make(chan bool)
	r := ui.ApproveImportRequest{
		Params: p,
		Response: ch,
		Reply: reply,
	}

	c.ui.ApproveImportRequest <- r
	<-ch

	return nil
}

func (c *ClefService) ApproveNewAccount(p []*params.ApproveNewAccountParams, reply *params.ApproveNewAccountResponse) error {
	ch := make(chan bool)
	r := ui.ApproveNewAccountRequest{
		Params: p,
		Response: ch,
		Reply: reply,
	}

	c.ui.ApproveNewAccountRequest <- r
	<-ch

	return nil
}

func (c *ClefService) ApproveTx(p []*params.ApproveTxParams, reply *params.ApproveTxResponse) error {
	ch := make(chan bool)
	r := ui.ApproveTxRequest{
		Params: p,
		Response: ch,
		Reply: reply,
	}

	c.ui.ApproveTxRequest <- r
	<-ch

	return nil
}

func (c *ClefService) ApproveListing(p []*params.ApproveListingParams, reply *AccountListResponse) error {
	ch := make(chan []params.ApproveListingAccount)
	r := ui.ApproveListingRequest{
		Params: p,
		Response: ch,
	}

	c.ui.ApproveListingRequest <- r

	res := <-ch
	reply.Accounts = res
	//reply.Password = res["password"]

	return nil
}

func (c *ClefService) ApproveSignData(params []*params.ApproveSignDataParams, reply *ClefResponse) error {
	ch := make(chan map[string]string)
	r := ui.ApproveSignDataRequest{
		Params: params,
		Response: ch,
	}

	c.ui.ApproveSignDataRequest <- r

	res := <-ch
	reply.Approved = res["approved"] == "true"
	reply.Password = res["password"]

	//c.ui.IncomingRequest <- ui.RpcRequest{ Method: "" }

	return nil
}


//map[jsonrpc:2.0 id:2 method:ClefService.ApproveSignData params:[map[hash:0xe35ba1e4664bb69c56eb414044a09c5f673aae2d54f29aafdd5978db1a643283 meta:map[scheme:HTTP/1.1 remote:127.0.0.1:58801 local:localhost:8550] address:0x9ab95a4082e537e5e6f82fcff41996077143029c raw_data:0xaabbccdd message:Ethereum Signed Message:
//4����]]]
