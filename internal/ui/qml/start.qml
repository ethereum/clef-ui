import QtQuick 2.7
import QtQuick.Controls 2.1
import "Style"
import "Controls" as ClefUI
import "Views" as ClefUI

Item {
    id: window

    ClefUI.Login {
        id: login
        anchors.fill: parent
        context: LoginContext
    }

    Component.onCompleted: {
//        LoginContext.defaultGoPath = "test"
        login.forceActiveFocus()
    }
}
