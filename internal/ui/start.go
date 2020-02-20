package ui

import (
	"context"
	"fmt"

	core2 "github.com/ethereum/go-ethereum/signer/core"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

// gfz: probably useless...
const (
	USR_BACK    = iota // User selecting menu 'back'
	USR_APPROVE = iota // User select menu 'approve'
	USR_REJECT  = iota // User select menu 'Reject'
)

// gfz: this is used in rpc.go
type ApproveSignDataRequest struct {
	Params   *core2.SignDataRequest
	Response chan *core2.SignDataResponse
}

type ApproveListingRequest struct {
	Params     core2.ListRequest
	ResponseCh chan *core2.ListResponse
}

type ApproveTxRequest struct {
	Params     *core2.SignTxRequest
	ResponseCh chan *core2.SignTxResponse
}

type ApproveNewAccountRequest struct {
	Params     *core2.NewAccountRequest
	ResponseCh chan *core2.NewAccountResponse
}

type ClefUI struct {
	currentView     string
	BackToMain      chan bool
	IncomingRequest chan IncomingRequestItem
	operationCh     chan requestInvocation // When user clicks an op in the list, it gets sent over this chan
	ErrorDialog     chan string
}

// gfz: I guess this is there to handle requests from the server, should probably be handled somewhere else.
// RequestUserInput synchronously asks for user input
func (c *ClefUI) RequestUserInput(title, message string, isPassword bool) (string, error) {
	ok := false
	echoMode := widgets.QLineEdit__Password
	if !isPassword {
		echoMode = widgets.QLineEdit__Normal
	}
	response := widgets.QInputDialog_GetText(nil, title, message, echoMode, "",
		&ok, core.Qt__Dialog, core.Qt__ImhNone)
	if ok {
		return response, nil
	}
	return "", fmt.Errorf("no input provided")
}

type contextInfo interface {
	SetTransport(string) // gfz: this is in fact calling the property setter "transport"
	SetRemote(string) // gfz: idem
	SetEndpoint(string) // gfz: idem
}

func setMeta(co contextInfo, metadata core2.Metadata) {
	co.SetTransport(metadata.Scheme)
	co.SetRemote(metadata.Remote)
	co.SetEndpoint(metadata.Local)
}

func (msg *ApproveListingRequest) handle(ui *ClefUI) {
	// co := ui.approvelisting.ContextObject
	// setMeta(co, msg.Params.Meta)

	// co.ExternalSetAccounts(msg.Params.Accounts)
	// co.ClickResponse(msg.ResponseCh)
	// ui.approvelisting.UI.Show()
}
func (msg *ApproveTxRequest) handle(ui *ClefUI) {
	// co := ui.approvetx.ContextObject
	// setMeta(co, msg.Params.Meta)

	// co.SetTransaction(msg.Params.Transaction)
	// co.ClickResponse(msg.ResponseCh)
	// ui.approvetx.UI.Show()
}

func (msg *ApproveNewAccountRequest) handle(ui *ClefUI) {
	
	// ui.approvenewaccount.UI.Show()
}

func NewClefUI(ctx context.Context, readyToStart chan string) *ClefUI {
	c := &ClefUI{}
	return c
}
