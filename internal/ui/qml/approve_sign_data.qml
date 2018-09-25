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

                MouseArea {
                    x: 0
                    y: 0
                    width: parent.width
                    height: parent.height
                    cursorShape: Qt.PointingHandCursor
                    onClicked: ctxObject.back()
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

        Rectangle {
            id: rectangle11
            x: 0
            y: 70
            width: 400
            height: 46
            color: "#00000000"

            Text {
                id: text6
                x: 12
                y: 4
                height: 10
                text: qsTr("From")
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
                width: 342
                height: 24
                color: "#efefef"
                radius: 2
                TextInput {
                    id: textInput
                    x: 4
                    y: 7
                    width: parent.width - 8
                    height: parent.height - 12
                    text: ctxObject.from
                    readOnly: true
                    font.family: "Courier"
                    horizontalAlignment: Text.AlignLeft
                    verticalAlignment: Text.AlignHCenter
                    font.pixelSize: 12
                }
            }

            Image {
                id: image
                x: 356
                y: 6
                width: 36
                height: 36
                source: ctxObject.fromSrc
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

                MouseArea {
                    x: 0
                    y: 0
                    width: parent.width
                    height: parent.height
                    cursorShape: Qt.PointingHandCursor
                    onClicked: ctxObject.clicked(1)
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
                id: control1
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
                    onClicked: pwInput2.state = "show"
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

        Rectangle {
            id: rectangle21
            x: 0
            y: 244
            width: 400
            height: 166
            color: "#00000000"
            Text {
                id: text14
                x: 12
                y: 4
                height: 10
                color: "#747474"
                text: qsTr("Message")
                font.bold: false
                font.pixelSize: 10
                font.family: "Verdana"
                font.capitalization: Font.AllUppercase
            }

            Rectangle {
                id: rectangle22
                x: 8
                y: 18
                width: 384
                height: 140
                color: "#efefef"
                radius: 2

                TextArea {
                    id: textArea
                    x: -4
                    y: 0
                    width: parent.width + 4
                    height: parent.height
                    text: ctxObject.message
                    font.pointSize: 12
                    font.family: "Courier"
                    readOnly: true
                }
            }
        }

        Rectangle {
            id: rectangle25
            x: 0
            y: 416
            width: 400
            height: 166
            color: "#00000000"
            Text {
                id: text15
                x: 12
                y: 4
                height: 10
                color: "#747474"
                text: qsTr("Raw Data")
                font.bold: false
                font.pixelSize: 10
                font.family: "Verdana"
                font.capitalization: Font.AllUppercase
            }

            Rectangle {
                id: rectangle24
                x: 8
                y: 18
                width: 384
                height: 140
                color: "#efefef"
                radius: 2
                TextArea {
                    id: textArea1
                    x: -4
                    y: 0
                    width: parent.width + 4
                    height: parent.height
                    text: ctxObject.rawData
                    font.family: "Courier"
                    font.pointSize: 12
                    readOnly: true
                }
            }
        }

        Rectangle {
            id: rectangle26
            x: 0
            y: 122
            width: 400
            height: 116
            color: "#00000000"
            Text {
                id: text16
                x: 12
                y: 4
                height: 10
                color: "#747474"
                text: qsTr("Hash")
                font.bold: false
                font.pixelSize: 10
                font.family: "Verdana"
                font.capitalization: Font.AllUppercase
            }

            Rectangle {
                id: rectangle27
                x: 8
                y: 18
                width: 384
                height: 90
                color: "#efefef"
                radius: 2
                TextArea {
                    id: textArea2
                    x: -4
                    y: 0
                    width: parent.width + 4
                    height: parent.height
                    text: ctxObject.hash
                    wrapMode: Text.WrapAnywhere
                    font.family: "Courier"
                    font.pointSize: 12
                    readOnly: true
                }
            }
        }


    }
    Rectangle {
        id: pwInput2
        x: 0
        y: 0
        width: 400
        height: 680
        color: "#00000000"
        opacity: 0.5
        visible: false

        states: [
            State {
                name: "show"
                PropertyChanges { target: pwInput2; opacity: 1; visible: true}
            },
            State {
                name: "hide"
                PropertyChanges { target: pwInput2; opacity: 0.5; visible: false}
            }
        ]

        transitions: Transition {
            NumberAnimation {
                duration: 150
                properties: "opacity"
            }
        }


        Rectangle {
            id: pwInputBg
            x: 0
            y: 0
            width: 400
            height: 680
            color: "#000000"
            opacity: 0.8
        }



        Rectangle {
            id: rectangle17
            x: 33
            y: 215
            width: 334
            height: 250
            color: "#ffffff"
            radius: 6

            Text {
                id: text3
                x: 105
                y: 89
                color: "#646464"
                text: qsTr("Enter Password")
                font.family: "Verdana"
                font.capitalization: Font.AllUppercase
            }

            Rectangle {
                id: rectangle23
                x: 31
                y: 111
                width: 274
                height: 29
                color: "#efefef"
                radius: 2

                TextInput {
                    id: textInput6
                    x: 0
                    y: 8
                    width: 274
                    height: 16
                    color: "#000000"
                    text: ctxObject.password
                    horizontalAlignment: Text.AlignHCenter
                    echoMode: TextInput.Password
                    font.pixelSize: 16
                    onTextChanged: ctxObject.edited("password", text)
                }
            }

            Button {
                id: button1
                x: 185
                y: 184
                width: 100
                height: 36
                function onClick() {
                    ctxObject.clicked(2)
                    pwInput2.state = "hide"
                }

                background: Rectangle {
                    x: 0
                    y: 0
                    width: parent.width
                    height: parent.height
                    color: "#b4b4b4"
                    radius: 18

                    MouseArea {
                        x: 0
                        y: 0
                        width: parent.width
                        height: parent.height
                        cursorShape: Qt.PointingHandCursor
                        onClicked: button1.onClick()
                    }
                }

                contentItem: Text {
                    color: "#ffffff"
                    text: "Confirm"
                    font.capitalization: Font.AllUppercase
                    font.bold: true
                    font.family: "Verdana"
                    verticalAlignment: Text.AlignVCenter
                    horizontalAlignment: Text.AlignHCenter
                }
            }

            Button {
                id: button2
                x: 50
                y: 184
                width: 100
                height: 36

                function onClick() {
                    pwInput2.state = "hide"
                    ctxObject.edited("password", "")
                }

                contentItem: Text {
                    color: "#b4b4b4"
                    text: "Cancel"
                    font.family: "Verdana"
                    font.bold: true
                    horizontalAlignment: Text.AlignHCenter
                    font.capitalization: Font.AllUppercase
                    verticalAlignment: Text.AlignVCenter
                }
                background: Rectangle {
                    x: 0
                    y: 0
                    width: parent.width
                    height: parent.height
                    color: "#ffffff"
                    radius: 18
                    border.color: "#b4b4b4"
                    border.width: 2
                    MouseArea {
                        x: 0
                        y: 0
                        width: parent.width
                        height: parent.height
                        onClicked: button2.onClick()
                        cursorShape: Qt.PointingHandCursor
                    }
                }
            }
        }
    }
}
