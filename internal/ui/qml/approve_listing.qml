import QtQuick 2.4
import QtQuick.Controls 2.3
import CustomQmlTypes 1.0		//CustomListModel

Item {
    id: item1
    width: 400
    height: 615

    Rectangle {
        id: rectangle
        x: 0
        y: 0
        width: 400
        height: 615
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
            height: 348
            color: "#00000000"

            Text {
                id: text6
                x: 16
                y: 8
                text: qsTr("Accounts")
                font.bold: true
                font.family: "Verdana"
                font.pixelSize: 11
            }

            Rectangle {
                id: rectangle9
                x: 8
                y: 28
                width: 384
                height: 312
                color: "#ffffff"
                radius: 4

                ListView {
                    id: listView
                    x: 8
                    y: 8
                    width: 368
                    height: 296
                    z: 1
                    highlightRangeMode: ListView.NoHighlightRange
                    model: CustomListModel{}
                    delegate: Item {
                        x: 5
                        width: 80
                        height: 40
                        Row {
                            id: row1
                            CheckBox {
                                text: Address
                                font.pointSize: 12
                                font.capitalization: Font.MixedCase
                                font.family: "Courier"
//                                checked: listView.model.modelData.checked
                            }
                            spacing: 10
                        }
                    }
                }
            }
        }

        Rectangle {
            id: row
            x: 0
            y: 529
            width: 400
            height: 85
            color: "#ffffff"

            Button {
                id: control
                x: 8
                y: 8
                height: 69
                width: 188
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
                    x: 0
                    y: 0
                    width: parent.width
                    height: parent.height
                    border.width: 0
                    radius: 4
                }

            }

            Button {
                id: control1
                x: 204
                y: 8
                height: 69
                width: 188
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
                    x: 0
                    y: 0
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
