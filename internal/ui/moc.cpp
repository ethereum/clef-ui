

#define protected public
#define private public

#include "moc.h"
#include "_cgo_export.h"

#include <QAbstractItemModel>
#include <QAbstractListModel>
#include <QByteArray>
#include <QCameraImageCapture>
#include <QChildEvent>
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
#include <QMetaObject>
#include <QMimeData>
#include <QModelIndex>
#include <QObject>
#include <QOffscreenSurface>
#include <QPaintDeviceWindow>
#include <QPdfWriter>
#include <QPersistentModelIndex>
#include <QQuickItem>
#include <QRadioData>
#include <QSize>
#include <QString>
#include <QTimerEvent>
#include <QVariant>
#include <QVector>
#include <QWidget>
#include <QWindow>


class TxListAccountsModel77da62: public QAbstractListModel
{
Q_OBJECT
public:
	TxListAccountsModel77da62(QObject *parent = Q_NULLPTR) : QAbstractListModel(parent) {qRegisterMetaType<quintptr>("quintptr");TxListAccountsModel77da62_TxListAccountsModel77da62_QRegisterMetaType();TxListAccountsModel77da62_TxListAccountsModel77da62_QRegisterMetaTypes();callbackTxListAccountsModel77da62_Constructor(this);};
	void Signal_Add(QString tx) { QByteArray t066bc1 = tx.toUtf8(); Moc_PackedString txPacked = { const_cast<char*>(t066bc1.prepend("WHITESPACE").constData()+10), t066bc1.size()-10 };callbackTxListAccountsModel77da62_Add(this, txPacked); };
	 ~TxListAccountsModel77da62() { callbackTxListAccountsModel77da62_DestroyTxListAccountsModel(this); };
	bool dropMimeData(const QMimeData * data, Qt::DropAction action, int row, int column, const QModelIndex & parent) { return callbackTxListAccountsModel77da62_DropMimeData(this, const_cast<QMimeData*>(data), action, row, column, const_cast<QModelIndex*>(&parent)) != 0; };
	QModelIndex index(int row, int column, const QModelIndex & parent) const { return *static_cast<QModelIndex*>(callbackTxListAccountsModel77da62_Index(const_cast<void*>(static_cast<const void*>(this)), row, column, const_cast<QModelIndex*>(&parent))); };
	QModelIndex sibling(int row, int column, const QModelIndex & idx) const { return *static_cast<QModelIndex*>(callbackTxListAccountsModel77da62_Sibling(const_cast<void*>(static_cast<const void*>(this)), row, column, const_cast<QModelIndex*>(&idx))); };
	Qt::ItemFlags flags(const QModelIndex & index) const { return static_cast<Qt::ItemFlag>(callbackTxListAccountsModel77da62_Flags(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	bool insertColumns(int column, int count, const QModelIndex & parent) { return callbackTxListAccountsModel77da62_InsertColumns(this, column, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool insertRows(int row, int count, const QModelIndex & parent) { return callbackTxListAccountsModel77da62_InsertRows(this, row, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool moveColumns(const QModelIndex & sourceParent, int sourceColumn, int count, const QModelIndex & destinationParent, int destinationChild) { return callbackTxListAccountsModel77da62_MoveColumns(this, const_cast<QModelIndex*>(&sourceParent), sourceColumn, count, const_cast<QModelIndex*>(&destinationParent), destinationChild) != 0; };
	bool moveRows(const QModelIndex & sourceParent, int sourceRow, int count, const QModelIndex & destinationParent, int destinationChild) { return callbackTxListAccountsModel77da62_MoveRows(this, const_cast<QModelIndex*>(&sourceParent), sourceRow, count, const_cast<QModelIndex*>(&destinationParent), destinationChild) != 0; };
	bool removeColumns(int column, int count, const QModelIndex & parent) { return callbackTxListAccountsModel77da62_RemoveColumns(this, column, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool removeRows(int row, int count, const QModelIndex & parent) { return callbackTxListAccountsModel77da62_RemoveRows(this, row, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool setData(const QModelIndex & index, const QVariant & value, int role) { return callbackTxListAccountsModel77da62_SetData(this, const_cast<QModelIndex*>(&index), const_cast<QVariant*>(&value), role) != 0; };
	bool setHeaderData(int section, Qt::Orientation orientation, const QVariant & value, int role) { return callbackTxListAccountsModel77da62_SetHeaderData(this, section, orientation, const_cast<QVariant*>(&value), role) != 0; };
	bool setItemData(const QModelIndex & index, const QMap<int, QVariant> & roles) { return callbackTxListAccountsModel77da62_SetItemData(this, const_cast<QModelIndex*>(&index), ({ QMap<int, QVariant>* tmpValue = const_cast<QMap<int, QVariant>*>(&roles); Moc_PackedList { tmpValue, tmpValue->size() }; })) != 0; };
	bool submit() { return callbackTxListAccountsModel77da62_Submit(this) != 0; };
	void Signal_ColumnsAboutToBeInserted(const QModelIndex & parent, int first, int last) { callbackTxListAccountsModel77da62_ColumnsAboutToBeInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsAboutToBeMoved(const QModelIndex & sourceParent, int sourceStart, int sourceEnd, const QModelIndex & destinationParent, int destinationColumn) { callbackTxListAccountsModel77da62_ColumnsAboutToBeMoved(this, const_cast<QModelIndex*>(&sourceParent), sourceStart, sourceEnd, const_cast<QModelIndex*>(&destinationParent), destinationColumn); };
	void Signal_ColumnsAboutToBeRemoved(const QModelIndex & parent, int first, int last) { callbackTxListAccountsModel77da62_ColumnsAboutToBeRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsInserted(const QModelIndex & parent, int first, int last) { callbackTxListAccountsModel77da62_ColumnsInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsMoved(const QModelIndex & parent, int start, int end, const QModelIndex & destination, int column) { callbackTxListAccountsModel77da62_ColumnsMoved(this, const_cast<QModelIndex*>(&parent), start, end, const_cast<QModelIndex*>(&destination), column); };
	void Signal_ColumnsRemoved(const QModelIndex & parent, int first, int last) { callbackTxListAccountsModel77da62_ColumnsRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_DataChanged(const QModelIndex & topLeft, const QModelIndex & bottomRight, const QVector<int> & roles) { callbackTxListAccountsModel77da62_DataChanged(this, const_cast<QModelIndex*>(&topLeft), const_cast<QModelIndex*>(&bottomRight), ({ QVector<int>* tmpValue = const_cast<QVector<int>*>(&roles); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	void fetchMore(const QModelIndex & parent) { callbackTxListAccountsModel77da62_FetchMore(this, const_cast<QModelIndex*>(&parent)); };
	void Signal_HeaderDataChanged(Qt::Orientation orientation, int first, int last) { callbackTxListAccountsModel77da62_HeaderDataChanged(this, orientation, first, last); };
	void Signal_LayoutAboutToBeChanged(const QList<QPersistentModelIndex> & parents, QAbstractItemModel::LayoutChangeHint hint) { callbackTxListAccountsModel77da62_LayoutAboutToBeChanged(this, ({ QList<QPersistentModelIndex>* tmpValue = const_cast<QList<QPersistentModelIndex>*>(&parents); Moc_PackedList { tmpValue, tmpValue->size() }; }), hint); };
	void Signal_LayoutChanged(const QList<QPersistentModelIndex> & parents, QAbstractItemModel::LayoutChangeHint hint) { callbackTxListAccountsModel77da62_LayoutChanged(this, ({ QList<QPersistentModelIndex>* tmpValue = const_cast<QList<QPersistentModelIndex>*>(&parents); Moc_PackedList { tmpValue, tmpValue->size() }; }), hint); };
	void Signal_ModelAboutToBeReset() { callbackTxListAccountsModel77da62_ModelAboutToBeReset(this); };
	void Signal_ModelReset() { callbackTxListAccountsModel77da62_ModelReset(this); };
	void resetInternalData() { callbackTxListAccountsModel77da62_ResetInternalData(this); };
	void revert() { callbackTxListAccountsModel77da62_Revert(this); };
	void Signal_RowsAboutToBeInserted(const QModelIndex & parent, int start, int end) { callbackTxListAccountsModel77da62_RowsAboutToBeInserted(this, const_cast<QModelIndex*>(&parent), start, end); };
	void Signal_RowsAboutToBeMoved(const QModelIndex & sourceParent, int sourceStart, int sourceEnd, const QModelIndex & destinationParent, int destinationRow) { callbackTxListAccountsModel77da62_RowsAboutToBeMoved(this, const_cast<QModelIndex*>(&sourceParent), sourceStart, sourceEnd, const_cast<QModelIndex*>(&destinationParent), destinationRow); };
	void Signal_RowsAboutToBeRemoved(const QModelIndex & parent, int first, int last) { callbackTxListAccountsModel77da62_RowsAboutToBeRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_RowsInserted(const QModelIndex & parent, int first, int last) { callbackTxListAccountsModel77da62_RowsInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_RowsMoved(const QModelIndex & parent, int start, int end, const QModelIndex & destination, int row) { callbackTxListAccountsModel77da62_RowsMoved(this, const_cast<QModelIndex*>(&parent), start, end, const_cast<QModelIndex*>(&destination), row); };
	void Signal_RowsRemoved(const QModelIndex & parent, int first, int last) { callbackTxListAccountsModel77da62_RowsRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void sort(int column, Qt::SortOrder order) { callbackTxListAccountsModel77da62_Sort(this, column, order); };
	QHash<int, QByteArray> roleNames() const { return ({ QHash<int, QByteArray>* tmpP = static_cast<QHash<int, QByteArray>*>(callbackTxListAccountsModel77da62_RoleNames(const_cast<void*>(static_cast<const void*>(this)))); QHash<int, QByteArray> tmpV = *tmpP; tmpP->~QHash(); free(tmpP); tmpV; }); };
	QMap<int, QVariant> itemData(const QModelIndex & index) const { return ({ QMap<int, QVariant>* tmpP = static_cast<QMap<int, QVariant>*>(callbackTxListAccountsModel77da62_ItemData(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); QMap<int, QVariant> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }); };
	QMimeData * mimeData(const QModelIndexList & indexes) const { return static_cast<QMimeData*>(callbackTxListAccountsModel77da62_MimeData(const_cast<void*>(static_cast<const void*>(this)), ({ QList<QModelIndex>* tmpValue = new QList<QModelIndex>(indexes); Moc_PackedList { tmpValue, tmpValue->size() }; }))); };
	QModelIndex buddy(const QModelIndex & index) const { return *static_cast<QModelIndex*>(callbackTxListAccountsModel77da62_Buddy(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	QModelIndex parent(const QModelIndex & index) const { return *static_cast<QModelIndex*>(callbackTxListAccountsModel77da62_Parent(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	QList<QModelIndex> match(const QModelIndex & start, int role, const QVariant & value, int hits, Qt::MatchFlags flags) const { return ({ QList<QModelIndex>* tmpP = static_cast<QList<QModelIndex>*>(callbackTxListAccountsModel77da62_Match(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&start), role, const_cast<QVariant*>(&value), hits, flags)); QList<QModelIndex> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }); };
	QSize span(const QModelIndex & index) const { return *static_cast<QSize*>(callbackTxListAccountsModel77da62_Span(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	QStringList mimeTypes() const { return ({ Moc_PackedString tempVal = callbackTxListAccountsModel77da62_MimeTypes(const_cast<void*>(static_cast<const void*>(this))); QStringList ret = QString::fromUtf8(tempVal.data, tempVal.len).split("|", QString::SkipEmptyParts); free(tempVal.data); ret; }); };
	QVariant data(const QModelIndex & index, int role) const { return *static_cast<QVariant*>(callbackTxListAccountsModel77da62_Data(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index), role)); };
	QVariant headerData(int section, Qt::Orientation orientation, int role) const { return *static_cast<QVariant*>(callbackTxListAccountsModel77da62_HeaderData(const_cast<void*>(static_cast<const void*>(this)), section, orientation, role)); };
	Qt::DropActions supportedDragActions() const { return static_cast<Qt::DropAction>(callbackTxListAccountsModel77da62_SupportedDragActions(const_cast<void*>(static_cast<const void*>(this)))); };
	Qt::DropActions supportedDropActions() const { return static_cast<Qt::DropAction>(callbackTxListAccountsModel77da62_SupportedDropActions(const_cast<void*>(static_cast<const void*>(this)))); };
	bool canDropMimeData(const QMimeData * data, Qt::DropAction action, int row, int column, const QModelIndex & parent) const { return callbackTxListAccountsModel77da62_CanDropMimeData(const_cast<void*>(static_cast<const void*>(this)), const_cast<QMimeData*>(data), action, row, column, const_cast<QModelIndex*>(&parent)) != 0; };
	bool canFetchMore(const QModelIndex & parent) const { return callbackTxListAccountsModel77da62_CanFetchMore(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)) != 0; };
	bool hasChildren(const QModelIndex & parent) const { return callbackTxListAccountsModel77da62_HasChildren(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)) != 0; };
	int columnCount(const QModelIndex & parent) const { return callbackTxListAccountsModel77da62_ColumnCount(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)); };
	int rowCount(const QModelIndex & parent) const { return callbackTxListAccountsModel77da62_RowCount(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)); };
	bool event(QEvent * e) { return callbackTxListAccountsModel77da62_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackTxListAccountsModel77da62_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackTxListAccountsModel77da62_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackTxListAccountsModel77da62_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackTxListAccountsModel77da62_CustomEvent(this, event); };
	void deleteLater() { callbackTxListAccountsModel77da62_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackTxListAccountsModel77da62_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackTxListAccountsModel77da62_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackTxListAccountsModel77da62_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackTxListAccountsModel77da62_TimerEvent(this, event); };
signals:
	void add(QString tx);
public slots:
private:
};

Q_DECLARE_METATYPE(TxListAccountsModel77da62*)


void TxListAccountsModel77da62_TxListAccountsModel77da62_QRegisterMetaTypes() {
}

class TxListModel77da62: public QAbstractListModel
{
Q_OBJECT
Q_PROPERTY(bool isEmpty READ isEmpty WRITE setIsEmpty NOTIFY isEmptyChanged)
public:
	TxListModel77da62(QObject *parent = Q_NULLPTR) : QAbstractListModel(parent) {qRegisterMetaType<quintptr>("quintptr");TxListModel77da62_TxListModel77da62_QRegisterMetaType();TxListModel77da62_TxListModel77da62_QRegisterMetaTypes();callbackTxListModel77da62_Constructor(this);};
	void Signal_Clear() { callbackTxListModel77da62_Clear(this); };
	void Signal_Add(quintptr tx) { callbackTxListModel77da62_Add(this, tx); };
	void Signal_Remove(qint32 i) { callbackTxListModel77da62_Remove(this, i); };
	bool isEmpty() { return callbackTxListModel77da62_IsEmpty(this) != 0; };
	void setIsEmpty(bool isEmpty) { callbackTxListModel77da62_SetIsEmpty(this, isEmpty); };
	void Signal_IsEmptyChanged(bool isEmpty) { callbackTxListModel77da62_IsEmptyChanged(this, isEmpty); };
	 ~TxListModel77da62() { callbackTxListModel77da62_DestroyTxListModel(this); };
	bool dropMimeData(const QMimeData * data, Qt::DropAction action, int row, int column, const QModelIndex & parent) { return callbackTxListModel77da62_DropMimeData(this, const_cast<QMimeData*>(data), action, row, column, const_cast<QModelIndex*>(&parent)) != 0; };
	QModelIndex index(int row, int column, const QModelIndex & parent) const { return *static_cast<QModelIndex*>(callbackTxListModel77da62_Index(const_cast<void*>(static_cast<const void*>(this)), row, column, const_cast<QModelIndex*>(&parent))); };
	QModelIndex sibling(int row, int column, const QModelIndex & idx) const { return *static_cast<QModelIndex*>(callbackTxListModel77da62_Sibling(const_cast<void*>(static_cast<const void*>(this)), row, column, const_cast<QModelIndex*>(&idx))); };
	Qt::ItemFlags flags(const QModelIndex & index) const { return static_cast<Qt::ItemFlag>(callbackTxListModel77da62_Flags(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	bool insertColumns(int column, int count, const QModelIndex & parent) { return callbackTxListModel77da62_InsertColumns(this, column, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool insertRows(int row, int count, const QModelIndex & parent) { return callbackTxListModel77da62_InsertRows(this, row, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool moveColumns(const QModelIndex & sourceParent, int sourceColumn, int count, const QModelIndex & destinationParent, int destinationChild) { return callbackTxListModel77da62_MoveColumns(this, const_cast<QModelIndex*>(&sourceParent), sourceColumn, count, const_cast<QModelIndex*>(&destinationParent), destinationChild) != 0; };
	bool moveRows(const QModelIndex & sourceParent, int sourceRow, int count, const QModelIndex & destinationParent, int destinationChild) { return callbackTxListModel77da62_MoveRows(this, const_cast<QModelIndex*>(&sourceParent), sourceRow, count, const_cast<QModelIndex*>(&destinationParent), destinationChild) != 0; };
	bool removeColumns(int column, int count, const QModelIndex & parent) { return callbackTxListModel77da62_RemoveColumns(this, column, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool removeRows(int row, int count, const QModelIndex & parent) { return callbackTxListModel77da62_RemoveRows(this, row, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool setData(const QModelIndex & index, const QVariant & value, int role) { return callbackTxListModel77da62_SetData(this, const_cast<QModelIndex*>(&index), const_cast<QVariant*>(&value), role) != 0; };
	bool setHeaderData(int section, Qt::Orientation orientation, const QVariant & value, int role) { return callbackTxListModel77da62_SetHeaderData(this, section, orientation, const_cast<QVariant*>(&value), role) != 0; };
	bool setItemData(const QModelIndex & index, const QMap<int, QVariant> & roles) { return callbackTxListModel77da62_SetItemData(this, const_cast<QModelIndex*>(&index), ({ QMap<int, QVariant>* tmpValue = const_cast<QMap<int, QVariant>*>(&roles); Moc_PackedList { tmpValue, tmpValue->size() }; })) != 0; };
	bool submit() { return callbackTxListModel77da62_Submit(this) != 0; };
	void Signal_ColumnsAboutToBeInserted(const QModelIndex & parent, int first, int last) { callbackTxListModel77da62_ColumnsAboutToBeInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsAboutToBeMoved(const QModelIndex & sourceParent, int sourceStart, int sourceEnd, const QModelIndex & destinationParent, int destinationColumn) { callbackTxListModel77da62_ColumnsAboutToBeMoved(this, const_cast<QModelIndex*>(&sourceParent), sourceStart, sourceEnd, const_cast<QModelIndex*>(&destinationParent), destinationColumn); };
	void Signal_ColumnsAboutToBeRemoved(const QModelIndex & parent, int first, int last) { callbackTxListModel77da62_ColumnsAboutToBeRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsInserted(const QModelIndex & parent, int first, int last) { callbackTxListModel77da62_ColumnsInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsMoved(const QModelIndex & parent, int start, int end, const QModelIndex & destination, int column) { callbackTxListModel77da62_ColumnsMoved(this, const_cast<QModelIndex*>(&parent), start, end, const_cast<QModelIndex*>(&destination), column); };
	void Signal_ColumnsRemoved(const QModelIndex & parent, int first, int last) { callbackTxListModel77da62_ColumnsRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_DataChanged(const QModelIndex & topLeft, const QModelIndex & bottomRight, const QVector<int> & roles) { callbackTxListModel77da62_DataChanged(this, const_cast<QModelIndex*>(&topLeft), const_cast<QModelIndex*>(&bottomRight), ({ QVector<int>* tmpValue = const_cast<QVector<int>*>(&roles); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	void fetchMore(const QModelIndex & parent) { callbackTxListModel77da62_FetchMore(this, const_cast<QModelIndex*>(&parent)); };
	void Signal_HeaderDataChanged(Qt::Orientation orientation, int first, int last) { callbackTxListModel77da62_HeaderDataChanged(this, orientation, first, last); };
	void Signal_LayoutAboutToBeChanged(const QList<QPersistentModelIndex> & parents, QAbstractItemModel::LayoutChangeHint hint) { callbackTxListModel77da62_LayoutAboutToBeChanged(this, ({ QList<QPersistentModelIndex>* tmpValue = const_cast<QList<QPersistentModelIndex>*>(&parents); Moc_PackedList { tmpValue, tmpValue->size() }; }), hint); };
	void Signal_LayoutChanged(const QList<QPersistentModelIndex> & parents, QAbstractItemModel::LayoutChangeHint hint) { callbackTxListModel77da62_LayoutChanged(this, ({ QList<QPersistentModelIndex>* tmpValue = const_cast<QList<QPersistentModelIndex>*>(&parents); Moc_PackedList { tmpValue, tmpValue->size() }; }), hint); };
	void Signal_ModelAboutToBeReset() { callbackTxListModel77da62_ModelAboutToBeReset(this); };
	void Signal_ModelReset() { callbackTxListModel77da62_ModelReset(this); };
	void resetInternalData() { callbackTxListModel77da62_ResetInternalData(this); };
	void revert() { callbackTxListModel77da62_Revert(this); };
	void Signal_RowsAboutToBeInserted(const QModelIndex & parent, int start, int end) { callbackTxListModel77da62_RowsAboutToBeInserted(this, const_cast<QModelIndex*>(&parent), start, end); };
	void Signal_RowsAboutToBeMoved(const QModelIndex & sourceParent, int sourceStart, int sourceEnd, const QModelIndex & destinationParent, int destinationRow) { callbackTxListModel77da62_RowsAboutToBeMoved(this, const_cast<QModelIndex*>(&sourceParent), sourceStart, sourceEnd, const_cast<QModelIndex*>(&destinationParent), destinationRow); };
	void Signal_RowsAboutToBeRemoved(const QModelIndex & parent, int first, int last) { callbackTxListModel77da62_RowsAboutToBeRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_RowsInserted(const QModelIndex & parent, int first, int last) { callbackTxListModel77da62_RowsInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_RowsMoved(const QModelIndex & parent, int start, int end, const QModelIndex & destination, int row) { callbackTxListModel77da62_RowsMoved(this, const_cast<QModelIndex*>(&parent), start, end, const_cast<QModelIndex*>(&destination), row); };
	void Signal_RowsRemoved(const QModelIndex & parent, int first, int last) { callbackTxListModel77da62_RowsRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void sort(int column, Qt::SortOrder order) { callbackTxListModel77da62_Sort(this, column, order); };
	QHash<int, QByteArray> roleNames() const { return ({ QHash<int, QByteArray>* tmpP = static_cast<QHash<int, QByteArray>*>(callbackTxListModel77da62_RoleNames(const_cast<void*>(static_cast<const void*>(this)))); QHash<int, QByteArray> tmpV = *tmpP; tmpP->~QHash(); free(tmpP); tmpV; }); };
	QMap<int, QVariant> itemData(const QModelIndex & index) const { return ({ QMap<int, QVariant>* tmpP = static_cast<QMap<int, QVariant>*>(callbackTxListModel77da62_ItemData(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); QMap<int, QVariant> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }); };
	QMimeData * mimeData(const QModelIndexList & indexes) const { return static_cast<QMimeData*>(callbackTxListModel77da62_MimeData(const_cast<void*>(static_cast<const void*>(this)), ({ QList<QModelIndex>* tmpValue = new QList<QModelIndex>(indexes); Moc_PackedList { tmpValue, tmpValue->size() }; }))); };
	QModelIndex buddy(const QModelIndex & index) const { return *static_cast<QModelIndex*>(callbackTxListModel77da62_Buddy(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	QModelIndex parent(const QModelIndex & index) const { return *static_cast<QModelIndex*>(callbackTxListModel77da62_Parent(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	QList<QModelIndex> match(const QModelIndex & start, int role, const QVariant & value, int hits, Qt::MatchFlags flags) const { return ({ QList<QModelIndex>* tmpP = static_cast<QList<QModelIndex>*>(callbackTxListModel77da62_Match(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&start), role, const_cast<QVariant*>(&value), hits, flags)); QList<QModelIndex> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }); };
	QSize span(const QModelIndex & index) const { return *static_cast<QSize*>(callbackTxListModel77da62_Span(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	QStringList mimeTypes() const { return ({ Moc_PackedString tempVal = callbackTxListModel77da62_MimeTypes(const_cast<void*>(static_cast<const void*>(this))); QStringList ret = QString::fromUtf8(tempVal.data, tempVal.len).split("|", QString::SkipEmptyParts); free(tempVal.data); ret; }); };
	QVariant data(const QModelIndex & index, int role) const { return *static_cast<QVariant*>(callbackTxListModel77da62_Data(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index), role)); };
	QVariant headerData(int section, Qt::Orientation orientation, int role) const { return *static_cast<QVariant*>(callbackTxListModel77da62_HeaderData(const_cast<void*>(static_cast<const void*>(this)), section, orientation, role)); };
	Qt::DropActions supportedDragActions() const { return static_cast<Qt::DropAction>(callbackTxListModel77da62_SupportedDragActions(const_cast<void*>(static_cast<const void*>(this)))); };
	Qt::DropActions supportedDropActions() const { return static_cast<Qt::DropAction>(callbackTxListModel77da62_SupportedDropActions(const_cast<void*>(static_cast<const void*>(this)))); };
	bool canDropMimeData(const QMimeData * data, Qt::DropAction action, int row, int column, const QModelIndex & parent) const { return callbackTxListModel77da62_CanDropMimeData(const_cast<void*>(static_cast<const void*>(this)), const_cast<QMimeData*>(data), action, row, column, const_cast<QModelIndex*>(&parent)) != 0; };
	bool canFetchMore(const QModelIndex & parent) const { return callbackTxListModel77da62_CanFetchMore(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)) != 0; };
	bool hasChildren(const QModelIndex & parent) const { return callbackTxListModel77da62_HasChildren(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)) != 0; };
	int columnCount(const QModelIndex & parent) const { return callbackTxListModel77da62_ColumnCount(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)); };
	int rowCount(const QModelIndex & parent) const { return callbackTxListModel77da62_RowCount(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)); };
	bool event(QEvent * e) { return callbackTxListModel77da62_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackTxListModel77da62_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackTxListModel77da62_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackTxListModel77da62_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackTxListModel77da62_CustomEvent(this, event); };
	void deleteLater() { callbackTxListModel77da62_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackTxListModel77da62_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackTxListModel77da62_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackTxListModel77da62_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackTxListModel77da62_TimerEvent(this, event); };
	bool isEmptyDefault() { return _isEmpty; };
	void setIsEmptyDefault(bool p) { if (p != _isEmpty) { _isEmpty = p; isEmptyChanged(_isEmpty); } };
signals:
	void clear();
	void add(quintptr tx);
	void remove(qint32 i);
	void isEmptyChanged(bool isEmpty);
public slots:
private:
	bool _isEmpty;
};

Q_DECLARE_METATYPE(TxListModel77da62*)


void TxListModel77da62_TxListModel77da62_QRegisterMetaTypes() {
}

class ApproveListingCtx77da62: public QObject
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
	ApproveListingCtx77da62(QObject *parent = Q_NULLPTR) : QObject(parent) {qRegisterMetaType<quintptr>("quintptr");ApproveListingCtx77da62_ApproveListingCtx77da62_QRegisterMetaType();ApproveListingCtx77da62_ApproveListingCtx77da62_QRegisterMetaTypes();callbackApproveListingCtx77da62_Constructor(this);};
	void Signal_Back() { callbackApproveListingCtx77da62_Back(this); };
	void Signal_Approve() { callbackApproveListingCtx77da62_Approve(this); };
	void Signal_Reject() { callbackApproveListingCtx77da62_Reject(this); };
	void Signal_OnCheckStateChanged(qint32 i, bool checked) { callbackApproveListingCtx77da62_OnCheckStateChanged(this, i, checked); };
	void Signal_TriggerUpdate() { callbackApproveListingCtx77da62_TriggerUpdate(this); };
	QString remote() { return ({ Moc_PackedString tempVal = callbackApproveListingCtx77da62_Remote(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setRemote(QString remote) { QByteArray t41ffe5 = remote.toUtf8(); Moc_PackedString remotePacked = { const_cast<char*>(t41ffe5.prepend("WHITESPACE").constData()+10), t41ffe5.size()-10 };callbackApproveListingCtx77da62_SetRemote(this, remotePacked); };
	void Signal_RemoteChanged(QString remote) { QByteArray t41ffe5 = remote.toUtf8(); Moc_PackedString remotePacked = { const_cast<char*>(t41ffe5.prepend("WHITESPACE").constData()+10), t41ffe5.size()-10 };callbackApproveListingCtx77da62_RemoteChanged(this, remotePacked); };
	QString transport() { return ({ Moc_PackedString tempVal = callbackApproveListingCtx77da62_Transport(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setTransport(QString transport) { QByteArray ta8e601 = transport.toUtf8(); Moc_PackedString transportPacked = { const_cast<char*>(ta8e601.prepend("WHITESPACE").constData()+10), ta8e601.size()-10 };callbackApproveListingCtx77da62_SetTransport(this, transportPacked); };
	void Signal_TransportChanged(QString transport) { QByteArray ta8e601 = transport.toUtf8(); Moc_PackedString transportPacked = { const_cast<char*>(ta8e601.prepend("WHITESPACE").constData()+10), ta8e601.size()-10 };callbackApproveListingCtx77da62_TransportChanged(this, transportPacked); };
	QString endpoint() { return ({ Moc_PackedString tempVal = callbackApproveListingCtx77da62_Endpoint(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setEndpoint(QString endpoint) { QByteArray te13fe4 = endpoint.toUtf8(); Moc_PackedString endpointPacked = { const_cast<char*>(te13fe4.prepend("WHITESPACE").constData()+10), te13fe4.size()-10 };callbackApproveListingCtx77da62_SetEndpoint(this, endpointPacked); };
	void Signal_EndpointChanged(QString endpoint) { QByteArray te13fe4 = endpoint.toUtf8(); Moc_PackedString endpointPacked = { const_cast<char*>(te13fe4.prepend("WHITESPACE").constData()+10), te13fe4.size()-10 };callbackApproveListingCtx77da62_EndpointChanged(this, endpointPacked); };
	QString from() { return ({ Moc_PackedString tempVal = callbackApproveListingCtx77da62_From(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setFrom(QString from) { QByteArray t0b1e95 = from.toUtf8(); Moc_PackedString fromPacked = { const_cast<char*>(t0b1e95.prepend("WHITESPACE").constData()+10), t0b1e95.size()-10 };callbackApproveListingCtx77da62_SetFrom(this, fromPacked); };
	void Signal_FromChanged(QString from) { QByteArray t0b1e95 = from.toUtf8(); Moc_PackedString fromPacked = { const_cast<char*>(t0b1e95.prepend("WHITESPACE").constData()+10), t0b1e95.size()-10 };callbackApproveListingCtx77da62_FromChanged(this, fromPacked); };
	QString message() { return ({ Moc_PackedString tempVal = callbackApproveListingCtx77da62_Message(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setMessage(QString message) { QByteArray t6f9b9a = message.toUtf8(); Moc_PackedString messagePacked = { const_cast<char*>(t6f9b9a.prepend("WHITESPACE").constData()+10), t6f9b9a.size()-10 };callbackApproveListingCtx77da62_SetMessage(this, messagePacked); };
	void Signal_MessageChanged(QString message) { QByteArray t6f9b9a = message.toUtf8(); Moc_PackedString messagePacked = { const_cast<char*>(t6f9b9a.prepend("WHITESPACE").constData()+10), t6f9b9a.size()-10 };callbackApproveListingCtx77da62_MessageChanged(this, messagePacked); };
	QString rawData() { return ({ Moc_PackedString tempVal = callbackApproveListingCtx77da62_RawData(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setRawData(QString rawData) { QByteArray tcacc10 = rawData.toUtf8(); Moc_PackedString rawDataPacked = { const_cast<char*>(tcacc10.prepend("WHITESPACE").constData()+10), tcacc10.size()-10 };callbackApproveListingCtx77da62_SetRawData(this, rawDataPacked); };
	void Signal_RawDataChanged(QString rawData) { QByteArray tcacc10 = rawData.toUtf8(); Moc_PackedString rawDataPacked = { const_cast<char*>(tcacc10.prepend("WHITESPACE").constData()+10), tcacc10.size()-10 };callbackApproveListingCtx77da62_RawDataChanged(this, rawDataPacked); };
	QString hash() { return ({ Moc_PackedString tempVal = callbackApproveListingCtx77da62_Hash(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setHash(QString hash) { QByteArray t2346ad = hash.toUtf8(); Moc_PackedString hashPacked = { const_cast<char*>(t2346ad.prepend("WHITESPACE").constData()+10), t2346ad.size()-10 };callbackApproveListingCtx77da62_SetHash(this, hashPacked); };
	void Signal_HashChanged(QString hash) { QByteArray t2346ad = hash.toUtf8(); Moc_PackedString hashPacked = { const_cast<char*>(t2346ad.prepend("WHITESPACE").constData()+10), t2346ad.size()-10 };callbackApproveListingCtx77da62_HashChanged(this, hashPacked); };
	 ~ApproveListingCtx77da62() { callbackApproveListingCtx77da62_DestroyApproveListingCtx(this); };
	bool event(QEvent * e) { return callbackApproveListingCtx77da62_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackApproveListingCtx77da62_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackApproveListingCtx77da62_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackApproveListingCtx77da62_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackApproveListingCtx77da62_CustomEvent(this, event); };
	void deleteLater() { callbackApproveListingCtx77da62_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackApproveListingCtx77da62_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackApproveListingCtx77da62_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackApproveListingCtx77da62_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackApproveListingCtx77da62_TimerEvent(this, event); };
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
	void back();
	void approve();
	void reject();
	void onCheckStateChanged(qint32 i, bool checked);
	void triggerUpdate();
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

Q_DECLARE_METATYPE(ApproveListingCtx77da62*)


void ApproveListingCtx77da62_ApproveListingCtx77da62_QRegisterMetaTypes() {
	qRegisterMetaType<QString>();
}

class ApproveSignDataCtx77da62: public QObject
{
Q_OBJECT
Q_PROPERTY(QString remote READ remote WRITE setRemote NOTIFY remoteChanged)
Q_PROPERTY(QString transport READ transport WRITE setTransport NOTIFY transportChanged)
Q_PROPERTY(QString endpoint READ endpoint WRITE setEndpoint NOTIFY endpointChanged)
Q_PROPERTY(QString from READ from WRITE setFrom NOTIFY fromChanged)
Q_PROPERTY(QString message READ message WRITE setMessage NOTIFY messageChanged)
Q_PROPERTY(QString rawData READ rawData WRITE setRawData NOTIFY rawDataChanged)
Q_PROPERTY(QString hash READ hash WRITE setHash NOTIFY hashChanged)
Q_PROPERTY(QString password READ password WRITE setPassword NOTIFY passwordChanged)
Q_PROPERTY(QString fromSrc READ fromSrc WRITE setFromSrc NOTIFY fromSrcChanged)
public:
	ApproveSignDataCtx77da62(QObject *parent = Q_NULLPTR) : QObject(parent) {qRegisterMetaType<quintptr>("quintptr");ApproveSignDataCtx77da62_ApproveSignDataCtx77da62_QRegisterMetaType();ApproveSignDataCtx77da62_ApproveSignDataCtx77da62_QRegisterMetaTypes();callbackApproveSignDataCtx77da62_Constructor(this);};
	void Signal_Clicked(qint32 b) { callbackApproveSignDataCtx77da62_Clicked(this, b); };
	void Signal_OnBack() { callbackApproveSignDataCtx77da62_OnBack(this); };
	void Signal_OnApprove() { callbackApproveSignDataCtx77da62_OnApprove(this); };
	void Signal_OnReject() { callbackApproveSignDataCtx77da62_OnReject(this); };
	void Signal_Edited(QString b, QString value) { QByteArray te9d71f = b.toUtf8(); Moc_PackedString bPacked = { const_cast<char*>(te9d71f.prepend("WHITESPACE").constData()+10), te9d71f.size()-10 };QByteArray tf32b67 = value.toUtf8(); Moc_PackedString valuePacked = { const_cast<char*>(tf32b67.prepend("WHITESPACE").constData()+10), tf32b67.size()-10 };callbackApproveSignDataCtx77da62_Edited(this, bPacked, valuePacked); };
	QString remote() { return ({ Moc_PackedString tempVal = callbackApproveSignDataCtx77da62_Remote(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setRemote(QString remote) { QByteArray t41ffe5 = remote.toUtf8(); Moc_PackedString remotePacked = { const_cast<char*>(t41ffe5.prepend("WHITESPACE").constData()+10), t41ffe5.size()-10 };callbackApproveSignDataCtx77da62_SetRemote(this, remotePacked); };
	void Signal_RemoteChanged(QString remote) { QByteArray t41ffe5 = remote.toUtf8(); Moc_PackedString remotePacked = { const_cast<char*>(t41ffe5.prepend("WHITESPACE").constData()+10), t41ffe5.size()-10 };callbackApproveSignDataCtx77da62_RemoteChanged(this, remotePacked); };
	QString transport() { return ({ Moc_PackedString tempVal = callbackApproveSignDataCtx77da62_Transport(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setTransport(QString transport) { QByteArray ta8e601 = transport.toUtf8(); Moc_PackedString transportPacked = { const_cast<char*>(ta8e601.prepend("WHITESPACE").constData()+10), ta8e601.size()-10 };callbackApproveSignDataCtx77da62_SetTransport(this, transportPacked); };
	void Signal_TransportChanged(QString transport) { QByteArray ta8e601 = transport.toUtf8(); Moc_PackedString transportPacked = { const_cast<char*>(ta8e601.prepend("WHITESPACE").constData()+10), ta8e601.size()-10 };callbackApproveSignDataCtx77da62_TransportChanged(this, transportPacked); };
	QString endpoint() { return ({ Moc_PackedString tempVal = callbackApproveSignDataCtx77da62_Endpoint(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setEndpoint(QString endpoint) { QByteArray te13fe4 = endpoint.toUtf8(); Moc_PackedString endpointPacked = { const_cast<char*>(te13fe4.prepend("WHITESPACE").constData()+10), te13fe4.size()-10 };callbackApproveSignDataCtx77da62_SetEndpoint(this, endpointPacked); };
	void Signal_EndpointChanged(QString endpoint) { QByteArray te13fe4 = endpoint.toUtf8(); Moc_PackedString endpointPacked = { const_cast<char*>(te13fe4.prepend("WHITESPACE").constData()+10), te13fe4.size()-10 };callbackApproveSignDataCtx77da62_EndpointChanged(this, endpointPacked); };
	QString from() { return ({ Moc_PackedString tempVal = callbackApproveSignDataCtx77da62_From(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setFrom(QString from) { QByteArray t0b1e95 = from.toUtf8(); Moc_PackedString fromPacked = { const_cast<char*>(t0b1e95.prepend("WHITESPACE").constData()+10), t0b1e95.size()-10 };callbackApproveSignDataCtx77da62_SetFrom(this, fromPacked); };
	void Signal_FromChanged(QString from) { QByteArray t0b1e95 = from.toUtf8(); Moc_PackedString fromPacked = { const_cast<char*>(t0b1e95.prepend("WHITESPACE").constData()+10), t0b1e95.size()-10 };callbackApproveSignDataCtx77da62_FromChanged(this, fromPacked); };
	QString message() { return ({ Moc_PackedString tempVal = callbackApproveSignDataCtx77da62_Message(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setMessage(QString message) { QByteArray t6f9b9a = message.toUtf8(); Moc_PackedString messagePacked = { const_cast<char*>(t6f9b9a.prepend("WHITESPACE").constData()+10), t6f9b9a.size()-10 };callbackApproveSignDataCtx77da62_SetMessage(this, messagePacked); };
	void Signal_MessageChanged(QString message) { QByteArray t6f9b9a = message.toUtf8(); Moc_PackedString messagePacked = { const_cast<char*>(t6f9b9a.prepend("WHITESPACE").constData()+10), t6f9b9a.size()-10 };callbackApproveSignDataCtx77da62_MessageChanged(this, messagePacked); };
	QString rawData() { return ({ Moc_PackedString tempVal = callbackApproveSignDataCtx77da62_RawData(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setRawData(QString rawData) { QByteArray tcacc10 = rawData.toUtf8(); Moc_PackedString rawDataPacked = { const_cast<char*>(tcacc10.prepend("WHITESPACE").constData()+10), tcacc10.size()-10 };callbackApproveSignDataCtx77da62_SetRawData(this, rawDataPacked); };
	void Signal_RawDataChanged(QString rawData) { QByteArray tcacc10 = rawData.toUtf8(); Moc_PackedString rawDataPacked = { const_cast<char*>(tcacc10.prepend("WHITESPACE").constData()+10), tcacc10.size()-10 };callbackApproveSignDataCtx77da62_RawDataChanged(this, rawDataPacked); };
	QString hash() { return ({ Moc_PackedString tempVal = callbackApproveSignDataCtx77da62_Hash(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setHash(QString hash) { QByteArray t2346ad = hash.toUtf8(); Moc_PackedString hashPacked = { const_cast<char*>(t2346ad.prepend("WHITESPACE").constData()+10), t2346ad.size()-10 };callbackApproveSignDataCtx77da62_SetHash(this, hashPacked); };
	void Signal_HashChanged(QString hash) { QByteArray t2346ad = hash.toUtf8(); Moc_PackedString hashPacked = { const_cast<char*>(t2346ad.prepend("WHITESPACE").constData()+10), t2346ad.size()-10 };callbackApproveSignDataCtx77da62_HashChanged(this, hashPacked); };
	QString password() { return ({ Moc_PackedString tempVal = callbackApproveSignDataCtx77da62_Password(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setPassword(QString password) { QByteArray t5baa61 = password.toUtf8(); Moc_PackedString passwordPacked = { const_cast<char*>(t5baa61.prepend("WHITESPACE").constData()+10), t5baa61.size()-10 };callbackApproveSignDataCtx77da62_SetPassword(this, passwordPacked); };
	void Signal_PasswordChanged(QString password) { QByteArray t5baa61 = password.toUtf8(); Moc_PackedString passwordPacked = { const_cast<char*>(t5baa61.prepend("WHITESPACE").constData()+10), t5baa61.size()-10 };callbackApproveSignDataCtx77da62_PasswordChanged(this, passwordPacked); };
	QString fromSrc() { return ({ Moc_PackedString tempVal = callbackApproveSignDataCtx77da62_FromSrc(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setFromSrc(QString fromSrc) { QByteArray ta8ded4 = fromSrc.toUtf8(); Moc_PackedString fromSrcPacked = { const_cast<char*>(ta8ded4.prepend("WHITESPACE").constData()+10), ta8ded4.size()-10 };callbackApproveSignDataCtx77da62_SetFromSrc(this, fromSrcPacked); };
	void Signal_FromSrcChanged(QString fromSrc) { QByteArray ta8ded4 = fromSrc.toUtf8(); Moc_PackedString fromSrcPacked = { const_cast<char*>(ta8ded4.prepend("WHITESPACE").constData()+10), ta8ded4.size()-10 };callbackApproveSignDataCtx77da62_FromSrcChanged(this, fromSrcPacked); };
	 ~ApproveSignDataCtx77da62() { callbackApproveSignDataCtx77da62_DestroyApproveSignDataCtx(this); };
	bool event(QEvent * e) { return callbackApproveSignDataCtx77da62_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackApproveSignDataCtx77da62_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackApproveSignDataCtx77da62_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackApproveSignDataCtx77da62_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackApproveSignDataCtx77da62_CustomEvent(this, event); };
	void deleteLater() { callbackApproveSignDataCtx77da62_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackApproveSignDataCtx77da62_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackApproveSignDataCtx77da62_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackApproveSignDataCtx77da62_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackApproveSignDataCtx77da62_TimerEvent(this, event); };
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
	QString passwordDefault() { return _password; };
	void setPasswordDefault(QString p) { if (p != _password) { _password = p; passwordChanged(_password); } };
	QString fromSrcDefault() { return _fromSrc; };
	void setFromSrcDefault(QString p) { if (p != _fromSrc) { _fromSrc = p; fromSrcChanged(_fromSrc); } };
signals:
	void clicked(qint32 b);
	void onBack();
	void onApprove();
	void onReject();
	void edited(QString b, QString value);
	void remoteChanged(QString remote);
	void transportChanged(QString transport);
	void endpointChanged(QString endpoint);
	void fromChanged(QString from);
	void messageChanged(QString message);
	void rawDataChanged(QString rawData);
	void hashChanged(QString hash);
	void passwordChanged(QString password);
	void fromSrcChanged(QString fromSrc);
public slots:
private:
	QString _remote;
	QString _transport;
	QString _endpoint;
	QString _from;
	QString _message;
	QString _rawData;
	QString _hash;
	QString _password;
	QString _fromSrc;
};

Q_DECLARE_METATYPE(ApproveSignDataCtx77da62*)


void ApproveSignDataCtx77da62_ApproveSignDataCtx77da62_QRegisterMetaTypes() {
	qRegisterMetaType<QString>();
}

class CustomListModel77da62: public QAbstractListModel
{
Q_OBJECT
Q_PROPERTY(QString updateRequired READ updateRequired WRITE setUpdateRequired NOTIFY updateRequiredChanged)
public:
	CustomListModel77da62(QObject *parent = Q_NULLPTR) : QAbstractListModel(parent) {qRegisterMetaType<quintptr>("quintptr");CustomListModel77da62_CustomListModel77da62_QRegisterMetaType();CustomListModel77da62_CustomListModel77da62_QRegisterMetaTypes();callbackCustomListModel77da62_Constructor(this);};
	void Signal_Clear() { callbackCustomListModel77da62_Clear(this); };
	void Signal_Add(quintptr account) { callbackCustomListModel77da62_Add(this, account); };
	QString updateRequired() { return ({ Moc_PackedString tempVal = callbackCustomListModel77da62_UpdateRequired(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setUpdateRequired(QString updateRequired) { QByteArray t4432f4 = updateRequired.toUtf8(); Moc_PackedString updateRequiredPacked = { const_cast<char*>(t4432f4.prepend("WHITESPACE").constData()+10), t4432f4.size()-10 };callbackCustomListModel77da62_SetUpdateRequired(this, updateRequiredPacked); };
	void Signal_UpdateRequiredChanged(QString updateRequired) { QByteArray t4432f4 = updateRequired.toUtf8(); Moc_PackedString updateRequiredPacked = { const_cast<char*>(t4432f4.prepend("WHITESPACE").constData()+10), t4432f4.size()-10 };callbackCustomListModel77da62_UpdateRequiredChanged(this, updateRequiredPacked); };
	 ~CustomListModel77da62() { callbackCustomListModel77da62_DestroyCustomListModel(this); };
	bool dropMimeData(const QMimeData * data, Qt::DropAction action, int row, int column, const QModelIndex & parent) { return callbackCustomListModel77da62_DropMimeData(this, const_cast<QMimeData*>(data), action, row, column, const_cast<QModelIndex*>(&parent)) != 0; };
	QModelIndex index(int row, int column, const QModelIndex & parent) const { return *static_cast<QModelIndex*>(callbackCustomListModel77da62_Index(const_cast<void*>(static_cast<const void*>(this)), row, column, const_cast<QModelIndex*>(&parent))); };
	QModelIndex sibling(int row, int column, const QModelIndex & idx) const { return *static_cast<QModelIndex*>(callbackCustomListModel77da62_Sibling(const_cast<void*>(static_cast<const void*>(this)), row, column, const_cast<QModelIndex*>(&idx))); };
	Qt::ItemFlags flags(const QModelIndex & index) const { return static_cast<Qt::ItemFlag>(callbackCustomListModel77da62_Flags(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	bool insertColumns(int column, int count, const QModelIndex & parent) { return callbackCustomListModel77da62_InsertColumns(this, column, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool insertRows(int row, int count, const QModelIndex & parent) { return callbackCustomListModel77da62_InsertRows(this, row, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool moveColumns(const QModelIndex & sourceParent, int sourceColumn, int count, const QModelIndex & destinationParent, int destinationChild) { return callbackCustomListModel77da62_MoveColumns(this, const_cast<QModelIndex*>(&sourceParent), sourceColumn, count, const_cast<QModelIndex*>(&destinationParent), destinationChild) != 0; };
	bool moveRows(const QModelIndex & sourceParent, int sourceRow, int count, const QModelIndex & destinationParent, int destinationChild) { return callbackCustomListModel77da62_MoveRows(this, const_cast<QModelIndex*>(&sourceParent), sourceRow, count, const_cast<QModelIndex*>(&destinationParent), destinationChild) != 0; };
	bool removeColumns(int column, int count, const QModelIndex & parent) { return callbackCustomListModel77da62_RemoveColumns(this, column, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool removeRows(int row, int count, const QModelIndex & parent) { return callbackCustomListModel77da62_RemoveRows(this, row, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool setData(const QModelIndex & index, const QVariant & value, int role) { return callbackCustomListModel77da62_SetData(this, const_cast<QModelIndex*>(&index), const_cast<QVariant*>(&value), role) != 0; };
	bool setHeaderData(int section, Qt::Orientation orientation, const QVariant & value, int role) { return callbackCustomListModel77da62_SetHeaderData(this, section, orientation, const_cast<QVariant*>(&value), role) != 0; };
	bool setItemData(const QModelIndex & index, const QMap<int, QVariant> & roles) { return callbackCustomListModel77da62_SetItemData(this, const_cast<QModelIndex*>(&index), ({ QMap<int, QVariant>* tmpValue = const_cast<QMap<int, QVariant>*>(&roles); Moc_PackedList { tmpValue, tmpValue->size() }; })) != 0; };
	bool submit() { return callbackCustomListModel77da62_Submit(this) != 0; };
	void Signal_ColumnsAboutToBeInserted(const QModelIndex & parent, int first, int last) { callbackCustomListModel77da62_ColumnsAboutToBeInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsAboutToBeMoved(const QModelIndex & sourceParent, int sourceStart, int sourceEnd, const QModelIndex & destinationParent, int destinationColumn) { callbackCustomListModel77da62_ColumnsAboutToBeMoved(this, const_cast<QModelIndex*>(&sourceParent), sourceStart, sourceEnd, const_cast<QModelIndex*>(&destinationParent), destinationColumn); };
	void Signal_ColumnsAboutToBeRemoved(const QModelIndex & parent, int first, int last) { callbackCustomListModel77da62_ColumnsAboutToBeRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsInserted(const QModelIndex & parent, int first, int last) { callbackCustomListModel77da62_ColumnsInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsMoved(const QModelIndex & parent, int start, int end, const QModelIndex & destination, int column) { callbackCustomListModel77da62_ColumnsMoved(this, const_cast<QModelIndex*>(&parent), start, end, const_cast<QModelIndex*>(&destination), column); };
	void Signal_ColumnsRemoved(const QModelIndex & parent, int first, int last) { callbackCustomListModel77da62_ColumnsRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_DataChanged(const QModelIndex & topLeft, const QModelIndex & bottomRight, const QVector<int> & roles) { callbackCustomListModel77da62_DataChanged(this, const_cast<QModelIndex*>(&topLeft), const_cast<QModelIndex*>(&bottomRight), ({ QVector<int>* tmpValue = const_cast<QVector<int>*>(&roles); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	void fetchMore(const QModelIndex & parent) { callbackCustomListModel77da62_FetchMore(this, const_cast<QModelIndex*>(&parent)); };
	void Signal_HeaderDataChanged(Qt::Orientation orientation, int first, int last) { callbackCustomListModel77da62_HeaderDataChanged(this, orientation, first, last); };
	void Signal_LayoutAboutToBeChanged(const QList<QPersistentModelIndex> & parents, QAbstractItemModel::LayoutChangeHint hint) { callbackCustomListModel77da62_LayoutAboutToBeChanged(this, ({ QList<QPersistentModelIndex>* tmpValue = const_cast<QList<QPersistentModelIndex>*>(&parents); Moc_PackedList { tmpValue, tmpValue->size() }; }), hint); };
	void Signal_LayoutChanged(const QList<QPersistentModelIndex> & parents, QAbstractItemModel::LayoutChangeHint hint) { callbackCustomListModel77da62_LayoutChanged(this, ({ QList<QPersistentModelIndex>* tmpValue = const_cast<QList<QPersistentModelIndex>*>(&parents); Moc_PackedList { tmpValue, tmpValue->size() }; }), hint); };
	void Signal_ModelAboutToBeReset() { callbackCustomListModel77da62_ModelAboutToBeReset(this); };
	void Signal_ModelReset() { callbackCustomListModel77da62_ModelReset(this); };
	void resetInternalData() { callbackCustomListModel77da62_ResetInternalData(this); };
	void revert() { callbackCustomListModel77da62_Revert(this); };
	void Signal_RowsAboutToBeInserted(const QModelIndex & parent, int start, int end) { callbackCustomListModel77da62_RowsAboutToBeInserted(this, const_cast<QModelIndex*>(&parent), start, end); };
	void Signal_RowsAboutToBeMoved(const QModelIndex & sourceParent, int sourceStart, int sourceEnd, const QModelIndex & destinationParent, int destinationRow) { callbackCustomListModel77da62_RowsAboutToBeMoved(this, const_cast<QModelIndex*>(&sourceParent), sourceStart, sourceEnd, const_cast<QModelIndex*>(&destinationParent), destinationRow); };
	void Signal_RowsAboutToBeRemoved(const QModelIndex & parent, int first, int last) { callbackCustomListModel77da62_RowsAboutToBeRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_RowsInserted(const QModelIndex & parent, int first, int last) { callbackCustomListModel77da62_RowsInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_RowsMoved(const QModelIndex & parent, int start, int end, const QModelIndex & destination, int row) { callbackCustomListModel77da62_RowsMoved(this, const_cast<QModelIndex*>(&parent), start, end, const_cast<QModelIndex*>(&destination), row); };
	void Signal_RowsRemoved(const QModelIndex & parent, int first, int last) { callbackCustomListModel77da62_RowsRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void sort(int column, Qt::SortOrder order) { callbackCustomListModel77da62_Sort(this, column, order); };
	QHash<int, QByteArray> roleNames() const { return ({ QHash<int, QByteArray>* tmpP = static_cast<QHash<int, QByteArray>*>(callbackCustomListModel77da62_RoleNames(const_cast<void*>(static_cast<const void*>(this)))); QHash<int, QByteArray> tmpV = *tmpP; tmpP->~QHash(); free(tmpP); tmpV; }); };
	QMap<int, QVariant> itemData(const QModelIndex & index) const { return ({ QMap<int, QVariant>* tmpP = static_cast<QMap<int, QVariant>*>(callbackCustomListModel77da62_ItemData(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); QMap<int, QVariant> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }); };
	QMimeData * mimeData(const QModelIndexList & indexes) const { return static_cast<QMimeData*>(callbackCustomListModel77da62_MimeData(const_cast<void*>(static_cast<const void*>(this)), ({ QList<QModelIndex>* tmpValue = new QList<QModelIndex>(indexes); Moc_PackedList { tmpValue, tmpValue->size() }; }))); };
	QModelIndex buddy(const QModelIndex & index) const { return *static_cast<QModelIndex*>(callbackCustomListModel77da62_Buddy(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	QModelIndex parent(const QModelIndex & index) const { return *static_cast<QModelIndex*>(callbackCustomListModel77da62_Parent(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	QList<QModelIndex> match(const QModelIndex & start, int role, const QVariant & value, int hits, Qt::MatchFlags flags) const { return ({ QList<QModelIndex>* tmpP = static_cast<QList<QModelIndex>*>(callbackCustomListModel77da62_Match(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&start), role, const_cast<QVariant*>(&value), hits, flags)); QList<QModelIndex> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }); };
	QSize span(const QModelIndex & index) const { return *static_cast<QSize*>(callbackCustomListModel77da62_Span(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	QStringList mimeTypes() const { return ({ Moc_PackedString tempVal = callbackCustomListModel77da62_MimeTypes(const_cast<void*>(static_cast<const void*>(this))); QStringList ret = QString::fromUtf8(tempVal.data, tempVal.len).split("|", QString::SkipEmptyParts); free(tempVal.data); ret; }); };
	QVariant data(const QModelIndex & index, int role) const { return *static_cast<QVariant*>(callbackCustomListModel77da62_Data(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index), role)); };
	QVariant headerData(int section, Qt::Orientation orientation, int role) const { return *static_cast<QVariant*>(callbackCustomListModel77da62_HeaderData(const_cast<void*>(static_cast<const void*>(this)), section, orientation, role)); };
	Qt::DropActions supportedDragActions() const { return static_cast<Qt::DropAction>(callbackCustomListModel77da62_SupportedDragActions(const_cast<void*>(static_cast<const void*>(this)))); };
	Qt::DropActions supportedDropActions() const { return static_cast<Qt::DropAction>(callbackCustomListModel77da62_SupportedDropActions(const_cast<void*>(static_cast<const void*>(this)))); };
	bool canDropMimeData(const QMimeData * data, Qt::DropAction action, int row, int column, const QModelIndex & parent) const { return callbackCustomListModel77da62_CanDropMimeData(const_cast<void*>(static_cast<const void*>(this)), const_cast<QMimeData*>(data), action, row, column, const_cast<QModelIndex*>(&parent)) != 0; };
	bool canFetchMore(const QModelIndex & parent) const { return callbackCustomListModel77da62_CanFetchMore(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)) != 0; };
	bool hasChildren(const QModelIndex & parent) const { return callbackCustomListModel77da62_HasChildren(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)) != 0; };
	int columnCount(const QModelIndex & parent) const { return callbackCustomListModel77da62_ColumnCount(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)); };
	int rowCount(const QModelIndex & parent) const { return callbackCustomListModel77da62_RowCount(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)); };
	bool event(QEvent * e) { return callbackCustomListModel77da62_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackCustomListModel77da62_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackCustomListModel77da62_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackCustomListModel77da62_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackCustomListModel77da62_CustomEvent(this, event); };
	void deleteLater() { callbackCustomListModel77da62_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackCustomListModel77da62_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackCustomListModel77da62_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackCustomListModel77da62_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackCustomListModel77da62_TimerEvent(this, event); };
	QString updateRequiredDefault() { return _updateRequired; };
	void setUpdateRequiredDefault(QString p) { if (p != _updateRequired) { _updateRequired = p; updateRequiredChanged(_updateRequired); } };
signals:
	void clear();
	void add(quintptr account);
	void updateRequiredChanged(QString updateRequired);
public slots:
	void callbackFromQml() { callbackCustomListModel77da62_CallbackFromQml(this); };
private:
	QString _updateRequired;
};

Q_DECLARE_METATYPE(CustomListModel77da62*)


void CustomListModel77da62_CustomListModel77da62_QRegisterMetaTypes() {
	qRegisterMetaType<QString>();
}

class TxListCtx77da62: public QObject
{
Q_OBJECT
Q_PROPERTY(QString shortenAddress READ shortenAddress WRITE setShortenAddress NOTIFY shortenAddressChanged)
Q_PROPERTY(QString selectedSrc READ selectedSrc WRITE setSelectedSrc NOTIFY selectedSrcChanged)
public:
	TxListCtx77da62(QObject *parent = Q_NULLPTR) : QObject(parent) {qRegisterMetaType<quintptr>("quintptr");TxListCtx77da62_TxListCtx77da62_QRegisterMetaType();TxListCtx77da62_TxListCtx77da62_QRegisterMetaTypes();callbackTxListCtx77da62_Constructor(this);};
	void Signal_Clicked(qint32 b) { callbackTxListCtx77da62_Clicked(this, b); };
	QString shortenAddress() { return ({ Moc_PackedString tempVal = callbackTxListCtx77da62_ShortenAddress(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setShortenAddress(QString shortenAddress) { QByteArray t3fdf7b = shortenAddress.toUtf8(); Moc_PackedString shortenAddressPacked = { const_cast<char*>(t3fdf7b.prepend("WHITESPACE").constData()+10), t3fdf7b.size()-10 };callbackTxListCtx77da62_SetShortenAddress(this, shortenAddressPacked); };
	void Signal_ShortenAddressChanged(QString shortenAddress) { QByteArray t3fdf7b = shortenAddress.toUtf8(); Moc_PackedString shortenAddressPacked = { const_cast<char*>(t3fdf7b.prepend("WHITESPACE").constData()+10), t3fdf7b.size()-10 };callbackTxListCtx77da62_ShortenAddressChanged(this, shortenAddressPacked); };
	QString selectedSrc() { return ({ Moc_PackedString tempVal = callbackTxListCtx77da62_SelectedSrc(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setSelectedSrc(QString selectedSrc) { QByteArray t5f7742 = selectedSrc.toUtf8(); Moc_PackedString selectedSrcPacked = { const_cast<char*>(t5f7742.prepend("WHITESPACE").constData()+10), t5f7742.size()-10 };callbackTxListCtx77da62_SetSelectedSrc(this, selectedSrcPacked); };
	void Signal_SelectedSrcChanged(QString selectedSrc) { QByteArray t5f7742 = selectedSrc.toUtf8(); Moc_PackedString selectedSrcPacked = { const_cast<char*>(t5f7742.prepend("WHITESPACE").constData()+10), t5f7742.size()-10 };callbackTxListCtx77da62_SelectedSrcChanged(this, selectedSrcPacked); };
	 ~TxListCtx77da62() { callbackTxListCtx77da62_DestroyTxListCtx(this); };
	bool event(QEvent * e) { return callbackTxListCtx77da62_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackTxListCtx77da62_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackTxListCtx77da62_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackTxListCtx77da62_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackTxListCtx77da62_CustomEvent(this, event); };
	void deleteLater() { callbackTxListCtx77da62_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackTxListCtx77da62_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackTxListCtx77da62_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackTxListCtx77da62_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackTxListCtx77da62_TimerEvent(this, event); };
	QString shortenAddressDefault() { return _shortenAddress; };
	void setShortenAddressDefault(QString p) { if (p != _shortenAddress) { _shortenAddress = p; shortenAddressChanged(_shortenAddress); } };
	QString selectedSrcDefault() { return _selectedSrc; };
	void setSelectedSrcDefault(QString p) { if (p != _selectedSrc) { _selectedSrc = p; selectedSrcChanged(_selectedSrc); } };
signals:
	void clicked(qint32 b);
	void shortenAddressChanged(QString shortenAddress);
	void selectedSrcChanged(QString selectedSrc);
public slots:
private:
	QString _shortenAddress;
	QString _selectedSrc;
};

Q_DECLARE_METATYPE(TxListCtx77da62*)


void TxListCtx77da62_TxListCtx77da62_QRegisterMetaTypes() {
	qRegisterMetaType<QString>();
}

class ApproveNewAccountCtx77da62: public QObject
{
Q_OBJECT
Q_PROPERTY(QString remote READ remote WRITE setRemote NOTIFY remoteChanged)
Q_PROPERTY(QString transport READ transport WRITE setTransport NOTIFY transportChanged)
Q_PROPERTY(QString endpoint READ endpoint WRITE setEndpoint NOTIFY endpointChanged)
Q_PROPERTY(QString password READ password WRITE setPassword NOTIFY passwordChanged)
Q_PROPERTY(QString confirmPassword READ confirmPassword WRITE setConfirmPassword NOTIFY confirmPasswordChanged)
public:
	ApproveNewAccountCtx77da62(QObject *parent = Q_NULLPTR) : QObject(parent) {qRegisterMetaType<quintptr>("quintptr");ApproveNewAccountCtx77da62_ApproveNewAccountCtx77da62_QRegisterMetaType();ApproveNewAccountCtx77da62_ApproveNewAccountCtx77da62_QRegisterMetaTypes();callbackApproveNewAccountCtx77da62_Constructor(this);};
	void Signal_Clicked(qint32 b) { callbackApproveNewAccountCtx77da62_Clicked(this, b); };
	void Signal_Back() { callbackApproveNewAccountCtx77da62_Back(this); };
	void Signal_PasswordEdited(QString b) { QByteArray te9d71f = b.toUtf8(); Moc_PackedString bPacked = { const_cast<char*>(te9d71f.prepend("WHITESPACE").constData()+10), te9d71f.size()-10 };callbackApproveNewAccountCtx77da62_PasswordEdited(this, bPacked); };
	void Signal_ConfirmPasswordEdited(QString b) { QByteArray te9d71f = b.toUtf8(); Moc_PackedString bPacked = { const_cast<char*>(te9d71f.prepend("WHITESPACE").constData()+10), te9d71f.size()-10 };callbackApproveNewAccountCtx77da62_ConfirmPasswordEdited(this, bPacked); };
	QString remote() { return ({ Moc_PackedString tempVal = callbackApproveNewAccountCtx77da62_Remote(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setRemote(QString remote) { QByteArray t41ffe5 = remote.toUtf8(); Moc_PackedString remotePacked = { const_cast<char*>(t41ffe5.prepend("WHITESPACE").constData()+10), t41ffe5.size()-10 };callbackApproveNewAccountCtx77da62_SetRemote(this, remotePacked); };
	void Signal_RemoteChanged(QString remote) { QByteArray t41ffe5 = remote.toUtf8(); Moc_PackedString remotePacked = { const_cast<char*>(t41ffe5.prepend("WHITESPACE").constData()+10), t41ffe5.size()-10 };callbackApproveNewAccountCtx77da62_RemoteChanged(this, remotePacked); };
	QString transport() { return ({ Moc_PackedString tempVal = callbackApproveNewAccountCtx77da62_Transport(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setTransport(QString transport) { QByteArray ta8e601 = transport.toUtf8(); Moc_PackedString transportPacked = { const_cast<char*>(ta8e601.prepend("WHITESPACE").constData()+10), ta8e601.size()-10 };callbackApproveNewAccountCtx77da62_SetTransport(this, transportPacked); };
	void Signal_TransportChanged(QString transport) { QByteArray ta8e601 = transport.toUtf8(); Moc_PackedString transportPacked = { const_cast<char*>(ta8e601.prepend("WHITESPACE").constData()+10), ta8e601.size()-10 };callbackApproveNewAccountCtx77da62_TransportChanged(this, transportPacked); };
	QString endpoint() { return ({ Moc_PackedString tempVal = callbackApproveNewAccountCtx77da62_Endpoint(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setEndpoint(QString endpoint) { QByteArray te13fe4 = endpoint.toUtf8(); Moc_PackedString endpointPacked = { const_cast<char*>(te13fe4.prepend("WHITESPACE").constData()+10), te13fe4.size()-10 };callbackApproveNewAccountCtx77da62_SetEndpoint(this, endpointPacked); };
	void Signal_EndpointChanged(QString endpoint) { QByteArray te13fe4 = endpoint.toUtf8(); Moc_PackedString endpointPacked = { const_cast<char*>(te13fe4.prepend("WHITESPACE").constData()+10), te13fe4.size()-10 };callbackApproveNewAccountCtx77da62_EndpointChanged(this, endpointPacked); };
	QString password() { return ({ Moc_PackedString tempVal = callbackApproveNewAccountCtx77da62_Password(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setPassword(QString password) { QByteArray t5baa61 = password.toUtf8(); Moc_PackedString passwordPacked = { const_cast<char*>(t5baa61.prepend("WHITESPACE").constData()+10), t5baa61.size()-10 };callbackApproveNewAccountCtx77da62_SetPassword(this, passwordPacked); };
	void Signal_PasswordChanged(QString password) { QByteArray t5baa61 = password.toUtf8(); Moc_PackedString passwordPacked = { const_cast<char*>(t5baa61.prepend("WHITESPACE").constData()+10), t5baa61.size()-10 };callbackApproveNewAccountCtx77da62_PasswordChanged(this, passwordPacked); };
	QString confirmPassword() { return ({ Moc_PackedString tempVal = callbackApproveNewAccountCtx77da62_ConfirmPassword(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setConfirmPassword(QString confirmPassword) { QByteArray t802b53 = confirmPassword.toUtf8(); Moc_PackedString confirmPasswordPacked = { const_cast<char*>(t802b53.prepend("WHITESPACE").constData()+10), t802b53.size()-10 };callbackApproveNewAccountCtx77da62_SetConfirmPassword(this, confirmPasswordPacked); };
	void Signal_ConfirmPasswordChanged(QString confirmPassword) { QByteArray t802b53 = confirmPassword.toUtf8(); Moc_PackedString confirmPasswordPacked = { const_cast<char*>(t802b53.prepend("WHITESPACE").constData()+10), t802b53.size()-10 };callbackApproveNewAccountCtx77da62_ConfirmPasswordChanged(this, confirmPasswordPacked); };
	 ~ApproveNewAccountCtx77da62() { callbackApproveNewAccountCtx77da62_DestroyApproveNewAccountCtx(this); };
	bool event(QEvent * e) { return callbackApproveNewAccountCtx77da62_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackApproveNewAccountCtx77da62_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackApproveNewAccountCtx77da62_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackApproveNewAccountCtx77da62_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackApproveNewAccountCtx77da62_CustomEvent(this, event); };
	void deleteLater() { callbackApproveNewAccountCtx77da62_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackApproveNewAccountCtx77da62_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackApproveNewAccountCtx77da62_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackApproveNewAccountCtx77da62_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackApproveNewAccountCtx77da62_TimerEvent(this, event); };
	QString remoteDefault() { return _remote; };
	void setRemoteDefault(QString p) { if (p != _remote) { _remote = p; remoteChanged(_remote); } };
	QString transportDefault() { return _transport; };
	void setTransportDefault(QString p) { if (p != _transport) { _transport = p; transportChanged(_transport); } };
	QString endpointDefault() { return _endpoint; };
	void setEndpointDefault(QString p) { if (p != _endpoint) { _endpoint = p; endpointChanged(_endpoint); } };
	QString passwordDefault() { return _password; };
	void setPasswordDefault(QString p) { if (p != _password) { _password = p; passwordChanged(_password); } };
	QString confirmPasswordDefault() { return _confirmPassword; };
	void setConfirmPasswordDefault(QString p) { if (p != _confirmPassword) { _confirmPassword = p; confirmPasswordChanged(_confirmPassword); } };
signals:
	void clicked(qint32 b);
	void back();
	void passwordEdited(QString b);
	void confirmPasswordEdited(QString b);
	void remoteChanged(QString remote);
	void transportChanged(QString transport);
	void endpointChanged(QString endpoint);
	void passwordChanged(QString password);
	void confirmPasswordChanged(QString confirmPassword);
public slots:
private:
	QString _remote;
	QString _transport;
	QString _endpoint;
	QString _password;
	QString _confirmPassword;
};

Q_DECLARE_METATYPE(ApproveNewAccountCtx77da62*)


void ApproveNewAccountCtx77da62_ApproveNewAccountCtx77da62_QRegisterMetaTypes() {
	qRegisterMetaType<QString>();
}

class ApproveTxCtx77da62: public QObject
{
Q_OBJECT
Q_PROPERTY(qint32 valueUnit READ valueUnit WRITE setValueUnit NOTIFY valueUnitChanged)
Q_PROPERTY(QString remote READ remote WRITE setRemote NOTIFY remoteChanged)
Q_PROPERTY(QString transport READ transport WRITE setTransport NOTIFY transportChanged)
Q_PROPERTY(QString endpoint READ endpoint WRITE setEndpoint NOTIFY endpointChanged)
Q_PROPERTY(QString data READ data WRITE setData NOTIFY dataChanged)
Q_PROPERTY(QString from READ from WRITE setFrom NOTIFY fromChanged)
Q_PROPERTY(QString fromWarning READ fromWarning WRITE setFromWarning NOTIFY fromWarningChanged)
Q_PROPERTY(bool fromVisible READ isFromVisible WRITE setFromVisible NOTIFY fromVisibleChanged)
Q_PROPERTY(QString to READ to WRITE setTo NOTIFY toChanged)
Q_PROPERTY(QString toWarning READ toWarning WRITE setToWarning NOTIFY toWarningChanged)
Q_PROPERTY(bool toVisible READ isToVisible WRITE setToVisible NOTIFY toVisibleChanged)
Q_PROPERTY(QString gas READ gas WRITE setGas NOTIFY gasChanged)
Q_PROPERTY(QString gasPrice READ gasPrice WRITE setGasPrice NOTIFY gasPriceChanged)
Q_PROPERTY(qint32 gasPriceUnit READ gasPriceUnit WRITE setGasPriceUnit NOTIFY gasPriceUnitChanged)
Q_PROPERTY(QString nonce READ nonce WRITE setNonce NOTIFY nonceChanged)
Q_PROPERTY(QString value READ value WRITE setValue NOTIFY valueChanged)
Q_PROPERTY(QString password READ password WRITE setPassword NOTIFY passwordChanged)
Q_PROPERTY(QString fromSrc READ fromSrc WRITE setFromSrc NOTIFY fromSrcChanged)
Q_PROPERTY(QString toSrc READ toSrc WRITE setToSrc NOTIFY toSrcChanged)
Q_PROPERTY(QString diff READ diff WRITE setDiff NOTIFY diffChanged)
public:
	ApproveTxCtx77da62(QObject *parent = Q_NULLPTR) : QObject(parent) {qRegisterMetaType<quintptr>("quintptr");ApproveTxCtx77da62_ApproveTxCtx77da62_QRegisterMetaType();ApproveTxCtx77da62_ApproveTxCtx77da62_QRegisterMetaTypes();callbackApproveTxCtx77da62_Constructor(this);};
	void Signal_Approve() { callbackApproveTxCtx77da62_Approve(this); };
	void Signal_Reject() { callbackApproveTxCtx77da62_Reject(this); };
	void Signal_CheckTxDiff() { callbackApproveTxCtx77da62_CheckTxDiff(this); };
	void Signal_Back() { callbackApproveTxCtx77da62_Back(this); };
	void Signal_Edited(QString s, QString v) { QByteArray ta0f149 = s.toUtf8(); Moc_PackedString sPacked = { const_cast<char*>(ta0f149.prepend("WHITESPACE").constData()+10), ta0f149.size()-10 };QByteArray t7a38d8 = v.toUtf8(); Moc_PackedString vPacked = { const_cast<char*>(t7a38d8.prepend("WHITESPACE").constData()+10), t7a38d8.size()-10 };callbackApproveTxCtx77da62_Edited(this, sPacked, vPacked); };
	void Signal_ChangeValueUnit(qint32 v) { callbackApproveTxCtx77da62_ChangeValueUnit(this, v); };
	void Signal_ChangeGasPriceUnit(qint32 v) { callbackApproveTxCtx77da62_ChangeGasPriceUnit(this, v); };
	qint32 valueUnit() { return callbackApproveTxCtx77da62_ValueUnit(this); };
	void setValueUnit(qint32 valueUnit) { callbackApproveTxCtx77da62_SetValueUnit(this, valueUnit); };
	void Signal_ValueUnitChanged(qint32 valueUnit) { callbackApproveTxCtx77da62_ValueUnitChanged(this, valueUnit); };
	QString remote() { return ({ Moc_PackedString tempVal = callbackApproveTxCtx77da62_Remote(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setRemote(QString remote) { QByteArray t41ffe5 = remote.toUtf8(); Moc_PackedString remotePacked = { const_cast<char*>(t41ffe5.prepend("WHITESPACE").constData()+10), t41ffe5.size()-10 };callbackApproveTxCtx77da62_SetRemote(this, remotePacked); };
	void Signal_RemoteChanged(QString remote) { QByteArray t41ffe5 = remote.toUtf8(); Moc_PackedString remotePacked = { const_cast<char*>(t41ffe5.prepend("WHITESPACE").constData()+10), t41ffe5.size()-10 };callbackApproveTxCtx77da62_RemoteChanged(this, remotePacked); };
	QString transport() { return ({ Moc_PackedString tempVal = callbackApproveTxCtx77da62_Transport(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setTransport(QString transport) { QByteArray ta8e601 = transport.toUtf8(); Moc_PackedString transportPacked = { const_cast<char*>(ta8e601.prepend("WHITESPACE").constData()+10), ta8e601.size()-10 };callbackApproveTxCtx77da62_SetTransport(this, transportPacked); };
	void Signal_TransportChanged(QString transport) { QByteArray ta8e601 = transport.toUtf8(); Moc_PackedString transportPacked = { const_cast<char*>(ta8e601.prepend("WHITESPACE").constData()+10), ta8e601.size()-10 };callbackApproveTxCtx77da62_TransportChanged(this, transportPacked); };
	QString endpoint() { return ({ Moc_PackedString tempVal = callbackApproveTxCtx77da62_Endpoint(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setEndpoint(QString endpoint) { QByteArray te13fe4 = endpoint.toUtf8(); Moc_PackedString endpointPacked = { const_cast<char*>(te13fe4.prepend("WHITESPACE").constData()+10), te13fe4.size()-10 };callbackApproveTxCtx77da62_SetEndpoint(this, endpointPacked); };
	void Signal_EndpointChanged(QString endpoint) { QByteArray te13fe4 = endpoint.toUtf8(); Moc_PackedString endpointPacked = { const_cast<char*>(te13fe4.prepend("WHITESPACE").constData()+10), te13fe4.size()-10 };callbackApproveTxCtx77da62_EndpointChanged(this, endpointPacked); };
	QString data() { return ({ Moc_PackedString tempVal = callbackApproveTxCtx77da62_Data(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setData(QString data) { QByteArray ta17c9a = data.toUtf8(); Moc_PackedString dataPacked = { const_cast<char*>(ta17c9a.prepend("WHITESPACE").constData()+10), ta17c9a.size()-10 };callbackApproveTxCtx77da62_SetData(this, dataPacked); };
	void Signal_DataChanged(QString data) { QByteArray ta17c9a = data.toUtf8(); Moc_PackedString dataPacked = { const_cast<char*>(ta17c9a.prepend("WHITESPACE").constData()+10), ta17c9a.size()-10 };callbackApproveTxCtx77da62_DataChanged(this, dataPacked); };
	QString from() { return ({ Moc_PackedString tempVal = callbackApproveTxCtx77da62_From(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setFrom(QString from) { QByteArray t0b1e95 = from.toUtf8(); Moc_PackedString fromPacked = { const_cast<char*>(t0b1e95.prepend("WHITESPACE").constData()+10), t0b1e95.size()-10 };callbackApproveTxCtx77da62_SetFrom(this, fromPacked); };
	void Signal_FromChanged(QString from) { QByteArray t0b1e95 = from.toUtf8(); Moc_PackedString fromPacked = { const_cast<char*>(t0b1e95.prepend("WHITESPACE").constData()+10), t0b1e95.size()-10 };callbackApproveTxCtx77da62_FromChanged(this, fromPacked); };
	QString fromWarning() { return ({ Moc_PackedString tempVal = callbackApproveTxCtx77da62_FromWarning(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setFromWarning(QString fromWarning) { QByteArray t388a5b = fromWarning.toUtf8(); Moc_PackedString fromWarningPacked = { const_cast<char*>(t388a5b.prepend("WHITESPACE").constData()+10), t388a5b.size()-10 };callbackApproveTxCtx77da62_SetFromWarning(this, fromWarningPacked); };
	void Signal_FromWarningChanged(QString fromWarning) { QByteArray t388a5b = fromWarning.toUtf8(); Moc_PackedString fromWarningPacked = { const_cast<char*>(t388a5b.prepend("WHITESPACE").constData()+10), t388a5b.size()-10 };callbackApproveTxCtx77da62_FromWarningChanged(this, fromWarningPacked); };
	bool isFromVisible() { return callbackApproveTxCtx77da62_IsFromVisible(this) != 0; };
	void setFromVisible(bool fromVisible) { callbackApproveTxCtx77da62_SetFromVisible(this, fromVisible); };
	void Signal_FromVisibleChanged(bool fromVisible) { callbackApproveTxCtx77da62_FromVisibleChanged(this, fromVisible); };
	QString to() { return ({ Moc_PackedString tempVal = callbackApproveTxCtx77da62_To(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setTo(QString to) { QByteArray t4374aa = to.toUtf8(); Moc_PackedString toPacked = { const_cast<char*>(t4374aa.prepend("WHITESPACE").constData()+10), t4374aa.size()-10 };callbackApproveTxCtx77da62_SetTo(this, toPacked); };
	void Signal_ToChanged(QString to) { QByteArray t4374aa = to.toUtf8(); Moc_PackedString toPacked = { const_cast<char*>(t4374aa.prepend("WHITESPACE").constData()+10), t4374aa.size()-10 };callbackApproveTxCtx77da62_ToChanged(this, toPacked); };
	QString toWarning() { return ({ Moc_PackedString tempVal = callbackApproveTxCtx77da62_ToWarning(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setToWarning(QString toWarning) { QByteArray t9fefd3 = toWarning.toUtf8(); Moc_PackedString toWarningPacked = { const_cast<char*>(t9fefd3.prepend("WHITESPACE").constData()+10), t9fefd3.size()-10 };callbackApproveTxCtx77da62_SetToWarning(this, toWarningPacked); };
	void Signal_ToWarningChanged(QString toWarning) { QByteArray t9fefd3 = toWarning.toUtf8(); Moc_PackedString toWarningPacked = { const_cast<char*>(t9fefd3.prepend("WHITESPACE").constData()+10), t9fefd3.size()-10 };callbackApproveTxCtx77da62_ToWarningChanged(this, toWarningPacked); };
	bool isToVisible() { return callbackApproveTxCtx77da62_IsToVisible(this) != 0; };
	void setToVisible(bool toVisible) { callbackApproveTxCtx77da62_SetToVisible(this, toVisible); };
	void Signal_ToVisibleChanged(bool toVisible) { callbackApproveTxCtx77da62_ToVisibleChanged(this, toVisible); };
	QString gas() { return ({ Moc_PackedString tempVal = callbackApproveTxCtx77da62_Gas(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setGas(QString gas) { QByteArray tfacafa = gas.toUtf8(); Moc_PackedString gasPacked = { const_cast<char*>(tfacafa.prepend("WHITESPACE").constData()+10), tfacafa.size()-10 };callbackApproveTxCtx77da62_SetGas(this, gasPacked); };
	void Signal_GasChanged(QString gas) { QByteArray tfacafa = gas.toUtf8(); Moc_PackedString gasPacked = { const_cast<char*>(tfacafa.prepend("WHITESPACE").constData()+10), tfacafa.size()-10 };callbackApproveTxCtx77da62_GasChanged(this, gasPacked); };
	QString gasPrice() { return ({ Moc_PackedString tempVal = callbackApproveTxCtx77da62_GasPrice(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setGasPrice(QString gasPrice) { QByteArray t72824c = gasPrice.toUtf8(); Moc_PackedString gasPricePacked = { const_cast<char*>(t72824c.prepend("WHITESPACE").constData()+10), t72824c.size()-10 };callbackApproveTxCtx77da62_SetGasPrice(this, gasPricePacked); };
	void Signal_GasPriceChanged(QString gasPrice) { QByteArray t72824c = gasPrice.toUtf8(); Moc_PackedString gasPricePacked = { const_cast<char*>(t72824c.prepend("WHITESPACE").constData()+10), t72824c.size()-10 };callbackApproveTxCtx77da62_GasPriceChanged(this, gasPricePacked); };
	qint32 gasPriceUnit() { return callbackApproveTxCtx77da62_GasPriceUnit(this); };
	void setGasPriceUnit(qint32 gasPriceUnit) { callbackApproveTxCtx77da62_SetGasPriceUnit(this, gasPriceUnit); };
	void Signal_GasPriceUnitChanged(qint32 gasPriceUnit) { callbackApproveTxCtx77da62_GasPriceUnitChanged(this, gasPriceUnit); };
	QString nonce() { return ({ Moc_PackedString tempVal = callbackApproveTxCtx77da62_Nonce(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setNonce(QString nonce) { QByteArray t49afa7 = nonce.toUtf8(); Moc_PackedString noncePacked = { const_cast<char*>(t49afa7.prepend("WHITESPACE").constData()+10), t49afa7.size()-10 };callbackApproveTxCtx77da62_SetNonce(this, noncePacked); };
	void Signal_NonceChanged(QString nonce) { QByteArray t49afa7 = nonce.toUtf8(); Moc_PackedString noncePacked = { const_cast<char*>(t49afa7.prepend("WHITESPACE").constData()+10), t49afa7.size()-10 };callbackApproveTxCtx77da62_NonceChanged(this, noncePacked); };
	QString value() { return ({ Moc_PackedString tempVal = callbackApproveTxCtx77da62_Value(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setValue(QString value) { QByteArray tf32b67 = value.toUtf8(); Moc_PackedString valuePacked = { const_cast<char*>(tf32b67.prepend("WHITESPACE").constData()+10), tf32b67.size()-10 };callbackApproveTxCtx77da62_SetValue(this, valuePacked); };
	void Signal_ValueChanged(QString value) { QByteArray tf32b67 = value.toUtf8(); Moc_PackedString valuePacked = { const_cast<char*>(tf32b67.prepend("WHITESPACE").constData()+10), tf32b67.size()-10 };callbackApproveTxCtx77da62_ValueChanged(this, valuePacked); };
	QString password() { return ({ Moc_PackedString tempVal = callbackApproveTxCtx77da62_Password(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setPassword(QString password) { QByteArray t5baa61 = password.toUtf8(); Moc_PackedString passwordPacked = { const_cast<char*>(t5baa61.prepend("WHITESPACE").constData()+10), t5baa61.size()-10 };callbackApproveTxCtx77da62_SetPassword(this, passwordPacked); };
	void Signal_PasswordChanged(QString password) { QByteArray t5baa61 = password.toUtf8(); Moc_PackedString passwordPacked = { const_cast<char*>(t5baa61.prepend("WHITESPACE").constData()+10), t5baa61.size()-10 };callbackApproveTxCtx77da62_PasswordChanged(this, passwordPacked); };
	QString fromSrc() { return ({ Moc_PackedString tempVal = callbackApproveTxCtx77da62_FromSrc(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setFromSrc(QString fromSrc) { QByteArray ta8ded4 = fromSrc.toUtf8(); Moc_PackedString fromSrcPacked = { const_cast<char*>(ta8ded4.prepend("WHITESPACE").constData()+10), ta8ded4.size()-10 };callbackApproveTxCtx77da62_SetFromSrc(this, fromSrcPacked); };
	void Signal_FromSrcChanged(QString fromSrc) { QByteArray ta8ded4 = fromSrc.toUtf8(); Moc_PackedString fromSrcPacked = { const_cast<char*>(ta8ded4.prepend("WHITESPACE").constData()+10), ta8ded4.size()-10 };callbackApproveTxCtx77da62_FromSrcChanged(this, fromSrcPacked); };
	QString toSrc() { return ({ Moc_PackedString tempVal = callbackApproveTxCtx77da62_ToSrc(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setToSrc(QString toSrc) { QByteArray t6f94e3 = toSrc.toUtf8(); Moc_PackedString toSrcPacked = { const_cast<char*>(t6f94e3.prepend("WHITESPACE").constData()+10), t6f94e3.size()-10 };callbackApproveTxCtx77da62_SetToSrc(this, toSrcPacked); };
	void Signal_ToSrcChanged(QString toSrc) { QByteArray t6f94e3 = toSrc.toUtf8(); Moc_PackedString toSrcPacked = { const_cast<char*>(t6f94e3.prepend("WHITESPACE").constData()+10), t6f94e3.size()-10 };callbackApproveTxCtx77da62_ToSrcChanged(this, toSrcPacked); };
	QString diff() { return ({ Moc_PackedString tempVal = callbackApproveTxCtx77da62_Diff(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setDiff(QString diff) { QByteArray t75a0ee = diff.toUtf8(); Moc_PackedString diffPacked = { const_cast<char*>(t75a0ee.prepend("WHITESPACE").constData()+10), t75a0ee.size()-10 };callbackApproveTxCtx77da62_SetDiff(this, diffPacked); };
	void Signal_DiffChanged(QString diff) { QByteArray t75a0ee = diff.toUtf8(); Moc_PackedString diffPacked = { const_cast<char*>(t75a0ee.prepend("WHITESPACE").constData()+10), t75a0ee.size()-10 };callbackApproveTxCtx77da62_DiffChanged(this, diffPacked); };
	 ~ApproveTxCtx77da62() { callbackApproveTxCtx77da62_DestroyApproveTxCtx(this); };
	bool event(QEvent * e) { return callbackApproveTxCtx77da62_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackApproveTxCtx77da62_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackApproveTxCtx77da62_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackApproveTxCtx77da62_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackApproveTxCtx77da62_CustomEvent(this, event); };
	void deleteLater() { callbackApproveTxCtx77da62_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackApproveTxCtx77da62_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackApproveTxCtx77da62_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackApproveTxCtx77da62_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackApproveTxCtx77da62_TimerEvent(this, event); };
	qint32 valueUnitDefault() { return _valueUnit; };
	void setValueUnitDefault(qint32 p) { if (p != _valueUnit) { _valueUnit = p; valueUnitChanged(_valueUnit); } };
	QString remoteDefault() { return _remote; };
	void setRemoteDefault(QString p) { if (p != _remote) { _remote = p; remoteChanged(_remote); } };
	QString transportDefault() { return _transport; };
	void setTransportDefault(QString p) { if (p != _transport) { _transport = p; transportChanged(_transport); } };
	QString endpointDefault() { return _endpoint; };
	void setEndpointDefault(QString p) { if (p != _endpoint) { _endpoint = p; endpointChanged(_endpoint); } };
	QString dataDefault() { return _data; };
	void setDataDefault(QString p) { if (p != _data) { _data = p; dataChanged(_data); } };
	QString fromDefault() { return _from; };
	void setFromDefault(QString p) { if (p != _from) { _from = p; fromChanged(_from); } };
	QString fromWarningDefault() { return _fromWarning; };
	void setFromWarningDefault(QString p) { if (p != _fromWarning) { _fromWarning = p; fromWarningChanged(_fromWarning); } };
	bool isFromVisibleDefault() { return _fromVisible; };
	void setFromVisibleDefault(bool p) { if (p != _fromVisible) { _fromVisible = p; fromVisibleChanged(_fromVisible); } };
	QString toDefault() { return _to; };
	void setToDefault(QString p) { if (p != _to) { _to = p; toChanged(_to); } };
	QString toWarningDefault() { return _toWarning; };
	void setToWarningDefault(QString p) { if (p != _toWarning) { _toWarning = p; toWarningChanged(_toWarning); } };
	bool isToVisibleDefault() { return _toVisible; };
	void setToVisibleDefault(bool p) { if (p != _toVisible) { _toVisible = p; toVisibleChanged(_toVisible); } };
	QString gasDefault() { return _gas; };
	void setGasDefault(QString p) { if (p != _gas) { _gas = p; gasChanged(_gas); } };
	QString gasPriceDefault() { return _gasPrice; };
	void setGasPriceDefault(QString p) { if (p != _gasPrice) { _gasPrice = p; gasPriceChanged(_gasPrice); } };
	qint32 gasPriceUnitDefault() { return _gasPriceUnit; };
	void setGasPriceUnitDefault(qint32 p) { if (p != _gasPriceUnit) { _gasPriceUnit = p; gasPriceUnitChanged(_gasPriceUnit); } };
	QString nonceDefault() { return _nonce; };
	void setNonceDefault(QString p) { if (p != _nonce) { _nonce = p; nonceChanged(_nonce); } };
	QString valueDefault() { return _value; };
	void setValueDefault(QString p) { if (p != _value) { _value = p; valueChanged(_value); } };
	QString passwordDefault() { return _password; };
	void setPasswordDefault(QString p) { if (p != _password) { _password = p; passwordChanged(_password); } };
	QString fromSrcDefault() { return _fromSrc; };
	void setFromSrcDefault(QString p) { if (p != _fromSrc) { _fromSrc = p; fromSrcChanged(_fromSrc); } };
	QString toSrcDefault() { return _toSrc; };
	void setToSrcDefault(QString p) { if (p != _toSrc) { _toSrc = p; toSrcChanged(_toSrc); } };
	QString diffDefault() { return _diff; };
	void setDiffDefault(QString p) { if (p != _diff) { _diff = p; diffChanged(_diff); } };
signals:
	void approve();
	void reject();
	void checkTxDiff();
	void back();
	void edited(QString s, QString v);
	void changeValueUnit(qint32 v);
	void changeGasPriceUnit(qint32 v);
	void valueUnitChanged(qint32 valueUnit);
	void remoteChanged(QString remote);
	void transportChanged(QString transport);
	void endpointChanged(QString endpoint);
	void dataChanged(QString data);
	void fromChanged(QString from);
	void fromWarningChanged(QString fromWarning);
	void fromVisibleChanged(bool fromVisible);
	void toChanged(QString to);
	void toWarningChanged(QString toWarning);
	void toVisibleChanged(bool toVisible);
	void gasChanged(QString gas);
	void gasPriceChanged(QString gasPrice);
	void gasPriceUnitChanged(qint32 gasPriceUnit);
	void nonceChanged(QString nonce);
	void valueChanged(QString value);
	void passwordChanged(QString password);
	void fromSrcChanged(QString fromSrc);
	void toSrcChanged(QString toSrc);
	void diffChanged(QString diff);
public slots:
private:
	qint32 _valueUnit;
	QString _remote;
	QString _transport;
	QString _endpoint;
	QString _data;
	QString _from;
	QString _fromWarning;
	bool _fromVisible;
	QString _to;
	QString _toWarning;
	bool _toVisible;
	QString _gas;
	QString _gasPrice;
	qint32 _gasPriceUnit;
	QString _nonce;
	QString _value;
	QString _password;
	QString _fromSrc;
	QString _toSrc;
	QString _diff;
};

Q_DECLARE_METATYPE(ApproveTxCtx77da62*)


void ApproveTxCtx77da62_ApproveTxCtx77da62_QRegisterMetaTypes() {
	qRegisterMetaType<QString>();
}

class LoginCtx77da62: public QObject
{
Q_OBJECT
Q_PROPERTY(QString remote READ remote WRITE setRemote NOTIFY remoteChanged)
Q_PROPERTY(QString transport READ transport WRITE setTransport NOTIFY transportChanged)
Q_PROPERTY(QString endpoint READ endpoint WRITE setEndpoint NOTIFY endpointChanged)
Q_PROPERTY(QString gopath READ gopath WRITE setGopath NOTIFY gopathChanged)
Q_PROPERTY(QString binaryHash READ binaryHash WRITE setBinaryHash NOTIFY binaryHashChanged)
Q_PROPERTY(bool isValid READ isValid WRITE setIsValid NOTIFY isValidChanged)
public:
	LoginCtx77da62(QObject *parent = Q_NULLPTR) : QObject(parent) {qRegisterMetaType<quintptr>("quintptr");LoginCtx77da62_LoginCtx77da62_QRegisterMetaType();LoginCtx77da62_LoginCtx77da62_QRegisterMetaTypes();callbackLoginCtx77da62_Constructor(this);};
	void Signal_Clicked() { callbackLoginCtx77da62_Clicked(this); };
	void Signal_Edited(QString b) { QByteArray te9d71f = b.toUtf8(); Moc_PackedString bPacked = { const_cast<char*>(te9d71f.prepend("WHITESPACE").constData()+10), te9d71f.size()-10 };callbackLoginCtx77da62_Edited(this, bPacked); };
	QString remote() { return ({ Moc_PackedString tempVal = callbackLoginCtx77da62_Remote(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setRemote(QString remote) { QByteArray t41ffe5 = remote.toUtf8(); Moc_PackedString remotePacked = { const_cast<char*>(t41ffe5.prepend("WHITESPACE").constData()+10), t41ffe5.size()-10 };callbackLoginCtx77da62_SetRemote(this, remotePacked); };
	void Signal_RemoteChanged(QString remote) { QByteArray t41ffe5 = remote.toUtf8(); Moc_PackedString remotePacked = { const_cast<char*>(t41ffe5.prepend("WHITESPACE").constData()+10), t41ffe5.size()-10 };callbackLoginCtx77da62_RemoteChanged(this, remotePacked); };
	QString transport() { return ({ Moc_PackedString tempVal = callbackLoginCtx77da62_Transport(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setTransport(QString transport) { QByteArray ta8e601 = transport.toUtf8(); Moc_PackedString transportPacked = { const_cast<char*>(ta8e601.prepend("WHITESPACE").constData()+10), ta8e601.size()-10 };callbackLoginCtx77da62_SetTransport(this, transportPacked); };
	void Signal_TransportChanged(QString transport) { QByteArray ta8e601 = transport.toUtf8(); Moc_PackedString transportPacked = { const_cast<char*>(ta8e601.prepend("WHITESPACE").constData()+10), ta8e601.size()-10 };callbackLoginCtx77da62_TransportChanged(this, transportPacked); };
	QString endpoint() { return ({ Moc_PackedString tempVal = callbackLoginCtx77da62_Endpoint(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setEndpoint(QString endpoint) { QByteArray te13fe4 = endpoint.toUtf8(); Moc_PackedString endpointPacked = { const_cast<char*>(te13fe4.prepend("WHITESPACE").constData()+10), te13fe4.size()-10 };callbackLoginCtx77da62_SetEndpoint(this, endpointPacked); };
	void Signal_EndpointChanged(QString endpoint) { QByteArray te13fe4 = endpoint.toUtf8(); Moc_PackedString endpointPacked = { const_cast<char*>(te13fe4.prepend("WHITESPACE").constData()+10), te13fe4.size()-10 };callbackLoginCtx77da62_EndpointChanged(this, endpointPacked); };
	QString gopath() { return ({ Moc_PackedString tempVal = callbackLoginCtx77da62_Gopath(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setGopath(QString gopath) { QByteArray tdbae25 = gopath.toUtf8(); Moc_PackedString gopathPacked = { const_cast<char*>(tdbae25.prepend("WHITESPACE").constData()+10), tdbae25.size()-10 };callbackLoginCtx77da62_SetGopath(this, gopathPacked); };
	void Signal_GopathChanged(QString gopath) { QByteArray tdbae25 = gopath.toUtf8(); Moc_PackedString gopathPacked = { const_cast<char*>(tdbae25.prepend("WHITESPACE").constData()+10), tdbae25.size()-10 };callbackLoginCtx77da62_GopathChanged(this, gopathPacked); };
	QString binaryHash() { return ({ Moc_PackedString tempVal = callbackLoginCtx77da62_BinaryHash(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setBinaryHash(QString binaryHash) { QByteArray t296c7d = binaryHash.toUtf8(); Moc_PackedString binaryHashPacked = { const_cast<char*>(t296c7d.prepend("WHITESPACE").constData()+10), t296c7d.size()-10 };callbackLoginCtx77da62_SetBinaryHash(this, binaryHashPacked); };
	void Signal_BinaryHashChanged(QString binaryHash) { QByteArray t296c7d = binaryHash.toUtf8(); Moc_PackedString binaryHashPacked = { const_cast<char*>(t296c7d.prepend("WHITESPACE").constData()+10), t296c7d.size()-10 };callbackLoginCtx77da62_BinaryHashChanged(this, binaryHashPacked); };
	bool isValid() { return callbackLoginCtx77da62_IsValid(this) != 0; };
	void setIsValid(bool isValid) { callbackLoginCtx77da62_SetIsValid(this, isValid); };
	void Signal_IsValidChanged(bool isValid) { callbackLoginCtx77da62_IsValidChanged(this, isValid); };
	 ~LoginCtx77da62() { callbackLoginCtx77da62_DestroyLoginCtx(this); };
	bool event(QEvent * e) { return callbackLoginCtx77da62_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackLoginCtx77da62_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackLoginCtx77da62_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackLoginCtx77da62_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackLoginCtx77da62_CustomEvent(this, event); };
	void deleteLater() { callbackLoginCtx77da62_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackLoginCtx77da62_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackLoginCtx77da62_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackLoginCtx77da62_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackLoginCtx77da62_TimerEvent(this, event); };
	QString remoteDefault() { return _remote; };
	void setRemoteDefault(QString p) { if (p != _remote) { _remote = p; remoteChanged(_remote); } };
	QString transportDefault() { return _transport; };
	void setTransportDefault(QString p) { if (p != _transport) { _transport = p; transportChanged(_transport); } };
	QString endpointDefault() { return _endpoint; };
	void setEndpointDefault(QString p) { if (p != _endpoint) { _endpoint = p; endpointChanged(_endpoint); } };
	QString gopathDefault() { return _gopath; };
	void setGopathDefault(QString p) { if (p != _gopath) { _gopath = p; gopathChanged(_gopath); } };
	QString binaryHashDefault() { return _binaryHash; };
	void setBinaryHashDefault(QString p) { if (p != _binaryHash) { _binaryHash = p; binaryHashChanged(_binaryHash); } };
	bool isValidDefault() { return _isValid; };
	void setIsValidDefault(bool p) { if (p != _isValid) { _isValid = p; isValidChanged(_isValid); } };
signals:
	void clicked();
	void edited(QString b);
	void remoteChanged(QString remote);
	void transportChanged(QString transport);
	void endpointChanged(QString endpoint);
	void gopathChanged(QString gopath);
	void binaryHashChanged(QString binaryHash);
	void isValidChanged(bool isValid);
public slots:
private:
	QString _remote;
	QString _transport;
	QString _endpoint;
	QString _gopath;
	QString _binaryHash;
	bool _isValid;
};

Q_DECLARE_METATYPE(LoginCtx77da62*)


void LoginCtx77da62_LoginCtx77da62_QRegisterMetaTypes() {
	qRegisterMetaType<QString>();
}

void ApproveListingCtx77da62_ConnectBack(void* ptr)
{
	QObject::connect(static_cast<ApproveListingCtx77da62*>(ptr), static_cast<void (ApproveListingCtx77da62::*)()>(&ApproveListingCtx77da62::back), static_cast<ApproveListingCtx77da62*>(ptr), static_cast<void (ApproveListingCtx77da62::*)()>(&ApproveListingCtx77da62::Signal_Back));
}

void ApproveListingCtx77da62_DisconnectBack(void* ptr)
{
	QObject::disconnect(static_cast<ApproveListingCtx77da62*>(ptr), static_cast<void (ApproveListingCtx77da62::*)()>(&ApproveListingCtx77da62::back), static_cast<ApproveListingCtx77da62*>(ptr), static_cast<void (ApproveListingCtx77da62::*)()>(&ApproveListingCtx77da62::Signal_Back));
}

void ApproveListingCtx77da62_Back(void* ptr)
{
	static_cast<ApproveListingCtx77da62*>(ptr)->back();
}

void ApproveListingCtx77da62_ConnectApprove(void* ptr)
{
	QObject::connect(static_cast<ApproveListingCtx77da62*>(ptr), static_cast<void (ApproveListingCtx77da62::*)()>(&ApproveListingCtx77da62::approve), static_cast<ApproveListingCtx77da62*>(ptr), static_cast<void (ApproveListingCtx77da62::*)()>(&ApproveListingCtx77da62::Signal_Approve));
}

void ApproveListingCtx77da62_DisconnectApprove(void* ptr)
{
	QObject::disconnect(static_cast<ApproveListingCtx77da62*>(ptr), static_cast<void (ApproveListingCtx77da62::*)()>(&ApproveListingCtx77da62::approve), static_cast<ApproveListingCtx77da62*>(ptr), static_cast<void (ApproveListingCtx77da62::*)()>(&ApproveListingCtx77da62::Signal_Approve));
}

void ApproveListingCtx77da62_Approve(void* ptr)
{
	static_cast<ApproveListingCtx77da62*>(ptr)->approve();
}

void ApproveListingCtx77da62_ConnectReject(void* ptr)
{
	QObject::connect(static_cast<ApproveListingCtx77da62*>(ptr), static_cast<void (ApproveListingCtx77da62::*)()>(&ApproveListingCtx77da62::reject), static_cast<ApproveListingCtx77da62*>(ptr), static_cast<void (ApproveListingCtx77da62::*)()>(&ApproveListingCtx77da62::Signal_Reject));
}

void ApproveListingCtx77da62_DisconnectReject(void* ptr)
{
	QObject::disconnect(static_cast<ApproveListingCtx77da62*>(ptr), static_cast<void (ApproveListingCtx77da62::*)()>(&ApproveListingCtx77da62::reject), static_cast<ApproveListingCtx77da62*>(ptr), static_cast<void (ApproveListingCtx77da62::*)()>(&ApproveListingCtx77da62::Signal_Reject));
}

void ApproveListingCtx77da62_Reject(void* ptr)
{
	static_cast<ApproveListingCtx77da62*>(ptr)->reject();
}

void ApproveListingCtx77da62_ConnectOnCheckStateChanged(void* ptr)
{
	QObject::connect(static_cast<ApproveListingCtx77da62*>(ptr), static_cast<void (ApproveListingCtx77da62::*)(qint32, bool)>(&ApproveListingCtx77da62::onCheckStateChanged), static_cast<ApproveListingCtx77da62*>(ptr), static_cast<void (ApproveListingCtx77da62::*)(qint32, bool)>(&ApproveListingCtx77da62::Signal_OnCheckStateChanged));
}

void ApproveListingCtx77da62_DisconnectOnCheckStateChanged(void* ptr)
{
	QObject::disconnect(static_cast<ApproveListingCtx77da62*>(ptr), static_cast<void (ApproveListingCtx77da62::*)(qint32, bool)>(&ApproveListingCtx77da62::onCheckStateChanged), static_cast<ApproveListingCtx77da62*>(ptr), static_cast<void (ApproveListingCtx77da62::*)(qint32, bool)>(&ApproveListingCtx77da62::Signal_OnCheckStateChanged));
}

void ApproveListingCtx77da62_OnCheckStateChanged(void* ptr, int i, char checked)
{
	static_cast<ApproveListingCtx77da62*>(ptr)->onCheckStateChanged(i, checked != 0);
}

void ApproveListingCtx77da62_ConnectTriggerUpdate(void* ptr)
{
	QObject::connect(static_cast<ApproveListingCtx77da62*>(ptr), static_cast<void (ApproveListingCtx77da62::*)()>(&ApproveListingCtx77da62::triggerUpdate), static_cast<ApproveListingCtx77da62*>(ptr), static_cast<void (ApproveListingCtx77da62::*)()>(&ApproveListingCtx77da62::Signal_TriggerUpdate));
}

void ApproveListingCtx77da62_DisconnectTriggerUpdate(void* ptr)
{
	QObject::disconnect(static_cast<ApproveListingCtx77da62*>(ptr), static_cast<void (ApproveListingCtx77da62::*)()>(&ApproveListingCtx77da62::triggerUpdate), static_cast<ApproveListingCtx77da62*>(ptr), static_cast<void (ApproveListingCtx77da62::*)()>(&ApproveListingCtx77da62::Signal_TriggerUpdate));
}

void ApproveListingCtx77da62_TriggerUpdate(void* ptr)
{
	static_cast<ApproveListingCtx77da62*>(ptr)->triggerUpdate();
}

struct Moc_PackedString ApproveListingCtx77da62_Remote(void* ptr)
{
	return ({ QByteArray t9605d3 = static_cast<ApproveListingCtx77da62*>(ptr)->remote().toUtf8(); Moc_PackedString { const_cast<char*>(t9605d3.prepend("WHITESPACE").constData()+10), t9605d3.size()-10 }; });
}

struct Moc_PackedString ApproveListingCtx77da62_RemoteDefault(void* ptr)
{
	return ({ QByteArray t4b8f13 = static_cast<ApproveListingCtx77da62*>(ptr)->remoteDefault().toUtf8(); Moc_PackedString { const_cast<char*>(t4b8f13.prepend("WHITESPACE").constData()+10), t4b8f13.size()-10 }; });
}

void ApproveListingCtx77da62_SetRemote(void* ptr, struct Moc_PackedString remote)
{
	static_cast<ApproveListingCtx77da62*>(ptr)->setRemote(QString::fromUtf8(remote.data, remote.len));
}

void ApproveListingCtx77da62_SetRemoteDefault(void* ptr, struct Moc_PackedString remote)
{
	static_cast<ApproveListingCtx77da62*>(ptr)->setRemoteDefault(QString::fromUtf8(remote.data, remote.len));
}

void ApproveListingCtx77da62_ConnectRemoteChanged(void* ptr)
{
	QObject::connect(static_cast<ApproveListingCtx77da62*>(ptr), static_cast<void (ApproveListingCtx77da62::*)(QString)>(&ApproveListingCtx77da62::remoteChanged), static_cast<ApproveListingCtx77da62*>(ptr), static_cast<void (ApproveListingCtx77da62::*)(QString)>(&ApproveListingCtx77da62::Signal_RemoteChanged));
}

void ApproveListingCtx77da62_DisconnectRemoteChanged(void* ptr)
{
	QObject::disconnect(static_cast<ApproveListingCtx77da62*>(ptr), static_cast<void (ApproveListingCtx77da62::*)(QString)>(&ApproveListingCtx77da62::remoteChanged), static_cast<ApproveListingCtx77da62*>(ptr), static_cast<void (ApproveListingCtx77da62::*)(QString)>(&ApproveListingCtx77da62::Signal_RemoteChanged));
}

void ApproveListingCtx77da62_RemoteChanged(void* ptr, struct Moc_PackedString remote)
{
	static_cast<ApproveListingCtx77da62*>(ptr)->remoteChanged(QString::fromUtf8(remote.data, remote.len));
}

struct Moc_PackedString ApproveListingCtx77da62_Transport(void* ptr)
{
	return ({ QByteArray t342796 = static_cast<ApproveListingCtx77da62*>(ptr)->transport().toUtf8(); Moc_PackedString { const_cast<char*>(t342796.prepend("WHITESPACE").constData()+10), t342796.size()-10 }; });
}

struct Moc_PackedString ApproveListingCtx77da62_TransportDefault(void* ptr)
{
	return ({ QByteArray t433cbf = static_cast<ApproveListingCtx77da62*>(ptr)->transportDefault().toUtf8(); Moc_PackedString { const_cast<char*>(t433cbf.prepend("WHITESPACE").constData()+10), t433cbf.size()-10 }; });
}

void ApproveListingCtx77da62_SetTransport(void* ptr, struct Moc_PackedString transport)
{
	static_cast<ApproveListingCtx77da62*>(ptr)->setTransport(QString::fromUtf8(transport.data, transport.len));
}

void ApproveListingCtx77da62_SetTransportDefault(void* ptr, struct Moc_PackedString transport)
{
	static_cast<ApproveListingCtx77da62*>(ptr)->setTransportDefault(QString::fromUtf8(transport.data, transport.len));
}

void ApproveListingCtx77da62_ConnectTransportChanged(void* ptr)
{
	QObject::connect(static_cast<ApproveListingCtx77da62*>(ptr), static_cast<void (ApproveListingCtx77da62::*)(QString)>(&ApproveListingCtx77da62::transportChanged), static_cast<ApproveListingCtx77da62*>(ptr), static_cast<void (ApproveListingCtx77da62::*)(QString)>(&ApproveListingCtx77da62::Signal_TransportChanged));
}

void ApproveListingCtx77da62_DisconnectTransportChanged(void* ptr)
{
	QObject::disconnect(static_cast<ApproveListingCtx77da62*>(ptr), static_cast<void (ApproveListingCtx77da62::*)(QString)>(&ApproveListingCtx77da62::transportChanged), static_cast<ApproveListingCtx77da62*>(ptr), static_cast<void (ApproveListingCtx77da62::*)(QString)>(&ApproveListingCtx77da62::Signal_TransportChanged));
}

void ApproveListingCtx77da62_TransportChanged(void* ptr, struct Moc_PackedString transport)
{
	static_cast<ApproveListingCtx77da62*>(ptr)->transportChanged(QString::fromUtf8(transport.data, transport.len));
}

struct Moc_PackedString ApproveListingCtx77da62_Endpoint(void* ptr)
{
	return ({ QByteArray t637b7c = static_cast<ApproveListingCtx77da62*>(ptr)->endpoint().toUtf8(); Moc_PackedString { const_cast<char*>(t637b7c.prepend("WHITESPACE").constData()+10), t637b7c.size()-10 }; });
}

struct Moc_PackedString ApproveListingCtx77da62_EndpointDefault(void* ptr)
{
	return ({ QByteArray t7e670e = static_cast<ApproveListingCtx77da62*>(ptr)->endpointDefault().toUtf8(); Moc_PackedString { const_cast<char*>(t7e670e.prepend("WHITESPACE").constData()+10), t7e670e.size()-10 }; });
}

void ApproveListingCtx77da62_SetEndpoint(void* ptr, struct Moc_PackedString endpoint)
{
	static_cast<ApproveListingCtx77da62*>(ptr)->setEndpoint(QString::fromUtf8(endpoint.data, endpoint.len));
}

void ApproveListingCtx77da62_SetEndpointDefault(void* ptr, struct Moc_PackedString endpoint)
{
	static_cast<ApproveListingCtx77da62*>(ptr)->setEndpointDefault(QString::fromUtf8(endpoint.data, endpoint.len));
}

void ApproveListingCtx77da62_ConnectEndpointChanged(void* ptr)
{
	QObject::connect(static_cast<ApproveListingCtx77da62*>(ptr), static_cast<void (ApproveListingCtx77da62::*)(QString)>(&ApproveListingCtx77da62::endpointChanged), static_cast<ApproveListingCtx77da62*>(ptr), static_cast<void (ApproveListingCtx77da62::*)(QString)>(&ApproveListingCtx77da62::Signal_EndpointChanged));
}

void ApproveListingCtx77da62_DisconnectEndpointChanged(void* ptr)
{
	QObject::disconnect(static_cast<ApproveListingCtx77da62*>(ptr), static_cast<void (ApproveListingCtx77da62::*)(QString)>(&ApproveListingCtx77da62::endpointChanged), static_cast<ApproveListingCtx77da62*>(ptr), static_cast<void (ApproveListingCtx77da62::*)(QString)>(&ApproveListingCtx77da62::Signal_EndpointChanged));
}

void ApproveListingCtx77da62_EndpointChanged(void* ptr, struct Moc_PackedString endpoint)
{
	static_cast<ApproveListingCtx77da62*>(ptr)->endpointChanged(QString::fromUtf8(endpoint.data, endpoint.len));
}

struct Moc_PackedString ApproveListingCtx77da62_From(void* ptr)
{
	return ({ QByteArray t68097f = static_cast<ApproveListingCtx77da62*>(ptr)->from().toUtf8(); Moc_PackedString { const_cast<char*>(t68097f.prepend("WHITESPACE").constData()+10), t68097f.size()-10 }; });
}

struct Moc_PackedString ApproveListingCtx77da62_FromDefault(void* ptr)
{
	return ({ QByteArray t5c8897 = static_cast<ApproveListingCtx77da62*>(ptr)->fromDefault().toUtf8(); Moc_PackedString { const_cast<char*>(t5c8897.prepend("WHITESPACE").constData()+10), t5c8897.size()-10 }; });
}

void ApproveListingCtx77da62_SetFrom(void* ptr, struct Moc_PackedString from)
{
	static_cast<ApproveListingCtx77da62*>(ptr)->setFrom(QString::fromUtf8(from.data, from.len));
}

void ApproveListingCtx77da62_SetFromDefault(void* ptr, struct Moc_PackedString from)
{
	static_cast<ApproveListingCtx77da62*>(ptr)->setFromDefault(QString::fromUtf8(from.data, from.len));
}

void ApproveListingCtx77da62_ConnectFromChanged(void* ptr)
{
	QObject::connect(static_cast<ApproveListingCtx77da62*>(ptr), static_cast<void (ApproveListingCtx77da62::*)(QString)>(&ApproveListingCtx77da62::fromChanged), static_cast<ApproveListingCtx77da62*>(ptr), static_cast<void (ApproveListingCtx77da62::*)(QString)>(&ApproveListingCtx77da62::Signal_FromChanged));
}

void ApproveListingCtx77da62_DisconnectFromChanged(void* ptr)
{
	QObject::disconnect(static_cast<ApproveListingCtx77da62*>(ptr), static_cast<void (ApproveListingCtx77da62::*)(QString)>(&ApproveListingCtx77da62::fromChanged), static_cast<ApproveListingCtx77da62*>(ptr), static_cast<void (ApproveListingCtx77da62::*)(QString)>(&ApproveListingCtx77da62::Signal_FromChanged));
}

void ApproveListingCtx77da62_FromChanged(void* ptr, struct Moc_PackedString from)
{
	static_cast<ApproveListingCtx77da62*>(ptr)->fromChanged(QString::fromUtf8(from.data, from.len));
}

struct Moc_PackedString ApproveListingCtx77da62_Message(void* ptr)
{
	return ({ QByteArray t87291a = static_cast<ApproveListingCtx77da62*>(ptr)->message().toUtf8(); Moc_PackedString { const_cast<char*>(t87291a.prepend("WHITESPACE").constData()+10), t87291a.size()-10 }; });
}

struct Moc_PackedString ApproveListingCtx77da62_MessageDefault(void* ptr)
{
	return ({ QByteArray t124a6d = static_cast<ApproveListingCtx77da62*>(ptr)->messageDefault().toUtf8(); Moc_PackedString { const_cast<char*>(t124a6d.prepend("WHITESPACE").constData()+10), t124a6d.size()-10 }; });
}

void ApproveListingCtx77da62_SetMessage(void* ptr, struct Moc_PackedString message)
{
	static_cast<ApproveListingCtx77da62*>(ptr)->setMessage(QString::fromUtf8(message.data, message.len));
}

void ApproveListingCtx77da62_SetMessageDefault(void* ptr, struct Moc_PackedString message)
{
	static_cast<ApproveListingCtx77da62*>(ptr)->setMessageDefault(QString::fromUtf8(message.data, message.len));
}

void ApproveListingCtx77da62_ConnectMessageChanged(void* ptr)
{
	QObject::connect(static_cast<ApproveListingCtx77da62*>(ptr), static_cast<void (ApproveListingCtx77da62::*)(QString)>(&ApproveListingCtx77da62::messageChanged), static_cast<ApproveListingCtx77da62*>(ptr), static_cast<void (ApproveListingCtx77da62::*)(QString)>(&ApproveListingCtx77da62::Signal_MessageChanged));
}

void ApproveListingCtx77da62_DisconnectMessageChanged(void* ptr)
{
	QObject::disconnect(static_cast<ApproveListingCtx77da62*>(ptr), static_cast<void (ApproveListingCtx77da62::*)(QString)>(&ApproveListingCtx77da62::messageChanged), static_cast<ApproveListingCtx77da62*>(ptr), static_cast<void (ApproveListingCtx77da62::*)(QString)>(&ApproveListingCtx77da62::Signal_MessageChanged));
}

void ApproveListingCtx77da62_MessageChanged(void* ptr, struct Moc_PackedString message)
{
	static_cast<ApproveListingCtx77da62*>(ptr)->messageChanged(QString::fromUtf8(message.data, message.len));
}

struct Moc_PackedString ApproveListingCtx77da62_RawData(void* ptr)
{
	return ({ QByteArray t409f0f = static_cast<ApproveListingCtx77da62*>(ptr)->rawData().toUtf8(); Moc_PackedString { const_cast<char*>(t409f0f.prepend("WHITESPACE").constData()+10), t409f0f.size()-10 }; });
}

struct Moc_PackedString ApproveListingCtx77da62_RawDataDefault(void* ptr)
{
	return ({ QByteArray t518430 = static_cast<ApproveListingCtx77da62*>(ptr)->rawDataDefault().toUtf8(); Moc_PackedString { const_cast<char*>(t518430.prepend("WHITESPACE").constData()+10), t518430.size()-10 }; });
}

void ApproveListingCtx77da62_SetRawData(void* ptr, struct Moc_PackedString rawData)
{
	static_cast<ApproveListingCtx77da62*>(ptr)->setRawData(QString::fromUtf8(rawData.data, rawData.len));
}

void ApproveListingCtx77da62_SetRawDataDefault(void* ptr, struct Moc_PackedString rawData)
{
	static_cast<ApproveListingCtx77da62*>(ptr)->setRawDataDefault(QString::fromUtf8(rawData.data, rawData.len));
}

void ApproveListingCtx77da62_ConnectRawDataChanged(void* ptr)
{
	QObject::connect(static_cast<ApproveListingCtx77da62*>(ptr), static_cast<void (ApproveListingCtx77da62::*)(QString)>(&ApproveListingCtx77da62::rawDataChanged), static_cast<ApproveListingCtx77da62*>(ptr), static_cast<void (ApproveListingCtx77da62::*)(QString)>(&ApproveListingCtx77da62::Signal_RawDataChanged));
}

void ApproveListingCtx77da62_DisconnectRawDataChanged(void* ptr)
{
	QObject::disconnect(static_cast<ApproveListingCtx77da62*>(ptr), static_cast<void (ApproveListingCtx77da62::*)(QString)>(&ApproveListingCtx77da62::rawDataChanged), static_cast<ApproveListingCtx77da62*>(ptr), static_cast<void (ApproveListingCtx77da62::*)(QString)>(&ApproveListingCtx77da62::Signal_RawDataChanged));
}

void ApproveListingCtx77da62_RawDataChanged(void* ptr, struct Moc_PackedString rawData)
{
	static_cast<ApproveListingCtx77da62*>(ptr)->rawDataChanged(QString::fromUtf8(rawData.data, rawData.len));
}

struct Moc_PackedString ApproveListingCtx77da62_Hash(void* ptr)
{
	return ({ QByteArray t74e19d = static_cast<ApproveListingCtx77da62*>(ptr)->hash().toUtf8(); Moc_PackedString { const_cast<char*>(t74e19d.prepend("WHITESPACE").constData()+10), t74e19d.size()-10 }; });
}

struct Moc_PackedString ApproveListingCtx77da62_HashDefault(void* ptr)
{
	return ({ QByteArray td3c517 = static_cast<ApproveListingCtx77da62*>(ptr)->hashDefault().toUtf8(); Moc_PackedString { const_cast<char*>(td3c517.prepend("WHITESPACE").constData()+10), td3c517.size()-10 }; });
}

void ApproveListingCtx77da62_SetHash(void* ptr, struct Moc_PackedString hash)
{
	static_cast<ApproveListingCtx77da62*>(ptr)->setHash(QString::fromUtf8(hash.data, hash.len));
}

void ApproveListingCtx77da62_SetHashDefault(void* ptr, struct Moc_PackedString hash)
{
	static_cast<ApproveListingCtx77da62*>(ptr)->setHashDefault(QString::fromUtf8(hash.data, hash.len));
}

void ApproveListingCtx77da62_ConnectHashChanged(void* ptr)
{
	QObject::connect(static_cast<ApproveListingCtx77da62*>(ptr), static_cast<void (ApproveListingCtx77da62::*)(QString)>(&ApproveListingCtx77da62::hashChanged), static_cast<ApproveListingCtx77da62*>(ptr), static_cast<void (ApproveListingCtx77da62::*)(QString)>(&ApproveListingCtx77da62::Signal_HashChanged));
}

void ApproveListingCtx77da62_DisconnectHashChanged(void* ptr)
{
	QObject::disconnect(static_cast<ApproveListingCtx77da62*>(ptr), static_cast<void (ApproveListingCtx77da62::*)(QString)>(&ApproveListingCtx77da62::hashChanged), static_cast<ApproveListingCtx77da62*>(ptr), static_cast<void (ApproveListingCtx77da62::*)(QString)>(&ApproveListingCtx77da62::Signal_HashChanged));
}

void ApproveListingCtx77da62_HashChanged(void* ptr, struct Moc_PackedString hash)
{
	static_cast<ApproveListingCtx77da62*>(ptr)->hashChanged(QString::fromUtf8(hash.data, hash.len));
}

int ApproveListingCtx77da62_ApproveListingCtx77da62_QRegisterMetaType()
{
	return qRegisterMetaType<ApproveListingCtx77da62*>();
}

int ApproveListingCtx77da62_ApproveListingCtx77da62_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<ApproveListingCtx77da62*>(const_cast<const char*>(typeName));
}

int ApproveListingCtx77da62_ApproveListingCtx77da62_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<ApproveListingCtx77da62>();
#else
	return 0;
#endif
}

int ApproveListingCtx77da62_ApproveListingCtx77da62_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<ApproveListingCtx77da62>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

void* ApproveListingCtx77da62___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void ApproveListingCtx77da62___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* ApproveListingCtx77da62___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* ApproveListingCtx77da62___findChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void ApproveListingCtx77da62___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* ApproveListingCtx77da62___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* ApproveListingCtx77da62___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void ApproveListingCtx77da62___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* ApproveListingCtx77da62___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* ApproveListingCtx77da62___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void ApproveListingCtx77da62___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* ApproveListingCtx77da62___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* ApproveListingCtx77da62___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void ApproveListingCtx77da62___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* ApproveListingCtx77da62___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* ApproveListingCtx77da62_NewApproveListingCtx(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new ApproveListingCtx77da62(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new ApproveListingCtx77da62(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new ApproveListingCtx77da62(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new ApproveListingCtx77da62(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new ApproveListingCtx77da62(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new ApproveListingCtx77da62(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new ApproveListingCtx77da62(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new ApproveListingCtx77da62(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new ApproveListingCtx77da62(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new ApproveListingCtx77da62(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new ApproveListingCtx77da62(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new ApproveListingCtx77da62(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new ApproveListingCtx77da62(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new ApproveListingCtx77da62(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new ApproveListingCtx77da62(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new ApproveListingCtx77da62(static_cast<QWindow*>(parent));
	} else {
		return new ApproveListingCtx77da62(static_cast<QObject*>(parent));
	}
}

void ApproveListingCtx77da62_DestroyApproveListingCtx(void* ptr)
{
	static_cast<ApproveListingCtx77da62*>(ptr)->~ApproveListingCtx77da62();
}

void ApproveListingCtx77da62_DestroyApproveListingCtxDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

char ApproveListingCtx77da62_EventDefault(void* ptr, void* e)
{
	return static_cast<ApproveListingCtx77da62*>(ptr)->QObject::event(static_cast<QEvent*>(e));
}

char ApproveListingCtx77da62_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<ApproveListingCtx77da62*>(ptr)->QObject::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void ApproveListingCtx77da62_ChildEventDefault(void* ptr, void* event)
{
	static_cast<ApproveListingCtx77da62*>(ptr)->QObject::childEvent(static_cast<QChildEvent*>(event));
}

void ApproveListingCtx77da62_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<ApproveListingCtx77da62*>(ptr)->QObject::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void ApproveListingCtx77da62_CustomEventDefault(void* ptr, void* event)
{
	static_cast<ApproveListingCtx77da62*>(ptr)->QObject::customEvent(static_cast<QEvent*>(event));
}

void ApproveListingCtx77da62_DeleteLaterDefault(void* ptr)
{
	static_cast<ApproveListingCtx77da62*>(ptr)->QObject::deleteLater();
}

void ApproveListingCtx77da62_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<ApproveListingCtx77da62*>(ptr)->QObject::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void ApproveListingCtx77da62_TimerEventDefault(void* ptr, void* event)
{
	static_cast<ApproveListingCtx77da62*>(ptr)->QObject::timerEvent(static_cast<QTimerEvent*>(event));
}



void ApproveNewAccountCtx77da62_ConnectClicked(void* ptr)
{
	QObject::connect(static_cast<ApproveNewAccountCtx77da62*>(ptr), static_cast<void (ApproveNewAccountCtx77da62::*)(qint32)>(&ApproveNewAccountCtx77da62::clicked), static_cast<ApproveNewAccountCtx77da62*>(ptr), static_cast<void (ApproveNewAccountCtx77da62::*)(qint32)>(&ApproveNewAccountCtx77da62::Signal_Clicked));
}

void ApproveNewAccountCtx77da62_DisconnectClicked(void* ptr)
{
	QObject::disconnect(static_cast<ApproveNewAccountCtx77da62*>(ptr), static_cast<void (ApproveNewAccountCtx77da62::*)(qint32)>(&ApproveNewAccountCtx77da62::clicked), static_cast<ApproveNewAccountCtx77da62*>(ptr), static_cast<void (ApproveNewAccountCtx77da62::*)(qint32)>(&ApproveNewAccountCtx77da62::Signal_Clicked));
}

void ApproveNewAccountCtx77da62_Clicked(void* ptr, int b)
{
	static_cast<ApproveNewAccountCtx77da62*>(ptr)->clicked(b);
}

void ApproveNewAccountCtx77da62_ConnectBack(void* ptr)
{
	QObject::connect(static_cast<ApproveNewAccountCtx77da62*>(ptr), static_cast<void (ApproveNewAccountCtx77da62::*)()>(&ApproveNewAccountCtx77da62::back), static_cast<ApproveNewAccountCtx77da62*>(ptr), static_cast<void (ApproveNewAccountCtx77da62::*)()>(&ApproveNewAccountCtx77da62::Signal_Back));
}

void ApproveNewAccountCtx77da62_DisconnectBack(void* ptr)
{
	QObject::disconnect(static_cast<ApproveNewAccountCtx77da62*>(ptr), static_cast<void (ApproveNewAccountCtx77da62::*)()>(&ApproveNewAccountCtx77da62::back), static_cast<ApproveNewAccountCtx77da62*>(ptr), static_cast<void (ApproveNewAccountCtx77da62::*)()>(&ApproveNewAccountCtx77da62::Signal_Back));
}

void ApproveNewAccountCtx77da62_Back(void* ptr)
{
	static_cast<ApproveNewAccountCtx77da62*>(ptr)->back();
}

void ApproveNewAccountCtx77da62_ConnectPasswordEdited(void* ptr)
{
	QObject::connect(static_cast<ApproveNewAccountCtx77da62*>(ptr), static_cast<void (ApproveNewAccountCtx77da62::*)(QString)>(&ApproveNewAccountCtx77da62::passwordEdited), static_cast<ApproveNewAccountCtx77da62*>(ptr), static_cast<void (ApproveNewAccountCtx77da62::*)(QString)>(&ApproveNewAccountCtx77da62::Signal_PasswordEdited));
}

void ApproveNewAccountCtx77da62_DisconnectPasswordEdited(void* ptr)
{
	QObject::disconnect(static_cast<ApproveNewAccountCtx77da62*>(ptr), static_cast<void (ApproveNewAccountCtx77da62::*)(QString)>(&ApproveNewAccountCtx77da62::passwordEdited), static_cast<ApproveNewAccountCtx77da62*>(ptr), static_cast<void (ApproveNewAccountCtx77da62::*)(QString)>(&ApproveNewAccountCtx77da62::Signal_PasswordEdited));
}

void ApproveNewAccountCtx77da62_PasswordEdited(void* ptr, struct Moc_PackedString b)
{
	static_cast<ApproveNewAccountCtx77da62*>(ptr)->passwordEdited(QString::fromUtf8(b.data, b.len));
}

void ApproveNewAccountCtx77da62_ConnectConfirmPasswordEdited(void* ptr)
{
	QObject::connect(static_cast<ApproveNewAccountCtx77da62*>(ptr), static_cast<void (ApproveNewAccountCtx77da62::*)(QString)>(&ApproveNewAccountCtx77da62::confirmPasswordEdited), static_cast<ApproveNewAccountCtx77da62*>(ptr), static_cast<void (ApproveNewAccountCtx77da62::*)(QString)>(&ApproveNewAccountCtx77da62::Signal_ConfirmPasswordEdited));
}

void ApproveNewAccountCtx77da62_DisconnectConfirmPasswordEdited(void* ptr)
{
	QObject::disconnect(static_cast<ApproveNewAccountCtx77da62*>(ptr), static_cast<void (ApproveNewAccountCtx77da62::*)(QString)>(&ApproveNewAccountCtx77da62::confirmPasswordEdited), static_cast<ApproveNewAccountCtx77da62*>(ptr), static_cast<void (ApproveNewAccountCtx77da62::*)(QString)>(&ApproveNewAccountCtx77da62::Signal_ConfirmPasswordEdited));
}

void ApproveNewAccountCtx77da62_ConfirmPasswordEdited(void* ptr, struct Moc_PackedString b)
{
	static_cast<ApproveNewAccountCtx77da62*>(ptr)->confirmPasswordEdited(QString::fromUtf8(b.data, b.len));
}

struct Moc_PackedString ApproveNewAccountCtx77da62_Remote(void* ptr)
{
	return ({ QByteArray t3094cf = static_cast<ApproveNewAccountCtx77da62*>(ptr)->remote().toUtf8(); Moc_PackedString { const_cast<char*>(t3094cf.prepend("WHITESPACE").constData()+10), t3094cf.size()-10 }; });
}

struct Moc_PackedString ApproveNewAccountCtx77da62_RemoteDefault(void* ptr)
{
	return ({ QByteArray t63aa2c = static_cast<ApproveNewAccountCtx77da62*>(ptr)->remoteDefault().toUtf8(); Moc_PackedString { const_cast<char*>(t63aa2c.prepend("WHITESPACE").constData()+10), t63aa2c.size()-10 }; });
}

void ApproveNewAccountCtx77da62_SetRemote(void* ptr, struct Moc_PackedString remote)
{
	static_cast<ApproveNewAccountCtx77da62*>(ptr)->setRemote(QString::fromUtf8(remote.data, remote.len));
}

void ApproveNewAccountCtx77da62_SetRemoteDefault(void* ptr, struct Moc_PackedString remote)
{
	static_cast<ApproveNewAccountCtx77da62*>(ptr)->setRemoteDefault(QString::fromUtf8(remote.data, remote.len));
}

void ApproveNewAccountCtx77da62_ConnectRemoteChanged(void* ptr)
{
	QObject::connect(static_cast<ApproveNewAccountCtx77da62*>(ptr), static_cast<void (ApproveNewAccountCtx77da62::*)(QString)>(&ApproveNewAccountCtx77da62::remoteChanged), static_cast<ApproveNewAccountCtx77da62*>(ptr), static_cast<void (ApproveNewAccountCtx77da62::*)(QString)>(&ApproveNewAccountCtx77da62::Signal_RemoteChanged));
}

void ApproveNewAccountCtx77da62_DisconnectRemoteChanged(void* ptr)
{
	QObject::disconnect(static_cast<ApproveNewAccountCtx77da62*>(ptr), static_cast<void (ApproveNewAccountCtx77da62::*)(QString)>(&ApproveNewAccountCtx77da62::remoteChanged), static_cast<ApproveNewAccountCtx77da62*>(ptr), static_cast<void (ApproveNewAccountCtx77da62::*)(QString)>(&ApproveNewAccountCtx77da62::Signal_RemoteChanged));
}

void ApproveNewAccountCtx77da62_RemoteChanged(void* ptr, struct Moc_PackedString remote)
{
	static_cast<ApproveNewAccountCtx77da62*>(ptr)->remoteChanged(QString::fromUtf8(remote.data, remote.len));
}

struct Moc_PackedString ApproveNewAccountCtx77da62_Transport(void* ptr)
{
	return ({ QByteArray t98cdda = static_cast<ApproveNewAccountCtx77da62*>(ptr)->transport().toUtf8(); Moc_PackedString { const_cast<char*>(t98cdda.prepend("WHITESPACE").constData()+10), t98cdda.size()-10 }; });
}

struct Moc_PackedString ApproveNewAccountCtx77da62_TransportDefault(void* ptr)
{
	return ({ QByteArray t73eeff = static_cast<ApproveNewAccountCtx77da62*>(ptr)->transportDefault().toUtf8(); Moc_PackedString { const_cast<char*>(t73eeff.prepend("WHITESPACE").constData()+10), t73eeff.size()-10 }; });
}

void ApproveNewAccountCtx77da62_SetTransport(void* ptr, struct Moc_PackedString transport)
{
	static_cast<ApproveNewAccountCtx77da62*>(ptr)->setTransport(QString::fromUtf8(transport.data, transport.len));
}

void ApproveNewAccountCtx77da62_SetTransportDefault(void* ptr, struct Moc_PackedString transport)
{
	static_cast<ApproveNewAccountCtx77da62*>(ptr)->setTransportDefault(QString::fromUtf8(transport.data, transport.len));
}

void ApproveNewAccountCtx77da62_ConnectTransportChanged(void* ptr)
{
	QObject::connect(static_cast<ApproveNewAccountCtx77da62*>(ptr), static_cast<void (ApproveNewAccountCtx77da62::*)(QString)>(&ApproveNewAccountCtx77da62::transportChanged), static_cast<ApproveNewAccountCtx77da62*>(ptr), static_cast<void (ApproveNewAccountCtx77da62::*)(QString)>(&ApproveNewAccountCtx77da62::Signal_TransportChanged));
}

void ApproveNewAccountCtx77da62_DisconnectTransportChanged(void* ptr)
{
	QObject::disconnect(static_cast<ApproveNewAccountCtx77da62*>(ptr), static_cast<void (ApproveNewAccountCtx77da62::*)(QString)>(&ApproveNewAccountCtx77da62::transportChanged), static_cast<ApproveNewAccountCtx77da62*>(ptr), static_cast<void (ApproveNewAccountCtx77da62::*)(QString)>(&ApproveNewAccountCtx77da62::Signal_TransportChanged));
}

void ApproveNewAccountCtx77da62_TransportChanged(void* ptr, struct Moc_PackedString transport)
{
	static_cast<ApproveNewAccountCtx77da62*>(ptr)->transportChanged(QString::fromUtf8(transport.data, transport.len));
}

struct Moc_PackedString ApproveNewAccountCtx77da62_Endpoint(void* ptr)
{
	return ({ QByteArray t7d2dd9 = static_cast<ApproveNewAccountCtx77da62*>(ptr)->endpoint().toUtf8(); Moc_PackedString { const_cast<char*>(t7d2dd9.prepend("WHITESPACE").constData()+10), t7d2dd9.size()-10 }; });
}

struct Moc_PackedString ApproveNewAccountCtx77da62_EndpointDefault(void* ptr)
{
	return ({ QByteArray t43ee71 = static_cast<ApproveNewAccountCtx77da62*>(ptr)->endpointDefault().toUtf8(); Moc_PackedString { const_cast<char*>(t43ee71.prepend("WHITESPACE").constData()+10), t43ee71.size()-10 }; });
}

void ApproveNewAccountCtx77da62_SetEndpoint(void* ptr, struct Moc_PackedString endpoint)
{
	static_cast<ApproveNewAccountCtx77da62*>(ptr)->setEndpoint(QString::fromUtf8(endpoint.data, endpoint.len));
}

void ApproveNewAccountCtx77da62_SetEndpointDefault(void* ptr, struct Moc_PackedString endpoint)
{
	static_cast<ApproveNewAccountCtx77da62*>(ptr)->setEndpointDefault(QString::fromUtf8(endpoint.data, endpoint.len));
}

void ApproveNewAccountCtx77da62_ConnectEndpointChanged(void* ptr)
{
	QObject::connect(static_cast<ApproveNewAccountCtx77da62*>(ptr), static_cast<void (ApproveNewAccountCtx77da62::*)(QString)>(&ApproveNewAccountCtx77da62::endpointChanged), static_cast<ApproveNewAccountCtx77da62*>(ptr), static_cast<void (ApproveNewAccountCtx77da62::*)(QString)>(&ApproveNewAccountCtx77da62::Signal_EndpointChanged));
}

void ApproveNewAccountCtx77da62_DisconnectEndpointChanged(void* ptr)
{
	QObject::disconnect(static_cast<ApproveNewAccountCtx77da62*>(ptr), static_cast<void (ApproveNewAccountCtx77da62::*)(QString)>(&ApproveNewAccountCtx77da62::endpointChanged), static_cast<ApproveNewAccountCtx77da62*>(ptr), static_cast<void (ApproveNewAccountCtx77da62::*)(QString)>(&ApproveNewAccountCtx77da62::Signal_EndpointChanged));
}

void ApproveNewAccountCtx77da62_EndpointChanged(void* ptr, struct Moc_PackedString endpoint)
{
	static_cast<ApproveNewAccountCtx77da62*>(ptr)->endpointChanged(QString::fromUtf8(endpoint.data, endpoint.len));
}

struct Moc_PackedString ApproveNewAccountCtx77da62_Password(void* ptr)
{
	return ({ QByteArray tcbaaee = static_cast<ApproveNewAccountCtx77da62*>(ptr)->password().toUtf8(); Moc_PackedString { const_cast<char*>(tcbaaee.prepend("WHITESPACE").constData()+10), tcbaaee.size()-10 }; });
}

struct Moc_PackedString ApproveNewAccountCtx77da62_PasswordDefault(void* ptr)
{
	return ({ QByteArray t2050de = static_cast<ApproveNewAccountCtx77da62*>(ptr)->passwordDefault().toUtf8(); Moc_PackedString { const_cast<char*>(t2050de.prepend("WHITESPACE").constData()+10), t2050de.size()-10 }; });
}

void ApproveNewAccountCtx77da62_SetPassword(void* ptr, struct Moc_PackedString password)
{
	static_cast<ApproveNewAccountCtx77da62*>(ptr)->setPassword(QString::fromUtf8(password.data, password.len));
}

void ApproveNewAccountCtx77da62_SetPasswordDefault(void* ptr, struct Moc_PackedString password)
{
	static_cast<ApproveNewAccountCtx77da62*>(ptr)->setPasswordDefault(QString::fromUtf8(password.data, password.len));
}

void ApproveNewAccountCtx77da62_ConnectPasswordChanged(void* ptr)
{
	QObject::connect(static_cast<ApproveNewAccountCtx77da62*>(ptr), static_cast<void (ApproveNewAccountCtx77da62::*)(QString)>(&ApproveNewAccountCtx77da62::passwordChanged), static_cast<ApproveNewAccountCtx77da62*>(ptr), static_cast<void (ApproveNewAccountCtx77da62::*)(QString)>(&ApproveNewAccountCtx77da62::Signal_PasswordChanged));
}

void ApproveNewAccountCtx77da62_DisconnectPasswordChanged(void* ptr)
{
	QObject::disconnect(static_cast<ApproveNewAccountCtx77da62*>(ptr), static_cast<void (ApproveNewAccountCtx77da62::*)(QString)>(&ApproveNewAccountCtx77da62::passwordChanged), static_cast<ApproveNewAccountCtx77da62*>(ptr), static_cast<void (ApproveNewAccountCtx77da62::*)(QString)>(&ApproveNewAccountCtx77da62::Signal_PasswordChanged));
}

void ApproveNewAccountCtx77da62_PasswordChanged(void* ptr, struct Moc_PackedString password)
{
	static_cast<ApproveNewAccountCtx77da62*>(ptr)->passwordChanged(QString::fromUtf8(password.data, password.len));
}

struct Moc_PackedString ApproveNewAccountCtx77da62_ConfirmPassword(void* ptr)
{
	return ({ QByteArray t47d33e = static_cast<ApproveNewAccountCtx77da62*>(ptr)->confirmPassword().toUtf8(); Moc_PackedString { const_cast<char*>(t47d33e.prepend("WHITESPACE").constData()+10), t47d33e.size()-10 }; });
}

struct Moc_PackedString ApproveNewAccountCtx77da62_ConfirmPasswordDefault(void* ptr)
{
	return ({ QByteArray te8e3d4 = static_cast<ApproveNewAccountCtx77da62*>(ptr)->confirmPasswordDefault().toUtf8(); Moc_PackedString { const_cast<char*>(te8e3d4.prepend("WHITESPACE").constData()+10), te8e3d4.size()-10 }; });
}

void ApproveNewAccountCtx77da62_SetConfirmPassword(void* ptr, struct Moc_PackedString confirmPassword)
{
	static_cast<ApproveNewAccountCtx77da62*>(ptr)->setConfirmPassword(QString::fromUtf8(confirmPassword.data, confirmPassword.len));
}

void ApproveNewAccountCtx77da62_SetConfirmPasswordDefault(void* ptr, struct Moc_PackedString confirmPassword)
{
	static_cast<ApproveNewAccountCtx77da62*>(ptr)->setConfirmPasswordDefault(QString::fromUtf8(confirmPassword.data, confirmPassword.len));
}

void ApproveNewAccountCtx77da62_ConnectConfirmPasswordChanged(void* ptr)
{
	QObject::connect(static_cast<ApproveNewAccountCtx77da62*>(ptr), static_cast<void (ApproveNewAccountCtx77da62::*)(QString)>(&ApproveNewAccountCtx77da62::confirmPasswordChanged), static_cast<ApproveNewAccountCtx77da62*>(ptr), static_cast<void (ApproveNewAccountCtx77da62::*)(QString)>(&ApproveNewAccountCtx77da62::Signal_ConfirmPasswordChanged));
}

void ApproveNewAccountCtx77da62_DisconnectConfirmPasswordChanged(void* ptr)
{
	QObject::disconnect(static_cast<ApproveNewAccountCtx77da62*>(ptr), static_cast<void (ApproveNewAccountCtx77da62::*)(QString)>(&ApproveNewAccountCtx77da62::confirmPasswordChanged), static_cast<ApproveNewAccountCtx77da62*>(ptr), static_cast<void (ApproveNewAccountCtx77da62::*)(QString)>(&ApproveNewAccountCtx77da62::Signal_ConfirmPasswordChanged));
}

void ApproveNewAccountCtx77da62_ConfirmPasswordChanged(void* ptr, struct Moc_PackedString confirmPassword)
{
	static_cast<ApproveNewAccountCtx77da62*>(ptr)->confirmPasswordChanged(QString::fromUtf8(confirmPassword.data, confirmPassword.len));
}

int ApproveNewAccountCtx77da62_ApproveNewAccountCtx77da62_QRegisterMetaType()
{
	return qRegisterMetaType<ApproveNewAccountCtx77da62*>();
}

int ApproveNewAccountCtx77da62_ApproveNewAccountCtx77da62_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<ApproveNewAccountCtx77da62*>(const_cast<const char*>(typeName));
}

int ApproveNewAccountCtx77da62_ApproveNewAccountCtx77da62_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<ApproveNewAccountCtx77da62>();
#else
	return 0;
#endif
}

int ApproveNewAccountCtx77da62_ApproveNewAccountCtx77da62_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<ApproveNewAccountCtx77da62>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

void* ApproveNewAccountCtx77da62___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void ApproveNewAccountCtx77da62___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* ApproveNewAccountCtx77da62___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* ApproveNewAccountCtx77da62___findChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void ApproveNewAccountCtx77da62___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* ApproveNewAccountCtx77da62___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* ApproveNewAccountCtx77da62___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void ApproveNewAccountCtx77da62___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* ApproveNewAccountCtx77da62___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* ApproveNewAccountCtx77da62___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void ApproveNewAccountCtx77da62___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* ApproveNewAccountCtx77da62___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* ApproveNewAccountCtx77da62___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void ApproveNewAccountCtx77da62___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* ApproveNewAccountCtx77da62___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* ApproveNewAccountCtx77da62_NewApproveNewAccountCtx(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new ApproveNewAccountCtx77da62(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new ApproveNewAccountCtx77da62(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new ApproveNewAccountCtx77da62(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new ApproveNewAccountCtx77da62(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new ApproveNewAccountCtx77da62(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new ApproveNewAccountCtx77da62(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new ApproveNewAccountCtx77da62(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new ApproveNewAccountCtx77da62(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new ApproveNewAccountCtx77da62(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new ApproveNewAccountCtx77da62(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new ApproveNewAccountCtx77da62(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new ApproveNewAccountCtx77da62(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new ApproveNewAccountCtx77da62(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new ApproveNewAccountCtx77da62(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new ApproveNewAccountCtx77da62(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new ApproveNewAccountCtx77da62(static_cast<QWindow*>(parent));
	} else {
		return new ApproveNewAccountCtx77da62(static_cast<QObject*>(parent));
	}
}

void ApproveNewAccountCtx77da62_DestroyApproveNewAccountCtx(void* ptr)
{
	static_cast<ApproveNewAccountCtx77da62*>(ptr)->~ApproveNewAccountCtx77da62();
}

void ApproveNewAccountCtx77da62_DestroyApproveNewAccountCtxDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

char ApproveNewAccountCtx77da62_EventDefault(void* ptr, void* e)
{
	return static_cast<ApproveNewAccountCtx77da62*>(ptr)->QObject::event(static_cast<QEvent*>(e));
}

char ApproveNewAccountCtx77da62_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<ApproveNewAccountCtx77da62*>(ptr)->QObject::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void ApproveNewAccountCtx77da62_ChildEventDefault(void* ptr, void* event)
{
	static_cast<ApproveNewAccountCtx77da62*>(ptr)->QObject::childEvent(static_cast<QChildEvent*>(event));
}

void ApproveNewAccountCtx77da62_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<ApproveNewAccountCtx77da62*>(ptr)->QObject::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void ApproveNewAccountCtx77da62_CustomEventDefault(void* ptr, void* event)
{
	static_cast<ApproveNewAccountCtx77da62*>(ptr)->QObject::customEvent(static_cast<QEvent*>(event));
}

void ApproveNewAccountCtx77da62_DeleteLaterDefault(void* ptr)
{
	static_cast<ApproveNewAccountCtx77da62*>(ptr)->QObject::deleteLater();
}

void ApproveNewAccountCtx77da62_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<ApproveNewAccountCtx77da62*>(ptr)->QObject::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void ApproveNewAccountCtx77da62_TimerEventDefault(void* ptr, void* event)
{
	static_cast<ApproveNewAccountCtx77da62*>(ptr)->QObject::timerEvent(static_cast<QTimerEvent*>(event));
}



void ApproveTxCtx77da62_ConnectApprove(void* ptr)
{
	QObject::connect(static_cast<ApproveTxCtx77da62*>(ptr), static_cast<void (ApproveTxCtx77da62::*)()>(&ApproveTxCtx77da62::approve), static_cast<ApproveTxCtx77da62*>(ptr), static_cast<void (ApproveTxCtx77da62::*)()>(&ApproveTxCtx77da62::Signal_Approve));
}

void ApproveTxCtx77da62_DisconnectApprove(void* ptr)
{
	QObject::disconnect(static_cast<ApproveTxCtx77da62*>(ptr), static_cast<void (ApproveTxCtx77da62::*)()>(&ApproveTxCtx77da62::approve), static_cast<ApproveTxCtx77da62*>(ptr), static_cast<void (ApproveTxCtx77da62::*)()>(&ApproveTxCtx77da62::Signal_Approve));
}

void ApproveTxCtx77da62_Approve(void* ptr)
{
	static_cast<ApproveTxCtx77da62*>(ptr)->approve();
}

void ApproveTxCtx77da62_ConnectReject(void* ptr)
{
	QObject::connect(static_cast<ApproveTxCtx77da62*>(ptr), static_cast<void (ApproveTxCtx77da62::*)()>(&ApproveTxCtx77da62::reject), static_cast<ApproveTxCtx77da62*>(ptr), static_cast<void (ApproveTxCtx77da62::*)()>(&ApproveTxCtx77da62::Signal_Reject));
}

void ApproveTxCtx77da62_DisconnectReject(void* ptr)
{
	QObject::disconnect(static_cast<ApproveTxCtx77da62*>(ptr), static_cast<void (ApproveTxCtx77da62::*)()>(&ApproveTxCtx77da62::reject), static_cast<ApproveTxCtx77da62*>(ptr), static_cast<void (ApproveTxCtx77da62::*)()>(&ApproveTxCtx77da62::Signal_Reject));
}

void ApproveTxCtx77da62_Reject(void* ptr)
{
	static_cast<ApproveTxCtx77da62*>(ptr)->reject();
}

void ApproveTxCtx77da62_ConnectCheckTxDiff(void* ptr)
{
	QObject::connect(static_cast<ApproveTxCtx77da62*>(ptr), static_cast<void (ApproveTxCtx77da62::*)()>(&ApproveTxCtx77da62::checkTxDiff), static_cast<ApproveTxCtx77da62*>(ptr), static_cast<void (ApproveTxCtx77da62::*)()>(&ApproveTxCtx77da62::Signal_CheckTxDiff));
}

void ApproveTxCtx77da62_DisconnectCheckTxDiff(void* ptr)
{
	QObject::disconnect(static_cast<ApproveTxCtx77da62*>(ptr), static_cast<void (ApproveTxCtx77da62::*)()>(&ApproveTxCtx77da62::checkTxDiff), static_cast<ApproveTxCtx77da62*>(ptr), static_cast<void (ApproveTxCtx77da62::*)()>(&ApproveTxCtx77da62::Signal_CheckTxDiff));
}

void ApproveTxCtx77da62_CheckTxDiff(void* ptr)
{
	static_cast<ApproveTxCtx77da62*>(ptr)->checkTxDiff();
}

void ApproveTxCtx77da62_ConnectBack(void* ptr)
{
	QObject::connect(static_cast<ApproveTxCtx77da62*>(ptr), static_cast<void (ApproveTxCtx77da62::*)()>(&ApproveTxCtx77da62::back), static_cast<ApproveTxCtx77da62*>(ptr), static_cast<void (ApproveTxCtx77da62::*)()>(&ApproveTxCtx77da62::Signal_Back));
}

void ApproveTxCtx77da62_DisconnectBack(void* ptr)
{
	QObject::disconnect(static_cast<ApproveTxCtx77da62*>(ptr), static_cast<void (ApproveTxCtx77da62::*)()>(&ApproveTxCtx77da62::back), static_cast<ApproveTxCtx77da62*>(ptr), static_cast<void (ApproveTxCtx77da62::*)()>(&ApproveTxCtx77da62::Signal_Back));
}

void ApproveTxCtx77da62_Back(void* ptr)
{
	static_cast<ApproveTxCtx77da62*>(ptr)->back();
}

void ApproveTxCtx77da62_ConnectEdited(void* ptr)
{
	QObject::connect(static_cast<ApproveTxCtx77da62*>(ptr), static_cast<void (ApproveTxCtx77da62::*)(QString, QString)>(&ApproveTxCtx77da62::edited), static_cast<ApproveTxCtx77da62*>(ptr), static_cast<void (ApproveTxCtx77da62::*)(QString, QString)>(&ApproveTxCtx77da62::Signal_Edited));
}

void ApproveTxCtx77da62_DisconnectEdited(void* ptr)
{
	QObject::disconnect(static_cast<ApproveTxCtx77da62*>(ptr), static_cast<void (ApproveTxCtx77da62::*)(QString, QString)>(&ApproveTxCtx77da62::edited), static_cast<ApproveTxCtx77da62*>(ptr), static_cast<void (ApproveTxCtx77da62::*)(QString, QString)>(&ApproveTxCtx77da62::Signal_Edited));
}

void ApproveTxCtx77da62_Edited(void* ptr, struct Moc_PackedString s, struct Moc_PackedString v)
{
	static_cast<ApproveTxCtx77da62*>(ptr)->edited(QString::fromUtf8(s.data, s.len), QString::fromUtf8(v.data, v.len));
}

void ApproveTxCtx77da62_ConnectChangeValueUnit(void* ptr)
{
	QObject::connect(static_cast<ApproveTxCtx77da62*>(ptr), static_cast<void (ApproveTxCtx77da62::*)(qint32)>(&ApproveTxCtx77da62::changeValueUnit), static_cast<ApproveTxCtx77da62*>(ptr), static_cast<void (ApproveTxCtx77da62::*)(qint32)>(&ApproveTxCtx77da62::Signal_ChangeValueUnit));
}

void ApproveTxCtx77da62_DisconnectChangeValueUnit(void* ptr)
{
	QObject::disconnect(static_cast<ApproveTxCtx77da62*>(ptr), static_cast<void (ApproveTxCtx77da62::*)(qint32)>(&ApproveTxCtx77da62::changeValueUnit), static_cast<ApproveTxCtx77da62*>(ptr), static_cast<void (ApproveTxCtx77da62::*)(qint32)>(&ApproveTxCtx77da62::Signal_ChangeValueUnit));
}

void ApproveTxCtx77da62_ChangeValueUnit(void* ptr, int v)
{
	static_cast<ApproveTxCtx77da62*>(ptr)->changeValueUnit(v);
}

void ApproveTxCtx77da62_ConnectChangeGasPriceUnit(void* ptr)
{
	QObject::connect(static_cast<ApproveTxCtx77da62*>(ptr), static_cast<void (ApproveTxCtx77da62::*)(qint32)>(&ApproveTxCtx77da62::changeGasPriceUnit), static_cast<ApproveTxCtx77da62*>(ptr), static_cast<void (ApproveTxCtx77da62::*)(qint32)>(&ApproveTxCtx77da62::Signal_ChangeGasPriceUnit));
}

void ApproveTxCtx77da62_DisconnectChangeGasPriceUnit(void* ptr)
{
	QObject::disconnect(static_cast<ApproveTxCtx77da62*>(ptr), static_cast<void (ApproveTxCtx77da62::*)(qint32)>(&ApproveTxCtx77da62::changeGasPriceUnit), static_cast<ApproveTxCtx77da62*>(ptr), static_cast<void (ApproveTxCtx77da62::*)(qint32)>(&ApproveTxCtx77da62::Signal_ChangeGasPriceUnit));
}

void ApproveTxCtx77da62_ChangeGasPriceUnit(void* ptr, int v)
{
	static_cast<ApproveTxCtx77da62*>(ptr)->changeGasPriceUnit(v);
}

int ApproveTxCtx77da62_ValueUnit(void* ptr)
{
	return static_cast<ApproveTxCtx77da62*>(ptr)->valueUnit();
}

int ApproveTxCtx77da62_ValueUnitDefault(void* ptr)
{
	return static_cast<ApproveTxCtx77da62*>(ptr)->valueUnitDefault();
}

void ApproveTxCtx77da62_SetValueUnit(void* ptr, int valueUnit)
{
	static_cast<ApproveTxCtx77da62*>(ptr)->setValueUnit(valueUnit);
}

void ApproveTxCtx77da62_SetValueUnitDefault(void* ptr, int valueUnit)
{
	static_cast<ApproveTxCtx77da62*>(ptr)->setValueUnitDefault(valueUnit);
}

void ApproveTxCtx77da62_ConnectValueUnitChanged(void* ptr)
{
	QObject::connect(static_cast<ApproveTxCtx77da62*>(ptr), static_cast<void (ApproveTxCtx77da62::*)(qint32)>(&ApproveTxCtx77da62::valueUnitChanged), static_cast<ApproveTxCtx77da62*>(ptr), static_cast<void (ApproveTxCtx77da62::*)(qint32)>(&ApproveTxCtx77da62::Signal_ValueUnitChanged));
}

void ApproveTxCtx77da62_DisconnectValueUnitChanged(void* ptr)
{
	QObject::disconnect(static_cast<ApproveTxCtx77da62*>(ptr), static_cast<void (ApproveTxCtx77da62::*)(qint32)>(&ApproveTxCtx77da62::valueUnitChanged), static_cast<ApproveTxCtx77da62*>(ptr), static_cast<void (ApproveTxCtx77da62::*)(qint32)>(&ApproveTxCtx77da62::Signal_ValueUnitChanged));
}

void ApproveTxCtx77da62_ValueUnitChanged(void* ptr, int valueUnit)
{
	static_cast<ApproveTxCtx77da62*>(ptr)->valueUnitChanged(valueUnit);
}

struct Moc_PackedString ApproveTxCtx77da62_Remote(void* ptr)
{
	return ({ QByteArray t90f7f1 = static_cast<ApproveTxCtx77da62*>(ptr)->remote().toUtf8(); Moc_PackedString { const_cast<char*>(t90f7f1.prepend("WHITESPACE").constData()+10), t90f7f1.size()-10 }; });
}

struct Moc_PackedString ApproveTxCtx77da62_RemoteDefault(void* ptr)
{
	return ({ QByteArray t5fca0a = static_cast<ApproveTxCtx77da62*>(ptr)->remoteDefault().toUtf8(); Moc_PackedString { const_cast<char*>(t5fca0a.prepend("WHITESPACE").constData()+10), t5fca0a.size()-10 }; });
}

void ApproveTxCtx77da62_SetRemote(void* ptr, struct Moc_PackedString remote)
{
	static_cast<ApproveTxCtx77da62*>(ptr)->setRemote(QString::fromUtf8(remote.data, remote.len));
}

void ApproveTxCtx77da62_SetRemoteDefault(void* ptr, struct Moc_PackedString remote)
{
	static_cast<ApproveTxCtx77da62*>(ptr)->setRemoteDefault(QString::fromUtf8(remote.data, remote.len));
}

void ApproveTxCtx77da62_ConnectRemoteChanged(void* ptr)
{
	QObject::connect(static_cast<ApproveTxCtx77da62*>(ptr), static_cast<void (ApproveTxCtx77da62::*)(QString)>(&ApproveTxCtx77da62::remoteChanged), static_cast<ApproveTxCtx77da62*>(ptr), static_cast<void (ApproveTxCtx77da62::*)(QString)>(&ApproveTxCtx77da62::Signal_RemoteChanged));
}

void ApproveTxCtx77da62_DisconnectRemoteChanged(void* ptr)
{
	QObject::disconnect(static_cast<ApproveTxCtx77da62*>(ptr), static_cast<void (ApproveTxCtx77da62::*)(QString)>(&ApproveTxCtx77da62::remoteChanged), static_cast<ApproveTxCtx77da62*>(ptr), static_cast<void (ApproveTxCtx77da62::*)(QString)>(&ApproveTxCtx77da62::Signal_RemoteChanged));
}

void ApproveTxCtx77da62_RemoteChanged(void* ptr, struct Moc_PackedString remote)
{
	static_cast<ApproveTxCtx77da62*>(ptr)->remoteChanged(QString::fromUtf8(remote.data, remote.len));
}

struct Moc_PackedString ApproveTxCtx77da62_Transport(void* ptr)
{
	return ({ QByteArray te6d44f = static_cast<ApproveTxCtx77da62*>(ptr)->transport().toUtf8(); Moc_PackedString { const_cast<char*>(te6d44f.prepend("WHITESPACE").constData()+10), te6d44f.size()-10 }; });
}

struct Moc_PackedString ApproveTxCtx77da62_TransportDefault(void* ptr)
{
	return ({ QByteArray t09c1b3 = static_cast<ApproveTxCtx77da62*>(ptr)->transportDefault().toUtf8(); Moc_PackedString { const_cast<char*>(t09c1b3.prepend("WHITESPACE").constData()+10), t09c1b3.size()-10 }; });
}

void ApproveTxCtx77da62_SetTransport(void* ptr, struct Moc_PackedString transport)
{
	static_cast<ApproveTxCtx77da62*>(ptr)->setTransport(QString::fromUtf8(transport.data, transport.len));
}

void ApproveTxCtx77da62_SetTransportDefault(void* ptr, struct Moc_PackedString transport)
{
	static_cast<ApproveTxCtx77da62*>(ptr)->setTransportDefault(QString::fromUtf8(transport.data, transport.len));
}

void ApproveTxCtx77da62_ConnectTransportChanged(void* ptr)
{
	QObject::connect(static_cast<ApproveTxCtx77da62*>(ptr), static_cast<void (ApproveTxCtx77da62::*)(QString)>(&ApproveTxCtx77da62::transportChanged), static_cast<ApproveTxCtx77da62*>(ptr), static_cast<void (ApproveTxCtx77da62::*)(QString)>(&ApproveTxCtx77da62::Signal_TransportChanged));
}

void ApproveTxCtx77da62_DisconnectTransportChanged(void* ptr)
{
	QObject::disconnect(static_cast<ApproveTxCtx77da62*>(ptr), static_cast<void (ApproveTxCtx77da62::*)(QString)>(&ApproveTxCtx77da62::transportChanged), static_cast<ApproveTxCtx77da62*>(ptr), static_cast<void (ApproveTxCtx77da62::*)(QString)>(&ApproveTxCtx77da62::Signal_TransportChanged));
}

void ApproveTxCtx77da62_TransportChanged(void* ptr, struct Moc_PackedString transport)
{
	static_cast<ApproveTxCtx77da62*>(ptr)->transportChanged(QString::fromUtf8(transport.data, transport.len));
}

struct Moc_PackedString ApproveTxCtx77da62_Endpoint(void* ptr)
{
	return ({ QByteArray taa8e81 = static_cast<ApproveTxCtx77da62*>(ptr)->endpoint().toUtf8(); Moc_PackedString { const_cast<char*>(taa8e81.prepend("WHITESPACE").constData()+10), taa8e81.size()-10 }; });
}

struct Moc_PackedString ApproveTxCtx77da62_EndpointDefault(void* ptr)
{
	return ({ QByteArray t2b1328 = static_cast<ApproveTxCtx77da62*>(ptr)->endpointDefault().toUtf8(); Moc_PackedString { const_cast<char*>(t2b1328.prepend("WHITESPACE").constData()+10), t2b1328.size()-10 }; });
}

void ApproveTxCtx77da62_SetEndpoint(void* ptr, struct Moc_PackedString endpoint)
{
	static_cast<ApproveTxCtx77da62*>(ptr)->setEndpoint(QString::fromUtf8(endpoint.data, endpoint.len));
}

void ApproveTxCtx77da62_SetEndpointDefault(void* ptr, struct Moc_PackedString endpoint)
{
	static_cast<ApproveTxCtx77da62*>(ptr)->setEndpointDefault(QString::fromUtf8(endpoint.data, endpoint.len));
}

void ApproveTxCtx77da62_ConnectEndpointChanged(void* ptr)
{
	QObject::connect(static_cast<ApproveTxCtx77da62*>(ptr), static_cast<void (ApproveTxCtx77da62::*)(QString)>(&ApproveTxCtx77da62::endpointChanged), static_cast<ApproveTxCtx77da62*>(ptr), static_cast<void (ApproveTxCtx77da62::*)(QString)>(&ApproveTxCtx77da62::Signal_EndpointChanged));
}

void ApproveTxCtx77da62_DisconnectEndpointChanged(void* ptr)
{
	QObject::disconnect(static_cast<ApproveTxCtx77da62*>(ptr), static_cast<void (ApproveTxCtx77da62::*)(QString)>(&ApproveTxCtx77da62::endpointChanged), static_cast<ApproveTxCtx77da62*>(ptr), static_cast<void (ApproveTxCtx77da62::*)(QString)>(&ApproveTxCtx77da62::Signal_EndpointChanged));
}

void ApproveTxCtx77da62_EndpointChanged(void* ptr, struct Moc_PackedString endpoint)
{
	static_cast<ApproveTxCtx77da62*>(ptr)->endpointChanged(QString::fromUtf8(endpoint.data, endpoint.len));
}

struct Moc_PackedString ApproveTxCtx77da62_Data(void* ptr)
{
	return ({ QByteArray t8072b7 = static_cast<ApproveTxCtx77da62*>(ptr)->data().toUtf8(); Moc_PackedString { const_cast<char*>(t8072b7.prepend("WHITESPACE").constData()+10), t8072b7.size()-10 }; });
}

struct Moc_PackedString ApproveTxCtx77da62_DataDefault(void* ptr)
{
	return ({ QByteArray t04fe35 = static_cast<ApproveTxCtx77da62*>(ptr)->dataDefault().toUtf8(); Moc_PackedString { const_cast<char*>(t04fe35.prepend("WHITESPACE").constData()+10), t04fe35.size()-10 }; });
}

void ApproveTxCtx77da62_SetData(void* ptr, struct Moc_PackedString data)
{
	static_cast<ApproveTxCtx77da62*>(ptr)->setData(QString::fromUtf8(data.data, data.len));
}

void ApproveTxCtx77da62_SetDataDefault(void* ptr, struct Moc_PackedString data)
{
	static_cast<ApproveTxCtx77da62*>(ptr)->setDataDefault(QString::fromUtf8(data.data, data.len));
}

void ApproveTxCtx77da62_ConnectDataChanged(void* ptr)
{
	QObject::connect(static_cast<ApproveTxCtx77da62*>(ptr), static_cast<void (ApproveTxCtx77da62::*)(QString)>(&ApproveTxCtx77da62::dataChanged), static_cast<ApproveTxCtx77da62*>(ptr), static_cast<void (ApproveTxCtx77da62::*)(QString)>(&ApproveTxCtx77da62::Signal_DataChanged));
}

void ApproveTxCtx77da62_DisconnectDataChanged(void* ptr)
{
	QObject::disconnect(static_cast<ApproveTxCtx77da62*>(ptr), static_cast<void (ApproveTxCtx77da62::*)(QString)>(&ApproveTxCtx77da62::dataChanged), static_cast<ApproveTxCtx77da62*>(ptr), static_cast<void (ApproveTxCtx77da62::*)(QString)>(&ApproveTxCtx77da62::Signal_DataChanged));
}

void ApproveTxCtx77da62_DataChanged(void* ptr, struct Moc_PackedString data)
{
	static_cast<ApproveTxCtx77da62*>(ptr)->dataChanged(QString::fromUtf8(data.data, data.len));
}

struct Moc_PackedString ApproveTxCtx77da62_From(void* ptr)
{
	return ({ QByteArray tc27f9b = static_cast<ApproveTxCtx77da62*>(ptr)->from().toUtf8(); Moc_PackedString { const_cast<char*>(tc27f9b.prepend("WHITESPACE").constData()+10), tc27f9b.size()-10 }; });
}

struct Moc_PackedString ApproveTxCtx77da62_FromDefault(void* ptr)
{
	return ({ QByteArray t805f9b = static_cast<ApproveTxCtx77da62*>(ptr)->fromDefault().toUtf8(); Moc_PackedString { const_cast<char*>(t805f9b.prepend("WHITESPACE").constData()+10), t805f9b.size()-10 }; });
}

void ApproveTxCtx77da62_SetFrom(void* ptr, struct Moc_PackedString from)
{
	static_cast<ApproveTxCtx77da62*>(ptr)->setFrom(QString::fromUtf8(from.data, from.len));
}

void ApproveTxCtx77da62_SetFromDefault(void* ptr, struct Moc_PackedString from)
{
	static_cast<ApproveTxCtx77da62*>(ptr)->setFromDefault(QString::fromUtf8(from.data, from.len));
}

void ApproveTxCtx77da62_ConnectFromChanged(void* ptr)
{
	QObject::connect(static_cast<ApproveTxCtx77da62*>(ptr), static_cast<void (ApproveTxCtx77da62::*)(QString)>(&ApproveTxCtx77da62::fromChanged), static_cast<ApproveTxCtx77da62*>(ptr), static_cast<void (ApproveTxCtx77da62::*)(QString)>(&ApproveTxCtx77da62::Signal_FromChanged));
}

void ApproveTxCtx77da62_DisconnectFromChanged(void* ptr)
{
	QObject::disconnect(static_cast<ApproveTxCtx77da62*>(ptr), static_cast<void (ApproveTxCtx77da62::*)(QString)>(&ApproveTxCtx77da62::fromChanged), static_cast<ApproveTxCtx77da62*>(ptr), static_cast<void (ApproveTxCtx77da62::*)(QString)>(&ApproveTxCtx77da62::Signal_FromChanged));
}

void ApproveTxCtx77da62_FromChanged(void* ptr, struct Moc_PackedString from)
{
	static_cast<ApproveTxCtx77da62*>(ptr)->fromChanged(QString::fromUtf8(from.data, from.len));
}

struct Moc_PackedString ApproveTxCtx77da62_FromWarning(void* ptr)
{
	return ({ QByteArray td56e55 = static_cast<ApproveTxCtx77da62*>(ptr)->fromWarning().toUtf8(); Moc_PackedString { const_cast<char*>(td56e55.prepend("WHITESPACE").constData()+10), td56e55.size()-10 }; });
}

struct Moc_PackedString ApproveTxCtx77da62_FromWarningDefault(void* ptr)
{
	return ({ QByteArray t2020a1 = static_cast<ApproveTxCtx77da62*>(ptr)->fromWarningDefault().toUtf8(); Moc_PackedString { const_cast<char*>(t2020a1.prepend("WHITESPACE").constData()+10), t2020a1.size()-10 }; });
}

void ApproveTxCtx77da62_SetFromWarning(void* ptr, struct Moc_PackedString fromWarning)
{
	static_cast<ApproveTxCtx77da62*>(ptr)->setFromWarning(QString::fromUtf8(fromWarning.data, fromWarning.len));
}

void ApproveTxCtx77da62_SetFromWarningDefault(void* ptr, struct Moc_PackedString fromWarning)
{
	static_cast<ApproveTxCtx77da62*>(ptr)->setFromWarningDefault(QString::fromUtf8(fromWarning.data, fromWarning.len));
}

void ApproveTxCtx77da62_ConnectFromWarningChanged(void* ptr)
{
	QObject::connect(static_cast<ApproveTxCtx77da62*>(ptr), static_cast<void (ApproveTxCtx77da62::*)(QString)>(&ApproveTxCtx77da62::fromWarningChanged), static_cast<ApproveTxCtx77da62*>(ptr), static_cast<void (ApproveTxCtx77da62::*)(QString)>(&ApproveTxCtx77da62::Signal_FromWarningChanged));
}

void ApproveTxCtx77da62_DisconnectFromWarningChanged(void* ptr)
{
	QObject::disconnect(static_cast<ApproveTxCtx77da62*>(ptr), static_cast<void (ApproveTxCtx77da62::*)(QString)>(&ApproveTxCtx77da62::fromWarningChanged), static_cast<ApproveTxCtx77da62*>(ptr), static_cast<void (ApproveTxCtx77da62::*)(QString)>(&ApproveTxCtx77da62::Signal_FromWarningChanged));
}

void ApproveTxCtx77da62_FromWarningChanged(void* ptr, struct Moc_PackedString fromWarning)
{
	static_cast<ApproveTxCtx77da62*>(ptr)->fromWarningChanged(QString::fromUtf8(fromWarning.data, fromWarning.len));
}

char ApproveTxCtx77da62_IsFromVisible(void* ptr)
{
	return static_cast<ApproveTxCtx77da62*>(ptr)->isFromVisible();
}

char ApproveTxCtx77da62_IsFromVisibleDefault(void* ptr)
{
	return static_cast<ApproveTxCtx77da62*>(ptr)->isFromVisibleDefault();
}

void ApproveTxCtx77da62_SetFromVisible(void* ptr, char fromVisible)
{
	static_cast<ApproveTxCtx77da62*>(ptr)->setFromVisible(fromVisible != 0);
}

void ApproveTxCtx77da62_SetFromVisibleDefault(void* ptr, char fromVisible)
{
	static_cast<ApproveTxCtx77da62*>(ptr)->setFromVisibleDefault(fromVisible != 0);
}

void ApproveTxCtx77da62_ConnectFromVisibleChanged(void* ptr)
{
	QObject::connect(static_cast<ApproveTxCtx77da62*>(ptr), static_cast<void (ApproveTxCtx77da62::*)(bool)>(&ApproveTxCtx77da62::fromVisibleChanged), static_cast<ApproveTxCtx77da62*>(ptr), static_cast<void (ApproveTxCtx77da62::*)(bool)>(&ApproveTxCtx77da62::Signal_FromVisibleChanged));
}

void ApproveTxCtx77da62_DisconnectFromVisibleChanged(void* ptr)
{
	QObject::disconnect(static_cast<ApproveTxCtx77da62*>(ptr), static_cast<void (ApproveTxCtx77da62::*)(bool)>(&ApproveTxCtx77da62::fromVisibleChanged), static_cast<ApproveTxCtx77da62*>(ptr), static_cast<void (ApproveTxCtx77da62::*)(bool)>(&ApproveTxCtx77da62::Signal_FromVisibleChanged));
}

void ApproveTxCtx77da62_FromVisibleChanged(void* ptr, char fromVisible)
{
	static_cast<ApproveTxCtx77da62*>(ptr)->fromVisibleChanged(fromVisible != 0);
}

struct Moc_PackedString ApproveTxCtx77da62_To(void* ptr)
{
	return ({ QByteArray t6e4b06 = static_cast<ApproveTxCtx77da62*>(ptr)->to().toUtf8(); Moc_PackedString { const_cast<char*>(t6e4b06.prepend("WHITESPACE").constData()+10), t6e4b06.size()-10 }; });
}

struct Moc_PackedString ApproveTxCtx77da62_ToDefault(void* ptr)
{
	return ({ QByteArray t7ceec8 = static_cast<ApproveTxCtx77da62*>(ptr)->toDefault().toUtf8(); Moc_PackedString { const_cast<char*>(t7ceec8.prepend("WHITESPACE").constData()+10), t7ceec8.size()-10 }; });
}

void ApproveTxCtx77da62_SetTo(void* ptr, struct Moc_PackedString to)
{
	static_cast<ApproveTxCtx77da62*>(ptr)->setTo(QString::fromUtf8(to.data, to.len));
}

void ApproveTxCtx77da62_SetToDefault(void* ptr, struct Moc_PackedString to)
{
	static_cast<ApproveTxCtx77da62*>(ptr)->setToDefault(QString::fromUtf8(to.data, to.len));
}

void ApproveTxCtx77da62_ConnectToChanged(void* ptr)
{
	QObject::connect(static_cast<ApproveTxCtx77da62*>(ptr), static_cast<void (ApproveTxCtx77da62::*)(QString)>(&ApproveTxCtx77da62::toChanged), static_cast<ApproveTxCtx77da62*>(ptr), static_cast<void (ApproveTxCtx77da62::*)(QString)>(&ApproveTxCtx77da62::Signal_ToChanged));
}

void ApproveTxCtx77da62_DisconnectToChanged(void* ptr)
{
	QObject::disconnect(static_cast<ApproveTxCtx77da62*>(ptr), static_cast<void (ApproveTxCtx77da62::*)(QString)>(&ApproveTxCtx77da62::toChanged), static_cast<ApproveTxCtx77da62*>(ptr), static_cast<void (ApproveTxCtx77da62::*)(QString)>(&ApproveTxCtx77da62::Signal_ToChanged));
}

void ApproveTxCtx77da62_ToChanged(void* ptr, struct Moc_PackedString to)
{
	static_cast<ApproveTxCtx77da62*>(ptr)->toChanged(QString::fromUtf8(to.data, to.len));
}

struct Moc_PackedString ApproveTxCtx77da62_ToWarning(void* ptr)
{
	return ({ QByteArray t7006fd = static_cast<ApproveTxCtx77da62*>(ptr)->toWarning().toUtf8(); Moc_PackedString { const_cast<char*>(t7006fd.prepend("WHITESPACE").constData()+10), t7006fd.size()-10 }; });
}

struct Moc_PackedString ApproveTxCtx77da62_ToWarningDefault(void* ptr)
{
	return ({ QByteArray t523608 = static_cast<ApproveTxCtx77da62*>(ptr)->toWarningDefault().toUtf8(); Moc_PackedString { const_cast<char*>(t523608.prepend("WHITESPACE").constData()+10), t523608.size()-10 }; });
}

void ApproveTxCtx77da62_SetToWarning(void* ptr, struct Moc_PackedString toWarning)
{
	static_cast<ApproveTxCtx77da62*>(ptr)->setToWarning(QString::fromUtf8(toWarning.data, toWarning.len));
}

void ApproveTxCtx77da62_SetToWarningDefault(void* ptr, struct Moc_PackedString toWarning)
{
	static_cast<ApproveTxCtx77da62*>(ptr)->setToWarningDefault(QString::fromUtf8(toWarning.data, toWarning.len));
}

void ApproveTxCtx77da62_ConnectToWarningChanged(void* ptr)
{
	QObject::connect(static_cast<ApproveTxCtx77da62*>(ptr), static_cast<void (ApproveTxCtx77da62::*)(QString)>(&ApproveTxCtx77da62::toWarningChanged), static_cast<ApproveTxCtx77da62*>(ptr), static_cast<void (ApproveTxCtx77da62::*)(QString)>(&ApproveTxCtx77da62::Signal_ToWarningChanged));
}

void ApproveTxCtx77da62_DisconnectToWarningChanged(void* ptr)
{
	QObject::disconnect(static_cast<ApproveTxCtx77da62*>(ptr), static_cast<void (ApproveTxCtx77da62::*)(QString)>(&ApproveTxCtx77da62::toWarningChanged), static_cast<ApproveTxCtx77da62*>(ptr), static_cast<void (ApproveTxCtx77da62::*)(QString)>(&ApproveTxCtx77da62::Signal_ToWarningChanged));
}

void ApproveTxCtx77da62_ToWarningChanged(void* ptr, struct Moc_PackedString toWarning)
{
	static_cast<ApproveTxCtx77da62*>(ptr)->toWarningChanged(QString::fromUtf8(toWarning.data, toWarning.len));
}

char ApproveTxCtx77da62_IsToVisible(void* ptr)
{
	return static_cast<ApproveTxCtx77da62*>(ptr)->isToVisible();
}

char ApproveTxCtx77da62_IsToVisibleDefault(void* ptr)
{
	return static_cast<ApproveTxCtx77da62*>(ptr)->isToVisibleDefault();
}

void ApproveTxCtx77da62_SetToVisible(void* ptr, char toVisible)
{
	static_cast<ApproveTxCtx77da62*>(ptr)->setToVisible(toVisible != 0);
}

void ApproveTxCtx77da62_SetToVisibleDefault(void* ptr, char toVisible)
{
	static_cast<ApproveTxCtx77da62*>(ptr)->setToVisibleDefault(toVisible != 0);
}

void ApproveTxCtx77da62_ConnectToVisibleChanged(void* ptr)
{
	QObject::connect(static_cast<ApproveTxCtx77da62*>(ptr), static_cast<void (ApproveTxCtx77da62::*)(bool)>(&ApproveTxCtx77da62::toVisibleChanged), static_cast<ApproveTxCtx77da62*>(ptr), static_cast<void (ApproveTxCtx77da62::*)(bool)>(&ApproveTxCtx77da62::Signal_ToVisibleChanged));
}

void ApproveTxCtx77da62_DisconnectToVisibleChanged(void* ptr)
{
	QObject::disconnect(static_cast<ApproveTxCtx77da62*>(ptr), static_cast<void (ApproveTxCtx77da62::*)(bool)>(&ApproveTxCtx77da62::toVisibleChanged), static_cast<ApproveTxCtx77da62*>(ptr), static_cast<void (ApproveTxCtx77da62::*)(bool)>(&ApproveTxCtx77da62::Signal_ToVisibleChanged));
}

void ApproveTxCtx77da62_ToVisibleChanged(void* ptr, char toVisible)
{
	static_cast<ApproveTxCtx77da62*>(ptr)->toVisibleChanged(toVisible != 0);
}

struct Moc_PackedString ApproveTxCtx77da62_Gas(void* ptr)
{
	return ({ QByteArray t3b252a = static_cast<ApproveTxCtx77da62*>(ptr)->gas().toUtf8(); Moc_PackedString { const_cast<char*>(t3b252a.prepend("WHITESPACE").constData()+10), t3b252a.size()-10 }; });
}

struct Moc_PackedString ApproveTxCtx77da62_GasDefault(void* ptr)
{
	return ({ QByteArray tdee409 = static_cast<ApproveTxCtx77da62*>(ptr)->gasDefault().toUtf8(); Moc_PackedString { const_cast<char*>(tdee409.prepend("WHITESPACE").constData()+10), tdee409.size()-10 }; });
}

void ApproveTxCtx77da62_SetGas(void* ptr, struct Moc_PackedString gas)
{
	static_cast<ApproveTxCtx77da62*>(ptr)->setGas(QString::fromUtf8(gas.data, gas.len));
}

void ApproveTxCtx77da62_SetGasDefault(void* ptr, struct Moc_PackedString gas)
{
	static_cast<ApproveTxCtx77da62*>(ptr)->setGasDefault(QString::fromUtf8(gas.data, gas.len));
}

void ApproveTxCtx77da62_ConnectGasChanged(void* ptr)
{
	QObject::connect(static_cast<ApproveTxCtx77da62*>(ptr), static_cast<void (ApproveTxCtx77da62::*)(QString)>(&ApproveTxCtx77da62::gasChanged), static_cast<ApproveTxCtx77da62*>(ptr), static_cast<void (ApproveTxCtx77da62::*)(QString)>(&ApproveTxCtx77da62::Signal_GasChanged));
}

void ApproveTxCtx77da62_DisconnectGasChanged(void* ptr)
{
	QObject::disconnect(static_cast<ApproveTxCtx77da62*>(ptr), static_cast<void (ApproveTxCtx77da62::*)(QString)>(&ApproveTxCtx77da62::gasChanged), static_cast<ApproveTxCtx77da62*>(ptr), static_cast<void (ApproveTxCtx77da62::*)(QString)>(&ApproveTxCtx77da62::Signal_GasChanged));
}

void ApproveTxCtx77da62_GasChanged(void* ptr, struct Moc_PackedString gas)
{
	static_cast<ApproveTxCtx77da62*>(ptr)->gasChanged(QString::fromUtf8(gas.data, gas.len));
}

struct Moc_PackedString ApproveTxCtx77da62_GasPrice(void* ptr)
{
	return ({ QByteArray te19d41 = static_cast<ApproveTxCtx77da62*>(ptr)->gasPrice().toUtf8(); Moc_PackedString { const_cast<char*>(te19d41.prepend("WHITESPACE").constData()+10), te19d41.size()-10 }; });
}

struct Moc_PackedString ApproveTxCtx77da62_GasPriceDefault(void* ptr)
{
	return ({ QByteArray tef8e7f = static_cast<ApproveTxCtx77da62*>(ptr)->gasPriceDefault().toUtf8(); Moc_PackedString { const_cast<char*>(tef8e7f.prepend("WHITESPACE").constData()+10), tef8e7f.size()-10 }; });
}

void ApproveTxCtx77da62_SetGasPrice(void* ptr, struct Moc_PackedString gasPrice)
{
	static_cast<ApproveTxCtx77da62*>(ptr)->setGasPrice(QString::fromUtf8(gasPrice.data, gasPrice.len));
}

void ApproveTxCtx77da62_SetGasPriceDefault(void* ptr, struct Moc_PackedString gasPrice)
{
	static_cast<ApproveTxCtx77da62*>(ptr)->setGasPriceDefault(QString::fromUtf8(gasPrice.data, gasPrice.len));
}

void ApproveTxCtx77da62_ConnectGasPriceChanged(void* ptr)
{
	QObject::connect(static_cast<ApproveTxCtx77da62*>(ptr), static_cast<void (ApproveTxCtx77da62::*)(QString)>(&ApproveTxCtx77da62::gasPriceChanged), static_cast<ApproveTxCtx77da62*>(ptr), static_cast<void (ApproveTxCtx77da62::*)(QString)>(&ApproveTxCtx77da62::Signal_GasPriceChanged));
}

void ApproveTxCtx77da62_DisconnectGasPriceChanged(void* ptr)
{
	QObject::disconnect(static_cast<ApproveTxCtx77da62*>(ptr), static_cast<void (ApproveTxCtx77da62::*)(QString)>(&ApproveTxCtx77da62::gasPriceChanged), static_cast<ApproveTxCtx77da62*>(ptr), static_cast<void (ApproveTxCtx77da62::*)(QString)>(&ApproveTxCtx77da62::Signal_GasPriceChanged));
}

void ApproveTxCtx77da62_GasPriceChanged(void* ptr, struct Moc_PackedString gasPrice)
{
	static_cast<ApproveTxCtx77da62*>(ptr)->gasPriceChanged(QString::fromUtf8(gasPrice.data, gasPrice.len));
}

int ApproveTxCtx77da62_GasPriceUnit(void* ptr)
{
	return static_cast<ApproveTxCtx77da62*>(ptr)->gasPriceUnit();
}

int ApproveTxCtx77da62_GasPriceUnitDefault(void* ptr)
{
	return static_cast<ApproveTxCtx77da62*>(ptr)->gasPriceUnitDefault();
}

void ApproveTxCtx77da62_SetGasPriceUnit(void* ptr, int gasPriceUnit)
{
	static_cast<ApproveTxCtx77da62*>(ptr)->setGasPriceUnit(gasPriceUnit);
}

void ApproveTxCtx77da62_SetGasPriceUnitDefault(void* ptr, int gasPriceUnit)
{
	static_cast<ApproveTxCtx77da62*>(ptr)->setGasPriceUnitDefault(gasPriceUnit);
}

void ApproveTxCtx77da62_ConnectGasPriceUnitChanged(void* ptr)
{
	QObject::connect(static_cast<ApproveTxCtx77da62*>(ptr), static_cast<void (ApproveTxCtx77da62::*)(qint32)>(&ApproveTxCtx77da62::gasPriceUnitChanged), static_cast<ApproveTxCtx77da62*>(ptr), static_cast<void (ApproveTxCtx77da62::*)(qint32)>(&ApproveTxCtx77da62::Signal_GasPriceUnitChanged));
}

void ApproveTxCtx77da62_DisconnectGasPriceUnitChanged(void* ptr)
{
	QObject::disconnect(static_cast<ApproveTxCtx77da62*>(ptr), static_cast<void (ApproveTxCtx77da62::*)(qint32)>(&ApproveTxCtx77da62::gasPriceUnitChanged), static_cast<ApproveTxCtx77da62*>(ptr), static_cast<void (ApproveTxCtx77da62::*)(qint32)>(&ApproveTxCtx77da62::Signal_GasPriceUnitChanged));
}

void ApproveTxCtx77da62_GasPriceUnitChanged(void* ptr, int gasPriceUnit)
{
	static_cast<ApproveTxCtx77da62*>(ptr)->gasPriceUnitChanged(gasPriceUnit);
}

struct Moc_PackedString ApproveTxCtx77da62_Nonce(void* ptr)
{
	return ({ QByteArray t532f58 = static_cast<ApproveTxCtx77da62*>(ptr)->nonce().toUtf8(); Moc_PackedString { const_cast<char*>(t532f58.prepend("WHITESPACE").constData()+10), t532f58.size()-10 }; });
}

struct Moc_PackedString ApproveTxCtx77da62_NonceDefault(void* ptr)
{
	return ({ QByteArray tf1f93d = static_cast<ApproveTxCtx77da62*>(ptr)->nonceDefault().toUtf8(); Moc_PackedString { const_cast<char*>(tf1f93d.prepend("WHITESPACE").constData()+10), tf1f93d.size()-10 }; });
}

void ApproveTxCtx77da62_SetNonce(void* ptr, struct Moc_PackedString nonce)
{
	static_cast<ApproveTxCtx77da62*>(ptr)->setNonce(QString::fromUtf8(nonce.data, nonce.len));
}

void ApproveTxCtx77da62_SetNonceDefault(void* ptr, struct Moc_PackedString nonce)
{
	static_cast<ApproveTxCtx77da62*>(ptr)->setNonceDefault(QString::fromUtf8(nonce.data, nonce.len));
}

void ApproveTxCtx77da62_ConnectNonceChanged(void* ptr)
{
	QObject::connect(static_cast<ApproveTxCtx77da62*>(ptr), static_cast<void (ApproveTxCtx77da62::*)(QString)>(&ApproveTxCtx77da62::nonceChanged), static_cast<ApproveTxCtx77da62*>(ptr), static_cast<void (ApproveTxCtx77da62::*)(QString)>(&ApproveTxCtx77da62::Signal_NonceChanged));
}

void ApproveTxCtx77da62_DisconnectNonceChanged(void* ptr)
{
	QObject::disconnect(static_cast<ApproveTxCtx77da62*>(ptr), static_cast<void (ApproveTxCtx77da62::*)(QString)>(&ApproveTxCtx77da62::nonceChanged), static_cast<ApproveTxCtx77da62*>(ptr), static_cast<void (ApproveTxCtx77da62::*)(QString)>(&ApproveTxCtx77da62::Signal_NonceChanged));
}

void ApproveTxCtx77da62_NonceChanged(void* ptr, struct Moc_PackedString nonce)
{
	static_cast<ApproveTxCtx77da62*>(ptr)->nonceChanged(QString::fromUtf8(nonce.data, nonce.len));
}

struct Moc_PackedString ApproveTxCtx77da62_Value(void* ptr)
{
	return ({ QByteArray te5a3fa = static_cast<ApproveTxCtx77da62*>(ptr)->value().toUtf8(); Moc_PackedString { const_cast<char*>(te5a3fa.prepend("WHITESPACE").constData()+10), te5a3fa.size()-10 }; });
}

struct Moc_PackedString ApproveTxCtx77da62_ValueDefault(void* ptr)
{
	return ({ QByteArray t9772be = static_cast<ApproveTxCtx77da62*>(ptr)->valueDefault().toUtf8(); Moc_PackedString { const_cast<char*>(t9772be.prepend("WHITESPACE").constData()+10), t9772be.size()-10 }; });
}

void ApproveTxCtx77da62_SetValue(void* ptr, struct Moc_PackedString value)
{
	static_cast<ApproveTxCtx77da62*>(ptr)->setValue(QString::fromUtf8(value.data, value.len));
}

void ApproveTxCtx77da62_SetValueDefault(void* ptr, struct Moc_PackedString value)
{
	static_cast<ApproveTxCtx77da62*>(ptr)->setValueDefault(QString::fromUtf8(value.data, value.len));
}

void ApproveTxCtx77da62_ConnectValueChanged(void* ptr)
{
	QObject::connect(static_cast<ApproveTxCtx77da62*>(ptr), static_cast<void (ApproveTxCtx77da62::*)(QString)>(&ApproveTxCtx77da62::valueChanged), static_cast<ApproveTxCtx77da62*>(ptr), static_cast<void (ApproveTxCtx77da62::*)(QString)>(&ApproveTxCtx77da62::Signal_ValueChanged));
}

void ApproveTxCtx77da62_DisconnectValueChanged(void* ptr)
{
	QObject::disconnect(static_cast<ApproveTxCtx77da62*>(ptr), static_cast<void (ApproveTxCtx77da62::*)(QString)>(&ApproveTxCtx77da62::valueChanged), static_cast<ApproveTxCtx77da62*>(ptr), static_cast<void (ApproveTxCtx77da62::*)(QString)>(&ApproveTxCtx77da62::Signal_ValueChanged));
}

void ApproveTxCtx77da62_ValueChanged(void* ptr, struct Moc_PackedString value)
{
	static_cast<ApproveTxCtx77da62*>(ptr)->valueChanged(QString::fromUtf8(value.data, value.len));
}

struct Moc_PackedString ApproveTxCtx77da62_Password(void* ptr)
{
	return ({ QByteArray t918077 = static_cast<ApproveTxCtx77da62*>(ptr)->password().toUtf8(); Moc_PackedString { const_cast<char*>(t918077.prepend("WHITESPACE").constData()+10), t918077.size()-10 }; });
}

struct Moc_PackedString ApproveTxCtx77da62_PasswordDefault(void* ptr)
{
	return ({ QByteArray t686f68 = static_cast<ApproveTxCtx77da62*>(ptr)->passwordDefault().toUtf8(); Moc_PackedString { const_cast<char*>(t686f68.prepend("WHITESPACE").constData()+10), t686f68.size()-10 }; });
}

void ApproveTxCtx77da62_SetPassword(void* ptr, struct Moc_PackedString password)
{
	static_cast<ApproveTxCtx77da62*>(ptr)->setPassword(QString::fromUtf8(password.data, password.len));
}

void ApproveTxCtx77da62_SetPasswordDefault(void* ptr, struct Moc_PackedString password)
{
	static_cast<ApproveTxCtx77da62*>(ptr)->setPasswordDefault(QString::fromUtf8(password.data, password.len));
}

void ApproveTxCtx77da62_ConnectPasswordChanged(void* ptr)
{
	QObject::connect(static_cast<ApproveTxCtx77da62*>(ptr), static_cast<void (ApproveTxCtx77da62::*)(QString)>(&ApproveTxCtx77da62::passwordChanged), static_cast<ApproveTxCtx77da62*>(ptr), static_cast<void (ApproveTxCtx77da62::*)(QString)>(&ApproveTxCtx77da62::Signal_PasswordChanged));
}

void ApproveTxCtx77da62_DisconnectPasswordChanged(void* ptr)
{
	QObject::disconnect(static_cast<ApproveTxCtx77da62*>(ptr), static_cast<void (ApproveTxCtx77da62::*)(QString)>(&ApproveTxCtx77da62::passwordChanged), static_cast<ApproveTxCtx77da62*>(ptr), static_cast<void (ApproveTxCtx77da62::*)(QString)>(&ApproveTxCtx77da62::Signal_PasswordChanged));
}

void ApproveTxCtx77da62_PasswordChanged(void* ptr, struct Moc_PackedString password)
{
	static_cast<ApproveTxCtx77da62*>(ptr)->passwordChanged(QString::fromUtf8(password.data, password.len));
}

struct Moc_PackedString ApproveTxCtx77da62_FromSrc(void* ptr)
{
	return ({ QByteArray t6aa1b7 = static_cast<ApproveTxCtx77da62*>(ptr)->fromSrc().toUtf8(); Moc_PackedString { const_cast<char*>(t6aa1b7.prepend("WHITESPACE").constData()+10), t6aa1b7.size()-10 }; });
}

struct Moc_PackedString ApproveTxCtx77da62_FromSrcDefault(void* ptr)
{
	return ({ QByteArray tc7a000 = static_cast<ApproveTxCtx77da62*>(ptr)->fromSrcDefault().toUtf8(); Moc_PackedString { const_cast<char*>(tc7a000.prepend("WHITESPACE").constData()+10), tc7a000.size()-10 }; });
}

void ApproveTxCtx77da62_SetFromSrc(void* ptr, struct Moc_PackedString fromSrc)
{
	static_cast<ApproveTxCtx77da62*>(ptr)->setFromSrc(QString::fromUtf8(fromSrc.data, fromSrc.len));
}

void ApproveTxCtx77da62_SetFromSrcDefault(void* ptr, struct Moc_PackedString fromSrc)
{
	static_cast<ApproveTxCtx77da62*>(ptr)->setFromSrcDefault(QString::fromUtf8(fromSrc.data, fromSrc.len));
}

void ApproveTxCtx77da62_ConnectFromSrcChanged(void* ptr)
{
	QObject::connect(static_cast<ApproveTxCtx77da62*>(ptr), static_cast<void (ApproveTxCtx77da62::*)(QString)>(&ApproveTxCtx77da62::fromSrcChanged), static_cast<ApproveTxCtx77da62*>(ptr), static_cast<void (ApproveTxCtx77da62::*)(QString)>(&ApproveTxCtx77da62::Signal_FromSrcChanged));
}

void ApproveTxCtx77da62_DisconnectFromSrcChanged(void* ptr)
{
	QObject::disconnect(static_cast<ApproveTxCtx77da62*>(ptr), static_cast<void (ApproveTxCtx77da62::*)(QString)>(&ApproveTxCtx77da62::fromSrcChanged), static_cast<ApproveTxCtx77da62*>(ptr), static_cast<void (ApproveTxCtx77da62::*)(QString)>(&ApproveTxCtx77da62::Signal_FromSrcChanged));
}

void ApproveTxCtx77da62_FromSrcChanged(void* ptr, struct Moc_PackedString fromSrc)
{
	static_cast<ApproveTxCtx77da62*>(ptr)->fromSrcChanged(QString::fromUtf8(fromSrc.data, fromSrc.len));
}

struct Moc_PackedString ApproveTxCtx77da62_ToSrc(void* ptr)
{
	return ({ QByteArray t03230d = static_cast<ApproveTxCtx77da62*>(ptr)->toSrc().toUtf8(); Moc_PackedString { const_cast<char*>(t03230d.prepend("WHITESPACE").constData()+10), t03230d.size()-10 }; });
}

struct Moc_PackedString ApproveTxCtx77da62_ToSrcDefault(void* ptr)
{
	return ({ QByteArray t20ab85 = static_cast<ApproveTxCtx77da62*>(ptr)->toSrcDefault().toUtf8(); Moc_PackedString { const_cast<char*>(t20ab85.prepend("WHITESPACE").constData()+10), t20ab85.size()-10 }; });
}

void ApproveTxCtx77da62_SetToSrc(void* ptr, struct Moc_PackedString toSrc)
{
	static_cast<ApproveTxCtx77da62*>(ptr)->setToSrc(QString::fromUtf8(toSrc.data, toSrc.len));
}

void ApproveTxCtx77da62_SetToSrcDefault(void* ptr, struct Moc_PackedString toSrc)
{
	static_cast<ApproveTxCtx77da62*>(ptr)->setToSrcDefault(QString::fromUtf8(toSrc.data, toSrc.len));
}

void ApproveTxCtx77da62_ConnectToSrcChanged(void* ptr)
{
	QObject::connect(static_cast<ApproveTxCtx77da62*>(ptr), static_cast<void (ApproveTxCtx77da62::*)(QString)>(&ApproveTxCtx77da62::toSrcChanged), static_cast<ApproveTxCtx77da62*>(ptr), static_cast<void (ApproveTxCtx77da62::*)(QString)>(&ApproveTxCtx77da62::Signal_ToSrcChanged));
}

void ApproveTxCtx77da62_DisconnectToSrcChanged(void* ptr)
{
	QObject::disconnect(static_cast<ApproveTxCtx77da62*>(ptr), static_cast<void (ApproveTxCtx77da62::*)(QString)>(&ApproveTxCtx77da62::toSrcChanged), static_cast<ApproveTxCtx77da62*>(ptr), static_cast<void (ApproveTxCtx77da62::*)(QString)>(&ApproveTxCtx77da62::Signal_ToSrcChanged));
}

void ApproveTxCtx77da62_ToSrcChanged(void* ptr, struct Moc_PackedString toSrc)
{
	static_cast<ApproveTxCtx77da62*>(ptr)->toSrcChanged(QString::fromUtf8(toSrc.data, toSrc.len));
}

struct Moc_PackedString ApproveTxCtx77da62_Diff(void* ptr)
{
	return ({ QByteArray t81763e = static_cast<ApproveTxCtx77da62*>(ptr)->diff().toUtf8(); Moc_PackedString { const_cast<char*>(t81763e.prepend("WHITESPACE").constData()+10), t81763e.size()-10 }; });
}

struct Moc_PackedString ApproveTxCtx77da62_DiffDefault(void* ptr)
{
	return ({ QByteArray td180f6 = static_cast<ApproveTxCtx77da62*>(ptr)->diffDefault().toUtf8(); Moc_PackedString { const_cast<char*>(td180f6.prepend("WHITESPACE").constData()+10), td180f6.size()-10 }; });
}

void ApproveTxCtx77da62_SetDiff(void* ptr, struct Moc_PackedString diff)
{
	static_cast<ApproveTxCtx77da62*>(ptr)->setDiff(QString::fromUtf8(diff.data, diff.len));
}

void ApproveTxCtx77da62_SetDiffDefault(void* ptr, struct Moc_PackedString diff)
{
	static_cast<ApproveTxCtx77da62*>(ptr)->setDiffDefault(QString::fromUtf8(diff.data, diff.len));
}

void ApproveTxCtx77da62_ConnectDiffChanged(void* ptr)
{
	QObject::connect(static_cast<ApproveTxCtx77da62*>(ptr), static_cast<void (ApproveTxCtx77da62::*)(QString)>(&ApproveTxCtx77da62::diffChanged), static_cast<ApproveTxCtx77da62*>(ptr), static_cast<void (ApproveTxCtx77da62::*)(QString)>(&ApproveTxCtx77da62::Signal_DiffChanged));
}

void ApproveTxCtx77da62_DisconnectDiffChanged(void* ptr)
{
	QObject::disconnect(static_cast<ApproveTxCtx77da62*>(ptr), static_cast<void (ApproveTxCtx77da62::*)(QString)>(&ApproveTxCtx77da62::diffChanged), static_cast<ApproveTxCtx77da62*>(ptr), static_cast<void (ApproveTxCtx77da62::*)(QString)>(&ApproveTxCtx77da62::Signal_DiffChanged));
}

void ApproveTxCtx77da62_DiffChanged(void* ptr, struct Moc_PackedString diff)
{
	static_cast<ApproveTxCtx77da62*>(ptr)->diffChanged(QString::fromUtf8(diff.data, diff.len));
}

int ApproveTxCtx77da62_ApproveTxCtx77da62_QRegisterMetaType()
{
	return qRegisterMetaType<ApproveTxCtx77da62*>();
}

int ApproveTxCtx77da62_ApproveTxCtx77da62_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<ApproveTxCtx77da62*>(const_cast<const char*>(typeName));
}

int ApproveTxCtx77da62_ApproveTxCtx77da62_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<ApproveTxCtx77da62>();
#else
	return 0;
#endif
}

int ApproveTxCtx77da62_ApproveTxCtx77da62_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<ApproveTxCtx77da62>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

void* ApproveTxCtx77da62___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void ApproveTxCtx77da62___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* ApproveTxCtx77da62___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* ApproveTxCtx77da62___findChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void ApproveTxCtx77da62___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* ApproveTxCtx77da62___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* ApproveTxCtx77da62___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void ApproveTxCtx77da62___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* ApproveTxCtx77da62___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* ApproveTxCtx77da62___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void ApproveTxCtx77da62___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* ApproveTxCtx77da62___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* ApproveTxCtx77da62___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void ApproveTxCtx77da62___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* ApproveTxCtx77da62___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* ApproveTxCtx77da62_NewApproveTxCtx(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new ApproveTxCtx77da62(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new ApproveTxCtx77da62(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new ApproveTxCtx77da62(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new ApproveTxCtx77da62(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new ApproveTxCtx77da62(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new ApproveTxCtx77da62(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new ApproveTxCtx77da62(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new ApproveTxCtx77da62(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new ApproveTxCtx77da62(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new ApproveTxCtx77da62(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new ApproveTxCtx77da62(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new ApproveTxCtx77da62(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new ApproveTxCtx77da62(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new ApproveTxCtx77da62(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new ApproveTxCtx77da62(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new ApproveTxCtx77da62(static_cast<QWindow*>(parent));
	} else {
		return new ApproveTxCtx77da62(static_cast<QObject*>(parent));
	}
}

void ApproveTxCtx77da62_DestroyApproveTxCtx(void* ptr)
{
	static_cast<ApproveTxCtx77da62*>(ptr)->~ApproveTxCtx77da62();
}

void ApproveTxCtx77da62_DestroyApproveTxCtxDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

char ApproveTxCtx77da62_EventDefault(void* ptr, void* e)
{
	return static_cast<ApproveTxCtx77da62*>(ptr)->QObject::event(static_cast<QEvent*>(e));
}

char ApproveTxCtx77da62_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<ApproveTxCtx77da62*>(ptr)->QObject::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void ApproveTxCtx77da62_ChildEventDefault(void* ptr, void* event)
{
	static_cast<ApproveTxCtx77da62*>(ptr)->QObject::childEvent(static_cast<QChildEvent*>(event));
}

void ApproveTxCtx77da62_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<ApproveTxCtx77da62*>(ptr)->QObject::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void ApproveTxCtx77da62_CustomEventDefault(void* ptr, void* event)
{
	static_cast<ApproveTxCtx77da62*>(ptr)->QObject::customEvent(static_cast<QEvent*>(event));
}

void ApproveTxCtx77da62_DeleteLaterDefault(void* ptr)
{
	static_cast<ApproveTxCtx77da62*>(ptr)->QObject::deleteLater();
}

void ApproveTxCtx77da62_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<ApproveTxCtx77da62*>(ptr)->QObject::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void ApproveTxCtx77da62_TimerEventDefault(void* ptr, void* event)
{
	static_cast<ApproveTxCtx77da62*>(ptr)->QObject::timerEvent(static_cast<QTimerEvent*>(event));
}



void LoginCtx77da62_ConnectClicked(void* ptr)
{
	QObject::connect(static_cast<LoginCtx77da62*>(ptr), static_cast<void (LoginCtx77da62::*)()>(&LoginCtx77da62::clicked), static_cast<LoginCtx77da62*>(ptr), static_cast<void (LoginCtx77da62::*)()>(&LoginCtx77da62::Signal_Clicked));
}

void LoginCtx77da62_DisconnectClicked(void* ptr)
{
	QObject::disconnect(static_cast<LoginCtx77da62*>(ptr), static_cast<void (LoginCtx77da62::*)()>(&LoginCtx77da62::clicked), static_cast<LoginCtx77da62*>(ptr), static_cast<void (LoginCtx77da62::*)()>(&LoginCtx77da62::Signal_Clicked));
}

void LoginCtx77da62_Clicked(void* ptr)
{
	static_cast<LoginCtx77da62*>(ptr)->clicked();
}

void LoginCtx77da62_ConnectEdited(void* ptr)
{
	QObject::connect(static_cast<LoginCtx77da62*>(ptr), static_cast<void (LoginCtx77da62::*)(QString)>(&LoginCtx77da62::edited), static_cast<LoginCtx77da62*>(ptr), static_cast<void (LoginCtx77da62::*)(QString)>(&LoginCtx77da62::Signal_Edited));
}

void LoginCtx77da62_DisconnectEdited(void* ptr)
{
	QObject::disconnect(static_cast<LoginCtx77da62*>(ptr), static_cast<void (LoginCtx77da62::*)(QString)>(&LoginCtx77da62::edited), static_cast<LoginCtx77da62*>(ptr), static_cast<void (LoginCtx77da62::*)(QString)>(&LoginCtx77da62::Signal_Edited));
}

void LoginCtx77da62_Edited(void* ptr, struct Moc_PackedString b)
{
	static_cast<LoginCtx77da62*>(ptr)->edited(QString::fromUtf8(b.data, b.len));
}

struct Moc_PackedString LoginCtx77da62_Remote(void* ptr)
{
	return ({ QByteArray tae1d12 = static_cast<LoginCtx77da62*>(ptr)->remote().toUtf8(); Moc_PackedString { const_cast<char*>(tae1d12.prepend("WHITESPACE").constData()+10), tae1d12.size()-10 }; });
}

struct Moc_PackedString LoginCtx77da62_RemoteDefault(void* ptr)
{
	return ({ QByteArray t330ea6 = static_cast<LoginCtx77da62*>(ptr)->remoteDefault().toUtf8(); Moc_PackedString { const_cast<char*>(t330ea6.prepend("WHITESPACE").constData()+10), t330ea6.size()-10 }; });
}

void LoginCtx77da62_SetRemote(void* ptr, struct Moc_PackedString remote)
{
	static_cast<LoginCtx77da62*>(ptr)->setRemote(QString::fromUtf8(remote.data, remote.len));
}

void LoginCtx77da62_SetRemoteDefault(void* ptr, struct Moc_PackedString remote)
{
	static_cast<LoginCtx77da62*>(ptr)->setRemoteDefault(QString::fromUtf8(remote.data, remote.len));
}

void LoginCtx77da62_ConnectRemoteChanged(void* ptr)
{
	QObject::connect(static_cast<LoginCtx77da62*>(ptr), static_cast<void (LoginCtx77da62::*)(QString)>(&LoginCtx77da62::remoteChanged), static_cast<LoginCtx77da62*>(ptr), static_cast<void (LoginCtx77da62::*)(QString)>(&LoginCtx77da62::Signal_RemoteChanged));
}

void LoginCtx77da62_DisconnectRemoteChanged(void* ptr)
{
	QObject::disconnect(static_cast<LoginCtx77da62*>(ptr), static_cast<void (LoginCtx77da62::*)(QString)>(&LoginCtx77da62::remoteChanged), static_cast<LoginCtx77da62*>(ptr), static_cast<void (LoginCtx77da62::*)(QString)>(&LoginCtx77da62::Signal_RemoteChanged));
}

void LoginCtx77da62_RemoteChanged(void* ptr, struct Moc_PackedString remote)
{
	static_cast<LoginCtx77da62*>(ptr)->remoteChanged(QString::fromUtf8(remote.data, remote.len));
}

struct Moc_PackedString LoginCtx77da62_Transport(void* ptr)
{
	return ({ QByteArray tb90150 = static_cast<LoginCtx77da62*>(ptr)->transport().toUtf8(); Moc_PackedString { const_cast<char*>(tb90150.prepend("WHITESPACE").constData()+10), tb90150.size()-10 }; });
}

struct Moc_PackedString LoginCtx77da62_TransportDefault(void* ptr)
{
	return ({ QByteArray t21023e = static_cast<LoginCtx77da62*>(ptr)->transportDefault().toUtf8(); Moc_PackedString { const_cast<char*>(t21023e.prepend("WHITESPACE").constData()+10), t21023e.size()-10 }; });
}

void LoginCtx77da62_SetTransport(void* ptr, struct Moc_PackedString transport)
{
	static_cast<LoginCtx77da62*>(ptr)->setTransport(QString::fromUtf8(transport.data, transport.len));
}

void LoginCtx77da62_SetTransportDefault(void* ptr, struct Moc_PackedString transport)
{
	static_cast<LoginCtx77da62*>(ptr)->setTransportDefault(QString::fromUtf8(transport.data, transport.len));
}

void LoginCtx77da62_ConnectTransportChanged(void* ptr)
{
	QObject::connect(static_cast<LoginCtx77da62*>(ptr), static_cast<void (LoginCtx77da62::*)(QString)>(&LoginCtx77da62::transportChanged), static_cast<LoginCtx77da62*>(ptr), static_cast<void (LoginCtx77da62::*)(QString)>(&LoginCtx77da62::Signal_TransportChanged));
}

void LoginCtx77da62_DisconnectTransportChanged(void* ptr)
{
	QObject::disconnect(static_cast<LoginCtx77da62*>(ptr), static_cast<void (LoginCtx77da62::*)(QString)>(&LoginCtx77da62::transportChanged), static_cast<LoginCtx77da62*>(ptr), static_cast<void (LoginCtx77da62::*)(QString)>(&LoginCtx77da62::Signal_TransportChanged));
}

void LoginCtx77da62_TransportChanged(void* ptr, struct Moc_PackedString transport)
{
	static_cast<LoginCtx77da62*>(ptr)->transportChanged(QString::fromUtf8(transport.data, transport.len));
}

struct Moc_PackedString LoginCtx77da62_Endpoint(void* ptr)
{
	return ({ QByteArray t5dca4f = static_cast<LoginCtx77da62*>(ptr)->endpoint().toUtf8(); Moc_PackedString { const_cast<char*>(t5dca4f.prepend("WHITESPACE").constData()+10), t5dca4f.size()-10 }; });
}

struct Moc_PackedString LoginCtx77da62_EndpointDefault(void* ptr)
{
	return ({ QByteArray t776b2b = static_cast<LoginCtx77da62*>(ptr)->endpointDefault().toUtf8(); Moc_PackedString { const_cast<char*>(t776b2b.prepend("WHITESPACE").constData()+10), t776b2b.size()-10 }; });
}

void LoginCtx77da62_SetEndpoint(void* ptr, struct Moc_PackedString endpoint)
{
	static_cast<LoginCtx77da62*>(ptr)->setEndpoint(QString::fromUtf8(endpoint.data, endpoint.len));
}

void LoginCtx77da62_SetEndpointDefault(void* ptr, struct Moc_PackedString endpoint)
{
	static_cast<LoginCtx77da62*>(ptr)->setEndpointDefault(QString::fromUtf8(endpoint.data, endpoint.len));
}

void LoginCtx77da62_ConnectEndpointChanged(void* ptr)
{
	QObject::connect(static_cast<LoginCtx77da62*>(ptr), static_cast<void (LoginCtx77da62::*)(QString)>(&LoginCtx77da62::endpointChanged), static_cast<LoginCtx77da62*>(ptr), static_cast<void (LoginCtx77da62::*)(QString)>(&LoginCtx77da62::Signal_EndpointChanged));
}

void LoginCtx77da62_DisconnectEndpointChanged(void* ptr)
{
	QObject::disconnect(static_cast<LoginCtx77da62*>(ptr), static_cast<void (LoginCtx77da62::*)(QString)>(&LoginCtx77da62::endpointChanged), static_cast<LoginCtx77da62*>(ptr), static_cast<void (LoginCtx77da62::*)(QString)>(&LoginCtx77da62::Signal_EndpointChanged));
}

void LoginCtx77da62_EndpointChanged(void* ptr, struct Moc_PackedString endpoint)
{
	static_cast<LoginCtx77da62*>(ptr)->endpointChanged(QString::fromUtf8(endpoint.data, endpoint.len));
}

struct Moc_PackedString LoginCtx77da62_Gopath(void* ptr)
{
	return ({ QByteArray t184777 = static_cast<LoginCtx77da62*>(ptr)->gopath().toUtf8(); Moc_PackedString { const_cast<char*>(t184777.prepend("WHITESPACE").constData()+10), t184777.size()-10 }; });
}

struct Moc_PackedString LoginCtx77da62_GopathDefault(void* ptr)
{
	return ({ QByteArray td9192e = static_cast<LoginCtx77da62*>(ptr)->gopathDefault().toUtf8(); Moc_PackedString { const_cast<char*>(td9192e.prepend("WHITESPACE").constData()+10), td9192e.size()-10 }; });
}

void LoginCtx77da62_SetGopath(void* ptr, struct Moc_PackedString gopath)
{
	static_cast<LoginCtx77da62*>(ptr)->setGopath(QString::fromUtf8(gopath.data, gopath.len));
}

void LoginCtx77da62_SetGopathDefault(void* ptr, struct Moc_PackedString gopath)
{
	static_cast<LoginCtx77da62*>(ptr)->setGopathDefault(QString::fromUtf8(gopath.data, gopath.len));
}

void LoginCtx77da62_ConnectGopathChanged(void* ptr)
{
	QObject::connect(static_cast<LoginCtx77da62*>(ptr), static_cast<void (LoginCtx77da62::*)(QString)>(&LoginCtx77da62::gopathChanged), static_cast<LoginCtx77da62*>(ptr), static_cast<void (LoginCtx77da62::*)(QString)>(&LoginCtx77da62::Signal_GopathChanged));
}

void LoginCtx77da62_DisconnectGopathChanged(void* ptr)
{
	QObject::disconnect(static_cast<LoginCtx77da62*>(ptr), static_cast<void (LoginCtx77da62::*)(QString)>(&LoginCtx77da62::gopathChanged), static_cast<LoginCtx77da62*>(ptr), static_cast<void (LoginCtx77da62::*)(QString)>(&LoginCtx77da62::Signal_GopathChanged));
}

void LoginCtx77da62_GopathChanged(void* ptr, struct Moc_PackedString gopath)
{
	static_cast<LoginCtx77da62*>(ptr)->gopathChanged(QString::fromUtf8(gopath.data, gopath.len));
}

struct Moc_PackedString LoginCtx77da62_BinaryHash(void* ptr)
{
	return ({ QByteArray t668e2d = static_cast<LoginCtx77da62*>(ptr)->binaryHash().toUtf8(); Moc_PackedString { const_cast<char*>(t668e2d.prepend("WHITESPACE").constData()+10), t668e2d.size()-10 }; });
}

struct Moc_PackedString LoginCtx77da62_BinaryHashDefault(void* ptr)
{
	return ({ QByteArray tccc185 = static_cast<LoginCtx77da62*>(ptr)->binaryHashDefault().toUtf8(); Moc_PackedString { const_cast<char*>(tccc185.prepend("WHITESPACE").constData()+10), tccc185.size()-10 }; });
}

void LoginCtx77da62_SetBinaryHash(void* ptr, struct Moc_PackedString binaryHash)
{
	static_cast<LoginCtx77da62*>(ptr)->setBinaryHash(QString::fromUtf8(binaryHash.data, binaryHash.len));
}

void LoginCtx77da62_SetBinaryHashDefault(void* ptr, struct Moc_PackedString binaryHash)
{
	static_cast<LoginCtx77da62*>(ptr)->setBinaryHashDefault(QString::fromUtf8(binaryHash.data, binaryHash.len));
}

void LoginCtx77da62_ConnectBinaryHashChanged(void* ptr)
{
	QObject::connect(static_cast<LoginCtx77da62*>(ptr), static_cast<void (LoginCtx77da62::*)(QString)>(&LoginCtx77da62::binaryHashChanged), static_cast<LoginCtx77da62*>(ptr), static_cast<void (LoginCtx77da62::*)(QString)>(&LoginCtx77da62::Signal_BinaryHashChanged));
}

void LoginCtx77da62_DisconnectBinaryHashChanged(void* ptr)
{
	QObject::disconnect(static_cast<LoginCtx77da62*>(ptr), static_cast<void (LoginCtx77da62::*)(QString)>(&LoginCtx77da62::binaryHashChanged), static_cast<LoginCtx77da62*>(ptr), static_cast<void (LoginCtx77da62::*)(QString)>(&LoginCtx77da62::Signal_BinaryHashChanged));
}

void LoginCtx77da62_BinaryHashChanged(void* ptr, struct Moc_PackedString binaryHash)
{
	static_cast<LoginCtx77da62*>(ptr)->binaryHashChanged(QString::fromUtf8(binaryHash.data, binaryHash.len));
}

char LoginCtx77da62_IsValid(void* ptr)
{
	return static_cast<LoginCtx77da62*>(ptr)->isValid();
}

char LoginCtx77da62_IsValidDefault(void* ptr)
{
	return static_cast<LoginCtx77da62*>(ptr)->isValidDefault();
}

void LoginCtx77da62_SetIsValid(void* ptr, char isValid)
{
	static_cast<LoginCtx77da62*>(ptr)->setIsValid(isValid != 0);
}

void LoginCtx77da62_SetIsValidDefault(void* ptr, char isValid)
{
	static_cast<LoginCtx77da62*>(ptr)->setIsValidDefault(isValid != 0);
}

void LoginCtx77da62_ConnectIsValidChanged(void* ptr)
{
	QObject::connect(static_cast<LoginCtx77da62*>(ptr), static_cast<void (LoginCtx77da62::*)(bool)>(&LoginCtx77da62::isValidChanged), static_cast<LoginCtx77da62*>(ptr), static_cast<void (LoginCtx77da62::*)(bool)>(&LoginCtx77da62::Signal_IsValidChanged));
}

void LoginCtx77da62_DisconnectIsValidChanged(void* ptr)
{
	QObject::disconnect(static_cast<LoginCtx77da62*>(ptr), static_cast<void (LoginCtx77da62::*)(bool)>(&LoginCtx77da62::isValidChanged), static_cast<LoginCtx77da62*>(ptr), static_cast<void (LoginCtx77da62::*)(bool)>(&LoginCtx77da62::Signal_IsValidChanged));
}

void LoginCtx77da62_IsValidChanged(void* ptr, char isValid)
{
	static_cast<LoginCtx77da62*>(ptr)->isValidChanged(isValid != 0);
}

int LoginCtx77da62_LoginCtx77da62_QRegisterMetaType()
{
	return qRegisterMetaType<LoginCtx77da62*>();
}

int LoginCtx77da62_LoginCtx77da62_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<LoginCtx77da62*>(const_cast<const char*>(typeName));
}

int LoginCtx77da62_LoginCtx77da62_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<LoginCtx77da62>();
#else
	return 0;
#endif
}

int LoginCtx77da62_LoginCtx77da62_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<LoginCtx77da62>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

void* LoginCtx77da62___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void LoginCtx77da62___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* LoginCtx77da62___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* LoginCtx77da62___findChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void LoginCtx77da62___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* LoginCtx77da62___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* LoginCtx77da62___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void LoginCtx77da62___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* LoginCtx77da62___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* LoginCtx77da62___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void LoginCtx77da62___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* LoginCtx77da62___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* LoginCtx77da62___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void LoginCtx77da62___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* LoginCtx77da62___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* LoginCtx77da62_NewLoginCtx(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new LoginCtx77da62(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new LoginCtx77da62(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new LoginCtx77da62(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new LoginCtx77da62(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new LoginCtx77da62(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new LoginCtx77da62(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new LoginCtx77da62(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new LoginCtx77da62(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new LoginCtx77da62(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new LoginCtx77da62(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new LoginCtx77da62(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new LoginCtx77da62(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new LoginCtx77da62(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new LoginCtx77da62(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new LoginCtx77da62(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new LoginCtx77da62(static_cast<QWindow*>(parent));
	} else {
		return new LoginCtx77da62(static_cast<QObject*>(parent));
	}
}

void LoginCtx77da62_DestroyLoginCtx(void* ptr)
{
	static_cast<LoginCtx77da62*>(ptr)->~LoginCtx77da62();
}

void LoginCtx77da62_DestroyLoginCtxDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

char LoginCtx77da62_EventDefault(void* ptr, void* e)
{
	return static_cast<LoginCtx77da62*>(ptr)->QObject::event(static_cast<QEvent*>(e));
}

char LoginCtx77da62_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<LoginCtx77da62*>(ptr)->QObject::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void LoginCtx77da62_ChildEventDefault(void* ptr, void* event)
{
	static_cast<LoginCtx77da62*>(ptr)->QObject::childEvent(static_cast<QChildEvent*>(event));
}

void LoginCtx77da62_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<LoginCtx77da62*>(ptr)->QObject::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void LoginCtx77da62_CustomEventDefault(void* ptr, void* event)
{
	static_cast<LoginCtx77da62*>(ptr)->QObject::customEvent(static_cast<QEvent*>(event));
}

void LoginCtx77da62_DeleteLaterDefault(void* ptr)
{
	static_cast<LoginCtx77da62*>(ptr)->QObject::deleteLater();
}

void LoginCtx77da62_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<LoginCtx77da62*>(ptr)->QObject::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void LoginCtx77da62_TimerEventDefault(void* ptr, void* event)
{
	static_cast<LoginCtx77da62*>(ptr)->QObject::timerEvent(static_cast<QTimerEvent*>(event));
}



void TxListCtx77da62_ConnectClicked(void* ptr)
{
	QObject::connect(static_cast<TxListCtx77da62*>(ptr), static_cast<void (TxListCtx77da62::*)(qint32)>(&TxListCtx77da62::clicked), static_cast<TxListCtx77da62*>(ptr), static_cast<void (TxListCtx77da62::*)(qint32)>(&TxListCtx77da62::Signal_Clicked));
}

void TxListCtx77da62_DisconnectClicked(void* ptr)
{
	QObject::disconnect(static_cast<TxListCtx77da62*>(ptr), static_cast<void (TxListCtx77da62::*)(qint32)>(&TxListCtx77da62::clicked), static_cast<TxListCtx77da62*>(ptr), static_cast<void (TxListCtx77da62::*)(qint32)>(&TxListCtx77da62::Signal_Clicked));
}

void TxListCtx77da62_Clicked(void* ptr, int b)
{
	static_cast<TxListCtx77da62*>(ptr)->clicked(b);
}

struct Moc_PackedString TxListCtx77da62_ShortenAddress(void* ptr)
{
	return ({ QByteArray tf699cf = static_cast<TxListCtx77da62*>(ptr)->shortenAddress().toUtf8(); Moc_PackedString { const_cast<char*>(tf699cf.prepend("WHITESPACE").constData()+10), tf699cf.size()-10 }; });
}

struct Moc_PackedString TxListCtx77da62_ShortenAddressDefault(void* ptr)
{
	return ({ QByteArray tde87a5 = static_cast<TxListCtx77da62*>(ptr)->shortenAddressDefault().toUtf8(); Moc_PackedString { const_cast<char*>(tde87a5.prepend("WHITESPACE").constData()+10), tde87a5.size()-10 }; });
}

void TxListCtx77da62_SetShortenAddress(void* ptr, struct Moc_PackedString shortenAddress)
{
	static_cast<TxListCtx77da62*>(ptr)->setShortenAddress(QString::fromUtf8(shortenAddress.data, shortenAddress.len));
}

void TxListCtx77da62_SetShortenAddressDefault(void* ptr, struct Moc_PackedString shortenAddress)
{
	static_cast<TxListCtx77da62*>(ptr)->setShortenAddressDefault(QString::fromUtf8(shortenAddress.data, shortenAddress.len));
}

void TxListCtx77da62_ConnectShortenAddressChanged(void* ptr)
{
	QObject::connect(static_cast<TxListCtx77da62*>(ptr), static_cast<void (TxListCtx77da62::*)(QString)>(&TxListCtx77da62::shortenAddressChanged), static_cast<TxListCtx77da62*>(ptr), static_cast<void (TxListCtx77da62::*)(QString)>(&TxListCtx77da62::Signal_ShortenAddressChanged));
}

void TxListCtx77da62_DisconnectShortenAddressChanged(void* ptr)
{
	QObject::disconnect(static_cast<TxListCtx77da62*>(ptr), static_cast<void (TxListCtx77da62::*)(QString)>(&TxListCtx77da62::shortenAddressChanged), static_cast<TxListCtx77da62*>(ptr), static_cast<void (TxListCtx77da62::*)(QString)>(&TxListCtx77da62::Signal_ShortenAddressChanged));
}

void TxListCtx77da62_ShortenAddressChanged(void* ptr, struct Moc_PackedString shortenAddress)
{
	static_cast<TxListCtx77da62*>(ptr)->shortenAddressChanged(QString::fromUtf8(shortenAddress.data, shortenAddress.len));
}

struct Moc_PackedString TxListCtx77da62_SelectedSrc(void* ptr)
{
	return ({ QByteArray tff7988 = static_cast<TxListCtx77da62*>(ptr)->selectedSrc().toUtf8(); Moc_PackedString { const_cast<char*>(tff7988.prepend("WHITESPACE").constData()+10), tff7988.size()-10 }; });
}

struct Moc_PackedString TxListCtx77da62_SelectedSrcDefault(void* ptr)
{
	return ({ QByteArray te8fe51 = static_cast<TxListCtx77da62*>(ptr)->selectedSrcDefault().toUtf8(); Moc_PackedString { const_cast<char*>(te8fe51.prepend("WHITESPACE").constData()+10), te8fe51.size()-10 }; });
}

void TxListCtx77da62_SetSelectedSrc(void* ptr, struct Moc_PackedString selectedSrc)
{
	static_cast<TxListCtx77da62*>(ptr)->setSelectedSrc(QString::fromUtf8(selectedSrc.data, selectedSrc.len));
}

void TxListCtx77da62_SetSelectedSrcDefault(void* ptr, struct Moc_PackedString selectedSrc)
{
	static_cast<TxListCtx77da62*>(ptr)->setSelectedSrcDefault(QString::fromUtf8(selectedSrc.data, selectedSrc.len));
}

void TxListCtx77da62_ConnectSelectedSrcChanged(void* ptr)
{
	QObject::connect(static_cast<TxListCtx77da62*>(ptr), static_cast<void (TxListCtx77da62::*)(QString)>(&TxListCtx77da62::selectedSrcChanged), static_cast<TxListCtx77da62*>(ptr), static_cast<void (TxListCtx77da62::*)(QString)>(&TxListCtx77da62::Signal_SelectedSrcChanged));
}

void TxListCtx77da62_DisconnectSelectedSrcChanged(void* ptr)
{
	QObject::disconnect(static_cast<TxListCtx77da62*>(ptr), static_cast<void (TxListCtx77da62::*)(QString)>(&TxListCtx77da62::selectedSrcChanged), static_cast<TxListCtx77da62*>(ptr), static_cast<void (TxListCtx77da62::*)(QString)>(&TxListCtx77da62::Signal_SelectedSrcChanged));
}

void TxListCtx77da62_SelectedSrcChanged(void* ptr, struct Moc_PackedString selectedSrc)
{
	static_cast<TxListCtx77da62*>(ptr)->selectedSrcChanged(QString::fromUtf8(selectedSrc.data, selectedSrc.len));
}

int TxListCtx77da62_TxListCtx77da62_QRegisterMetaType()
{
	return qRegisterMetaType<TxListCtx77da62*>();
}

int TxListCtx77da62_TxListCtx77da62_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<TxListCtx77da62*>(const_cast<const char*>(typeName));
}

int TxListCtx77da62_TxListCtx77da62_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<TxListCtx77da62>();
#else
	return 0;
#endif
}

int TxListCtx77da62_TxListCtx77da62_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<TxListCtx77da62>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

void* TxListCtx77da62___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void TxListCtx77da62___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* TxListCtx77da62___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* TxListCtx77da62___findChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void TxListCtx77da62___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* TxListCtx77da62___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* TxListCtx77da62___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void TxListCtx77da62___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* TxListCtx77da62___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* TxListCtx77da62___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void TxListCtx77da62___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* TxListCtx77da62___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* TxListCtx77da62___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void TxListCtx77da62___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* TxListCtx77da62___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* TxListCtx77da62_NewTxListCtx(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new TxListCtx77da62(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new TxListCtx77da62(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new TxListCtx77da62(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new TxListCtx77da62(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new TxListCtx77da62(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new TxListCtx77da62(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new TxListCtx77da62(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new TxListCtx77da62(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new TxListCtx77da62(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new TxListCtx77da62(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new TxListCtx77da62(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new TxListCtx77da62(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new TxListCtx77da62(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new TxListCtx77da62(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new TxListCtx77da62(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new TxListCtx77da62(static_cast<QWindow*>(parent));
	} else {
		return new TxListCtx77da62(static_cast<QObject*>(parent));
	}
}

void TxListCtx77da62_DestroyTxListCtx(void* ptr)
{
	static_cast<TxListCtx77da62*>(ptr)->~TxListCtx77da62();
}

void TxListCtx77da62_DestroyTxListCtxDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

char TxListCtx77da62_EventDefault(void* ptr, void* e)
{
	return static_cast<TxListCtx77da62*>(ptr)->QObject::event(static_cast<QEvent*>(e));
}

char TxListCtx77da62_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<TxListCtx77da62*>(ptr)->QObject::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void TxListCtx77da62_ChildEventDefault(void* ptr, void* event)
{
	static_cast<TxListCtx77da62*>(ptr)->QObject::childEvent(static_cast<QChildEvent*>(event));
}

void TxListCtx77da62_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<TxListCtx77da62*>(ptr)->QObject::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void TxListCtx77da62_CustomEventDefault(void* ptr, void* event)
{
	static_cast<TxListCtx77da62*>(ptr)->QObject::customEvent(static_cast<QEvent*>(event));
}

void TxListCtx77da62_DeleteLaterDefault(void* ptr)
{
	static_cast<TxListCtx77da62*>(ptr)->QObject::deleteLater();
}

void TxListCtx77da62_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<TxListCtx77da62*>(ptr)->QObject::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void TxListCtx77da62_TimerEventDefault(void* ptr, void* event)
{
	static_cast<TxListCtx77da62*>(ptr)->QObject::timerEvent(static_cast<QTimerEvent*>(event));
}



void ApproveSignDataCtx77da62_ConnectClicked(void* ptr)
{
	QObject::connect(static_cast<ApproveSignDataCtx77da62*>(ptr), static_cast<void (ApproveSignDataCtx77da62::*)(qint32)>(&ApproveSignDataCtx77da62::clicked), static_cast<ApproveSignDataCtx77da62*>(ptr), static_cast<void (ApproveSignDataCtx77da62::*)(qint32)>(&ApproveSignDataCtx77da62::Signal_Clicked));
}

void ApproveSignDataCtx77da62_DisconnectClicked(void* ptr)
{
	QObject::disconnect(static_cast<ApproveSignDataCtx77da62*>(ptr), static_cast<void (ApproveSignDataCtx77da62::*)(qint32)>(&ApproveSignDataCtx77da62::clicked), static_cast<ApproveSignDataCtx77da62*>(ptr), static_cast<void (ApproveSignDataCtx77da62::*)(qint32)>(&ApproveSignDataCtx77da62::Signal_Clicked));
}

void ApproveSignDataCtx77da62_Clicked(void* ptr, int b)
{
	static_cast<ApproveSignDataCtx77da62*>(ptr)->clicked(b);
}

void ApproveSignDataCtx77da62_ConnectOnBack(void* ptr)
{
	QObject::connect(static_cast<ApproveSignDataCtx77da62*>(ptr), static_cast<void (ApproveSignDataCtx77da62::*)()>(&ApproveSignDataCtx77da62::onBack), static_cast<ApproveSignDataCtx77da62*>(ptr), static_cast<void (ApproveSignDataCtx77da62::*)()>(&ApproveSignDataCtx77da62::Signal_OnBack));
}

void ApproveSignDataCtx77da62_DisconnectOnBack(void* ptr)
{
	QObject::disconnect(static_cast<ApproveSignDataCtx77da62*>(ptr), static_cast<void (ApproveSignDataCtx77da62::*)()>(&ApproveSignDataCtx77da62::onBack), static_cast<ApproveSignDataCtx77da62*>(ptr), static_cast<void (ApproveSignDataCtx77da62::*)()>(&ApproveSignDataCtx77da62::Signal_OnBack));
}

void ApproveSignDataCtx77da62_OnBack(void* ptr)
{
	static_cast<ApproveSignDataCtx77da62*>(ptr)->onBack();
}

void ApproveSignDataCtx77da62_ConnectOnApprove(void* ptr)
{
	QObject::connect(static_cast<ApproveSignDataCtx77da62*>(ptr), static_cast<void (ApproveSignDataCtx77da62::*)()>(&ApproveSignDataCtx77da62::onApprove), static_cast<ApproveSignDataCtx77da62*>(ptr), static_cast<void (ApproveSignDataCtx77da62::*)()>(&ApproveSignDataCtx77da62::Signal_OnApprove));
}

void ApproveSignDataCtx77da62_DisconnectOnApprove(void* ptr)
{
	QObject::disconnect(static_cast<ApproveSignDataCtx77da62*>(ptr), static_cast<void (ApproveSignDataCtx77da62::*)()>(&ApproveSignDataCtx77da62::onApprove), static_cast<ApproveSignDataCtx77da62*>(ptr), static_cast<void (ApproveSignDataCtx77da62::*)()>(&ApproveSignDataCtx77da62::Signal_OnApprove));
}

void ApproveSignDataCtx77da62_OnApprove(void* ptr)
{
	static_cast<ApproveSignDataCtx77da62*>(ptr)->onApprove();
}

void ApproveSignDataCtx77da62_ConnectOnReject(void* ptr)
{
	QObject::connect(static_cast<ApproveSignDataCtx77da62*>(ptr), static_cast<void (ApproveSignDataCtx77da62::*)()>(&ApproveSignDataCtx77da62::onReject), static_cast<ApproveSignDataCtx77da62*>(ptr), static_cast<void (ApproveSignDataCtx77da62::*)()>(&ApproveSignDataCtx77da62::Signal_OnReject));
}

void ApproveSignDataCtx77da62_DisconnectOnReject(void* ptr)
{
	QObject::disconnect(static_cast<ApproveSignDataCtx77da62*>(ptr), static_cast<void (ApproveSignDataCtx77da62::*)()>(&ApproveSignDataCtx77da62::onReject), static_cast<ApproveSignDataCtx77da62*>(ptr), static_cast<void (ApproveSignDataCtx77da62::*)()>(&ApproveSignDataCtx77da62::Signal_OnReject));
}

void ApproveSignDataCtx77da62_OnReject(void* ptr)
{
	static_cast<ApproveSignDataCtx77da62*>(ptr)->onReject();
}

void ApproveSignDataCtx77da62_ConnectEdited(void* ptr)
{
	QObject::connect(static_cast<ApproveSignDataCtx77da62*>(ptr), static_cast<void (ApproveSignDataCtx77da62::*)(QString, QString)>(&ApproveSignDataCtx77da62::edited), static_cast<ApproveSignDataCtx77da62*>(ptr), static_cast<void (ApproveSignDataCtx77da62::*)(QString, QString)>(&ApproveSignDataCtx77da62::Signal_Edited));
}

void ApproveSignDataCtx77da62_DisconnectEdited(void* ptr)
{
	QObject::disconnect(static_cast<ApproveSignDataCtx77da62*>(ptr), static_cast<void (ApproveSignDataCtx77da62::*)(QString, QString)>(&ApproveSignDataCtx77da62::edited), static_cast<ApproveSignDataCtx77da62*>(ptr), static_cast<void (ApproveSignDataCtx77da62::*)(QString, QString)>(&ApproveSignDataCtx77da62::Signal_Edited));
}

void ApproveSignDataCtx77da62_Edited(void* ptr, struct Moc_PackedString b, struct Moc_PackedString value)
{
	static_cast<ApproveSignDataCtx77da62*>(ptr)->edited(QString::fromUtf8(b.data, b.len), QString::fromUtf8(value.data, value.len));
}

struct Moc_PackedString ApproveSignDataCtx77da62_Remote(void* ptr)
{
	return ({ QByteArray td166de = static_cast<ApproveSignDataCtx77da62*>(ptr)->remote().toUtf8(); Moc_PackedString { const_cast<char*>(td166de.prepend("WHITESPACE").constData()+10), td166de.size()-10 }; });
}

struct Moc_PackedString ApproveSignDataCtx77da62_RemoteDefault(void* ptr)
{
	return ({ QByteArray t110aa3 = static_cast<ApproveSignDataCtx77da62*>(ptr)->remoteDefault().toUtf8(); Moc_PackedString { const_cast<char*>(t110aa3.prepend("WHITESPACE").constData()+10), t110aa3.size()-10 }; });
}

void ApproveSignDataCtx77da62_SetRemote(void* ptr, struct Moc_PackedString remote)
{
	static_cast<ApproveSignDataCtx77da62*>(ptr)->setRemote(QString::fromUtf8(remote.data, remote.len));
}

void ApproveSignDataCtx77da62_SetRemoteDefault(void* ptr, struct Moc_PackedString remote)
{
	static_cast<ApproveSignDataCtx77da62*>(ptr)->setRemoteDefault(QString::fromUtf8(remote.data, remote.len));
}

void ApproveSignDataCtx77da62_ConnectRemoteChanged(void* ptr)
{
	QObject::connect(static_cast<ApproveSignDataCtx77da62*>(ptr), static_cast<void (ApproveSignDataCtx77da62::*)(QString)>(&ApproveSignDataCtx77da62::remoteChanged), static_cast<ApproveSignDataCtx77da62*>(ptr), static_cast<void (ApproveSignDataCtx77da62::*)(QString)>(&ApproveSignDataCtx77da62::Signal_RemoteChanged));
}

void ApproveSignDataCtx77da62_DisconnectRemoteChanged(void* ptr)
{
	QObject::disconnect(static_cast<ApproveSignDataCtx77da62*>(ptr), static_cast<void (ApproveSignDataCtx77da62::*)(QString)>(&ApproveSignDataCtx77da62::remoteChanged), static_cast<ApproveSignDataCtx77da62*>(ptr), static_cast<void (ApproveSignDataCtx77da62::*)(QString)>(&ApproveSignDataCtx77da62::Signal_RemoteChanged));
}

void ApproveSignDataCtx77da62_RemoteChanged(void* ptr, struct Moc_PackedString remote)
{
	static_cast<ApproveSignDataCtx77da62*>(ptr)->remoteChanged(QString::fromUtf8(remote.data, remote.len));
}

struct Moc_PackedString ApproveSignDataCtx77da62_Transport(void* ptr)
{
	return ({ QByteArray tf8900c = static_cast<ApproveSignDataCtx77da62*>(ptr)->transport().toUtf8(); Moc_PackedString { const_cast<char*>(tf8900c.prepend("WHITESPACE").constData()+10), tf8900c.size()-10 }; });
}

struct Moc_PackedString ApproveSignDataCtx77da62_TransportDefault(void* ptr)
{
	return ({ QByteArray t04ad21 = static_cast<ApproveSignDataCtx77da62*>(ptr)->transportDefault().toUtf8(); Moc_PackedString { const_cast<char*>(t04ad21.prepend("WHITESPACE").constData()+10), t04ad21.size()-10 }; });
}

void ApproveSignDataCtx77da62_SetTransport(void* ptr, struct Moc_PackedString transport)
{
	static_cast<ApproveSignDataCtx77da62*>(ptr)->setTransport(QString::fromUtf8(transport.data, transport.len));
}

void ApproveSignDataCtx77da62_SetTransportDefault(void* ptr, struct Moc_PackedString transport)
{
	static_cast<ApproveSignDataCtx77da62*>(ptr)->setTransportDefault(QString::fromUtf8(transport.data, transport.len));
}

void ApproveSignDataCtx77da62_ConnectTransportChanged(void* ptr)
{
	QObject::connect(static_cast<ApproveSignDataCtx77da62*>(ptr), static_cast<void (ApproveSignDataCtx77da62::*)(QString)>(&ApproveSignDataCtx77da62::transportChanged), static_cast<ApproveSignDataCtx77da62*>(ptr), static_cast<void (ApproveSignDataCtx77da62::*)(QString)>(&ApproveSignDataCtx77da62::Signal_TransportChanged));
}

void ApproveSignDataCtx77da62_DisconnectTransportChanged(void* ptr)
{
	QObject::disconnect(static_cast<ApproveSignDataCtx77da62*>(ptr), static_cast<void (ApproveSignDataCtx77da62::*)(QString)>(&ApproveSignDataCtx77da62::transportChanged), static_cast<ApproveSignDataCtx77da62*>(ptr), static_cast<void (ApproveSignDataCtx77da62::*)(QString)>(&ApproveSignDataCtx77da62::Signal_TransportChanged));
}

void ApproveSignDataCtx77da62_TransportChanged(void* ptr, struct Moc_PackedString transport)
{
	static_cast<ApproveSignDataCtx77da62*>(ptr)->transportChanged(QString::fromUtf8(transport.data, transport.len));
}

struct Moc_PackedString ApproveSignDataCtx77da62_Endpoint(void* ptr)
{
	return ({ QByteArray tb0291e = static_cast<ApproveSignDataCtx77da62*>(ptr)->endpoint().toUtf8(); Moc_PackedString { const_cast<char*>(tb0291e.prepend("WHITESPACE").constData()+10), tb0291e.size()-10 }; });
}

struct Moc_PackedString ApproveSignDataCtx77da62_EndpointDefault(void* ptr)
{
	return ({ QByteArray t55e053 = static_cast<ApproveSignDataCtx77da62*>(ptr)->endpointDefault().toUtf8(); Moc_PackedString { const_cast<char*>(t55e053.prepend("WHITESPACE").constData()+10), t55e053.size()-10 }; });
}

void ApproveSignDataCtx77da62_SetEndpoint(void* ptr, struct Moc_PackedString endpoint)
{
	static_cast<ApproveSignDataCtx77da62*>(ptr)->setEndpoint(QString::fromUtf8(endpoint.data, endpoint.len));
}

void ApproveSignDataCtx77da62_SetEndpointDefault(void* ptr, struct Moc_PackedString endpoint)
{
	static_cast<ApproveSignDataCtx77da62*>(ptr)->setEndpointDefault(QString::fromUtf8(endpoint.data, endpoint.len));
}

void ApproveSignDataCtx77da62_ConnectEndpointChanged(void* ptr)
{
	QObject::connect(static_cast<ApproveSignDataCtx77da62*>(ptr), static_cast<void (ApproveSignDataCtx77da62::*)(QString)>(&ApproveSignDataCtx77da62::endpointChanged), static_cast<ApproveSignDataCtx77da62*>(ptr), static_cast<void (ApproveSignDataCtx77da62::*)(QString)>(&ApproveSignDataCtx77da62::Signal_EndpointChanged));
}

void ApproveSignDataCtx77da62_DisconnectEndpointChanged(void* ptr)
{
	QObject::disconnect(static_cast<ApproveSignDataCtx77da62*>(ptr), static_cast<void (ApproveSignDataCtx77da62::*)(QString)>(&ApproveSignDataCtx77da62::endpointChanged), static_cast<ApproveSignDataCtx77da62*>(ptr), static_cast<void (ApproveSignDataCtx77da62::*)(QString)>(&ApproveSignDataCtx77da62::Signal_EndpointChanged));
}

void ApproveSignDataCtx77da62_EndpointChanged(void* ptr, struct Moc_PackedString endpoint)
{
	static_cast<ApproveSignDataCtx77da62*>(ptr)->endpointChanged(QString::fromUtf8(endpoint.data, endpoint.len));
}

struct Moc_PackedString ApproveSignDataCtx77da62_From(void* ptr)
{
	return ({ QByteArray t1ec6d2 = static_cast<ApproveSignDataCtx77da62*>(ptr)->from().toUtf8(); Moc_PackedString { const_cast<char*>(t1ec6d2.prepend("WHITESPACE").constData()+10), t1ec6d2.size()-10 }; });
}

struct Moc_PackedString ApproveSignDataCtx77da62_FromDefault(void* ptr)
{
	return ({ QByteArray t8143a5 = static_cast<ApproveSignDataCtx77da62*>(ptr)->fromDefault().toUtf8(); Moc_PackedString { const_cast<char*>(t8143a5.prepend("WHITESPACE").constData()+10), t8143a5.size()-10 }; });
}

void ApproveSignDataCtx77da62_SetFrom(void* ptr, struct Moc_PackedString from)
{
	static_cast<ApproveSignDataCtx77da62*>(ptr)->setFrom(QString::fromUtf8(from.data, from.len));
}

void ApproveSignDataCtx77da62_SetFromDefault(void* ptr, struct Moc_PackedString from)
{
	static_cast<ApproveSignDataCtx77da62*>(ptr)->setFromDefault(QString::fromUtf8(from.data, from.len));
}

void ApproveSignDataCtx77da62_ConnectFromChanged(void* ptr)
{
	QObject::connect(static_cast<ApproveSignDataCtx77da62*>(ptr), static_cast<void (ApproveSignDataCtx77da62::*)(QString)>(&ApproveSignDataCtx77da62::fromChanged), static_cast<ApproveSignDataCtx77da62*>(ptr), static_cast<void (ApproveSignDataCtx77da62::*)(QString)>(&ApproveSignDataCtx77da62::Signal_FromChanged));
}

void ApproveSignDataCtx77da62_DisconnectFromChanged(void* ptr)
{
	QObject::disconnect(static_cast<ApproveSignDataCtx77da62*>(ptr), static_cast<void (ApproveSignDataCtx77da62::*)(QString)>(&ApproveSignDataCtx77da62::fromChanged), static_cast<ApproveSignDataCtx77da62*>(ptr), static_cast<void (ApproveSignDataCtx77da62::*)(QString)>(&ApproveSignDataCtx77da62::Signal_FromChanged));
}

void ApproveSignDataCtx77da62_FromChanged(void* ptr, struct Moc_PackedString from)
{
	static_cast<ApproveSignDataCtx77da62*>(ptr)->fromChanged(QString::fromUtf8(from.data, from.len));
}

struct Moc_PackedString ApproveSignDataCtx77da62_Message(void* ptr)
{
	return ({ QByteArray tf8c345 = static_cast<ApproveSignDataCtx77da62*>(ptr)->message().toUtf8(); Moc_PackedString { const_cast<char*>(tf8c345.prepend("WHITESPACE").constData()+10), tf8c345.size()-10 }; });
}

struct Moc_PackedString ApproveSignDataCtx77da62_MessageDefault(void* ptr)
{
	return ({ QByteArray t59a7b5 = static_cast<ApproveSignDataCtx77da62*>(ptr)->messageDefault().toUtf8(); Moc_PackedString { const_cast<char*>(t59a7b5.prepend("WHITESPACE").constData()+10), t59a7b5.size()-10 }; });
}

void ApproveSignDataCtx77da62_SetMessage(void* ptr, struct Moc_PackedString message)
{
	static_cast<ApproveSignDataCtx77da62*>(ptr)->setMessage(QString::fromUtf8(message.data, message.len));
}

void ApproveSignDataCtx77da62_SetMessageDefault(void* ptr, struct Moc_PackedString message)
{
	static_cast<ApproveSignDataCtx77da62*>(ptr)->setMessageDefault(QString::fromUtf8(message.data, message.len));
}

void ApproveSignDataCtx77da62_ConnectMessageChanged(void* ptr)
{
	QObject::connect(static_cast<ApproveSignDataCtx77da62*>(ptr), static_cast<void (ApproveSignDataCtx77da62::*)(QString)>(&ApproveSignDataCtx77da62::messageChanged), static_cast<ApproveSignDataCtx77da62*>(ptr), static_cast<void (ApproveSignDataCtx77da62::*)(QString)>(&ApproveSignDataCtx77da62::Signal_MessageChanged));
}

void ApproveSignDataCtx77da62_DisconnectMessageChanged(void* ptr)
{
	QObject::disconnect(static_cast<ApproveSignDataCtx77da62*>(ptr), static_cast<void (ApproveSignDataCtx77da62::*)(QString)>(&ApproveSignDataCtx77da62::messageChanged), static_cast<ApproveSignDataCtx77da62*>(ptr), static_cast<void (ApproveSignDataCtx77da62::*)(QString)>(&ApproveSignDataCtx77da62::Signal_MessageChanged));
}

void ApproveSignDataCtx77da62_MessageChanged(void* ptr, struct Moc_PackedString message)
{
	static_cast<ApproveSignDataCtx77da62*>(ptr)->messageChanged(QString::fromUtf8(message.data, message.len));
}

struct Moc_PackedString ApproveSignDataCtx77da62_RawData(void* ptr)
{
	return ({ QByteArray t71547a = static_cast<ApproveSignDataCtx77da62*>(ptr)->rawData().toUtf8(); Moc_PackedString { const_cast<char*>(t71547a.prepend("WHITESPACE").constData()+10), t71547a.size()-10 }; });
}

struct Moc_PackedString ApproveSignDataCtx77da62_RawDataDefault(void* ptr)
{
	return ({ QByteArray tf54367 = static_cast<ApproveSignDataCtx77da62*>(ptr)->rawDataDefault().toUtf8(); Moc_PackedString { const_cast<char*>(tf54367.prepend("WHITESPACE").constData()+10), tf54367.size()-10 }; });
}

void ApproveSignDataCtx77da62_SetRawData(void* ptr, struct Moc_PackedString rawData)
{
	static_cast<ApproveSignDataCtx77da62*>(ptr)->setRawData(QString::fromUtf8(rawData.data, rawData.len));
}

void ApproveSignDataCtx77da62_SetRawDataDefault(void* ptr, struct Moc_PackedString rawData)
{
	static_cast<ApproveSignDataCtx77da62*>(ptr)->setRawDataDefault(QString::fromUtf8(rawData.data, rawData.len));
}

void ApproveSignDataCtx77da62_ConnectRawDataChanged(void* ptr)
{
	QObject::connect(static_cast<ApproveSignDataCtx77da62*>(ptr), static_cast<void (ApproveSignDataCtx77da62::*)(QString)>(&ApproveSignDataCtx77da62::rawDataChanged), static_cast<ApproveSignDataCtx77da62*>(ptr), static_cast<void (ApproveSignDataCtx77da62::*)(QString)>(&ApproveSignDataCtx77da62::Signal_RawDataChanged));
}

void ApproveSignDataCtx77da62_DisconnectRawDataChanged(void* ptr)
{
	QObject::disconnect(static_cast<ApproveSignDataCtx77da62*>(ptr), static_cast<void (ApproveSignDataCtx77da62::*)(QString)>(&ApproveSignDataCtx77da62::rawDataChanged), static_cast<ApproveSignDataCtx77da62*>(ptr), static_cast<void (ApproveSignDataCtx77da62::*)(QString)>(&ApproveSignDataCtx77da62::Signal_RawDataChanged));
}

void ApproveSignDataCtx77da62_RawDataChanged(void* ptr, struct Moc_PackedString rawData)
{
	static_cast<ApproveSignDataCtx77da62*>(ptr)->rawDataChanged(QString::fromUtf8(rawData.data, rawData.len));
}

struct Moc_PackedString ApproveSignDataCtx77da62_Hash(void* ptr)
{
	return ({ QByteArray tf9f901 = static_cast<ApproveSignDataCtx77da62*>(ptr)->hash().toUtf8(); Moc_PackedString { const_cast<char*>(tf9f901.prepend("WHITESPACE").constData()+10), tf9f901.size()-10 }; });
}

struct Moc_PackedString ApproveSignDataCtx77da62_HashDefault(void* ptr)
{
	return ({ QByteArray tcc83f5 = static_cast<ApproveSignDataCtx77da62*>(ptr)->hashDefault().toUtf8(); Moc_PackedString { const_cast<char*>(tcc83f5.prepend("WHITESPACE").constData()+10), tcc83f5.size()-10 }; });
}

void ApproveSignDataCtx77da62_SetHash(void* ptr, struct Moc_PackedString hash)
{
	static_cast<ApproveSignDataCtx77da62*>(ptr)->setHash(QString::fromUtf8(hash.data, hash.len));
}

void ApproveSignDataCtx77da62_SetHashDefault(void* ptr, struct Moc_PackedString hash)
{
	static_cast<ApproveSignDataCtx77da62*>(ptr)->setHashDefault(QString::fromUtf8(hash.data, hash.len));
}

void ApproveSignDataCtx77da62_ConnectHashChanged(void* ptr)
{
	QObject::connect(static_cast<ApproveSignDataCtx77da62*>(ptr), static_cast<void (ApproveSignDataCtx77da62::*)(QString)>(&ApproveSignDataCtx77da62::hashChanged), static_cast<ApproveSignDataCtx77da62*>(ptr), static_cast<void (ApproveSignDataCtx77da62::*)(QString)>(&ApproveSignDataCtx77da62::Signal_HashChanged));
}

void ApproveSignDataCtx77da62_DisconnectHashChanged(void* ptr)
{
	QObject::disconnect(static_cast<ApproveSignDataCtx77da62*>(ptr), static_cast<void (ApproveSignDataCtx77da62::*)(QString)>(&ApproveSignDataCtx77da62::hashChanged), static_cast<ApproveSignDataCtx77da62*>(ptr), static_cast<void (ApproveSignDataCtx77da62::*)(QString)>(&ApproveSignDataCtx77da62::Signal_HashChanged));
}

void ApproveSignDataCtx77da62_HashChanged(void* ptr, struct Moc_PackedString hash)
{
	static_cast<ApproveSignDataCtx77da62*>(ptr)->hashChanged(QString::fromUtf8(hash.data, hash.len));
}

struct Moc_PackedString ApproveSignDataCtx77da62_Password(void* ptr)
{
	return ({ QByteArray t833eaa = static_cast<ApproveSignDataCtx77da62*>(ptr)->password().toUtf8(); Moc_PackedString { const_cast<char*>(t833eaa.prepend("WHITESPACE").constData()+10), t833eaa.size()-10 }; });
}

struct Moc_PackedString ApproveSignDataCtx77da62_PasswordDefault(void* ptr)
{
	return ({ QByteArray tbe2949 = static_cast<ApproveSignDataCtx77da62*>(ptr)->passwordDefault().toUtf8(); Moc_PackedString { const_cast<char*>(tbe2949.prepend("WHITESPACE").constData()+10), tbe2949.size()-10 }; });
}

void ApproveSignDataCtx77da62_SetPassword(void* ptr, struct Moc_PackedString password)
{
	static_cast<ApproveSignDataCtx77da62*>(ptr)->setPassword(QString::fromUtf8(password.data, password.len));
}

void ApproveSignDataCtx77da62_SetPasswordDefault(void* ptr, struct Moc_PackedString password)
{
	static_cast<ApproveSignDataCtx77da62*>(ptr)->setPasswordDefault(QString::fromUtf8(password.data, password.len));
}

void ApproveSignDataCtx77da62_ConnectPasswordChanged(void* ptr)
{
	QObject::connect(static_cast<ApproveSignDataCtx77da62*>(ptr), static_cast<void (ApproveSignDataCtx77da62::*)(QString)>(&ApproveSignDataCtx77da62::passwordChanged), static_cast<ApproveSignDataCtx77da62*>(ptr), static_cast<void (ApproveSignDataCtx77da62::*)(QString)>(&ApproveSignDataCtx77da62::Signal_PasswordChanged));
}

void ApproveSignDataCtx77da62_DisconnectPasswordChanged(void* ptr)
{
	QObject::disconnect(static_cast<ApproveSignDataCtx77da62*>(ptr), static_cast<void (ApproveSignDataCtx77da62::*)(QString)>(&ApproveSignDataCtx77da62::passwordChanged), static_cast<ApproveSignDataCtx77da62*>(ptr), static_cast<void (ApproveSignDataCtx77da62::*)(QString)>(&ApproveSignDataCtx77da62::Signal_PasswordChanged));
}

void ApproveSignDataCtx77da62_PasswordChanged(void* ptr, struct Moc_PackedString password)
{
	static_cast<ApproveSignDataCtx77da62*>(ptr)->passwordChanged(QString::fromUtf8(password.data, password.len));
}

struct Moc_PackedString ApproveSignDataCtx77da62_FromSrc(void* ptr)
{
	return ({ QByteArray tc808ad = static_cast<ApproveSignDataCtx77da62*>(ptr)->fromSrc().toUtf8(); Moc_PackedString { const_cast<char*>(tc808ad.prepend("WHITESPACE").constData()+10), tc808ad.size()-10 }; });
}

struct Moc_PackedString ApproveSignDataCtx77da62_FromSrcDefault(void* ptr)
{
	return ({ QByteArray t981347 = static_cast<ApproveSignDataCtx77da62*>(ptr)->fromSrcDefault().toUtf8(); Moc_PackedString { const_cast<char*>(t981347.prepend("WHITESPACE").constData()+10), t981347.size()-10 }; });
}

void ApproveSignDataCtx77da62_SetFromSrc(void* ptr, struct Moc_PackedString fromSrc)
{
	static_cast<ApproveSignDataCtx77da62*>(ptr)->setFromSrc(QString::fromUtf8(fromSrc.data, fromSrc.len));
}

void ApproveSignDataCtx77da62_SetFromSrcDefault(void* ptr, struct Moc_PackedString fromSrc)
{
	static_cast<ApproveSignDataCtx77da62*>(ptr)->setFromSrcDefault(QString::fromUtf8(fromSrc.data, fromSrc.len));
}

void ApproveSignDataCtx77da62_ConnectFromSrcChanged(void* ptr)
{
	QObject::connect(static_cast<ApproveSignDataCtx77da62*>(ptr), static_cast<void (ApproveSignDataCtx77da62::*)(QString)>(&ApproveSignDataCtx77da62::fromSrcChanged), static_cast<ApproveSignDataCtx77da62*>(ptr), static_cast<void (ApproveSignDataCtx77da62::*)(QString)>(&ApproveSignDataCtx77da62::Signal_FromSrcChanged));
}

void ApproveSignDataCtx77da62_DisconnectFromSrcChanged(void* ptr)
{
	QObject::disconnect(static_cast<ApproveSignDataCtx77da62*>(ptr), static_cast<void (ApproveSignDataCtx77da62::*)(QString)>(&ApproveSignDataCtx77da62::fromSrcChanged), static_cast<ApproveSignDataCtx77da62*>(ptr), static_cast<void (ApproveSignDataCtx77da62::*)(QString)>(&ApproveSignDataCtx77da62::Signal_FromSrcChanged));
}

void ApproveSignDataCtx77da62_FromSrcChanged(void* ptr, struct Moc_PackedString fromSrc)
{
	static_cast<ApproveSignDataCtx77da62*>(ptr)->fromSrcChanged(QString::fromUtf8(fromSrc.data, fromSrc.len));
}

int ApproveSignDataCtx77da62_ApproveSignDataCtx77da62_QRegisterMetaType()
{
	return qRegisterMetaType<ApproveSignDataCtx77da62*>();
}

int ApproveSignDataCtx77da62_ApproveSignDataCtx77da62_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<ApproveSignDataCtx77da62*>(const_cast<const char*>(typeName));
}

int ApproveSignDataCtx77da62_ApproveSignDataCtx77da62_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<ApproveSignDataCtx77da62>();
#else
	return 0;
#endif
}

int ApproveSignDataCtx77da62_ApproveSignDataCtx77da62_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<ApproveSignDataCtx77da62>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

void* ApproveSignDataCtx77da62___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void ApproveSignDataCtx77da62___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* ApproveSignDataCtx77da62___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* ApproveSignDataCtx77da62___findChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void ApproveSignDataCtx77da62___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* ApproveSignDataCtx77da62___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* ApproveSignDataCtx77da62___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void ApproveSignDataCtx77da62___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* ApproveSignDataCtx77da62___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* ApproveSignDataCtx77da62___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void ApproveSignDataCtx77da62___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* ApproveSignDataCtx77da62___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* ApproveSignDataCtx77da62___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void ApproveSignDataCtx77da62___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* ApproveSignDataCtx77da62___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* ApproveSignDataCtx77da62_NewApproveSignDataCtx(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new ApproveSignDataCtx77da62(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new ApproveSignDataCtx77da62(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new ApproveSignDataCtx77da62(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new ApproveSignDataCtx77da62(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new ApproveSignDataCtx77da62(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new ApproveSignDataCtx77da62(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new ApproveSignDataCtx77da62(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new ApproveSignDataCtx77da62(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new ApproveSignDataCtx77da62(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new ApproveSignDataCtx77da62(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new ApproveSignDataCtx77da62(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new ApproveSignDataCtx77da62(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new ApproveSignDataCtx77da62(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new ApproveSignDataCtx77da62(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new ApproveSignDataCtx77da62(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new ApproveSignDataCtx77da62(static_cast<QWindow*>(parent));
	} else {
		return new ApproveSignDataCtx77da62(static_cast<QObject*>(parent));
	}
}

void ApproveSignDataCtx77da62_DestroyApproveSignDataCtx(void* ptr)
{
	static_cast<ApproveSignDataCtx77da62*>(ptr)->~ApproveSignDataCtx77da62();
}

void ApproveSignDataCtx77da62_DestroyApproveSignDataCtxDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

char ApproveSignDataCtx77da62_EventDefault(void* ptr, void* e)
{
	return static_cast<ApproveSignDataCtx77da62*>(ptr)->QObject::event(static_cast<QEvent*>(e));
}

char ApproveSignDataCtx77da62_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<ApproveSignDataCtx77da62*>(ptr)->QObject::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void ApproveSignDataCtx77da62_ChildEventDefault(void* ptr, void* event)
{
	static_cast<ApproveSignDataCtx77da62*>(ptr)->QObject::childEvent(static_cast<QChildEvent*>(event));
}

void ApproveSignDataCtx77da62_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<ApproveSignDataCtx77da62*>(ptr)->QObject::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void ApproveSignDataCtx77da62_CustomEventDefault(void* ptr, void* event)
{
	static_cast<ApproveSignDataCtx77da62*>(ptr)->QObject::customEvent(static_cast<QEvent*>(event));
}

void ApproveSignDataCtx77da62_DeleteLaterDefault(void* ptr)
{
	static_cast<ApproveSignDataCtx77da62*>(ptr)->QObject::deleteLater();
}

void ApproveSignDataCtx77da62_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<ApproveSignDataCtx77da62*>(ptr)->QObject::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void ApproveSignDataCtx77da62_TimerEventDefault(void* ptr, void* event)
{
	static_cast<ApproveSignDataCtx77da62*>(ptr)->QObject::timerEvent(static_cast<QTimerEvent*>(event));
}



void CustomListModel77da62_ConnectClear(void* ptr)
{
	QObject::connect(static_cast<CustomListModel77da62*>(ptr), static_cast<void (CustomListModel77da62::*)()>(&CustomListModel77da62::clear), static_cast<CustomListModel77da62*>(ptr), static_cast<void (CustomListModel77da62::*)()>(&CustomListModel77da62::Signal_Clear));
}

void CustomListModel77da62_DisconnectClear(void* ptr)
{
	QObject::disconnect(static_cast<CustomListModel77da62*>(ptr), static_cast<void (CustomListModel77da62::*)()>(&CustomListModel77da62::clear), static_cast<CustomListModel77da62*>(ptr), static_cast<void (CustomListModel77da62::*)()>(&CustomListModel77da62::Signal_Clear));
}

void CustomListModel77da62_Clear(void* ptr)
{
	static_cast<CustomListModel77da62*>(ptr)->clear();
}

void CustomListModel77da62_ConnectAdd(void* ptr)
{
	QObject::connect(static_cast<CustomListModel77da62*>(ptr), static_cast<void (CustomListModel77da62::*)(quintptr)>(&CustomListModel77da62::add), static_cast<CustomListModel77da62*>(ptr), static_cast<void (CustomListModel77da62::*)(quintptr)>(&CustomListModel77da62::Signal_Add));
}

void CustomListModel77da62_DisconnectAdd(void* ptr)
{
	QObject::disconnect(static_cast<CustomListModel77da62*>(ptr), static_cast<void (CustomListModel77da62::*)(quintptr)>(&CustomListModel77da62::add), static_cast<CustomListModel77da62*>(ptr), static_cast<void (CustomListModel77da62::*)(quintptr)>(&CustomListModel77da62::Signal_Add));
}

void CustomListModel77da62_Add(void* ptr, uintptr_t account)
{
	static_cast<CustomListModel77da62*>(ptr)->add(account);
}

void CustomListModel77da62_CallbackFromQml(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<CustomListModel77da62*>(ptr), "callbackFromQml");
}

struct Moc_PackedString CustomListModel77da62_UpdateRequired(void* ptr)
{
	return ({ QByteArray t5b7d2c = static_cast<CustomListModel77da62*>(ptr)->updateRequired().toUtf8(); Moc_PackedString { const_cast<char*>(t5b7d2c.prepend("WHITESPACE").constData()+10), t5b7d2c.size()-10 }; });
}

struct Moc_PackedString CustomListModel77da62_UpdateRequiredDefault(void* ptr)
{
	return ({ QByteArray t4d3d34 = static_cast<CustomListModel77da62*>(ptr)->updateRequiredDefault().toUtf8(); Moc_PackedString { const_cast<char*>(t4d3d34.prepend("WHITESPACE").constData()+10), t4d3d34.size()-10 }; });
}

void CustomListModel77da62_SetUpdateRequired(void* ptr, struct Moc_PackedString updateRequired)
{
	static_cast<CustomListModel77da62*>(ptr)->setUpdateRequired(QString::fromUtf8(updateRequired.data, updateRequired.len));
}

void CustomListModel77da62_SetUpdateRequiredDefault(void* ptr, struct Moc_PackedString updateRequired)
{
	static_cast<CustomListModel77da62*>(ptr)->setUpdateRequiredDefault(QString::fromUtf8(updateRequired.data, updateRequired.len));
}

void CustomListModel77da62_ConnectUpdateRequiredChanged(void* ptr)
{
	QObject::connect(static_cast<CustomListModel77da62*>(ptr), static_cast<void (CustomListModel77da62::*)(QString)>(&CustomListModel77da62::updateRequiredChanged), static_cast<CustomListModel77da62*>(ptr), static_cast<void (CustomListModel77da62::*)(QString)>(&CustomListModel77da62::Signal_UpdateRequiredChanged));
}

void CustomListModel77da62_DisconnectUpdateRequiredChanged(void* ptr)
{
	QObject::disconnect(static_cast<CustomListModel77da62*>(ptr), static_cast<void (CustomListModel77da62::*)(QString)>(&CustomListModel77da62::updateRequiredChanged), static_cast<CustomListModel77da62*>(ptr), static_cast<void (CustomListModel77da62::*)(QString)>(&CustomListModel77da62::Signal_UpdateRequiredChanged));
}

void CustomListModel77da62_UpdateRequiredChanged(void* ptr, struct Moc_PackedString updateRequired)
{
	static_cast<CustomListModel77da62*>(ptr)->updateRequiredChanged(QString::fromUtf8(updateRequired.data, updateRequired.len));
}

int CustomListModel77da62_CustomListModel77da62_QRegisterMetaType()
{
	return qRegisterMetaType<CustomListModel77da62*>();
}

int CustomListModel77da62_CustomListModel77da62_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<CustomListModel77da62*>(const_cast<const char*>(typeName));
}

int CustomListModel77da62_CustomListModel77da62_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<CustomListModel77da62>();
#else
	return 0;
#endif
}

int CustomListModel77da62_CustomListModel77da62_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<CustomListModel77da62>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

int CustomListModel77da62_____setItemData_roles_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void CustomListModel77da62_____setItemData_roles_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* CustomListModel77da62_____setItemData_roles_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int CustomListModel77da62_____roleNames_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void CustomListModel77da62_____roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* CustomListModel77da62_____roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int CustomListModel77da62_____itemData_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void CustomListModel77da62_____itemData_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* CustomListModel77da62_____itemData_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

void* CustomListModel77da62___setItemData_roles_atList(void* ptr, int v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<int, QVariant>*>(ptr)->value(v); if (i == static_cast<QMap<int, QVariant>*>(ptr)->size()-1) { static_cast<QMap<int, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void CustomListModel77da62___setItemData_roles_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<int, QVariant>*>(ptr)->insert(key, *static_cast<QVariant*>(i));
}

void* CustomListModel77da62___setItemData_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<int, QVariant>();
}

struct Moc_PackedList CustomListModel77da62___setItemData_roles_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QMap<int, QVariant>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* CustomListModel77da62___changePersistentIndexList_from_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void CustomListModel77da62___changePersistentIndexList_from_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* CustomListModel77da62___changePersistentIndexList_from_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* CustomListModel77da62___changePersistentIndexList_to_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void CustomListModel77da62___changePersistentIndexList_to_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* CustomListModel77da62___changePersistentIndexList_to_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

int CustomListModel77da62___dataChanged_roles_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QVector<int>*>(ptr)->at(i); if (i == static_cast<QVector<int>*>(ptr)->size()-1) { static_cast<QVector<int>*>(ptr)->~QVector(); free(ptr); }; tmp; });
}

void CustomListModel77da62___dataChanged_roles_setList(void* ptr, int i)
{
	static_cast<QVector<int>*>(ptr)->append(i);
}

void* CustomListModel77da62___dataChanged_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<int>();
}

void* CustomListModel77da62___layoutAboutToBeChanged_parents_atList(void* ptr, int i)
{
	return new QPersistentModelIndex(({QPersistentModelIndex tmp = static_cast<QList<QPersistentModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QPersistentModelIndex>*>(ptr)->size()-1) { static_cast<QList<QPersistentModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void CustomListModel77da62___layoutAboutToBeChanged_parents_setList(void* ptr, void* i)
{
	static_cast<QList<QPersistentModelIndex>*>(ptr)->append(*static_cast<QPersistentModelIndex*>(i));
}

void* CustomListModel77da62___layoutAboutToBeChanged_parents_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QPersistentModelIndex>();
}

void* CustomListModel77da62___layoutChanged_parents_atList(void* ptr, int i)
{
	return new QPersistentModelIndex(({QPersistentModelIndex tmp = static_cast<QList<QPersistentModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QPersistentModelIndex>*>(ptr)->size()-1) { static_cast<QList<QPersistentModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void CustomListModel77da62___layoutChanged_parents_setList(void* ptr, void* i)
{
	static_cast<QList<QPersistentModelIndex>*>(ptr)->append(*static_cast<QPersistentModelIndex*>(i));
}

void* CustomListModel77da62___layoutChanged_parents_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QPersistentModelIndex>();
}

void* CustomListModel77da62___roleNames_atList(void* ptr, int v, int i)
{
	return new QByteArray(({ QByteArray tmp = static_cast<QHash<int, QByteArray>*>(ptr)->value(v); if (i == static_cast<QHash<int, QByteArray>*>(ptr)->size()-1) { static_cast<QHash<int, QByteArray>*>(ptr)->~QHash(); free(ptr); }; tmp; }));
}

void CustomListModel77da62___roleNames_setList(void* ptr, int key, void* i)
{
	static_cast<QHash<int, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* CustomListModel77da62___roleNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QHash<int, QByteArray>();
}

struct Moc_PackedList CustomListModel77da62___roleNames_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QHash<int, QByteArray>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* CustomListModel77da62___itemData_atList(void* ptr, int v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<int, QVariant>*>(ptr)->value(v); if (i == static_cast<QMap<int, QVariant>*>(ptr)->size()-1) { static_cast<QMap<int, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void CustomListModel77da62___itemData_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<int, QVariant>*>(ptr)->insert(key, *static_cast<QVariant*>(i));
}

void* CustomListModel77da62___itemData_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<int, QVariant>();
}

struct Moc_PackedList CustomListModel77da62___itemData_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QMap<int, QVariant>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* CustomListModel77da62___mimeData_indexes_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void CustomListModel77da62___mimeData_indexes_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* CustomListModel77da62___mimeData_indexes_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* CustomListModel77da62___match_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void CustomListModel77da62___match_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* CustomListModel77da62___match_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* CustomListModel77da62___persistentIndexList_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void CustomListModel77da62___persistentIndexList_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* CustomListModel77da62___persistentIndexList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

int CustomListModel77da62_____doSetRoleNames_roleNames_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void CustomListModel77da62_____doSetRoleNames_roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* CustomListModel77da62_____doSetRoleNames_roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int CustomListModel77da62_____setRoleNames_roleNames_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void CustomListModel77da62_____setRoleNames_roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* CustomListModel77da62_____setRoleNames_roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

void* CustomListModel77da62___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void CustomListModel77da62___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* CustomListModel77da62___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* CustomListModel77da62___findChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void CustomListModel77da62___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* CustomListModel77da62___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* CustomListModel77da62___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void CustomListModel77da62___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* CustomListModel77da62___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* CustomListModel77da62___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void CustomListModel77da62___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* CustomListModel77da62___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* CustomListModel77da62___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void CustomListModel77da62___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* CustomListModel77da62___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* CustomListModel77da62_NewCustomListModel(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new CustomListModel77da62(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new CustomListModel77da62(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new CustomListModel77da62(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new CustomListModel77da62(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new CustomListModel77da62(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new CustomListModel77da62(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new CustomListModel77da62(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new CustomListModel77da62(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new CustomListModel77da62(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new CustomListModel77da62(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new CustomListModel77da62(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new CustomListModel77da62(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new CustomListModel77da62(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new CustomListModel77da62(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new CustomListModel77da62(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new CustomListModel77da62(static_cast<QWindow*>(parent));
	} else {
		return new CustomListModel77da62(static_cast<QObject*>(parent));
	}
}

void CustomListModel77da62_DestroyCustomListModel(void* ptr)
{
	static_cast<CustomListModel77da62*>(ptr)->~CustomListModel77da62();
}

void CustomListModel77da62_DestroyCustomListModelDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

char CustomListModel77da62_DropMimeDataDefault(void* ptr, void* data, long long action, int row, int column, void* parent)
{
	return static_cast<CustomListModel77da62*>(ptr)->QAbstractListModel::dropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

void* CustomListModel77da62_IndexDefault(void* ptr, int row, int column, void* parent)
{
	return new QModelIndex(static_cast<CustomListModel77da62*>(ptr)->QAbstractListModel::index(row, column, *static_cast<QModelIndex*>(parent)));
}

void* CustomListModel77da62_SiblingDefault(void* ptr, int row, int column, void* idx)
{
	return new QModelIndex(static_cast<CustomListModel77da62*>(ptr)->QAbstractListModel::sibling(row, column, *static_cast<QModelIndex*>(idx)));
}

long long CustomListModel77da62_FlagsDefault(void* ptr, void* index)
{
	return static_cast<CustomListModel77da62*>(ptr)->QAbstractListModel::flags(*static_cast<QModelIndex*>(index));
}



char CustomListModel77da62_InsertColumnsDefault(void* ptr, int column, int count, void* parent)
{
	return static_cast<CustomListModel77da62*>(ptr)->QAbstractListModel::insertColumns(column, count, *static_cast<QModelIndex*>(parent));
}

char CustomListModel77da62_InsertRowsDefault(void* ptr, int row, int count, void* parent)
{
	return static_cast<CustomListModel77da62*>(ptr)->QAbstractListModel::insertRows(row, count, *static_cast<QModelIndex*>(parent));
}

char CustomListModel77da62_MoveColumnsDefault(void* ptr, void* sourceParent, int sourceColumn, int count, void* destinationParent, int destinationChild)
{
	return static_cast<CustomListModel77da62*>(ptr)->QAbstractListModel::moveColumns(*static_cast<QModelIndex*>(sourceParent), sourceColumn, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
}

char CustomListModel77da62_MoveRowsDefault(void* ptr, void* sourceParent, int sourceRow, int count, void* destinationParent, int destinationChild)
{
	return static_cast<CustomListModel77da62*>(ptr)->QAbstractListModel::moveRows(*static_cast<QModelIndex*>(sourceParent), sourceRow, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
}

char CustomListModel77da62_RemoveColumnsDefault(void* ptr, int column, int count, void* parent)
{
	return static_cast<CustomListModel77da62*>(ptr)->QAbstractListModel::removeColumns(column, count, *static_cast<QModelIndex*>(parent));
}

char CustomListModel77da62_RemoveRowsDefault(void* ptr, int row, int count, void* parent)
{
	return static_cast<CustomListModel77da62*>(ptr)->QAbstractListModel::removeRows(row, count, *static_cast<QModelIndex*>(parent));
}

char CustomListModel77da62_SetDataDefault(void* ptr, void* index, void* value, int role)
{
	return static_cast<CustomListModel77da62*>(ptr)->QAbstractListModel::setData(*static_cast<QModelIndex*>(index), *static_cast<QVariant*>(value), role);
}

char CustomListModel77da62_SetHeaderDataDefault(void* ptr, int section, long long orientation, void* value, int role)
{
	return static_cast<CustomListModel77da62*>(ptr)->QAbstractListModel::setHeaderData(section, static_cast<Qt::Orientation>(orientation), *static_cast<QVariant*>(value), role);
}

char CustomListModel77da62_SetItemDataDefault(void* ptr, void* index, void* roles)
{
	return static_cast<CustomListModel77da62*>(ptr)->QAbstractListModel::setItemData(*static_cast<QModelIndex*>(index), *static_cast<QMap<int, QVariant>*>(roles));
}

char CustomListModel77da62_SubmitDefault(void* ptr)
{
	return static_cast<CustomListModel77da62*>(ptr)->QAbstractListModel::submit();
}

void CustomListModel77da62_FetchMoreDefault(void* ptr, void* parent)
{
	static_cast<CustomListModel77da62*>(ptr)->QAbstractListModel::fetchMore(*static_cast<QModelIndex*>(parent));
}

void CustomListModel77da62_ResetInternalDataDefault(void* ptr)
{
	static_cast<CustomListModel77da62*>(ptr)->QAbstractListModel::resetInternalData();
}

void CustomListModel77da62_RevertDefault(void* ptr)
{
	static_cast<CustomListModel77da62*>(ptr)->QAbstractListModel::revert();
}

void CustomListModel77da62_SortDefault(void* ptr, int column, long long order)
{
	static_cast<CustomListModel77da62*>(ptr)->QAbstractListModel::sort(column, static_cast<Qt::SortOrder>(order));
}

struct Moc_PackedList CustomListModel77da62_RoleNamesDefault(void* ptr)
{
	return ({ QHash<int, QByteArray>* tmpValue = new QHash<int, QByteArray>(static_cast<CustomListModel77da62*>(ptr)->QAbstractListModel::roleNames()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

struct Moc_PackedList CustomListModel77da62_ItemDataDefault(void* ptr, void* index)
{
	return ({ QMap<int, QVariant>* tmpValue = new QMap<int, QVariant>(static_cast<CustomListModel77da62*>(ptr)->QAbstractListModel::itemData(*static_cast<QModelIndex*>(index))); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* CustomListModel77da62_MimeDataDefault(void* ptr, void* indexes)
{
	return static_cast<CustomListModel77da62*>(ptr)->QAbstractListModel::mimeData(({ QList<QModelIndex>* tmpP = static_cast<QList<QModelIndex>*>(indexes); QList<QModelIndex> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

void* CustomListModel77da62_BuddyDefault(void* ptr, void* index)
{
	return new QModelIndex(static_cast<CustomListModel77da62*>(ptr)->QAbstractListModel::buddy(*static_cast<QModelIndex*>(index)));
}

void* CustomListModel77da62_ParentDefault(void* ptr, void* index)
{
	return new QModelIndex(static_cast<CustomListModel77da62*>(ptr)->QAbstractListModel::parent(*static_cast<QModelIndex*>(index)));
}

struct Moc_PackedList CustomListModel77da62_MatchDefault(void* ptr, void* start, int role, void* value, int hits, long long flags)
{
	return ({ QList<QModelIndex>* tmpValue = new QList<QModelIndex>(static_cast<CustomListModel77da62*>(ptr)->QAbstractListModel::match(*static_cast<QModelIndex*>(start), role, *static_cast<QVariant*>(value), hits, static_cast<Qt::MatchFlag>(flags))); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* CustomListModel77da62_SpanDefault(void* ptr, void* index)
{
	return ({ QSize tmpValue = static_cast<CustomListModel77da62*>(ptr)->QAbstractListModel::span(*static_cast<QModelIndex*>(index)); new QSize(tmpValue.width(), tmpValue.height()); });
}

struct Moc_PackedString CustomListModel77da62_MimeTypesDefault(void* ptr)
{
	return ({ QByteArray t9c6749 = static_cast<CustomListModel77da62*>(ptr)->QAbstractListModel::mimeTypes().join("|").toUtf8(); Moc_PackedString { const_cast<char*>(t9c6749.prepend("WHITESPACE").constData()+10), t9c6749.size()-10 }; });
}

void* CustomListModel77da62_DataDefault(void* ptr, void* index, int role)
{
	Q_UNUSED(ptr);
	Q_UNUSED(index);
	Q_UNUSED(role);

}

void* CustomListModel77da62_HeaderDataDefault(void* ptr, int section, long long orientation, int role)
{
	return new QVariant(static_cast<CustomListModel77da62*>(ptr)->QAbstractListModel::headerData(section, static_cast<Qt::Orientation>(orientation), role));
}

long long CustomListModel77da62_SupportedDragActionsDefault(void* ptr)
{
	return static_cast<CustomListModel77da62*>(ptr)->QAbstractListModel::supportedDragActions();
}

long long CustomListModel77da62_SupportedDropActionsDefault(void* ptr)
{
	return static_cast<CustomListModel77da62*>(ptr)->QAbstractListModel::supportedDropActions();
}

char CustomListModel77da62_CanDropMimeDataDefault(void* ptr, void* data, long long action, int row, int column, void* parent)
{
	return static_cast<CustomListModel77da62*>(ptr)->QAbstractListModel::canDropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

char CustomListModel77da62_CanFetchMoreDefault(void* ptr, void* parent)
{
	return static_cast<CustomListModel77da62*>(ptr)->QAbstractListModel::canFetchMore(*static_cast<QModelIndex*>(parent));
}

char CustomListModel77da62_HasChildrenDefault(void* ptr, void* parent)
{
	return static_cast<CustomListModel77da62*>(ptr)->QAbstractListModel::hasChildren(*static_cast<QModelIndex*>(parent));
}

int CustomListModel77da62_ColumnCountDefault(void* ptr, void* parent)
{
	return static_cast<CustomListModel77da62*>(ptr)->QAbstractListModel::columnCount(*static_cast<QModelIndex*>(parent));
}

int CustomListModel77da62_RowCountDefault(void* ptr, void* parent)
{
	Q_UNUSED(ptr);
	Q_UNUSED(parent);

}

char CustomListModel77da62_EventDefault(void* ptr, void* e)
{
	return static_cast<CustomListModel77da62*>(ptr)->QAbstractListModel::event(static_cast<QEvent*>(e));
}

char CustomListModel77da62_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<CustomListModel77da62*>(ptr)->QAbstractListModel::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void CustomListModel77da62_ChildEventDefault(void* ptr, void* event)
{
	static_cast<CustomListModel77da62*>(ptr)->QAbstractListModel::childEvent(static_cast<QChildEvent*>(event));
}

void CustomListModel77da62_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<CustomListModel77da62*>(ptr)->QAbstractListModel::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void CustomListModel77da62_CustomEventDefault(void* ptr, void* event)
{
	static_cast<CustomListModel77da62*>(ptr)->QAbstractListModel::customEvent(static_cast<QEvent*>(event));
}

void CustomListModel77da62_DeleteLaterDefault(void* ptr)
{
	static_cast<CustomListModel77da62*>(ptr)->QAbstractListModel::deleteLater();
}

void CustomListModel77da62_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<CustomListModel77da62*>(ptr)->QAbstractListModel::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void CustomListModel77da62_TimerEventDefault(void* ptr, void* event)
{
	static_cast<CustomListModel77da62*>(ptr)->QAbstractListModel::timerEvent(static_cast<QTimerEvent*>(event));
}

void TxListAccountsModel77da62_ConnectAdd(void* ptr)
{
	QObject::connect(static_cast<TxListAccountsModel77da62*>(ptr), static_cast<void (TxListAccountsModel77da62::*)(QString)>(&TxListAccountsModel77da62::add), static_cast<TxListAccountsModel77da62*>(ptr), static_cast<void (TxListAccountsModel77da62::*)(QString)>(&TxListAccountsModel77da62::Signal_Add));
}

void TxListAccountsModel77da62_DisconnectAdd(void* ptr)
{
	QObject::disconnect(static_cast<TxListAccountsModel77da62*>(ptr), static_cast<void (TxListAccountsModel77da62::*)(QString)>(&TxListAccountsModel77da62::add), static_cast<TxListAccountsModel77da62*>(ptr), static_cast<void (TxListAccountsModel77da62::*)(QString)>(&TxListAccountsModel77da62::Signal_Add));
}

void TxListAccountsModel77da62_Add(void* ptr, struct Moc_PackedString tx)
{
	static_cast<TxListAccountsModel77da62*>(ptr)->add(QString::fromUtf8(tx.data, tx.len));
}

int TxListAccountsModel77da62_TxListAccountsModel77da62_QRegisterMetaType()
{
	return qRegisterMetaType<TxListAccountsModel77da62*>();
}

int TxListAccountsModel77da62_TxListAccountsModel77da62_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<TxListAccountsModel77da62*>(const_cast<const char*>(typeName));
}

int TxListAccountsModel77da62_TxListAccountsModel77da62_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<TxListAccountsModel77da62>();
#else
	return 0;
#endif
}

int TxListAccountsModel77da62_TxListAccountsModel77da62_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<TxListAccountsModel77da62>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

int TxListAccountsModel77da62_____setItemData_roles_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void TxListAccountsModel77da62_____setItemData_roles_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* TxListAccountsModel77da62_____setItemData_roles_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int TxListAccountsModel77da62_____roleNames_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void TxListAccountsModel77da62_____roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* TxListAccountsModel77da62_____roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int TxListAccountsModel77da62_____itemData_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void TxListAccountsModel77da62_____itemData_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* TxListAccountsModel77da62_____itemData_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

void* TxListAccountsModel77da62___setItemData_roles_atList(void* ptr, int v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<int, QVariant>*>(ptr)->value(v); if (i == static_cast<QMap<int, QVariant>*>(ptr)->size()-1) { static_cast<QMap<int, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void TxListAccountsModel77da62___setItemData_roles_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<int, QVariant>*>(ptr)->insert(key, *static_cast<QVariant*>(i));
}

void* TxListAccountsModel77da62___setItemData_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<int, QVariant>();
}

struct Moc_PackedList TxListAccountsModel77da62___setItemData_roles_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QMap<int, QVariant>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* TxListAccountsModel77da62___changePersistentIndexList_from_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void TxListAccountsModel77da62___changePersistentIndexList_from_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* TxListAccountsModel77da62___changePersistentIndexList_from_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* TxListAccountsModel77da62___changePersistentIndexList_to_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void TxListAccountsModel77da62___changePersistentIndexList_to_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* TxListAccountsModel77da62___changePersistentIndexList_to_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

int TxListAccountsModel77da62___dataChanged_roles_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QVector<int>*>(ptr)->at(i); if (i == static_cast<QVector<int>*>(ptr)->size()-1) { static_cast<QVector<int>*>(ptr)->~QVector(); free(ptr); }; tmp; });
}

void TxListAccountsModel77da62___dataChanged_roles_setList(void* ptr, int i)
{
	static_cast<QVector<int>*>(ptr)->append(i);
}

void* TxListAccountsModel77da62___dataChanged_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<int>();
}

void* TxListAccountsModel77da62___layoutAboutToBeChanged_parents_atList(void* ptr, int i)
{
	return new QPersistentModelIndex(({QPersistentModelIndex tmp = static_cast<QList<QPersistentModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QPersistentModelIndex>*>(ptr)->size()-1) { static_cast<QList<QPersistentModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void TxListAccountsModel77da62___layoutAboutToBeChanged_parents_setList(void* ptr, void* i)
{
	static_cast<QList<QPersistentModelIndex>*>(ptr)->append(*static_cast<QPersistentModelIndex*>(i));
}

void* TxListAccountsModel77da62___layoutAboutToBeChanged_parents_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QPersistentModelIndex>();
}

void* TxListAccountsModel77da62___layoutChanged_parents_atList(void* ptr, int i)
{
	return new QPersistentModelIndex(({QPersistentModelIndex tmp = static_cast<QList<QPersistentModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QPersistentModelIndex>*>(ptr)->size()-1) { static_cast<QList<QPersistentModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void TxListAccountsModel77da62___layoutChanged_parents_setList(void* ptr, void* i)
{
	static_cast<QList<QPersistentModelIndex>*>(ptr)->append(*static_cast<QPersistentModelIndex*>(i));
}

void* TxListAccountsModel77da62___layoutChanged_parents_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QPersistentModelIndex>();
}

void* TxListAccountsModel77da62___roleNames_atList(void* ptr, int v, int i)
{
	return new QByteArray(({ QByteArray tmp = static_cast<QHash<int, QByteArray>*>(ptr)->value(v); if (i == static_cast<QHash<int, QByteArray>*>(ptr)->size()-1) { static_cast<QHash<int, QByteArray>*>(ptr)->~QHash(); free(ptr); }; tmp; }));
}

void TxListAccountsModel77da62___roleNames_setList(void* ptr, int key, void* i)
{
	static_cast<QHash<int, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* TxListAccountsModel77da62___roleNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QHash<int, QByteArray>();
}

struct Moc_PackedList TxListAccountsModel77da62___roleNames_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QHash<int, QByteArray>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* TxListAccountsModel77da62___itemData_atList(void* ptr, int v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<int, QVariant>*>(ptr)->value(v); if (i == static_cast<QMap<int, QVariant>*>(ptr)->size()-1) { static_cast<QMap<int, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void TxListAccountsModel77da62___itemData_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<int, QVariant>*>(ptr)->insert(key, *static_cast<QVariant*>(i));
}

void* TxListAccountsModel77da62___itemData_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<int, QVariant>();
}

struct Moc_PackedList TxListAccountsModel77da62___itemData_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QMap<int, QVariant>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* TxListAccountsModel77da62___mimeData_indexes_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void TxListAccountsModel77da62___mimeData_indexes_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* TxListAccountsModel77da62___mimeData_indexes_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* TxListAccountsModel77da62___match_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void TxListAccountsModel77da62___match_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* TxListAccountsModel77da62___match_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* TxListAccountsModel77da62___persistentIndexList_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void TxListAccountsModel77da62___persistentIndexList_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* TxListAccountsModel77da62___persistentIndexList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

int TxListAccountsModel77da62_____doSetRoleNames_roleNames_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void TxListAccountsModel77da62_____doSetRoleNames_roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* TxListAccountsModel77da62_____doSetRoleNames_roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int TxListAccountsModel77da62_____setRoleNames_roleNames_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void TxListAccountsModel77da62_____setRoleNames_roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* TxListAccountsModel77da62_____setRoleNames_roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

void* TxListAccountsModel77da62___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void TxListAccountsModel77da62___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* TxListAccountsModel77da62___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* TxListAccountsModel77da62___findChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void TxListAccountsModel77da62___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* TxListAccountsModel77da62___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* TxListAccountsModel77da62___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void TxListAccountsModel77da62___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* TxListAccountsModel77da62___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* TxListAccountsModel77da62___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void TxListAccountsModel77da62___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* TxListAccountsModel77da62___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* TxListAccountsModel77da62___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void TxListAccountsModel77da62___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* TxListAccountsModel77da62___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* TxListAccountsModel77da62_NewTxListAccountsModel(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new TxListAccountsModel77da62(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new TxListAccountsModel77da62(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new TxListAccountsModel77da62(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new TxListAccountsModel77da62(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new TxListAccountsModel77da62(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new TxListAccountsModel77da62(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new TxListAccountsModel77da62(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new TxListAccountsModel77da62(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new TxListAccountsModel77da62(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new TxListAccountsModel77da62(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new TxListAccountsModel77da62(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new TxListAccountsModel77da62(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new TxListAccountsModel77da62(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new TxListAccountsModel77da62(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new TxListAccountsModel77da62(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new TxListAccountsModel77da62(static_cast<QWindow*>(parent));
	} else {
		return new TxListAccountsModel77da62(static_cast<QObject*>(parent));
	}
}

void TxListAccountsModel77da62_DestroyTxListAccountsModel(void* ptr)
{
	static_cast<TxListAccountsModel77da62*>(ptr)->~TxListAccountsModel77da62();
}

void TxListAccountsModel77da62_DestroyTxListAccountsModelDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

char TxListAccountsModel77da62_DropMimeDataDefault(void* ptr, void* data, long long action, int row, int column, void* parent)
{
	return static_cast<TxListAccountsModel77da62*>(ptr)->QAbstractListModel::dropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

void* TxListAccountsModel77da62_IndexDefault(void* ptr, int row, int column, void* parent)
{
	return new QModelIndex(static_cast<TxListAccountsModel77da62*>(ptr)->QAbstractListModel::index(row, column, *static_cast<QModelIndex*>(parent)));
}

void* TxListAccountsModel77da62_SiblingDefault(void* ptr, int row, int column, void* idx)
{
	return new QModelIndex(static_cast<TxListAccountsModel77da62*>(ptr)->QAbstractListModel::sibling(row, column, *static_cast<QModelIndex*>(idx)));
}

long long TxListAccountsModel77da62_FlagsDefault(void* ptr, void* index)
{
	return static_cast<TxListAccountsModel77da62*>(ptr)->QAbstractListModel::flags(*static_cast<QModelIndex*>(index));
}



char TxListAccountsModel77da62_InsertColumnsDefault(void* ptr, int column, int count, void* parent)
{
	return static_cast<TxListAccountsModel77da62*>(ptr)->QAbstractListModel::insertColumns(column, count, *static_cast<QModelIndex*>(parent));
}

char TxListAccountsModel77da62_InsertRowsDefault(void* ptr, int row, int count, void* parent)
{
	return static_cast<TxListAccountsModel77da62*>(ptr)->QAbstractListModel::insertRows(row, count, *static_cast<QModelIndex*>(parent));
}

char TxListAccountsModel77da62_MoveColumnsDefault(void* ptr, void* sourceParent, int sourceColumn, int count, void* destinationParent, int destinationChild)
{
	return static_cast<TxListAccountsModel77da62*>(ptr)->QAbstractListModel::moveColumns(*static_cast<QModelIndex*>(sourceParent), sourceColumn, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
}

char TxListAccountsModel77da62_MoveRowsDefault(void* ptr, void* sourceParent, int sourceRow, int count, void* destinationParent, int destinationChild)
{
	return static_cast<TxListAccountsModel77da62*>(ptr)->QAbstractListModel::moveRows(*static_cast<QModelIndex*>(sourceParent), sourceRow, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
}

char TxListAccountsModel77da62_RemoveColumnsDefault(void* ptr, int column, int count, void* parent)
{
	return static_cast<TxListAccountsModel77da62*>(ptr)->QAbstractListModel::removeColumns(column, count, *static_cast<QModelIndex*>(parent));
}

char TxListAccountsModel77da62_RemoveRowsDefault(void* ptr, int row, int count, void* parent)
{
	return static_cast<TxListAccountsModel77da62*>(ptr)->QAbstractListModel::removeRows(row, count, *static_cast<QModelIndex*>(parent));
}

char TxListAccountsModel77da62_SetDataDefault(void* ptr, void* index, void* value, int role)
{
	return static_cast<TxListAccountsModel77da62*>(ptr)->QAbstractListModel::setData(*static_cast<QModelIndex*>(index), *static_cast<QVariant*>(value), role);
}

char TxListAccountsModel77da62_SetHeaderDataDefault(void* ptr, int section, long long orientation, void* value, int role)
{
	return static_cast<TxListAccountsModel77da62*>(ptr)->QAbstractListModel::setHeaderData(section, static_cast<Qt::Orientation>(orientation), *static_cast<QVariant*>(value), role);
}

char TxListAccountsModel77da62_SetItemDataDefault(void* ptr, void* index, void* roles)
{
	return static_cast<TxListAccountsModel77da62*>(ptr)->QAbstractListModel::setItemData(*static_cast<QModelIndex*>(index), *static_cast<QMap<int, QVariant>*>(roles));
}

char TxListAccountsModel77da62_SubmitDefault(void* ptr)
{
	return static_cast<TxListAccountsModel77da62*>(ptr)->QAbstractListModel::submit();
}

void TxListAccountsModel77da62_FetchMoreDefault(void* ptr, void* parent)
{
	static_cast<TxListAccountsModel77da62*>(ptr)->QAbstractListModel::fetchMore(*static_cast<QModelIndex*>(parent));
}

void TxListAccountsModel77da62_ResetInternalDataDefault(void* ptr)
{
	static_cast<TxListAccountsModel77da62*>(ptr)->QAbstractListModel::resetInternalData();
}

void TxListAccountsModel77da62_RevertDefault(void* ptr)
{
	static_cast<TxListAccountsModel77da62*>(ptr)->QAbstractListModel::revert();
}

void TxListAccountsModel77da62_SortDefault(void* ptr, int column, long long order)
{
	static_cast<TxListAccountsModel77da62*>(ptr)->QAbstractListModel::sort(column, static_cast<Qt::SortOrder>(order));
}

struct Moc_PackedList TxListAccountsModel77da62_RoleNamesDefault(void* ptr)
{
	return ({ QHash<int, QByteArray>* tmpValue = new QHash<int, QByteArray>(static_cast<TxListAccountsModel77da62*>(ptr)->QAbstractListModel::roleNames()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

struct Moc_PackedList TxListAccountsModel77da62_ItemDataDefault(void* ptr, void* index)
{
	return ({ QMap<int, QVariant>* tmpValue = new QMap<int, QVariant>(static_cast<TxListAccountsModel77da62*>(ptr)->QAbstractListModel::itemData(*static_cast<QModelIndex*>(index))); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* TxListAccountsModel77da62_MimeDataDefault(void* ptr, void* indexes)
{
	return static_cast<TxListAccountsModel77da62*>(ptr)->QAbstractListModel::mimeData(({ QList<QModelIndex>* tmpP = static_cast<QList<QModelIndex>*>(indexes); QList<QModelIndex> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

void* TxListAccountsModel77da62_BuddyDefault(void* ptr, void* index)
{
	return new QModelIndex(static_cast<TxListAccountsModel77da62*>(ptr)->QAbstractListModel::buddy(*static_cast<QModelIndex*>(index)));
}

void* TxListAccountsModel77da62_ParentDefault(void* ptr, void* index)
{
	return new QModelIndex(static_cast<TxListAccountsModel77da62*>(ptr)->QAbstractListModel::parent(*static_cast<QModelIndex*>(index)));
}

struct Moc_PackedList TxListAccountsModel77da62_MatchDefault(void* ptr, void* start, int role, void* value, int hits, long long flags)
{
	return ({ QList<QModelIndex>* tmpValue = new QList<QModelIndex>(static_cast<TxListAccountsModel77da62*>(ptr)->QAbstractListModel::match(*static_cast<QModelIndex*>(start), role, *static_cast<QVariant*>(value), hits, static_cast<Qt::MatchFlag>(flags))); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* TxListAccountsModel77da62_SpanDefault(void* ptr, void* index)
{
	return ({ QSize tmpValue = static_cast<TxListAccountsModel77da62*>(ptr)->QAbstractListModel::span(*static_cast<QModelIndex*>(index)); new QSize(tmpValue.width(), tmpValue.height()); });
}

struct Moc_PackedString TxListAccountsModel77da62_MimeTypesDefault(void* ptr)
{
	return ({ QByteArray tb2f4e1 = static_cast<TxListAccountsModel77da62*>(ptr)->QAbstractListModel::mimeTypes().join("|").toUtf8(); Moc_PackedString { const_cast<char*>(tb2f4e1.prepend("WHITESPACE").constData()+10), tb2f4e1.size()-10 }; });
}

void* TxListAccountsModel77da62_DataDefault(void* ptr, void* index, int role)
{
	Q_UNUSED(ptr);
	Q_UNUSED(index);
	Q_UNUSED(role);

}

void* TxListAccountsModel77da62_HeaderDataDefault(void* ptr, int section, long long orientation, int role)
{
	return new QVariant(static_cast<TxListAccountsModel77da62*>(ptr)->QAbstractListModel::headerData(section, static_cast<Qt::Orientation>(orientation), role));
}

long long TxListAccountsModel77da62_SupportedDragActionsDefault(void* ptr)
{
	return static_cast<TxListAccountsModel77da62*>(ptr)->QAbstractListModel::supportedDragActions();
}

long long TxListAccountsModel77da62_SupportedDropActionsDefault(void* ptr)
{
	return static_cast<TxListAccountsModel77da62*>(ptr)->QAbstractListModel::supportedDropActions();
}

char TxListAccountsModel77da62_CanDropMimeDataDefault(void* ptr, void* data, long long action, int row, int column, void* parent)
{
	return static_cast<TxListAccountsModel77da62*>(ptr)->QAbstractListModel::canDropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

char TxListAccountsModel77da62_CanFetchMoreDefault(void* ptr, void* parent)
{
	return static_cast<TxListAccountsModel77da62*>(ptr)->QAbstractListModel::canFetchMore(*static_cast<QModelIndex*>(parent));
}

char TxListAccountsModel77da62_HasChildrenDefault(void* ptr, void* parent)
{
	return static_cast<TxListAccountsModel77da62*>(ptr)->QAbstractListModel::hasChildren(*static_cast<QModelIndex*>(parent));
}

int TxListAccountsModel77da62_ColumnCountDefault(void* ptr, void* parent)
{
	return static_cast<TxListAccountsModel77da62*>(ptr)->QAbstractListModel::columnCount(*static_cast<QModelIndex*>(parent));
}

int TxListAccountsModel77da62_RowCountDefault(void* ptr, void* parent)
{
	Q_UNUSED(ptr);
	Q_UNUSED(parent);

}

char TxListAccountsModel77da62_EventDefault(void* ptr, void* e)
{
	return static_cast<TxListAccountsModel77da62*>(ptr)->QAbstractListModel::event(static_cast<QEvent*>(e));
}

char TxListAccountsModel77da62_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<TxListAccountsModel77da62*>(ptr)->QAbstractListModel::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void TxListAccountsModel77da62_ChildEventDefault(void* ptr, void* event)
{
	static_cast<TxListAccountsModel77da62*>(ptr)->QAbstractListModel::childEvent(static_cast<QChildEvent*>(event));
}

void TxListAccountsModel77da62_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<TxListAccountsModel77da62*>(ptr)->QAbstractListModel::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void TxListAccountsModel77da62_CustomEventDefault(void* ptr, void* event)
{
	static_cast<TxListAccountsModel77da62*>(ptr)->QAbstractListModel::customEvent(static_cast<QEvent*>(event));
}

void TxListAccountsModel77da62_DeleteLaterDefault(void* ptr)
{
	static_cast<TxListAccountsModel77da62*>(ptr)->QAbstractListModel::deleteLater();
}

void TxListAccountsModel77da62_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<TxListAccountsModel77da62*>(ptr)->QAbstractListModel::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void TxListAccountsModel77da62_TimerEventDefault(void* ptr, void* event)
{
	static_cast<TxListAccountsModel77da62*>(ptr)->QAbstractListModel::timerEvent(static_cast<QTimerEvent*>(event));
}

void TxListModel77da62_ConnectClear(void* ptr)
{
	QObject::connect(static_cast<TxListModel77da62*>(ptr), static_cast<void (TxListModel77da62::*)()>(&TxListModel77da62::clear), static_cast<TxListModel77da62*>(ptr), static_cast<void (TxListModel77da62::*)()>(&TxListModel77da62::Signal_Clear));
}

void TxListModel77da62_DisconnectClear(void* ptr)
{
	QObject::disconnect(static_cast<TxListModel77da62*>(ptr), static_cast<void (TxListModel77da62::*)()>(&TxListModel77da62::clear), static_cast<TxListModel77da62*>(ptr), static_cast<void (TxListModel77da62::*)()>(&TxListModel77da62::Signal_Clear));
}

void TxListModel77da62_Clear(void* ptr)
{
	static_cast<TxListModel77da62*>(ptr)->clear();
}

void TxListModel77da62_ConnectAdd(void* ptr)
{
	QObject::connect(static_cast<TxListModel77da62*>(ptr), static_cast<void (TxListModel77da62::*)(quintptr)>(&TxListModel77da62::add), static_cast<TxListModel77da62*>(ptr), static_cast<void (TxListModel77da62::*)(quintptr)>(&TxListModel77da62::Signal_Add));
}

void TxListModel77da62_DisconnectAdd(void* ptr)
{
	QObject::disconnect(static_cast<TxListModel77da62*>(ptr), static_cast<void (TxListModel77da62::*)(quintptr)>(&TxListModel77da62::add), static_cast<TxListModel77da62*>(ptr), static_cast<void (TxListModel77da62::*)(quintptr)>(&TxListModel77da62::Signal_Add));
}

void TxListModel77da62_Add(void* ptr, uintptr_t tx)
{
	static_cast<TxListModel77da62*>(ptr)->add(tx);
}

void TxListModel77da62_ConnectRemove(void* ptr)
{
	QObject::connect(static_cast<TxListModel77da62*>(ptr), static_cast<void (TxListModel77da62::*)(qint32)>(&TxListModel77da62::remove), static_cast<TxListModel77da62*>(ptr), static_cast<void (TxListModel77da62::*)(qint32)>(&TxListModel77da62::Signal_Remove));
}

void TxListModel77da62_DisconnectRemove(void* ptr)
{
	QObject::disconnect(static_cast<TxListModel77da62*>(ptr), static_cast<void (TxListModel77da62::*)(qint32)>(&TxListModel77da62::remove), static_cast<TxListModel77da62*>(ptr), static_cast<void (TxListModel77da62::*)(qint32)>(&TxListModel77da62::Signal_Remove));
}

void TxListModel77da62_Remove(void* ptr, int i)
{
	static_cast<TxListModel77da62*>(ptr)->remove(i);
}

char TxListModel77da62_IsEmpty(void* ptr)
{
	return static_cast<TxListModel77da62*>(ptr)->isEmpty();
}

char TxListModel77da62_IsEmptyDefault(void* ptr)
{
	return static_cast<TxListModel77da62*>(ptr)->isEmptyDefault();
}

void TxListModel77da62_SetIsEmpty(void* ptr, char isEmpty)
{
	static_cast<TxListModel77da62*>(ptr)->setIsEmpty(isEmpty != 0);
}

void TxListModel77da62_SetIsEmptyDefault(void* ptr, char isEmpty)
{
	static_cast<TxListModel77da62*>(ptr)->setIsEmptyDefault(isEmpty != 0);
}

void TxListModel77da62_ConnectIsEmptyChanged(void* ptr)
{
	QObject::connect(static_cast<TxListModel77da62*>(ptr), static_cast<void (TxListModel77da62::*)(bool)>(&TxListModel77da62::isEmptyChanged), static_cast<TxListModel77da62*>(ptr), static_cast<void (TxListModel77da62::*)(bool)>(&TxListModel77da62::Signal_IsEmptyChanged));
}

void TxListModel77da62_DisconnectIsEmptyChanged(void* ptr)
{
	QObject::disconnect(static_cast<TxListModel77da62*>(ptr), static_cast<void (TxListModel77da62::*)(bool)>(&TxListModel77da62::isEmptyChanged), static_cast<TxListModel77da62*>(ptr), static_cast<void (TxListModel77da62::*)(bool)>(&TxListModel77da62::Signal_IsEmptyChanged));
}

void TxListModel77da62_IsEmptyChanged(void* ptr, char isEmpty)
{
	static_cast<TxListModel77da62*>(ptr)->isEmptyChanged(isEmpty != 0);
}

int TxListModel77da62_TxListModel77da62_QRegisterMetaType()
{
	return qRegisterMetaType<TxListModel77da62*>();
}

int TxListModel77da62_TxListModel77da62_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<TxListModel77da62*>(const_cast<const char*>(typeName));
}

int TxListModel77da62_TxListModel77da62_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<TxListModel77da62>();
#else
	return 0;
#endif
}

int TxListModel77da62_TxListModel77da62_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<TxListModel77da62>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

int TxListModel77da62_____setItemData_roles_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void TxListModel77da62_____setItemData_roles_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* TxListModel77da62_____setItemData_roles_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int TxListModel77da62_____roleNames_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void TxListModel77da62_____roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* TxListModel77da62_____roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int TxListModel77da62_____itemData_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void TxListModel77da62_____itemData_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* TxListModel77da62_____itemData_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

void* TxListModel77da62___setItemData_roles_atList(void* ptr, int v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<int, QVariant>*>(ptr)->value(v); if (i == static_cast<QMap<int, QVariant>*>(ptr)->size()-1) { static_cast<QMap<int, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void TxListModel77da62___setItemData_roles_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<int, QVariant>*>(ptr)->insert(key, *static_cast<QVariant*>(i));
}

void* TxListModel77da62___setItemData_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<int, QVariant>();
}

struct Moc_PackedList TxListModel77da62___setItemData_roles_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QMap<int, QVariant>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* TxListModel77da62___changePersistentIndexList_from_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void TxListModel77da62___changePersistentIndexList_from_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* TxListModel77da62___changePersistentIndexList_from_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* TxListModel77da62___changePersistentIndexList_to_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void TxListModel77da62___changePersistentIndexList_to_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* TxListModel77da62___changePersistentIndexList_to_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

int TxListModel77da62___dataChanged_roles_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QVector<int>*>(ptr)->at(i); if (i == static_cast<QVector<int>*>(ptr)->size()-1) { static_cast<QVector<int>*>(ptr)->~QVector(); free(ptr); }; tmp; });
}

void TxListModel77da62___dataChanged_roles_setList(void* ptr, int i)
{
	static_cast<QVector<int>*>(ptr)->append(i);
}

void* TxListModel77da62___dataChanged_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<int>();
}

void* TxListModel77da62___layoutAboutToBeChanged_parents_atList(void* ptr, int i)
{
	return new QPersistentModelIndex(({QPersistentModelIndex tmp = static_cast<QList<QPersistentModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QPersistentModelIndex>*>(ptr)->size()-1) { static_cast<QList<QPersistentModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void TxListModel77da62___layoutAboutToBeChanged_parents_setList(void* ptr, void* i)
{
	static_cast<QList<QPersistentModelIndex>*>(ptr)->append(*static_cast<QPersistentModelIndex*>(i));
}

void* TxListModel77da62___layoutAboutToBeChanged_parents_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QPersistentModelIndex>();
}

void* TxListModel77da62___layoutChanged_parents_atList(void* ptr, int i)
{
	return new QPersistentModelIndex(({QPersistentModelIndex tmp = static_cast<QList<QPersistentModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QPersistentModelIndex>*>(ptr)->size()-1) { static_cast<QList<QPersistentModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void TxListModel77da62___layoutChanged_parents_setList(void* ptr, void* i)
{
	static_cast<QList<QPersistentModelIndex>*>(ptr)->append(*static_cast<QPersistentModelIndex*>(i));
}

void* TxListModel77da62___layoutChanged_parents_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QPersistentModelIndex>();
}

void* TxListModel77da62___roleNames_atList(void* ptr, int v, int i)
{
	return new QByteArray(({ QByteArray tmp = static_cast<QHash<int, QByteArray>*>(ptr)->value(v); if (i == static_cast<QHash<int, QByteArray>*>(ptr)->size()-1) { static_cast<QHash<int, QByteArray>*>(ptr)->~QHash(); free(ptr); }; tmp; }));
}

void TxListModel77da62___roleNames_setList(void* ptr, int key, void* i)
{
	static_cast<QHash<int, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* TxListModel77da62___roleNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QHash<int, QByteArray>();
}

struct Moc_PackedList TxListModel77da62___roleNames_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QHash<int, QByteArray>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* TxListModel77da62___itemData_atList(void* ptr, int v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<int, QVariant>*>(ptr)->value(v); if (i == static_cast<QMap<int, QVariant>*>(ptr)->size()-1) { static_cast<QMap<int, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void TxListModel77da62___itemData_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<int, QVariant>*>(ptr)->insert(key, *static_cast<QVariant*>(i));
}

void* TxListModel77da62___itemData_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<int, QVariant>();
}

struct Moc_PackedList TxListModel77da62___itemData_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QMap<int, QVariant>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* TxListModel77da62___mimeData_indexes_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void TxListModel77da62___mimeData_indexes_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* TxListModel77da62___mimeData_indexes_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* TxListModel77da62___match_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void TxListModel77da62___match_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* TxListModel77da62___match_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* TxListModel77da62___persistentIndexList_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void TxListModel77da62___persistentIndexList_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* TxListModel77da62___persistentIndexList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

int TxListModel77da62_____doSetRoleNames_roleNames_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void TxListModel77da62_____doSetRoleNames_roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* TxListModel77da62_____doSetRoleNames_roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int TxListModel77da62_____setRoleNames_roleNames_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void TxListModel77da62_____setRoleNames_roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* TxListModel77da62_____setRoleNames_roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

void* TxListModel77da62___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void TxListModel77da62___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* TxListModel77da62___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* TxListModel77da62___findChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void TxListModel77da62___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* TxListModel77da62___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* TxListModel77da62___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void TxListModel77da62___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* TxListModel77da62___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* TxListModel77da62___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void TxListModel77da62___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* TxListModel77da62___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* TxListModel77da62___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void TxListModel77da62___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* TxListModel77da62___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* TxListModel77da62_NewTxListModel(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new TxListModel77da62(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new TxListModel77da62(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new TxListModel77da62(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new TxListModel77da62(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new TxListModel77da62(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new TxListModel77da62(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new TxListModel77da62(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new TxListModel77da62(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new TxListModel77da62(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new TxListModel77da62(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new TxListModel77da62(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new TxListModel77da62(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new TxListModel77da62(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new TxListModel77da62(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new TxListModel77da62(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new TxListModel77da62(static_cast<QWindow*>(parent));
	} else {
		return new TxListModel77da62(static_cast<QObject*>(parent));
	}
}

void TxListModel77da62_DestroyTxListModel(void* ptr)
{
	static_cast<TxListModel77da62*>(ptr)->~TxListModel77da62();
}

void TxListModel77da62_DestroyTxListModelDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

char TxListModel77da62_DropMimeDataDefault(void* ptr, void* data, long long action, int row, int column, void* parent)
{
	return static_cast<TxListModel77da62*>(ptr)->QAbstractListModel::dropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

void* TxListModel77da62_IndexDefault(void* ptr, int row, int column, void* parent)
{
	return new QModelIndex(static_cast<TxListModel77da62*>(ptr)->QAbstractListModel::index(row, column, *static_cast<QModelIndex*>(parent)));
}

void* TxListModel77da62_SiblingDefault(void* ptr, int row, int column, void* idx)
{
	return new QModelIndex(static_cast<TxListModel77da62*>(ptr)->QAbstractListModel::sibling(row, column, *static_cast<QModelIndex*>(idx)));
}

long long TxListModel77da62_FlagsDefault(void* ptr, void* index)
{
	return static_cast<TxListModel77da62*>(ptr)->QAbstractListModel::flags(*static_cast<QModelIndex*>(index));
}



char TxListModel77da62_InsertColumnsDefault(void* ptr, int column, int count, void* parent)
{
	return static_cast<TxListModel77da62*>(ptr)->QAbstractListModel::insertColumns(column, count, *static_cast<QModelIndex*>(parent));
}

char TxListModel77da62_InsertRowsDefault(void* ptr, int row, int count, void* parent)
{
	return static_cast<TxListModel77da62*>(ptr)->QAbstractListModel::insertRows(row, count, *static_cast<QModelIndex*>(parent));
}

char TxListModel77da62_MoveColumnsDefault(void* ptr, void* sourceParent, int sourceColumn, int count, void* destinationParent, int destinationChild)
{
	return static_cast<TxListModel77da62*>(ptr)->QAbstractListModel::moveColumns(*static_cast<QModelIndex*>(sourceParent), sourceColumn, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
}

char TxListModel77da62_MoveRowsDefault(void* ptr, void* sourceParent, int sourceRow, int count, void* destinationParent, int destinationChild)
{
	return static_cast<TxListModel77da62*>(ptr)->QAbstractListModel::moveRows(*static_cast<QModelIndex*>(sourceParent), sourceRow, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
}

char TxListModel77da62_RemoveColumnsDefault(void* ptr, int column, int count, void* parent)
{
	return static_cast<TxListModel77da62*>(ptr)->QAbstractListModel::removeColumns(column, count, *static_cast<QModelIndex*>(parent));
}

char TxListModel77da62_RemoveRowsDefault(void* ptr, int row, int count, void* parent)
{
	return static_cast<TxListModel77da62*>(ptr)->QAbstractListModel::removeRows(row, count, *static_cast<QModelIndex*>(parent));
}

char TxListModel77da62_SetDataDefault(void* ptr, void* index, void* value, int role)
{
	return static_cast<TxListModel77da62*>(ptr)->QAbstractListModel::setData(*static_cast<QModelIndex*>(index), *static_cast<QVariant*>(value), role);
}

char TxListModel77da62_SetHeaderDataDefault(void* ptr, int section, long long orientation, void* value, int role)
{
	return static_cast<TxListModel77da62*>(ptr)->QAbstractListModel::setHeaderData(section, static_cast<Qt::Orientation>(orientation), *static_cast<QVariant*>(value), role);
}

char TxListModel77da62_SetItemDataDefault(void* ptr, void* index, void* roles)
{
	return static_cast<TxListModel77da62*>(ptr)->QAbstractListModel::setItemData(*static_cast<QModelIndex*>(index), *static_cast<QMap<int, QVariant>*>(roles));
}

char TxListModel77da62_SubmitDefault(void* ptr)
{
	return static_cast<TxListModel77da62*>(ptr)->QAbstractListModel::submit();
}

void TxListModel77da62_FetchMoreDefault(void* ptr, void* parent)
{
	static_cast<TxListModel77da62*>(ptr)->QAbstractListModel::fetchMore(*static_cast<QModelIndex*>(parent));
}

void TxListModel77da62_ResetInternalDataDefault(void* ptr)
{
	static_cast<TxListModel77da62*>(ptr)->QAbstractListModel::resetInternalData();
}

void TxListModel77da62_RevertDefault(void* ptr)
{
	static_cast<TxListModel77da62*>(ptr)->QAbstractListModel::revert();
}

void TxListModel77da62_SortDefault(void* ptr, int column, long long order)
{
	static_cast<TxListModel77da62*>(ptr)->QAbstractListModel::sort(column, static_cast<Qt::SortOrder>(order));
}

struct Moc_PackedList TxListModel77da62_RoleNamesDefault(void* ptr)
{
	return ({ QHash<int, QByteArray>* tmpValue = new QHash<int, QByteArray>(static_cast<TxListModel77da62*>(ptr)->QAbstractListModel::roleNames()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

struct Moc_PackedList TxListModel77da62_ItemDataDefault(void* ptr, void* index)
{
	return ({ QMap<int, QVariant>* tmpValue = new QMap<int, QVariant>(static_cast<TxListModel77da62*>(ptr)->QAbstractListModel::itemData(*static_cast<QModelIndex*>(index))); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* TxListModel77da62_MimeDataDefault(void* ptr, void* indexes)
{
	return static_cast<TxListModel77da62*>(ptr)->QAbstractListModel::mimeData(({ QList<QModelIndex>* tmpP = static_cast<QList<QModelIndex>*>(indexes); QList<QModelIndex> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

void* TxListModel77da62_BuddyDefault(void* ptr, void* index)
{
	return new QModelIndex(static_cast<TxListModel77da62*>(ptr)->QAbstractListModel::buddy(*static_cast<QModelIndex*>(index)));
}

void* TxListModel77da62_ParentDefault(void* ptr, void* index)
{
	return new QModelIndex(static_cast<TxListModel77da62*>(ptr)->QAbstractListModel::parent(*static_cast<QModelIndex*>(index)));
}

struct Moc_PackedList TxListModel77da62_MatchDefault(void* ptr, void* start, int role, void* value, int hits, long long flags)
{
	return ({ QList<QModelIndex>* tmpValue = new QList<QModelIndex>(static_cast<TxListModel77da62*>(ptr)->QAbstractListModel::match(*static_cast<QModelIndex*>(start), role, *static_cast<QVariant*>(value), hits, static_cast<Qt::MatchFlag>(flags))); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* TxListModel77da62_SpanDefault(void* ptr, void* index)
{
	return ({ QSize tmpValue = static_cast<TxListModel77da62*>(ptr)->QAbstractListModel::span(*static_cast<QModelIndex*>(index)); new QSize(tmpValue.width(), tmpValue.height()); });
}

struct Moc_PackedString TxListModel77da62_MimeTypesDefault(void* ptr)
{
	return ({ QByteArray t525b22 = static_cast<TxListModel77da62*>(ptr)->QAbstractListModel::mimeTypes().join("|").toUtf8(); Moc_PackedString { const_cast<char*>(t525b22.prepend("WHITESPACE").constData()+10), t525b22.size()-10 }; });
}

void* TxListModel77da62_DataDefault(void* ptr, void* index, int role)
{
	Q_UNUSED(ptr);
	Q_UNUSED(index);
	Q_UNUSED(role);

}

void* TxListModel77da62_HeaderDataDefault(void* ptr, int section, long long orientation, int role)
{
	return new QVariant(static_cast<TxListModel77da62*>(ptr)->QAbstractListModel::headerData(section, static_cast<Qt::Orientation>(orientation), role));
}

long long TxListModel77da62_SupportedDragActionsDefault(void* ptr)
{
	return static_cast<TxListModel77da62*>(ptr)->QAbstractListModel::supportedDragActions();
}

long long TxListModel77da62_SupportedDropActionsDefault(void* ptr)
{
	return static_cast<TxListModel77da62*>(ptr)->QAbstractListModel::supportedDropActions();
}

char TxListModel77da62_CanDropMimeDataDefault(void* ptr, void* data, long long action, int row, int column, void* parent)
{
	return static_cast<TxListModel77da62*>(ptr)->QAbstractListModel::canDropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

char TxListModel77da62_CanFetchMoreDefault(void* ptr, void* parent)
{
	return static_cast<TxListModel77da62*>(ptr)->QAbstractListModel::canFetchMore(*static_cast<QModelIndex*>(parent));
}

char TxListModel77da62_HasChildrenDefault(void* ptr, void* parent)
{
	return static_cast<TxListModel77da62*>(ptr)->QAbstractListModel::hasChildren(*static_cast<QModelIndex*>(parent));
}

int TxListModel77da62_ColumnCountDefault(void* ptr, void* parent)
{
	return static_cast<TxListModel77da62*>(ptr)->QAbstractListModel::columnCount(*static_cast<QModelIndex*>(parent));
}

int TxListModel77da62_RowCountDefault(void* ptr, void* parent)
{
	Q_UNUSED(ptr);
	Q_UNUSED(parent);

}

char TxListModel77da62_EventDefault(void* ptr, void* e)
{
	return static_cast<TxListModel77da62*>(ptr)->QAbstractListModel::event(static_cast<QEvent*>(e));
}

char TxListModel77da62_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<TxListModel77da62*>(ptr)->QAbstractListModel::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void TxListModel77da62_ChildEventDefault(void* ptr, void* event)
{
	static_cast<TxListModel77da62*>(ptr)->QAbstractListModel::childEvent(static_cast<QChildEvent*>(event));
}

void TxListModel77da62_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<TxListModel77da62*>(ptr)->QAbstractListModel::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void TxListModel77da62_CustomEventDefault(void* ptr, void* event)
{
	static_cast<TxListModel77da62*>(ptr)->QAbstractListModel::customEvent(static_cast<QEvent*>(event));
}

void TxListModel77da62_DeleteLaterDefault(void* ptr)
{
	static_cast<TxListModel77da62*>(ptr)->QAbstractListModel::deleteLater();
}

void TxListModel77da62_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<TxListModel77da62*>(ptr)->QAbstractListModel::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void TxListModel77da62_TimerEventDefault(void* ptr, void* event)
{
	static_cast<TxListModel77da62*>(ptr)->QAbstractListModel::timerEvent(static_cast<QTimerEvent*>(event));
}

#include "moc_moc.h"
