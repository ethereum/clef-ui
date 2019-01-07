/****************************************************************************
** Meta object code from reading C++ file 'moc.cpp'
**
** Created by: The Qt Meta Object Compiler version 67 (Qt 5.12.0)
**
** WARNING! All changes made in this file will be lost!
*****************************************************************************/

#include <QtCore/qbytearray.h>
#include <QtCore/qmetatype.h>
#if !defined(Q_MOC_OUTPUT_REVISION)
#error "The header file 'moc.cpp' doesn't include <QObject>."
#elif Q_MOC_OUTPUT_REVISION != 67
#error "This file was generated using the moc from 5.12.0. It"
#error "cannot be used with the include files from this version of Qt."
#error "(The moc has changed too much.)"
#endif

QT_BEGIN_MOC_NAMESPACE
QT_WARNING_PUSH
QT_WARNING_DISABLE_DEPRECATED
struct qt_meta_stringdata_ApproveExportCtx77da62_t {
    QByteArrayData data[18];
    char stringdata0[198];
};
#define QT_MOC_LITERAL(idx, ofs, len) \
    Q_STATIC_BYTE_ARRAY_DATA_HEADER_INITIALIZER_WITH_OFFSET(len, \
    qptrdiff(offsetof(qt_meta_stringdata_ApproveExportCtx77da62_t, stringdata0) + ofs \
        - idx * sizeof(QByteArrayData)) \
    )
static const qt_meta_stringdata_ApproveExportCtx77da62_t qt_meta_stringdata_ApproveExportCtx77da62 = {
    {
QT_MOC_LITERAL(0, 0, 22), // "ApproveExportCtx77da62"
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
QT_MOC_LITERAL(15, 166, 8), // "password"
QT_MOC_LITERAL(16, 175, 14), // "fromSrcChanged"
QT_MOC_LITERAL(17, 190, 7) // "fromSrc"

    },
    "ApproveExportCtx77da62\0clicked\0\0b\0"
    "back\0passwordEdited\0remoteChanged\0"
    "remote\0transportChanged\0transport\0"
    "endpointChanged\0endpoint\0addressChanged\0"
    "address\0passwordChanged\0password\0"
    "fromSrcChanged\0fromSrc"
};
#undef QT_MOC_LITERAL

static const uint qt_meta_data_ApproveExportCtx77da62[] = {

 // content:
       8,       // revision
       0,       // classname
       0,    0, // classinfo
       9,   14, // methods
       6,   84, // properties
       0,    0, // enums/sets
       0,    0, // constructors
       0,       // flags
       9,       // signalCount

 // signals: name, argc, parameters, tag, flags
       1,    1,   59,    2, 0x06 /* Public */,
       4,    0,   62,    2, 0x06 /* Public */,
       5,    1,   63,    2, 0x06 /* Public */,
       6,    1,   66,    2, 0x06 /* Public */,
       8,    1,   69,    2, 0x06 /* Public */,
      10,    1,   72,    2, 0x06 /* Public */,
      12,    1,   75,    2, 0x06 /* Public */,
      14,    1,   78,    2, 0x06 /* Public */,
      16,    1,   81,    2, 0x06 /* Public */,

 // signals: parameters
    QMetaType::Void, QMetaType::Int,    3,
    QMetaType::Void,
    QMetaType::Void, QMetaType::QString,    3,
    QMetaType::Void, QMetaType::QString,    7,
    QMetaType::Void, QMetaType::QString,    9,
    QMetaType::Void, QMetaType::QString,   11,
    QMetaType::Void, QMetaType::QString,   13,
    QMetaType::Void, QMetaType::QString,   15,
    QMetaType::Void, QMetaType::QString,   17,

 // properties: name, type, flags
       7, QMetaType::QString, 0x00495103,
       9, QMetaType::QString, 0x00495103,
      11, QMetaType::QString, 0x00495103,
      13, QMetaType::QString, 0x00495103,
      15, QMetaType::QString, 0x00495103,
      17, QMetaType::QString, 0x00495103,

 // properties: notify_signal_id
       3,
       4,
       5,
       6,
       7,
       8,

       0        // eod
};

void ApproveExportCtx77da62::qt_static_metacall(QObject *_o, QMetaObject::Call _c, int _id, void **_a)
{
    if (_c == QMetaObject::InvokeMetaMethod) {
        ApproveExportCtx77da62 *_t = static_cast<ApproveExportCtx77da62 *>(_o);
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
        case 8: _t->fromSrcChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        default: ;
        }
    } else if (_c == QMetaObject::IndexOfMethod) {
        int *result = reinterpret_cast<int *>(_a[0]);
        {
            using _t = void (ApproveExportCtx77da62::*)(qint32 );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveExportCtx77da62::clicked)) {
                *result = 0;
                return;
            }
        }
        {
            using _t = void (ApproveExportCtx77da62::*)();
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveExportCtx77da62::back)) {
                *result = 1;
                return;
            }
        }
        {
            using _t = void (ApproveExportCtx77da62::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveExportCtx77da62::passwordEdited)) {
                *result = 2;
                return;
            }
        }
        {
            using _t = void (ApproveExportCtx77da62::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveExportCtx77da62::remoteChanged)) {
                *result = 3;
                return;
            }
        }
        {
            using _t = void (ApproveExportCtx77da62::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveExportCtx77da62::transportChanged)) {
                *result = 4;
                return;
            }
        }
        {
            using _t = void (ApproveExportCtx77da62::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveExportCtx77da62::endpointChanged)) {
                *result = 5;
                return;
            }
        }
        {
            using _t = void (ApproveExportCtx77da62::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveExportCtx77da62::addressChanged)) {
                *result = 6;
                return;
            }
        }
        {
            using _t = void (ApproveExportCtx77da62::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveExportCtx77da62::passwordChanged)) {
                *result = 7;
                return;
            }
        }
        {
            using _t = void (ApproveExportCtx77da62::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveExportCtx77da62::fromSrcChanged)) {
                *result = 8;
                return;
            }
        }
    }
#ifndef QT_NO_PROPERTIES
    else if (_c == QMetaObject::ReadProperty) {
        ApproveExportCtx77da62 *_t = static_cast<ApproveExportCtx77da62 *>(_o);
        Q_UNUSED(_t)
        void *_v = _a[0];
        switch (_id) {
        case 0: *reinterpret_cast< QString*>(_v) = _t->remote(); break;
        case 1: *reinterpret_cast< QString*>(_v) = _t->transport(); break;
        case 2: *reinterpret_cast< QString*>(_v) = _t->endpoint(); break;
        case 3: *reinterpret_cast< QString*>(_v) = _t->address(); break;
        case 4: *reinterpret_cast< QString*>(_v) = _t->password(); break;
        case 5: *reinterpret_cast< QString*>(_v) = _t->fromSrc(); break;
        default: break;
        }
    } else if (_c == QMetaObject::WriteProperty) {
        ApproveExportCtx77da62 *_t = static_cast<ApproveExportCtx77da62 *>(_o);
        Q_UNUSED(_t)
        void *_v = _a[0];
        switch (_id) {
        case 0: _t->setRemote(*reinterpret_cast< QString*>(_v)); break;
        case 1: _t->setTransport(*reinterpret_cast< QString*>(_v)); break;
        case 2: _t->setEndpoint(*reinterpret_cast< QString*>(_v)); break;
        case 3: _t->setAddress(*reinterpret_cast< QString*>(_v)); break;
        case 4: _t->setPassword(*reinterpret_cast< QString*>(_v)); break;
        case 5: _t->setFromSrc(*reinterpret_cast< QString*>(_v)); break;
        default: break;
        }
    } else if (_c == QMetaObject::ResetProperty) {
    }
#endif // QT_NO_PROPERTIES
}

QT_INIT_METAOBJECT const QMetaObject ApproveExportCtx77da62::staticMetaObject = { {
    &QObject::staticMetaObject,
    qt_meta_stringdata_ApproveExportCtx77da62.data,
    qt_meta_data_ApproveExportCtx77da62,
    qt_static_metacall,
    nullptr,
    nullptr
} };


const QMetaObject *ApproveExportCtx77da62::metaObject() const
{
    return QObject::d_ptr->metaObject ? QObject::d_ptr->dynamicMetaObject() : &staticMetaObject;
}

void *ApproveExportCtx77da62::qt_metacast(const char *_clname)
{
    if (!_clname) return nullptr;
    if (!strcmp(_clname, qt_meta_stringdata_ApproveExportCtx77da62.stringdata0))
        return static_cast<void*>(this);
    return QObject::qt_metacast(_clname);
}

int ApproveExportCtx77da62::qt_metacall(QMetaObject::Call _c, int _id, void **_a)
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
void ApproveExportCtx77da62::clicked(qint32 _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 0, _a);
}

// SIGNAL 1
void ApproveExportCtx77da62::back()
{
    QMetaObject::activate(this, &staticMetaObject, 1, nullptr);
}

// SIGNAL 2
void ApproveExportCtx77da62::passwordEdited(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 2, _a);
}

// SIGNAL 3
void ApproveExportCtx77da62::remoteChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 3, _a);
}

// SIGNAL 4
void ApproveExportCtx77da62::transportChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 4, _a);
}

// SIGNAL 5
void ApproveExportCtx77da62::endpointChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 5, _a);
}

// SIGNAL 6
void ApproveExportCtx77da62::addressChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 6, _a);
}

// SIGNAL 7
void ApproveExportCtx77da62::passwordChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 7, _a);
}

// SIGNAL 8
void ApproveExportCtx77da62::fromSrcChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 8, _a);
}
struct qt_meta_stringdata_CustomListModel77da62_t {
    QByteArrayData data[6];
    char stringdata0[50];
};
#define QT_MOC_LITERAL(idx, ofs, len) \
    Q_STATIC_BYTE_ARRAY_DATA_HEADER_INITIALIZER_WITH_OFFSET(len, \
    qptrdiff(offsetof(qt_meta_stringdata_CustomListModel77da62_t, stringdata0) + ofs \
        - idx * sizeof(QByteArrayData)) \
    )
static const qt_meta_stringdata_CustomListModel77da62_t qt_meta_stringdata_CustomListModel77da62 = {
    {
QT_MOC_LITERAL(0, 0, 21), // "CustomListModel77da62"
QT_MOC_LITERAL(1, 22, 5), // "clear"
QT_MOC_LITERAL(2, 28, 0), // ""
QT_MOC_LITERAL(3, 29, 3), // "add"
QT_MOC_LITERAL(4, 33, 8), // "quintptr"
QT_MOC_LITERAL(5, 42, 7) // "account"

    },
    "CustomListModel77da62\0clear\0\0add\0"
    "quintptr\0account"
};
#undef QT_MOC_LITERAL

