package ui

import (
	"context"
	"github.com/kyokan/clef-ui/internal/identicon"
	"github.com/kyokan/clef-ui/internal/params"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/quick"
	"github.com/therecipe/qt/quickcontrols2"
	"github.com/therecipe/qt/widgets"
	"log"
	"os"
)

type RpcRequest struct {
	Params 		map[string]string
	Response 	chan map[string]string
	Method 		string
}

type ApproveSignDataRequest struct {
	Params 		[]*params.ApproveSignDataParams
	Response 	chan bool
	Reply 		*params.ApproveSignDataResponse
}

type ApproveListingRequest struct {
	Params 		[]*params.ApproveListingParams
	Response 	chan bool
	Reply 		*params.ApproveListingResponse
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

type ApproveImportRequest struct {
	Params 		[]*params.ApproveImportParams
	Reply 		*params.ApproveImportResponse
	Response 	chan bool
}

type ApproveExportRequest struct {
	Params 		[]*params.ApproveExportParams
	Reply 		*params.ApproveExportResponse
	Response 	chan bool
}

type ClefUI struct {
	App 						*widgets.QApplication
	Mainw 						*widgets.QWidget
	currentView 				string
	IncomingRequest	 			chan *TxListItem
	BackToMain 					chan bool
	ApproveListingRequest 		chan ApproveListingRequest
	ApproveSignDataRequest 		chan ApproveSignDataRequest
	ApproveTxRequest 			chan ApproveTxRequest
	ApproveNewAccountRequest 	chan ApproveNewAccountRequest
	ApproveImportRequest 		chan ApproveImportRequest
	ApproveExportRequest 		chan ApproveExportRequest

	approvesigndata 			*quick.QQuickWidget
	approvetx	 	 			*quick.QQuickWidget
	approvelisting 				*quick.QQuickWidget
	approvenewaccount 			*quick.QQuickWidget
	approveimport 				*quick.QQuickWidget
	approveexport 				*quick.QQuickWidget
	txlist 						*quick.QQuickWidget
	login 						*quick.QQuickWidget
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
	c.IncomingRequest = make(chan *TxListItem)
	c.ApproveListingRequest = make(chan ApproveListingRequest)
	c.ApproveSignDataRequest = make(chan ApproveSignDataRequest)
	c.ApproveTxRequest = make(chan ApproveTxRequest)
	c.ApproveNewAccountRequest = make(chan ApproveNewAccountRequest)
	c.ApproveImportRequest = make(chan ApproveImportRequest)
	c.ApproveExportRequest = make(chan ApproveExportRequest)
	c.BackToMain = make(chan bool)
}

func (c *ClefUI) hideAll() {
	c.approvesigndata.Hide()
	c.approvelisting.Hide()
	c.approvetx.Hide()
	c.approveimport.Hide()
	c.approveexport.Hide()
	c.approvenewaccount.Hide()
	c.txlist.Hide()
	c.login.Hide()
}

func NewClefUI(ctx context.Context, uiClose chan bool, readyToStart chan string) *ClefUI {
	c := &ClefUI{}
	c.initApp()

	approvesigndata := NewApproveSignDataUI(c)
	approvelisting := NewApproveListingUI(c)
	approvetx := NewApproveTxUI(c)
	approvenewaccount := NewApproveNewAccountUI(c)
	approveimport := NewApproveImportUI(c)
	approveexport := NewApproveExportUI(c)
	txlist := NewTxListUI(c)
	login := NewLoginUI(c, readyToStart)

	c.approvesigndata = approvesigndata.UI
	c.approvelisting = approvelisting.UI
	c.approvetx = approvetx.UI
	c.approvenewaccount = approvenewaccount.UI
	c.approveimport = approveimport.UI
	c.approveexport = approveexport.UI
	c.login = login.UI
	c.txlist = txlist.UI

	c.Mainw.Layout().AddWidget(approvesigndata.UI)
	c.Mainw.Layout().AddWidget(approvelisting.UI)
	c.Mainw.Layout().AddWidget(approvetx.UI)
	c.Mainw.Layout().AddWidget(approvenewaccount.UI)
	c.Mainw.Layout().AddWidget(approveimport.UI)
	c.Mainw.Layout().AddWidget(approveexport.UI)
	c.Mainw.Layout().AddWidget(txlist.UI)
	c.Mainw.Layout().AddWidget(login.UI)
	c.Mainw.SetFixedSize2(400, 680	)

	//c.txlist.Show()
	c.hideAll()
	login.UI.Show()
	login.ContextObject.SetGopath(os.Getenv("GOPATH"))

	go func() {
		for {
			select {
			case <-c.BackToMain:
				log.Println("Back to Main")
				c.hideAll()
				txlist.CtxObject.RequestAccounts()
				txlist.UI.Show()
			case req := <-c.IncomingRequest:
				log.Println("New Request")
				txlist.CtxObject.transactions.Add(req)
			case req := <-c.ApproveListingRequest:
				log.Println("Approve Listing")
				c.hideAll()
				param := req.Params[0]

				co := approvelisting.ContextObject
				co.SetTransport(param.Meta.Transport)
				co.SetRemote(param.Meta.Remote)
				co.SetEndpoint(param.Meta.Local)

				model := co.accounts

				for _, account := range param.Accounts {
					model.Add(account)
				}

				co.ClickResponse(req.Reply, req.Response)

				approvelisting.UI.Show()
			case req := <-c.ApproveSignDataRequest:
				c.hideAll()
				data := req.Params
				param := data[0]

				co := approvesigndata.ContextObject

				co.SetTransport(param.Meta.Transport)
				co.SetRemote(param.Meta.Remote)
				co.SetEndpoint(param.Meta.Local)

				co.SetHash(param.Hash)
				co.SetMessage(param.Message)
				co.SetRawData(param.Raw_data)
				co.SetFrom(param.Address)
				co.SetFromSrc(identicon.ToBase64Img(param.Address))

				co.ClickResponse(req.Reply, req.Response)

				approvesigndata.UI.Show()
			case req := <-c.ApproveTxRequest:
				c.hideAll()
				data := req.Params
				param := data[0]

				co := approvetx.ContextObject
				co.SetTransport(param.Meta.Transport)
				co.SetRemote(param.Meta.Remote)
				co.SetEndpoint(param.Meta.Local)
				co.SetTransaction(param.Transaction)
				co.ClickResponse(req.Reply, req.Response)

				approvetx.UI.Show()
			case req := <-c.ApproveNewAccountRequest:
				c.hideAll()
				data := req.Params
				param := data[0]

				co := approvenewaccount.ContextObject
				co.SetTransport(param.Meta.Transport)
				co.SetRemote(param.Meta.Remote)
				co.SetEndpoint(param.Meta.Local)
				co.ClickResponse(req.Reply, req.Response)

				approvenewaccount.UI.Show()
			case req := <-c.ApproveImportRequest:
				c.hideAll()
				data := req.Params
				param := data[0]

				co := approveimport.ContextObject
				co.SetTransport(param.Meta.Transport)
				co.SetRemote(param.Meta.Remote)
				co.SetEndpoint(param.Meta.Local)
				co.ClickResponse(req.Reply, req.Response)

				approveimport.UI.Show()
			case req := <-c.ApproveExportRequest:
				c.hideAll()
				data := req.Params
				param := data[0]

				co := approveexport.ContextObject
				co.SetTransport(param.Meta.Transport)
				co.SetRemote(param.Meta.Remote)
				co.SetEndpoint(param.Meta.Local)
				co.SetAddress(param.Address)
				co.SetFromSrc(identicon.ToBase64Img(param.Address))
				co.ClickResponse(req.Reply, req.Response)

				approveexport.UI.Show()
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
