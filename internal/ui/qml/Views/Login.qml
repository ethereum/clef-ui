import QtQuick 2.7
import QtQuick.Controls 2.1
import "../Style"
import "../Controls" as Controls

Item {
    id: root

    implicitHeight: 200
    implicitWidth: 100

    onActiveFocusChanged: {
        if (activeFocus)
            input.forceActiveFocus()
    }

    Text {
        anchors.top: parent.top
        anchors.topMargin: 24
        anchors.horizontalCenter: parent.horizontalCenter
        text: "Clef"
        font.capitalization: Font.AllUppercase
        font.family: "Verdana"
        font.bold: true
        font.pixelSize: 14
    }

    Text {
        id: logo
        anchors.top: parent.top
        anchors.topMargin: 64
        anchors.horizontalCenter: parent.horizontalCenter
        color: LoginContext.error === "" ? "#76e09f" : "#e07688"
        text: "𝄞"
        style: Text.Normal
        font.family: "Tahoma"
        font.pixelSize: 128
    }

    Text {
        anchors.top: logo.bottom
        anchors.horizontalCenter: parent.horizontalCenter
        anchors.topMargin: 16
        color: logo.color
        visible: text !== ""
        text: LoginContext.error
    }

    Timer {
        id: checkPathTimer
        interval: 700
        onTriggered: LoginContext.checkPath(input.text)
    }

    Controls.DirectionField {
        id: input
        anchors.left: parent.left
        anchors.right: parent.right
        anchors.verticalCenter: parent.verticalCenter
        anchors.margins: 2 * Style.defaultMargin

        title: qsTr("enter path to clef") // + Translation.trigger
        onTextChanged: checkPathTimer.restart()
    }

    Controls.Text {
        anchors.left: parent.left
        anchors.right: parent.right
        anchors.top: input.bottom
        anchors.margins: 2 * Style.defaultMargin

        visible: text !== ""

        title: qsTr("signer SHA256 HASH") // + Translation.trigger
        text: LoginContext.binaryHash
    }

    Controls.Footer {
        id: footer
        anchors.left: parent.left
        anchors.right: parent.right
        anchors.bottom: parent.bottom
        enabled: LoginContext.error === "" && LoginContext.clefPath !== ""
        startButtonVisible: true
        onStartClicked: LoginContext.start()
    }
}