static const uint qt_meta_data_CustomListModel77da62[] = {

 // content:
       8,       // revision
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

void CustomListModel77da62::qt_static_metacall(QObject *_o, QMetaObject::Call _c, int _id, void **_a)
{
    if (_c == QMetaObject::InvokeMetaMethod) {
        CustomListModel77da62 *_t = static_cast<CustomListModel77da62 *>(_o);
        Q_UNUSED(_t)
        switch (_id) {
        case 0: _t->clear(); break;
        case 1: _t->add((*reinterpret_cast< quintptr(*)>(_a[1]))); break;
        default: ;
        }
    } else if (_c == QMetaObject::IndexOfMethod) {
        int *result = reinterpret_cast<int *>(_a[0]);
        {
            using _t = void (CustomListModel77da62::*)();
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&CustomListModel77da62::clear)) {
                *result = 0;
                return;
            }
        }
        {
            using _t = void (CustomListModel77da62::*)(quintptr );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&CustomListModel77da62::add)) {
                *result = 1;
                return;
            }
        }
    }
}

QT_INIT_METAOBJECT const QMetaObject CustomListModel77da62::staticMetaObject = { {
    &QAbstractListModel::staticMetaObject,
    qt_meta_stringdata_CustomListModel77da62.data,
    qt_meta_data_CustomListModel77da62,
    qt_static_metacall,
    nullptr,
    nullptr
} };


const QMetaObject *CustomListModel77da62::metaObject() const
{
    return QObject::d_ptr->metaObject ? QObject::d_ptr->dynamicMetaObject() : &staticMetaObject;
}

void *CustomListModel77da62::qt_metacast(const char *_clname)
{
    if (!_clname) return nullptr;
    if (!strcmp(_clname, qt_meta_stringdata_CustomListModel77da62.stringdata0))
        return static_cast<void*>(this);
    return QAbstractListModel::qt_metacast(_clname);
}

int CustomListModel77da62::qt_metacall(QMetaObject::Call _c, int _id, void **_a)
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
void CustomListModel77da62::clear()
{
    QMetaObject::activate(this, &staticMetaObject, 0, nullptr);
}

// SIGNAL 1
void CustomListModel77da62::add(quintptr _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 1, _a);
}
struct qt_meta_stringdata_LoginCtx77da62_t {
    QByteArrayData data[17];
    char stringdata0[179];
};
#define QT_MOC_LITERAL(idx, ofs, len) \
    Q_STATIC_BYTE_ARRAY_DATA_HEADER_INITIALIZER_WITH_OFFSET(len, \
    qptrdiff(offsetof(qt_meta_stringdata_LoginCtx77da62_t, stringdata0) + ofs \
        - idx * sizeof(QByteArrayData)) \
    )
static const qt_meta_stringdata_LoginCtx77da62_t qt_meta_stringdata_LoginCtx77da62 = {
    {
QT_MOC_LITERAL(0, 0, 14), // "LoginCtx77da62"
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
    "LoginCtx77da62\0clicked\0\0edited\0b\0"
    "remoteChanged\0remote\0transportChanged\0"
    "transport\0endpointChanged\0endpoint\0"
    "gopathChanged\0gopath\0binaryHashChanged\0"
    "binaryHash\0isValidChanged\0isValid"
};
#undef QT_MOC_LITERAL

static const uint qt_meta_data_LoginCtx77da62[] = {

 // content:
       8,       // revision
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

void LoginCtx77da62::qt_static_metacall(QObject *_o, QMetaObject::Call _c, int _id, void **_a)
{
    if (_c == QMetaObject::InvokeMetaMethod) {
        LoginCtx77da62 *_t = static_cast<LoginCtx77da62 *>(_o);
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
            using _t = void (LoginCtx77da62::*)();
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&LoginCtx77da62::clicked)) {
                *result = 0;
                return;
            }
        }
        {
            using _t = void (LoginCtx77da62::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&LoginCtx77da62::edited)) {
                *result = 1;
                return;
            }
        }
        {
            using _t = void (LoginCtx77da62::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&LoginCtx77da62::remoteChanged)) {
                *result = 2;
                return;
            }
        }
        {
            using _t = void (LoginCtx77da62::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&LoginCtx77da62::transportChanged)) {
                *result = 3;
                return;
            }
        }
        {
            using _t = void (LoginCtx77da62::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&LoginCtx77da62::endpointChanged)) {
                *result = 4;
                return;
            }
        }
        {
            using _t = void (LoginCtx77da62::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&LoginCtx77da62::gopathChanged)) {
                *result = 5;
                return;
            }
        }
        {
            using _t = void (LoginCtx77da62::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&LoginCtx77da62::binaryHashChanged)) {
                *result = 6;
                return;
            }
        }
        {
            using _t = void (LoginCtx77da62::*)(bool );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&LoginCtx77da62::isValidChanged)) {
                *result = 7;
                return;
            }
        }
    }
#ifndef QT_NO_PROPERTIES
    else if (_c == QMetaObject::ReadProperty) {
        LoginCtx77da62 *_t = static_cast<LoginCtx77da62 *>(_o);
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
        LoginCtx77da62 *_t = static_cast<LoginCtx77da62 *>(_o);
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

QT_INIT_METAOBJECT const QMetaObject LoginCtx77da62::staticMetaObject = { {
    &QObject::staticMetaObject,
    qt_meta_stringdata_LoginCtx77da62.data,
    qt_meta_data_LoginCtx77da62,
    qt_static_metacall,
    nullptr,
    nullptr
} };


const QMetaObject *LoginCtx77da62::metaObject() const
{
    return QObject::d_ptr->metaObject ? QObject::d_ptr->dynamicMetaObject() : &staticMetaObject;
}

void *LoginCtx77da62::qt_metacast(const char *_clname)
{
    if (!_clname) return nullptr;
    if (!strcmp(_clname, qt_meta_stringdata_LoginCtx77da62.stringdata0))
        return static_cast<void*>(this);
    return QObject::qt_metacast(_clname);
}

int LoginCtx77da62::qt_metacall(QMetaObject::Call _c, int _id, void **_a)
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
void LoginCtx77da62::clicked()
{
    QMetaObject::activate(this, &staticMetaObject, 0, nullptr);
}

// SIGNAL 1
void LoginCtx77da62::edited(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 1, _a);
}

// SIGNAL 2
void LoginCtx77da62::remoteChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 2, _a);
}

// SIGNAL 3
void LoginCtx77da62::transportChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 3, _a);
}

// SIGNAL 4
void LoginCtx77da62::endpointChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 4, _a);
}

// SIGNAL 5
void LoginCtx77da62::gopathChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 5, _a);
}

// SIGNAL 6
void LoginCtx77da62::binaryHashChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 6, _a);
}

// SIGNAL 7
void LoginCtx77da62::isValidChanged(bool _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 7, _a);
}
struct qt_meta_stringdata_TxListAccountsModel77da62_t {
    QByteArrayData data[4];
    char stringdata0[34];
};
#define QT_MOC_LITERAL(idx, ofs, len) \
    Q_STATIC_BYTE_ARRAY_DATA_HEADER_INITIALIZER_WITH_OFFSET(len, \
    qptrdiff(offsetof(qt_meta_stringdata_TxListAccountsModel77da62_t, stringdata0) + ofs \
        - idx * sizeof(QByteArrayData)) \
    )
static const qt_meta_stringdata_TxListAccountsModel77da62_t qt_meta_stringdata_TxListAccountsModel77da62 = {
    {
QT_MOC_LITERAL(0, 0, 25), // "TxListAccountsModel77da62"
QT_MOC_LITERAL(1, 26, 3), // "add"
QT_MOC_LITERAL(2, 30, 0), // ""
QT_MOC_LITERAL(3, 31, 2) // "tx"

    },
    "TxListAccountsModel77da62\0add\0\0tx"
};
#undef QT_MOC_LITERAL

static const uint qt_meta_data_TxListAccountsModel77da62[] = {

 // content:
       8,       // revision
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
    QMetaType::Void, QMetaType::QString,    3,

       0        // eod
};

void TxListAccountsModel77da62::qt_static_metacall(QObject *_o, QMetaObject::Call _c, int _id, void **_a)
{
    if (_c == QMetaObject::InvokeMetaMethod) {
        TxListAccountsModel77da62 *_t = static_cast<TxListAccountsModel77da62 *>(_o);
        Q_UNUSED(_t)
        switch (_id) {
        case 0: _t->add((*reinterpret_cast< QString(*)>(_a[1]))); break;
        default: ;
        }
    } else if (_c == QMetaObject::IndexOfMethod) {
        int *result = reinterpret_cast<int *>(_a[0]);
        {
            using _t = void (TxListAccountsModel77da62::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&TxListAccountsModel77da62::add)) {
                *result = 0;
                return;
            }
        }
    }
}

QT_INIT_METAOBJECT const QMetaObject TxListAccountsModel77da62::staticMetaObject = { {
    &QAbstractListModel::staticMetaObject,
    qt_meta_stringdata_TxListAccountsModel77da62.data,
    qt_meta_data_TxListAccountsModel77da62,
    qt_static_metacall,
    nullptr,
    nullptr
} };


const QMetaObject *TxListAccountsModel77da62::metaObject() const
{
    return QObject::d_ptr->metaObject ? QObject::d_ptr->dynamicMetaObject() : &staticMetaObject;
}

void *TxListAccountsModel77da62::qt_metacast(const char *_clname)
{
    if (!_clname) return nullptr;
    if (!strcmp(_clname, qt_meta_stringdata_TxListAccountsModel77da62.stringdata0))
        return static_cast<void*>(this);
    return QAbstractListModel::qt_metacast(_clname);
}

int TxListAccountsModel77da62::qt_metacall(QMetaObject::Call _c, int _id, void **_a)
{
    _id = QAbstractListModel::qt_metacall(_c, _id, _a);
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
void TxListAccountsModel77da62::add(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 0, _a);
}
struct qt_meta_stringdata_ApproveImportCtx77da62_t {
    QByteArrayData data[20];
    char stringdata0[262];
};
#define QT_MOC_LITERAL(idx, ofs, len) \
    Q_STATIC_BYTE_ARRAY_DATA_HEADER_INITIALIZER_WITH_OFFSET(len, \
    qptrdiff(offsetof(qt_meta_stringdata_ApproveImportCtx77da62_t, stringdata0) + ofs \
        - idx * sizeof(QByteArrayData)) \
    )
