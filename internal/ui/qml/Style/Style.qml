pragma Singleton	
import QtQuick 2.3	

// can be splitted into multiple files if growing too big...	
QtObject {	
    id: root	

    readonly property int defaultMargin: 16	
    readonly property int defaultRadius: 2	

    readonly property QtObject buttons: QtObject {	
        readonly property color backgroundColor: "#5b5b5b"	
        readonly property color validationColor: "#48b877"	
        readonly property color textColor: "white"	
        readonly property font font: Qt.font({	
            family: "Arial",	
            bold: true,	
            capitalization: Font.AllUppercase,	
            pixelSize: 14	
        })	
        readonly property bool hoverEnabled: true	
        readonly property int cursorShape: Qt.PointingHandCursor	
    }	

    readonly property QtObject separators: QtObject {	
        readonly property color color: "#efefef"	
        readonly property int thickness: 1	
    }	

    readonly property QtObject texts: QtObject {	
        readonly property color backgroundColor: "#efefef"	
        readonly property int backgroundRadius: root.defaultRadius	
        readonly property color titleColor: "#646464"	
        readonly property font titleFont: Qt.font({	
            capitalization: Font.AllUppercase,	
            pixelSize: 12	
        })	
        readonly property font font: Qt.font({	
            family: "Courier",	
            pixelSize: 14	
        })	
        readonly property color color: "#646464"	
    }	

    readonly property QtObject textFields: QtObject {	
        readonly property color backgroundColor: texts.backgroundColor	
        readonly property int backgroundRadius: root.defaultRadius	
        readonly property color titleColor: texts.titleColor	
        readonly property font titleFont: texts.titleFont	
        readonly property font font: Qt.font({	
            family: "Verdana",	
            pixelSize: 12	
        })	
    }	

    readonly property QtObject disclaimers: QtObject {	
        readonly property color borderColor: "#c9c9c9"	
        readonly property int borderRadius: 4	
        readonly property int borderWidth: 1	
        readonly property color textColor: "#747474"	
        readonly property font font: textFields.font	
    }	

    function rgb(r, g, b) {	
        return Qt.rgba(r / 255, g / 255, b / 255, 1)	
    }	

    function rgba(r, g, b, a) {	
        return Qt.rgba(r / 255, g / 255, b / 255, a)	
    }	
}
