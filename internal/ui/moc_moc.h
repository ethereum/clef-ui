/****************************************************************************
** Meta object code from reading C++ file 'moc.cpp'
**
** Created by: The Qt Meta Object Compiler version 67 (Qt 5.11.1)
**
** WARNING! All changes made in this file will be lost!
*****************************************************************************/

#include <QtCore/qbytearray.h>
#include <QtCore/qmetatype.h>
#include <QtCore/QList>
#if !defined(Q_MOC_OUTPUT_REVISION)
#error "The header file 'moc.cpp' doesn't include <QObject>."
#elif Q_MOC_OUTPUT_REVISION != 67
#error "This file was generated using the moc from 5.11.1. It"
#error "cannot be used with the include files from this version of Qt."
#error "(The moc has changed too much.)"
#endif

QT_BEGIN_MOC_NAMESPACE
QT_WARNING_PUSH
QT_WARNING_DISABLE_DEPRECATED
struct qt_meta_stringdata_CtxObject721036_t {
    QByteArrayData data[18];
    char stringdata0[180];
};
#define QT_MOC_LITERAL(idx, ofs, len) \
    Q_STATIC_BYTE_ARRAY_DATA_HEADER_INITIALIZER_WITH_OFFSET(len, \
    qptrdiff(offsetof(qt_meta_stringdata_CtxObject721036_t, stringdata0) + ofs \
        - idx * sizeof(QByteArrayData)) \
    )
static const qt_meta_stringdata_CtxObject721036_t qt_meta_stringdata_CtxObject721036 = {
    {
QT_MOC_LITERAL(0, 0, 15), // "CtxObject721036"
QT_MOC_LITERAL(1, 16, 7), // "clicked"
QT_MOC_LITERAL(2, 24, 0), // ""
QT_MOC_LITERAL(3, 25, 1), // "b"
QT_MOC_LITERAL(4, 27, 13), // "remoteChanged"
QT_MOC_LITERAL(5, 41, 6), // "remote"
QT_MOC_LITERAL(6, 48, 16), // "transportChanged"
QT_MOC_LITERAL(7, 65, 9), // "transport"
QT_MOC_LITERAL(8, 75, 15), // "endpointChanged"
QT_MOC_LITERAL(9, 91, 8), // "endpoint"
QT_MOC_LITERAL(10, 100, 11), // "fromChanged"
QT_MOC_LITERAL(11, 112, 4), // "from"
QT_MOC_LITERAL(12, 117, 14), // "messageChanged"
QT_MOC_LITERAL(13, 132, 7), // "message"
QT_MOC_LITERAL(14, 140, 14), // "rawDataChanged"
QT_MOC_LITERAL(15, 155, 7), // "rawData"
QT_MOC_LITERAL(16, 163, 11), // "hashChanged"
QT_MOC_LITERAL(17, 175, 4) // "hash"

    },
    "CtxObject721036\0clicked\0\0b\0remoteChanged\0"
    "remote\0transportChanged\0transport\0"
    "endpointChanged\0endpoint\0fromChanged\0"
    "from\0messageChanged\0message\0rawDataChanged\0"
    "rawData\0hashChanged\0hash"
};
#undef QT_MOC_LITERAL

static const uint qt_meta_data_CtxObject721036[] = {

 // content:
       7,       // revision
       0,       // classname
       0,    0, // classinfo
       8,   14, // methods
       7,   78, // properties
       0,    0, // enums/sets
       0,    0, // constructors
       0,       // flags
       8,       // signalCount

 // signals: name, argc, parameters, tag, flags
       1,    1,   54,    2, 0x06 /* Public */,
       4,    1,   57,    2, 0x06 /* Public */,
       6,    1,   60,    2, 0x06 /* Public */,
       8,    1,   63,    2, 0x06 /* Public */,
      10,    1,   66,    2, 0x06 /* Public */,
      12,    1,   69,    2, 0x06 /* Public */,
      14,    1,   72,    2, 0x06 /* Public */,
      16,    1,   75,    2, 0x06 /* Public */,

 // signals: parameters
    QMetaType::Void, QMetaType::Int,    3,
    QMetaType::Void, QMetaType::QString,    5,
    QMetaType::Void, QMetaType::QString,    7,
    QMetaType::Void, QMetaType::QString,    9,
    QMetaType::Void, QMetaType::QString,   11,
    QMetaType::Void, QMetaType::QString,   13,
    QMetaType::Void, QMetaType::QString,   15,
    QMetaType::Void, QMetaType::QString,   17,

 // properties: name, type, flags
       5, QMetaType::QString, 0x00495103,
       7, QMetaType::QString, 0x00495103,
       9, QMetaType::QString, 0x00495103,
      11, QMetaType::QString, 0x00495103,
      13, QMetaType::QString, 0x00495103,
      15, QMetaType::QString, 0x00495103,
      17, QMetaType::QString, 0x00495103,

 // properties: notify_signal_id
       1,
       2,
       3,
       4,
       5,
       6,
       7,

       0        // eod
};

