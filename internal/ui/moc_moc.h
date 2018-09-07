/****************************************************************************
** Meta object code from reading C++ file 'moc.cpp'
**
** Created by: The Qt Meta Object Compiler version 67 (Qt 5.11.1)
**
** WARNING! All changes made in this file will be lost!
*****************************************************************************/

#include <QtCore/qbytearray.h>
#include <QtCore/qmetatype.h>
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
struct qt_meta_stringdata_ApproveSignDataCtx721036_t {
    QByteArrayData data[18];
    char stringdata0[189];
};
#define QT_MOC_LITERAL(idx, ofs, len) \
    Q_STATIC_BYTE_ARRAY_DATA_HEADER_INITIALIZER_WITH_OFFSET(len, \
    qptrdiff(offsetof(qt_meta_stringdata_ApproveSignDataCtx721036_t, stringdata0) + ofs \
        - idx * sizeof(QByteArrayData)) \
    )
static const qt_meta_stringdata_ApproveSignDataCtx721036_t qt_meta_stringdata_ApproveSignDataCtx721036 = {
    {
QT_MOC_LITERAL(0, 0, 24), // "ApproveSignDataCtx721036"
QT_MOC_LITERAL(1, 25, 7), // "clicked"
QT_MOC_LITERAL(2, 33, 0), // ""
QT_MOC_LITERAL(3, 34, 1), // "b"
QT_MOC_LITERAL(4, 36, 13), // "remoteChanged"
QT_MOC_LITERAL(5, 50, 6), // "remote"
QT_MOC_LITERAL(6, 57, 16), // "transportChanged"
QT_MOC_LITERAL(7, 74, 9), // "transport"
QT_MOC_LITERAL(8, 84, 15), // "endpointChanged"
QT_MOC_LITERAL(9, 100, 8), // "endpoint"
QT_MOC_LITERAL(10, 109, 11), // "fromChanged"
QT_MOC_LITERAL(11, 121, 4), // "from"
QT_MOC_LITERAL(12, 126, 14), // "messageChanged"
QT_MOC_LITERAL(13, 141, 7), // "message"
QT_MOC_LITERAL(14, 149, 14), // "rawDataChanged"
QT_MOC_LITERAL(15, 164, 7), // "rawData"
QT_MOC_LITERAL(16, 172, 11), // "hashChanged"
QT_MOC_LITERAL(17, 184, 4) // "hash"

    },
    "ApproveSignDataCtx721036\0clicked\0\0b\0"
    "remoteChanged\0remote\0transportChanged\0"
    "transport\0endpointChanged\0endpoint\0"
    "fromChanged\0from\0messageChanged\0message\0"
    "rawDataChanged\0rawData\0hashChanged\0"
    "hash"
};
#undef QT_MOC_LITERAL

static const uint qt_meta_data_ApproveSignDataCtx721036[] = {

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

void ApproveSignDataCtx721036::qt_static_metacall(QObject *_o, QMetaObject::Call _c, int _id, void **_a)
{
    if (_c == QMetaObject::InvokeMetaMethod) {
        ApproveSignDataCtx721036 *_t = static_cast<ApproveSignDataCtx721036 *>(_o);
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
            using _t = void (ApproveSignDataCtx721036::*)(qint32 );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveSignDataCtx721036::clicked)) {
                *result = 0;
                return;
            }
        }
        {
            using _t = void (ApproveSignDataCtx721036::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveSignDataCtx721036::remoteChanged)) {
                *result = 1;
                return;
            }
        }
        {
            using _t = void (ApproveSignDataCtx721036::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveSignDataCtx721036::transportChanged)) {
                *result = 2;
                return;
            }
        }
        {
            using _t = void (ApproveSignDataCtx721036::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveSignDataCtx721036::endpointChanged)) {
                *result = 3;
                return;
            }
        }
        {
            using _t = void (ApproveSignDataCtx721036::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveSignDataCtx721036::fromChanged)) {
                *result = 4;
                return;
            }
        }
        {
            using _t = void (ApproveSignDataCtx721036::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveSignDataCtx721036::messageChanged)) {
                *result = 5;
                return;
            }
        }
        {
            using _t = void (ApproveSignDataCtx721036::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveSignDataCtx721036::rawDataChanged)) {
                *result = 6;
                return;
            }
        }
        {
            using _t = void (ApproveSignDataCtx721036::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveSignDataCtx721036::hashChanged)) {
                *result = 7;
                return;
            }
        }
    }
#ifndef QT_NO_PROPERTIES
    else if (_c == QMetaObject::ReadProperty) {
        ApproveSignDataCtx721036 *_t = static_cast<ApproveSignDataCtx721036 *>(_o);
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
        ApproveSignDataCtx721036 *_t = static_cast<ApproveSignDataCtx721036 *>(_o);
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

QT_INIT_METAOBJECT const QMetaObject ApproveSignDataCtx721036::staticMetaObject = {
    { &QObject::staticMetaObject, qt_meta_stringdata_ApproveSignDataCtx721036.data,
      qt_meta_data_ApproveSignDataCtx721036,  qt_static_metacall, nullptr, nullptr}
};


const QMetaObject *ApproveSignDataCtx721036::metaObject() const
{
    return QObject::d_ptr->metaObject ? QObject::d_ptr->dynamicMetaObject() : &staticMetaObject;
}

void *ApproveSignDataCtx721036::qt_metacast(const char *_clname)
{
    if (!_clname) return nullptr;
    if (!strcmp(_clname, qt_meta_stringdata_ApproveSignDataCtx721036.stringdata0))
        return static_cast<void*>(this);
    return QObject::qt_metacast(_clname);
}

int ApproveSignDataCtx721036::qt_metacall(QMetaObject::Call _c, int _id, void **_a)
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
void ApproveSignDataCtx721036::clicked(qint32 _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 0, _a);
}

// SIGNAL 1
void ApproveSignDataCtx721036::remoteChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 1, _a);
}

// SIGNAL 2
void ApproveSignDataCtx721036::transportChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 2, _a);
}

// SIGNAL 3
void ApproveSignDataCtx721036::endpointChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 3, _a);
}

// SIGNAL 4
void ApproveSignDataCtx721036::fromChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 4, _a);
}

// SIGNAL 5
void ApproveSignDataCtx721036::messageChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 5, _a);
}

// SIGNAL 6
void ApproveSignDataCtx721036::rawDataChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 6, _a);
}

// SIGNAL 7
void ApproveSignDataCtx721036::hashChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 7, _a);
}
struct qt_meta_stringdata_ApproveTxCtx721036_t {
    QByteArrayData data[27];
    char stringdata0[239];
};
#define QT_MOC_LITERAL(idx, ofs, len) \
    Q_STATIC_BYTE_ARRAY_DATA_HEADER_INITIALIZER_WITH_OFFSET(len, \
    qptrdiff(offsetof(qt_meta_stringdata_ApproveTxCtx721036_t, stringdata0) + ofs \
        - idx * sizeof(QByteArrayData)) \
    )
