import QtQuick 2.7	
import QtQuick.Controls 2.1	
import "../Style"

Item {	
    id: root	

    implicitHeight: 121	
    implicitWidth: 288	

    Rectangle {	
        anchors.fill: parent	
        radius: Style.disclaimers.borderRadius	
        border.color: Style.disclaimers.borderColor	
        border.width: Style.disclaimers.borderWidth	
    }	

    Text {	
        anchors.fill: parent	
        anchors.margins: Style.defaultMargin / 2	
        font: Style.disclaimers.font	
        color: Style.disclaimers.textColor	
        wrapMode: Text.WordWrap	
        horizontalAlignment: Text.AlignHCenter	
        verticalAlignment: Text.AlignVCenter	
        text: qsTr("Clef will generate a new private key, encrypts it according to web3 keystore spec and stores it in the keystore directory. Please create a backup of the keystore. If the keystore is lost there is no method of retrieving lost accounts.")	

    }	
}