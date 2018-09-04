

#define protected public
#define private public

#include "moc.h"
#include "_cgo_export.h"

#include <QAction>
#include <QActionEvent>
#include <QByteArray>
#include <QCamera>
#include <QCameraImageCapture>
#include <QChildEvent>
#include <QCloseEvent>
#include <QContextMenuEvent>
#include <QDBusPendingCall>
#include <QDBusPendingCallWatcher>
#include <QDrag>
#include <QDragEnterEvent>
#include <QDragLeaveEvent>
#include <QDragMoveEvent>
#include <QDropEvent>
#include <QEvent>
#include <QExtensionFactory>
#include <QExtensionManager>
#include <QFocusEvent>
#include <QGraphicsObject>
#include <QGraphicsWidget>
#include <QHideEvent>
#include <QIcon>
#include <QInputMethod>
#include <QInputMethodEvent>
#include <QKeyEvent>
#include <QLabel>
#include <QLayout>
#include <QMediaPlaylist>
#include <QMediaRecorder>
#include <QMetaMethod>
#include <QMouseEvent>
#include <QMoveEvent>
#include <QMovie>
#include <QObject>
#include <QOffscreenSurface>
#include <QPaintDevice>
#include <QPaintDeviceWindow>
#include <QPaintEngine>
#include <QPaintEvent>
#include <QPainter>
#include <QPdfWriter>
#include <QPicture>
#include <QPixmap>
#include <QPoint>
#include <QQuickItem>
#include <QRadioData>
#include <QResizeEvent>
#include <QShowEvent>
#include <QSize>
#include <QString>
#include <QTabletEvent>
#include <QTime>
#include <QTimer>
#include <QTimerEvent>
#include <QVariant>
#include <QWheelEvent>
#include <QWidget>
#include <QWindow>


