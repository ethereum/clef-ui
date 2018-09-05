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
        color: "#ededed"

        Rectangle {
            id: rectangle1
            x: 0
            y: 0
            width: 400
            height: 140
            color: "#62a0dd"

            Rectangle {
                id: rectangle10
                x: 144
                y: 14
                width: 113
                height: 112
                color: "#ffeac9"
                radius: 8
            }
        }

        Rectangle {
            id: rectangle8
            x: 0
            y: 138
            width: 400
            height: 31
            color: "#b2d6f1"
            border.color: "#00000000"

            Rectangle {
                id: rectangle4
                x: 4
                y: 6
                width: 145
                height: 20
                color: "#00000000"
                border.width: 0
                border.color: "#00000000"

                Rectangle {
                    id: rectangle2
                    x: 0
                    y: 0
                    width: 53
                    height: 20
                    color: "#04273f"
                    radius: 1
                    opacity: 1

                    Text {
                        id: text1
                        x: 0
                        y: 0
                        width: parent.width
                        height: parent.height
                        color: "#ffffff"
                        text: qsTr("REMOTE")
                        font.bold: true
                        styleColor: "#00000000"
                        z: 1
                        font.family: "Verdana"
                        verticalAlignment: Text.AlignVCenter
                        horizontalAlignment: Text.AlignHCenter
                        font.pixelSize: 7
                    }
                }

                Rectangle {
                    id: rectangle3
                    x: 53
                    y: 0
                    width: 92
                    height: parent.height
                    color: "#d2ebff"
                    opacity: 1

                    Text {
                        id: remote
                        x: 0
                        y: 0
                        width: parent.width
                        height: parent.height
                        text: ctxObject.remote
                        styleColor: "#00000000"
                        font.family: "Verdana"
                        verticalAlignment: Text.AlignVCenter
                        horizontalAlignment: Text.AlignHCenter
                        font.pixelSize: 8
                    }
                }
            }

            Rectangle {
                id: rectangle5
                x: 251
                y: 6
                width: 145
                height: 20
                color: "#00000000"
                Rectangle {
                    id: rectangle6
                    x: 0
                    y: 0
                    width: 53
                    height: 20
                    color: "#04273f"
                    radius: 1
                    Text {
                        id: text4
                        x: 0
                        y: 0
                        width: parent.width
                        height: parent.height
                        color: "#ffffff"
                        text: qsTr("ENDPOINT")
                        font.pixelSize: 7
                        horizontalAlignment: Text.AlignHCenter
                        font.bold: true
                        verticalAlignment: Text.AlignVCenter
                        font.family: "Verdana"
                        styleColor: "#00000000"
                        z: 1
                    }
                    opacity: 1
                }

                Rectangle {
                    id: rectangle7
                    x: 53
                    y: 0
                    width: 92
                    height: parent.height
                    color: "#d2ebff"
                    Text {
                        id: endpoint
                        x: 0
                        y: 0
                        width: parent.width
                        height: parent.height
                        text: ctxObject.endpoint
                        font.pixelSize: 8
                        horizontalAlignment: Text.AlignHCenter
                        verticalAlignment: Text.AlignVCenter
                        font.family: "Verdana"
                        styleColor: "#00000000"
                    }
                    opacity: 1
                }
                border.width: 0
                border.color: "#00000000"
            }

            Text {
                id: transport
                x: 155
                y: 6
                width: 90
                height: 20
                text: ctxObject.transport
                font.bold: true
                font.pixelSize: 8
                verticalAlignment: Text.AlignVCenter
                horizontalAlignment: Text.AlignHCenter
            }
        }

        Rectangle {
            id: rectangle11
            x: 0
            y: 175
            width: 400
            height: 62
            color: "#00000000"

            Text {
                id: text6
                x: 16
                y: 8
                text: qsTr("From")
                font.bold: true
                font.family: "Verdana"
                font.pixelSize: 11
            }

            Rectangle {
                id: rectangle9
                x: 8
                y: 28
                width: parent.width - 16

                height: 34
                color: "#ffffff"
                radius: 4

                Text {
                    id: text7
                    x: 8
                    y: 8
                    width: parent.width - 16
                    height: parent.height - 16
                    color: "#202020"
                    text: ctxObject.from
                    font.letterSpacing: 0
                    font.wordSpacing: 0
                    styleColor: "#00000000"
                    verticalAlignment: Text.AlignVCenter
                    horizontalAlignment: Text.AlignLeft
                    font.family: "Courier"
                    font.pixelSize: 14
                }
            }
        }

        Rectangle {
            id: rectangle16
            x: 0
            y: 493
            width: 400
            height: 96
            color: "#00000000"
            Text {
                id: text12
                x: 16
                y: 8
                text: qsTr("Data")
                font.pixelSize: 11
                font.bold: true
                font.family: "Verdana"
            }

            Rectangle {
                id: rectangle17
                x: 8
                y: 28
                width: 384
                height: 60
                color: "#ffffff"
                radius: 4
                Text {
                    id: text13
                    x: 8
                    y: 8
                    width: parent.width - 16
                    height: parent.height - 16
                    color: "#202020"
                    text: ctxObject.data
                    font.pixelSize: 14
                    fontSizeMode: Text.FixedSize
                    scale: 1
                    lineHeight: 1
                    font.capitalization: Font.MixedCase
                    horizontalAlignment: Text.AlignLeft
                    verticalAlignment: Text.AlignTop
                    wrapMode: Text.WordWrap
                    font.family: "Courier"
                    rotation: 0
                }
            }
        }

        Rectangle {
            id: row
            x: 0
            y: 595
            width: 400
            height: 85
            color: "#ffffff"

            Button {
                id: control
                x: 8
                y: 8
                height: parent.height - 16
                width: parent.width/2 - 12
                onClicked: ctxObject.clicked(1)
                contentItem: Text {
                    x: 1
                    y: 1
                    width: parent.width
                    height: parent.height
                    color: "#ffffff"
                    text: qsTr("Reject")
                    font.family: "Verdana"
                    font.capitalization: Font.AllUppercase
                    font.bold: true
                    horizontalAlignment: Text.AlignHCenter
                    verticalAlignment: Text.AlignVCenter
                    elide: Text.ElideRight
                    font.pixelSize: 24
                }

                background: Rectangle {
                    color: "#bbbbbb"
                    width: parent.width
                    height: parent.height
                    border.width: 0
                    radius: 4
                }

            }

            Button {
                id: control1
                x: parent.width/2 + 4
                y: 8
                height: parent.height - 16
                width: parent.width/2 - 12
                onClicked: ctxObject.clicked(2)
                contentItem: Text {
                    x: 1
                    y: 1
                    width: parent.width
                    height: parent.height
                    color: "#ffffff"
                    text: qsTr("Approve")
                    font.pixelSize: 24
                    horizontalAlignment: Text.AlignHCenter
                    font.capitalization: Font.AllUppercase
                    font.bold: true
                    verticalAlignment: Text.AlignVCenter
                    elide: Text.ElideRight
                    font.family: "Verdana"
                }
                background: Rectangle {
                    width: parent.width
                    height: parent.height
                    color: "#62a0dd"
                    radius: 4
                    border.width: 0
                }
            }

        }

        Rectangle {
            id: rectangle18
            x: 0
            y: 243
            width: 400
            height: 62
            color: "#00000000"
            Text {
                id: text14
                x: 16
                y: 8
                text: qsTr("To")
                font.bold: true
                font.pixelSize: 11
                font.family: "Verdana"
            }

            Rectangle {
                id: rectangle19
                x: 8
                y: 28
                width: parent.width - 16
                height: 34
                color: "#ffffff"
                radius: 4
                Text {
                    id: text15
                    x: 8
                    y: 8
                    width: parent.width - 16
                    height: parent.height - 16
                    color: "#202020"
                    text: ctxObject.to
                    font.letterSpacing: 0
                    font.wordSpacing: 0
                    styleColor: "#00000000"
                    font.pixelSize: 14
                    horizontalAlignment: Text.AlignLeft
                    font.family: "Courier"
                    verticalAlignment: Text.AlignVCenter
                }
            }
        }

        Rectangle {
            id: rectangle20
            x: 0
            y: 325
            width: 400
            height: 31
            color: "#00000000"
            Text {
                id: text16
                x: 16
                y: 8
                text: qsTr("Value")
                font.pixelSize: 11
                font.bold: true
                font.family: "Verdana"
            }

            Rectangle {
                id: rectangle21
                x: 90
                y: 0
                width: 302
                height: 30
                color: "#ffffff"
                radius: 4
                Text {
                    id: text17
                    x: 8
                    y: 8
                    width: parent.width - 16
                    height: parent.height - 16
                    color: "#202020"
                    text: ctxObject.value
                    font.wordSpacing: 0
                    font.letterSpacing: 0
                    styleColor: "#00000000"
                    font.pixelSize: 14
                    horizontalAlignment: Text.AlignLeft
                    font.family: "Courier"
                    verticalAlignment: Text.AlignVCenter
                }
            }
        }

        Rectangle {
            id: rectangle22
            x: 0
            y: 369
            width: 400
            height: 31
            color: "#00000000"
            Text {
                id: text18
                x: 16
                y: 8
                text: qsTr("Gas")
                font.bold: true
                font.pixelSize: 11
                font.family: "Verdana"
            }

            Rectangle {
                id: rectangle23
                x: 90
                y: 0
                width: 302
                height: 30
                color: "#ffffff"
                radius: 4
                Text {
                    id: text19
                    x: 8
                    y: 8
                    width: parent.width - 16
                    height: parent.height - 16
                    color: "#202020"
                    text: ctxObject.gas
                    font.letterSpacing: 0
                    font.wordSpacing: 0
                    styleColor: "#00000000"
                    font.pixelSize: 14
                    horizontalAlignment: Text.AlignLeft
                    font.family: "Courier"
                    verticalAlignment: Text.AlignVCenter
                }
            }
        }

        Rectangle {
            id: rectangle24
            x: 0
            y: 413
            width: 400
            height: 31
            color: "#00000000"
            Text {
                id: text20
                x: 16
                y: 8
                text: qsTr("Gas Price")
                font.pixelSize: 11
                font.bold: true
                font.family: "Verdana"
            }

            Rectangle {
                id: rectangle25
                x: 90
                y: 0
                width: 302
                height: 30
                color: "#ffffff"
                radius: 4
                Text {
                    id: text21
                    x: 8
                    y: 8
                    width: parent.width - 16
                    height: parent.height - 16
                    color: "#202020"
                    text: ctxObject.gasPrice
                    font.wordSpacing: 0
                    font.letterSpacing: 0
                    styleColor: "#00000000"
                    font.pixelSize: 14
                    horizontalAlignment: Text.AlignLeft
                    font.family: "Courier"
                    verticalAlignment: Text.AlignVCenter
                }
            }
        }

        Rectangle {
            id: rectangle26
            x: 0
            y: 456
            width: 400
            height: 31
            color: "#00000000"
            Text {
                id: text22
                x: 16
                y: 8
                text: qsTr("Nonce")
                font.bold: true
                font.pixelSize: 11
                font.family: "Verdana"
            }

            Rectangle {
                id: rectangle27
                x: 90
                y: 0
                width: 302
                height: 30
                color: "#ffffff"
                radius: 4
                Text {
                    id: text23
                    x: 8
                    y: 8
                    width: parent.width - 16
                    height: parent.height - 16
                    color: "#202020"
                    text: ctxObject.nonce
                    font.letterSpacing: 0
                    font.wordSpacing: 0
                    styleColor: "#00000000"
                    font.pixelSize: 14
                    horizontalAlignment: Text.AlignLeft
                    font.family: "Courier"
                    verticalAlignment: Text.AlignVCenter
                }
            }
        }

    }
}
