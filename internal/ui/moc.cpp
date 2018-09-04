

#define protected public
#define private public

#include "moc.h"
#include "_cgo_export.h"

#include <QAbstractItemModel>
#include <QAbstractListModel>
#include <QByteArray>
#include <QCamera>
#include <QCameraImageCapture>
#include <QChildEvent>
#include <QDBusPendingCall>
#include <QDBusPendingCallWatcher>
#include <QEvent>
#include <QExtensionFactory>
#include <QExtensionManager>
#include <QGraphicsObject>
#include <QGraphicsWidget>
#include <QHash>
#include <QLayout>
#include <QMap>
#include <QMediaPlaylist>
#include <QMediaRecorder>
#include <QMetaMethod>
#include <QMimeData>
#include <QModelIndex>
#include <QObject>
#include <QOffscreenSurface>
#include <QPaintDevice>
#include <QPaintDeviceWindow>
#include <QPdfWriter>
#include <QPersistentModelIndex>
#include <QQuickItem>
#include <QRadioData>
#include <QSize>
#include <QString>
#include <QTime>
#include <QTimer>
#include <QTimerEvent>
#include <QVariant>
#include <QVector>
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

class CustomListModel721036: public QAbstractListModel
{
Q_OBJECT
public:
	CustomListModel721036(QObject *parent = Q_NULLPTR) : QAbstractListModel(parent) {qRegisterMetaType<quintptr>("quintptr");CustomListModel721036_CustomListModel721036_QRegisterMetaType();CustomListModel721036_CustomListModel721036_QRegisterMetaTypes();callbackCustomListModel721036_Constructor(this);};
	void Signal_Remove() { callbackCustomListModel721036_Remove(this); };
	void Signal_Add(QList<QVariant> obj) { callbackCustomListModel721036_Add(this, ({ QList<QVariant>* tmpValue = new QList<QVariant>(obj); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	void Signal_Edit(QString address, QString checked) { QByteArray tc66218 = address.toUtf8(); Moc_PackedString addressPacked = { const_cast<char*>(tc66218.prepend("WHITESPACE").constData()+10), tc66218.size()-10 };QByteArray t75e4ae = checked.toUtf8(); Moc_PackedString checkedPacked = { const_cast<char*>(t75e4ae.prepend("WHITESPACE").constData()+10), t75e4ae.size()-10 };callbackCustomListModel721036_Edit(this, addressPacked, checkedPacked); };
	 ~CustomListModel721036() { callbackCustomListModel721036_DestroyCustomListModel(this); };
	bool dropMimeData(const QMimeData * data, Qt::DropAction action, int row, int column, const QModelIndex & parent) { return callbackCustomListModel721036_DropMimeData(this, const_cast<QMimeData*>(data), action, row, column, const_cast<QModelIndex*>(&parent)) != 0; };
	
	QModelIndex index(int row, int column, const QModelIndex & parent) const { return *static_cast<QModelIndex*>(callbackCustomListModel721036_Index(const_cast<void*>(static_cast<const void*>(this)), row, column, const_cast<QModelIndex*>(&parent))); };
	QModelIndex sibling(int row, int column, const QModelIndex & idx) const { return *static_cast<QModelIndex*>(callbackCustomListModel721036_Sibling(const_cast<void*>(static_cast<const void*>(this)), row, column, const_cast<QModelIndex*>(&idx))); };
	Qt::ItemFlags flags(const QModelIndex & index) const { return static_cast<Qt::ItemFlag>(callbackCustomListModel721036_Flags(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	
	bool insertColumns(int column, int count, const QModelIndex & parent) { return callbackCustomListModel721036_InsertColumns(this, column, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool insertRows(int row, int count, const QModelIndex & parent) { return callbackCustomListModel721036_InsertRows(this, row, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool moveColumns(const QModelIndex & sourceParent, int sourceColumn, int count, const QModelIndex & destinationParent, int destinationChild) { return callbackCustomListModel721036_MoveColumns(this, const_cast<QModelIndex*>(&sourceParent), sourceColumn, count, const_cast<QModelIndex*>(&destinationParent), destinationChild) != 0; };
	bool moveRows(const QModelIndex & sourceParent, int sourceRow, int count, const QModelIndex & destinationParent, int destinationChild) { return callbackCustomListModel721036_MoveRows(this, const_cast<QModelIndex*>(&sourceParent), sourceRow, count, const_cast<QModelIndex*>(&destinationParent), destinationChild) != 0; };
	bool removeColumns(int column, int count, const QModelIndex & parent) { return callbackCustomListModel721036_RemoveColumns(this, column, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool removeRows(int row, int count, const QModelIndex & parent) { return callbackCustomListModel721036_RemoveRows(this, row, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool setData(const QModelIndex & index, const QVariant & value, int role) { return callbackCustomListModel721036_SetData(this, const_cast<QModelIndex*>(&index), const_cast<QVariant*>(&value), role) != 0; };
	bool setHeaderData(int section, Qt::Orientation orientation, const QVariant & value, int role) { return callbackCustomListModel721036_SetHeaderData(this, section, orientation, const_cast<QVariant*>(&value), role) != 0; };
	bool setItemData(const QModelIndex & index, const QMap<int, QVariant> & roles) { return callbackCustomListModel721036_SetItemData(this, const_cast<QModelIndex*>(&index), ({ QMap<int, QVariant>* tmpValue = const_cast<QMap<int, QVariant>*>(&roles); Moc_PackedList { tmpValue, tmpValue->size() }; })) != 0; };
	bool submit() { return callbackCustomListModel721036_Submit(this) != 0; };
	void Signal_ColumnsAboutToBeInserted(const QModelIndex & parent, int first, int last) { callbackCustomListModel721036_ColumnsAboutToBeInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsAboutToBeMoved(const QModelIndex & sourceParent, int sourceStart, int sourceEnd, const QModelIndex & destinationParent, int destinationColumn) { callbackCustomListModel721036_ColumnsAboutToBeMoved(this, const_cast<QModelIndex*>(&sourceParent), sourceStart, sourceEnd, const_cast<QModelIndex*>(&destinationParent), destinationColumn); };
	void Signal_ColumnsAboutToBeRemoved(const QModelIndex & parent, int first, int last) { callbackCustomListModel721036_ColumnsAboutToBeRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsInserted(const QModelIndex & parent, int first, int last) { callbackCustomListModel721036_ColumnsInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsMoved(const QModelIndex & parent, int start, int end, const QModelIndex & destination, int column) { callbackCustomListModel721036_ColumnsMoved(this, const_cast<QModelIndex*>(&parent), start, end, const_cast<QModelIndex*>(&destination), column); };
	void Signal_ColumnsRemoved(const QModelIndex & parent, int first, int last) { callbackCustomListModel721036_ColumnsRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_DataChanged(const QModelIndex & topLeft, const QModelIndex & bottomRight, const QVector<int> & roles) { callbackCustomListModel721036_DataChanged(this, const_cast<QModelIndex*>(&topLeft), const_cast<QModelIndex*>(&bottomRight), ({ QVector<int>* tmpValue = const_cast<QVector<int>*>(&roles); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	void fetchMore(const QModelIndex & parent) { callbackCustomListModel721036_FetchMore(this, const_cast<QModelIndex*>(&parent)); };
	void Signal_HeaderDataChanged(Qt::Orientation orientation, int first, int last) { callbackCustomListModel721036_HeaderDataChanged(this, orientation, first, last); };
	void Signal_LayoutAboutToBeChanged(const QList<QPersistentModelIndex> & parents, QAbstractItemModel::LayoutChangeHint hint) { callbackCustomListModel721036_LayoutAboutToBeChanged(this, ({ QList<QPersistentModelIndex>* tmpValue = const_cast<QList<QPersistentModelIndex>*>(&parents); Moc_PackedList { tmpValue, tmpValue->size() }; }), hint); };
	void Signal_LayoutChanged(const QList<QPersistentModelIndex> & parents, QAbstractItemModel::LayoutChangeHint hint) { callbackCustomListModel721036_LayoutChanged(this, ({ QList<QPersistentModelIndex>* tmpValue = const_cast<QList<QPersistentModelIndex>*>(&parents); Moc_PackedList { tmpValue, tmpValue->size() }; }), hint); };
	void Signal_ModelAboutToBeReset() { callbackCustomListModel721036_ModelAboutToBeReset(this); };
	void Signal_ModelReset() { callbackCustomListModel721036_ModelReset(this); };
	void resetInternalData() { callbackCustomListModel721036_ResetInternalData(this); };
	void revert() { callbackCustomListModel721036_Revert(this); };
	void Signal_RowsAboutToBeInserted(const QModelIndex & parent, int start, int end) { callbackCustomListModel721036_RowsAboutToBeInserted(this, const_cast<QModelIndex*>(&parent), start, end); };
	void Signal_RowsAboutToBeMoved(const QModelIndex & sourceParent, int sourceStart, int sourceEnd, const QModelIndex & destinationParent, int destinationRow) { callbackCustomListModel721036_RowsAboutToBeMoved(this, const_cast<QModelIndex*>(&sourceParent), sourceStart, sourceEnd, const_cast<QModelIndex*>(&destinationParent), destinationRow); };
	void Signal_RowsAboutToBeRemoved(const QModelIndex & parent, int first, int last) { callbackCustomListModel721036_RowsAboutToBeRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_RowsInserted(const QModelIndex & parent, int first, int last) { callbackCustomListModel721036_RowsInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_RowsMoved(const QModelIndex & parent, int start, int end, const QModelIndex & destination, int row) { callbackCustomListModel721036_RowsMoved(this, const_cast<QModelIndex*>(&parent), start, end, const_cast<QModelIndex*>(&destination), row); };
	void Signal_RowsRemoved(const QModelIndex & parent, int first, int last) { callbackCustomListModel721036_RowsRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void sort(int column, Qt::SortOrder order) { callbackCustomListModel721036_Sort(this, column, order); };
	QHash<int, QByteArray> roleNames() const { return ({ QHash<int, QByteArray>* tmpP = static_cast<QHash<int, QByteArray>*>(callbackCustomListModel721036_RoleNames(const_cast<void*>(static_cast<const void*>(this)))); QHash<int, QByteArray> tmpV = *tmpP; tmpP->~QHash(); free(tmpP); tmpV; }); };
	QMap<int, QVariant> itemData(const QModelIndex & index) const { return ({ QMap<int, QVariant>* tmpP = static_cast<QMap<int, QVariant>*>(callbackCustomListModel721036_ItemData(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); QMap<int, QVariant> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }); };
	QMimeData * mimeData(const QModelIndexList & indexes) const { return static_cast<QMimeData*>(callbackCustomListModel721036_MimeData(const_cast<void*>(static_cast<const void*>(this)), ({ QList<QModelIndex>* tmpValue = new QList<QModelIndex>(indexes); Moc_PackedList { tmpValue, tmpValue->size() }; }))); };
	QModelIndex buddy(const QModelIndex & index) const { return *static_cast<QModelIndex*>(callbackCustomListModel721036_Buddy(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	QModelIndex parent(const QModelIndex & index) const { return *static_cast<QModelIndex*>(callbackCustomListModel721036_Parent(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	QList<QModelIndex> match(const QModelIndex & start, int role, const QVariant & value, int hits, Qt::MatchFlags flags) const { return ({ QList<QModelIndex>* tmpP = static_cast<QList<QModelIndex>*>(callbackCustomListModel721036_Match(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&start), role, const_cast<QVariant*>(&value), hits, flags)); QList<QModelIndex> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }); };
	QSize span(const QModelIndex & index) const { return *static_cast<QSize*>(callbackCustomListModel721036_Span(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	QStringList mimeTypes() const { return ({ Moc_PackedString tempVal = callbackCustomListModel721036_MimeTypes(const_cast<void*>(static_cast<const void*>(this))); QStringList ret = QString::fromUtf8(tempVal.data, tempVal.len).split("|", QString::SkipEmptyParts); free(tempVal.data); ret; }); };
	QVariant data(const QModelIndex & index, int role) const { return *static_cast<QVariant*>(callbackCustomListModel721036_Data(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index), role)); };
	QVariant headerData(int section, Qt::Orientation orientation, int role) const { return *static_cast<QVariant*>(callbackCustomListModel721036_HeaderData(const_cast<void*>(static_cast<const void*>(this)), section, orientation, role)); };
	Qt::DropActions supportedDragActions() const { return static_cast<Qt::DropAction>(callbackCustomListModel721036_SupportedDragActions(const_cast<void*>(static_cast<const void*>(this)))); };
	Qt::DropActions supportedDropActions() const { return static_cast<Qt::DropAction>(callbackCustomListModel721036_SupportedDropActions(const_cast<void*>(static_cast<const void*>(this)))); };
	bool canDropMimeData(const QMimeData * data, Qt::DropAction action, int row, int column, const QModelIndex & parent) const { return callbackCustomListModel721036_CanDropMimeData(const_cast<void*>(static_cast<const void*>(this)), const_cast<QMimeData*>(data), action, row, column, const_cast<QModelIndex*>(&parent)) != 0; };
	bool canFetchMore(const QModelIndex & parent) const { return callbackCustomListModel721036_CanFetchMore(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)) != 0; };
	bool hasChildren(const QModelIndex & parent) const { return callbackCustomListModel721036_HasChildren(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)) != 0; };
	int columnCount(const QModelIndex & parent) const { return callbackCustomListModel721036_ColumnCount(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)); };
	int rowCount(const QModelIndex & parent) const { return callbackCustomListModel721036_RowCount(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)); };
	bool event(QEvent * e) { return callbackCustomListModel721036_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackCustomListModel721036_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackCustomListModel721036_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackCustomListModel721036_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackCustomListModel721036_CustomEvent(this, event); };
	void deleteLater() { callbackCustomListModel721036_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackCustomListModel721036_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackCustomListModel721036_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackCustomListModel721036_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackCustomListModel721036_TimerEvent(this, event); };
signals:
	void remove();
	void add(QList<QVariant> obj);
	void edit(QString address, QString checked);
public slots:
private:
};

Q_DECLARE_METATYPE(CustomListModel721036*)


void CustomListModel721036_CustomListModel721036_QRegisterMetaTypes() {
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



void CustomListModel721036_ConnectRemove(void* ptr)
{
	QObject::connect(static_cast<CustomListModel721036*>(ptr), static_cast<void (CustomListModel721036::*)()>(&CustomListModel721036::remove), static_cast<CustomListModel721036*>(ptr), static_cast<void (CustomListModel721036::*)()>(&CustomListModel721036::Signal_Remove));
}

void CustomListModel721036_DisconnectRemove(void* ptr)
{
	QObject::disconnect(static_cast<CustomListModel721036*>(ptr), static_cast<void (CustomListModel721036::*)()>(&CustomListModel721036::remove), static_cast<CustomListModel721036*>(ptr), static_cast<void (CustomListModel721036::*)()>(&CustomListModel721036::Signal_Remove));
}

void CustomListModel721036_Remove(void* ptr)
{
	static_cast<CustomListModel721036*>(ptr)->remove();
}

void CustomListModel721036_ConnectAdd(void* ptr)
{
	QObject::connect(static_cast<CustomListModel721036*>(ptr), static_cast<void (CustomListModel721036::*)(QList<QVariant>)>(&CustomListModel721036::add), static_cast<CustomListModel721036*>(ptr), static_cast<void (CustomListModel721036::*)(QList<QVariant>)>(&CustomListModel721036::Signal_Add));
}

void CustomListModel721036_DisconnectAdd(void* ptr)
{
	QObject::disconnect(static_cast<CustomListModel721036*>(ptr), static_cast<void (CustomListModel721036::*)(QList<QVariant>)>(&CustomListModel721036::add), static_cast<CustomListModel721036*>(ptr), static_cast<void (CustomListModel721036::*)(QList<QVariant>)>(&CustomListModel721036::Signal_Add));
}

void CustomListModel721036_Add(void* ptr, void* obj)
{
	static_cast<CustomListModel721036*>(ptr)->add(({ QList<QVariant>* tmpP = static_cast<QList<QVariant>*>(obj); QList<QVariant> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

void CustomListModel721036_ConnectEdit(void* ptr)
{
	QObject::connect(static_cast<CustomListModel721036*>(ptr), static_cast<void (CustomListModel721036::*)(QString, QString)>(&CustomListModel721036::edit), static_cast<CustomListModel721036*>(ptr), static_cast<void (CustomListModel721036::*)(QString, QString)>(&CustomListModel721036::Signal_Edit));
}

void CustomListModel721036_DisconnectEdit(void* ptr)
{
	QObject::disconnect(static_cast<CustomListModel721036*>(ptr), static_cast<void (CustomListModel721036::*)(QString, QString)>(&CustomListModel721036::edit), static_cast<CustomListModel721036*>(ptr), static_cast<void (CustomListModel721036::*)(QString, QString)>(&CustomListModel721036::Signal_Edit));
}

void CustomListModel721036_Edit(void* ptr, struct Moc_PackedString address, struct Moc_PackedString checked)
{
	static_cast<CustomListModel721036*>(ptr)->edit(QString::fromUtf8(address.data, address.len), QString::fromUtf8(checked.data, checked.len));
}

int CustomListModel721036_CustomListModel721036_QRegisterMetaType()
{
	return qRegisterMetaType<CustomListModel721036*>();
}

int CustomListModel721036_CustomListModel721036_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<CustomListModel721036*>(const_cast<const char*>(typeName));
}

int CustomListModel721036_CustomListModel721036_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<CustomListModel721036>();
#else
	return 0;
#endif
}

int CustomListModel721036_CustomListModel721036_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<CustomListModel721036>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

int CustomListModel721036_____setItemData_roles_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void CustomListModel721036_____setItemData_roles_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* CustomListModel721036_____setItemData_roles_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int CustomListModel721036_____roleNames_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void CustomListModel721036_____roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* CustomListModel721036_____roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int CustomListModel721036_____itemData_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void CustomListModel721036_____itemData_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* CustomListModel721036_____itemData_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

void* CustomListModel721036___setItemData_roles_atList(void* ptr, int v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<int, QVariant>*>(ptr)->value(v); if (i == static_cast<QMap<int, QVariant>*>(ptr)->size()-1) { static_cast<QMap<int, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void CustomListModel721036___setItemData_roles_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<int, QVariant>*>(ptr)->insert(key, *static_cast<QVariant*>(i));
}

void* CustomListModel721036___setItemData_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<int, QVariant>();
}

struct Moc_PackedList CustomListModel721036___setItemData_roles_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QMap<int, QVariant>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* CustomListModel721036___changePersistentIndexList_from_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void CustomListModel721036___changePersistentIndexList_from_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* CustomListModel721036___changePersistentIndexList_from_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* CustomListModel721036___changePersistentIndexList_to_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void CustomListModel721036___changePersistentIndexList_to_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* CustomListModel721036___changePersistentIndexList_to_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

int CustomListModel721036___dataChanged_roles_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QVector<int>*>(ptr)->at(i); if (i == static_cast<QVector<int>*>(ptr)->size()-1) { static_cast<QVector<int>*>(ptr)->~QVector(); free(ptr); }; tmp; });
}

void CustomListModel721036___dataChanged_roles_setList(void* ptr, int i)
{
	static_cast<QVector<int>*>(ptr)->append(i);
}

void* CustomListModel721036___dataChanged_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<int>();
}

void* CustomListModel721036___layoutAboutToBeChanged_parents_atList(void* ptr, int i)
{
	return new QPersistentModelIndex(({QPersistentModelIndex tmp = static_cast<QList<QPersistentModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QPersistentModelIndex>*>(ptr)->size()-1) { static_cast<QList<QPersistentModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void CustomListModel721036___layoutAboutToBeChanged_parents_setList(void* ptr, void* i)
{
	static_cast<QList<QPersistentModelIndex>*>(ptr)->append(*static_cast<QPersistentModelIndex*>(i));
}

void* CustomListModel721036___layoutAboutToBeChanged_parents_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QPersistentModelIndex>();
}

void* CustomListModel721036___layoutChanged_parents_atList(void* ptr, int i)
{
	return new QPersistentModelIndex(({QPersistentModelIndex tmp = static_cast<QList<QPersistentModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QPersistentModelIndex>*>(ptr)->size()-1) { static_cast<QList<QPersistentModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void CustomListModel721036___layoutChanged_parents_setList(void* ptr, void* i)
{
	static_cast<QList<QPersistentModelIndex>*>(ptr)->append(*static_cast<QPersistentModelIndex*>(i));
}

void* CustomListModel721036___layoutChanged_parents_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QPersistentModelIndex>();
}

void* CustomListModel721036___roleNames_atList(void* ptr, int v, int i)
{
	return new QByteArray(({ QByteArray tmp = static_cast<QHash<int, QByteArray>*>(ptr)->value(v); if (i == static_cast<QHash<int, QByteArray>*>(ptr)->size()-1) { static_cast<QHash<int, QByteArray>*>(ptr)->~QHash(); free(ptr); }; tmp; }));
}

void CustomListModel721036___roleNames_setList(void* ptr, int key, void* i)
{
	static_cast<QHash<int, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* CustomListModel721036___roleNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QHash<int, QByteArray>();
}

struct Moc_PackedList CustomListModel721036___roleNames_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QHash<int, QByteArray>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* CustomListModel721036___itemData_atList(void* ptr, int v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<int, QVariant>*>(ptr)->value(v); if (i == static_cast<QMap<int, QVariant>*>(ptr)->size()-1) { static_cast<QMap<int, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void CustomListModel721036___itemData_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<int, QVariant>*>(ptr)->insert(key, *static_cast<QVariant*>(i));
}

void* CustomListModel721036___itemData_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<int, QVariant>();
}

struct Moc_PackedList CustomListModel721036___itemData_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QMap<int, QVariant>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* CustomListModel721036___mimeData_indexes_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void CustomListModel721036___mimeData_indexes_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* CustomListModel721036___mimeData_indexes_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* CustomListModel721036___match_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void CustomListModel721036___match_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* CustomListModel721036___match_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* CustomListModel721036___persistentIndexList_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void CustomListModel721036___persistentIndexList_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* CustomListModel721036___persistentIndexList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

int CustomListModel721036_____doSetRoleNames_roleNames_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void CustomListModel721036_____doSetRoleNames_roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* CustomListModel721036_____doSetRoleNames_roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int CustomListModel721036_____setRoleNames_roleNames_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void CustomListModel721036_____setRoleNames_roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* CustomListModel721036_____setRoleNames_roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

void* CustomListModel721036___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void CustomListModel721036___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* CustomListModel721036___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* CustomListModel721036___findChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void CustomListModel721036___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* CustomListModel721036___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* CustomListModel721036___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void CustomListModel721036___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* CustomListModel721036___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* CustomListModel721036___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void CustomListModel721036___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* CustomListModel721036___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* CustomListModel721036___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void CustomListModel721036___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* CustomListModel721036___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* CustomListModel721036___add_obj_atList(void* ptr, int i)
{
	return new QVariant(({QVariant tmp = static_cast<QList<QVariant>*>(ptr)->at(i); if (i == static_cast<QList<QVariant>*>(ptr)->size()-1) { static_cast<QList<QVariant>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void CustomListModel721036___add_obj_setList(void* ptr, void* i)
{
	static_cast<QList<QVariant>*>(ptr)->append(*static_cast<QVariant*>(i));
}

void* CustomListModel721036___add_obj_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QVariant>();
}

void* CustomListModel721036_NewCustomListModel(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new CustomListModel721036(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new CustomListModel721036(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new CustomListModel721036(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new CustomListModel721036(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new CustomListModel721036(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new CustomListModel721036(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new CustomListModel721036(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new CustomListModel721036(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new CustomListModel721036(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new CustomListModel721036(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new CustomListModel721036(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new CustomListModel721036(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new CustomListModel721036(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new CustomListModel721036(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new CustomListModel721036(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new CustomListModel721036(static_cast<QWindow*>(parent));
	} else {
		return new CustomListModel721036(static_cast<QObject*>(parent));
	}
}

void CustomListModel721036_DestroyCustomListModel(void* ptr)
{
	static_cast<CustomListModel721036*>(ptr)->~CustomListModel721036();
}

void CustomListModel721036_DestroyCustomListModelDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

char CustomListModel721036_DropMimeDataDefault(void* ptr, void* data, long long action, int row, int column, void* parent)
{
	return static_cast<CustomListModel721036*>(ptr)->QAbstractListModel::dropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

void* CustomListModel721036_IndexDefault(void* ptr, int row, int column, void* parent)
{
	return new QModelIndex(static_cast<CustomListModel721036*>(ptr)->QAbstractListModel::index(row, column, *static_cast<QModelIndex*>(parent)));
}

void* CustomListModel721036_SiblingDefault(void* ptr, int row, int column, void* idx)
{
	return new QModelIndex(static_cast<CustomListModel721036*>(ptr)->QAbstractListModel::sibling(row, column, *static_cast<QModelIndex*>(idx)));
}

long long CustomListModel721036_FlagsDefault(void* ptr, void* index)
{
	return static_cast<CustomListModel721036*>(ptr)->QAbstractListModel::flags(*static_cast<QModelIndex*>(index));
}



char CustomListModel721036_InsertColumnsDefault(void* ptr, int column, int count, void* parent)
{
	return static_cast<CustomListModel721036*>(ptr)->QAbstractListModel::insertColumns(column, count, *static_cast<QModelIndex*>(parent));
}

char CustomListModel721036_InsertRowsDefault(void* ptr, int row, int count, void* parent)
{
	return static_cast<CustomListModel721036*>(ptr)->QAbstractListModel::insertRows(row, count, *static_cast<QModelIndex*>(parent));
}

char CustomListModel721036_MoveColumnsDefault(void* ptr, void* sourceParent, int sourceColumn, int count, void* destinationParent, int destinationChild)
{
	return static_cast<CustomListModel721036*>(ptr)->QAbstractListModel::moveColumns(*static_cast<QModelIndex*>(sourceParent), sourceColumn, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
}

char CustomListModel721036_MoveRowsDefault(void* ptr, void* sourceParent, int sourceRow, int count, void* destinationParent, int destinationChild)
{
	return static_cast<CustomListModel721036*>(ptr)->QAbstractListModel::moveRows(*static_cast<QModelIndex*>(sourceParent), sourceRow, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
}

char CustomListModel721036_RemoveColumnsDefault(void* ptr, int column, int count, void* parent)
{
	return static_cast<CustomListModel721036*>(ptr)->QAbstractListModel::removeColumns(column, count, *static_cast<QModelIndex*>(parent));
}

char CustomListModel721036_RemoveRowsDefault(void* ptr, int row, int count, void* parent)
{
	return static_cast<CustomListModel721036*>(ptr)->QAbstractListModel::removeRows(row, count, *static_cast<QModelIndex*>(parent));
}

char CustomListModel721036_SetDataDefault(void* ptr, void* index, void* value, int role)
{
	return static_cast<CustomListModel721036*>(ptr)->QAbstractListModel::setData(*static_cast<QModelIndex*>(index), *static_cast<QVariant*>(value), role);
}

char CustomListModel721036_SetHeaderDataDefault(void* ptr, int section, long long orientation, void* value, int role)
{
	return static_cast<CustomListModel721036*>(ptr)->QAbstractListModel::setHeaderData(section, static_cast<Qt::Orientation>(orientation), *static_cast<QVariant*>(value), role);
}

char CustomListModel721036_SetItemDataDefault(void* ptr, void* index, void* roles)
{
	return static_cast<CustomListModel721036*>(ptr)->QAbstractListModel::setItemData(*static_cast<QModelIndex*>(index), *static_cast<QMap<int, QVariant>*>(roles));
}

char CustomListModel721036_SubmitDefault(void* ptr)
{
	return static_cast<CustomListModel721036*>(ptr)->QAbstractListModel::submit();
}

void CustomListModel721036_FetchMoreDefault(void* ptr, void* parent)
{
	static_cast<CustomListModel721036*>(ptr)->QAbstractListModel::fetchMore(*static_cast<QModelIndex*>(parent));
}

void CustomListModel721036_ResetInternalDataDefault(void* ptr)
{
	static_cast<CustomListModel721036*>(ptr)->QAbstractListModel::resetInternalData();
}

void CustomListModel721036_RevertDefault(void* ptr)
{
	static_cast<CustomListModel721036*>(ptr)->QAbstractListModel::revert();
}

void CustomListModel721036_SortDefault(void* ptr, int column, long long order)
{
	static_cast<CustomListModel721036*>(ptr)->QAbstractListModel::sort(column, static_cast<Qt::SortOrder>(order));
}

struct Moc_PackedList CustomListModel721036_RoleNamesDefault(void* ptr)
{
	return ({ QHash<int, QByteArray>* tmpValue = new QHash<int, QByteArray>(static_cast<CustomListModel721036*>(ptr)->QAbstractListModel::roleNames()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

struct Moc_PackedList CustomListModel721036_ItemDataDefault(void* ptr, void* index)
{
	return ({ QMap<int, QVariant>* tmpValue = new QMap<int, QVariant>(static_cast<CustomListModel721036*>(ptr)->QAbstractListModel::itemData(*static_cast<QModelIndex*>(index))); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* CustomListModel721036_MimeDataDefault(void* ptr, void* indexes)
{
	return static_cast<CustomListModel721036*>(ptr)->QAbstractListModel::mimeData(({ QList<QModelIndex>* tmpP = static_cast<QList<QModelIndex>*>(indexes); QList<QModelIndex> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

void* CustomListModel721036_BuddyDefault(void* ptr, void* index)
{
	return new QModelIndex(static_cast<CustomListModel721036*>(ptr)->QAbstractListModel::buddy(*static_cast<QModelIndex*>(index)));
}

void* CustomListModel721036_ParentDefault(void* ptr, void* index)
{
	return new QModelIndex(static_cast<CustomListModel721036*>(ptr)->QAbstractListModel::parent(*static_cast<QModelIndex*>(index)));
}

struct Moc_PackedList CustomListModel721036_MatchDefault(void* ptr, void* start, int role, void* value, int hits, long long flags)
{
	return ({ QList<QModelIndex>* tmpValue = new QList<QModelIndex>(static_cast<CustomListModel721036*>(ptr)->QAbstractListModel::match(*static_cast<QModelIndex*>(start), role, *static_cast<QVariant*>(value), hits, static_cast<Qt::MatchFlag>(flags))); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* CustomListModel721036_SpanDefault(void* ptr, void* index)
{
	return ({ QSize tmpValue = static_cast<CustomListModel721036*>(ptr)->QAbstractListModel::span(*static_cast<QModelIndex*>(index)); new QSize(tmpValue.width(), tmpValue.height()); });
}

struct Moc_PackedString CustomListModel721036_MimeTypesDefault(void* ptr)
{
	return ({ QByteArray t9c6749 = static_cast<CustomListModel721036*>(ptr)->QAbstractListModel::mimeTypes().join("|").toUtf8(); Moc_PackedString { const_cast<char*>(t9c6749.prepend("WHITESPACE").constData()+10), t9c6749.size()-10 }; });
}

void* CustomListModel721036_DataDefault(void* ptr, void* index, int role)
{
	Q_UNUSED(ptr);
	Q_UNUSED(index);
	Q_UNUSED(role);

}

void* CustomListModel721036_HeaderDataDefault(void* ptr, int section, long long orientation, int role)
{
	return new QVariant(static_cast<CustomListModel721036*>(ptr)->QAbstractListModel::headerData(section, static_cast<Qt::Orientation>(orientation), role));
}

long long CustomListModel721036_SupportedDragActionsDefault(void* ptr)
{
	return static_cast<CustomListModel721036*>(ptr)->QAbstractListModel::supportedDragActions();
}

long long CustomListModel721036_SupportedDropActionsDefault(void* ptr)
{
	return static_cast<CustomListModel721036*>(ptr)->QAbstractListModel::supportedDropActions();
}

char CustomListModel721036_CanDropMimeDataDefault(void* ptr, void* data, long long action, int row, int column, void* parent)
{
	return static_cast<CustomListModel721036*>(ptr)->QAbstractListModel::canDropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

char CustomListModel721036_CanFetchMoreDefault(void* ptr, void* parent)
{
	return static_cast<CustomListModel721036*>(ptr)->QAbstractListModel::canFetchMore(*static_cast<QModelIndex*>(parent));
}

char CustomListModel721036_HasChildrenDefault(void* ptr, void* parent)
{
	return static_cast<CustomListModel721036*>(ptr)->QAbstractListModel::hasChildren(*static_cast<QModelIndex*>(parent));
}

int CustomListModel721036_ColumnCountDefault(void* ptr, void* parent)
{
	return static_cast<CustomListModel721036*>(ptr)->QAbstractListModel::columnCount(*static_cast<QModelIndex*>(parent));
}

int CustomListModel721036_RowCountDefault(void* ptr, void* parent)
{
	Q_UNUSED(ptr);
	Q_UNUSED(parent);

}

char CustomListModel721036_EventDefault(void* ptr, void* e)
{
	return static_cast<CustomListModel721036*>(ptr)->QAbstractListModel::event(static_cast<QEvent*>(e));
}

char CustomListModel721036_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<CustomListModel721036*>(ptr)->QAbstractListModel::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void CustomListModel721036_ChildEventDefault(void* ptr, void* event)
{
	static_cast<CustomListModel721036*>(ptr)->QAbstractListModel::childEvent(static_cast<QChildEvent*>(event));
}

void CustomListModel721036_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<CustomListModel721036*>(ptr)->QAbstractListModel::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void CustomListModel721036_CustomEventDefault(void* ptr, void* event)
{
	static_cast<CustomListModel721036*>(ptr)->QAbstractListModel::customEvent(static_cast<QEvent*>(event));
}

void CustomListModel721036_DeleteLaterDefault(void* ptr)
{
	static_cast<CustomListModel721036*>(ptr)->QAbstractListModel::deleteLater();
}

void CustomListModel721036_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<CustomListModel721036*>(ptr)->QAbstractListModel::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void CustomListModel721036_TimerEventDefault(void* ptr, void* event)
{
	static_cast<CustomListModel721036*>(ptr)->QAbstractListModel::timerEvent(static_cast<QTimerEvent*>(event));
}

#include "moc_moc.h"
