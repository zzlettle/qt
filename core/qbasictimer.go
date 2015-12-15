package core

//#include "core.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QBasicTimer struct {
	ptr unsafe.Pointer
}

type QBasicTimer_ITF interface {
	QBasicTimer_PTR() *QBasicTimer
}

func (p *QBasicTimer) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QBasicTimer) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQBasicTimer(ptr QBasicTimer_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QBasicTimer_PTR().Pointer()
	}
	return nil
}

func NewQBasicTimerFromPointer(ptr unsafe.Pointer) *QBasicTimer {
	var n = new(QBasicTimer)
	n.SetPointer(ptr)
	return n
}

func (ptr *QBasicTimer) QBasicTimer_PTR() *QBasicTimer {
	return ptr
}

func (ptr *QBasicTimer) Start(msec int, object QObject_ITF) {
	defer qt.Recovering("QBasicTimer::start")

	if ptr.Pointer() != nil {
		C.QBasicTimer_Start(ptr.Pointer(), C.int(msec), PointerFromQObject(object))
	}
}

func NewQBasicTimer() *QBasicTimer {
	defer qt.Recovering("QBasicTimer::QBasicTimer")

	return NewQBasicTimerFromPointer(C.QBasicTimer_NewQBasicTimer())
}

func (ptr *QBasicTimer) IsActive() bool {
	defer qt.Recovering("QBasicTimer::isActive")

	if ptr.Pointer() != nil {
		return C.QBasicTimer_IsActive(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QBasicTimer) Start2(msec int, timerType Qt__TimerType, obj QObject_ITF) {
	defer qt.Recovering("QBasicTimer::start")

	if ptr.Pointer() != nil {
		C.QBasicTimer_Start2(ptr.Pointer(), C.int(msec), C.int(timerType), PointerFromQObject(obj))
	}
}

func (ptr *QBasicTimer) Stop() {
	defer qt.Recovering("QBasicTimer::stop")

	if ptr.Pointer() != nil {
		C.QBasicTimer_Stop(ptr.Pointer())
	}
}

func (ptr *QBasicTimer) TimerId() int {
	defer qt.Recovering("QBasicTimer::timerId")

	if ptr.Pointer() != nil {
		return int(C.QBasicTimer_TimerId(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBasicTimer) DestroyQBasicTimer() {
	defer qt.Recovering("QBasicTimer::~QBasicTimer")

	if ptr.Pointer() != nil {
		C.QBasicTimer_DestroyQBasicTimer(ptr.Pointer())
	}
}
