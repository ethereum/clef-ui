package ui

import (
	"context"
	"github.com/kyokan/clef-ui/internal/params"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/quickcontrols2"
	"github.com/therecipe/qt/widgets"
	"os"
)

const (
	Default string = ""
	ApproveSignData string	= "ApproveSignData"
	OnSignerStartup string = "OnSignerStartup"
)

type RpcRequest struct {
	Params 		map[string]string
	Response 	chan map[string]string
	Method 		string
}

type ApproveSignDataRequest struct {
	Params 		[]*params.ApproveSignDataParams
	Response 	chan map[string]string
}

type ApproveListingRequest struct {
	Params 		[]*params.ApproveListingParams
	Response 	chan []params.ApproveListingAccount
}

type ApproveTxRequest struct {
	Params 		[]*params.ApproveTxParams
	Response 	chan bool
	Reply 		*params.ApproveTxResponse
}

type ApproveNewAccountRequest struct {
	Params 		[]*params.ApproveNewAccountParams
	Reply 		*params.ApproveNewAccountResponse
	Response 	chan bool
}

type ClefUI struct {
	App 						*widgets.QApplication
	Mainw 						*widgets.QWidget
	currentView 				string
	//IncomingRequest	 		chan RpcRequest
	ApproveListingRequest 		chan ApproveListingRequest
	ApproveSignDataRequest 		chan ApproveSignDataRequest
	ApproveTxRequest 			chan ApproveTxRequest
	ApproveNewAccountRequest 	chan ApproveNewAccountRequest
	views 						map[string]interface{}
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
	box.SetContentsMargins(0,0,0,0)
	widget.SetLayout(box)
	mainw.SetCentralWidget(widget)

	mainw.Show()

	// use the material style
	// the other inbuild styles are:
	// Default, Fusion, Imagine, Universal
	quickcontrols2.QQuickStyle_SetStyle("Default")

	c.App = app
	c.Mainw = widget
	//c.IncomingRequest = incomingRequest
	c.ApproveListingRequest = make(chan ApproveListingRequest)
	c.ApproveSignDataRequest = make(chan ApproveSignDataRequest)
	c.ApproveTxRequest = make(chan ApproveTxRequest)
	c.ApproveNewAccountRequest = make(chan ApproveNewAccountRequest)
}

func NewClefUI(ctx context.Context, uiClose chan bool) *ClefUI {
	c := &ClefUI{}
	c.initApp()

	login := NewLoginUI()
	approvesigndata := NewApproveSignDataUI()
	approvelisting := NewApproveListingUI()
	approvetx := NewApproveTxUI()
	approvenewaccount := NewApproveNewAccountUI()

	c.Mainw.Layout().AddWidget(login)
	c.Mainw.Layout().AddWidget(approvesigndata.UI)
	c.Mainw.Layout().AddWidget(approvelisting.UI)
	c.Mainw.Layout().AddWidget(approvetx.UI)
	c.Mainw.Layout().AddWidget(approvenewaccount.UI)
	c.Mainw.SetFixedSize2(400, 680	)

	go func() {
		for {
			select {
			case req := <-c.ApproveListingRequest:
				c.Mainw.SetWindowTitle("List Account")
				param := req.Params[0]

				co := approvelisting.ContextObject
				co.SetTransport(param.Meta.Transport)
				co.SetRemote(param.Meta.Remote)
				co.SetEndpoint(param.Meta.Local)

				model := approvelisting.AccountListModel
				model.Clear()

				for _, account := range param.Accounts {
					model.Add(account)
				}
				model.ClickResponse(req.Response)

				login.Hide()
				approvesigndata.UI.Hide()
				approvetx.UI.Hide()
				approvenewaccount.UI.Hide()
				approvelisting.UI.Show()
			case req := <-c.ApproveSignDataRequest:
				c.Mainw.SetWindowTitle("Sign Data")
				data := req.Params
				param := data[0]

				co := approvesigndata.ContextObject
				co.Reset()
				co.SetTransport(param.Meta.Transport)
				co.SetRemote(param.Meta.Remote)
				co.SetEndpoint(param.Meta.Local)

				co.SetHash(param.Hash)
				co.SetMessage(param.Message)
				co.SetRawData(param.Raw_data)
				co.SetFrom(param.Address)

				co.ClickResponse(req.Response)
				login.Hide()
				approvelisting.UI.Hide()
				approvetx.UI.Hide()
				approvenewaccount.UI.Hide()
				approvesigndata.UI.Show()
			case req := <-c.ApproveTxRequest:
				c.Mainw.SetWindowTitle("Send Transaction")
				data := req.Params
				param := data[0]

				co := approvetx.ContextObject
				co.Reset()
				co.SetTransport(param.Meta.Transport)
				co.SetRemote(param.Meta.Remote)
				co.SetEndpoint(param.Meta.Local)
				co.SetTransaction(param.Transaction)
				//co.SetFrom(param.Transaction.From)
				//co.SetData(param.Transaction.Data)
				//co.SetTo(param.Transaction.To)
				//co.SetGas(param.Transaction.Gas)
				//co.SetGasPrice(param.Transaction.GasPrice)
				//co.SetValue(param.Transaction.Value)
				//co.SetNonce(param.Transaction.Nonce)
				//co.formData = param.Transaction
				co.ClickResponse(req.Reply, req.Response)


				login.Hide()
				approvelisting.UI.Hide()
				approvesigndata.UI.Hide()
				approvenewaccount.UI.Hide()
				approvetx.UI.Show()
			case req := <-c.ApproveNewAccountRequest:
				c.Mainw.SetWindowTitle("New Account")
				data := req.Params
				param := data[0]

				co := approvenewaccount.ContextObject
				co.Reset()
				co.SetTransport(param.Meta.Transport)
				co.SetRemote(param.Meta.Remote)
				co.SetEndpoint(param.Meta.Local)
				co.ClickResponse(req.Reply, req.Response)

				login.Hide()
				approvelisting.UI.Hide()
				approvesigndata.UI.Hide()
				approvetx.UI.Hide()
				approvenewaccount.UI.Show()
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
