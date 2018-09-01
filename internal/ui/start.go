package ui

import (
	"context"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/quickcontrols2"
	"github.com/therecipe/qt/widgets"
	"log"
	"os"
)

const (
	Default string = ""
	ApproveSignData string	= "ApproveSignData"
	OnSignerStartup string = "OnSignerStartup"
)

type RpcRequest struct {
	Params 		map[string]string
	Channel 	chan map[string]string
	Method 		string
}

type ClefUI struct {
	App 				*widgets.QApplication
	Mainw 				*widgets.QMainWindow
	currentView 		string
	IncomingRequest	 	chan RpcRequest
	views 				map[string]interface{}
}

func (c *ClefUI) initApp() {
	incomingRequest := make(chan RpcRequest)

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
	widget.SetLayout(widgets.NewQVBoxLayout())

	mainw.SetCentralWidget(widget)
	mainw.Show()

	// use the material style
	// the other inbuild styles are:
	// Default, Fusion, Imagine, Universal
	quickcontrols2.QQuickStyle_SetStyle("Material")

	c.App = app
	c.Mainw = mainw
	c.IncomingRequest = incomingRequest
}

func NewClefUI(ctx context.Context, uiClose chan bool) *ClefUI {
	c := &ClefUI{}
	c.initApp()

	login := NewLoginUI()
	approvesigndata := NewApproveSignDataUI()

	c.Mainw.Layout().AddWidget(login)
	c.Mainw.Layout().AddWidget(approvesigndata.UI)

	go func() {
		for {
			req := <-c.IncomingRequest
			log.Println(req)
			switch req.Method {
			case ApproveSignData:
				c.Mainw.SetWindowTitle("Sign Data")

				approvesigndata.SetFrom(req.Params["address"])
				approvesigndata.SetMessage(req.Params["message"])
				approvesigndata.SetEndpoint(req.Params["local"])
				approvesigndata.SetRawData(req.Params["raw_data"])
				approvesigndata.SetRemote(req.Params["remote"])
				approvesigndata.SetTransport(req.Params["transport"])
				approvesigndata.SetTxHash(req.Params["hash"])
				approvesigndata.SetResponseChannel(req.Channel)

				login.Hide()
				approvesigndata.UI.Show()
			case OnSignerStartup:
				c.Mainw.SetWindowTitle("Start Clef")
				login.Show()
				approvesigndata.UI.Hide()
			default:
				c.Mainw.SetWindowTitle("Clef")
				login.Hide()
				approvesigndata.UI.Hide()
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
