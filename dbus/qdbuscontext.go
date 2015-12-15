package dbus

//#include "dbus.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QDBusContext struct {
	ptr unsafe.Pointer
}

type QDBusContext_ITF interface {
	QDBusContext_PTR() *QDBusContext
}

func (p *QDBusContext) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QDBusContext) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQDBusContext(ptr QDBusContext_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusContext_PTR().Pointer()
	}
	return nil
}

func NewQDBusContextFromPointer(ptr unsafe.Pointer) *QDBusContext {
	var n = new(QDBusContext)
	n.SetPointer(ptr)
	return n
}

func (ptr *QDBusContext) QDBusContext_PTR() *QDBusContext {
	return ptr
}

func NewQDBusContext() *QDBusContext {
	defer qt.Recovering("QDBusContext::QDBusContext")

	return NewQDBusContextFromPointer(C.QDBusContext_NewQDBusContext())
}

func (ptr *QDBusContext) CalledFromDBus() bool {
	defer qt.Recovering("QDBusContext::calledFromDBus")

	if ptr.Pointer() != nil {
		return C.QDBusContext_CalledFromDBus(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDBusContext) IsDelayedReply() bool {
	defer qt.Recovering("QDBusContext::isDelayedReply")

	if ptr.Pointer() != nil {
		return C.QDBusContext_IsDelayedReply(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDBusContext) SendErrorReply2(ty QDBusError__ErrorType, msg string) {
	defer qt.Recovering("QDBusContext::sendErrorReply")

	if ptr.Pointer() != nil {
		C.QDBusContext_SendErrorReply2(ptr.Pointer(), C.int(ty), C.CString(msg))
	}
}

func (ptr *QDBusContext) SendErrorReply(name string, msg string) {
	defer qt.Recovering("QDBusContext::sendErrorReply")

	if ptr.Pointer() != nil {
		C.QDBusContext_SendErrorReply(ptr.Pointer(), C.CString(name), C.CString(msg))
	}
}

func (ptr *QDBusContext) SetDelayedReply(enable bool) {
	defer qt.Recovering("QDBusContext::setDelayedReply")

	if ptr.Pointer() != nil {
		C.QDBusContext_SetDelayedReply(ptr.Pointer(), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QDBusContext) DestroyQDBusContext() {
	defer qt.Recovering("QDBusContext::~QDBusContext")

	if ptr.Pointer() != nil {
		C.QDBusContext_DestroyQDBusContext(ptr.Pointer())
	}
}
