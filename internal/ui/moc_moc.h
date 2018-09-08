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
struct qt_meta_stringdata_TxListCtx721036_t {
    QByteArrayData data[4];
    char stringdata0[27];
};
#define QT_MOC_LITERAL(idx, ofs, len) \
    Q_STATIC_BYTE_ARRAY_DATA_HEADER_INITIALIZER_WITH_OFFSET(len, \
    qptrdiff(offsetof(qt_meta_stringdata_TxListCtx721036_t, stringdata0) + ofs \
        - idx * sizeof(QByteArrayData)) \
    )
static const qt_meta_stringdata_TxListCtx721036_t qt_meta_stringdata_TxListCtx721036 = {
    {
QT_MOC_LITERAL(0, 0, 15), // "TxListCtx721036"
QT_MOC_LITERAL(1, 16, 7), // "clicked"
QT_MOC_LITERAL(2, 24, 0), // ""
QT_MOC_LITERAL(3, 25, 1) // "b"

    },
    "TxListCtx721036\0clicked\0\0b"
};
#undef QT_MOC_LITERAL

static const uint qt_meta_data_TxListCtx721036[] = {

 // content:
       7,       // revision
       0,       // classname
       0,    0, // classinfo
       1,   14, // methods
       0,    0, // properties
       0,    0, // enums/sets
       0,    0, // constructors
       0,       // flags
       1,       // signalCount

 // signals: name, argc, parameters, tag, flags
       1,    1,   19,    2, 0x06 /* Public */,

 // signals: parameters
    QMetaType::Void, QMetaType::Int,    3,

       0        // eod
};

void TxListCtx721036::qt_static_metacall(QObject *_o, QMetaObject::Call _c, int _id, void **_a)
{
    if (_c == QMetaObject::InvokeMetaMethod) {
        TxListCtx721036 *_t = static_cast<TxListCtx721036 *>(_o);
        Q_UNUSED(_t)
        switch (_id) {
        case 0: _t->clicked((*reinterpret_cast< qint32(*)>(_a[1]))); break;
        default: ;
        }
    } else if (_c == QMetaObject::IndexOfMethod) {
        int *result = reinterpret_cast<int *>(_a[0]);
        {
            using _t = void (TxListCtx721036::*)(qint32 );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&TxListCtx721036::clicked)) {
                *result = 0;
                return;
            }
        }
    }
}

QT_INIT_METAOBJECT const QMetaObject TxListCtx721036::staticMetaObject = {
    { &QObject::staticMetaObject, qt_meta_stringdata_TxListCtx721036.data,
      qt_meta_data_TxListCtx721036,  qt_static_metacall, nullptr, nullptr}
};


const QMetaObject *TxListCtx721036::metaObject() const
{
    return QObject::d_ptr->metaObject ? QObject::d_ptr->dynamicMetaObject() : &staticMetaObject;
}

void *TxListCtx721036::qt_metacast(const char *_clname)
{
    if (!_clname) return nullptr;
    if (!strcmp(_clname, qt_meta_stringdata_TxListCtx721036.stringdata0))
        return static_cast<void*>(this);
    return QObject::qt_metacast(_clname);
}

int TxListCtx721036::qt_metacall(QMetaObject::Call _c, int _id, void **_a)
{
    _id = QObject::qt_metacall(_c, _id, _a);
    if (_id < 0)
        return _id;
    if (_c == QMetaObject::InvokeMetaMethod) {
        if (_id < 1)
            qt_static_metacall(this, _c, _id, _a);
        _id -= 1;
    } else if (_c == QMetaObject::RegisterMethodArgumentMetaType) {
        if (_id < 1)
            *reinterpret_cast<int*>(_a[0]) = -1;
        _id -= 1;
    }
    return _id;
}

// SIGNAL 0
void TxListCtx721036::clicked(qint32 _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 0, _a);
}
struct qt_meta_stringdata_TxListModel721036_t {
    QByteArrayData data[8];
    char stringdata0[50];
};
#define QT_MOC_LITERAL(idx, ofs, len) \
    Q_STATIC_BYTE_ARRAY_DATA_HEADER_INITIALIZER_WITH_OFFSET(len, \
    qptrdiff(offsetof(qt_meta_stringdata_TxListModel721036_t, stringdata0) + ofs \
        - idx * sizeof(QByteArrayData)) \
    )
static const qt_meta_stringdata_TxListModel721036_t qt_meta_stringdata_TxListModel721036 = {
    {
QT_MOC_LITERAL(0, 0, 17), // "TxListModel721036"
QT_MOC_LITERAL(1, 18, 5), // "clear"
QT_MOC_LITERAL(2, 24, 0), // ""
QT_MOC_LITERAL(3, 25, 3), // "add"
QT_MOC_LITERAL(4, 29, 8), // "quintptr"
QT_MOC_LITERAL(5, 38, 2), // "tx"
QT_MOC_LITERAL(6, 41, 6), // "remove"
QT_MOC_LITERAL(7, 48, 1) // "i"

    },
    "TxListModel721036\0clear\0\0add\0quintptr\0"
    "tx\0remove\0i"
};
#undef QT_MOC_LITERAL

static const uint qt_meta_data_TxListModel721036[] = {

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
       6,    1,   33,    2, 0x06 /* Public */,

 // signals: parameters
    QMetaType::Void,
    QMetaType::Void, 0x80000000 | 4,    5,
    QMetaType::Void, QMetaType::Int,    7,

       0        // eod
};

void TxListModel721036::qt_static_metacall(QObject *_o, QMetaObject::Call _c, int _id, void **_a)
{
    if (_c == QMetaObject::InvokeMetaMethod) {
        TxListModel721036 *_t = static_cast<TxListModel721036 *>(_o);
        Q_UNUSED(_t)
        switch (_id) {
        case 0: _t->clear(); break;
        case 1: _t->add((*reinterpret_cast< quintptr(*)>(_a[1]))); break;
        case 2: _t->remove((*reinterpret_cast< qint32(*)>(_a[1]))); break;
        default: ;
        }
    } else if (_c == QMetaObject::IndexOfMethod) {
        int *result = reinterpret_cast<int *>(_a[0]);
        {
            using _t = void (TxListModel721036::*)();
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&TxListModel721036::clear)) {
                *result = 0;
                return;
            }
        }
        {
            using _t = void (TxListModel721036::*)(quintptr );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&TxListModel721036::add)) {
                *result = 1;
                return;
            }
        }
        {
            using _t = void (TxListModel721036::*)(qint32 );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&TxListModel721036::remove)) {
                *result = 2;
                return;
            }
        }
    }
}

QT_INIT_METAOBJECT const QMetaObject TxListModel721036::staticMetaObject = {
    { &QAbstractListModel::staticMetaObject, qt_meta_stringdata_TxListModel721036.data,
      qt_meta_data_TxListModel721036,  qt_static_metacall, nullptr, nullptr}
};


const QMetaObject *TxListModel721036::metaObject() const
{
    return QObject::d_ptr->metaObject ? QObject::d_ptr->dynamicMetaObject() : &staticMetaObject;
}

void *TxListModel721036::qt_metacast(const char *_clname)
{
    if (!_clname) return nullptr;
    if (!strcmp(_clname, qt_meta_stringdata_TxListModel721036.stringdata0))
        return static_cast<void*>(this);
    return QAbstractListModel::qt_metacast(_clname);
}

int TxListModel721036::qt_metacall(QMetaObject::Call _c, int _id, void **_a)
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
void TxListModel721036::clear()
{
    QMetaObject::activate(this, &staticMetaObject, 0, nullptr);
}

// SIGNAL 1
void TxListModel721036::add(quintptr _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 1, _a);
}

// SIGNAL 2
void TxListModel721036::remove(qint32 _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 2, _a);
}
struct qt_meta_stringdata_ApproveExportCtx721036_t {
    QByteArrayData data[16];
    char stringdata0[175];
};
#define QT_MOC_LITERAL(idx, ofs, len) \
    Q_STATIC_BYTE_ARRAY_DATA_HEADER_INITIALIZER_WITH_OFFSET(len, \
    qptrdiff(offsetof(qt_meta_stringdata_ApproveExportCtx721036_t, stringdata0) + ofs \
        - idx * sizeof(QByteArrayData)) \
    )
