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

func (c *ClefService) OnSignerStartup(params []*params.OnSignerStartupParam, _ *struct{}) error {
	//fmt.Println(params[0].Info.Extapi_http)
	fmt.Println("hihihihihihihi")
	//r := ui.RpcRequest{
	//	Method: ui.OnSignerStartup,
	//}

	//c.ui.IncomingRequest <- r
	return nil
}

func (c *ClefService) ApproveExport(p []*params.ApproveExportParams, reply *params.ApproveExportResponse) error {
	ch := make(chan bool)
	r := ui.ApproveExportRequest{
		Params: p,
		Response: ch,
		Reply: reply,
	}

	c.ui.IncomingRequest <- ui.TxListItem{
		From: p[0].Address,
		Method: "ApproveExport",
		RPC: r,
	}
	//c.ui.ApproveExportRequest <- r
	<-ch

	return nil
}

func (c *ClefService) ApproveImport(p []*params.ApproveImportParams, reply *params.ApproveImportResponse) error {
	ch := make(chan bool)
	r := ui.ApproveImportRequest{
		Params: p,
		Response: ch,
		Reply: reply,
	}

	c.ui.IncomingRequest <- ui.TxListItem{
		From: "",
		Method: "ApproveImport",
		RPC: r,
	}
	//c.ui.ApproveImportRequest <- r
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

	c.ui.IncomingRequest <- ui.TxListItem{
		From: "",
		Method: "ApproveNewAccount",
		RPC: r,
	}
	//c.ui.ApproveNewAccountRequest <- r
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

	c.ui.IncomingRequest <- ui.TxListItem{
		From: p[0].Transaction.From,
		Method: "ApproveTx",
		RPC: r,
	}
	//c.ui.ApproveTxRequest <- r
	<-ch

	return nil
}

func (c *ClefService) ApproveListing(p []*params.ApproveListingParams, reply *params.ApproveListingResponse) error {
	ch := make(chan bool)
	r := ui.ApproveListingRequest{
		Params: p,
		Response: ch,
		Reply: reply,
	}

	c.ui.IncomingRequest <- ui.TxListItem{
		From: "",
		Method: "ApproveListing",
		RPC: r,
	}
	//c.ui.ApproveListingRequest <- r
	<-ch

	return nil
}

func (c *ClefService) ApproveSignData(params []*params.ApproveSignDataParams, reply *params.ApproveSignDataResponse) error {
	ch := make(chan bool)
	r := ui.ApproveSignDataRequest{
		Params: params,
		Response: ch,
		Reply: reply,
	}

	c.ui.IncomingRequest <- ui.TxListItem{
		From: params[0].Address,
		Method: "ApproveSignData",
		RPC: r,
	}
	//c.ui.ApproveSignDataRequest <- r
	<-ch

	return nil
}


//map[jsonrpc:2.0 id:2 method:ClefService.ApproveSignData params:[map[hash:0xe35ba1e4664bb69c56eb414044a09c5f673aae2d54f29aafdd5978db1a643283 meta:map[scheme:HTTP/1.1 remote:127.0.0.1:58801 local:localhost:8550] address:0x9ab95a4082e537e5e6f82fcff41996077143029c raw_data:0xaabbccdd message:Ethereum Signed Message:
//4����]]]
