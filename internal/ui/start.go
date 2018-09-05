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

type ClefUI struct {
	App 					*widgets.QApplication
	Mainw 					*widgets.QWidget
	currentView 			string
	//IncomingRequest	 		chan RpcRequest
	ApproveListingRequest 	chan ApproveListingRequest
	ApproveSignDataRequest 	chan ApproveSignDataRequest
	views 					map[string]interface{}
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
}

func NewClefUI(ctx context.Context, uiClose chan bool) *ClefUI {
	c := &ClefUI{}
	c.initApp()

	login := NewLoginUI()
	approvesigndata := NewApproveSignDataUI()
	approvelisting := NewApproveListingUI()

	c.Mainw.Layout().AddWidget(login)
	c.Mainw.Layout().AddWidget(approvesigndata.UI)
	c.Mainw.Layout().AddWidget(approvelisting.UI)
	c.Mainw.SetFixedSize2(400, 680	)

	go func() {
		for {
			select {
			case req := <-c.ApproveListingRequest:
				c.Mainw.SetWindowTitle("Account Listing")
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
				approvelisting.UI.Show()
			case req := <-c.ApproveSignDataRequest:
				c.Mainw.SetWindowTitle("Sign Data")
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

				co.ClickResponse(req.Response)
				login.Hide()
				approvelisting.UI.Hide()
				approvesigndata.UI.Show()
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