static const qt_meta_stringdata_ApproveExportCtx721036_t qt_meta_stringdata_ApproveExportCtx721036 = {
    {
QT_MOC_LITERAL(0, 0, 22), // "ApproveExportCtx721036"
QT_MOC_LITERAL(1, 23, 7), // "clicked"
QT_MOC_LITERAL(2, 31, 0), // ""
QT_MOC_LITERAL(3, 32, 1), // "b"
QT_MOC_LITERAL(4, 34, 4), // "back"
QT_MOC_LITERAL(5, 39, 14), // "passwordEdited"
QT_MOC_LITERAL(6, 54, 13), // "remoteChanged"
QT_MOC_LITERAL(7, 68, 6), // "remote"
QT_MOC_LITERAL(8, 75, 16), // "transportChanged"
QT_MOC_LITERAL(9, 92, 9), // "transport"
QT_MOC_LITERAL(10, 102, 15), // "endpointChanged"
QT_MOC_LITERAL(11, 118, 8), // "endpoint"
QT_MOC_LITERAL(12, 127, 14), // "addressChanged"
QT_MOC_LITERAL(13, 142, 7), // "address"
QT_MOC_LITERAL(14, 150, 15), // "passwordChanged"
QT_MOC_LITERAL(15, 166, 8) // "password"

    },
    "ApproveExportCtx721036\0clicked\0\0b\0"
    "back\0passwordEdited\0remoteChanged\0"
    "remote\0transportChanged\0transport\0"
    "endpointChanged\0endpoint\0addressChanged\0"
    "address\0passwordChanged\0password"
};
#undef QT_MOC_LITERAL

static const uint qt_meta_data_ApproveExportCtx721036[] = {

 // content:
       7,       // revision
       0,       // classname
       0,    0, // classinfo
       8,   14, // methods
       5,   76, // properties
       0,    0, // enums/sets
       0,    0, // constructors
       0,       // flags
       8,       // signalCount

 // signals: name, argc, parameters, tag, flags
       1,    1,   54,    2, 0x06 /* Public */,
       4,    0,   57,    2, 0x06 /* Public */,
       5,    1,   58,    2, 0x06 /* Public */,
       6,    1,   61,    2, 0x06 /* Public */,
       8,    1,   64,    2, 0x06 /* Public */,
      10,    1,   67,    2, 0x06 /* Public */,
      12,    1,   70,    2, 0x06 /* Public */,
      14,    1,   73,    2, 0x06 /* Public */,

 // signals: parameters
    QMetaType::Void, QMetaType::Int,    3,
    QMetaType::Void,
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

void ApproveExportCtx721036::qt_static_metacall(QObject *_o, QMetaObject::Call _c, int _id, void **_a)
{
    if (_c == QMetaObject::InvokeMetaMethod) {
        ApproveExportCtx721036 *_t = static_cast<ApproveExportCtx721036 *>(_o);
        Q_UNUSED(_t)
        switch (_id) {
        case 0: _t->clicked((*reinterpret_cast< qint32(*)>(_a[1]))); break;
        case 1: _t->back(); break;
        case 2: _t->passwordEdited((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 3: _t->remoteChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 4: _t->transportChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 5: _t->endpointChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 6: _t->addressChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 7: _t->passwordChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        default: ;
        }
    } else if (_c == QMetaObject::IndexOfMethod) {
        int *result = reinterpret_cast<int *>(_a[0]);
        {
            using _t = void (ApproveExportCtx721036::*)(qint32 );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveExportCtx721036::clicked)) {
                *result = 0;
                return;
            }
        }
        {
            using _t = void (ApproveExportCtx721036::*)();
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveExportCtx721036::back)) {
                *result = 1;
                return;
            }
        }
        {
            using _t = void (ApproveExportCtx721036::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveExportCtx721036::passwordEdited)) {
                *result = 2;
                return;
            }
        }
        {
            using _t = void (ApproveExportCtx721036::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveExportCtx721036::remoteChanged)) {
                *result = 3;
                return;
            }
        }
        {
            using _t = void (ApproveExportCtx721036::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveExportCtx721036::transportChanged)) {
                *result = 4;
                return;
            }
        }
        {
            using _t = void (ApproveExportCtx721036::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveExportCtx721036::endpointChanged)) {
                *result = 5;
                return;
            }
        }
        {
            using _t = void (ApproveExportCtx721036::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveExportCtx721036::addressChanged)) {
                *result = 6;
                return;
            }
        }
        {
            using _t = void (ApproveExportCtx721036::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveExportCtx721036::passwordChanged)) {
                *result = 7;
                return;
            }
        }
    }
#ifndef QT_NO_PROPERTIES
    else if (_c == QMetaObject::ReadProperty) {
        ApproveExportCtx721036 *_t = static_cast<ApproveExportCtx721036 *>(_o);
        Q_UNUSED(_t)
        void *_v = _a[0];
        switch (_id) {
        case 0: *reinterpret_cast< QString*>(_v) = _t->remote(); break;
        case 1: *reinterpret_cast< QString*>(_v) = _t->transport(); break;
        case 2: *reinterpret_cast< QString*>(_v) = _t->endpoint(); break;
        case 3: *reinterpret_cast< QString*>(_v) = _t->address(); break;
        case 4: *reinterpret_cast< QString*>(_v) = _t->password(); break;
        default: break;
        }
    } else if (_c == QMetaObject::WriteProperty) {
        ApproveExportCtx721036 *_t = static_cast<ApproveExportCtx721036 *>(_o);
        Q_UNUSED(_t)
        void *_v = _a[0];
        switch (_id) {
        case 0: _t->setRemote(*reinterpret_cast< QString*>(_v)); break;
        case 1: _t->setTransport(*reinterpret_cast< QString*>(_v)); break;
        case 2: _t->setEndpoint(*reinterpret_cast< QString*>(_v)); break;
        case 3: _t->setAddress(*reinterpret_cast< QString*>(_v)); break;
        case 4: _t->setPassword(*reinterpret_cast< QString*>(_v)); break;
        default: break;
        }
    } else if (_c == QMetaObject::ResetProperty) {
    }
#endif // QT_NO_PROPERTIES
}

QT_INIT_METAOBJECT const QMetaObject ApproveExportCtx721036::staticMetaObject = {
    { &QObject::staticMetaObject, qt_meta_stringdata_ApproveExportCtx721036.data,
      qt_meta_data_ApproveExportCtx721036,  qt_static_metacall, nullptr, nullptr}
};


const QMetaObject *ApproveExportCtx721036::metaObject() const
{
    return QObject::d_ptr->metaObject ? QObject::d_ptr->dynamicMetaObject() : &staticMetaObject;
}

void *ApproveExportCtx721036::qt_metacast(const char *_clname)
{
    if (!_clname) return nullptr;
    if (!strcmp(_clname, qt_meta_stringdata_ApproveExportCtx721036.stringdata0))
        return static_cast<void*>(this);
    return QObject::qt_metacast(_clname);
}

int ApproveExportCtx721036::qt_metacall(QMetaObject::Call _c, int _id, void **_a)
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
void ApproveExportCtx721036::clicked(qint32 _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 0, _a);
}

// SIGNAL 1
void ApproveExportCtx721036::back()
{
    QMetaObject::activate(this, &staticMetaObject, 1, nullptr);
}

// SIGNAL 2
void ApproveExportCtx721036::passwordEdited(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 2, _a);
}

// SIGNAL 3
void ApproveExportCtx721036::remoteChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 3, _a);
}

// SIGNAL 4
void ApproveExportCtx721036::transportChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 4, _a);
}

// SIGNAL 5
void ApproveExportCtx721036::endpointChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 5, _a);
}

// SIGNAL 6
void ApproveExportCtx721036::addressChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 6, _a);
}

// SIGNAL 7
void ApproveExportCtx721036::passwordChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 7, _a);
}
struct qt_meta_stringdata_ApproveImportCtx721036_t {
    QByteArrayData data[20];
    char stringdata0[262];
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
QT_MOC_LITERAL(4, 34, 4), // "back"
QT_MOC_LITERAL(5, 39, 14), // "passwordEdited"
QT_MOC_LITERAL(6, 54, 21), // "confirmPasswordEdited"
QT_MOC_LITERAL(7, 76, 17), // "oldPasswordEdited"
QT_MOC_LITERAL(8, 94, 13), // "remoteChanged"
QT_MOC_LITERAL(9, 108, 6), // "remote"
QT_MOC_LITERAL(10, 115, 16), // "transportChanged"
QT_MOC_LITERAL(11, 132, 9), // "transport"
QT_MOC_LITERAL(12, 142, 15), // "endpointChanged"
QT_MOC_LITERAL(13, 158, 8), // "endpoint"
QT_MOC_LITERAL(14, 167, 18), // "oldPasswordChanged"
QT_MOC_LITERAL(15, 186, 11), // "oldPassword"
QT_MOC_LITERAL(16, 198, 15), // "passwordChanged"
QT_MOC_LITERAL(17, 214, 8), // "password"
QT_MOC_LITERAL(18, 223, 22), // "confirmPasswordChanged"
QT_MOC_LITERAL(19, 246, 15) // "confirmPassword"

    },
    "ApproveImportCtx721036\0clicked\0\0b\0"
    "back\0passwordEdited\0confirmPasswordEdited\0"
    "oldPasswordEdited\0remoteChanged\0remote\0"
    "transportChanged\0transport\0endpointChanged\0"
    "endpoint\0oldPasswordChanged\0oldPassword\0"
    "passwordChanged\0password\0"
    "confirmPasswordChanged\0confirmPassword"
};
#undef QT_MOC_LITERAL

