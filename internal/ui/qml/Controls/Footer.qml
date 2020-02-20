import QtQuick 2.7
import QtQuick.Controls 2.1
import "../Style"
import "." as ClefUI

Item {
    id: root

    implicitHeight: 55
    implicitWidth: row.implicitWidth

    // StackView?

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
//            onReleased: console.log("test")
        }

        ClefUI.Button {
            text: qsTr("approve") // + Translation.trigger
            color: Style.buttons.validationColor
        }
    }
}