static const qt_meta_stringdata_ApproveImportCtx77da62_t qt_meta_stringdata_ApproveImportCtx77da62 = {
    {
QT_MOC_LITERAL(0, 0, 22), // "ApproveImportCtx77da62"
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
    "ApproveImportCtx77da62\0clicked\0\0b\0"
    "back\0passwordEdited\0confirmPasswordEdited\0"
    "oldPasswordEdited\0remoteChanged\0remote\0"
    "transportChanged\0transport\0endpointChanged\0"
    "endpoint\0oldPasswordChanged\0oldPassword\0"
    "passwordChanged\0password\0"
    "confirmPasswordChanged\0confirmPassword"
};
#undef QT_MOC_LITERAL

static const uint qt_meta_data_ApproveImportCtx77da62[] = {

 // content:
       8,       // revision
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

void ApproveImportCtx77da62::qt_static_metacall(QObject *_o, QMetaObject::Call _c, int _id, void **_a)
{
    if (_c == QMetaObject::InvokeMetaMethod) {
        ApproveImportCtx77da62 *_t = static_cast<ApproveImportCtx77da62 *>(_o);
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
            using _t = void (ApproveImportCtx77da62::*)(qint32 );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveImportCtx77da62::clicked)) {
                *result = 0;
                return;
            }
        }
        {
            using _t = void (ApproveImportCtx77da62::*)();
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveImportCtx77da62::back)) {
                *result = 1;
                return;
            }
        }
        {
            using _t = void (ApproveImportCtx77da62::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveImportCtx77da62::passwordEdited)) {
                *result = 2;
                return;
            }
        }
        {
            using _t = void (ApproveImportCtx77da62::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveImportCtx77da62::confirmPasswordEdited)) {
                *result = 3;
                return;
            }
        }
        {
            using _t = void (ApproveImportCtx77da62::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveImportCtx77da62::oldPasswordEdited)) {
                *result = 4;
                return;
            }
        }
        {
            using _t = void (ApproveImportCtx77da62::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveImportCtx77da62::remoteChanged)) {
                *result = 5;
                return;
            }
        }
        {
            using _t = void (ApproveImportCtx77da62::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveImportCtx77da62::transportChanged)) {
                *result = 6;
                return;
            }
        }
        {
            using _t = void (ApproveImportCtx77da62::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveImportCtx77da62::endpointChanged)) {
                *result = 7;
                return;
            }
        }
        {
            using _t = void (ApproveImportCtx77da62::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveImportCtx77da62::oldPasswordChanged)) {
                *result = 8;
                return;
            }
        }
        {
            using _t = void (ApproveImportCtx77da62::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveImportCtx77da62::passwordChanged)) {
                *result = 9;
                return;
            }
        }
        {
            using _t = void (ApproveImportCtx77da62::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveImportCtx77da62::confirmPasswordChanged)) {
                *result = 10;
                return;
            }
        }
    }
#ifndef QT_NO_PROPERTIES
    else if (_c == QMetaObject::ReadProperty) {
        ApproveImportCtx77da62 *_t = static_cast<ApproveImportCtx77da62 *>(_o);
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
        ApproveImportCtx77da62 *_t = static_cast<ApproveImportCtx77da62 *>(_o);
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

QT_INIT_METAOBJECT const QMetaObject ApproveImportCtx77da62::staticMetaObject = { {
    &QObject::staticMetaObject,
    qt_meta_stringdata_ApproveImportCtx77da62.data,
    qt_meta_data_ApproveImportCtx77da62,
    qt_static_metacall,
    nullptr,
    nullptr
} };


const QMetaObject *ApproveImportCtx77da62::metaObject() const
{
    return QObject::d_ptr->metaObject ? QObject::d_ptr->dynamicMetaObject() : &staticMetaObject;
}

void *ApproveImportCtx77da62::qt_metacast(const char *_clname)
{
    if (!_clname) return nullptr;
    if (!strcmp(_clname, qt_meta_stringdata_ApproveImportCtx77da62.stringdata0))
        return static_cast<void*>(this);
    return QObject::qt_metacast(_clname);
}

int ApproveImportCtx77da62::qt_metacall(QMetaObject::Call _c, int _id, void **_a)
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
void ApproveImportCtx77da62::clicked(qint32 _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 0, _a);
}

// SIGNAL 1
void ApproveImportCtx77da62::back()
{
    QMetaObject::activate(this, &staticMetaObject, 1, nullptr);
}

// SIGNAL 2
void ApproveImportCtx77da62::passwordEdited(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 2, _a);
}

// SIGNAL 3
void ApproveImportCtx77da62::confirmPasswordEdited(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 3, _a);
}

// SIGNAL 4
void ApproveImportCtx77da62::oldPasswordEdited(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 4, _a);
}

// SIGNAL 5
void ApproveImportCtx77da62::remoteChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 5, _a);
}

// SIGNAL 6
void ApproveImportCtx77da62::transportChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 6, _a);
}

// SIGNAL 7
void ApproveImportCtx77da62::endpointChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 7, _a);
}

// SIGNAL 8
void ApproveImportCtx77da62::oldPasswordChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 8, _a);
}

// SIGNAL 9
void ApproveImportCtx77da62::passwordChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 9, _a);
}

// SIGNAL 10
void ApproveImportCtx77da62::confirmPasswordChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 10, _a);
}
struct qt_meta_stringdata_ApproveListingCtx77da62_t {
    QByteArrayData data[22];
    char stringdata0[223];
};
#define QT_MOC_LITERAL(idx, ofs, len) \
    Q_STATIC_BYTE_ARRAY_DATA_HEADER_INITIALIZER_WITH_OFFSET(len, \
    qptrdiff(offsetof(qt_meta_stringdata_ApproveListingCtx77da62_t, stringdata0) + ofs \
        - idx * sizeof(QByteArrayData)) \
    )
static const qt_meta_stringdata_ApproveListingCtx77da62_t qt_meta_stringdata_ApproveListingCtx77da62 = {
    {
QT_MOC_LITERAL(0, 0, 23), // "ApproveListingCtx77da62"
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
    "ApproveListingCtx77da62\0clicked\0\0b\0"
    "back\0onCheckStateChanged\0i\0checked\0"
    "remoteChanged\0remote\0transportChanged\0"
    "transport\0endpointChanged\0endpoint\0"
    "fromChanged\0from\0messageChanged\0message\0"
    "rawDataChanged\0rawData\0hashChanged\0"
    "hash"
};
#undef QT_MOC_LITERAL

static const uint qt_meta_data_ApproveListingCtx77da62[] = {

 // content:
       8,       // revision
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

void ApproveListingCtx77da62::qt_static_metacall(QObject *_o, QMetaObject::Call _c, int _id, void **_a)
{
    if (_c == QMetaObject::InvokeMetaMethod) {
        ApproveListingCtx77da62 *_t = static_cast<ApproveListingCtx77da62 *>(_o);
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
            using _t = void (ApproveListingCtx77da62::*)(qint32 );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveListingCtx77da62::clicked)) {
                *result = 0;
                return;
            }
        }
        {
            using _t = void (ApproveListingCtx77da62::*)();
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveListingCtx77da62::back)) {
                *result = 1;
                return;
            }
        }
        {
            using _t = void (ApproveListingCtx77da62::*)(qint32 , bool );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveListingCtx77da62::onCheckStateChanged)) {
                *result = 2;
                return;
            }
        }
        {
            using _t = void (ApproveListingCtx77da62::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveListingCtx77da62::remoteChanged)) {
                *result = 3;
                return;
            }
        }
        {
            using _t = void (ApproveListingCtx77da62::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveListingCtx77da62::transportChanged)) {
                *result = 4;
                return;
            }
        }
        {
            using _t = void (ApproveListingCtx77da62::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveListingCtx77da62::endpointChanged)) {
                *result = 5;
                return;
            }
        }
        {
            using _t = void (ApproveListingCtx77da62::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveListingCtx77da62::fromChanged)) {
                *result = 6;
                return;
            }
        }
        {
            using _t = void (ApproveListingCtx77da62::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveListingCtx77da62::messageChanged)) {
                *result = 7;
                return;
            }
        }
        {
            using _t = void (ApproveListingCtx77da62::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveListingCtx77da62::rawDataChanged)) {
                *result = 8;
                return;
            }
        }
        {
            using _t = void (ApproveListingCtx77da62::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveListingCtx77da62::hashChanged)) {
                *result = 9;
                return;
            }
        }
    }
#ifndef QT_NO_PROPERTIES
    else if (_c == QMetaObject::ReadProperty) {
        ApproveListingCtx77da62 *_t = static_cast<ApproveListingCtx77da62 *>(_o);
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
        ApproveListingCtx77da62 *_t = static_cast<ApproveListingCtx77da62 *>(_o);
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

QT_INIT_METAOBJECT const QMetaObject ApproveListingCtx77da62::staticMetaObject = { {
    &QObject::staticMetaObject,
    qt_meta_stringdata_ApproveListingCtx77da62.data,
    qt_meta_data_ApproveListingCtx77da62,
    qt_static_metacall,
    nullptr,
    nullptr
} };


const QMetaObject *ApproveListingCtx77da62::metaObject() const
{
    return QObject::d_ptr->metaObject ? QObject::d_ptr->dynamicMetaObject() : &staticMetaObject;
}

void *ApproveListingCtx77da62::qt_metacast(const char *_clname)
{
    if (!_clname) return nullptr;
    if (!strcmp(_clname, qt_meta_stringdata_ApproveListingCtx77da62.stringdata0))
        return static_cast<void*>(this);
    return QObject::qt_metacast(_clname);
}

int ApproveListingCtx77da62::qt_metacall(QMetaObject::Call _c, int _id, void **_a)
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
void ApproveListingCtx77da62::clicked(qint32 _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 0, _a);
}

// SIGNAL 1
void ApproveListingCtx77da62::back()
{
    QMetaObject::activate(this, &staticMetaObject, 1, nullptr);
}

// SIGNAL 2
void ApproveListingCtx77da62::onCheckStateChanged(qint32 _t1, bool _t2)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)), const_cast<void*>(reinterpret_cast<const void*>(&_t2)) };
    QMetaObject::activate(this, &staticMetaObject, 2, _a);
}

// SIGNAL 3
void ApproveListingCtx77da62::remoteChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 3, _a);
}

// SIGNAL 4
void ApproveListingCtx77da62::transportChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 4, _a);
}

// SIGNAL 5
void ApproveListingCtx77da62::endpointChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 5, _a);
}

// SIGNAL 6
void ApproveListingCtx77da62::fromChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 6, _a);
}

// SIGNAL 7
void ApproveListingCtx77da62::messageChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 7, _a);
}

// SIGNAL 8
void ApproveListingCtx77da62::rawDataChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 8, _a);
}