static const uint qt_meta_data_ApproveImportCtx721036[] = {

 // content:
       7,       // revision
       0,       // classname
       0,    0, // classinfo
      11,   14, // methods
       6,  100, // properties
       0,    0, // enums/sets
       0,    0, // constructors
       0,       // flags
      11,       // signalCount

 // signals: name, argc, parameters, tag, flags
       1,    1,   69,    2, 0x06 /* Public */,
       4,    0,   72,    2, 0x06 /* Public */,
       5,    1,   73,    2, 0x06 /* Public */,
       6,    1,   76,    2, 0x06 /* Public */,
       7,    1,   79,    2, 0x06 /* Public */,
       8,    1,   82,    2, 0x06 /* Public */,
      10,    1,   85,    2, 0x06 /* Public */,
      12,    1,   88,    2, 0x06 /* Public */,
      14,    1,   91,    2, 0x06 /* Public */,
      16,    1,   94,    2, 0x06 /* Public */,
      18,    1,   97,    2, 0x06 /* Public */,

 // signals: parameters
    QMetaType::Void, QMetaType::Int,    3,
    QMetaType::Void,
    QMetaType::Void, QMetaType::QString,    3,
    QMetaType::Void, QMetaType::QString,    3,
    QMetaType::Void, QMetaType::QString,    3,
    QMetaType::Void, QMetaType::QString,    9,
    QMetaType::Void, QMetaType::QString,   11,
    QMetaType::Void, QMetaType::QString,   13,
    QMetaType::Void, QMetaType::QString,   15,
    QMetaType::Void, QMetaType::QString,   17,
    QMetaType::Void, QMetaType::QString,   19,

 // properties: name, type, flags
       9, QMetaType::QString, 0x00495103,
      11, QMetaType::QString, 0x00495103,
      13, QMetaType::QString, 0x00495103,
      15, QMetaType::QString, 0x00495103,
      17, QMetaType::QString, 0x00495103,
      19, QMetaType::QString, 0x00495103,

 // properties: notify_signal_id
       5,
       6,
       7,
       8,
       9,
      10,

       0        // eod
};

void ApproveImportCtx721036::qt_static_metacall(QObject *_o, QMetaObject::Call _c, int _id, void **_a)
{
    if (_c == QMetaObject::InvokeMetaMethod) {
        ApproveImportCtx721036 *_t = static_cast<ApproveImportCtx721036 *>(_o);
        Q_UNUSED(_t)
        switch (_id) {
        case 0: _t->clicked((*reinterpret_cast< qint32(*)>(_a[1]))); break;
        case 1: _t->back(); break;
        case 2: _t->passwordEdited((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 3: _t->confirmPasswordEdited((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 4: _t->oldPasswordEdited((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 5: _t->remoteChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 6: _t->transportChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 7: _t->endpointChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 8: _t->oldPasswordChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 9: _t->passwordChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 10: _t->confirmPasswordChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
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
            using _t = void (ApproveImportCtx721036::*)();
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveImportCtx721036::back)) {
                *result = 1;
                return;
            }
        }
        {
            using _t = void (ApproveImportCtx721036::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveImportCtx721036::passwordEdited)) {
                *result = 2;
                return;
            }
        }
        {
            using _t = void (ApproveImportCtx721036::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveImportCtx721036::confirmPasswordEdited)) {
                *result = 3;
                return;
            }
        }
        {
            using _t = void (ApproveImportCtx721036::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveImportCtx721036::oldPasswordEdited)) {
                *result = 4;
                return;
            }
        }
        {
            using _t = void (ApproveImportCtx721036::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveImportCtx721036::remoteChanged)) {
                *result = 5;
                return;
            }
        }
        {
            using _t = void (ApproveImportCtx721036::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveImportCtx721036::transportChanged)) {
                *result = 6;
                return;
            }
        }
        {
            using _t = void (ApproveImportCtx721036::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveImportCtx721036::endpointChanged)) {
                *result = 7;
                return;
            }
        }
        {
            using _t = void (ApproveImportCtx721036::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveImportCtx721036::oldPasswordChanged)) {
                *result = 8;
                return;
            }
        }
        {
            using _t = void (ApproveImportCtx721036::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveImportCtx721036::passwordChanged)) {
                *result = 9;
                return;
            }
        }
        {
            using _t = void (ApproveImportCtx721036::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveImportCtx721036::confirmPasswordChanged)) {
                *result = 10;
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
        case 3: *reinterpret_cast< QString*>(_v) = _t->oldPassword(); break;
        case 4: *reinterpret_cast< QString*>(_v) = _t->password(); break;
        case 5: *reinterpret_cast< QString*>(_v) = _t->confirmPassword(); break;
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
        case 3: _t->setOldPassword(*reinterpret_cast< QString*>(_v)); break;
        case 4: _t->setPassword(*reinterpret_cast< QString*>(_v)); break;
        case 5: _t->setConfirmPassword(*reinterpret_cast< QString*>(_v)); break;
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
        if (_id < 11)
            qt_static_metacall(this, _c, _id, _a);
        _id -= 11;
    } else if (_c == QMetaObject::RegisterMethodArgumentMetaType) {
        if (_id < 11)
            *reinterpret_cast<int*>(_a[0]) = -1;
        _id -= 11;
    }
#ifndef QT_NO_PROPERTIES
   else if (_c == QMetaObject::ReadProperty || _c == QMetaObject::WriteProperty
            || _c == QMetaObject::ResetProperty || _c == QMetaObject::RegisterPropertyMetaType) {
        qt_static_metacall(this, _c, _id, _a);
        _id -= 6;
    } else if (_c == QMetaObject::QueryPropertyDesignable) {
        _id -= 6;
    } else if (_c == QMetaObject::QueryPropertyScriptable) {
        _id -= 6;
    } else if (_c == QMetaObject::QueryPropertyStored) {
        _id -= 6;
    } else if (_c == QMetaObject::QueryPropertyEditable) {
        _id -= 6;
    } else if (_c == QMetaObject::QueryPropertyUser) {
        _id -= 6;
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
void ApproveImportCtx721036::back()
{
    QMetaObject::activate(this, &staticMetaObject, 1, nullptr);
}

// SIGNAL 2
void ApproveImportCtx721036::passwordEdited(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 2, _a);
}

// SIGNAL 3
void ApproveImportCtx721036::confirmPasswordEdited(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 3, _a);
}

// SIGNAL 4
void ApproveImportCtx721036::oldPasswordEdited(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 4, _a);
}

// SIGNAL 5
void ApproveImportCtx721036::remoteChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 5, _a);
}

// SIGNAL 6
void ApproveImportCtx721036::transportChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 6, _a);
}

// SIGNAL 7
void ApproveImportCtx721036::endpointChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 7, _a);
}

// SIGNAL 8
void ApproveImportCtx721036::oldPasswordChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 8, _a);
}

// SIGNAL 9
void ApproveImportCtx721036::passwordChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 9, _a);
}

// SIGNAL 10
void ApproveImportCtx721036::confirmPasswordChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 10, _a);
}
struct qt_meta_stringdata_ApproveListingCtx721036_t {
    QByteArrayData data[22];
    char stringdata0[223];
};
#define QT_MOC_LITERAL(idx, ofs, len) \
    Q_STATIC_BYTE_ARRAY_DATA_HEADER_INITIALIZER_WITH_OFFSET(len, \
    qptrdiff(offsetof(qt_meta_stringdata_ApproveListingCtx721036_t, stringdata0) + ofs \
        - idx * sizeof(QByteArrayData)) \
    )
static const qt_meta_stringdata_ApproveListingCtx721036_t qt_meta_stringdata_ApproveListingCtx721036 = {
    {
QT_MOC_LITERAL(0, 0, 23), // "ApproveListingCtx721036"
QT_MOC_LITERAL(1, 24, 7), // "clicked"
QT_MOC_LITERAL(2, 32, 0), // ""
QT_MOC_LITERAL(3, 33, 1), // "b"
QT_MOC_LITERAL(4, 35, 4), // "back"
QT_MOC_LITERAL(5, 40, 19), // "onCheckStateChanged"
QT_MOC_LITERAL(6, 60, 1), // "i"
QT_MOC_LITERAL(7, 62, 7), // "checked"
QT_MOC_LITERAL(8, 70, 13), // "remoteChanged"
QT_MOC_LITERAL(9, 84, 6), // "remote"
QT_MOC_LITERAL(10, 91, 16), // "transportChanged"
QT_MOC_LITERAL(11, 108, 9), // "transport"
QT_MOC_LITERAL(12, 118, 15), // "endpointChanged"
QT_MOC_LITERAL(13, 134, 8), // "endpoint"
QT_MOC_LITERAL(14, 143, 11), // "fromChanged"
QT_MOC_LITERAL(15, 155, 4), // "from"
QT_MOC_LITERAL(16, 160, 14), // "messageChanged"
QT_MOC_LITERAL(17, 175, 7), // "message"
QT_MOC_LITERAL(18, 183, 14), // "rawDataChanged"
QT_MOC_LITERAL(19, 198, 7), // "rawData"
QT_MOC_LITERAL(20, 206, 11), // "hashChanged"
QT_MOC_LITERAL(21, 218, 4) // "hash"

    },
    "ApproveListingCtx721036\0clicked\0\0b\0"
    "back\0onCheckStateChanged\0i\0checked\0"
    "remoteChanged\0remote\0transportChanged\0"
    "transport\0endpointChanged\0endpoint\0"
    "fromChanged\0from\0messageChanged\0message\0"
    "rawDataChanged\0rawData\0hashChanged\0"
    "hash"
};
#undef QT_MOC_LITERAL