static const qt_meta_stringdata_ApproveTxCtx721036_t qt_meta_stringdata_ApproveTxCtx721036 = {
    {
QT_MOC_LITERAL(0, 0, 18), // "ApproveTxCtx721036"
QT_MOC_LITERAL(1, 19, 7), // "clicked"
QT_MOC_LITERAL(2, 27, 0), // ""
QT_MOC_LITERAL(3, 28, 1), // "b"
QT_MOC_LITERAL(4, 30, 6), // "edited"
QT_MOC_LITERAL(5, 37, 1), // "s"
QT_MOC_LITERAL(6, 39, 1), // "v"
QT_MOC_LITERAL(7, 41, 13), // "remoteChanged"
QT_MOC_LITERAL(8, 55, 6), // "remote"
QT_MOC_LITERAL(9, 62, 16), // "transportChanged"
QT_MOC_LITERAL(10, 79, 9), // "transport"
QT_MOC_LITERAL(11, 89, 15), // "endpointChanged"
QT_MOC_LITERAL(12, 105, 8), // "endpoint"
QT_MOC_LITERAL(13, 114, 11), // "dataChanged"
QT_MOC_LITERAL(14, 126, 4), // "data"
QT_MOC_LITERAL(15, 131, 11), // "fromChanged"
QT_MOC_LITERAL(16, 143, 4), // "from"
QT_MOC_LITERAL(17, 148, 9), // "toChanged"
QT_MOC_LITERAL(18, 158, 2), // "to"
QT_MOC_LITERAL(19, 161, 10), // "gasChanged"
QT_MOC_LITERAL(20, 172, 3), // "gas"
QT_MOC_LITERAL(21, 176, 15), // "gasPriceChanged"
QT_MOC_LITERAL(22, 192, 8), // "gasPrice"
QT_MOC_LITERAL(23, 201, 12), // "nonceChanged"
QT_MOC_LITERAL(24, 214, 5), // "nonce"
QT_MOC_LITERAL(25, 220, 12), // "valueChanged"
QT_MOC_LITERAL(26, 233, 5) // "value"

    },
    "ApproveTxCtx721036\0clicked\0\0b\0edited\0"
    "s\0v\0remoteChanged\0remote\0transportChanged\0"
    "transport\0endpointChanged\0endpoint\0"
    "dataChanged\0data\0fromChanged\0from\0"
    "toChanged\0to\0gasChanged\0gas\0gasPriceChanged\0"
    "gasPrice\0nonceChanged\0nonce\0valueChanged\0"
    "value"
};
#undef QT_MOC_LITERAL

static const uint qt_meta_data_ApproveTxCtx721036[] = {

 // content:
       7,       // revision
       0,       // classname
       0,    0, // classinfo
      12,   14, // methods
      10,  112, // properties
       0,    0, // enums/sets
       0,    0, // constructors
       0,       // flags
      12,       // signalCount

 // signals: name, argc, parameters, tag, flags
       1,    1,   74,    2, 0x06 /* Public */,
       4,    2,   77,    2, 0x06 /* Public */,
       7,    1,   82,    2, 0x06 /* Public */,
       9,    1,   85,    2, 0x06 /* Public */,
      11,    1,   88,    2, 0x06 /* Public */,
      13,    1,   91,    2, 0x06 /* Public */,
      15,    1,   94,    2, 0x06 /* Public */,
      17,    1,   97,    2, 0x06 /* Public */,
      19,    1,  100,    2, 0x06 /* Public */,
      21,    1,  103,    2, 0x06 /* Public */,
      23,    1,  106,    2, 0x06 /* Public */,
      25,    1,  109,    2, 0x06 /* Public */,

 // signals: parameters
    QMetaType::Void, QMetaType::Int,    3,
    QMetaType::Void, QMetaType::QString, QMetaType::QString,    5,    6,
    QMetaType::Void, QMetaType::QString,    8,
    QMetaType::Void, QMetaType::QString,   10,
    QMetaType::Void, QMetaType::QString,   12,
    QMetaType::Void, QMetaType::QString,   14,
    QMetaType::Void, QMetaType::QString,   16,
    QMetaType::Void, QMetaType::QString,   18,
    QMetaType::Void, QMetaType::QString,   20,
    QMetaType::Void, QMetaType::QString,   22,
    QMetaType::Void, QMetaType::QString,   24,
    QMetaType::Void, QMetaType::QString,   26,

 // properties: name, type, flags
       8, QMetaType::QString, 0x00495103,
      10, QMetaType::QString, 0x00495103,
      12, QMetaType::QString, 0x00495103,
      14, QMetaType::QString, 0x00495103,
      16, QMetaType::QString, 0x00495103,
      18, QMetaType::QString, 0x00495103,
      20, QMetaType::QString, 0x00495103,
      22, QMetaType::QString, 0x00495103,
      24, QMetaType::QString, 0x00495103,
      26, QMetaType::QString, 0x00495103,

 // properties: notify_signal_id
       2,
       3,
       4,
       5,
       6,
       7,
       8,
       9,
      10,
      11,

       0        // eod
};

