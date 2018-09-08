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
                    onClicked: ctxObject.back()
                    cursorShape: Qt.PointingHandCursor
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
                    onClicked: ctxObject.clicked(1)
                    cursorShape: Qt.PointingHandCursor
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
                    onClicked: ctxObject.clicked(2)
                    cursorShape: Qt.PointingHandCursor
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
            id: rectangle7
            x: 56
            y: 109
            width: 288
            height: 121
            color: "#ffffff"
            radius: 4
            border.width: 0
            border.color: "#c9c9c9"

            Rectangle {
                x: -37
                y: 240
                width: 363
                height: 24
                color: "#efefef"
                radius: 2

                Text {
                    id: text10
                    x: 0
                    y: 0
                    width: 363
                    height: 24
                    color: "#747474"
                    text: ctxObject.address
                    font.pixelSize: 14
                    font.letterSpacing: 0
                    font.bold: false
                    font.capitalization: Font.MixedCase
                    font.family: "Courier"
                    verticalAlignment: Text.AlignVCenter
                    wrapMode: Text.WordWrap
                    horizontalAlignment: Text.AlignHCenter
                }
            }

            Text {
                id: text9
                x: 8
                y: 15
                width: 272
                height: 45
                color: "#747474"
                text: qsTr("Clef is exporting the following account:")
                verticalAlignment: Text.AlignVCenter
                wrapMode: Text.WordWrap
                font.pixelSize: 12
                font.bold: false
                horizontalAlignment: Text.AlignHCenter
                font.family: "Verdana"
                font.capitalization: Font.AllUppercase
            }

        }

        Image {
            id: image
            x: 133
            y: 206
            width: 134
            height: 131
            source: ctxObject.fromSrc
//            source: "data:image/png;base64,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J2l0OUMg45SytWfQ7E5BdshulxF+r0SgrLAyF/MARnbIzpM0upwhkHFK2dozaHanIDtkt8sIv1ciUFZYmYt5ACM7ZOdJGl3OEMg4pWztGTS7U5AdsttlhN8rESgrrMzFPIB9svO8hy4QgAAEPgkgu08s/BECEJhGANlNmyjvgQAEPgkgu08s/BECEJhGANlNmyjvgQAEPgkgu08s/BECEJhGANlNmyjvgQAEPgkgu08s/BECEJhGANlNmyjvgQAEPgkgu08s/BECEJhGANlNmyjvgQAEPgkgu08s/BECEJhGANlNmyjvgQAEPgkgu08s/BECEJhGANlNmyjvgQAEPgkgu08s/BECEJhGANlNmyjvgQAEPgkgu08s/BECEJhGANlNmyjvgQAEPgkgu08s/BECEJhGANlNmyjvgQAEPgkgu08s/BECEJhGANlNmyjvgQAEPgkgu08s/BECEJhGANlNmyjvgQAEPgkgu08s/BECEJhGANlNmyjvgQAEPgkgu08s/BECEJhGANlNmyjvgQAEPgkgu08s/BECEJhGANlNmyjvgQAEPgkgu08s/BECEJhGANlNmyjvgQAEPgkgu08s/BECEJhG4F8hykI7z95C0gAAAABJRU5ErkJggg=="
        }


    }
}
