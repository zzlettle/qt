package core

//#include "core.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QWriteLocker struct {
	ptr unsafe.Pointer
}

type QWriteLocker_ITF interface {
	QWriteLocker_PTR() *QWriteLocker
}

func (p *QWriteLocker) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QWriteLocker) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQWriteLocker(ptr QWriteLocker_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWriteLocker_PTR().Pointer()
	}
	return nil
}

func NewQWriteLockerFromPointer(ptr unsafe.Pointer) *QWriteLocker {
	var n = new(QWriteLocker)
	n.SetPointer(ptr)
	return n
}

func (ptr *QWriteLocker) QWriteLocker_PTR() *QWriteLocker {
	return ptr
}

func NewQWriteLocker(lock QReadWriteLock_ITF) *QWriteLocker {
	defer qt.Recovering("QWriteLocker::QWriteLocker")

	return NewQWriteLockerFromPointer(C.QWriteLocker_NewQWriteLocker(PointerFromQReadWriteLock(lock)))
}

func (ptr *QWriteLocker) ReadWriteLock() *QReadWriteLock {
	defer qt.Recovering("QWriteLocker::readWriteLock")

	if ptr.Pointer() != nil {
		return NewQReadWriteLockFromPointer(C.QWriteLocker_ReadWriteLock(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWriteLocker) Relock() {
	defer qt.Recovering("QWriteLocker::relock")

	if ptr.Pointer() != nil {
		C.QWriteLocker_Relock(ptr.Pointer())
	}
}

func (ptr *QWriteLocker) Unlock() {
	defer qt.Recovering("QWriteLocker::unlock")

	if ptr.Pointer() != nil {
		C.QWriteLocker_Unlock(ptr.Pointer())
	}
}

func (ptr *QWriteLocker) DestroyQWriteLocker() {
	defer qt.Recovering("QWriteLocker::~QWriteLocker")

	if ptr.Pointer() != nil {
		C.QWriteLocker_DestroyQWriteLocker(ptr.Pointer())
	}
}