void ApproveTxCtx721036::qt_static_metacall(QObject *_o, QMetaObject::Call _c, int _id, void **_a)
{
    if (_c == QMetaObject::InvokeMetaMethod) {
        ApproveTxCtx721036 *_t = static_cast<ApproveTxCtx721036 *>(_o);
        Q_UNUSED(_t)
        switch (_id) {
        case 0: _t->clicked((*reinterpret_cast< qint32(*)>(_a[1]))); break;
        case 1: _t->edited((*reinterpret_cast< QString(*)>(_a[1])),(*reinterpret_cast< QString(*)>(_a[2]))); break;
        case 2: _t->remoteChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 3: _t->transportChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 4: _t->endpointChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 5: _t->dataChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 6: _t->fromChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 7: _t->toChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 8: _t->gasChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 9: _t->gasPriceChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 10: _t->nonceChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 11: _t->valueChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        default: ;
        }
    } else if (_c == QMetaObject::IndexOfMethod) {
        int *result = reinterpret_cast<int *>(_a[0]);
        {
            using _t = void (ApproveTxCtx721036::*)(qint32 );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveTxCtx721036::clicked)) {
                *result = 0;
                return;
            }
        }
        {
            using _t = void (ApproveTxCtx721036::*)(QString , QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveTxCtx721036::edited)) {
                *result = 1;
                return;
            }
        }
        {
            using _t = void (ApproveTxCtx721036::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveTxCtx721036::remoteChanged)) {
                *result = 2;
                return;
            }
        }
        {
            using _t = void (ApproveTxCtx721036::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveTxCtx721036::transportChanged)) {
                *result = 3;
                return;
            }
        }
        {
            using _t = void (ApproveTxCtx721036::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveTxCtx721036::endpointChanged)) {
                *result = 4;
                return;
            }
        }
        {
            using _t = void (ApproveTxCtx721036::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveTxCtx721036::dataChanged)) {
                *result = 5;
                return;
            }
        }
        {
            using _t = void (ApproveTxCtx721036::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveTxCtx721036::fromChanged)) {
                *result = 6;
                return;
            }
        }
        {
            using _t = void (ApproveTxCtx721036::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveTxCtx721036::toChanged)) {
                *result = 7;
                return;
            }
        }
        {
            using _t = void (ApproveTxCtx721036::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveTxCtx721036::gasChanged)) {
                *result = 8;
                return;
            }
        }
        {
            using _t = void (ApproveTxCtx721036::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveTxCtx721036::gasPriceChanged)) {
                *result = 9;
                return;
            }
        }
        {
            using _t = void (ApproveTxCtx721036::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveTxCtx721036::nonceChanged)) {
                *result = 10;
                return;
            }
        }
        {
            using _t = void (ApproveTxCtx721036::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveTxCtx721036::valueChanged)) {
                *result = 11;
                return;
            }
        }
    }
#ifndef QT_NO_PROPERTIES
    else if (_c == QMetaObject::ReadProperty) {
        ApproveTxCtx721036 *_t = static_cast<ApproveTxCtx721036 *>(_o);
        Q_UNUSED(_t)
        void *_v = _a[0];
        switch (_id) {
        case 0: *reinterpret_cast< QString*>(_v) = _t->remote(); break;
        case 1: *reinterpret_cast< QString*>(_v) = _t->transport(); break;
        case 2: *reinterpret_cast< QString*>(_v) = _t->endpoint(); break;
        case 3: *reinterpret_cast< QString*>(_v) = _t->data(); break;
        case 4: *reinterpret_cast< QString*>(_v) = _t->from(); break;
        case 5: *reinterpret_cast< QString*>(_v) = _t->to(); break;
        case 6: *reinterpret_cast< QString*>(_v) = _t->gas(); break;
        case 7: *reinterpret_cast< QString*>(_v) = _t->gasPrice(); break;
        case 8: *reinterpret_cast< QString*>(_v) = _t->nonce(); break;
        case 9: *reinterpret_cast< QString*>(_v) = _t->value(); break;
        default: break;
        }
    } else if (_c == QMetaObject::WriteProperty) {
        ApproveTxCtx721036 *_t = static_cast<ApproveTxCtx721036 *>(_o);
        Q_UNUSED(_t)
        void *_v = _a[0];
        switch (_id) {
        case 0: _t->setRemote(*reinterpret_cast< QString*>(_v)); break;
        case 1: _t->setTransport(*reinterpret_cast< QString*>(_v)); break;
        case 2: _t->setEndpoint(*reinterpret_cast< QString*>(_v)); break;
        case 3: _t->setData(*reinterpret_cast< QString*>(_v)); break;
        case 4: _t->setFrom(*reinterpret_cast< QString*>(_v)); break;
        case 5: _t->setTo(*reinterpret_cast< QString*>(_v)); break;
        case 6: _t->setGas(*reinterpret_cast< QString*>(_v)); break;
        case 7: _t->setGasPrice(*reinterpret_cast< QString*>(_v)); break;
        case 8: _t->setNonce(*reinterpret_cast< QString*>(_v)); break;
        case 9: _t->setValue(*reinterpret_cast< QString*>(_v)); break;
        default: break;
        }
    } else if (_c == QMetaObject::ResetProperty) {
    }
#endif // QT_NO_PROPERTIES
}

QT_INIT_METAOBJECT const QMetaObject ApproveTxCtx721036::staticMetaObject = {
    { &QObject::staticMetaObject, qt_meta_stringdata_ApproveTxCtx721036.data,
      qt_meta_data_ApproveTxCtx721036,  qt_static_metacall, nullptr, nullptr}
};


const QMetaObject *ApproveTxCtx721036::metaObject() const
{
    return QObject::d_ptr->metaObject ? QObject::d_ptr->dynamicMetaObject() : &staticMetaObject;
}

void *ApproveTxCtx721036::qt_metacast(const char *_clname)
{
    if (!_clname) return nullptr;
    if (!strcmp(_clname, qt_meta_stringdata_ApproveTxCtx721036.stringdata0))
        return static_cast<void*>(this);
    return QObject::qt_metacast(_clname);
}

int ApproveTxCtx721036::qt_metacall(QMetaObject::Call _c, int _id, void **_a)
{
    _id = QObject::qt_metacall(_c, _id, _a);
    if (_id < 0)
        return _id;
    if (_c == QMetaObject::InvokeMetaMethod) {
        if (_id < 12)
            qt_static_metacall(this, _c, _id, _a);
        _id -= 12;
    } else if (_c == QMetaObject::RegisterMethodArgumentMetaType) {
        if (_id < 12)
            *reinterpret_cast<int*>(_a[0]) = -1;
        _id -= 12;
    }
#ifndef QT_NO_PROPERTIES
   else if (_c == QMetaObject::ReadProperty || _c == QMetaObject::WriteProperty
            || _c == QMetaObject::ResetProperty || _c == QMetaObject::RegisterPropertyMetaType) {
        qt_static_metacall(this, _c, _id, _a);
        _id -= 10;
    } else if (_c == QMetaObject::QueryPropertyDesignable) {
        _id -= 10;
    } else if (_c == QMetaObject::QueryPropertyScriptable) {
        _id -= 10;
    } else if (_c == QMetaObject::QueryPropertyStored) {
        _id -= 10;
    } else if (_c == QMetaObject::QueryPropertyEditable) {
        _id -= 10;
    } else if (_c == QMetaObject::QueryPropertyUser) {
        _id -= 10;
    }
#endif // QT_NO_PROPERTIES
    return _id;
}

// SIGNAL 0
void ApproveTxCtx721036::clicked(qint32 _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 0, _a);
}

// SIGNAL 1
void ApproveTxCtx721036::edited(QString _t1, QString _t2)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)), const_cast<void*>(reinterpret_cast<const void*>(&_t2)) };
    QMetaObject::activate(this, &staticMetaObject, 1, _a);
}

// SIGNAL 2
void ApproveTxCtx721036::remoteChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 2, _a);
}

// SIGNAL 3
void ApproveTxCtx721036::transportChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 3, _a);
}

// SIGNAL 4
void ApproveTxCtx721036::endpointChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 4, _a);
}

