import QtQuick 2.7
import QtQuick.Controls 2.1
import "Style"
import "Views" as ClefUI

ApplicationWindow {
    id: window

	visible: true
    title: "Clef UI"
	minimumWidth: 400
	minimumHeight: 680

    ClefUI.Login {
        anchors.fill: parent
    }

//    ClefUI.BackButton {
//        id: backButton
//        anchors.top: parent.top
//        anchors.left: parent.left
//        anchors.margins: Style.defaultMargin
//    }

//    ClefUI.Disclaimer {
//        anchors.left: parent.left
//        anchors.right: parent.right
//        anchors.top: backButton.bottom
//        anchors.margins: 54
//    }

//    ClefUI.DirectionField {
//        id: input
//        anchors.left: parent.left
//        anchors.right: parent.right
//        anchors.verticalCenter: parent.verticalCenter
//        anchors.margins: 2 * Style.defaultMargin

//        title: qsTr("enter path to clef") // + Translation.trigger
////        password: true
//        onTextChanged: console.log("text " + text)
//    }

//    ClefUI.TextField {
//        anchors.left: parent.left
//        anchors.right: parent.right
//        anchors.top: input.bottom
//        anchors.margins: 2 * Style.defaultMargin

//        title: qsTr("Alternative") // + Translation.trigger
//        text: loginContext.error
//    }

//    ClefUI.Text {
////        id: text
//        anchors.left: parent.left
//        anchors.right: parent.right
//        anchors.top: input.bottom
//        anchors.margins: 2 * Style.defaultMargin

//        title: qsTr("signer SHA256 HASH") // + Translation.trigger
//        //text: loginContext.error
//    }

//    ClefUI.Footer {
//        anchors.left: parent.left
//        anchors.right: parent.right
//        anchors.bottom: parent.bottom
//    }

//    Component.onCompleted: {
//        input.forceActiveFocus()
//    }
}
