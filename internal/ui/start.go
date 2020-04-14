package ui

import (
	"context"
	"fmt"
	"os"

	core2 "github.com/ethereum/go-ethereum/signer/core"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/qml"
	"github.com/therecipe/qt/quick"
	"github.com/therecipe/qt/quickcontrols2"
	"github.com/therecipe/qt/widgets"
)

const (
	USR_BACK    = iota // User selecting menu 'back'
	USR_APPROVE = iota // User select menu 'approve'
	USR_REJECT  = iota // User select menu 'Reject'
)

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
	App             *widgets.QApplication
	Mainw           *widgets.QWidget
	currentView     string
	BackToMain      chan bool
	IncomingRequest chan IncomingRequestItem
	operationCh     chan requestInvocation // When user clicks an op in the list, it gets sent over this chan
	ErrorDialog     chan string

	approvesigndata   *quick.QQuickWidget
	approvetx         *ApproveTxUI
	approvelisting    *ApproveListingUI
	approvenewaccount *ApproveNewAccount
	txlist            *quick.QQuickWidget
	// login             *quick.QQuickWidget
}

func (c *ClefUI) hideAll() {
	c.approvesigndata.Hide()
	c.approvelisting.UI.Hide()
	c.approvetx.UI.Hide()
	c.approvenewaccount.UI.Hide()
	c.txlist.Hide()
	// c.login.Hide()
}

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
	SetTransport(string)
	SetRemote(string)
	SetEndpoint(string)
}

func setMeta(co contextInfo, metadata core2.Metadata) {
	co.SetTransport(metadata.Scheme)
	co.SetRemote(metadata.Remote)
	co.SetEndpoint(metadata.Local)
}

func (msg *ApproveListingRequest) handle(ui *ClefUI) {
	co := ui.approvelisting.ContextObject
	setMeta(co, msg.Params.Meta)

	co.ExternalSetAccounts(msg.Params.Accounts)
	co.ClickResponse(msg.ResponseCh)
	ui.approvelisting.UI.Show()
}
func (msg *ApproveTxRequest) handle(ui *ClefUI) {
	co := ui.approvetx.ContextObject
	setMeta(co, msg.Params.Meta)

	co.SetTransaction(msg.Params.Transaction)
	co.ClickResponse(msg.ResponseCh)
	ui.approvetx.UI.Show()
}

func (msg *ApproveNewAccountRequest) handle(ui *ClefUI) {
	co := ui.approvenewaccount.ContextObject
	setMeta(co, msg.Params.Meta)

	co.ClickResponse(msg.ResponseCh)
	ui.approvenewaccount.UI.Show()

}

func (c *ClefUI) Start() {
	gui.QGuiApplication_Exec()
}

func NewClefUI(ctx context.Context, uiClose chan bool, readyToStart chan string) *ClefUI {
	c := &ClefUI{}

	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)
	app := gui.NewQGuiApplication(len(os.Args), os.Args)
	quickcontrols2.QQuickStyle_SetStyle("Material")
	engine := qml.NewQQmlApplicationEngine(nil)

	NewLogin(engine.RootContext(), readyToStart)

	// engine.Load(core.NewQUrl3("qrc:/qml/main.qml", 0))
	engine.Load(core.NewQUrl3("internal/ui/qml/main.qml", 0))

	app.ConnectLastWindowClosed(func() {
		app.Quit()
	})

	return c
}