// SIGNAL 5
void ApproveTxCtx721036::dataChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 5, _a);
}

// SIGNAL 6
void ApproveTxCtx721036::fromChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 6, _a);
}

// SIGNAL 7
void ApproveTxCtx721036::toChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 7, _a);
}

// SIGNAL 8
void ApproveTxCtx721036::gasChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 8, _a);
}

// SIGNAL 9
void ApproveTxCtx721036::gasPriceChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 9, _a);
}

// SIGNAL 10
void ApproveTxCtx721036::nonceChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 10, _a);
}

// SIGNAL 11
void ApproveTxCtx721036::valueChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 11, _a);
}
struct qt_meta_stringdata_CustomListModel721036_t {
    QByteArrayData data[11];
    char stringdata0[90];
};
#define QT_MOC_LITERAL(idx, ofs, len) \
    Q_STATIC_BYTE_ARRAY_DATA_HEADER_INITIALIZER_WITH_OFFSET(len, \
    qptrdiff(offsetof(qt_meta_stringdata_CustomListModel721036_t, stringdata0) + ofs \
        - idx * sizeof(QByteArrayData)) \
    )
static const qt_meta_stringdata_CustomListModel721036_t qt_meta_stringdata_CustomListModel721036 = {
    {
QT_MOC_LITERAL(0, 0, 21), // "CustomListModel721036"
QT_MOC_LITERAL(1, 22, 5), // "clear"
QT_MOC_LITERAL(2, 28, 0), // ""
QT_MOC_LITERAL(3, 29, 3), // "add"
QT_MOC_LITERAL(4, 33, 8), // "quintptr"
QT_MOC_LITERAL(5, 42, 7), // "account"
QT_MOC_LITERAL(6, 50, 7), // "clicked"
QT_MOC_LITERAL(7, 58, 1), // "b"
QT_MOC_LITERAL(8, 60, 19), // "onCheckStateChanged"
QT_MOC_LITERAL(9, 80, 1), // "i"
QT_MOC_LITERAL(10, 82, 7) // "checked"

    },
    "CustomListModel721036\0clear\0\0add\0"
    "quintptr\0account\0clicked\0b\0"
    "onCheckStateChanged\0i\0checked"
};
#undef QT_MOC_LITERAL

static const uint qt_meta_data_CustomListModel721036[] = {

 // content:
       7,       // revision
       0,       // classname
       0,    0, // classinfo
       4,   14, // methods
       0,    0, // properties
       0,    0, // enums/sets
       0,    0, // constructors
       0,       // flags
       4,       // signalCount

 // signals: name, argc, parameters, tag, flags
       1,    0,   34,    2, 0x06 /* Public */,
       3,    1,   35,    2, 0x06 /* Public */,
       6,    1,   38,    2, 0x06 /* Public */,
       8,    2,   41,    2, 0x06 /* Public */,

 // signals: parameters
    QMetaType::Void,
    QMetaType::Void, 0x80000000 | 4,    5,
    QMetaType::Void, QMetaType::Int,    7,
    QMetaType::Void, QMetaType::Int, QMetaType::Bool,    9,   10,

       0        // eod
};

void CustomListModel721036::qt_static_metacall(QObject *_o, QMetaObject::Call _c, int _id, void **_a)
{
    if (_c == QMetaObject::InvokeMetaMethod) {
        CustomListModel721036 *_t = static_cast<CustomListModel721036 *>(_o);
        Q_UNUSED(_t)
        switch (_id) {
        case 0: _t->clear(); break;
        case 1: _t->add((*reinterpret_cast< quintptr(*)>(_a[1]))); break;
        case 2: _t->clicked((*reinterpret_cast< qint32(*)>(_a[1]))); break;
        case 3: _t->onCheckStateChanged((*reinterpret_cast< qint32(*)>(_a[1])),(*reinterpret_cast< bool(*)>(_a[2]))); break;
        default: ;
        }
    } else if (_c == QMetaObject::IndexOfMethod) {
        int *result = reinterpret_cast<int *>(_a[0]);
        {
            using _t = void (CustomListModel721036::*)();
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&CustomListModel721036::clear)) {
                *result = 0;
                return;
            }
        }
        {
            using _t = void (CustomListModel721036::*)(quintptr );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&CustomListModel721036::add)) {
                *result = 1;
                return;
            }
        }
        {
            using _t = void (CustomListModel721036::*)(qint32 );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&CustomListModel721036::clicked)) {
                *result = 2;
                return;
            }
        }
        {
            using _t = void (CustomListModel721036::*)(qint32 , bool );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&CustomListModel721036::onCheckStateChanged)) {
                *result = 3;
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
        if (_id < 4)
            qt_static_metacall(this, _c, _id, _a);
        _id -= 4;
    } else if (_c == QMetaObject::RegisterMethodArgumentMetaType) {
        if (_id < 4)
            *reinterpret_cast<int*>(_a[0]) = -1;
        _id -= 4;
    }
    return _id;
}

// SIGNAL 0
void CustomListModel721036::clear()
{
    QMetaObject::activate(this, &staticMetaObject, 0, nullptr);
}

// SIGNAL 1
void CustomListModel721036::add(quintptr _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 1, _a);
}

// SIGNAL 2
void CustomListModel721036::clicked(qint32 _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 2, _a);
}

// SIGNAL 3
void CustomListModel721036::onCheckStateChanged(qint32 _t1, bool _t2)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)), const_cast<void*>(reinterpret_cast<const void*>(&_t2)) };
    QMetaObject::activate(this, &staticMetaObject, 3, _a);
}
struct qt_meta_stringdata_ApproveImportCtx721036_t {
    QByteArrayData data[16];
    char stringdata0[208];
};
#define QT_MOC_LITERAL(idx, ofs, len) \
    Q_STATIC_BYTE_ARRAY_DATA_HEADER_INITIALIZER_WITH_OFFSET(len, \
    qptrdiff(offsetof(qt_meta_stringdata_ApproveImportCtx721036_t, stringdata0) + ofs \
        - idx * sizeof(QByteArrayData)) \
    )