class CtxObject721036: public QObject
{
Q_OBJECT
Q_PROPERTY(QString remote READ remote WRITE setRemote NOTIFY remoteChanged)
Q_PROPERTY(QString transport READ transport WRITE setTransport NOTIFY transportChanged)
Q_PROPERTY(QString endpoint READ endpoint WRITE setEndpoint NOTIFY endpointChanged)
Q_PROPERTY(QString from READ from WRITE setFrom NOTIFY fromChanged)
Q_PROPERTY(QString message READ message WRITE setMessage NOTIFY messageChanged)
Q_PROPERTY(QString rawData READ rawData WRITE setRawData NOTIFY rawDataChanged)
Q_PROPERTY(QString hash READ hash WRITE setHash NOTIFY hashChanged)
public:
	CtxObject721036(QObject *parent = Q_NULLPTR) : QObject(parent) {qRegisterMetaType<quintptr>("quintptr");CtxObject721036_CtxObject721036_QRegisterMetaType();CtxObject721036_CtxObject721036_QRegisterMetaTypes();callbackCtxObject721036_Constructor(this);};
	void Signal_Clicked(qint32 b) { callbackCtxObject721036_Clicked(this, b); };
	QString remote() { return ({ Moc_PackedString tempVal = callbackCtxObject721036_Remote(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setRemote(QString remote) { QByteArray t41ffe5 = remote.toUtf8(); Moc_PackedString remotePacked = { const_cast<char*>(t41ffe5.prepend("WHITESPACE").constData()+10), t41ffe5.size()-10 };callbackCtxObject721036_SetRemote(this, remotePacked); };
	void Signal_RemoteChanged(QString remote) { QByteArray t41ffe5 = remote.toUtf8(); Moc_PackedString remotePacked = { const_cast<char*>(t41ffe5.prepend("WHITESPACE").constData()+10), t41ffe5.size()-10 };callbackCtxObject721036_RemoteChanged(this, remotePacked); };
	QString transport() { return ({ Moc_PackedString tempVal = callbackCtxObject721036_Transport(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setTransport(QString transport) { QByteArray ta8e601 = transport.toUtf8(); Moc_PackedString transportPacked = { const_cast<char*>(ta8e601.prepend("WHITESPACE").constData()+10), ta8e601.size()-10 };callbackCtxObject721036_SetTransport(this, transportPacked); };
	void Signal_TransportChanged(QString transport) { QByteArray ta8e601 = transport.toUtf8(); Moc_PackedString transportPacked = { const_cast<char*>(ta8e601.prepend("WHITESPACE").constData()+10), ta8e601.size()-10 };callbackCtxObject721036_TransportChanged(this, transportPacked); };
	QString endpoint() { return ({ Moc_PackedString tempVal = callbackCtxObject721036_Endpoint(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setEndpoint(QString endpoint) { QByteArray te13fe4 = endpoint.toUtf8(); Moc_PackedString endpointPacked = { const_cast<char*>(te13fe4.prepend("WHITESPACE").constData()+10), te13fe4.size()-10 };callbackCtxObject721036_SetEndpoint(this, endpointPacked); };
	void Signal_EndpointChanged(QString endpoint) { QByteArray te13fe4 = endpoint.toUtf8(); Moc_PackedString endpointPacked = { const_cast<char*>(te13fe4.prepend("WHITESPACE").constData()+10), te13fe4.size()-10 };callbackCtxObject721036_EndpointChanged(this, endpointPacked); };
	QString from() { return ({ Moc_PackedString tempVal = callbackCtxObject721036_From(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setFrom(QString from) { QByteArray t0b1e95 = from.toUtf8(); Moc_PackedString fromPacked = { const_cast<char*>(t0b1e95.prepend("WHITESPACE").constData()+10), t0b1e95.size()-10 };callbackCtxObject721036_SetFrom(this, fromPacked); };
	void Signal_FromChanged(QString from) { QByteArray t0b1e95 = from.toUtf8(); Moc_PackedString fromPacked = { const_cast<char*>(t0b1e95.prepend("WHITESPACE").constData()+10), t0b1e95.size()-10 };callbackCtxObject721036_FromChanged(this, fromPacked); };
	QString message() { return ({ Moc_PackedString tempVal = callbackCtxObject721036_Message(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setMessage(QString message) { QByteArray t6f9b9a = message.toUtf8(); Moc_PackedString messagePacked = { const_cast<char*>(t6f9b9a.prepend("WHITESPACE").constData()+10), t6f9b9a.size()-10 };callbackCtxObject721036_SetMessage(this, messagePacked); };
	void Signal_MessageChanged(QString message) { QByteArray t6f9b9a = message.toUtf8(); Moc_PackedString messagePacked = { const_cast<char*>(t6f9b9a.prepend("WHITESPACE").constData()+10), t6f9b9a.size()-10 };callbackCtxObject721036_MessageChanged(this, messagePacked); };
	QString rawData() { return ({ Moc_PackedString tempVal = callbackCtxObject721036_RawData(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setRawData(QString rawData) { QByteArray tcacc10 = rawData.toUtf8(); Moc_PackedString rawDataPacked = { const_cast<char*>(tcacc10.prepend("WHITESPACE").constData()+10), tcacc10.size()-10 };callbackCtxObject721036_SetRawData(this, rawDataPacked); };
	void Signal_RawDataChanged(QString rawData) { QByteArray tcacc10 = rawData.toUtf8(); Moc_PackedString rawDataPacked = { const_cast<char*>(tcacc10.prepend("WHITESPACE").constData()+10), tcacc10.size()-10 };callbackCtxObject721036_RawDataChanged(this, rawDataPacked); };
	QString hash() { return ({ Moc_PackedString tempVal = callbackCtxObject721036_Hash(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setHash(QString hash) { QByteArray t2346ad = hash.toUtf8(); Moc_PackedString hashPacked = { const_cast<char*>(t2346ad.prepend("WHITESPACE").constData()+10), t2346ad.size()-10 };callbackCtxObject721036_SetHash(this, hashPacked); };
	void Signal_HashChanged(QString hash) { QByteArray t2346ad = hash.toUtf8(); Moc_PackedString hashPacked = { const_cast<char*>(t2346ad.prepend("WHITESPACE").constData()+10), t2346ad.size()-10 };callbackCtxObject721036_HashChanged(this, hashPacked); };
	 ~CtxObject721036() { callbackCtxObject721036_DestroyCtxObject(this); };
	bool event(QEvent * e) { return callbackCtxObject721036_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackCtxObject721036_EventFilter(this, watched, event) != 0; };
	
	void childEvent(QChildEvent * event) { callbackCtxObject721036_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackCtxObject721036_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackCtxObject721036_CustomEvent(this, event); };
	void deleteLater() { callbackCtxObject721036_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackCtxObject721036_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackCtxObject721036_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackCtxObject721036_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackCtxObject721036_TimerEvent(this, event); };
	
	QString remoteDefault() { return _remote; };
	void setRemoteDefault(QString p) { if (p != _remote) { _remote = p; remoteChanged(_remote); } };
	QString transportDefault() { return _transport; };
	void setTransportDefault(QString p) { if (p != _transport) { _transport = p; transportChanged(_transport); } };
	QString endpointDefault() { return _endpoint; };
	void setEndpointDefault(QString p) { if (p != _endpoint) { _endpoint = p; endpointChanged(_endpoint); } };
	QString fromDefault() { return _from; };
	void setFromDefault(QString p) { if (p != _from) { _from = p; fromChanged(_from); } };
	QString messageDefault() { return _message; };
	void setMessageDefault(QString p) { if (p != _message) { _message = p; messageChanged(_message); } };
	QString rawDataDefault() { return _rawData; };
	void setRawDataDefault(QString p) { if (p != _rawData) { _rawData = p; rawDataChanged(_rawData); } };
	QString hashDefault() { return _hash; };
	void setHashDefault(QString p) { if (p != _hash) { _hash = p; hashChanged(_hash); } };
signals:
	void clicked(qint32 b);
	void remoteChanged(QString remote);
	void transportChanged(QString transport);
	void endpointChanged(QString endpoint);
	void fromChanged(QString from);
	void messageChanged(QString message);
	void rawDataChanged(QString rawData);
	void hashChanged(QString hash);
public slots:
private:
	QString _remote;
	QString _transport;
	QString _endpoint;
	QString _from;
	QString _message;
	QString _rawData;
	QString _hash;
};

Q_DECLARE_METATYPE(CtxObject721036*)


void CtxObject721036_CtxObject721036_QRegisterMetaTypes() {
	qRegisterMetaType<QString>();
}

class CustomLabel721036: public QLabel
{
Q_OBJECT
public:
	CustomLabel721036(QWidget *parent = Q_NULLPTR, Qt::WindowFlags fo = Qt::WindowFlags()) : QLabel(parent, fo) {qRegisterMetaType<quintptr>("quintptr");CustomLabel721036_CustomLabel721036_QRegisterMetaType();CustomLabel721036_CustomLabel721036_QRegisterMetaTypes();callbackCustomLabel721036_Constructor(this);};
	CustomLabel721036(const QString &text, QWidget *parent = Q_NULLPTR, Qt::WindowFlags fo = Qt::WindowFlags()) : QLabel(text, parent, fo) {qRegisterMetaType<quintptr>("quintptr");CustomLabel721036_CustomLabel721036_QRegisterMetaType();CustomLabel721036_CustomLabel721036_QRegisterMetaTypes();callbackCustomLabel721036_Constructor(this);};
	void Signal_UpdateText(QString v0) { QByteArray tea1dd7 = v0.toUtf8(); Moc_PackedString v0Packed = { const_cast<char*>(tea1dd7.prepend("WHITESPACE").constData()+10), tea1dd7.size()-10 };callbackCustomLabel721036_UpdateText(this, v0Packed); };
	 ~CustomLabel721036() { callbackCustomLabel721036_DestroyCustomLabel(this); };
	bool event(QEvent * e) { return callbackCustomLabel721036_Event(this, e) != 0; };
	bool focusNextPrevChild(bool next) { return callbackCustomLabel721036_FocusNextPrevChild(this, next) != 0; };
	
	void changeEvent(QEvent * ev) { callbackCustomLabel721036_ChangeEvent(this, ev); };
	void clear() { callbackCustomLabel721036_Clear(this); };
	void contextMenuEvent(QContextMenuEvent * ev) { callbackCustomLabel721036_ContextMenuEvent(this, ev); };
	void focusInEvent(QFocusEvent * ev) { callbackCustomLabel721036_FocusInEvent(this, ev); };
	void focusOutEvent(QFocusEvent * ev) { callbackCustomLabel721036_FocusOutEvent(this, ev); };
	void keyPressEvent(QKeyEvent * ev) { callbackCustomLabel721036_KeyPressEvent(this, ev); };
	void Signal_LinkActivated(const QString & link) { QByteArray t4f0aa5 = link.toUtf8(); Moc_PackedString linkPacked = { const_cast<char*>(t4f0aa5.prepend("WHITESPACE").constData()+10), t4f0aa5.size()-10 };callbackCustomLabel721036_LinkActivated(this, linkPacked); };
	void Signal_LinkHovered(const QString & link) { QByteArray t4f0aa5 = link.toUtf8(); Moc_PackedString linkPacked = { const_cast<char*>(t4f0aa5.prepend("WHITESPACE").constData()+10), t4f0aa5.size()-10 };callbackCustomLabel721036_LinkHovered(this, linkPacked); };
	void mouseMoveEvent(QMouseEvent * ev) { callbackCustomLabel721036_MouseMoveEvent(this, ev); };
	void mousePressEvent(QMouseEvent * ev) { callbackCustomLabel721036_MousePressEvent(this, ev); };
	void mouseReleaseEvent(QMouseEvent * ev) { callbackCustomLabel721036_MouseReleaseEvent(this, ev); };
	void paintEvent(QPaintEvent * vqp) { callbackCustomLabel721036_PaintEvent(this, vqp); };
	void setMovie(QMovie * movie) { callbackCustomLabel721036_SetMovie(this, movie); };
	void setNum(double num) { callbackCustomLabel721036_SetNum2(this, num); };
	void setNum(int num) { callbackCustomLabel721036_SetNum(this, num); };
	void setPicture(const QPicture & picture) { callbackCustomLabel721036_SetPicture(this, const_cast<QPicture*>(&picture)); };
	void setPixmap(const QPixmap & vqp) { callbackCustomLabel721036_SetPixmap(this, const_cast<QPixmap*>(&vqp)); };
	void setText(const QString & vqs) { QByteArray tda39a3 = vqs.toUtf8(); Moc_PackedString vqsPacked = { const_cast<char*>(tda39a3.prepend("WHITESPACE").constData()+10), tda39a3.size()-10 };callbackCustomLabel721036_SetText(this, vqsPacked); };
	QSize minimumSizeHint() const { return *static_cast<QSize*>(callbackCustomLabel721036_MinimumSizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	QSize sizeHint() const { return *static_cast<QSize*>(callbackCustomLabel721036_SizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	
	int heightForWidth(int w) const { return callbackCustomLabel721036_HeightForWidth(const_cast<void*>(static_cast<const void*>(this)), w); };
	bool close() { return callbackCustomLabel721036_Close(this) != 0; };
	void actionEvent(QActionEvent * event) { callbackCustomLabel721036_ActionEvent(this, event); };
	void closeEvent(QCloseEvent * event) { callbackCustomLabel721036_CloseEvent(this, event); };
	void Signal_CustomContextMenuRequested(const QPoint & pos) { callbackCustomLabel721036_CustomContextMenuRequested(this, const_cast<QPoint*>(&pos)); };
	void dragEnterEvent(QDragEnterEvent * event) { callbackCustomLabel721036_DragEnterEvent(this, event); };
	void dragLeaveEvent(QDragLeaveEvent * event) { callbackCustomLabel721036_DragLeaveEvent(this, event); };
	void dragMoveEvent(QDragMoveEvent * event) { callbackCustomLabel721036_DragMoveEvent(this, event); };
	void dropEvent(QDropEvent * event) { callbackCustomLabel721036_DropEvent(this, event); };
	void enterEvent(QEvent * event) { callbackCustomLabel721036_EnterEvent(this, event); };
	void hide() { callbackCustomLabel721036_Hide(this); };
	void hideEvent(QHideEvent * event) { callbackCustomLabel721036_HideEvent(this, event); };
	void inputMethodEvent(QInputMethodEvent * event) { callbackCustomLabel721036_InputMethodEvent(this, event); };
	void keyReleaseEvent(QKeyEvent * event) { callbackCustomLabel721036_KeyReleaseEvent(this, event); };
	void leaveEvent(QEvent * event) { callbackCustomLabel721036_LeaveEvent(this, event); };
	void lower() { callbackCustomLabel721036_Lower(this); };
	void mouseDoubleClickEvent(QMouseEvent * event) { callbackCustomLabel721036_MouseDoubleClickEvent(this, event); };
	void moveEvent(QMoveEvent * event) { callbackCustomLabel721036_MoveEvent(this, event); };
	void raise() { callbackCustomLabel721036_Raise(this); };
	void repaint() { callbackCustomLabel721036_Repaint(this); };
	void resizeEvent(QResizeEvent * event) { callbackCustomLabel721036_ResizeEvent(this, event); };
	void setDisabled(bool disable) { callbackCustomLabel721036_SetDisabled(this, disable); };
	void setEnabled(bool vbo) { callbackCustomLabel721036_SetEnabled(this, vbo); };
	void setFocus() { callbackCustomLabel721036_SetFocus2(this); };
	void setHidden(bool hidden) { callbackCustomLabel721036_SetHidden(this, hidden); };
	void setStyleSheet(const QString & styleSheet) { QByteArray t728ae7 = styleSheet.toUtf8(); Moc_PackedString styleSheetPacked = { const_cast<char*>(t728ae7.prepend("WHITESPACE").constData()+10), t728ae7.size()-10 };callbackCustomLabel721036_SetStyleSheet(this, styleSheetPacked); };
	void setVisible(bool visible) { callbackCustomLabel721036_SetVisible(this, visible); };
	void setWindowModified(bool vbo) { callbackCustomLabel721036_SetWindowModified(this, vbo); };
	void setWindowTitle(const QString & vqs) { QByteArray tda39a3 = vqs.toUtf8(); Moc_PackedString vqsPacked = { const_cast<char*>(tda39a3.prepend("WHITESPACE").constData()+10), tda39a3.size()-10 };callbackCustomLabel721036_SetWindowTitle(this, vqsPacked); };
	void show() { callbackCustomLabel721036_Show(this); };
	void showEvent(QShowEvent * event) { callbackCustomLabel721036_ShowEvent(this, event); };
	void showFullScreen() { callbackCustomLabel721036_ShowFullScreen(this); };
	void showMaximized() { callbackCustomLabel721036_ShowMaximized(this); };
	void showMinimized() { callbackCustomLabel721036_ShowMinimized(this); };
	void showNormal() { callbackCustomLabel721036_ShowNormal(this); };
	void tabletEvent(QTabletEvent * event) { callbackCustomLabel721036_TabletEvent(this, event); };
	void update() { callbackCustomLabel721036_Update(this); };
	void updateMicroFocus() { callbackCustomLabel721036_UpdateMicroFocus(this); };
	void wheelEvent(QWheelEvent * event) { callbackCustomLabel721036_WheelEvent(this, event); };
	void Signal_WindowIconChanged(const QIcon & icon) { callbackCustomLabel721036_WindowIconChanged(this, const_cast<QIcon*>(&icon)); };
	void Signal_WindowTitleChanged(const QString & title) { QByteArray t3c6de1 = title.toUtf8(); Moc_PackedString titlePacked = { const_cast<char*>(t3c6de1.prepend("WHITESPACE").constData()+10), t3c6de1.size()-10 };callbackCustomLabel721036_WindowTitleChanged(this, titlePacked); };
	QPaintEngine * paintEngine() const { return static_cast<QPaintEngine*>(callbackCustomLabel721036_PaintEngine(const_cast<void*>(static_cast<const void*>(this)))); };
	QVariant inputMethodQuery(Qt::InputMethodQuery query) const { return *static_cast<QVariant*>(callbackCustomLabel721036_InputMethodQuery(const_cast<void*>(static_cast<const void*>(this)), query)); };
	bool hasHeightForWidth() const { return callbackCustomLabel721036_HasHeightForWidth(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	int metric(QPaintDevice::PaintDeviceMetric m) const { return callbackCustomLabel721036_Metric(const_cast<void*>(static_cast<const void*>(this)), m); };
	void initPainter(QPainter * painter) const { callbackCustomLabel721036_InitPainter(const_cast<void*>(static_cast<const void*>(this)), painter); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackCustomLabel721036_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackCustomLabel721036_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackCustomLabel721036_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackCustomLabel721036_CustomEvent(this, event); };
	void deleteLater() { callbackCustomLabel721036_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackCustomLabel721036_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackCustomLabel721036_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackCustomLabel721036_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackCustomLabel721036_TimerEvent(this, event); };
signals:
	void updateText(QString v0);
public slots:
private:
};

Q_DECLARE_METATYPE(CustomLabel721036*)


void CustomLabel721036_CustomLabel721036_QRegisterMetaTypes() {
}

void CtxObject721036_ConnectClicked(void* ptr)
{
	QObject::connect(static_cast<CtxObject721036*>(ptr), static_cast<void (CtxObject721036::*)(qint32)>(&CtxObject721036::clicked), static_cast<CtxObject721036*>(ptr), static_cast<void (CtxObject721036::*)(qint32)>(&CtxObject721036::Signal_Clicked));
}

void CtxObject721036_DisconnectClicked(void* ptr)
{
	QObject::disconnect(static_cast<CtxObject721036*>(ptr), static_cast<void (CtxObject721036::*)(qint32)>(&CtxObject721036::clicked), static_cast<CtxObject721036*>(ptr), static_cast<void (CtxObject721036::*)(qint32)>(&CtxObject721036::Signal_Clicked));
}

void CtxObject721036_Clicked(void* ptr, int b)
{
	static_cast<CtxObject721036*>(ptr)->clicked(b);
}

struct Moc_PackedString CtxObject721036_Remote(void* ptr)
{
	return ({ QByteArray t7c9bd6 = static_cast<CtxObject721036*>(ptr)->remote().toUtf8(); Moc_PackedString { const_cast<char*>(t7c9bd6.prepend("WHITESPACE").constData()+10), t7c9bd6.size()-10 }; });
}

struct Moc_PackedString CtxObject721036_RemoteDefault(void* ptr)
{
	return ({ QByteArray t014c6a = static_cast<CtxObject721036*>(ptr)->remoteDefault().toUtf8(); Moc_PackedString { const_cast<char*>(t014c6a.prepend("WHITESPACE").constData()+10), t014c6a.size()-10 }; });
}

void CtxObject721036_SetRemote(void* ptr, struct Moc_PackedString remote)
{
	static_cast<CtxObject721036*>(ptr)->setRemote(QString::fromUtf8(remote.data, remote.len));
}

void CtxObject721036_SetRemoteDefault(void* ptr, struct Moc_PackedString remote)
{
	static_cast<CtxObject721036*>(ptr)->setRemoteDefault(QString::fromUtf8(remote.data, remote.len));
}

void CtxObject721036_ConnectRemoteChanged(void* ptr)
{
	QObject::connect(static_cast<CtxObject721036*>(ptr), static_cast<void (CtxObject721036::*)(QString)>(&CtxObject721036::remoteChanged), static_cast<CtxObject721036*>(ptr), static_cast<void (CtxObject721036::*)(QString)>(&CtxObject721036::Signal_RemoteChanged));
}

void CtxObject721036_DisconnectRemoteChanged(void* ptr)
{
	QObject::disconnect(static_cast<CtxObject721036*>(ptr), static_cast<void (CtxObject721036::*)(QString)>(&CtxObject721036::remoteChanged), static_cast<CtxObject721036*>(ptr), static_cast<void (CtxObject721036::*)(QString)>(&CtxObject721036::Signal_RemoteChanged));
}

void CtxObject721036_RemoteChanged(void* ptr, struct Moc_PackedString remote)
{
	static_cast<CtxObject721036*>(ptr)->remoteChanged(QString::fromUtf8(remote.data, remote.len));
}

struct Moc_PackedString CtxObject721036_Transport(void* ptr)
{
	return ({ QByteArray tb4e796 = static_cast<CtxObject721036*>(ptr)->transport().toUtf8(); Moc_PackedString { const_cast<char*>(tb4e796.prepend("WHITESPACE").constData()+10), tb4e796.size()-10 }; });
}

struct Moc_PackedString CtxObject721036_TransportDefault(void* ptr)
{
	return ({ QByteArray t957ae9 = static_cast<CtxObject721036*>(ptr)->transportDefault().toUtf8(); Moc_PackedString { const_cast<char*>(t957ae9.prepend("WHITESPACE").constData()+10), t957ae9.size()-10 }; });
}

void CtxObject721036_SetTransport(void* ptr, struct Moc_PackedString transport)
{
	static_cast<CtxObject721036*>(ptr)->setTransport(QString::fromUtf8(transport.data, transport.len));
}

void CtxObject721036_SetTransportDefault(void* ptr, struct Moc_PackedString transport)
{
	static_cast<CtxObject721036*>(ptr)->setTransportDefault(QString::fromUtf8(transport.data, transport.len));
}

void CtxObject721036_ConnectTransportChanged(void* ptr)
{
	QObject::connect(static_cast<CtxObject721036*>(ptr), static_cast<void (CtxObject721036::*)(QString)>(&CtxObject721036::transportChanged), static_cast<CtxObject721036*>(ptr), static_cast<void (CtxObject721036::*)(QString)>(&CtxObject721036::Signal_TransportChanged));
}

void CtxObject721036_DisconnectTransportChanged(void* ptr)
{
	QObject::disconnect(static_cast<CtxObject721036*>(ptr), static_cast<void (CtxObject721036::*)(QString)>(&CtxObject721036::transportChanged), static_cast<CtxObject721036*>(ptr), static_cast<void (CtxObject721036::*)(QString)>(&CtxObject721036::Signal_TransportChanged));
}

void CtxObject721036_TransportChanged(void* ptr, struct Moc_PackedString transport)
{
	static_cast<CtxObject721036*>(ptr)->transportChanged(QString::fromUtf8(transport.data, transport.len));
}

struct Moc_PackedString CtxObject721036_Endpoint(void* ptr)
{
	return ({ QByteArray t1856a9 = static_cast<CtxObject721036*>(ptr)->endpoint().toUtf8(); Moc_PackedString { const_cast<char*>(t1856a9.prepend("WHITESPACE").constData()+10), t1856a9.size()-10 }; });
}

struct Moc_PackedString CtxObject721036_EndpointDefault(void* ptr)
{
	return ({ QByteArray tcc850f = static_cast<CtxObject721036*>(ptr)->endpointDefault().toUtf8(); Moc_PackedString { const_cast<char*>(tcc850f.prepend("WHITESPACE").constData()+10), tcc850f.size()-10 }; });
}

void CtxObject721036_SetEndpoint(void* ptr, struct Moc_PackedString endpoint)
{
	static_cast<CtxObject721036*>(ptr)->setEndpoint(QString::fromUtf8(endpoint.data, endpoint.len));
}

void CtxObject721036_SetEndpointDefault(void* ptr, struct Moc_PackedString endpoint)
{
	static_cast<CtxObject721036*>(ptr)->setEndpointDefault(QString::fromUtf8(endpoint.data, endpoint.len));
}

void CtxObject721036_ConnectEndpointChanged(void* ptr)
{
	QObject::connect(static_cast<CtxObject721036*>(ptr), static_cast<void (CtxObject721036::*)(QString)>(&CtxObject721036::endpointChanged), static_cast<CtxObject721036*>(ptr), static_cast<void (CtxObject721036::*)(QString)>(&CtxObject721036::Signal_EndpointChanged));
}

void CtxObject721036_DisconnectEndpointChanged(void* ptr)
{
	QObject::disconnect(static_cast<CtxObject721036*>(ptr), static_cast<void (CtxObject721036::*)(QString)>(&CtxObject721036::endpointChanged), static_cast<CtxObject721036*>(ptr), static_cast<void (CtxObject721036::*)(QString)>(&CtxObject721036::Signal_EndpointChanged));
}

void CtxObject721036_EndpointChanged(void* ptr, struct Moc_PackedString endpoint)
{
	static_cast<CtxObject721036*>(ptr)->endpointChanged(QString::fromUtf8(endpoint.data, endpoint.len));
}

struct Moc_PackedString CtxObject721036_From(void* ptr)
{
	return ({ QByteArray tacf4a6 = static_cast<CtxObject721036*>(ptr)->from().toUtf8(); Moc_PackedString { const_cast<char*>(tacf4a6.prepend("WHITESPACE").constData()+10), tacf4a6.size()-10 }; });
}

struct Moc_PackedString CtxObject721036_FromDefault(void* ptr)
{
	return ({ QByteArray t174a18 = static_cast<CtxObject721036*>(ptr)->fromDefault().toUtf8(); Moc_PackedString { const_cast<char*>(t174a18.prepend("WHITESPACE").constData()+10), t174a18.size()-10 }; });
}

void CtxObject721036_SetFrom(void* ptr, struct Moc_PackedString from)
{
	static_cast<CtxObject721036*>(ptr)->setFrom(QString::fromUtf8(from.data, from.len));
}

void CtxObject721036_SetFromDefault(void* ptr, struct Moc_PackedString from)
{
	static_cast<CtxObject721036*>(ptr)->setFromDefault(QString::fromUtf8(from.data, from.len));
}

void CtxObject721036_ConnectFromChanged(void* ptr)
{
	QObject::connect(static_cast<CtxObject721036*>(ptr), static_cast<void (CtxObject721036::*)(QString)>(&CtxObject721036::fromChanged), static_cast<CtxObject721036*>(ptr), static_cast<void (CtxObject721036::*)(QString)>(&CtxObject721036::Signal_FromChanged));
}

void CtxObject721036_DisconnectFromChanged(void* ptr)
{
	QObject::disconnect(static_cast<CtxObject721036*>(ptr), static_cast<void (CtxObject721036::*)(QString)>(&CtxObject721036::fromChanged), static_cast<CtxObject721036*>(ptr), static_cast<void (CtxObject721036::*)(QString)>(&CtxObject721036::Signal_FromChanged));
}

void CtxObject721036_FromChanged(void* ptr, struct Moc_PackedString from)
{
	static_cast<CtxObject721036*>(ptr)->fromChanged(QString::fromUtf8(from.data, from.len));
}

struct Moc_PackedString CtxObject721036_Message(void* ptr)
{
	return ({ QByteArray t1245a9 = static_cast<CtxObject721036*>(ptr)->message().toUtf8(); Moc_PackedString { const_cast<char*>(t1245a9.prepend("WHITESPACE").constData()+10), t1245a9.size()-10 }; });
}

struct Moc_PackedString CtxObject721036_MessageDefault(void* ptr)
{
	return ({ QByteArray t422ffb = static_cast<CtxObject721036*>(ptr)->messageDefault().toUtf8(); Moc_PackedString { const_cast<char*>(t422ffb.prepend("WHITESPACE").constData()+10), t422ffb.size()-10 }; });
}

void CtxObject721036_SetMessage(void* ptr, struct Moc_PackedString message)
{
	static_cast<CtxObject721036*>(ptr)->setMessage(QString::fromUtf8(message.data, message.len));
}

void CtxObject721036_SetMessageDefault(void* ptr, struct Moc_PackedString message)
{
	static_cast<CtxObject721036*>(ptr)->setMessageDefault(QString::fromUtf8(message.data, message.len));
}

void CtxObject721036_ConnectMessageChanged(void* ptr)
{
	QObject::connect(static_cast<CtxObject721036*>(ptr), static_cast<void (CtxObject721036::*)(QString)>(&CtxObject721036::messageChanged), static_cast<CtxObject721036*>(ptr), static_cast<void (CtxObject721036::*)(QString)>(&CtxObject721036::Signal_MessageChanged));
}

void CtxObject721036_DisconnectMessageChanged(void* ptr)
{
	QObject::disconnect(static_cast<CtxObject721036*>(ptr), static_cast<void (CtxObject721036::*)(QString)>(&CtxObject721036::messageChanged), static_cast<CtxObject721036*>(ptr), static_cast<void (CtxObject721036::*)(QString)>(&CtxObject721036::Signal_MessageChanged));
}

void CtxObject721036_MessageChanged(void* ptr, struct Moc_PackedString message)
{
	static_cast<CtxObject721036*>(ptr)->messageChanged(QString::fromUtf8(message.data, message.len));
}

struct Moc_PackedString CtxObject721036_RawData(void* ptr)
{
	return ({ QByteArray tf70d6c = static_cast<CtxObject721036*>(ptr)->rawData().toUtf8(); Moc_PackedString { const_cast<char*>(tf70d6c.prepend("WHITESPACE").constData()+10), tf70d6c.size()-10 }; });
}

struct Moc_PackedString CtxObject721036_RawDataDefault(void* ptr)
{
	return ({ QByteArray t9270de = static_cast<CtxObject721036*>(ptr)->rawDataDefault().toUtf8(); Moc_PackedString { const_cast<char*>(t9270de.prepend("WHITESPACE").constData()+10), t9270de.size()-10 }; });
}

void CtxObject721036_SetRawData(void* ptr, struct Moc_PackedString rawData)
{
	static_cast<CtxObject721036*>(ptr)->setRawData(QString::fromUtf8(rawData.data, rawData.len));
}

void CtxObject721036_SetRawDataDefault(void* ptr, struct Moc_PackedString rawData)
{
	static_cast<CtxObject721036*>(ptr)->setRawDataDefault(QString::fromUtf8(rawData.data, rawData.len));
}

void CtxObject721036_ConnectRawDataChanged(void* ptr)
{
	QObject::connect(static_cast<CtxObject721036*>(ptr), static_cast<void (CtxObject721036::*)(QString)>(&CtxObject721036::rawDataChanged), static_cast<CtxObject721036*>(ptr), static_cast<void (CtxObject721036::*)(QString)>(&CtxObject721036::Signal_RawDataChanged));
}

void CtxObject721036_DisconnectRawDataChanged(void* ptr)
{
	QObject::disconnect(static_cast<CtxObject721036*>(ptr), static_cast<void (CtxObject721036::*)(QString)>(&CtxObject721036::rawDataChanged), static_cast<CtxObject721036*>(ptr), static_cast<void (CtxObject721036::*)(QString)>(&CtxObject721036::Signal_RawDataChanged));
}

void CtxObject721036_RawDataChanged(void* ptr, struct Moc_PackedString rawData)
{
	static_cast<CtxObject721036*>(ptr)->rawDataChanged(QString::fromUtf8(rawData.data, rawData.len));
}

struct Moc_PackedString CtxObject721036_Hash(void* ptr)
{
	return ({ QByteArray tea5ef5 = static_cast<CtxObject721036*>(ptr)->hash().toUtf8(); Moc_PackedString { const_cast<char*>(tea5ef5.prepend("WHITESPACE").constData()+10), tea5ef5.size()-10 }; });
}

struct Moc_PackedString CtxObject721036_HashDefault(void* ptr)
{
	return ({ QByteArray t6a9fe6 = static_cast<CtxObject721036*>(ptr)->hashDefault().toUtf8(); Moc_PackedString { const_cast<char*>(t6a9fe6.prepend("WHITESPACE").constData()+10), t6a9fe6.size()-10 }; });
}

void CtxObject721036_SetHash(void* ptr, struct Moc_PackedString hash)
{
	static_cast<CtxObject721036*>(ptr)->setHash(QString::fromUtf8(hash.data, hash.len));
}

void CtxObject721036_SetHashDefault(void* ptr, struct Moc_PackedString hash)
{
	static_cast<CtxObject721036*>(ptr)->setHashDefault(QString::fromUtf8(hash.data, hash.len));
}

void CtxObject721036_ConnectHashChanged(void* ptr)
{
	QObject::connect(static_cast<CtxObject721036*>(ptr), static_cast<void (CtxObject721036::*)(QString)>(&CtxObject721036::hashChanged), static_cast<CtxObject721036*>(ptr), static_cast<void (CtxObject721036::*)(QString)>(&CtxObject721036::Signal_HashChanged));
}

void CtxObject721036_DisconnectHashChanged(void* ptr)
{
	QObject::disconnect(static_cast<CtxObject721036*>(ptr), static_cast<void (CtxObject721036::*)(QString)>(&CtxObject721036::hashChanged), static_cast<CtxObject721036*>(ptr), static_cast<void (CtxObject721036::*)(QString)>(&CtxObject721036::Signal_HashChanged));
}

void CtxObject721036_HashChanged(void* ptr, struct Moc_PackedString hash)
{
	static_cast<CtxObject721036*>(ptr)->hashChanged(QString::fromUtf8(hash.data, hash.len));
}

int CtxObject721036_CtxObject721036_QRegisterMetaType()
{
	return qRegisterMetaType<CtxObject721036*>();
}

int CtxObject721036_CtxObject721036_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<CtxObject721036*>(const_cast<const char*>(typeName));
}

int CtxObject721036_CtxObject721036_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<CtxObject721036>();
#else
	return 0;
#endif
}

int CtxObject721036_CtxObject721036_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<CtxObject721036>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

void* CtxObject721036___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void CtxObject721036___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* CtxObject721036___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* CtxObject721036___findChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void CtxObject721036___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* CtxObject721036___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* CtxObject721036___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void CtxObject721036___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* CtxObject721036___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* CtxObject721036___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void CtxObject721036___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* CtxObject721036___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* CtxObject721036___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void CtxObject721036___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* CtxObject721036___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* CtxObject721036_NewCtxObject(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new CtxObject721036(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new CtxObject721036(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new CtxObject721036(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new CtxObject721036(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new CtxObject721036(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new CtxObject721036(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new CtxObject721036(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new CtxObject721036(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new CtxObject721036(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new CtxObject721036(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new CtxObject721036(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new CtxObject721036(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new CtxObject721036(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new CtxObject721036(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new CtxObject721036(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new CtxObject721036(static_cast<QWindow*>(parent));
	} else {
		return new CtxObject721036(static_cast<QObject*>(parent));
	}
}

void CtxObject721036_DestroyCtxObject(void* ptr)
{
	static_cast<CtxObject721036*>(ptr)->~CtxObject721036();
}

void CtxObject721036_DestroyCtxObjectDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

char CtxObject721036_EventDefault(void* ptr, void* e)
{
	return static_cast<CtxObject721036*>(ptr)->QObject::event(static_cast<QEvent*>(e));
}

char CtxObject721036_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<CtxObject721036*>(ptr)->QObject::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void CtxObject721036_ChildEventDefault(void* ptr, void* event)
{
	static_cast<CtxObject721036*>(ptr)->QObject::childEvent(static_cast<QChildEvent*>(event));
}

void CtxObject721036_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<CtxObject721036*>(ptr)->QObject::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void CtxObject721036_CustomEventDefault(void* ptr, void* event)
{
	static_cast<CtxObject721036*>(ptr)->QObject::customEvent(static_cast<QEvent*>(event));
}

void CtxObject721036_DeleteLaterDefault(void* ptr)
{
	static_cast<CtxObject721036*>(ptr)->QObject::deleteLater();
}

void CtxObject721036_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<CtxObject721036*>(ptr)->QObject::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void CtxObject721036_TimerEventDefault(void* ptr, void* event)
{
	static_cast<CtxObject721036*>(ptr)->QObject::timerEvent(static_cast<QTimerEvent*>(event));
}



void CustomLabel721036_ConnectUpdateText(void* ptr)
{
	QObject::connect(static_cast<CustomLabel721036*>(ptr), static_cast<void (CustomLabel721036::*)(QString)>(&CustomLabel721036::updateText), static_cast<CustomLabel721036*>(ptr), static_cast<void (CustomLabel721036::*)(QString)>(&CustomLabel721036::Signal_UpdateText));
}

void CustomLabel721036_DisconnectUpdateText(void* ptr)
{
	QObject::disconnect(static_cast<CustomLabel721036*>(ptr), static_cast<void (CustomLabel721036::*)(QString)>(&CustomLabel721036::updateText), static_cast<CustomLabel721036*>(ptr), static_cast<void (CustomLabel721036::*)(QString)>(&CustomLabel721036::Signal_UpdateText));
}

void CustomLabel721036_UpdateText(void* ptr, struct Moc_PackedString v0)
{
	static_cast<CustomLabel721036*>(ptr)->updateText(QString::fromUtf8(v0.data, v0.len));
}

int CustomLabel721036_CustomLabel721036_QRegisterMetaType()
{
	return qRegisterMetaType<CustomLabel721036*>();
}

int CustomLabel721036_CustomLabel721036_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<CustomLabel721036*>(const_cast<const char*>(typeName));
}

int CustomLabel721036_CustomLabel721036_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<CustomLabel721036>();
#else
	return 0;
#endif
}

int CustomLabel721036_CustomLabel721036_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<CustomLabel721036>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

void* CustomLabel721036___addActions_actions_atList(void* ptr, int i)
{
	return ({QAction * tmp = static_cast<QList<QAction *>*>(ptr)->at(i); if (i == static_cast<QList<QAction *>*>(ptr)->size()-1) { static_cast<QList<QAction *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void CustomLabel721036___addActions_actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* CustomLabel721036___addActions_actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAction *>();
}

void* CustomLabel721036___insertActions_actions_atList(void* ptr, int i)
{
	return ({QAction * tmp = static_cast<QList<QAction *>*>(ptr)->at(i); if (i == static_cast<QList<QAction *>*>(ptr)->size()-1) { static_cast<QList<QAction *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void CustomLabel721036___insertActions_actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* CustomLabel721036___insertActions_actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAction *>();
}

void* CustomLabel721036___actions_atList(void* ptr, int i)
{
	return ({QAction * tmp = static_cast<QList<QAction *>*>(ptr)->at(i); if (i == static_cast<QList<QAction *>*>(ptr)->size()-1) { static_cast<QList<QAction *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void CustomLabel721036___actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* CustomLabel721036___actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAction *>();
}

void* CustomLabel721036___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void CustomLabel721036___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* CustomLabel721036___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* CustomLabel721036___findChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void CustomLabel721036___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* CustomLabel721036___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* CustomLabel721036___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void CustomLabel721036___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* CustomLabel721036___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* CustomLabel721036___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void CustomLabel721036___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* CustomLabel721036___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* CustomLabel721036___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void CustomLabel721036___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* CustomLabel721036___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* CustomLabel721036_NewCustomLabel(void* parent, long long fo)
{
		return new CustomLabel721036(static_cast<QWidget*>(parent), static_cast<Qt::WindowType>(fo));
}

void* CustomLabel721036_NewCustomLabel2(struct Moc_PackedString text, void* parent, long long fo)
{
		return new CustomLabel721036(QString::fromUtf8(text.data, text.len), static_cast<QWidget*>(parent), static_cast<Qt::WindowType>(fo));
}

void CustomLabel721036_DestroyCustomLabel(void* ptr)
{
	static_cast<CustomLabel721036*>(ptr)->~CustomLabel721036();
}

void CustomLabel721036_DestroyCustomLabelDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

char CustomLabel721036_EventDefault(void* ptr, void* e)
{
	return static_cast<CustomLabel721036*>(ptr)->QLabel::event(static_cast<QEvent*>(e));
}

char CustomLabel721036_FocusNextPrevChildDefault(void* ptr, char next)
{
	return static_cast<CustomLabel721036*>(ptr)->QLabel::focusNextPrevChild(next != 0);
}

void CustomLabel721036_ChangeEventDefault(void* ptr, void* ev)
{
	static_cast<CustomLabel721036*>(ptr)->QLabel::changeEvent(static_cast<QEvent*>(ev));
}

void CustomLabel721036_ClearDefault(void* ptr)
{
	static_cast<CustomLabel721036*>(ptr)->QLabel::clear();
}

void CustomLabel721036_ContextMenuEventDefault(void* ptr, void* ev)
{
	static_cast<CustomLabel721036*>(ptr)->QLabel::contextMenuEvent(static_cast<QContextMenuEvent*>(ev));
}

void CustomLabel721036_FocusInEventDefault(void* ptr, void* ev)
{
	static_cast<CustomLabel721036*>(ptr)->QLabel::focusInEvent(static_cast<QFocusEvent*>(ev));
}

void CustomLabel721036_FocusOutEventDefault(void* ptr, void* ev)
{
	static_cast<CustomLabel721036*>(ptr)->QLabel::focusOutEvent(static_cast<QFocusEvent*>(ev));
}

void CustomLabel721036_KeyPressEventDefault(void* ptr, void* ev)
{
	static_cast<CustomLabel721036*>(ptr)->QLabel::keyPressEvent(static_cast<QKeyEvent*>(ev));
}

void CustomLabel721036_MouseMoveEventDefault(void* ptr, void* ev)
{
	static_cast<CustomLabel721036*>(ptr)->QLabel::mouseMoveEvent(static_cast<QMouseEvent*>(ev));
}

void CustomLabel721036_MousePressEventDefault(void* ptr, void* ev)
{
	static_cast<CustomLabel721036*>(ptr)->QLabel::mousePressEvent(static_cast<QMouseEvent*>(ev));
}

void CustomLabel721036_MouseReleaseEventDefault(void* ptr, void* ev)
{
	static_cast<CustomLabel721036*>(ptr)->QLabel::mouseReleaseEvent(static_cast<QMouseEvent*>(ev));
}

void CustomLabel721036_PaintEventDefault(void* ptr, void* vqp)
{
	static_cast<CustomLabel721036*>(ptr)->QLabel::paintEvent(static_cast<QPaintEvent*>(vqp));
}

void CustomLabel721036_SetMovieDefault(void* ptr, void* movie)
{
	static_cast<CustomLabel721036*>(ptr)->QLabel::setMovie(static_cast<QMovie*>(movie));
}

void CustomLabel721036_SetNum2Default(void* ptr, double num)
{
	static_cast<CustomLabel721036*>(ptr)->QLabel::setNum(num);
}

void CustomLabel721036_SetNumDefault(void* ptr, int num)
{
	static_cast<CustomLabel721036*>(ptr)->QLabel::setNum(num);
}

void CustomLabel721036_SetPictureDefault(void* ptr, void* picture)
{
	static_cast<CustomLabel721036*>(ptr)->QLabel::setPicture(*static_cast<QPicture*>(picture));
}

void CustomLabel721036_SetPixmapDefault(void* ptr, void* vqp)
{
	static_cast<CustomLabel721036*>(ptr)->QLabel::setPixmap(*static_cast<QPixmap*>(vqp));
}

void CustomLabel721036_SetTextDefault(void* ptr, struct Moc_PackedString vqs)
{
	static_cast<CustomLabel721036*>(ptr)->QLabel::setText(QString::fromUtf8(vqs.data, vqs.len));
}

void* CustomLabel721036_MinimumSizeHintDefault(void* ptr)
{
	return ({ QSize tmpValue = static_cast<CustomLabel721036*>(ptr)->QLabel::minimumSizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* CustomLabel721036_SizeHintDefault(void* ptr)
{
	return ({ QSize tmpValue = static_cast<CustomLabel721036*>(ptr)->QLabel::sizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}



int CustomLabel721036_HeightForWidthDefault(void* ptr, int w)
{
	return static_cast<CustomLabel721036*>(ptr)->QLabel::heightForWidth(w);
}

char CustomLabel721036_CloseDefault(void* ptr)
{
	return static_cast<CustomLabel721036*>(ptr)->QLabel::close();
}

void CustomLabel721036_ActionEventDefault(void* ptr, void* event)
{
	static_cast<CustomLabel721036*>(ptr)->QLabel::actionEvent(static_cast<QActionEvent*>(event));
}

void CustomLabel721036_CloseEventDefault(void* ptr, void* event)
{
	static_cast<CustomLabel721036*>(ptr)->QLabel::closeEvent(static_cast<QCloseEvent*>(event));
}

void CustomLabel721036_DragEnterEventDefault(void* ptr, void* event)
{
	static_cast<CustomLabel721036*>(ptr)->QLabel::dragEnterEvent(static_cast<QDragEnterEvent*>(event));
}

void CustomLabel721036_DragLeaveEventDefault(void* ptr, void* event)
{
	static_cast<CustomLabel721036*>(ptr)->QLabel::dragLeaveEvent(static_cast<QDragLeaveEvent*>(event));
}

void CustomLabel721036_DragMoveEventDefault(void* ptr, void* event)
{
	static_cast<CustomLabel721036*>(ptr)->QLabel::dragMoveEvent(static_cast<QDragMoveEvent*>(event));
}

void CustomLabel721036_DropEventDefault(void* ptr, void* event)
{
	static_cast<CustomLabel721036*>(ptr)->QLabel::dropEvent(static_cast<QDropEvent*>(event));
}

void CustomLabel721036_EnterEventDefault(void* ptr, void* event)
{
	static_cast<CustomLabel721036*>(ptr)->QLabel::enterEvent(static_cast<QEvent*>(event));
}

void CustomLabel721036_HideDefault(void* ptr)
{
	static_cast<CustomLabel721036*>(ptr)->QLabel::hide();
}

void CustomLabel721036_HideEventDefault(void* ptr, void* event)
{
	static_cast<CustomLabel721036*>(ptr)->QLabel::hideEvent(static_cast<QHideEvent*>(event));
}

void CustomLabel721036_InputMethodEventDefault(void* ptr, void* event)
{
	static_cast<CustomLabel721036*>(ptr)->QLabel::inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void CustomLabel721036_KeyReleaseEventDefault(void* ptr, void* event)
{
	static_cast<CustomLabel721036*>(ptr)->QLabel::keyReleaseEvent(static_cast<QKeyEvent*>(event));
}

void CustomLabel721036_LeaveEventDefault(void* ptr, void* event)
{
	static_cast<CustomLabel721036*>(ptr)->QLabel::leaveEvent(static_cast<QEvent*>(event));
}

void CustomLabel721036_LowerDefault(void* ptr)
{
	static_cast<CustomLabel721036*>(ptr)->QLabel::lower();
}

void CustomLabel721036_MouseDoubleClickEventDefault(void* ptr, void* event)
{
	static_cast<CustomLabel721036*>(ptr)->QLabel::mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
}

void CustomLabel721036_MoveEventDefault(void* ptr, void* event)
{
	static_cast<CustomLabel721036*>(ptr)->QLabel::moveEvent(static_cast<QMoveEvent*>(event));
}

void CustomLabel721036_RaiseDefault(void* ptr)
{
	static_cast<CustomLabel721036*>(ptr)->QLabel::raise();
}

void CustomLabel721036_RepaintDefault(void* ptr)
{
	static_cast<CustomLabel721036*>(ptr)->QLabel::repaint();
}

void CustomLabel721036_ResizeEventDefault(void* ptr, void* event)
{
	static_cast<CustomLabel721036*>(ptr)->QLabel::resizeEvent(static_cast<QResizeEvent*>(event));
}

void CustomLabel721036_SetDisabledDefault(void* ptr, char disable)
{
	static_cast<CustomLabel721036*>(ptr)->QLabel::setDisabled(disable != 0);
}

void CustomLabel721036_SetEnabledDefault(void* ptr, char vbo)
{
	static_cast<CustomLabel721036*>(ptr)->QLabel::setEnabled(vbo != 0);
}

void CustomLabel721036_SetFocus2Default(void* ptr)
{
	static_cast<CustomLabel721036*>(ptr)->QLabel::setFocus();
}

void CustomLabel721036_SetHiddenDefault(void* ptr, char hidden)
{
	static_cast<CustomLabel721036*>(ptr)->QLabel::setHidden(hidden != 0);
}

void CustomLabel721036_SetStyleSheetDefault(void* ptr, struct Moc_PackedString styleSheet)
{
	static_cast<CustomLabel721036*>(ptr)->QLabel::setStyleSheet(QString::fromUtf8(styleSheet.data, styleSheet.len));
}

void CustomLabel721036_SetVisibleDefault(void* ptr, char visible)
{
	static_cast<CustomLabel721036*>(ptr)->QLabel::setVisible(visible != 0);
}

void CustomLabel721036_SetWindowModifiedDefault(void* ptr, char vbo)
{
	static_cast<CustomLabel721036*>(ptr)->QLabel::setWindowModified(vbo != 0);
}

void CustomLabel721036_SetWindowTitleDefault(void* ptr, struct Moc_PackedString vqs)
{
	static_cast<CustomLabel721036*>(ptr)->QLabel::setWindowTitle(QString::fromUtf8(vqs.data, vqs.len));
}

void CustomLabel721036_ShowDefault(void* ptr)
{
	static_cast<CustomLabel721036*>(ptr)->QLabel::show();
}

void CustomLabel721036_ShowEventDefault(void* ptr, void* event)
{
	static_cast<CustomLabel721036*>(ptr)->QLabel::showEvent(static_cast<QShowEvent*>(event));
}

void CustomLabel721036_ShowFullScreenDefault(void* ptr)
{
	static_cast<CustomLabel721036*>(ptr)->QLabel::showFullScreen();
}

void CustomLabel721036_ShowMaximizedDefault(void* ptr)
{
	static_cast<CustomLabel721036*>(ptr)->QLabel::showMaximized();
}

void CustomLabel721036_ShowMinimizedDefault(void* ptr)
{
	static_cast<CustomLabel721036*>(ptr)->QLabel::showMinimized();
}

void CustomLabel721036_ShowNormalDefault(void* ptr)
{
	static_cast<CustomLabel721036*>(ptr)->QLabel::showNormal();
}

void CustomLabel721036_TabletEventDefault(void* ptr, void* event)
{
	static_cast<CustomLabel721036*>(ptr)->QLabel::tabletEvent(static_cast<QTabletEvent*>(event));
}

void CustomLabel721036_UpdateDefault(void* ptr)
{
	static_cast<CustomLabel721036*>(ptr)->QLabel::update();
}

void CustomLabel721036_UpdateMicroFocusDefault(void* ptr)
{
	static_cast<CustomLabel721036*>(ptr)->QLabel::updateMicroFocus();
}

void CustomLabel721036_WheelEventDefault(void* ptr, void* event)
{
	static_cast<CustomLabel721036*>(ptr)->QLabel::wheelEvent(static_cast<QWheelEvent*>(event));
}

void* CustomLabel721036_PaintEngineDefault(void* ptr)
{
	return static_cast<CustomLabel721036*>(ptr)->QLabel::paintEngine();
}

void* CustomLabel721036_InputMethodQueryDefault(void* ptr, long long query)
{
	return new QVariant(static_cast<CustomLabel721036*>(ptr)->QLabel::inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
}

char CustomLabel721036_HasHeightForWidthDefault(void* ptr)
{
	return static_cast<CustomLabel721036*>(ptr)->QLabel::hasHeightForWidth();
}

int CustomLabel721036_MetricDefault(void* ptr, long long m)
{
	return static_cast<CustomLabel721036*>(ptr)->QLabel::metric(static_cast<QPaintDevice::PaintDeviceMetric>(m));
}

void CustomLabel721036_InitPainterDefault(void* ptr, void* painter)
{
	static_cast<CustomLabel721036*>(ptr)->QLabel::initPainter(static_cast<QPainter*>(painter));
}

char CustomLabel721036_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<CustomLabel721036*>(ptr)->QLabel::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void CustomLabel721036_ChildEventDefault(void* ptr, void* event)
{
	static_cast<CustomLabel721036*>(ptr)->QLabel::childEvent(static_cast<QChildEvent*>(event));
}

void CustomLabel721036_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<CustomLabel721036*>(ptr)->QLabel::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void CustomLabel721036_CustomEventDefault(void* ptr, void* event)
{
	static_cast<CustomLabel721036*>(ptr)->QLabel::customEvent(static_cast<QEvent*>(event));
}

void CustomLabel721036_DeleteLaterDefault(void* ptr)
{
	static_cast<CustomLabel721036*>(ptr)->QLabel::deleteLater();
}

void CustomLabel721036_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<CustomLabel721036*>(ptr)->QLabel::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void CustomLabel721036_TimerEventDefault(void* ptr, void* event)
{
	static_cast<CustomLabel721036*>(ptr)->QLabel::timerEvent(static_cast<QTimerEvent*>(event));
}

#include "moc_moc.h"