void CtxObject721036::qt_static_metacall(QObject *_o, QMetaObject::Call _c, int _id, void **_a)
{
    if (_c == QMetaObject::InvokeMetaMethod) {
        CtxObject721036 *_t = static_cast<CtxObject721036 *>(_o);
        Q_UNUSED(_t)
        switch (_id) {
        case 0: _t->clicked((*reinterpret_cast< qint32(*)>(_a[1]))); break;
        case 1: _t->remoteChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 2: _t->transportChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 3: _t->endpointChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 4: _t->fromChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 5: _t->messageChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 6: _t->rawDataChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 7: _t->hashChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        default: ;
        }
    } else if (_c == QMetaObject::IndexOfMethod) {
        int *result = reinterpret_cast<int *>(_a[0]);
        {
            using _t = void (CtxObject721036::*)(qint32 );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&CtxObject721036::clicked)) {
                *result = 0;
                return;
            }
        }
        {
            using _t = void (CtxObject721036::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&CtxObject721036::remoteChanged)) {
                *result = 1;
                return;
            }
        }
        {
            using _t = void (CtxObject721036::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&CtxObject721036::transportChanged)) {
                *result = 2;
                return;
            }
        }
        {
            using _t = void (CtxObject721036::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&CtxObject721036::endpointChanged)) {
                *result = 3;
                return;
            }
        }
        {
            using _t = void (CtxObject721036::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&CtxObject721036::fromChanged)) {
                *result = 4;
                return;
            }
        }
        {
            using _t = void (CtxObject721036::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&CtxObject721036::messageChanged)) {
                *result = 5;
                return;
            }
        }
        {
            using _t = void (CtxObject721036::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&CtxObject721036::rawDataChanged)) {
                *result = 6;
                return;
            }
        }
        {
            using _t = void (CtxObject721036::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&CtxObject721036::hashChanged)) {
                *result = 7;
                return;
            }
        }
    }
#ifndef QT_NO_PROPERTIES
    else if (_c == QMetaObject::ReadProperty) {
        CtxObject721036 *_t = static_cast<CtxObject721036 *>(_o);
        Q_UNUSED(_t)
        void *_v = _a[0];
        switch (_id) {
        case 0: *reinterpret_cast< QString*>(_v) = _t->remote(); break;
        case 1: *reinterpret_cast< QString*>(_v) = _t->transport(); break;
        case 2: *reinterpret_cast< QString*>(_v) = _t->endpoint(); break;
        case 3: *reinterpret_cast< QString*>(_v) = _t->from(); break;
        case 4: *reinterpret_cast< QString*>(_v) = _t->message(); break;
        case 5: *reinterpret_cast< QString*>(_v) = _t->rawData(); break;
        case 6: *reinterpret_cast< QString*>(_v) = _t->hash(); break;
        default: break;
        }
    } else if (_c == QMetaObject::WriteProperty) {
        CtxObject721036 *_t = static_cast<CtxObject721036 *>(_o);
        Q_UNUSED(_t)
        void *_v = _a[0];
        switch (_id) {
        case 0: _t->setRemote(*reinterpret_cast< QString*>(_v)); break;
        case 1: _t->setTransport(*reinterpret_cast< QString*>(_v)); break;
        case 2: _t->setEndpoint(*reinterpret_cast< QString*>(_v)); break;
        case 3: _t->setFrom(*reinterpret_cast< QString*>(_v)); break;
        case 4: _t->setMessage(*reinterpret_cast< QString*>(_v)); break;
        case 5: _t->setRawData(*reinterpret_cast< QString*>(_v)); break;
        case 6: _t->setHash(*reinterpret_cast< QString*>(_v)); break;
        default: break;
        }
    } else if (_c == QMetaObject::ResetProperty) {
    }