static const uint qt_meta_data_ApproveListingCtx721036[] = {

 // content:
       7,       // revision
       0,       // classname
       0,    0, // classinfo
      10,   14, // methods
       7,   94, // properties
       0,    0, // enums/sets
       0,    0, // constructors
       0,       // flags
      10,       // signalCount

 // signals: name, argc, parameters, tag, flags
       1,    1,   64,    2, 0x06 /* Public */,
       4,    0,   67,    2, 0x06 /* Public */,
       5,    2,   68,    2, 0x06 /* Public */,
       8,    1,   73,    2, 0x06 /* Public */,
      10,    1,   76,    2, 0x06 /* Public */,
      12,    1,   79,    2, 0x06 /* Public */,
      14,    1,   82,    2, 0x06 /* Public */,
      16,    1,   85,    2, 0x06 /* Public */,
      18,    1,   88,    2, 0x06 /* Public */,
      20,    1,   91,    2, 0x06 /* Public */,

 // signals: parameters
    QMetaType::Void, QMetaType::Int,    3,
    QMetaType::Void,
    QMetaType::Void, QMetaType::Int, QMetaType::Bool,    6,    7,
    QMetaType::Void, QMetaType::QString,    9,
    QMetaType::Void, QMetaType::QString,   11,
    QMetaType::Void, QMetaType::QString,   13,
    QMetaType::Void, QMetaType::QString,   15,
    QMetaType::Void, QMetaType::QString,   17,
    QMetaType::Void, QMetaType::QString,   19,
    QMetaType::Void, QMetaType::QString,   21,

 // properties: name, type, flags
       9, QMetaType::QString, 0x00495103,
      11, QMetaType::QString, 0x00495103,
      13, QMetaType::QString, 0x00495103,
      15, QMetaType::QString, 0x00495103,
      17, QMetaType::QString, 0x00495103,
      19, QMetaType::QString, 0x00495103,
      21, QMetaType::QString, 0x00495103,

 // properties: notify_signal_id
       3,
       4,
       5,
       6,
       7,
       8,
       9,

       0        // eod
};

void ApproveListingCtx721036::qt_static_metacall(QObject *_o, QMetaObject::Call _c, int _id, void **_a)
{
    if (_c == QMetaObject::InvokeMetaMethod) {
        ApproveListingCtx721036 *_t = static_cast<ApproveListingCtx721036 *>(_o);
        Q_UNUSED(_t)
        switch (_id) {
        case 0: _t->clicked((*reinterpret_cast< qint32(*)>(_a[1]))); break;
        case 1: _t->back(); break;
        case 2: _t->onCheckStateChanged((*reinterpret_cast< qint32(*)>(_a[1])),(*reinterpret_cast< bool(*)>(_a[2]))); break;
        case 3: _t->remoteChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 4: _t->transportChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 5: _t->endpointChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 6: _t->fromChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 7: _t->messageChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 8: _t->rawDataChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 9: _t->hashChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        default: ;
        }
    } else if (_c == QMetaObject::IndexOfMethod) {
        int *result = reinterpret_cast<int *>(_a[0]);
        {
            using _t = void (ApproveListingCtx721036::*)(qint32 );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveListingCtx721036::clicked)) {
                *result = 0;
                return;
            }
        }
        {
            using _t = void (ApproveListingCtx721036::*)();
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveListingCtx721036::back)) {
                *result = 1;
                return;
            }
        }
        {
            using _t = void (ApproveListingCtx721036::*)(qint32 , bool );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveListingCtx721036::onCheckStateChanged)) {
                *result = 2;
                return;
            }
        }
        {
            using _t = void (ApproveListingCtx721036::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveListingCtx721036::remoteChanged)) {
                *result = 3;
                return;
            }
        }
        {
            using _t = void (ApproveListingCtx721036::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveListingCtx721036::transportChanged)) {
                *result = 4;
                return;
            }
        }
        {
            using _t = void (ApproveListingCtx721036::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveListingCtx721036::endpointChanged)) {
                *result = 5;
                return;
            }
        }
        {
            using _t = void (ApproveListingCtx721036::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveListingCtx721036::fromChanged)) {
                *result = 6;
                return;
            }
        }
        {
            using _t = void (ApproveListingCtx721036::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveListingCtx721036::messageChanged)) {
                *result = 7;
                return;
            }
        }
        {
            using _t = void (ApproveListingCtx721036::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveListingCtx721036::rawDataChanged)) {
                *result = 8;
                return;
            }
        }
        {
            using _t = void (ApproveListingCtx721036::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveListingCtx721036::hashChanged)) {
                *result = 9;
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
        if (_id < 10)
            qt_static_metacall(this, _c, _id, _a);
        _id -= 10;
    } else if (_c == QMetaObject::RegisterMethodArgumentMetaType) {
        if (_id < 10)
            *reinterpret_cast<int*>(_a[0]) = -1;
        _id -= 10;
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
void ApproveListingCtx721036::clicked(qint32 _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 0, _a);
}

// SIGNAL 1
void ApproveListingCtx721036::back()
{
    QMetaObject::activate(this, &staticMetaObject, 1, nullptr);
}

// SIGNAL 2
void ApproveListingCtx721036::onCheckStateChanged(qint32 _t1, bool _t2)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)), const_cast<void*>(reinterpret_cast<const void*>(&_t2)) };
    QMetaObject::activate(this, &staticMetaObject, 2, _a);
}

// SIGNAL 3
void ApproveListingCtx721036::remoteChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 3, _a);
}

// SIGNAL 4
void ApproveListingCtx721036::transportChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 4, _a);
}

// SIGNAL 5
void ApproveListingCtx721036::endpointChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 5, _a);
}

// SIGNAL 6
void ApproveListingCtx721036::fromChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 6, _a);
}

// SIGNAL 7
void ApproveListingCtx721036::messageChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 7, _a);
}

// SIGNAL 8
void ApproveListingCtx721036::rawDataChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 8, _a);
}

// SIGNAL 9
void ApproveListingCtx721036::hashChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 9, _a);
}
struct qt_meta_stringdata_ApproveNewAccountCtx721036_t {
    QByteArrayData data[17];
    char stringdata0[217];
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
QT_MOC_LITERAL(4, 38, 4), // "back"
QT_MOC_LITERAL(5, 43, 14), // "passwordEdited"
QT_MOC_LITERAL(6, 58, 21), // "confirmPasswordEdited"
QT_MOC_LITERAL(7, 80, 13), // "remoteChanged"
QT_MOC_LITERAL(8, 94, 6), // "remote"
QT_MOC_LITERAL(9, 101, 16), // "transportChanged"
QT_MOC_LITERAL(10, 118, 9), // "transport"
QT_MOC_LITERAL(11, 128, 15), // "endpointChanged"
QT_MOC_LITERAL(12, 144, 8), // "endpoint"
QT_MOC_LITERAL(13, 153, 15), // "passwordChanged"
QT_MOC_LITERAL(14, 169, 8), // "password"
QT_MOC_LITERAL(15, 178, 22), // "confirmPasswordChanged"
QT_MOC_LITERAL(16, 201, 15) // "confirmPassword"

    },
    "ApproveNewAccountCtx721036\0clicked\0\0"
    "b\0back\0passwordEdited\0confirmPasswordEdited\0"
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
       9,   14, // methods
       5,   84, // properties
       0,    0, // enums/sets
       0,    0, // constructors
       0,       // flags
       9,       // signalCount

 // signals: name, argc, parameters, tag, flags
       1,    1,   59,    2, 0x06 /* Public */,
       4,    0,   62,    2, 0x06 /* Public */,
       5,    1,   63,    2, 0x06 /* Public */,
       6,    1,   66,    2, 0x06 /* Public */,
       7,    1,   69,    2, 0x06 /* Public */,
       9,    1,   72,    2, 0x06 /* Public */,
      11,    1,   75,    2, 0x06 /* Public */,
      13,    1,   78,    2, 0x06 /* Public */,
      15,    1,   81,    2, 0x06 /* Public */,

 // signals: parameters
    QMetaType::Void, QMetaType::Int,    3,
    QMetaType::Void,
    QMetaType::Void, QMetaType::QString,    3,
    QMetaType::Void, QMetaType::QString,    3,
    QMetaType::Void, QMetaType::QString,    8,
    QMetaType::Void, QMetaType::QString,   10,
    QMetaType::Void, QMetaType::QString,   12,
    QMetaType::Void, QMetaType::QString,   14,
    QMetaType::Void, QMetaType::QString,   16,

 // properties: name, type, flags
       8, QMetaType::QString, 0x00495103,
      10, QMetaType::QString, 0x00495103,
      12, QMetaType::QString, 0x00495103,
      14, QMetaType::QString, 0x00495103,
      16, QMetaType::QString, 0x00495103,

 // properties: notify_signal_id
       4,
       5,
       6,
       7,
       8,

       0        // eod
};