// SIGNAL 9
void ApproveListingCtx77da62::hashChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 9, _a);
}
struct qt_meta_stringdata_ApproveNewAccountCtx77da62_t {
    QByteArrayData data[17];
    char stringdata0[217];
};
#define QT_MOC_LITERAL(idx, ofs, len) \
    Q_STATIC_BYTE_ARRAY_DATA_HEADER_INITIALIZER_WITH_OFFSET(len, \
    qptrdiff(offsetof(qt_meta_stringdata_ApproveNewAccountCtx77da62_t, stringdata0) + ofs \
        - idx * sizeof(QByteArrayData)) \
    )
static const qt_meta_stringdata_ApproveNewAccountCtx77da62_t qt_meta_stringdata_ApproveNewAccountCtx77da62 = {
    {
QT_MOC_LITERAL(0, 0, 26), // "ApproveNewAccountCtx77da62"
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
    "ApproveNewAccountCtx77da62\0clicked\0\0"
    "b\0back\0passwordEdited\0confirmPasswordEdited\0"
    "remoteChanged\0remote\0transportChanged\0"
    "transport\0endpointChanged\0endpoint\0"
    "passwordChanged\0password\0"
    "confirmPasswordChanged\0confirmPassword"
};
#undef QT_MOC_LITERAL

static const uint qt_meta_data_ApproveNewAccountCtx77da62[] = {

 // content:
       8,       // revision
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

void ApproveNewAccountCtx77da62::qt_static_metacall(QObject *_o, QMetaObject::Call _c, int _id, void **_a)
{
    if (_c == QMetaObject::InvokeMetaMethod) {
        ApproveNewAccountCtx77da62 *_t = static_cast<ApproveNewAccountCtx77da62 *>(_o);
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
            using _t = void (ApproveNewAccountCtx77da62::*)(qint32 );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveNewAccountCtx77da62::clicked)) {
                *result = 0;
                return;
            }
        }
        {
            using _t = void (ApproveNewAccountCtx77da62::*)();
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveNewAccountCtx77da62::back)) {
                *result = 1;
                return;
            }
        }
        {
            using _t = void (ApproveNewAccountCtx77da62::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveNewAccountCtx77da62::passwordEdited)) {
                *result = 2;
                return;
            }
        }
        {
            using _t = void (ApproveNewAccountCtx77da62::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveNewAccountCtx77da62::confirmPasswordEdited)) {
                *result = 3;
                return;
            }
        }
        {
            using _t = void (ApproveNewAccountCtx77da62::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveNewAccountCtx77da62::remoteChanged)) {
                *result = 4;
                return;
            }
        }
        {
            using _t = void (ApproveNewAccountCtx77da62::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveNewAccountCtx77da62::transportChanged)) {
                *result = 5;
                return;
            }
        }
        {
            using _t = void (ApproveNewAccountCtx77da62::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveNewAccountCtx77da62::endpointChanged)) {
                *result = 6;
                return;
            }
        }
        {
            using _t = void (ApproveNewAccountCtx77da62::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveNewAccountCtx77da62::passwordChanged)) {
                *result = 7;
                return;
            }
        }
        {
            using _t = void (ApproveNewAccountCtx77da62::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveNewAccountCtx77da62::confirmPasswordChanged)) {
                *result = 8;
                return;
            }
        }
    }
#ifndef QT_NO_PROPERTIES
    else if (_c == QMetaObject::ReadProperty) {
        ApproveNewAccountCtx77da62 *_t = static_cast<ApproveNewAccountCtx77da62 *>(_o);
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
        ApproveNewAccountCtx77da62 *_t = static_cast<ApproveNewAccountCtx77da62 *>(_o);
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

QT_INIT_METAOBJECT const QMetaObject ApproveNewAccountCtx77da62::staticMetaObject = { {
    &QObject::staticMetaObject,
    qt_meta_stringdata_ApproveNewAccountCtx77da62.data,
    qt_meta_data_ApproveNewAccountCtx77da62,
    qt_static_metacall,
    nullptr,
    nullptr
} };


const QMetaObject *ApproveNewAccountCtx77da62::metaObject() const
{
    return QObject::d_ptr->metaObject ? QObject::d_ptr->dynamicMetaObject() : &staticMetaObject;
}

void *ApproveNewAccountCtx77da62::qt_metacast(const char *_clname)
{
    if (!_clname) return nullptr;
    if (!strcmp(_clname, qt_meta_stringdata_ApproveNewAccountCtx77da62.stringdata0))
        return static_cast<void*>(this);
    return QObject::qt_metacast(_clname);
}

int ApproveNewAccountCtx77da62::qt_metacall(QMetaObject::Call _c, int _id, void **_a)
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
void ApproveNewAccountCtx77da62::clicked(qint32 _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 0, _a);
}

// SIGNAL 1
void ApproveNewAccountCtx77da62::back()
{
    QMetaObject::activate(this, &staticMetaObject, 1, nullptr);
}

// SIGNAL 2
void ApproveNewAccountCtx77da62::passwordEdited(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 2, _a);
}

// SIGNAL 3
void ApproveNewAccountCtx77da62::confirmPasswordEdited(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 3, _a);
}

// SIGNAL 4
void ApproveNewAccountCtx77da62::remoteChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 4, _a);
}

// SIGNAL 5
void ApproveNewAccountCtx77da62::transportChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 5, _a);
}

// SIGNAL 6
void ApproveNewAccountCtx77da62::endpointChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 6, _a);
}

// SIGNAL 7
void ApproveNewAccountCtx77da62::passwordChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 7, _a);
}

// SIGNAL 8
void ApproveNewAccountCtx77da62::confirmPasswordChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 8, _a);
}
struct qt_meta_stringdata_ApproveSignDataCtx77da62_t {
    QByteArrayData data[25];
    char stringdata0[255];
};
#define QT_MOC_LITERAL(idx, ofs, len) \
    Q_STATIC_BYTE_ARRAY_DATA_HEADER_INITIALIZER_WITH_OFFSET(len, \
    qptrdiff(offsetof(qt_meta_stringdata_ApproveSignDataCtx77da62_t, stringdata0) + ofs \
        - idx * sizeof(QByteArrayData)) \
    )
static const qt_meta_stringdata_ApproveSignDataCtx77da62_t qt_meta_stringdata_ApproveSignDataCtx77da62 = {
    {
QT_MOC_LITERAL(0, 0, 24), // "ApproveSignDataCtx77da62"
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
QT_MOC_LITERAL(22, 223, 8), // "password"
QT_MOC_LITERAL(23, 232, 14), // "fromSrcChanged"
QT_MOC_LITERAL(24, 247, 7) // "fromSrc"

    },
    "ApproveSignDataCtx77da62\0clicked\0\0b\0"
    "back\0edited\0value\0remoteChanged\0remote\0"
    "transportChanged\0transport\0endpointChanged\0"
    "endpoint\0fromChanged\0from\0messageChanged\0"
    "message\0rawDataChanged\0rawData\0"
    "hashChanged\0hash\0passwordChanged\0"
    "password\0fromSrcChanged\0fromSrc"
};
#undef QT_MOC_LITERAL

static const uint qt_meta_data_ApproveSignDataCtx77da62[] = {

 // content:
       8,       // revision
       0,       // classname
       0,    0, // classinfo
      12,   14, // methods
       9,  110, // properties
       0,    0, // enums/sets
       0,    0, // constructors
       0,       // flags
      12,       // signalCount

 // signals: name, argc, parameters, tag, flags
       1,    1,   74,    2, 0x06 /* Public */,
       4,    0,   77,    2, 0x06 /* Public */,
       5,    2,   78,    2, 0x06 /* Public */,
       7,    1,   83,    2, 0x06 /* Public */,
       9,    1,   86,    2, 0x06 /* Public */,
      11,    1,   89,    2, 0x06 /* Public */,
      13,    1,   92,    2, 0x06 /* Public */,
      15,    1,   95,    2, 0x06 /* Public */,
      17,    1,   98,    2, 0x06 /* Public */,
      19,    1,  101,    2, 0x06 /* Public */,
      21,    1,  104,    2, 0x06 /* Public */,
      23,    1,  107,    2, 0x06 /* Public */,

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
    QMetaType::Void, QMetaType::QString,   24,

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

       0        // eod
};

void ApproveSignDataCtx77da62::qt_static_metacall(QObject *_o, QMetaObject::Call _c, int _id, void **_a)
{
    if (_c == QMetaObject::InvokeMetaMethod) {
        ApproveSignDataCtx77da62 *_t = static_cast<ApproveSignDataCtx77da62 *>(_o);
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
        case 11: _t->fromSrcChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        default: ;
        }
    } else if (_c == QMetaObject::IndexOfMethod) {
        int *result = reinterpret_cast<int *>(_a[0]);
        {
            using _t = void (ApproveSignDataCtx77da62::*)(qint32 );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveSignDataCtx77da62::clicked)) {
                *result = 0;
                return;
            }
        }
        {
            using _t = void (ApproveSignDataCtx77da62::*)();
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveSignDataCtx77da62::back)) {
                *result = 1;
                return;
            }
        }
        {
            using _t = void (ApproveSignDataCtx77da62::*)(QString , QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveSignDataCtx77da62::edited)) {
                *result = 2;
                return;
            }
        }
        {
            using _t = void (ApproveSignDataCtx77da62::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveSignDataCtx77da62::remoteChanged)) {
                *result = 3;
                return;
            }
        }
        {
            using _t = void (ApproveSignDataCtx77da62::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveSignDataCtx77da62::transportChanged)) {
                *result = 4;
                return;
            }
        }
        {
            using _t = void (ApproveSignDataCtx77da62::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveSignDataCtx77da62::endpointChanged)) {
                *result = 5;
                return;
            }
        }
        {
            using _t = void (ApproveSignDataCtx77da62::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveSignDataCtx77da62::fromChanged)) {
                *result = 6;
                return;
            }
        }
        {
            using _t = void (ApproveSignDataCtx77da62::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveSignDataCtx77da62::messageChanged)) {
                *result = 7;
                return;
            }
        }
        {
            using _t = void (ApproveSignDataCtx77da62::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveSignDataCtx77da62::rawDataChanged)) {
                *result = 8;
                return;
            }
        }
        {
            using _t = void (ApproveSignDataCtx77da62::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveSignDataCtx77da62::hashChanged)) {
                *result = 9;
                return;
            }
        }
        {
            using _t = void (ApproveSignDataCtx77da62::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveSignDataCtx77da62::passwordChanged)) {
                *result = 10;
                return;
            }
        }
        {
            using _t = void (ApproveSignDataCtx77da62::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveSignDataCtx77da62::fromSrcChanged)) {
                *result = 11;
                return;
            }
        }
    }
