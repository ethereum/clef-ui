import QtQuick 2.4
import QtQuick.Controls 2.3

Item {
    id: item1
    width: 400
    height: 680

    Rectangle {
        id: rectangle
        x: 0
        y: 0
        width: 400
        height: 680
        color: "#ffffff"

        Rectangle {
            id: rectangle1
            x: 0
            y: 0
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
                onClicked: ctxObject.back()
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

        Rectangle {
            id: rectangle11
            x: 27
            y: 269
            width: 347
            height: 58
            color: "#00000000"

            Text {
                id: text6
                x: 12
                y: 4
                height: 10
                text: qsTr("Old Password")
                font.capitalization: Font.AllUppercase
                font.bold: false
                font.family: "Verdana"
                font.pixelSize: 10
                color: "#747474"
            }

            Rectangle {
                id: rectangle5
                x: 8
                y: 18
                width: 331
                height: 24
                color: "#efefef"
                radius: 2
                TextInput {
                    id: textInput
                    x: 4
                    y: 7
                    width: parent.width - 8
                    height: parent.height - 12
                    font.family: "Courier"
                    horizontalAlignment: Text.AlignLeft
                    verticalAlignment: Text.AlignHCenter
                    font.pixelSize: 12
                    echoMode: TextInput.Password
                    passwordMaskDelay: 1000
                    text: ctxObject.oldPassword
                    onTextChanged: ctxObject.oldPasswordEdited(text)
                }
            }


        }

        Rectangle {
            id: row
            x: 0
            y: 625
            width: 400
            height: 55
            color: "#ffffff"

            Button {
                id: control
                x: 25
                y: 10
                height: 36
                width: 150
                onClicked: ctxObject.clicked(1)
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
                id: control1
                x: 225
                y: 10
                height: 36
                width: 150
                onClicked: ctxObject.clicked(2)
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

        Rectangle {
            id: rectangle7
            x: 56
            y: 109
            width: 288
            height: 121
            color: "#ffffff"
            radius: 4
            border.color: "#c9c9c9"

            Text {
                id: text9
                x: 8
                y: 8
                width: 272
                height: 105
                color: "#747474"
                text: qsTr("Clef will import a new private key, encrypts it according to web3 keystore spec and stores it in the keystore directory. Please create a backup of the keystore. If the keystore is lost there is no method of retrieving lost accounts.")
                verticalAlignment: Text.AlignVCenter
                wrapMode: Text.WordWrap
                font.pixelSize: 12
                font.bold: false
                horizontalAlignment: Text.AlignHCenter
                font.family: "Verdana"
                font.capitalization: Font.MixedCase
            }
        }

        Rectangle {
            id: rectangle12
            x: 27
            y: 333
            width: 347
            height: 58
            color: "#00000000"
            Text {
                id: text7
                x: 12
                y: 4
                height: 10
                color: "#747474"
                text: qsTr("New Password")
                font.bold: false
                font.pixelSize: 10
                font.family: "Verdana"
                font.capitalization: Font.AllUppercase
            }

            Rectangle {
                id: rectangle6
                x: 8
                y: 18
                width: 331
                height: 24
                color: "#efefef"
                radius: 2
                TextInput {
                    id: textInput1
                    x: 4
                    y: 7
                    width: parent.width - 8
                    height: parent.height - 12
                    font.pixelSize: 12
                    horizontalAlignment: Text.AlignLeft
                    font.family: "Courier"
                    verticalAlignment: Text.AlignHCenter
                    echoMode: TextInput.Password
                    passwordMaskDelay: 1000
                    text: ctxObject.password
                    onTextChanged: ctxObject.passwordEdited(text)
                }
            }
        }

        Rectangle {
            id: rectangle13
            x: 27
            y: 397
            width: 347
            height: 58
            color: "#00000000"
            Text {
                id: text8
                x: 12
                y: 4
                height: 10
                color: "#747474"
                text: qsTr("Confirm New Password")
                font.bold: false
                font.capitalization: Font.AllUppercase
                font.family: "Verdana"
                font.pixelSize: 10
            }

            Rectangle {
                id: rectangle9
                x: 8
                y: 18
                width: 331
                height: 24
                color: "#efefef"
                radius: 2
                TextInput {
                    id: textInput2
                    x: 4
                    y: 7
                    width: parent.width - 8
                    height: parent.height - 12
                    text: ctxObject.confirmPassword
                    font.family: "Courier"
                    passwordMaskDelay: 1000
                    verticalAlignment: Text.AlignHCenter
                    font.pixelSize: 12
                    echoMode: TextInput.Password
                    horizontalAlignment: Text.AlignLeft
                    onTextChanged: ctxObject.confirmPasswordEdited(text)
                }
            }
        }


    }
}
