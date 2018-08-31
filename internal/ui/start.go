package ui

import (
	"context"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/quickcontrols2"
	"github.com/therecipe/qt/widgets"
	"log"
	"os"
)

type RpcRequest struct {
	Params 		map[string]string
	Channel 	chan map[string]string
	Method 		string
}

type ClefUI struct {
	App 	*widgets.QApplication
	View 	*widgets.QMainWindow
	Channel chan RpcRequest
}

func NewGethDirUI() *widgets.QWidget {
	// Create View Widget
	view := widgets.NewQWidget(nil, 0)
	view.SetLayout(widgets.NewQVBoxLayout())

	// Create Input
	input := widgets.NewQLineEdit(nil)
	input.SetPlaceholderText("Write something ...")

	// Create Submit Button
	button := widgets.NewQPushButton2("Submit", nil)
	button.ConnectClicked(func(bool) {
		log.Println("I am clicked")
	})

	// Add widgets to view
	view.Layout().AddWidget(input)
	view.Layout().AddWidget(button)
	view.Hide()

	return view
}

func NewClefUI(ctx context.Context, stopChan chan bool) *ClefUI {
	uiChannel := make(chan RpcRequest)

	// enable high dpi scaling
	// useful for devices with high pixel density displays
	// such as smartphones, retina displays, ...
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)
	core.QCoreApplication_SetApplicationName("Clef")

	app := widgets.NewQApplication(len(os.Args), os.Args)
	view := widgets.NewQMainWindow(nil, 0)
	view.SetWindowTitle(core.QCoreApplication_ApplicationName())
	view.SetStyleSheetDefault("background-color: #ecf0f1;")

	widget := widgets.NewQWidget(nil, 0)
	widget.SetLayout(widgets.NewQVBoxLayout())

	gethDirView := NewGethDirUI()
	approveSignDataView := NewApproveSignDataUI()
	widget.Layout().AddWidget(gethDirView)
	widget.Layout().AddWidget(approveSignDataView.UI)

	view.SetCentralWidget(widget)


	// use the material style
	// the other inbuild styles are:
	// Default, Fusion, Imagine, Universal
	quickcontrols2.QQuickStyle_SetStyle("Material")

	c := &ClefUI{
		App:     app,
		View:    view,
		Channel: uiChannel,
	}

	go func() {
		for {
			req := <-uiChannel
			log.Println(req)
			switch req.Method {
			case "ApproveSignData":
				view.SetWindowTitle("Approve Sign Data")

				approveSignDataView.SetFrom(req.Params["address"])
				approveSignDataView.SetMessage(req.Params["message"])
				approveSignDataView.SetEndpoint(req.Params["local"])
				approveSignDataView.SetRawData(req.Params["raw_data"])
				approveSignDataView.SetRemote(req.Params["remote"])
				approveSignDataView.SetTransport(req.Params["transport"])
				approveSignDataView.SetTxHash(req.Params["hash"])
				approveSignDataView.SetResponseChannel(req.Channel)

				gethDirView.Hide()
				approveSignDataView.UI.Show()
			default:
				view.SetWindowTitle("Clef")
				approveSignDataView.UI.Hide()
			}
		}
	}()

	var didQuit bool
	go func() {
		<-ctx.Done()
		if didQuit {
			return
		}
		app.Quit()
	}()
	app.ConnectLastWindowClosed(func() {
		didQuit = true
		stopChan <- true
	})

	return c
}