void ApproveNewAccountCtx721036::qt_static_metacall(QObject *_o, QMetaObject::Call _c, int _id, void **_a)
{
    if (_c == QMetaObject::InvokeMetaMethod) {
        ApproveNewAccountCtx721036 *_t = static_cast<ApproveNewAccountCtx721036 *>(_o);
        Q_UNUSED(_t)
        switch (_id) {
        case 0: _t->clicked((*reinterpret_cast< qint32(*)>(_a[1]))); break;
        case 1: _t->back(); break;
        case 2: _t->passwordEdited((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 3: _t->confirmPasswordEdited((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 4: _t->remoteChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 5: _t->transportChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 6: _t->endpointChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 7: _t->passwordChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 8: _t->confirmPasswordChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
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
            using _t = void (ApproveNewAccountCtx721036::*)();
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveNewAccountCtx721036::back)) {
                *result = 1;
                return;
            }
        }
        {
            using _t = void (ApproveNewAccountCtx721036::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveNewAccountCtx721036::passwordEdited)) {
                *result = 2;
                return;
            }
        }
        {
            using _t = void (ApproveNewAccountCtx721036::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveNewAccountCtx721036::confirmPasswordEdited)) {
                *result = 3;
                return;
            }
        }
        {
            using _t = void (ApproveNewAccountCtx721036::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveNewAccountCtx721036::remoteChanged)) {
                *result = 4;
                return;
            }
        }
        {
            using _t = void (ApproveNewAccountCtx721036::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveNewAccountCtx721036::transportChanged)) {
                *result = 5;
                return;
            }
        }
        {
            using _t = void (ApproveNewAccountCtx721036::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveNewAccountCtx721036::endpointChanged)) {
                *result = 6;
                return;
            }
        }
        {
            using _t = void (ApproveNewAccountCtx721036::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveNewAccountCtx721036::passwordChanged)) {
                *result = 7;
                return;
            }
        }
        {
            using _t = void (ApproveNewAccountCtx721036::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveNewAccountCtx721036::confirmPasswordChanged)) {
                *result = 8;
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
        if (_id < 9)
            qt_static_metacall(this, _c, _id, _a);
        _id -= 9;
    } else if (_c == QMetaObject::RegisterMethodArgumentMetaType) {
        if (_id < 9)
            *reinterpret_cast<int*>(_a[0]) = -1;
        _id -= 9;
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
void ApproveNewAccountCtx721036::back()
{
    QMetaObject::activate(this, &staticMetaObject, 1, nullptr);
}

// SIGNAL 2
void ApproveNewAccountCtx721036::passwordEdited(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 2, _a);
}

// SIGNAL 3
void ApproveNewAccountCtx721036::confirmPasswordEdited(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 3, _a);
}

// SIGNAL 4
void ApproveNewAccountCtx721036::remoteChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 4, _a);
}

// SIGNAL 5
void ApproveNewAccountCtx721036::transportChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 5, _a);
}

// SIGNAL 6
void ApproveNewAccountCtx721036::endpointChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 6, _a);
}

// SIGNAL 7
void ApproveNewAccountCtx721036::passwordChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 7, _a);
}

// SIGNAL 8
void ApproveNewAccountCtx721036::confirmPasswordChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 8, _a);
}
struct qt_meta_stringdata_ApproveTxCtx721036_t {
    QByteArrayData data[34];
    char stringdata0[311];
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
QT_MOC_LITERAL(4, 30, 4), // "back"
QT_MOC_LITERAL(5, 35, 6), // "edited"
QT_MOC_LITERAL(6, 42, 1), // "s"
QT_MOC_LITERAL(7, 44, 1), // "v"
QT_MOC_LITERAL(8, 46, 13), // "remoteChanged"
QT_MOC_LITERAL(9, 60, 6), // "remote"
QT_MOC_LITERAL(10, 67, 16), // "transportChanged"
QT_MOC_LITERAL(11, 84, 9), // "transport"
QT_MOC_LITERAL(12, 94, 15), // "endpointChanged"
QT_MOC_LITERAL(13, 110, 8), // "endpoint"
QT_MOC_LITERAL(14, 119, 11), // "dataChanged"
QT_MOC_LITERAL(15, 131, 4), // "data"
QT_MOC_LITERAL(16, 136, 11), // "fromChanged"
QT_MOC_LITERAL(17, 148, 4), // "from"
QT_MOC_LITERAL(18, 153, 9), // "toChanged"
QT_MOC_LITERAL(19, 163, 2), // "to"
QT_MOC_LITERAL(20, 166, 10), // "gasChanged"
QT_MOC_LITERAL(21, 177, 3), // "gas"
QT_MOC_LITERAL(22, 181, 15), // "gasPriceChanged"
QT_MOC_LITERAL(23, 197, 8), // "gasPrice"
QT_MOC_LITERAL(24, 206, 12), // "nonceChanged"
QT_MOC_LITERAL(25, 219, 5), // "nonce"
QT_MOC_LITERAL(26, 225, 12), // "valueChanged"
QT_MOC_LITERAL(27, 238, 5), // "value"
QT_MOC_LITERAL(28, 244, 15), // "passwordChanged"
QT_MOC_LITERAL(29, 260, 8), // "password"
QT_MOC_LITERAL(30, 269, 14), // "fromSrcChanged"
QT_MOC_LITERAL(31, 284, 7), // "fromSrc"
QT_MOC_LITERAL(32, 292, 12), // "toSrcChanged"
QT_MOC_LITERAL(33, 305, 5) // "toSrc"

    },
    "ApproveTxCtx721036\0clicked\0\0b\0back\0"
    "edited\0s\0v\0remoteChanged\0remote\0"
    "transportChanged\0transport\0endpointChanged\0"
    "endpoint\0dataChanged\0data\0fromChanged\0"
    "from\0toChanged\0to\0gasChanged\0gas\0"
    "gasPriceChanged\0gasPrice\0nonceChanged\0"
    "nonce\0valueChanged\0value\0passwordChanged\0"
    "password\0fromSrcChanged\0fromSrc\0"
    "toSrcChanged\0toSrc"
};
#undef QT_MOC_LITERAL

static const uint qt_meta_data_ApproveTxCtx721036[] = {

 // content:
       7,       // revision
       0,       // classname
       0,    0, // classinfo
      16,   14, // methods
      13,  142, // properties
       0,    0, // enums/sets
       0,    0, // constructors
       0,       // flags
      16,       // signalCount

 // signals: name, argc, parameters, tag, flags
       1,    1,   94,    2, 0x06 /* Public */,
       4,    0,   97,    2, 0x06 /* Public */,
       5,    2,   98,    2, 0x06 /* Public */,
       8,    1,  103,    2, 0x06 /* Public */,
      10,    1,  106,    2, 0x06 /* Public */,
      12,    1,  109,    2, 0x06 /* Public */,
      14,    1,  112,    2, 0x06 /* Public */,
      16,    1,  115,    2, 0x06 /* Public */,
      18,    1,  118,    2, 0x06 /* Public */,
      20,    1,  121,    2, 0x06 /* Public */,
      22,    1,  124,    2, 0x06 /* Public */,
      24,    1,  127,    2, 0x06 /* Public */,
      26,    1,  130,    2, 0x06 /* Public */,
      28,    1,  133,    2, 0x06 /* Public */,
      30,    1,  136,    2, 0x06 /* Public */,
      32,    1,  139,    2, 0x06 /* Public */,

 // signals: parameters
    QMetaType::Void, QMetaType::Int,    3,
    QMetaType::Void,
    QMetaType::Void, QMetaType::QString, QMetaType::QString,    6,    7,
    QMetaType::Void, QMetaType::QString,    9,
    QMetaType::Void, QMetaType::QString,   11,
    QMetaType::Void, QMetaType::QString,   13,
    QMetaType::Void, QMetaType::QString,   15,
    QMetaType::Void, QMetaType::QString,   17,
    QMetaType::Void, QMetaType::QString,   19,
    QMetaType::Void, QMetaType::QString,   21,
    QMetaType::Void, QMetaType::QString,   23,
    QMetaType::Void, QMetaType::QString,   25,
    QMetaType::Void, QMetaType::QString,   27,
    QMetaType::Void, QMetaType::QString,   29,
    QMetaType::Void, QMetaType::QString,   31,
    QMetaType::Void, QMetaType::QString,   33,

 // properties: name, type, flags
       9, QMetaType::QString, 0x00495103,
      11, QMetaType::QString, 0x00495103,
      13, QMetaType::QString, 0x00495103,
      15, QMetaType::QString, 0x00495103,
      17, QMetaType::QString, 0x00495103,
      19, QMetaType::QString, 0x00495103,
      21, QMetaType::QString, 0x00495103,
      23, QMetaType::QString, 0x00495103,
      25, QMetaType::QString, 0x00495103,
      27, QMetaType::QString, 0x00495103,
      29, QMetaType::QString, 0x00495103,
      31, QMetaType::QString, 0x00495103,
      33, QMetaType::QString, 0x00495103,

 // properties: notify_signal_id
       3,
       4,
       5,
       6,
       7,
       8,
       9,
      10,
      11,
      12,
      13,
      14,
      15,

       0        // eod
};