#endif // QT_NO_PROPERTIES
}

QT_INIT_METAOBJECT const QMetaObject CtxObject721036::staticMetaObject = {
    { &QObject::staticMetaObject, qt_meta_stringdata_CtxObject721036.data,
      qt_meta_data_CtxObject721036,  qt_static_metacall, nullptr, nullptr}
};


const QMetaObject *CtxObject721036::metaObject() const
{
    return QObject::d_ptr->metaObject ? QObject::d_ptr->dynamicMetaObject() : &staticMetaObject;
}

void *CtxObject721036::qt_metacast(const char *_clname)
{
    if (!_clname) return nullptr;
    if (!strcmp(_clname, qt_meta_stringdata_CtxObject721036.stringdata0))
        return static_cast<void*>(this);
    return QObject::qt_metacast(_clname);
}

int CtxObject721036::qt_metacall(QMetaObject::Call _c, int _id, void **_a)
{
    _id = QObject::qt_metacall(_c, _id, _a);
    if (_id < 0)
        return _id;
    if (_c == QMetaObject::InvokeMetaMethod) {
        if (_id < 8)
            qt_static_metacall(this, _c, _id, _a);
        _id -= 8;
    } else if (_c == QMetaObject::RegisterMethodArgumentMetaType) {
        if (_id < 8)
            *reinterpret_cast<int*>(_a[0]) = -1;
        _id -= 8;
    }
#ifndef QT_NO_PROPERTIES
   else if (_c == QMetaObject::ReadProperty || _c == QMetaObject::WriteProperty
            || _c == QMetaObject::ResetProperty || _c == QMetaObject::RegisterPropertyMetaType) {
        qt_static_metacall(this, _c, _id, _a);
        _id -= 7;
    } else if (_c == QMetaObject::QueryPropertyDesignable) {
        _id -= 7;
    } else if (_c == QMetaObject::QueryPropertyScriptable) {
        _id -= 7;
    } else if (_c == QMetaObject::QueryPropertyStored) {
        _id -= 7;
    } else if (_c == QMetaObject::QueryPropertyEditable) {
        _id -= 7;
    } else if (_c == QMetaObject::QueryPropertyUser) {
        _id -= 7;
    }
#endif // QT_NO_PROPERTIES
    return _id;
}

// SIGNAL 0
void CtxObject721036::clicked(qint32 _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 0, _a);
}

// SIGNAL 1
void CtxObject721036::remoteChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 1, _a);
}

// SIGNAL 2
void CtxObject721036::transportChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 2, _a);
}

// SIGNAL 3
void CtxObject721036::endpointChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 3, _a);
}

// SIGNAL 4
void CtxObject721036::fromChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 4, _a);
}

// SIGNAL 5
void CtxObject721036::messageChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 5, _a);
}

// SIGNAL 6
void CtxObject721036::rawDataChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 6, _a);
}

// SIGNAL 7
void CtxObject721036::hashChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 7, _a);
}
struct qt_meta_stringdata_CustomListModel721036_t {
    QByteArrayData data[8];
    char stringdata0[59];
};
#define QT_MOC_LITERAL(idx, ofs, len) \
    Q_STATIC_BYTE_ARRAY_DATA_HEADER_INITIALIZER_WITH_OFFSET(len, \
    qptrdiff(offsetof(qt_meta_stringdata_CustomListModel721036_t, stringdata0) + ofs \
        - idx * sizeof(QByteArrayData)) \
    )
static const qt_meta_stringdata_CustomListModel721036_t qt_meta_stringdata_CustomListModel721036 = {
    {
QT_MOC_LITERAL(0, 0, 21), // "CustomListModel721036"
QT_MOC_LITERAL(1, 22, 6), // "remove"
QT_MOC_LITERAL(2, 29, 0), // ""
QT_MOC_LITERAL(3, 30, 3), // "add"
QT_MOC_LITERAL(4, 34, 3), // "obj"
QT_MOC_LITERAL(5, 38, 4), // "edit"
QT_MOC_LITERAL(6, 43, 7), // "address"
QT_MOC_LITERAL(7, 51, 7) // "checked"

    },
    "CustomListModel721036\0remove\0\0add\0obj\0"
    "edit\0address\0checked"
};
#undef QT_MOC_LITERAL

