package ui

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

type CustomLabel struct {
	widgets.QLabel

	_ func(string) `signal:"updateText,auto(this.QLabel.setText)"` //TODO: support this.setText as well
}

type ApproveSignDataUI struct {
	UI 					*widgets.QWidget
	SetRemote			func(string)
	SetTransport		func(string)
	SetEndpoint			func(string)
	SetFrom				func(string)
	SetMessage			func(string)
	SetTxHash			func(string)
	SetRawData			func(string)
	responseChannel		chan map[string]string
}

func addNewRow(wrapper *widgets.QWidget, key string, value string) (updateText func(string)) {
	row := widgets.NewQWidget(nil, 0)
	box := widgets.NewQHBoxLayout()
	box.SetContentsMargins(8, 8, 0, 0)
	row.SetLayout(box)

	keyLabel := NewCustomLabel(nil, 0)
	keyLabel.SetAlignment(core.Qt__AlignLeft)
	keyLabel.UpdateText(key)
	//keyLabel.SetStyleSheetDefault("margin: 0px;")

	valueLabel := NewCustomLabel(nil, 0)
	valueLabel.SetAlignment(core.Qt__AlignLeft)
	valueLabel.UpdateText(value)
	//valueLabel.SetStyleSheetDefault("margin: 0px;")

	row.Layout().AddWidget(keyLabel)
	row.Layout().AddWidget(valueLabel)
	wrapper.Layout().AddWidget(row)

	return valueLabel.UpdateText
}

func addTitleLabel(wrapper *widgets.QWidget, title string) {
	requestInfoHeaderLabel := NewCustomLabel(nil, 0)
	requestInfoHeaderLabel.SetAlignment(core.Qt__AlignLeft)
	requestInfoHeaderLabel.UpdateText(title)
	requestInfoHeaderLabel.SetStyleSheetDefault("font-weight: bold;")
	wrapper.Layout().AddWidget(requestInfoHeaderLabel)
}

func (c *ApproveSignDataUI) addButtons (wrapper *widgets.QWidget) {
	row := widgets.NewQWidget(nil, 0)
	box := widgets.NewQHBoxLayout()
	box.SetContentsMargins(0, 0, 0, 0)
	row.SetLayout(box)
	wrapper.Layout().AddWidget(row)

	decline := widgets.NewQPushButton2("Reject", nil)
	decline.ConnectClicked(func(bool) {
		c.responseChannel <- map[string]string{
			"approved": "false",
			"password": "",
		}
	})
	row.Layout().AddWidget(decline)

	approve := widgets.NewQPushButton2("Approve", nil)
	approve.ConnectClicked(func(bool) {
		c.responseChannel <- map[string]string{
			"approved": "true",
			"password": "asdfasdf",
		}
	})

	row.Show()

	row.Layout().AddWidget(approve)
}

func (c *ApproveSignDataUI) SetResponseChannel (responseChannel chan map[string]string) {
	c.responseChannel = responseChannel
}

func NewApproveSignDataUI() *ApproveSignDataUI {
	widget := widgets.NewQWidget(nil, 0)
	widget.SetLayout(widgets.NewQVBoxLayout())
	widget.SetMinimumSize2(400, 0)
	widget.SetStyleSheetDefault("background-color: #ecf0f1;")

	addTitleLabel(widget, "Request Info")

	setRemote := addNewRow(widget,"Remote", "")
	setRemote("127.0.0.1:123")

	setTransport := addNewRow(widget,"Transport", "")
	setTransport("127.0.0.1:123")

	setEndpoint := addNewRow(widget,"Local Endpoint", "")
	setEndpoint("127.0.0.1:123")

	addTitleLabel(widget, "Transaction Info")

	setFrom := addNewRow(widget,"Address", "")
	setTxHash := addNewRow(widget,"Hash", "")
	setMessage := addNewRow(widget,"Message", "")
	setRawData := addNewRow(widget,"Raw Data", "")

	widget.Hide()

	view := &ApproveSignDataUI{
		UI: widget,
		SetRemote: setRemote,
		SetTransport: setTransport,
		SetEndpoint: setEndpoint,
		SetFrom: setFrom,
		SetMessage: setMessage,
		SetTxHash: setTxHash,
		SetRawData: setRawData,
	}

	view.addButtons(widget)

	return view
}
