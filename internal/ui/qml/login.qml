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

            Text {
                id: transport
                x: 155
                y: 22
                width: 90
                height: 20
                text: "Clef"
                font.capitalization: Font.AllUppercase
                font.family: "Verdana"
                font.bold: true
                font.pixelSize: 14
                verticalAlignment: Text.AlignVCenter
                horizontalAlignment: Text.AlignHCenter
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
                id: cta
                x: 125
                y: 11
                width: 150
                height: 36
                opacity: ctxObject.isValid ? 1 : 0.5

                MouseArea {
                    id: ctama
                    x: 0
                    y: 0
                    width: parent.width
                    height: parent.height
                    hoverEnabled: true
                    cursorShape: Qt.PointingHandCursor
                    onClicked: ctxObject.clicked()
                    onPressed: cta.state = "active"
                    onReleased: cta.state = ""
                }

                states: [
                    State {
                        name: "valid"
                        PropertyChanges { target: cta; opacity: 1 }
                    },
                    State {
                        name: "invalid"
                        PropertyChanges { target: cta; opacity: 0.5 }
                    },
                    State {
                        name: "active"
                        PropertyChanges { target: ctabg; opacity: 1 }
                    },
                    State {
                        name: "inactive"
                        PropertyChanges { target: ctabg; opacity: 0.8 }
                    }
                ]

                transitions: [
                    Transition {
                        NumberAnimation {
                            duration: 200
                            properties: "opacity"
                        }
                    }
                ]

                contentItem: Text {
                    x: 1
                    y: 1
                    width: 150
                    height: 24
                    color: "#ffffff"
                    text: qsTr("Start")
                    font.bold: true
                    font.family: "Arial"
                    horizontalAlignment: Text.AlignHCenter
                    font.pixelSize: 14
                    font.capitalization: Font.AllUppercase
                    verticalAlignment: Text.AlignVCenter
                    elide: Text.ElideRight
                }

                background: Rectangle {
                    id: ctabg
                    width: parent.width
                    height: parent.height
                    color: "#48b877"
                    opacity: ctama.containsMouse ? 0.6 : 0.8
                    radius: 18
                    border.width: 0
                    PropertyAnimation on opacity {
                        duration: 200
                        easing: Easing.InOutElastic
                    }
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
            id: logo
            x: 158
            y: 137
            width: 84
            height: 115
            color: "#ffffff"
            border.color: "#00000000"

            states: [
                State {
                    name: "valid"
                    PropertyChanges { target: text1; color: "#76e09f" }
                },
                State {
                    name: "invalid"
                    PropertyChanges { target: text1; color: "#e07688" }
                }
            ]

            transitions: Transition {
                ColorAnimation { duration: 200 }
            }

            Text {
                id: text1
                x: -2
                y: -11
                width: 84
                height: 115
                color: "#76e09f"
                text: "𝄞"
                font.wordSpacing: 0
                style: Text.Normal
                font.family: "Tahoma"
                verticalAlignment: Text.AlignVCenter
                horizontalAlignment: Text.AlignHCenter
                font.pixelSize: 128
            }
        }

        Text {
            id: text2
            x: 37
            y: 284
            color: "#646464"
            text: qsTr("Enter your go-ethereum directory")
            font.capitalization: Font.AllUppercase
            font.pixelSize: 12
        }

        Rectangle {
            id: rectangle3
            x: 29
            y: 299
            width: 342
            height: 28
            color: "#efefef"
            radius: 2

            function onChange(text) {
                ctxObject.edited(text)
                logo.state = ctxObject.isValid ? "valid" : "invalid"
                cta.state = ctxObject.isValid ? "valid" : "invalid"
            }

            TextInput {
                id: textInput
                x: 7
                y: 8
                width: 329
                height: 14
                color: "#646464"
                text: ctxObject.gopath
                font.family: "Verdana"
                cursorVisible: true
                onTextChanged: rectangle3.onChange(text)
                horizontalAlignment: Text.AlignLeft
                verticalAlignment: Text.Center
                font.pixelSize: 14
            }
        }

        Text {
            id: text3
            x: 37
            y: 339
            color: "#646464"
            text: qsTr("Signer SHA256 Hash")
            font.pixelSize: 12
            font.capitalization: Font.AllUppercase
        }

        Rectangle {
            id: rectangle4
            x: 29
            y: 354
            width: 342
            height: 69
            color: "#efefef"
            radius: 2
            TextInput {
                id: textInput1
                x: 7
                y: 8
                width: 327
                height: 53
                color: "#646464"
                text: ctxObject.binaryHash
                font.family: "Courier"
                activeFocusOnPress: false
                readOnly: true
                horizontalAlignment: Text.AlignLeft
                font.pixelSize: 14
                verticalAlignment: Text.Center
                wrapMode: TextInput.WrapAnywhere
            }
        }






    }

}
