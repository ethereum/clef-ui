package ui

import (
	"context"
	"fmt"
	"os"

	core2 "github.com/ethereum/go-ethereum/signer/core"
	"github.com/therecipe/qt/core"
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
	IncomingRequest chan *IncomingRequestItem
	operationCh     chan requestInvocation // When user clicks an op in the list, it gets sent over this chan
	ErrorDialog     chan string

	approvesigndata   *quick.QQuickWidget
	approvetx         *ApproveTxUI
	approvelisting    *ApproveListingUI
	approvenewaccount *ApproveNewAccount
	txlist            *quick.QQuickWidget
	login             *quick.QQuickWidget
}

func (c *ClefUI) initApp() {
	//incomingRequest := make(chan RpcRequest)

	// enable high dpi scaling
	// useful for devices with high pixel density displays
	// such as smartphones, retina displays, ...
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)
	core.QCoreApplication_SetApplicationName("Clef")

	app := widgets.NewQApplication(len(os.Args), os.Args)
	mainw := widgets.NewQMainWindow(nil, 0)
	mainw.SetWindowTitle(core.QCoreApplication_ApplicationName())
	mainw.SetStyleSheetDefault("background-color: #ecf0f1;")

	widget := widgets.NewQWidget(nil, 0)
	box := widgets.NewQVBoxLayout()
	box.SetSpacing(0)
	box.SetContentsMargins(0, 0, 0, 0)
	widget.SetLayout(box)
	mainw.SetCentralWidget(widget)

	mainw.Show()

	// use the material style
	// the other inbuild styles are:
	// Default, Fusion, Imagine, Universal
	quickcontrols2.QQuickStyle_SetStyle("Default")

	c.App = app
	c.Mainw = widget
	c.IncomingRequest = make(chan *IncomingRequestItem)
	c.operationCh = make(chan requestInvocation)
	c.BackToMain = make(chan bool)
	c.ErrorDialog = make(chan string)
}

func (c *ClefUI) hideAll() {
	c.approvesigndata.Hide()
	c.approvelisting.UI.Hide()
	c.approvetx.UI.Hide()
	c.approvenewaccount.UI.Hide()
	c.txlist.Hide()
	c.login.Hide()
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

func NewClefUI(ctx context.Context, uiClose chan bool, readyToStart chan string) *ClefUI {
	c := &ClefUI{}
	c.initApp()

	approvesigndata := NewApproveSignDataUI(c)
	approvelisting := NewApproveListingUI(c)
	approvetx := NewApproveTxUI(c)
	approvenewaccount := NewApproveNewAccountUI(c)
	txlist := NewTxListUI(c)
	login := NewLoginUI(c, readyToStart)
	errordialog := widgets.NewQErrorMessage(nil)
	errordialog.SetFixedSize2(350, 200)

	c.approvelisting = approvelisting
	c.approvetx = approvetx
	c.approvesigndata = approvesigndata.UI
	c.approvenewaccount = approvenewaccount
	c.login = login.UI
	c.txlist = txlist.UI

	c.Mainw.Layout().AddWidget(approvesigndata.UI)
	c.Mainw.Layout().AddWidget(approvelisting.UI)
	c.Mainw.Layout().AddWidget(approvetx.UI)
	c.Mainw.Layout().AddWidget(approvenewaccount.UI)
	c.Mainw.Layout().AddWidget(txlist.UI)
	c.Mainw.Layout().AddWidget(login.UI)
	c.Mainw.SetFixedSize2(400, 680)

	c.hideAll()
	login.UI.Show()

	go func() {
		for {
			select {
			case text := <-c.ErrorDialog:
				errordialog.ShowMessage(text)
			case <-c.BackToMain:
				// User clicked 'back' to the main listing
				c.hideAll()
				txlist.UI.Show()
			case req := <-c.IncomingRequest:
				// New request came in from the 'network'
				txlist.CtxObject.transactions.Add(req)
			case op := <-c.operationCh:
				// User selected an item for processing
				c.hideAll()
				op.handle(c)
			}
		}
	}()

	var didQuit bool
	go func() {
		<-ctx.Done()
		if didQuit {
			return
		}
		c.App.Quit()
	}()
	c.App.ConnectLastWindowClosed(func() {
		didQuit = true
		uiClose <- true
	})

	return c
}