void ApproveTxCtx721036::qt_static_metacall(QObject *_o, QMetaObject::Call _c, int _id, void **_a)
{
    if (_c == QMetaObject::InvokeMetaMethod) {
        ApproveTxCtx721036 *_t = static_cast<ApproveTxCtx721036 *>(_o);
        Q_UNUSED(_t)
        switch (_id) {
        case 0: _t->clicked((*reinterpret_cast< qint32(*)>(_a[1]))); break;
        case 1: _t->back(); break;
        case 2: _t->edited((*reinterpret_cast< QString(*)>(_a[1])),(*reinterpret_cast< QString(*)>(_a[2]))); break;
        case 3: _t->remoteChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 4: _t->transportChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 5: _t->endpointChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 6: _t->dataChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 7: _t->fromChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 8: _t->toChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 9: _t->gasChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 10: _t->gasPriceChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 11: _t->nonceChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 12: _t->valueChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 13: _t->passwordChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 14: _t->fromSrcChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 15: _t->toSrcChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
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
            using _t = void (ApproveTxCtx721036::*)();
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveTxCtx721036::back)) {
                *result = 1;
                return;
            }
        }
        {
            using _t = void (ApproveTxCtx721036::*)(QString , QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveTxCtx721036::edited)) {
                *result = 2;
                return;
            }
        }
        {
            using _t = void (ApproveTxCtx721036::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveTxCtx721036::remoteChanged)) {
                *result = 3;
                return;
            }
        }
        {
            using _t = void (ApproveTxCtx721036::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveTxCtx721036::transportChanged)) {
                *result = 4;
                return;
            }
        }
        {
            using _t = void (ApproveTxCtx721036::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveTxCtx721036::endpointChanged)) {
                *result = 5;
                return;
            }
        }
        {
            using _t = void (ApproveTxCtx721036::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveTxCtx721036::dataChanged)) {
                *result = 6;
                return;
            }
        }
        {
            using _t = void (ApproveTxCtx721036::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveTxCtx721036::fromChanged)) {
                *result = 7;
                return;
            }
        }
        {
            using _t = void (ApproveTxCtx721036::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveTxCtx721036::toChanged)) {
                *result = 8;
                return;
            }
        }
        {
            using _t = void (ApproveTxCtx721036::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveTxCtx721036::gasChanged)) {
                *result = 9;
                return;
            }
        }
        {
            using _t = void (ApproveTxCtx721036::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveTxCtx721036::gasPriceChanged)) {
                *result = 10;
                return;
            }
        }
        {
            using _t = void (ApproveTxCtx721036::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveTxCtx721036::nonceChanged)) {
                *result = 11;
                return;
            }
        }
        {
            using _t = void (ApproveTxCtx721036::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveTxCtx721036::valueChanged)) {
                *result = 12;
                return;
            }
        }
        {
            using _t = void (ApproveTxCtx721036::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveTxCtx721036::passwordChanged)) {
                *result = 13;
                return;
            }
        }
        {
            using _t = void (ApproveTxCtx721036::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveTxCtx721036::fromSrcChanged)) {
                *result = 14;
                return;
            }
        }
        {
            using _t = void (ApproveTxCtx721036::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveTxCtx721036::toSrcChanged)) {
                *result = 15;
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
        case 10: *reinterpret_cast< QString*>(_v) = _t->password(); break;
        case 11: *reinterpret_cast< QString*>(_v) = _t->fromSrc(); break;
        case 12: *reinterpret_cast< QString*>(_v) = _t->toSrc(); break;
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
        case 10: _t->setPassword(*reinterpret_cast< QString*>(_v)); break;
        case 11: _t->setFromSrc(*reinterpret_cast< QString*>(_v)); break;
        case 12: _t->setToSrc(*reinterpret_cast< QString*>(_v)); break;
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
        if (_id < 16)
            qt_static_metacall(this, _c, _id, _a);
        _id -= 16;
    } else if (_c == QMetaObject::RegisterMethodArgumentMetaType) {
        if (_id < 16)
            *reinterpret_cast<int*>(_a[0]) = -1;
        _id -= 16;
    }
#ifndef QT_NO_PROPERTIES
   else if (_c == QMetaObject::ReadProperty || _c == QMetaObject::WriteProperty
            || _c == QMetaObject::ResetProperty || _c == QMetaObject::RegisterPropertyMetaType) {
        qt_static_metacall(this, _c, _id, _a);
        _id -= 13;
    } else if (_c == QMetaObject::QueryPropertyDesignable) {
        _id -= 13;
    } else if (_c == QMetaObject::QueryPropertyScriptable) {
        _id -= 13;
    } else if (_c == QMetaObject::QueryPropertyStored) {
        _id -= 13;
    } else if (_c == QMetaObject::QueryPropertyEditable) {
        _id -= 13;
    } else if (_c == QMetaObject::QueryPropertyUser) {
        _id -= 13;
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
void ApproveTxCtx721036::back()
{
    QMetaObject::activate(this, &staticMetaObject, 1, nullptr);
}

// SIGNAL 2
void ApproveTxCtx721036::edited(QString _t1, QString _t2)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)), const_cast<void*>(reinterpret_cast<const void*>(&_t2)) };
    QMetaObject::activate(this, &staticMetaObject, 2, _a);
}

// SIGNAL 3
void ApproveTxCtx721036::remoteChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 3, _a);
}

// SIGNAL 4
void ApproveTxCtx721036::transportChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 4, _a);
}

// SIGNAL 5
void ApproveTxCtx721036::endpointChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 5, _a);
}

// SIGNAL 6
void ApproveTxCtx721036::dataChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 6, _a);
}

// SIGNAL 7
void ApproveTxCtx721036::fromChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 7, _a);
}

// SIGNAL 8
void ApproveTxCtx721036::toChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 8, _a);
}

// SIGNAL 9
void ApproveTxCtx721036::gasChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 9, _a);
}

// SIGNAL 10
void ApproveTxCtx721036::gasPriceChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 10, _a);
}

// SIGNAL 11
void ApproveTxCtx721036::nonceChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 11, _a);
}

// SIGNAL 12
void ApproveTxCtx721036::valueChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 12, _a);
}

// SIGNAL 13
void ApproveTxCtx721036::passwordChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 13, _a);
}

// SIGNAL 14
void ApproveTxCtx721036::fromSrcChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 14, _a);
}

// SIGNAL 15
void ApproveTxCtx721036::toSrcChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 15, _a);
}
struct qt_meta_stringdata_CustomListModel721036_t {
    QByteArrayData data[6];
    char stringdata0[50];
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
QT_MOC_LITERAL(5, 42, 7) // "account"

    },
    "CustomListModel721036\0clear\0\0add\0"
    "quintptr\0account"
};
#undef QT_MOC_LITERAL

static const uint qt_meta_data_CustomListModel721036[] = {

 // content:
       7,       // revision
       0,       // classname
       0,    0, // classinfo
       2,   14, // methods
       0,    0, // properties
       0,    0, // enums/sets
       0,    0, // constructors
       0,       // flags
       2,       // signalCount

 // signals: name, argc, parameters, tag, flags
       1,    0,   24,    2, 0x06 /* Public */,
       3,    1,   25,    2, 0x06 /* Public */,

 // signals: parameters
    QMetaType::Void,
    QMetaType::Void, 0x80000000 | 4,    5,

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
        if (_id < 2)
            qt_static_metacall(this, _c, _id, _a);
        _id -= 2;
    } else if (_c == QMetaObject::RegisterMethodArgumentMetaType) {
        if (_id < 2)
            *reinterpret_cast<int*>(_a[0]) = -1;
        _id -= 2;
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
struct qt_meta_stringdata_ApproveSignDataCtx721036_t {
    QByteArrayData data[23];
    char stringdata0[232];
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
QT_MOC_LITERAL(4, 36, 4), // "back"
QT_MOC_LITERAL(5, 41, 6), // "edited"
QT_MOC_LITERAL(6, 48, 5), // "value"
QT_MOC_LITERAL(7, 54, 13), // "remoteChanged"
QT_MOC_LITERAL(8, 68, 6), // "remote"
QT_MOC_LITERAL(9, 75, 16), // "transportChanged"
QT_MOC_LITERAL(10, 92, 9), // "transport"
QT_MOC_LITERAL(11, 102, 15), // "endpointChanged"
QT_MOC_LITERAL(12, 118, 8), // "endpoint"
QT_MOC_LITERAL(13, 127, 11), // "fromChanged"
QT_MOC_LITERAL(14, 139, 4), // "from"
QT_MOC_LITERAL(15, 144, 14), // "messageChanged"
QT_MOC_LITERAL(16, 159, 7), // "message"
QT_MOC_LITERAL(17, 167, 14), // "rawDataChanged"
QT_MOC_LITERAL(18, 182, 7), // "rawData"
QT_MOC_LITERAL(19, 190, 11), // "hashChanged"
QT_MOC_LITERAL(20, 202, 4), // "hash"
QT_MOC_LITERAL(21, 207, 15), // "passwordChanged"
QT_MOC_LITERAL(22, 223, 8) // "password"

    },
    "ApproveSignDataCtx721036\0clicked\0\0b\0"
    "back\0edited\0value\0remoteChanged\0remote\0"
    "transportChanged\0transport\0endpointChanged\0"
    "endpoint\0fromChanged\0from\0messageChanged\0"
    "message\0rawDataChanged\0rawData\0"
    "hashChanged\0hash\0passwordChanged\0"
    "password"
};
#undef QT_MOC_LITERAL

