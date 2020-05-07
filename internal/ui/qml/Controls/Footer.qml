import QtQuick 2.7
import QtQuick.Controls 2.1
import "." as ClefUI
import "../Style"

Item {
    id: root

    implicitHeight: 55
    implicitWidth: row.implicitWidth

    property bool rejectButtonVisible: false
    property bool approveButtonVisible: false
    property bool startButtonVisible: false

    property alias startButtonEnabled: start.enabled

    signal rejectClicked()
    signal approveClicked()
    signal startClicked()

    Rectangle {
        anchors.left: parent.left
        anchors.right: parent.right
        anchors.top: parent.top
        height: Style.separators.thickness
        color: Style.separators.color
    }

    Row {
        id: row
        anchors.centerIn: parent
        spacing: 50

        ClefUI.Button {
            text: qsTr("reject") // + Translation.trigger
            visible: rejectButtonVisible
            onClicked: root.rejectClicked()
        }

        ClefUI.Button {
            text: qsTr("approve") // + Translation.trigger
            color: Style.buttons.validationColor
            visible: approveButtonVisible
            onClicked: root.approveClicked()
        }

        ClefUI.Button {
            id: start
            text: qsTr("start") // + Translation.trigger
            color: Style.buttons.validationColor
            visible: startButtonVisible
            onClicked: root.startClicked()
        }
    }
}
