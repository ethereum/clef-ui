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

        Headerfield {
            id: rectangle1
            x: 0
            y: 0
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
                text: qsTr("Password")
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
                    text: ctxObject.password
                    onTextChanged: ctxObject.passwordEdited(text)
                }
            }


        }

        Footer {
            id: row_footer
            x: 0
            y: 625
            onApprove: {
                ctxObject.approve()
            }
            onReject: {
                ctxObject.reject()
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
                text: qsTr("Clef will generate a new private key, encrypts it according to web3 keystore spec and stores it in the keystore directory. Please create a backup of the keystore. If the keystore is lost there is no method of retrieving lost accounts.")
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
                text: qsTr("Confirm Password")
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
                    text: ctxObject.confirmPassword
                    onTextChanged: ctxObject.confirmPasswordEdited(text)
                }
            }
        }


    }
}
