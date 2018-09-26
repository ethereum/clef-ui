package rpc

import (
	"github.com/kyokan/clef-ui/internal/params"
	"github.com/kyokan/clef-ui/internal/ui"
	"strings"
)

type ClefService struct {
	ui 		ui.ClefUI
}

type ClefResponse struct {
	Approved		bool
	Password		string
}

var (
	Counter = 0
)

func (c *ClefService) OnSignerStartup(params []*params.OnSignerStartupParam, _ *struct{}) error {
	Counter++
	c.ui.BackToMain <- true
	return nil
}

func (c *ClefService) ShowError(p []*params.ShowErrorParam, _ *struct{}) error {
	text := strings.Replace(p[0].Text, "\u003c", "less than ", -1)
	c.ui.ErrorDialog <- text
	return nil
}

func (c *ClefService) ApproveExport(p []*params.ApproveExportParams, reply *params.ApproveExportResponse) error {
	Counter++
	ch := make(chan bool)
	r := ui.ApproveExportRequest{
		Params: p,
		Response: ch,
		Reply: reply,
	}

	item := &ui.TxListItem{
		From: p[0].Address,
		Method: "ApproveExport",
		RPC: r,
	}

	c.ui.IncomingRequest <- item
	<-ch

	item.Remove()

	return nil
}

func (c *ClefService) ApproveImport(p []*params.ApproveImportParams, reply *params.ApproveImportResponse) error {
	Counter++
	ch := make(chan bool)
	r := ui.ApproveImportRequest{
		Params: p,
		Response: ch,
		Reply: reply,
	}
	item := &ui.TxListItem{
		From: " - ",
		Method: "ApproveImport",
		RPC: r,
	}
	c.ui.IncomingRequest <- item
	<-ch
	item.Remove()
	return nil
}

func (c *ClefService) ApproveNewAccount(p []*params.ApproveNewAccountParams, reply *params.ApproveNewAccountResponse) error {
	Counter++
	ch := make(chan bool)
	r := ui.ApproveNewAccountRequest{
		Params: p,
		Response: ch,
		Reply: reply,
	}
	item := &ui.TxListItem{
		From: " - ",
		Method: "ApproveNewAccount",
		RPC: r,
	}
	c.ui.IncomingRequest <- item
	<-ch
	item.Remove()
	return nil
}

func (c *ClefService) ApproveTx(p []*params.ApproveTxParams, reply *params.ApproveTxResponse) error {
	Counter++
	ch := make(chan bool)
	r := ui.ApproveTxRequest{
		Params: p,
		Response: ch,
		Reply: reply,
	}
	item := &ui.TxListItem{
		From: p[0].Transaction.From,
		Method: "ApproveTx",
		RPC: r,
	}

	c.ui.IncomingRequest <- item
	<-ch
	item.Remove()

	return nil
}

func (c *ClefService) ApproveListing(p []*params.ApproveListingParams, reply *params.ApproveListingResponse) error {
	Counter++

	if Counter == 2 {
		reply.Accounts = p[0].Accounts
		return nil
	}

	ch := make(chan bool)
	r := ui.ApproveListingRequest{
		Params: p,
		Response: ch,
		Reply: reply,
	}
	item := &ui.TxListItem{
		From: " - ",
		Method: "ApproveListing",
		RPC: r,
	}
	c.ui.IncomingRequest <- item
	//c.ui.ApproveListingRequest <- r
	<-ch
	item.Remove()
	return nil
}

func (c *ClefService) ApproveSignData(params []*params.ApproveSignDataParams, reply *params.ApproveSignDataResponse) error {
	Counter++
	ch := make(chan bool)
	r := ui.ApproveSignDataRequest{
		Params: params,
		Response: ch,
		Reply: reply,
	}

	item := &ui.TxListItem {
		From: params[0].Address,
		Method: "ApproveSignData",
		RPC: r,
	}
	c.ui.IncomingRequest <- item
	<-ch
	item.Remove()
	return nil
}


//map[jsonrpc:2.0 id:2 method:ClefService.ApproveSignData params:[map[hash:0xe35ba1e4664bb69c56eb414044a09c5f673aae2d54f29aafdd5978db1a643283 meta:map[scheme:HTTP/1.1 remote:127.0.0.1:58801 local:localhost:8550] address:0x9ab95a4082e537e5e6f82fcff41996077143029c raw_data:0xaabbccdd message:Ethereum Signed Message:
//4����]]]
