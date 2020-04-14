import QtQuick 2.7
import QtQuick.Controls 2.1
import "../Style"

Item {
    id: root
    implicitHeight: column.implicitHeight
    implicitWidth: 240

    property string title: "" // better property name?
    property string text: ""
    property var wrapMode: Text.NoWrap

    Column {
        id: column

        anchors.left: parent.left
        anchors.right: parent.right
        spacing: 2

        Text {
            anchors.left: parent.left
            anchors.margins: 8
            height: visible ? contentHeight : 0
            text: root.title
            font: Style.texts.titleFont
            color: Style.texts.titleColor
            visible: root.title !== "" && root.text !== ""
        }

        Rectangle {
            anchors.left: parent.left
            anchors.right: parent.right
            height: root.text !== "" ? textField.contentHeight + 14 : 0
            radius: Style.texts.backgroundRadius
            color: Style.texts.backgroundColor

            Text {
                id: textField
                anchors.left: parent.left
                anchors.right: parent.right
                anchors.margins: 8
                anchors.verticalCenter: parent.verticalCenter
                font: Style.texts.font
                text: root.text
                color: Style.texts.color
                wrapMode: root.wrapMode
            }
        }
    }
}