#ifndef QT_NO_PROPERTIES
    else if (_c == QMetaObject::ReadProperty) {
        ApproveSignDataCtx77da62 *_t = static_cast<ApproveSignDataCtx77da62 *>(_o);
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
        case 8: *reinterpret_cast< QString*>(_v) = _t->fromSrc(); break;
        default: break;
        }
    } else if (_c == QMetaObject::WriteProperty) {
        ApproveSignDataCtx77da62 *_t = static_cast<ApproveSignDataCtx77da62 *>(_o);
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
        case 8: _t->setFromSrc(*reinterpret_cast< QString*>(_v)); break;
        default: break;
        }
    } else if (_c == QMetaObject::ResetProperty) {
    }
#endif // QT_NO_PROPERTIES
}

QT_INIT_METAOBJECT const QMetaObject ApproveSignDataCtx77da62::staticMetaObject = { {
    &QObject::staticMetaObject,
    qt_meta_stringdata_ApproveSignDataCtx77da62.data,
    qt_meta_data_ApproveSignDataCtx77da62,
    qt_static_metacall,
    nullptr,
    nullptr
} };


const QMetaObject *ApproveSignDataCtx77da62::metaObject() const
{
    return QObject::d_ptr->metaObject ? QObject::d_ptr->dynamicMetaObject() : &staticMetaObject;
}

void *ApproveSignDataCtx77da62::qt_metacast(const char *_clname)
{
    if (!_clname) return nullptr;
    if (!strcmp(_clname, qt_meta_stringdata_ApproveSignDataCtx77da62.stringdata0))
        return static_cast<void*>(this);
    return QObject::qt_metacast(_clname);
}

int ApproveSignDataCtx77da62::qt_metacall(QMetaObject::Call _c, int _id, void **_a)
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
        _id -= 9;
    } else if (_c == QMetaObject::QueryPropertyDesignable) {
        _id -= 9;
    } else if (_c == QMetaObject::QueryPropertyScriptable) {
        _id -= 9;
    } else if (_c == QMetaObject::QueryPropertyStored) {
        _id -= 9;
    } else if (_c == QMetaObject::QueryPropertyEditable) {
        _id -= 9;
    } else if (_c == QMetaObject::QueryPropertyUser) {
        _id -= 9;
    }
#endif // QT_NO_PROPERTIES
    return _id;
}

// SIGNAL 0
void ApproveSignDataCtx77da62::clicked(qint32 _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 0, _a);
}

// SIGNAL 1
void ApproveSignDataCtx77da62::back()
{
    QMetaObject::activate(this, &staticMetaObject, 1, nullptr);
}

// SIGNAL 2
void ApproveSignDataCtx77da62::edited(QString _t1, QString _t2)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)), const_cast<void*>(reinterpret_cast<const void*>(&_t2)) };
    QMetaObject::activate(this, &staticMetaObject, 2, _a);
}

// SIGNAL 3
void ApproveSignDataCtx77da62::remoteChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 3, _a);
}

// SIGNAL 4
void ApproveSignDataCtx77da62::transportChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 4, _a);
}

// SIGNAL 5
void ApproveSignDataCtx77da62::endpointChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 5, _a);
}

// SIGNAL 6
void ApproveSignDataCtx77da62::fromChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 6, _a);
}

// SIGNAL 7
void ApproveSignDataCtx77da62::messageChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 7, _a);
}

// SIGNAL 8
void ApproveSignDataCtx77da62::rawDataChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 8, _a);
}

// SIGNAL 9
void ApproveSignDataCtx77da62::hashChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 9, _a);
}

// SIGNAL 10
void ApproveSignDataCtx77da62::passwordChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 10, _a);
}

// SIGNAL 11
void ApproveSignDataCtx77da62::fromSrcChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 11, _a);
}
struct qt_meta_stringdata_ApproveTxCtx77da62_t {
    QByteArrayData data[51];
    char stringdata0[551];
};
#define QT_MOC_LITERAL(idx, ofs, len) \
    Q_STATIC_BYTE_ARRAY_DATA_HEADER_INITIALIZER_WITH_OFFSET(len, \
    qptrdiff(offsetof(qt_meta_stringdata_ApproveTxCtx77da62_t, stringdata0) + ofs \
        - idx * sizeof(QByteArrayData)) \
    )
static const qt_meta_stringdata_ApproveTxCtx77da62_t qt_meta_stringdata_ApproveTxCtx77da62 = {
    {
QT_MOC_LITERAL(0, 0, 18), // "ApproveTxCtx77da62"
QT_MOC_LITERAL(1, 19, 7), // "clicked"
QT_MOC_LITERAL(2, 27, 0), // ""
QT_MOC_LITERAL(3, 28, 1), // "b"
QT_MOC_LITERAL(4, 30, 11), // "checkTxDiff"
QT_MOC_LITERAL(5, 42, 4), // "back"
QT_MOC_LITERAL(6, 47, 6), // "edited"
QT_MOC_LITERAL(7, 54, 1), // "s"
QT_MOC_LITERAL(8, 56, 1), // "v"
QT_MOC_LITERAL(9, 58, 15), // "changeValueUnit"
QT_MOC_LITERAL(10, 74, 18), // "changeGasPriceUnit"
QT_MOC_LITERAL(11, 93, 16), // "valueUnitChanged"
QT_MOC_LITERAL(12, 110, 9), // "valueUnit"
QT_MOC_LITERAL(13, 120, 13), // "remoteChanged"
QT_MOC_LITERAL(14, 134, 6), // "remote"
QT_MOC_LITERAL(15, 141, 16), // "transportChanged"
QT_MOC_LITERAL(16, 158, 9), // "transport"
QT_MOC_LITERAL(17, 168, 15), // "endpointChanged"
QT_MOC_LITERAL(18, 184, 8), // "endpoint"
QT_MOC_LITERAL(19, 193, 11), // "dataChanged"
QT_MOC_LITERAL(20, 205, 4), // "data"
QT_MOC_LITERAL(21, 210, 11), // "fromChanged"
QT_MOC_LITERAL(22, 222, 4), // "from"
QT_MOC_LITERAL(23, 227, 18), // "fromWarningChanged"
QT_MOC_LITERAL(24, 246, 11), // "fromWarning"
QT_MOC_LITERAL(25, 258, 18), // "fromVisibleChanged"
QT_MOC_LITERAL(26, 277, 11), // "fromVisible"
QT_MOC_LITERAL(27, 289, 9), // "toChanged"
QT_MOC_LITERAL(28, 299, 2), // "to"
QT_MOC_LITERAL(29, 302, 16), // "toWarningChanged"
QT_MOC_LITERAL(30, 319, 9), // "toWarning"
QT_MOC_LITERAL(31, 329, 16), // "toVisibleChanged"
QT_MOC_LITERAL(32, 346, 9), // "toVisible"
QT_MOC_LITERAL(33, 356, 10), // "gasChanged"
QT_MOC_LITERAL(34, 367, 3), // "gas"
QT_MOC_LITERAL(35, 371, 15), // "gasPriceChanged"
QT_MOC_LITERAL(36, 387, 8), // "gasPrice"
QT_MOC_LITERAL(37, 396, 19), // "gasPriceUnitChanged"
QT_MOC_LITERAL(38, 416, 12), // "gasPriceUnit"
QT_MOC_LITERAL(39, 429, 12), // "nonceChanged"
QT_MOC_LITERAL(40, 442, 5), // "nonce"
QT_MOC_LITERAL(41, 448, 12), // "valueChanged"
QT_MOC_LITERAL(42, 461, 5), // "value"
QT_MOC_LITERAL(43, 467, 15), // "passwordChanged"
QT_MOC_LITERAL(44, 483, 8), // "password"
QT_MOC_LITERAL(45, 492, 14), // "fromSrcChanged"
QT_MOC_LITERAL(46, 507, 7), // "fromSrc"
QT_MOC_LITERAL(47, 515, 12), // "toSrcChanged"
QT_MOC_LITERAL(48, 528, 5), // "toSrc"
QT_MOC_LITERAL(49, 534, 11), // "diffChanged"
QT_MOC_LITERAL(50, 546, 4) // "diff"

    },
    "ApproveTxCtx77da62\0clicked\0\0b\0checkTxDiff\0"
    "back\0edited\0s\0v\0changeValueUnit\0"
    "changeGasPriceUnit\0valueUnitChanged\0"
    "valueUnit\0remoteChanged\0remote\0"
    "transportChanged\0transport\0endpointChanged\0"
    "endpoint\0dataChanged\0data\0fromChanged\0"
    "from\0fromWarningChanged\0fromWarning\0"
    "fromVisibleChanged\0fromVisible\0toChanged\0"
    "to\0toWarningChanged\0toWarning\0"
    "toVisibleChanged\0toVisible\0gasChanged\0"
    "gas\0gasPriceChanged\0gasPrice\0"
    "gasPriceUnitChanged\0gasPriceUnit\0"
    "nonceChanged\0nonce\0valueChanged\0value\0"
    "passwordChanged\0password\0fromSrcChanged\0"
    "fromSrc\0toSrcChanged\0toSrc\0diffChanged\0"
    "diff"
};
#undef QT_MOC_LITERAL