static const qt_meta_stringdata_ApproveImportCtx721036_t qt_meta_stringdata_ApproveImportCtx721036 = {
    {
QT_MOC_LITERAL(0, 0, 22), // "ApproveImportCtx721036"
QT_MOC_LITERAL(1, 23, 7), // "clicked"
QT_MOC_LITERAL(2, 31, 0), // ""
QT_MOC_LITERAL(3, 32, 1), // "b"
QT_MOC_LITERAL(4, 34, 14), // "passwordEdited"
QT_MOC_LITERAL(5, 49, 21), // "confirmPasswordEdited"
QT_MOC_LITERAL(6, 71, 13), // "remoteChanged"
QT_MOC_LITERAL(7, 85, 6), // "remote"
QT_MOC_LITERAL(8, 92, 16), // "transportChanged"
QT_MOC_LITERAL(9, 109, 9), // "transport"
QT_MOC_LITERAL(10, 119, 15), // "endpointChanged"
QT_MOC_LITERAL(11, 135, 8), // "endpoint"
QT_MOC_LITERAL(12, 144, 15), // "passwordChanged"
QT_MOC_LITERAL(13, 160, 8), // "password"
QT_MOC_LITERAL(14, 169, 22), // "confirmPasswordChanged"
QT_MOC_LITERAL(15, 192, 15) // "confirmPassword"

    },
    "ApproveImportCtx721036\0clicked\0\0b\0"
    "passwordEdited\0confirmPasswordEdited\0"
    "remoteChanged\0remote\0transportChanged\0"
    "transport\0endpointChanged\0endpoint\0"
    "passwordChanged\0password\0"
    "confirmPasswordChanged\0confirmPassword"
};
#undef QT_MOC_LITERAL

static const uint qt_meta_data_ApproveImportCtx721036[] = {

 // content:
       7,       // revision
       0,       // classname
       0,    0, // classinfo
       8,   14, // methods
       5,   78, // properties
       0,    0, // enums/sets
       0,    0, // constructors
       0,       // flags
       8,       // signalCount

 // signals: name, argc, parameters, tag, flags
       1,    1,   54,    2, 0x06 /* Public */,
       4,    1,   57,    2, 0x06 /* Public */,
       5,    1,   60,    2, 0x06 /* Public */,
       6,    1,   63,    2, 0x06 /* Public */,
       8,    1,   66,    2, 0x06 /* Public */,
      10,    1,   69,    2, 0x06 /* Public */,
      12,    1,   72,    2, 0x06 /* Public */,
      14,    1,   75,    2, 0x06 /* Public */,

 // signals: parameters
    QMetaType::Void, QMetaType::Int,    3,
    QMetaType::Void, QMetaType::QString,    3,
    QMetaType::Void, QMetaType::QString,    3,
    QMetaType::Void, QMetaType::QString,    7,
    QMetaType::Void, QMetaType::QString,    9,
    QMetaType::Void, QMetaType::QString,   11,
    QMetaType::Void, QMetaType::QString,   13,
    QMetaType::Void, QMetaType::QString,   15,

 // properties: name, type, flags
       7, QMetaType::QString, 0x00495103,
       9, QMetaType::QString, 0x00495103,
      11, QMetaType::QString, 0x00495103,
      13, QMetaType::QString, 0x00495103,
      15, QMetaType::QString, 0x00495103,

 // properties: notify_signal_id
       3,
       4,
       5,
       6,
       7,

       0        // eod
};

void ApproveImportCtx721036::qt_static_metacall(QObject *_o, QMetaObject::Call _c, int _id, void **_a)
{
    if (_c == QMetaObject::InvokeMetaMethod) {
        ApproveImportCtx721036 *_t = static_cast<ApproveImportCtx721036 *>(_o);
        Q_UNUSED(_t)
        switch (_id) {
        case 0: _t->clicked((*reinterpret_cast< qint32(*)>(_a[1]))); break;
        case 1: _t->passwordEdited((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 2: _t->confirmPasswordEdited((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 3: _t->remoteChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 4: _t->transportChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 5: _t->endpointChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 6: _t->passwordChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 7: _t->confirmPasswordChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        default: ;
        }
    } else if (_c == QMetaObject::IndexOfMethod) {
        int *result = reinterpret_cast<int *>(_a[0]);
        {
            using _t = void (ApproveImportCtx721036::*)(qint32 );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveImportCtx721036::clicked)) {
                *result = 0;
                return;
            }
        }
        {
            using _t = void (ApproveImportCtx721036::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveImportCtx721036::passwordEdited)) {
                *result = 1;
                return;
            }
        }
        {
            using _t = void (ApproveImportCtx721036::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveImportCtx721036::confirmPasswordEdited)) {
                *result = 2;
                return;
            }
        }
        {
            using _t = void (ApproveImportCtx721036::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveImportCtx721036::remoteChanged)) {
                *result = 3;
                return;
            }
        }
        {
            using _t = void (ApproveImportCtx721036::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveImportCtx721036::transportChanged)) {
                *result = 4;
                return;
            }
        }
        {
            using _t = void (ApproveImportCtx721036::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveImportCtx721036::endpointChanged)) {
                *result = 5;
                return;
            }
        }
        {
            using _t = void (ApproveImportCtx721036::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveImportCtx721036::passwordChanged)) {
                *result = 6;
                return;
            }
        }
        {
            using _t = void (ApproveImportCtx721036::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveImportCtx721036::confirmPasswordChanged)) {
                *result = 7;
                return;
            }
        }
    }
#ifndef QT_NO_PROPERTIES
    else if (_c == QMetaObject::ReadProperty) {
        ApproveImportCtx721036 *_t = static_cast<ApproveImportCtx721036 *>(_o);
        Q_UNUSED(_t)
        void *_v = _a[0];
        switch (_id) {
        case 0: *reinterpret_cast< QString*>(_v) = _t->remote(); break;
        case 1: *reinterpret_cast< QString*>(_v) = _t->transport(); break;
        case 2: *reinterpret_cast< QString*>(_v) = _t->endpoint(); break;
        case 3: *reinterpret_cast< QString*>(_v) = _t->password(); break;
        case 4: *reinterpret_cast< QString*>(_v) = _t->confirmPassword(); break;
        default: break;
        }
    } else if (_c == QMetaObject::WriteProperty) {
        ApproveImportCtx721036 *_t = static_cast<ApproveImportCtx721036 *>(_o);
        Q_UNUSED(_t)
        void *_v = _a[0];
        switch (_id) {
        case 0: _t->setRemote(*reinterpret_cast< QString*>(_v)); break;
        case 1: _t->setTransport(*reinterpret_cast< QString*>(_v)); break;
        case 2: _t->setEndpoint(*reinterpret_cast< QString*>(_v)); break;
        case 3: _t->setPassword(*reinterpret_cast< QString*>(_v)); break;
        case 4: _t->setConfirmPassword(*reinterpret_cast< QString*>(_v)); break;
        default: break;
        }
    } else if (_c == QMetaObject::ResetProperty) {
    }
#endif // QT_NO_PROPERTIES
}

QT_INIT_METAOBJECT const QMetaObject ApproveImportCtx721036::staticMetaObject = {
    { &QObject::staticMetaObject, qt_meta_stringdata_ApproveImportCtx721036.data,
      qt_meta_data_ApproveImportCtx721036,  qt_static_metacall, nullptr, nullptr}
};


