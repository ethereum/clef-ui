import QtQuick 2.7
import QtQuick.Controls 2.1
import "../Style"

Item {
    id: root

    implicitHeight: column.implicitHeight
    implicitWidth: 240

    property string title: "" // better property name?
    property alias text: input.text
    property bool password: false

    onActiveFocusChanged: {
        if (activeFocus)
            input.forceActiveFocus()
    }

    Column {
        id: column

        anchors.left: parent.left
        anchors.right: parent.right
        spacing: 2

        Text {
            anchors.left: parent.left
            anchors.margins: 8
            height: root.title !== "" ? contentHeight : 0
            text: root.title
            font: Style.textFields.titleFont
            color: Style.textFields.titleColor
        }

        Rectangle {
            anchors.left: parent.left
            anchors.right: parent.right
            height: input.contentHeight + 14
            radius: Style.textFields.backgroundRadius
            color: Style.textFields.backgroundColor

            TextInput {
                id: input
                anchors.left: parent.left
                anchors.right: parent.right
                anchors.margins: 8
                anchors.verticalCenter: parent.verticalCenter
                height: font.pixelSize
                font: Style.textFields.font

                echoMode: root.password ? TextInput.Password : TextInput.Normal
                passwordMaskDelay: 1000
            }
        }
    }
}