static const uint qt_meta_data_ApproveTxCtx77da62[] = {

 // content:
       8,       // revision
       0,       // classname
       0,    0, // classinfo
      26,   14, // methods
      20,  220, // properties
       0,    0, // enums/sets
       0,    0, // constructors
       0,       // flags
      26,       // signalCount

 // signals: name, argc, parameters, tag, flags
       1,    1,  144,    2, 0x06 /* Public */,
       4,    0,  147,    2, 0x06 /* Public */,
       5,    0,  148,    2, 0x06 /* Public */,
       6,    2,  149,    2, 0x06 /* Public */,
       9,    1,  154,    2, 0x06 /* Public */,
      10,    1,  157,    2, 0x06 /* Public */,
      11,    1,  160,    2, 0x06 /* Public */,
      13,    1,  163,    2, 0x06 /* Public */,
      15,    1,  166,    2, 0x06 /* Public */,
      17,    1,  169,    2, 0x06 /* Public */,
      19,    1,  172,    2, 0x06 /* Public */,
      21,    1,  175,    2, 0x06 /* Public */,
      23,    1,  178,    2, 0x06 /* Public */,
      25,    1,  181,    2, 0x06 /* Public */,
      27,    1,  184,    2, 0x06 /* Public */,
      29,    1,  187,    2, 0x06 /* Public */,
      31,    1,  190,    2, 0x06 /* Public */,
      33,    1,  193,    2, 0x06 /* Public */,
      35,    1,  196,    2, 0x06 /* Public */,
      37,    1,  199,    2, 0x06 /* Public */,
      39,    1,  202,    2, 0x06 /* Public */,
      41,    1,  205,    2, 0x06 /* Public */,
      43,    1,  208,    2, 0x06 /* Public */,
      45,    1,  211,    2, 0x06 /* Public */,
      47,    1,  214,    2, 0x06 /* Public */,
      49,    1,  217,    2, 0x06 /* Public */,

 // signals: parameters
    QMetaType::Void, QMetaType::Int,    3,
    QMetaType::Void,
    QMetaType::Void,
    QMetaType::Void, QMetaType::QString, QMetaType::QString,    7,    8,
    QMetaType::Void, QMetaType::Int,    8,
    QMetaType::Void, QMetaType::Int,    8,
    QMetaType::Void, QMetaType::QReal,   12,
    QMetaType::Void, QMetaType::QString,   14,
    QMetaType::Void, QMetaType::QString,   16,
    QMetaType::Void, QMetaType::QString,   18,
    QMetaType::Void, QMetaType::QString,   20,
    QMetaType::Void, QMetaType::QString,   22,
    QMetaType::Void, QMetaType::QString,   24,
    QMetaType::Void, QMetaType::Bool,   26,
    QMetaType::Void, QMetaType::QString,   28,
    QMetaType::Void, QMetaType::QString,   30,
    QMetaType::Void, QMetaType::Bool,   32,
    QMetaType::Void, QMetaType::QString,   34,
    QMetaType::Void, QMetaType::QString,   36,
    QMetaType::Void, QMetaType::QReal,   38,
    QMetaType::Void, QMetaType::QString,   40,
    QMetaType::Void, QMetaType::QString,   42,
    QMetaType::Void, QMetaType::QString,   44,
    QMetaType::Void, QMetaType::QString,   46,
    QMetaType::Void, QMetaType::QString,   48,
    QMetaType::Void, QMetaType::QString,   50,

 // properties: name, type, flags
      12, QMetaType::QReal, 0x00495103,
      14, QMetaType::QString, 0x00495103,
      16, QMetaType::QString, 0x00495103,
      18, QMetaType::QString, 0x00495103,
      20, QMetaType::QString, 0x00495103,
      22, QMetaType::QString, 0x00495103,
      24, QMetaType::QString, 0x00495103,
      26, QMetaType::Bool, 0x00495103,
      28, QMetaType::QString, 0x00495103,
      30, QMetaType::QString, 0x00495103,
      32, QMetaType::Bool, 0x00495103,
      34, QMetaType::QString, 0x00495103,
      36, QMetaType::QString, 0x00495103,
      38, QMetaType::QReal, 0x00495103,
      40, QMetaType::QString, 0x00495103,
      42, QMetaType::QString, 0x00495103,
      44, QMetaType::QString, 0x00495103,
      46, QMetaType::QString, 0x00495103,
      48, QMetaType::QString, 0x00495103,
      50, QMetaType::QString, 0x00495103,

 // properties: notify_signal_id
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
      16,
      17,
      18,
      19,
      20,
      21,
      22,
      23,
      24,
      25,

       0        // eod
};

void ApproveTxCtx77da62::qt_static_metacall(QObject *_o, QMetaObject::Call _c, int _id, void **_a)
{
    if (_c == QMetaObject::InvokeMetaMethod) {
        ApproveTxCtx77da62 *_t = static_cast<ApproveTxCtx77da62 *>(_o);
        Q_UNUSED(_t)
        switch (_id) {
        case 0: _t->clicked((*reinterpret_cast< qint32(*)>(_a[1]))); break;
        case 1: _t->checkTxDiff(); break;
        case 2: _t->back(); break;
        case 3: _t->edited((*reinterpret_cast< QString(*)>(_a[1])),(*reinterpret_cast< QString(*)>(_a[2]))); break;
        case 4: _t->changeValueUnit((*reinterpret_cast< qint32(*)>(_a[1]))); break;
        case 5: _t->changeGasPriceUnit((*reinterpret_cast< qint32(*)>(_a[1]))); break;
        case 6: _t->valueUnitChanged((*reinterpret_cast< qreal(*)>(_a[1]))); break;
        case 7: _t->remoteChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 8: _t->transportChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 9: _t->endpointChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 10: _t->dataChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 11: _t->fromChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 12: _t->fromWarningChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 13: _t->fromVisibleChanged((*reinterpret_cast< bool(*)>(_a[1]))); break;
        case 14: _t->toChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 15: _t->toWarningChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 16: _t->toVisibleChanged((*reinterpret_cast< bool(*)>(_a[1]))); break;
        case 17: _t->gasChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 18: _t->gasPriceChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 19: _t->gasPriceUnitChanged((*reinterpret_cast< qreal(*)>(_a[1]))); break;
        case 20: _t->nonceChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 21: _t->valueChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 22: _t->passwordChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 23: _t->fromSrcChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 24: _t->toSrcChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 25: _t->diffChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        default: ;
        }
    } else if (_c == QMetaObject::IndexOfMethod) {
        int *result = reinterpret_cast<int *>(_a[0]);
        {
            using _t = void (ApproveTxCtx77da62::*)(qint32 );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveTxCtx77da62::clicked)) {
                *result = 0;
                return;
            }
        }
        {
            using _t = void (ApproveTxCtx77da62::*)();
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveTxCtx77da62::checkTxDiff)) {
                *result = 1;
                return;
            }
        }
        {
            using _t = void (ApproveTxCtx77da62::*)();
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveTxCtx77da62::back)) {
                *result = 2;
                return;
            }
        }
        {
            using _t = void (ApproveTxCtx77da62::*)(QString , QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveTxCtx77da62::edited)) {
                *result = 3;
                return;
            }
        }
        {
            using _t = void (ApproveTxCtx77da62::*)(qint32 );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveTxCtx77da62::changeValueUnit)) {
                *result = 4;
                return;
            }
        }
        {
            using _t = void (ApproveTxCtx77da62::*)(qint32 );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveTxCtx77da62::changeGasPriceUnit)) {
                *result = 5;
                return;
            }
        }
        {
            using _t = void (ApproveTxCtx77da62::*)(qreal );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveTxCtx77da62::valueUnitChanged)) {
                *result = 6;
                return;
            }
        }
        {
            using _t = void (ApproveTxCtx77da62::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveTxCtx77da62::remoteChanged)) {
                *result = 7;
                return;
            }
        }
        {
            using _t = void (ApproveTxCtx77da62::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveTxCtx77da62::transportChanged)) {
                *result = 8;
                return;
            }
        }
        {
            using _t = void (ApproveTxCtx77da62::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveTxCtx77da62::endpointChanged)) {
                *result = 9;
                return;
            }
        }
        {
            using _t = void (ApproveTxCtx77da62::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveTxCtx77da62::dataChanged)) {
                *result = 10;
                return;
            }
        }
        {
            using _t = void (ApproveTxCtx77da62::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveTxCtx77da62::fromChanged)) {
                *result = 11;
                return;
            }
        }
        {
            using _t = void (ApproveTxCtx77da62::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveTxCtx77da62::fromWarningChanged)) {
                *result = 12;
                return;
            }
        }
        {
            using _t = void (ApproveTxCtx77da62::*)(bool );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveTxCtx77da62::fromVisibleChanged)) {
                *result = 13;
                return;
            }
        }
        {
            using _t = void (ApproveTxCtx77da62::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveTxCtx77da62::toChanged)) {
                *result = 14;
                return;
            }
        }
        {
            using _t = void (ApproveTxCtx77da62::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveTxCtx77da62::toWarningChanged)) {
                *result = 15;
                return;
            }
        }
        {
            using _t = void (ApproveTxCtx77da62::*)(bool );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveTxCtx77da62::toVisibleChanged)) {
                *result = 16;
                return;
            }
        }
        {
            using _t = void (ApproveTxCtx77da62::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveTxCtx77da62::gasChanged)) {
                *result = 17;
                return;
            }
        }
        {
            using _t = void (ApproveTxCtx77da62::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveTxCtx77da62::gasPriceChanged)) {
                *result = 18;
                return;
            }
        }
        {
            using _t = void (ApproveTxCtx77da62::*)(qreal );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveTxCtx77da62::gasPriceUnitChanged)) {
                *result = 19;
                return;
            }
        }
        {
            using _t = void (ApproveTxCtx77da62::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveTxCtx77da62::nonceChanged)) {
                *result = 20;
                return;
            }
        }
        {
            using _t = void (ApproveTxCtx77da62::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveTxCtx77da62::valueChanged)) {
                *result = 21;
                return;
            }
        }
        {
            using _t = void (ApproveTxCtx77da62::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveTxCtx77da62::passwordChanged)) {
                *result = 22;
                return;
            }
        }
        {
            using _t = void (ApproveTxCtx77da62::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveTxCtx77da62::fromSrcChanged)) {
                *result = 23;
                return;
            }
        }
        {
            using _t = void (ApproveTxCtx77da62::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveTxCtx77da62::toSrcChanged)) {
                *result = 24;
                return;
            }
        }
        {
            using _t = void (ApproveTxCtx77da62::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&ApproveTxCtx77da62::diffChanged)) {
                *result = 25;
                return;
            }
        }
    }
