import QtQuick 2.7
import QtQuick.Controls 2.1
import "Style"
import "Controls" as Controls
import "Views" as ClefUI

ApplicationWindow {
    id: window

	visible: true
    title: "Clef UI"
	minimumWidth: 400
	minimumHeight: 680

    // StackView?

    ClefUI.Login {
        id: login
        anchors.fill: parent
    }

    Component.onCompleted: {
        login.forceActiveFocus()
    }
}
