import QtQuick 2.7	
import QtQuick.Controls 2.1	
import "../Style"

Item {	
    id: root	

    implicitHeight: 36	
    implicitWidth: 150	

    property string text: "click me"	
    property color color: Style.buttons.backgroundColor	

    signal clicked()	
    signal pressed()	
    signal pressAndHold()	
    signal released()	

    Button {	
        anchors.fill: parent	
        background: Rectangle {	
            anchors.fill: parent	
            color: root.color	
            radius: height / 2	
        }	
        contentItem: Text {	
            anchors.fill: parent	
            horizontalAlignment: Text.AlignHCenter	
            verticalAlignment: Text.AlignVCenter	
            text: root.text	
            color: Style.buttons.textColor	
            font: Style.buttons.font	
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