#ifndef QT_NO_PROPERTIES
    else if (_c == QMetaObject::ReadProperty) {
        ApproveTxCtx77da62 *_t = static_cast<ApproveTxCtx77da62 *>(_o);
        Q_UNUSED(_t)
        void *_v = _a[0];
        switch (_id) {
        case 0: *reinterpret_cast< qreal*>(_v) = _t->valueUnit(); break;
        case 1: *reinterpret_cast< QString*>(_v) = _t->remote(); break;
        case 2: *reinterpret_cast< QString*>(_v) = _t->transport(); break;
        case 3: *reinterpret_cast< QString*>(_v) = _t->endpoint(); break;
        case 4: *reinterpret_cast< QString*>(_v) = _t->data(); break;
        case 5: *reinterpret_cast< QString*>(_v) = _t->from(); break;
        case 6: *reinterpret_cast< QString*>(_v) = _t->fromWarning(); break;
        case 7: *reinterpret_cast< bool*>(_v) = _t->isFromVisible(); break;
        case 8: *reinterpret_cast< QString*>(_v) = _t->to(); break;
        case 9: *reinterpret_cast< QString*>(_v) = _t->toWarning(); break;
        case 10: *reinterpret_cast< bool*>(_v) = _t->isToVisible(); break;
        case 11: *reinterpret_cast< QString*>(_v) = _t->gas(); break;
        case 12: *reinterpret_cast< QString*>(_v) = _t->gasPrice(); break;
        case 13: *reinterpret_cast< qreal*>(_v) = _t->gasPriceUnit(); break;
        case 14: *reinterpret_cast< QString*>(_v) = _t->nonce(); break;
        case 15: *reinterpret_cast< QString*>(_v) = _t->value(); break;
        case 16: *reinterpret_cast< QString*>(_v) = _t->password(); break;
        case 17: *reinterpret_cast< QString*>(_v) = _t->fromSrc(); break;
        case 18: *reinterpret_cast< QString*>(_v) = _t->toSrc(); break;
        case 19: *reinterpret_cast< QString*>(_v) = _t->diff(); break;
        default: break;
        }
    } else if (_c == QMetaObject::WriteProperty) {
        ApproveTxCtx77da62 *_t = static_cast<ApproveTxCtx77da62 *>(_o);
        Q_UNUSED(_t)
        void *_v = _a[0];
        switch (_id) {
        case 0: _t->setValueUnit(*reinterpret_cast< qreal*>(_v)); break;
        case 1: _t->setRemote(*reinterpret_cast< QString*>(_v)); break;
        case 2: _t->setTransport(*reinterpret_cast< QString*>(_v)); break;
        case 3: _t->setEndpoint(*reinterpret_cast< QString*>(_v)); break;
        case 4: _t->setData(*reinterpret_cast< QString*>(_v)); break;
        case 5: _t->setFrom(*reinterpret_cast< QString*>(_v)); break;
        case 6: _t->setFromWarning(*reinterpret_cast< QString*>(_v)); break;
        case 7: _t->setFromVisible(*reinterpret_cast< bool*>(_v)); break;
        case 8: _t->setTo(*reinterpret_cast< QString*>(_v)); break;
        case 9: _t->setToWarning(*reinterpret_cast< QString*>(_v)); break;
        case 10: _t->setToVisible(*reinterpret_cast< bool*>(_v)); break;
        case 11: _t->setGas(*reinterpret_cast< QString*>(_v)); break;
        case 12: _t->setGasPrice(*reinterpret_cast< QString*>(_v)); break;
        case 13: _t->setGasPriceUnit(*reinterpret_cast< qreal*>(_v)); break;
        case 14: _t->setNonce(*reinterpret_cast< QString*>(_v)); break;
        case 15: _t->setValue(*reinterpret_cast< QString*>(_v)); break;
        case 16: _t->setPassword(*reinterpret_cast< QString*>(_v)); break;
        case 17: _t->setFromSrc(*reinterpret_cast< QString*>(_v)); break;
        case 18: _t->setToSrc(*reinterpret_cast< QString*>(_v)); break;
        case 19: _t->setDiff(*reinterpret_cast< QString*>(_v)); break;
        default: break;
        }
    } else if (_c == QMetaObject::ResetProperty) {
    }
#endif // QT_NO_PROPERTIES
}

QT_INIT_METAOBJECT const QMetaObject ApproveTxCtx77da62::staticMetaObject = { {
    &QObject::staticMetaObject,
    qt_meta_stringdata_ApproveTxCtx77da62.data,
    qt_meta_data_ApproveTxCtx77da62,
    qt_static_metacall,
    nullptr,
    nullptr
} };


const QMetaObject *ApproveTxCtx77da62::metaObject() const
{
    return QObject::d_ptr->metaObject ? QObject::d_ptr->dynamicMetaObject() : &staticMetaObject;
}

void *ApproveTxCtx77da62::qt_metacast(const char *_clname)
{
    if (!_clname) return nullptr;
    if (!strcmp(_clname, qt_meta_stringdata_ApproveTxCtx77da62.stringdata0))
        return static_cast<void*>(this);
    return QObject::qt_metacast(_clname);
}

int ApproveTxCtx77da62::qt_metacall(QMetaObject::Call _c, int _id, void **_a)
{
    _id = QObject::qt_metacall(_c, _id, _a);
    if (_id < 0)
        return _id;
    if (_c == QMetaObject::InvokeMetaMethod) {
        if (_id < 26)
            qt_static_metacall(this, _c, _id, _a);
        _id -= 26;
    } else if (_c == QMetaObject::RegisterMethodArgumentMetaType) {
        if (_id < 26)
            *reinterpret_cast<int*>(_a[0]) = -1;
        _id -= 26;
    }
#ifndef QT_NO_PROPERTIES
   else if (_c == QMetaObject::ReadProperty || _c == QMetaObject::WriteProperty
            || _c == QMetaObject::ResetProperty || _c == QMetaObject::RegisterPropertyMetaType) {
        qt_static_metacall(this, _c, _id, _a);
        _id -= 20;
    } else if (_c == QMetaObject::QueryPropertyDesignable) {
        _id -= 20;
    } else if (_c == QMetaObject::QueryPropertyScriptable) {
        _id -= 20;
    } else if (_c == QMetaObject::QueryPropertyStored) {
        _id -= 20;
    } else if (_c == QMetaObject::QueryPropertyEditable) {
        _id -= 20;
    } else if (_c == QMetaObject::QueryPropertyUser) {
        _id -= 20;
    }
#endif // QT_NO_PROPERTIES
    return _id;
}

// SIGNAL 0
void ApproveTxCtx77da62::clicked(qint32 _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 0, _a);
}

// SIGNAL 1
void ApproveTxCtx77da62::checkTxDiff()
{
    QMetaObject::activate(this, &staticMetaObject, 1, nullptr);
}

// SIGNAL 2
void ApproveTxCtx77da62::back()
{
    QMetaObject::activate(this, &staticMetaObject, 2, nullptr);
}

// SIGNAL 3
void ApproveTxCtx77da62::edited(QString _t1, QString _t2)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)), const_cast<void*>(reinterpret_cast<const void*>(&_t2)) };
    QMetaObject::activate(this, &staticMetaObject, 3, _a);
}

// SIGNAL 4
void ApproveTxCtx77da62::changeValueUnit(qint32 _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 4, _a);
}

// SIGNAL 5
void ApproveTxCtx77da62::changeGasPriceUnit(qint32 _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 5, _a);
}

// SIGNAL 6
void ApproveTxCtx77da62::valueUnitChanged(qreal _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 6, _a);
}

// SIGNAL 7
void ApproveTxCtx77da62::remoteChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 7, _a);
}

// SIGNAL 8
void ApproveTxCtx77da62::transportChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 8, _a);
}

// SIGNAL 9
void ApproveTxCtx77da62::endpointChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 9, _a);
}

// SIGNAL 10
void ApproveTxCtx77da62::dataChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 10, _a);
}

// SIGNAL 11
void ApproveTxCtx77da62::fromChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 11, _a);
}

// SIGNAL 12
void ApproveTxCtx77da62::fromWarningChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 12, _a);
}

// SIGNAL 13
void ApproveTxCtx77da62::fromVisibleChanged(bool _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 13, _a);
}

// SIGNAL 14
void ApproveTxCtx77da62::toChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 14, _a);
}

// SIGNAL 15
void ApproveTxCtx77da62::toWarningChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 15, _a);
}

// SIGNAL 16
void ApproveTxCtx77da62::toVisibleChanged(bool _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 16, _a);
}

// SIGNAL 17
void ApproveTxCtx77da62::gasChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 17, _a);
}

// SIGNAL 18
void ApproveTxCtx77da62::gasPriceChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 18, _a);
}

// SIGNAL 19
void ApproveTxCtx77da62::gasPriceUnitChanged(qreal _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 19, _a);
}

// SIGNAL 20
void ApproveTxCtx77da62::nonceChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 20, _a);
}

// SIGNAL 21
void ApproveTxCtx77da62::valueChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 21, _a);
}

// SIGNAL 22
void ApproveTxCtx77da62::passwordChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 22, _a);
}

// SIGNAL 23
void ApproveTxCtx77da62::fromSrcChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 23, _a);
}

// SIGNAL 24
void ApproveTxCtx77da62::toSrcChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 24, _a);
}

// SIGNAL 25
void ApproveTxCtx77da62::diffChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 25, _a);
}
struct qt_meta_stringdata_TxListCtx77da62_t {
    QByteArrayData data[8];
    char stringdata0[95];
};
#define QT_MOC_LITERAL(idx, ofs, len) \
    Q_STATIC_BYTE_ARRAY_DATA_HEADER_INITIALIZER_WITH_OFFSET(len, \
    qptrdiff(offsetof(qt_meta_stringdata_TxListCtx77da62_t, stringdata0) + ofs \
        - idx * sizeof(QByteArrayData)) \
    )
static const qt_meta_stringdata_TxListCtx77da62_t qt_meta_stringdata_TxListCtx77da62 = {
    {
QT_MOC_LITERAL(0, 0, 15), // "TxListCtx77da62"
QT_MOC_LITERAL(1, 16, 7), // "clicked"
QT_MOC_LITERAL(2, 24, 0), // ""
QT_MOC_LITERAL(3, 25, 1), // "b"
QT_MOC_LITERAL(4, 27, 21), // "shortenAddressChanged"
QT_MOC_LITERAL(5, 49, 14), // "shortenAddress"
QT_MOC_LITERAL(6, 64, 18), // "selectedSrcChanged"
QT_MOC_LITERAL(7, 83, 11) // "selectedSrc"

    },
    "TxListCtx77da62\0clicked\0\0b\0"
    "shortenAddressChanged\0shortenAddress\0"
    "selectedSrcChanged\0selectedSrc"
};
#undef QT_MOC_LITERAL

static const uint qt_meta_data_TxListCtx77da62[] = {

 // content:
       8,       // revision
       0,       // classname
       0,    0, // classinfo
       3,   14, // methods
       2,   38, // properties
       0,    0, // enums/sets
       0,    0, // constructors
       0,       // flags
       3,       // signalCount

 // signals: name, argc, parameters, tag, flags
       1,    1,   29,    2, 0x06 /* Public */,
       4,    1,   32,    2, 0x06 /* Public */,
       6,    1,   35,    2, 0x06 /* Public */,

 // signals: parameters
    QMetaType::Void, QMetaType::Int,    3,
    QMetaType::Void, QMetaType::QString,    5,
    QMetaType::Void, QMetaType::QString,    7,

 // properties: name, type, flags
       5, QMetaType::QString, 0x00495103,
       7, QMetaType::QString, 0x00495103,

 // properties: notify_signal_id
       1,
       2,

       0        // eod
};

void TxListCtx77da62::qt_static_metacall(QObject *_o, QMetaObject::Call _c, int _id, void **_a)
{
    if (_c == QMetaObject::InvokeMetaMethod) {
        TxListCtx77da62 *_t = static_cast<TxListCtx77da62 *>(_o);
        Q_UNUSED(_t)
        switch (_id) {
        case 0: _t->clicked((*reinterpret_cast< qint32(*)>(_a[1]))); break;
        case 1: _t->shortenAddressChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 2: _t->selectedSrcChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        default: ;
        }
    } else if (_c == QMetaObject::IndexOfMethod) {
        int *result = reinterpret_cast<int *>(_a[0]);
        {
            using _t = void (TxListCtx77da62::*)(qint32 );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&TxListCtx77da62::clicked)) {
                *result = 0;
                return;
            }
        }
        {
            using _t = void (TxListCtx77da62::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&TxListCtx77da62::shortenAddressChanged)) {
                *result = 1;
                return;
            }
        }
        {
            using _t = void (TxListCtx77da62::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&TxListCtx77da62::selectedSrcChanged)) {
                *result = 2;
                return;
            }
        }
    }