static const uint qt_meta_data_ApproveSignDataCtx721036[] = {

 // content:
       7,       // revision
       0,       // classname
       0,    0, // classinfo
      11,   14, // methods
       8,  102, // properties
       0,    0, // enums/sets
       0,    0, // constructors
       0,       // flags
      11,       // signalCount

 // signals: name, argc, parameters, tag, flags
       1,    1,   69,    2, 0x06 /* Public */,
       4,    0,   72,    2, 0x06 /* Public */,
       5,    2,   73,    2, 0x06 /* Public */,
       7,    1,   78,    2, 0x06 /* Public */,
       9,    1,   81,    2, 0x06 /* Public */,
      11,    1,   84,    2, 0x06 /* Public */,
      13,    1,   87,    2, 0x06 /* Public */,
      15,    1,   90,    2, 0x06 /* Public */,
      17,    1,   93,    2, 0x06 /* Public */,
      19,    1,   96,    2, 0x06 /* Public */,
      21,    1,   99,    2, 0x06 /* Public */,

 // signals: parameters
    QMetaType::Void, QMetaType::Int,    3,
    QMetaType::Void,
    QMetaType::Void, QMetaType::QString, QMetaType::QString,    3,    6,
    QMetaType::Void, QMetaType::QString,    8,
    QMetaType::Void, QMetaType::QString,   10,
    QMetaType::Void, QMetaType::QString,   12,
    QMetaType::Void, QMetaType::QString,   14,
    QMetaType::Void, QMetaType::QString,   16,
    QMetaType::Void, QMetaType::QString,   18,
    QMetaType::Void, QMetaType::QString,   20,
    QMetaType::Void, QMetaType::QString,   22,

 // properties: name, type, flags
       8, QMetaType::QString, 0x00495103,
      10, QMetaType::QString, 0x00495103,
      12, QMetaType::QString, 0x00495103,
      14, QMetaType::QString, 0x00495103,
      16, QMetaType::QString, 0x00495103,
      18, QMetaType::QString, 0x00495103,
      20, QMetaType::QString, 0x00495103,
      22, QMetaType::QString, 0x00495103,

 // properties: notify_signal_id
       3,
       4,
       5,
       6,
       7,
       8,
       9,
      10,

       0        // eod
};

void ApproveSignDataCtx721036::qt_static_metacall(QObject *_o, QMetaObject::Call _c, int _id, void **_a)
{
    if (_c == QMetaObject::InvokeMetaMethod) {
        ApproveSignDataCtx721036 *_t = static_cast<ApproveSignDataCtx721036 *>(_o);
        Q_UNUSED(_t)
        switch (_id) {
        case 0: _t->clicked((*reinterpret_cast< qint32(*)>(_a[1]))); break;
        case 1: _t->back(); break;
        case 2: _t->edited((*reinterpret_cast< QString(*)>(_a[1])),(*reinterpret_cast< QString(*)>(_a[2]))); break;
        case 3: _t->remoteChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 4: _t->transportChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 5: _t->endpointChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 6: _t->fromChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 7: _t->messageChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 8: _t->rawDataChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 9: _t->hashChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 10: _t->passwordChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
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
            using _t = void (ApproveSignDataCtx721036::*)();
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveSignDataCtx721036::back)) {
                *result = 1;
                return;
            }
        }
        {
            using _t = void (ApproveSignDataCtx721036::*)(QString , QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveSignDataCtx721036::edited)) {
                *result = 2;
                return;
            }
        }
        {
            using _t = void (ApproveSignDataCtx721036::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveSignDataCtx721036::remoteChanged)) {
                *result = 3;
                return;
            }
        }
        {
            using _t = void (ApproveSignDataCtx721036::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveSignDataCtx721036::transportChanged)) {
                *result = 4;
                return;
            }
        }
        {
            using _t = void (ApproveSignDataCtx721036::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveSignDataCtx721036::endpointChanged)) {
                *result = 5;
                return;
            }
        }
        {
            using _t = void (ApproveSignDataCtx721036::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveSignDataCtx721036::fromChanged)) {
                *result = 6;
                return;
            }
        }
        {
            using _t = void (ApproveSignDataCtx721036::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveSignDataCtx721036::messageChanged)) {
                *result = 7;
                return;
            }
        }
        {
            using _t = void (ApproveSignDataCtx721036::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveSignDataCtx721036::rawDataChanged)) {
                *result = 8;
                return;
            }
        }
        {
            using _t = void (ApproveSignDataCtx721036::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveSignDataCtx721036::hashChanged)) {
                *result = 9;
                return;
            }
        }
        {
            using _t = void (ApproveSignDataCtx721036::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveSignDataCtx721036::passwordChanged)) {
                *result = 10;
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
        case 7: *reinterpret_cast< QString*>(_v) = _t->password(); break;
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
        case 7: _t->setPassword(*reinterpret_cast< QString*>(_v)); break;
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
        if (_id < 11)
            qt_static_metacall(this, _c, _id, _a);
        _id -= 11;
    } else if (_c == QMetaObject::RegisterMethodArgumentMetaType) {
        if (_id < 11)
            *reinterpret_cast<int*>(_a[0]) = -1;
        _id -= 11;
    }
#ifndef QT_NO_PROPERTIES
   else if (_c == QMetaObject::ReadProperty || _c == QMetaObject::WriteProperty
            || _c == QMetaObject::ResetProperty || _c == QMetaObject::RegisterPropertyMetaType) {
        qt_static_metacall(this, _c, _id, _a);
        _id -= 8;
    } else if (_c == QMetaObject::QueryPropertyDesignable) {
        _id -= 8;
    } else if (_c == QMetaObject::QueryPropertyScriptable) {
        _id -= 8;
    } else if (_c == QMetaObject::QueryPropertyStored) {
        _id -= 8;
    } else if (_c == QMetaObject::QueryPropertyEditable) {
        _id -= 8;
    } else if (_c == QMetaObject::QueryPropertyUser) {
        _id -= 8;
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
void ApproveSignDataCtx721036::back()
{
    QMetaObject::activate(this, &staticMetaObject, 1, nullptr);
}

// SIGNAL 2
void ApproveSignDataCtx721036::edited(QString _t1, QString _t2)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)), const_cast<void*>(reinterpret_cast<const void*>(&_t2)) };
    QMetaObject::activate(this, &staticMetaObject, 2, _a);
}

// SIGNAL 3
void ApproveSignDataCtx721036::remoteChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 3, _a);
}

// SIGNAL 4
void ApproveSignDataCtx721036::transportChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 4, _a);
}

// SIGNAL 5
void ApproveSignDataCtx721036::endpointChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 5, _a);
}

// SIGNAL 6
void ApproveSignDataCtx721036::fromChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 6, _a);
}

// SIGNAL 7
void ApproveSignDataCtx721036::messageChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 7, _a);
}

// SIGNAL 8
void ApproveSignDataCtx721036::rawDataChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 8, _a);
}

// SIGNAL 9
void ApproveSignDataCtx721036::hashChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 9, _a);
}

// SIGNAL 10
void ApproveSignDataCtx721036::passwordChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 10, _a);
}
struct qt_meta_stringdata_LoginCtx721036_t {
    QByteArrayData data[17];
    char stringdata0[179];
};
#define QT_MOC_LITERAL(idx, ofs, len) \
    Q_STATIC_BYTE_ARRAY_DATA_HEADER_INITIALIZER_WITH_OFFSET(len, \
    qptrdiff(offsetof(qt_meta_stringdata_LoginCtx721036_t, stringdata0) + ofs \
        - idx * sizeof(QByteArrayData)) \
    )