const QMetaObject *ApproveImportCtx721036::metaObject() const
{
    return QObject::d_ptr->metaObject ? QObject::d_ptr->dynamicMetaObject() : &staticMetaObject;
}

void *ApproveImportCtx721036::qt_metacast(const char *_clname)
{
    if (!_clname) return nullptr;
    if (!strcmp(_clname, qt_meta_stringdata_ApproveImportCtx721036.stringdata0))
        return static_cast<void*>(this);
    return QObject::qt_metacast(_clname);
}

int ApproveImportCtx721036::qt_metacall(QMetaObject::Call _c, int _id, void **_a)
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
        _id -= 5;
    } else if (_c == QMetaObject::QueryPropertyDesignable) {
        _id -= 5;
    } else if (_c == QMetaObject::QueryPropertyScriptable) {
        _id -= 5;
    } else if (_c == QMetaObject::QueryPropertyStored) {
        _id -= 5;
    } else if (_c == QMetaObject::QueryPropertyEditable) {
        _id -= 5;
    } else if (_c == QMetaObject::QueryPropertyUser) {
        _id -= 5;
    }
#endif // QT_NO_PROPERTIES
    return _id;
}

// SIGNAL 0
void ApproveImportCtx721036::clicked(qint32 _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 0, _a);
}

// SIGNAL 1
void ApproveImportCtx721036::passwordEdited(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 1, _a);
}

// SIGNAL 2
void ApproveImportCtx721036::confirmPasswordEdited(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 2, _a);
}

// SIGNAL 3
void ApproveImportCtx721036::remoteChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 3, _a);
}

// SIGNAL 4
void ApproveImportCtx721036::transportChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 4, _a);
}

// SIGNAL 5
void ApproveImportCtx721036::endpointChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 5, _a);
}

// SIGNAL 6
void ApproveImportCtx721036::passwordChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 6, _a);
}

// SIGNAL 7
void ApproveImportCtx721036::confirmPasswordChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 7, _a);
}
struct qt_meta_stringdata_ApproveListingCtx721036_t {
    QByteArrayData data[16];
    char stringdata0[178];
};
#define QT_MOC_LITERAL(idx, ofs, len) \
    Q_STATIC_BYTE_ARRAY_DATA_HEADER_INITIALIZER_WITH_OFFSET(len, \
    qptrdiff(offsetof(qt_meta_stringdata_ApproveListingCtx721036_t, stringdata0) + ofs \
        - idx * sizeof(QByteArrayData)) \
    )
static const qt_meta_stringdata_ApproveListingCtx721036_t qt_meta_stringdata_ApproveListingCtx721036 = {
    {
QT_MOC_LITERAL(0, 0, 23), // "ApproveListingCtx721036"
QT_MOC_LITERAL(1, 24, 13), // "remoteChanged"
QT_MOC_LITERAL(2, 38, 0), // ""
QT_MOC_LITERAL(3, 39, 6), // "remote"
QT_MOC_LITERAL(4, 46, 16), // "transportChanged"
QT_MOC_LITERAL(5, 63, 9), // "transport"
QT_MOC_LITERAL(6, 73, 15), // "endpointChanged"
QT_MOC_LITERAL(7, 89, 8), // "endpoint"
QT_MOC_LITERAL(8, 98, 11), // "fromChanged"
QT_MOC_LITERAL(9, 110, 4), // "from"
QT_MOC_LITERAL(10, 115, 14), // "messageChanged"
QT_MOC_LITERAL(11, 130, 7), // "message"
QT_MOC_LITERAL(12, 138, 14), // "rawDataChanged"
QT_MOC_LITERAL(13, 153, 7), // "rawData"
QT_MOC_LITERAL(14, 161, 11), // "hashChanged"
QT_MOC_LITERAL(15, 173, 4) // "hash"

    },
    "ApproveListingCtx721036\0remoteChanged\0"
    "\0remote\0transportChanged\0transport\0"
    "endpointChanged\0endpoint\0fromChanged\0"
    "from\0messageChanged\0message\0rawDataChanged\0"
    "rawData\0hashChanged\0hash"
};
#undef QT_MOC_LITERAL

static const uint qt_meta_data_ApproveListingCtx721036[] = {

 // content:
       7,       // revision
       0,       // classname
       0,    0, // classinfo
       7,   14, // methods
       7,   70, // properties
       0,    0, // enums/sets
       0,    0, // constructors
       0,       // flags
       7,       // signalCount

 // signals: name, argc, parameters, tag, flags
       1,    1,   49,    2, 0x06 /* Public */,
       4,    1,   52,    2, 0x06 /* Public */,
       6,    1,   55,    2, 0x06 /* Public */,
       8,    1,   58,    2, 0x06 /* Public */,
      10,    1,   61,    2, 0x06 /* Public */,
      12,    1,   64,    2, 0x06 /* Public */,
      14,    1,   67,    2, 0x06 /* Public */,

 // signals: parameters
    QMetaType::Void, QMetaType::QString,    3,
    QMetaType::Void, QMetaType::QString,    5,
    QMetaType::Void, QMetaType::QString,    7,
    QMetaType::Void, QMetaType::QString,    9,
    QMetaType::Void, QMetaType::QString,   11,
    QMetaType::Void, QMetaType::QString,   13,
    QMetaType::Void, QMetaType::QString,   15,

 // properties: name, type, flags
       3, QMetaType::QString, 0x00495103,
       5, QMetaType::QString, 0x00495103,
       7, QMetaType::QString, 0x00495103,
       9, QMetaType::QString, 0x00495103,
      11, QMetaType::QString, 0x00495103,
      13, QMetaType::QString, 0x00495103,
      15, QMetaType::QString, 0x00495103,

 // properties: notify_signal_id
       0,
       1,
       2,
       3,
       4,
       5,
       6,

       0        // eod
};

void ApproveListingCtx721036::qt_static_metacall(QObject *_o, QMetaObject::Call _c, int _id, void **_a)
{
    if (_c == QMetaObject::InvokeMetaMethod) {
        ApproveListingCtx721036 *_t = static_cast<ApproveListingCtx721036 *>(_o);
        Q_UNUSED(_t)
        switch (_id) {
        case 0: _t->remoteChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 1: _t->transportChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 2: _t->endpointChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 3: _t->fromChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 4: _t->messageChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 5: _t->rawDataChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 6: _t->hashChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        default: ;
        }
    } else if (_c == QMetaObject::IndexOfMethod) {
        int *result = reinterpret_cast<int *>(_a[0]);
        {
            using _t = void (ApproveListingCtx721036::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveListingCtx721036::remoteChanged)) {
                *result = 0;
                return;
            }
        }
        {
            using _t = void (ApproveListingCtx721036::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveListingCtx721036::transportChanged)) {
                *result = 1;
                return;
            }
        }
        {
            using _t = void (ApproveListingCtx721036::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveListingCtx721036::endpointChanged)) {
                *result = 2;
                return;
            }
        }
        {
            using _t = void (ApproveListingCtx721036::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveListingCtx721036::fromChanged)) {
                *result = 3;
                return;
            }
        }
        {
            using _t = void (ApproveListingCtx721036::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveListingCtx721036::messageChanged)) {
                *result = 4;
                return;
            }
        }
        {
            using _t = void (ApproveListingCtx721036::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveListingCtx721036::rawDataChanged)) {
                *result = 5;
                return;
            }
        }
        {
            using _t = void (ApproveListingCtx721036::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveListingCtx721036::hashChanged)) {
                *result = 6;
                return;
            }
        }
    }
