import QtQuick 2.4
import QtQuick.Controls 2.3

Rectangle {
    id: root
    width: 400
    height: 55
    color: "#ffffff"
    signal approve()
    signal reject()

    Button {
        id: button_reject
        x: 25
        y: 10
        height: 36
        width: 150
        
        MouseArea {
            x: 0
            y: 0
            width: parent.width
            height: parent.height
            onClicked: root.reject()
        }
        
        contentItem: Text {
            x: 1
            y: 1
            width: 150
            height: 24
            color: "#ffffff"
            text: qsTr("Reject")
            font.bold: true
            font.family: "Arial"
            font.capitalization: Font.AllUppercase
            horizontalAlignment: Text.AlignHCenter
            verticalAlignment: Text.AlignVCenter
            elide: Text.ElideRight
            font.pixelSize: 14
        }
        
        background: Rectangle {
            color: "#5b5b5b"
            width: parent.width
            height: parent.height
            border.width: 0
            radius: 18
        }
        
    }
    
    Button {
        id: button_approve
        x: 225
        y: 10
        height: 36
        width: 150
        
        MouseArea {
            x: 0
            y: 0
            width: parent.width
            height: parent.height
            cursorShape: Qt.PointingHandCursor
            onClicked: root.approve()
        }
        contentItem: Text {
            x: 1
            y: 1
            width: 150
            height: 24
            color: "#ffffff"
            text: qsTr("Approve")
            font.bold: true
            font.family: "Arial"
            font.capitalization: Font.AllUppercase
            horizontalAlignment: Text.AlignHCenter
            verticalAlignment: Text.AlignVCenter
            elide: Text.ElideRight
            font.pixelSize: 14
        }
        
        background: Rectangle {
            color: "#48b877"
            width: parent.width
            height: parent.height
            border.width: 0
            radius: 18
        }
        
    }
    
    Rectangle {
        id: rectangle16
        x: 0
        y: 0
        width: 400
        height: 1
        color: "#efefef"
    }
    
}
