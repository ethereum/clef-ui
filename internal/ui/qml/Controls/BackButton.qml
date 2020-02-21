import QtQuick 2.7
import QtQuick.Controls 2.1
import "../Style"

Item {
    id: root

    implicitHeight: 32
    implicitWidth: 32

    signal clicked()
    signal pressed()
    signal pressAndHold()
    signal released()

    Button {
        anchors.fill: parent
        contentItem: Text {
            text: "<"
            anchors.centerIn: parent
            font.family: "Verdana"
            font.pixelSize: 16
            color: "#c9c9c9"
        }
        background: Rectangle {
            anchors.fill: parent
            radius: height / 2
            border.color: "#c9c9c9"
            border.width: 1
        }

        onClicked: root.clicked()
        onPressed: root.pressed()
        onPressAndHold: root.pressAndHold()
        onReleased: root.released()

        // Button in Qt 5.12.0 doesn't seem to give access cursorShape property
        MouseArea {
            anchors.fill: parent
            hoverEnabled: Style.buttons.hoverEnabled
            cursorShape: Style.buttons.cursorShape
            acceptedButtons: Qt.NoButton
        }
    }
}
