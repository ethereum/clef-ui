import QtQuick 2.4
import QtQuick.Controls 2.3

Rectangle {
    id: rectangle1
    width: 400
    height: 64
    color: "#ffffff"
    border.width: 0
    
    Rectangle {
        id: rectangle10
        x: 0
        y: 64
        width: 400
        height: 1
        color: "#efefef"
    }
    
    Button {
        id: button
        x: 8
        y: 16
        width: 32
        height: 32
        z: 1
        
        MouseArea {
            x: 0
            y: 0
            width: parent.width
            height: parent.height
            onClicked: ctxObject.back()
            cursorShape: Qt.PointingHandCursor
        }
        
        contentItem: Text {
            text: qsTr("<")
            verticalAlignment: Text.AlignVCenter
            horizontalAlignment: Text.AlignHCenter
            font.family: "Verdana"
            font.pointSize: 16
            color: "#c9c9c9"
        }
        background: Rectangle {
            height: 32
            radius: 16
            border.color: "#c9c9c9"
            border.width: 1
            width: 32
        }
    }
    
    Rectangle {
        id: rectangle4
        x: 290
        y: 8
        width: 102
        height: 19
        color: "#00000000"
        z: 1
        border.width: 0
        border.color: "#00000000"
        
        Rectangle {
            id: rectangle2
            x: 0
            y: 8
            width: 100
            height: 1
            color: "#c9c9c9"
            radius: 1
            opacity: 1
        }
        
        Text {
            id: text1
            x: 0
            y: 0
            width: 100
            height: 8
            color: "#646464"
            text: qsTr("REMOTE")
            font.bold: false
            styleColor: "#747474"
            z: 1
            font.family: "Verdana"
            verticalAlignment: Text.AlignBottom
            horizontalAlignment: Text.AlignLeft
            font.pixelSize: 8
        }
        
        Text {
            id: remote
            x: 0
            y: 10
            width: 100
            height: 8
            color: "#747474"
            text: ctxObject.remote
            styleColor: "#646464"
            font.family: "Verdana"
            verticalAlignment: Text.AlignTop
            horizontalAlignment: Text.AlignLeft
            font.pixelSize: 8
        }
    }
    
    Text {
        id: transport
        x: 155
        y: 22
        width: 90
        height: 20
        text: ctxObject.transport
        font.bold: true
        font.pixelSize: 12
        verticalAlignment: Text.AlignVCenter
        horizontalAlignment: Text.AlignHCenter
    }
    
    Rectangle {
        id: rectangle8
        x: 290
        y: 39
        width: 102
        height: 19
        color: "#00000000"
        border.width: 0
        Rectangle {
            id: rectangle3
            x: 0
            y: 8
            width: 100
            height: 1
            color: "#c9c9c9"
            radius: 1
            opacity: 1
        }
        
        Text {
            id: text2
            x: 0
            y: 0
            width: 100
            height: 8
            color: "#646464"
            text: qsTr("LOCAL ENDPOINT")
            styleColor: "#747474"
            font.bold: false
            font.pixelSize: 8
            horizontalAlignment: Text.AlignLeft
            font.family: "Verdana"
            verticalAlignment: Text.AlignBottom
            z: 1
        }
        
        Text {
            id: remote1
            x: 0
            y: 10
            width: 100
            height: 8
            color: "#747474"
            text: ctxObject.endpoint
            styleColor: "#646464"
            font.pixelSize: 8
            horizontalAlignment: Text.AlignLeft
            font.family: "Verdana"
            verticalAlignment: Text.AlignTop
        }
        border.color: "#00000000"
        z: 1
    }
}