#ifndef QT_NO_PROPERTIES
    else if (_c == QMetaObject::ReadProperty) {
        ApproveListingCtx721036 *_t = static_cast<ApproveListingCtx721036 *>(_o);
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
        ApproveListingCtx721036 *_t = static_cast<ApproveListingCtx721036 *>(_o);
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

QT_INIT_METAOBJECT const QMetaObject ApproveListingCtx721036::staticMetaObject = {
    { &QObject::staticMetaObject, qt_meta_stringdata_ApproveListingCtx721036.data,
      qt_meta_data_ApproveListingCtx721036,  qt_static_metacall, nullptr, nullptr}
};


const QMetaObject *ApproveListingCtx721036::metaObject() const
{
    return QObject::d_ptr->metaObject ? QObject::d_ptr->dynamicMetaObject() : &staticMetaObject;
}

void *ApproveListingCtx721036::qt_metacast(const char *_clname)
{
    if (!_clname) return nullptr;
    if (!strcmp(_clname, qt_meta_stringdata_ApproveListingCtx721036.stringdata0))
        return static_cast<void*>(this);
    return QObject::qt_metacast(_clname);
}

int ApproveListingCtx721036::qt_metacall(QMetaObject::Call _c, int _id, void **_a)
{
    _id = QObject::qt_metacall(_c, _id, _a);
    if (_id < 0)
        return _id;
    if (_c == QMetaObject::InvokeMetaMethod) {
        if (_id < 7)
            qt_static_metacall(this, _c, _id, _a);
        _id -= 7;
    } else if (_c == QMetaObject::RegisterMethodArgumentMetaType) {
        if (_id < 7)
            *reinterpret_cast<int*>(_a[0]) = -1;
        _id -= 7;
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
void ApproveListingCtx721036::remoteChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 0, _a);
}

// SIGNAL 1
void ApproveListingCtx721036::transportChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 1, _a);
}

// SIGNAL 2
void ApproveListingCtx721036::endpointChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 2, _a);
}

// SIGNAL 3
void ApproveListingCtx721036::fromChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 3, _a);
}

// SIGNAL 4
void ApproveListingCtx721036::messageChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 4, _a);
}

// SIGNAL 5
void ApproveListingCtx721036::rawDataChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 5, _a);
}

// SIGNAL 6
void ApproveListingCtx721036::hashChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 6, _a);
}
struct qt_meta_stringdata_ApproveNewAccountCtx721036_t {
    QByteArrayData data[16];
    char stringdata0[212];
};
#define QT_MOC_LITERAL(idx, ofs, len) \
    Q_STATIC_BYTE_ARRAY_DATA_HEADER_INITIALIZER_WITH_OFFSET(len, \
    qptrdiff(offsetof(qt_meta_stringdata_ApproveNewAccountCtx721036_t, stringdata0) + ofs \
        - idx * sizeof(QByteArrayData)) \
    )
static const qt_meta_stringdata_ApproveNewAccountCtx721036_t qt_meta_stringdata_ApproveNewAccountCtx721036 = {
    {
QT_MOC_LITERAL(0, 0, 26), // "ApproveNewAccountCtx721036"
QT_MOC_LITERAL(1, 27, 7), // "clicked"
QT_MOC_LITERAL(2, 35, 0), // ""
QT_MOC_LITERAL(3, 36, 1), // "b"
QT_MOC_LITERAL(4, 38, 14), // "passwordEdited"
QT_MOC_LITERAL(5, 53, 21), // "confirmPasswordEdited"
QT_MOC_LITERAL(6, 75, 13), // "remoteChanged"
QT_MOC_LITERAL(7, 89, 6), // "remote"
QT_MOC_LITERAL(8, 96, 16), // "transportChanged"
QT_MOC_LITERAL(9, 113, 9), // "transport"
QT_MOC_LITERAL(10, 123, 15), // "endpointChanged"
QT_MOC_LITERAL(11, 139, 8), // "endpoint"
QT_MOC_LITERAL(12, 148, 15), // "passwordChanged"
QT_MOC_LITERAL(13, 164, 8), // "password"
QT_MOC_LITERAL(14, 173, 22), // "confirmPasswordChanged"
QT_MOC_LITERAL(15, 196, 15) // "confirmPassword"

    },
    "ApproveNewAccountCtx721036\0clicked\0\0"
    "b\0passwordEdited\0confirmPasswordEdited\0"
    "remoteChanged\0remote\0transportChanged\0"
    "transport\0endpointChanged\0endpoint\0"
    "passwordChanged\0password\0"
    "confirmPasswordChanged\0confirmPassword"
};
#undef QT_MOC_LITERAL

static const uint qt_meta_data_ApproveNewAccountCtx721036[] = {

 // content:
       7,       // revision
       0,       // classname
       0,    0, // classinfo
       8,   14, // methods
       5,   78, // properties
       0,    0, // enums/sets
       0,    0, // constructors
       0,       // flags
       8,       // signalCount

 // signals: name, argc, parameters, tag, flags
       1,    1,   54,    2, 0x06 /* Public */,
       4,    1,   57,    2, 0x06 /* Public */,
       5,    1,   60,    2, 0x06 /* Public */,
       6,    1,   63,    2, 0x06 /* Public */,
       8,    1,   66,    2, 0x06 /* Public */,
      10,    1,   69,    2, 0x06 /* Public */,
      12,    1,   72,    2, 0x06 /* Public */,
      14,    1,   75,    2, 0x06 /* Public */,

 // signals: parameters
    QMetaType::Void, QMetaType::Int,    3,
    QMetaType::Void, QMetaType::QString,    3,
    QMetaType::Void, QMetaType::QString,    3,
    QMetaType::Void, QMetaType::QString,    7,
    QMetaType::Void, QMetaType::QString,    9,
    QMetaType::Void, QMetaType::QString,   11,
    QMetaType::Void, QMetaType::QString,   13,
    QMetaType::Void, QMetaType::QString,   15,

 // properties: name, type, flags
       7, QMetaType::QString, 0x00495103,
       9, QMetaType::QString, 0x00495103,
      11, QMetaType::QString, 0x00495103,
      13, QMetaType::QString, 0x00495103,
      15, QMetaType::QString, 0x00495103,

 // properties: notify_signal_id
       3,
       4,
       5,
       6,
       7,

       0        // eod
};