static const uint qt_meta_data_CustomListModel721036[] = {

 // content:
       7,       // revision
       0,       // classname
       0,    0, // classinfo
       3,   14, // methods
       0,    0, // properties
       0,    0, // enums/sets
       0,    0, // constructors
       0,       // flags
       3,       // signalCount

 // signals: name, argc, parameters, tag, flags
       1,    0,   29,    2, 0x06 /* Public */,
       3,    1,   30,    2, 0x06 /* Public */,
       5,    2,   33,    2, 0x06 /* Public */,

 // signals: parameters
    QMetaType::Void,
    QMetaType::Void, QMetaType::QVariantList,    4,
    QMetaType::Void, QMetaType::QString, QMetaType::QString,    6,    7,

       0        // eod
};

void CustomListModel721036::qt_static_metacall(QObject *_o, QMetaObject::Call _c, int _id, void **_a)
{
    if (_c == QMetaObject::InvokeMetaMethod) {
        CustomListModel721036 *_t = static_cast<CustomListModel721036 *>(_o);
        Q_UNUSED(_t)
        switch (_id) {
        case 0: _t->remove(); break;
        case 1: _t->add((*reinterpret_cast< QList<QVariant>(*)>(_a[1]))); break;
        case 2: _t->edit((*reinterpret_cast< QString(*)>(_a[1])),(*reinterpret_cast< QString(*)>(_a[2]))); break;
        default: ;
        }
    } else if (_c == QMetaObject::IndexOfMethod) {
        int *result = reinterpret_cast<int *>(_a[0]);
        {
            using _t = void (CustomListModel721036::*)();
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&CustomListModel721036::remove)) {
                *result = 0;
                return;
            }
        }
        {
            using _t = void (CustomListModel721036::*)(QList<QVariant> );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&CustomListModel721036::add)) {
                *result = 1;
                return;
            }
        }
        {
            using _t = void (CustomListModel721036::*)(QString , QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&CustomListModel721036::edit)) {
                *result = 2;
                return;
            }
        }
    }
}

QT_INIT_METAOBJECT const QMetaObject CustomListModel721036::staticMetaObject = {
    { &QAbstractListModel::staticMetaObject, qt_meta_stringdata_CustomListModel721036.data,
      qt_meta_data_CustomListModel721036,  qt_static_metacall, nullptr, nullptr}
};


const QMetaObject *CustomListModel721036::metaObject() const
{
    return QObject::d_ptr->metaObject ? QObject::d_ptr->dynamicMetaObject() : &staticMetaObject;
}

void *CustomListModel721036::qt_metacast(const char *_clname)
{
    if (!_clname) return nullptr;
    if (!strcmp(_clname, qt_meta_stringdata_CustomListModel721036.stringdata0))
        return static_cast<void*>(this);
    return QAbstractListModel::qt_metacast(_clname);
}

int CustomListModel721036::qt_metacall(QMetaObject::Call _c, int _id, void **_a)
{
    _id = QAbstractListModel::qt_metacall(_c, _id, _a);
    if (_id < 0)
        return _id;
    if (_c == QMetaObject::InvokeMetaMethod) {
        if (_id < 3)
            qt_static_metacall(this, _c, _id, _a);
        _id -= 3;
    } else if (_c == QMetaObject::RegisterMethodArgumentMetaType) {
        if (_id < 3)
            *reinterpret_cast<int*>(_a[0]) = -1;
        _id -= 3;
    }
    return _id;
}

// SIGNAL 0
void CustomListModel721036::remove()
{
    QMetaObject::activate(this, &staticMetaObject, 0, nullptr);
}

// SIGNAL 1
void CustomListModel721036::add(QList<QVariant> _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 1, _a);
}

// SIGNAL 2
void CustomListModel721036::edit(QString _t1, QString _t2)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)), const_cast<void*>(reinterpret_cast<const void*>(&_t2)) };
    QMetaObject::activate(this, &staticMetaObject, 2, _a);
}
QT_WARNING_POP
QT_END_MOC_NAMESPACE