static const qt_meta_stringdata_LoginCtx721036_t qt_meta_stringdata_LoginCtx721036 = {
    {
QT_MOC_LITERAL(0, 0, 14), // "LoginCtx721036"
QT_MOC_LITERAL(1, 15, 7), // "clicked"
QT_MOC_LITERAL(2, 23, 0), // ""
QT_MOC_LITERAL(3, 24, 6), // "edited"
QT_MOC_LITERAL(4, 31, 1), // "b"
QT_MOC_LITERAL(5, 33, 13), // "remoteChanged"
QT_MOC_LITERAL(6, 47, 6), // "remote"
QT_MOC_LITERAL(7, 54, 16), // "transportChanged"
QT_MOC_LITERAL(8, 71, 9), // "transport"
QT_MOC_LITERAL(9, 81, 15), // "endpointChanged"
QT_MOC_LITERAL(10, 97, 8), // "endpoint"
QT_MOC_LITERAL(11, 106, 13), // "gopathChanged"
QT_MOC_LITERAL(12, 120, 6), // "gopath"
QT_MOC_LITERAL(13, 127, 17), // "binaryHashChanged"
QT_MOC_LITERAL(14, 145, 10), // "binaryHash"
QT_MOC_LITERAL(15, 156, 14), // "isValidChanged"
QT_MOC_LITERAL(16, 171, 7) // "isValid"

    },
    "LoginCtx721036\0clicked\0\0edited\0b\0"
    "remoteChanged\0remote\0transportChanged\0"
    "transport\0endpointChanged\0endpoint\0"
    "gopathChanged\0gopath\0binaryHashChanged\0"
    "binaryHash\0isValidChanged\0isValid"
};
#undef QT_MOC_LITERAL

static const uint qt_meta_data_LoginCtx721036[] = {

 // content:
       7,       // revision
       0,       // classname
       0,    0, // classinfo
       8,   14, // methods
       6,   76, // properties
       0,    0, // enums/sets
       0,    0, // constructors
       0,       // flags
       8,       // signalCount

 // signals: name, argc, parameters, tag, flags
       1,    0,   54,    2, 0x06 /* Public */,
       3,    1,   55,    2, 0x06 /* Public */,
       5,    1,   58,    2, 0x06 /* Public */,
       7,    1,   61,    2, 0x06 /* Public */,
       9,    1,   64,    2, 0x06 /* Public */,
      11,    1,   67,    2, 0x06 /* Public */,
      13,    1,   70,    2, 0x06 /* Public */,
      15,    1,   73,    2, 0x06 /* Public */,

 // signals: parameters
    QMetaType::Void,
    QMetaType::Void, QMetaType::QString,    4,
    QMetaType::Void, QMetaType::QString,    6,
    QMetaType::Void, QMetaType::QString,    8,
    QMetaType::Void, QMetaType::QString,   10,
    QMetaType::Void, QMetaType::QString,   12,
    QMetaType::Void, QMetaType::QString,   14,
    QMetaType::Void, QMetaType::Bool,   16,

 // properties: name, type, flags
       6, QMetaType::QString, 0x00495103,
       8, QMetaType::QString, 0x00495103,
      10, QMetaType::QString, 0x00495103,
      12, QMetaType::QString, 0x00495103,
      14, QMetaType::QString, 0x00495103,
      16, QMetaType::Bool, 0x00495103,

 // properties: notify_signal_id
       2,
       3,
       4,
       5,
       6,
       7,

       0        // eod
};

void LoginCtx721036::qt_static_metacall(QObject *_o, QMetaObject::Call _c, int _id, void **_a)
{
    if (_c == QMetaObject::InvokeMetaMethod) {
        LoginCtx721036 *_t = static_cast<LoginCtx721036 *>(_o);
        Q_UNUSED(_t)
        switch (_id) {
        case 0: _t->clicked(); break;
        case 1: _t->edited((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 2: _t->remoteChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 3: _t->transportChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 4: _t->endpointChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 5: _t->gopathChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 6: _t->binaryHashChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 7: _t->isValidChanged((*reinterpret_cast< bool(*)>(_a[1]))); break;
        default: ;
        }
    } else if (_c == QMetaObject::IndexOfMethod) {
        int *result = reinterpret_cast<int *>(_a[0]);
        {
            using _t = void (LoginCtx721036::*)();
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&LoginCtx721036::clicked)) {
                *result = 0;
                return;
            }
        }
        {
            using _t = void (LoginCtx721036::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&LoginCtx721036::edited)) {
                *result = 1;
                return;
            }
        }
        {
            using _t = void (LoginCtx721036::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&LoginCtx721036::remoteChanged)) {
                *result = 2;
                return;
            }
        }
        {
            using _t = void (LoginCtx721036::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&LoginCtx721036::transportChanged)) {
                *result = 3;
                return;
            }
        }
        {
            using _t = void (LoginCtx721036::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&LoginCtx721036::endpointChanged)) {
                *result = 4;
                return;
            }
        }
        {
            using _t = void (LoginCtx721036::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&LoginCtx721036::gopathChanged)) {
                *result = 5;
                return;
            }
        }
        {
            using _t = void (LoginCtx721036::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&LoginCtx721036::binaryHashChanged)) {
                *result = 6;
                return;
            }
        }
        {
            using _t = void (LoginCtx721036::*)(bool );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&LoginCtx721036::isValidChanged)) {
                *result = 7;
                return;
            }
        }
    }
#ifndef QT_NO_PROPERTIES
    else if (_c == QMetaObject::ReadProperty) {
        LoginCtx721036 *_t = static_cast<LoginCtx721036 *>(_o);
        Q_UNUSED(_t)
        void *_v = _a[0];
        switch (_id) {
        case 0: *reinterpret_cast< QString*>(_v) = _t->remote(); break;
        case 1: *reinterpret_cast< QString*>(_v) = _t->transport(); break;
        case 2: *reinterpret_cast< QString*>(_v) = _t->endpoint(); break;
        case 3: *reinterpret_cast< QString*>(_v) = _t->gopath(); break;
        case 4: *reinterpret_cast< QString*>(_v) = _t->binaryHash(); break;
        case 5: *reinterpret_cast< bool*>(_v) = _t->isValid(); break;
        default: break;
        }
    } else if (_c == QMetaObject::WriteProperty) {
        LoginCtx721036 *_t = static_cast<LoginCtx721036 *>(_o);
        Q_UNUSED(_t)
        void *_v = _a[0];
        switch (_id) {
        case 0: _t->setRemote(*reinterpret_cast< QString*>(_v)); break;
        case 1: _t->setTransport(*reinterpret_cast< QString*>(_v)); break;
        case 2: _t->setEndpoint(*reinterpret_cast< QString*>(_v)); break;
        case 3: _t->setGopath(*reinterpret_cast< QString*>(_v)); break;
        case 4: _t->setBinaryHash(*reinterpret_cast< QString*>(_v)); break;
        case 5: _t->setIsValid(*reinterpret_cast< bool*>(_v)); break;
        default: break;
        }
    } else if (_c == QMetaObject::ResetProperty) {
    }
#endif // QT_NO_PROPERTIES
}

QT_INIT_METAOBJECT const QMetaObject LoginCtx721036::staticMetaObject = {
    { &QObject::staticMetaObject, qt_meta_stringdata_LoginCtx721036.data,
      qt_meta_data_LoginCtx721036,  qt_static_metacall, nullptr, nullptr}
};


const QMetaObject *LoginCtx721036::metaObject() const
{
    return QObject::d_ptr->metaObject ? QObject::d_ptr->dynamicMetaObject() : &staticMetaObject;
}

void *LoginCtx721036::qt_metacast(const char *_clname)
{
    if (!_clname) return nullptr;
    if (!strcmp(_clname, qt_meta_stringdata_LoginCtx721036.stringdata0))
        return static_cast<void*>(this);
    return QObject::qt_metacast(_clname);
}

int LoginCtx721036::qt_metacall(QMetaObject::Call _c, int _id, void **_a)
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
        _id -= 6;
    } else if (_c == QMetaObject::QueryPropertyDesignable) {
        _id -= 6;
    } else if (_c == QMetaObject::QueryPropertyScriptable) {
        _id -= 6;
    } else if (_c == QMetaObject::QueryPropertyStored) {
        _id -= 6;
    } else if (_c == QMetaObject::QueryPropertyEditable) {
        _id -= 6;
    } else if (_c == QMetaObject::QueryPropertyUser) {
        _id -= 6;
    }
#endif // QT_NO_PROPERTIES
    return _id;
}

// SIGNAL 0
void LoginCtx721036::clicked()
{
    QMetaObject::activate(this, &staticMetaObject, 0, nullptr);
}

// SIGNAL 1
void LoginCtx721036::edited(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 1, _a);
}

// SIGNAL 2
void LoginCtx721036::remoteChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 2, _a);
}

// SIGNAL 3
void LoginCtx721036::transportChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 3, _a);
}

// SIGNAL 4
void LoginCtx721036::endpointChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 4, _a);
}

// SIGNAL 5
void LoginCtx721036::gopathChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 5, _a);
}

// SIGNAL 6
void LoginCtx721036::binaryHashChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 6, _a);
}

// SIGNAL 7
void LoginCtx721036::isValidChanged(bool _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 7, _a);
}
QT_WARNING_POP
QT_END_MOC_NAMESPACE
