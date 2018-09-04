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
            id: rectangle12
            x: 0
            y: 243
            width: 400
            height: 103
            color: "#00000000"
            Text {
                id: text8
                x: 16
                y: 8
                text: qsTr("Hash")
                font.pixelSize: 11
                font.bold: true
                font.family: "Verdana"
            }

            Rectangle {
                id: rectangle13
                x: 8
                y: 28
                width: 384
                height: 75
                color: "#ffffff"
                radius: 4
                Text {
                    id: text9
                    x: 8
                    y: 8
                    width: parent.width - 16
                    height: parent.height - 16
                    color: "#202020"
                    text: ctxObject.hash
                    lineHeight: 1
                    wrapMode: Text.WrapAnywhere
                    fontSizeMode: Text.FixedSize
                    font.capitalization: Font.MixedCase
                    scale: 1
                    rotation: 0
                    font.pixelSize: 14
                    horizontalAlignment: Text.AlignLeft
                    verticalAlignment: Text.AlignTop
                    font.family: "Courier"
                }
            }
        }

        Rectangle {
            id: rectangle14
            x: 0
            y: 352
            width: 400
            height: 141
            color: "#00000000"
            Text {
                id: text10
                x: 16
                y: 8
                text: qsTr("Message")
                font.pixelSize: 11
                font.bold: true
                font.family: "Verdana"
            }

            Rectangle {
                id: rectangle15
                x: 8
                y: 28
                width: 384
                height: 113
                color: "#ffffff"
                radius: 4

                Text {
                    id: text11
                    x: 8
                    y: 8
                    width: parent.width - 16
                    height: parent.height - 16
                    color: "#202020"
                    text: ctxObject.message
                    font.pixelSize: 14
                    scale: 1
                    fontSizeMode: Text.FixedSize
                    horizontalAlignment: Text.AlignLeft
                    font.capitalization: Font.MixedCase
                    lineHeight: 1
                    verticalAlignment: Text.AlignTop
                    wrapMode: Text.WordWrap
                    font.family: "Courier"
                    rotation: 0
                }
            }
        }

        Rectangle {
            id: rectangle16
            x: 0
            y: 501
            width: 400
            height: 82
            color: "#00000000"
            Text {
                id: text12
                x: 16
                y: 8
                text: qsTr("Raw Data")
                font.pixelSize: 11
                font.bold: true
                font.family: "Verdana"
            }

            Rectangle {
                id: rectangle17
                x: 8
                y: 28
                width: 384
                height: 53
                color: "#ffffff"
                radius: 4
                Text {
                    id: text13
                    x: 8
                    y: 8
                    width: parent.width - 16
                    height: parent.height - 16
                    color: "#202020"
                    text: ctxObject.rawData
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
                    x: 0
                    y: 0
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
                    x: 0
                    y: 0
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

    }
}
