import QtQuick 2.7
import QtQuick.Controls 2.1
import "../Style"

Item {
    id: root

    implicitHeight: 24
    implicitWidth: 112

    property int defaultIndex: 0
    property alias currentIndex: combobox.currentIndex
    property alias currentText: combobox.currentText

    onDefaultIndexChanged: combobox.currentIndex = defaultIndex

    ComboBox {
        id: combobox
        anchors.fill: parent
        font.pointSize: 13
        font.bold: true
        font.family: "Verdana"
        currentIndex: 2

        model: ListModel {
            ListElement { text: "WEI" }
            ListElement { text: "GWEI" }
            ListElement { text: "ETH" }
        }

        onCurrentIndexChanged: approvetx.setUnit("VALUE", currentIndex)

        background: Rectangle {
            x:0
            y:0
            width: parent.width
            height:parent.height
            color: "#efefef"
            radius: 0
            border.color: "#393939"
            border.width: 0
        }
    }
}