#ifndef QT_NO_PROPERTIES
    else if (_c == QMetaObject::ReadProperty) {
        TxListCtx77da62 *_t = static_cast<TxListCtx77da62 *>(_o);
        Q_UNUSED(_t)
        void *_v = _a[0];
        switch (_id) {
        case 0: *reinterpret_cast< QString*>(_v) = _t->shortenAddress(); break;
        case 1: *reinterpret_cast< QString*>(_v) = _t->selectedSrc(); break;
        default: break;
        }
    } else if (_c == QMetaObject::WriteProperty) {
        TxListCtx77da62 *_t = static_cast<TxListCtx77da62 *>(_o);
        Q_UNUSED(_t)
        void *_v = _a[0];
        switch (_id) {
        case 0: _t->setShortenAddress(*reinterpret_cast< QString*>(_v)); break;
        case 1: _t->setSelectedSrc(*reinterpret_cast< QString*>(_v)); break;
        default: break;
        }
    } else if (_c == QMetaObject::ResetProperty) {
    }
#endif // QT_NO_PROPERTIES
}

QT_INIT_METAOBJECT const QMetaObject TxListCtx77da62::staticMetaObject = { {
    &QObject::staticMetaObject,
    qt_meta_stringdata_TxListCtx77da62.data,
    qt_meta_data_TxListCtx77da62,
    qt_static_metacall,
    nullptr,
    nullptr
} };


const QMetaObject *TxListCtx77da62::metaObject() const
{
    return QObject::d_ptr->metaObject ? QObject::d_ptr->dynamicMetaObject() : &staticMetaObject;
}

void *TxListCtx77da62::qt_metacast(const char *_clname)
{
    if (!_clname) return nullptr;
    if (!strcmp(_clname, qt_meta_stringdata_TxListCtx77da62.stringdata0))
        return static_cast<void*>(this);
    return QObject::qt_metacast(_clname);
}

int TxListCtx77da62::qt_metacall(QMetaObject::Call _c, int _id, void **_a)
{
    _id = QObject::qt_metacall(_c, _id, _a);
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
#ifndef QT_NO_PROPERTIES
   else if (_c == QMetaObject::ReadProperty || _c == QMetaObject::WriteProperty
            || _c == QMetaObject::ResetProperty || _c == QMetaObject::RegisterPropertyMetaType) {
        qt_static_metacall(this, _c, _id, _a);
        _id -= 2;
    } else if (_c == QMetaObject::QueryPropertyDesignable) {
        _id -= 2;
    } else if (_c == QMetaObject::QueryPropertyScriptable) {
        _id -= 2;
    } else if (_c == QMetaObject::QueryPropertyStored) {
        _id -= 2;
    } else if (_c == QMetaObject::QueryPropertyEditable) {
        _id -= 2;
    } else if (_c == QMetaObject::QueryPropertyUser) {
        _id -= 2;
    }
#endif // QT_NO_PROPERTIES
    return _id;
}

// SIGNAL 0
void TxListCtx77da62::clicked(qint32 _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 0, _a);
}

// SIGNAL 1
void TxListCtx77da62::shortenAddressChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 1, _a);
}

// SIGNAL 2
void TxListCtx77da62::selectedSrcChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 2, _a);
}
struct qt_meta_stringdata_TxListModel77da62_t {
    QByteArrayData data[10];
    char stringdata0[73];
};
#define QT_MOC_LITERAL(idx, ofs, len) \
    Q_STATIC_BYTE_ARRAY_DATA_HEADER_INITIALIZER_WITH_OFFSET(len, \
    qptrdiff(offsetof(qt_meta_stringdata_TxListModel77da62_t, stringdata0) + ofs \
        - idx * sizeof(QByteArrayData)) \
    )
static const qt_meta_stringdata_TxListModel77da62_t qt_meta_stringdata_TxListModel77da62 = {
    {
QT_MOC_LITERAL(0, 0, 17), // "TxListModel77da62"
QT_MOC_LITERAL(1, 18, 5), // "clear"
QT_MOC_LITERAL(2, 24, 0), // ""
QT_MOC_LITERAL(3, 25, 3), // "add"
QT_MOC_LITERAL(4, 29, 8), // "quintptr"
QT_MOC_LITERAL(5, 38, 2), // "tx"
QT_MOC_LITERAL(6, 41, 6), // "remove"
QT_MOC_LITERAL(7, 48, 1), // "i"
QT_MOC_LITERAL(8, 50, 14), // "isEmptyChanged"
QT_MOC_LITERAL(9, 65, 7) // "isEmpty"

    },
    "TxListModel77da62\0clear\0\0add\0quintptr\0"
    "tx\0remove\0i\0isEmptyChanged\0isEmpty"
};
#undef QT_MOC_LITERAL

static const uint qt_meta_data_TxListModel77da62[] = {

 // content:
       8,       // revision
       0,       // classname
       0,    0, // classinfo
       4,   14, // methods
       1,   44, // properties
       0,    0, // enums/sets
       0,    0, // constructors
       0,       // flags
       4,       // signalCount

 // signals: name, argc, parameters, tag, flags
       1,    0,   34,    2, 0x06 /* Public */,
       3,    1,   35,    2, 0x06 /* Public */,
       6,    1,   38,    2, 0x06 /* Public */,
       8,    1,   41,    2, 0x06 /* Public */,

 // signals: parameters
    QMetaType::Void,
    QMetaType::Void, 0x80000000 | 4,    5,
    QMetaType::Void, QMetaType::Int,    7,
    QMetaType::Void, QMetaType::Bool,    9,

 // properties: name, type, flags
       9, QMetaType::Bool, 0x00495103,

 // properties: notify_signal_id
       3,

       0        // eod
};

void TxListModel77da62::qt_static_metacall(QObject *_o, QMetaObject::Call _c, int _id, void **_a)
{
    if (_c == QMetaObject::InvokeMetaMethod) {
        TxListModel77da62 *_t = static_cast<TxListModel77da62 *>(_o);
        Q_UNUSED(_t)
        switch (_id) {
        case 0: _t->clear(); break;
        case 1: _t->add((*reinterpret_cast< quintptr(*)>(_a[1]))); break;
        case 2: _t->remove((*reinterpret_cast< qint32(*)>(_a[1]))); break;
        case 3: _t->isEmptyChanged((*reinterpret_cast< bool(*)>(_a[1]))); break;
        default: ;
        }
    } else if (_c == QMetaObject::IndexOfMethod) {
        int *result = reinterpret_cast<int *>(_a[0]);
        {
            using _t = void (TxListModel77da62::*)();
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&TxListModel77da62::clear)) {
                *result = 0;
                return;
            }
        }
        {
            using _t = void (TxListModel77da62::*)(quintptr );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&TxListModel77da62::add)) {
                *result = 1;
                return;
            }
        }
        {
            using _t = void (TxListModel77da62::*)(qint32 );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&TxListModel77da62::remove)) {
                *result = 2;
                return;
            }
        }
        {
            using _t = void (TxListModel77da62::*)(bool );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&TxListModel77da62::isEmptyChanged)) {
                *result = 3;
                return;
            }
        }
    }
#ifndef QT_NO_PROPERTIES
    else if (_c == QMetaObject::ReadProperty) {
        TxListModel77da62 *_t = static_cast<TxListModel77da62 *>(_o);
        Q_UNUSED(_t)
        void *_v = _a[0];
        switch (_id) {
        case 0: *reinterpret_cast< bool*>(_v) = _t->isEmpty(); break;
        default: break;
        }
    } else if (_c == QMetaObject::WriteProperty) {
        TxListModel77da62 *_t = static_cast<TxListModel77da62 *>(_o);
        Q_UNUSED(_t)
        void *_v = _a[0];
        switch (_id) {
        case 0: _t->setIsEmpty(*reinterpret_cast< bool*>(_v)); break;
        default: break;
        }
    } else if (_c == QMetaObject::ResetProperty) {
    }
#endif // QT_NO_PROPERTIES
}

QT_INIT_METAOBJECT const QMetaObject TxListModel77da62::staticMetaObject = { {
    &QAbstractListModel::staticMetaObject,
    qt_meta_stringdata_TxListModel77da62.data,
    qt_meta_data_TxListModel77da62,
    qt_static_metacall,
    nullptr,
    nullptr
} };


const QMetaObject *TxListModel77da62::metaObject() const
{
    return QObject::d_ptr->metaObject ? QObject::d_ptr->dynamicMetaObject() : &staticMetaObject;
}

void *TxListModel77da62::qt_metacast(const char *_clname)
{
    if (!_clname) return nullptr;
    if (!strcmp(_clname, qt_meta_stringdata_TxListModel77da62.stringdata0))
        return static_cast<void*>(this);
    return QAbstractListModel::qt_metacast(_clname);
}

int TxListModel77da62::qt_metacall(QMetaObject::Call _c, int _id, void **_a)
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
#ifndef QT_NO_PROPERTIES
   else if (_c == QMetaObject::ReadProperty || _c == QMetaObject::WriteProperty
            || _c == QMetaObject::ResetProperty || _c == QMetaObject::RegisterPropertyMetaType) {
        qt_static_metacall(this, _c, _id, _a);
        _id -= 1;
    } else if (_c == QMetaObject::QueryPropertyDesignable) {
        _id -= 1;
    } else if (_c == QMetaObject::QueryPropertyScriptable) {
        _id -= 1;
    } else if (_c == QMetaObject::QueryPropertyStored) {
        _id -= 1;
    } else if (_c == QMetaObject::QueryPropertyEditable) {
        _id -= 1;
    } else if (_c == QMetaObject::QueryPropertyUser) {
        _id -= 1;
    }
#endif // QT_NO_PROPERTIES
    return _id;
}

// SIGNAL 0
void TxListModel77da62::clear()
{
    QMetaObject::activate(this, &staticMetaObject, 0, nullptr);
}

// SIGNAL 1
void TxListModel77da62::add(quintptr _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 1, _a);
}

// SIGNAL 2
void TxListModel77da62::remove(qint32 _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 2, _a);
}

// SIGNAL 3
void TxListModel77da62::isEmptyChanged(bool _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 3, _a);
}
QT_WARNING_POP
QT_END_MOC_NAMESPACE