void ApproveNewAccountCtx721036::qt_static_metacall(QObject *_o, QMetaObject::Call _c, int _id, void **_a)
{
    if (_c == QMetaObject::InvokeMetaMethod) {
        ApproveNewAccountCtx721036 *_t = static_cast<ApproveNewAccountCtx721036 *>(_o);
        Q_UNUSED(_t)
        switch (_id) {
        case 0: _t->clicked((*reinterpret_cast< qint32(*)>(_a[1]))); break;
        case 1: _t->passwordEdited((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 2: _t->confirmPasswordEdited((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 3: _t->remoteChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 4: _t->transportChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 5: _t->endpointChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 6: _t->passwordChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 7: _t->confirmPasswordChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        default: ;
        }
    } else if (_c == QMetaObject::IndexOfMethod) {
        int *result = reinterpret_cast<int *>(_a[0]);
        {
            using _t = void (ApproveNewAccountCtx721036::*)(qint32 );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveNewAccountCtx721036::clicked)) {
                *result = 0;
                return;
            }
        }
        {
            using _t = void (ApproveNewAccountCtx721036::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveNewAccountCtx721036::passwordEdited)) {
                *result = 1;
                return;
            }
        }
        {
            using _t = void (ApproveNewAccountCtx721036::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveNewAccountCtx721036::confirmPasswordEdited)) {
                *result = 2;
                return;
            }
        }
        {
            using _t = void (ApproveNewAccountCtx721036::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveNewAccountCtx721036::remoteChanged)) {
                *result = 3;
                return;
            }
        }
        {
            using _t = void (ApproveNewAccountCtx721036::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveNewAccountCtx721036::transportChanged)) {
                *result = 4;
                return;
            }
        }
        {
            using _t = void (ApproveNewAccountCtx721036::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveNewAccountCtx721036::endpointChanged)) {
                *result = 5;
                return;
            }
        }
        {
            using _t = void (ApproveNewAccountCtx721036::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveNewAccountCtx721036::passwordChanged)) {
                *result = 6;
                return;
            }
        }
        {
            using _t = void (ApproveNewAccountCtx721036::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveNewAccountCtx721036::confirmPasswordChanged)) {
                *result = 7;
                return;
            }
        }
    }
#ifndef QT_NO_PROPERTIES
    else if (_c == QMetaObject::ReadProperty) {
        ApproveNewAccountCtx721036 *_t = static_cast<ApproveNewAccountCtx721036 *>(_o);
        Q_UNUSED(_t)
        void *_v = _a[0];
        switch (_id) {
        case 0: *reinterpret_cast< QString*>(_v) = _t->remote(); break;
        case 1: *reinterpret_cast< QString*>(_v) = _t->transport(); break;
        case 2: *reinterpret_cast< QString*>(_v) = _t->endpoint(); break;
        case 3: *reinterpret_cast< QString*>(_v) = _t->password(); break;
        case 4: *reinterpret_cast< QString*>(_v) = _t->confirmPassword(); break;
        default: break;
        }
    } else if (_c == QMetaObject::WriteProperty) {
        ApproveNewAccountCtx721036 *_t = static_cast<ApproveNewAccountCtx721036 *>(_o);
        Q_UNUSED(_t)
        void *_v = _a[0];
        switch (_id) {
        case 0: _t->setRemote(*reinterpret_cast< QString*>(_v)); break;
        case 1: _t->setTransport(*reinterpret_cast< QString*>(_v)); break;
        case 2: _t->setEndpoint(*reinterpret_cast< QString*>(_v)); break;
        case 3: _t->setPassword(*reinterpret_cast< QString*>(_v)); break;
        case 4: _t->setConfirmPassword(*reinterpret_cast< QString*>(_v)); break;
        default: break;
        }
    } else if (_c == QMetaObject::ResetProperty) {
    }
#endif // QT_NO_PROPERTIES
}

QT_INIT_METAOBJECT const QMetaObject ApproveNewAccountCtx721036::staticMetaObject = {
    { &QObject::staticMetaObject, qt_meta_stringdata_ApproveNewAccountCtx721036.data,
      qt_meta_data_ApproveNewAccountCtx721036,  qt_static_metacall, nullptr, nullptr}
};


const QMetaObject *ApproveNewAccountCtx721036::metaObject() const
{
    return QObject::d_ptr->metaObject ? QObject::d_ptr->dynamicMetaObject() : &staticMetaObject;
}

void *ApproveNewAccountCtx721036::qt_metacast(const char *_clname)
{
    if (!_clname) return nullptr;
    if (!strcmp(_clname, qt_meta_stringdata_ApproveNewAccountCtx721036.stringdata0))
        return static_cast<void*>(this);
    return QObject::qt_metacast(_clname);
}

int ApproveNewAccountCtx721036::qt_metacall(QMetaObject::Call _c, int _id, void **_a)
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
        _id -= 5;
    } else if (_c == QMetaObject::QueryPropertyDesignable) {
        _id -= 5;
    } else if (_c == QMetaObject::QueryPropertyScriptable) {
        _id -= 5;
    } else if (_c == QMetaObject::QueryPropertyStored) {
        _id -= 5;
    } else if (_c == QMetaObject::QueryPropertyEditable) {
        _id -= 5;
    } else if (_c == QMetaObject::QueryPropertyUser) {
        _id -= 5;
    }
#endif // QT_NO_PROPERTIES
    return _id;
}

// SIGNAL 0
void ApproveNewAccountCtx721036::clicked(qint32 _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 0, _a);
}

// SIGNAL 1
void ApproveNewAccountCtx721036::passwordEdited(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 1, _a);
}

// SIGNAL 2
void ApproveNewAccountCtx721036::confirmPasswordEdited(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 2, _a);
}

// SIGNAL 3
void ApproveNewAccountCtx721036::remoteChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 3, _a);
}

// SIGNAL 4
void ApproveNewAccountCtx721036::transportChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 4, _a);
}

// SIGNAL 5
void ApproveNewAccountCtx721036::endpointChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 5, _a);
}

// SIGNAL 6
void ApproveNewAccountCtx721036::passwordChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 6, _a);
}

// SIGNAL 7
void ApproveNewAccountCtx721036::confirmPasswordChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 7, _a);
}
QT_WARNING_POP
QT_END_MOC_NAMESPACE
