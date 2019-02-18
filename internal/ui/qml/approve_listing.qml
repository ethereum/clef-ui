import QtQuick 2.4
import QtQuick.Controls 2.3
//import CustomQmlTypes 1.0		//CustomListModel

Item {
    id: root
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
            id: rect_header
            x: 0
            y: 0
            onBack: {
                ctxObject.back()
            }

        }

        Rectangle {
            id: rectangle16
            x: 0
            y: 70
            width: 400
            height: 549
            color: "#ffffff"

            Text {
                id: text14
                x: 12
                y: 4
                height: 10
                color: "#747474"
                text: qsTr("Accounts")
                font.bold: false
                font.pixelSize: 10
                font.family: "Verdana"
                font.capitalization: Font.AllUppercase
            }

            Rectangle {
                id: rectangle17
                x: 8
                y: 18
                width: 384
                height: 523
                color: "#ffffff"
                radius: 4

                Rectangle {
                    x: 0
                    y: 0
                    radius: 2
                    width: parent.width
                    height: parent.height
                    color: "#efefef"
                }

                ListView {
                    Connections
                    {
                        target: ctxObject
                        onTriggerUpdate: {
                            accounts.callbackFromQml()

                        }
                    }

                    id: listView_accounts
                    x: 0
                    y: 0
                    width: 384
                    height: 523
                    z: 1
                    highlightRangeMode: ListView.NoHighlightRange
                    model: accounts
                    delegate: Item {
                        x: 0
                        width: 384
                        height: 40

                        Row {
                            id: row1
                            CheckBox {
                                display: AbstractButton.IconOnly
                                focusPolicy: Qt.StrongFocus
                                font.pointSize: 12
                                font.capitalization: Font.MixedCase
                                font.family: "Courier"
                                checked: selected
                                onCheckStateChanged: ctxObject.onCheckStateChanged(index, checked)
                            }
                            Text {
                                text: address
                                font.pointSize: 12
                                height: 40
                                verticalAlignment: Text.AlignVCenter
                                horizontalAlignment: Text.AlignLeft
                                font.family: "Courier"
                            }
                        }
                    }
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
    }
}
