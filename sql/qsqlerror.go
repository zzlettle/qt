package sql

//#include "sql.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QSqlError struct {
	ptr unsafe.Pointer
}

type QSqlError_ITF interface {
	QSqlError_PTR() *QSqlError
}

func (p *QSqlError) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QSqlError) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQSqlError(ptr QSqlError_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSqlError_PTR().Pointer()
	}
	return nil
}

func NewQSqlErrorFromPointer(ptr unsafe.Pointer) *QSqlError {
	var n = new(QSqlError)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSqlError) QSqlError_PTR() *QSqlError {
	return ptr
}

//QSqlError::ErrorType
type QSqlError__ErrorType int64

const (
	QSqlError__NoError          = QSqlError__ErrorType(0)
	QSqlError__ConnectionError  = QSqlError__ErrorType(1)
	QSqlError__StatementError   = QSqlError__ErrorType(2)
	QSqlError__TransactionError = QSqlError__ErrorType(3)
	QSqlError__UnknownError     = QSqlError__ErrorType(4)
)

func NewQSqlError3(other QSqlError_ITF) *QSqlError {
	defer qt.Recovering("QSqlError::QSqlError")

	return NewQSqlErrorFromPointer(C.QSqlError_NewQSqlError3(PointerFromQSqlError(other)))
}

func NewQSqlError(driverText string, databaseText string, ty QSqlError__ErrorType, code string) *QSqlError {
	defer qt.Recovering("QSqlError::QSqlError")

	return NewQSqlErrorFromPointer(C.QSqlError_NewQSqlError(C.CString(driverText), C.CString(databaseText), C.int(ty), C.CString(code)))
}

func (ptr *QSqlError) DatabaseText() string {
	defer qt.Recovering("QSqlError::databaseText")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlError_DatabaseText(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSqlError) DriverText() string {
	defer qt.Recovering("QSqlError::driverText")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlError_DriverText(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSqlError) IsValid() bool {
	defer qt.Recovering("QSqlError::isValid")

	if ptr.Pointer() != nil {
		return C.QSqlError_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlError) NativeErrorCode() string {
	defer qt.Recovering("QSqlError::nativeErrorCode")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlError_NativeErrorCode(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSqlError) Text() string {
	defer qt.Recovering("QSqlError::text")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlError_Text(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSqlError) Type() QSqlError__ErrorType {
	defer qt.Recovering("QSqlError::type")

	if ptr.Pointer() != nil {
		return QSqlError__ErrorType(C.QSqlError_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSqlError) DestroyQSqlError() {
	defer qt.Recovering("QSqlError::~QSqlError")

	if ptr.Pointer() != nil {
		C.QSqlError_DestroyQSqlError(ptr.Pointer())
	}
}
