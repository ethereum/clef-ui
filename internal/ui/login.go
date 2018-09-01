package ui

import (
	"github.com/therecipe/qt/widgets"
	"log"
)

func NewLoginUI() *widgets.QWidget {
	// Create View Widget
	widget := widgets.NewQWidget(nil, 0)
	widget.SetLayout(widgets.NewQVBoxLayout())

	// Create Input
	input := widgets.NewQLineEdit(nil)
	input.SetPlaceholderText("Write something ...")

	// Create Submit Button
	button := widgets.NewQPushButton2("Submit", nil)
	button.ConnectClicked(func(bool) {
		log.Println("I am clicked")
	})

	// Add widgets to widget
	widget.Layout().AddWidget(input)
	widget.Layout().AddWidget(button)
	widget.SetMinimumSize2(400, 0)
	widget.Hide()

	return widget
}
