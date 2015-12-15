package macextras

//#include "macextras.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QMacToolBar struct {
	core.QObject
}

type QMacToolBar_ITF interface {
	core.QObject_ITF
	QMacToolBar_PTR() *QMacToolBar
}

func PointerFromQMacToolBar(ptr QMacToolBar_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMacToolBar_PTR().Pointer()
	}
	return nil
}

func NewQMacToolBarFromPointer(ptr unsafe.Pointer) *QMacToolBar {
	var n = new(QMacToolBar)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QMacToolBar_") {
		n.SetObjectName("QMacToolBar_" + qt.Identifier())
	}
	return n
}

func (ptr *QMacToolBar) QMacToolBar_PTR() *QMacToolBar {
	return ptr
}

func NewQMacToolBar(parent core.QObject_ITF) *QMacToolBar {
	defer qt.Recovering("QMacToolBar::QMacToolBar")

	return NewQMacToolBarFromPointer(C.QMacToolBar_NewQMacToolBar(core.PointerFromQObject(parent)))
}

func NewQMacToolBar2(identifier string, parent core.QObject_ITF) *QMacToolBar {
	defer qt.Recovering("QMacToolBar::QMacToolBar")

	return NewQMacToolBarFromPointer(C.QMacToolBar_NewQMacToolBar2(C.CString(identifier), core.PointerFromQObject(parent)))
}

func (ptr *QMacToolBar) AddAllowedItem(icon gui.QIcon_ITF, text string) *QMacToolBarItem {
	defer qt.Recovering("QMacToolBar::addAllowedItem")

	if ptr.Pointer() != nil {
		return NewQMacToolBarItemFromPointer(C.QMacToolBar_AddAllowedItem(ptr.Pointer(), gui.PointerFromQIcon(icon), C.CString(text)))
	}
	return nil
}

func (ptr *QMacToolBar) AddItem(icon gui.QIcon_ITF, text string) *QMacToolBarItem {
	defer qt.Recovering("QMacToolBar::addItem")

	if ptr.Pointer() != nil {
		return NewQMacToolBarItemFromPointer(C.QMacToolBar_AddItem(ptr.Pointer(), gui.PointerFromQIcon(icon), C.CString(text)))
	}
	return nil
}

func (ptr *QMacToolBar) AddSeparator() {
	defer qt.Recovering("QMacToolBar::addSeparator")

	if ptr.Pointer() != nil {
		C.QMacToolBar_AddSeparator(ptr.Pointer())
	}
}

func (ptr *QMacToolBar) AttachToWindow(window gui.QWindow_ITF) {
	defer qt.Recovering("QMacToolBar::attachToWindow")

	if ptr.Pointer() != nil {
		C.QMacToolBar_AttachToWindow(ptr.Pointer(), gui.PointerFromQWindow(window))
	}
}

func (ptr *QMacToolBar) DetachFromWindow() {
	defer qt.Recovering("QMacToolBar::detachFromWindow")

	if ptr.Pointer() != nil {
		C.QMacToolBar_DetachFromWindow(ptr.Pointer())
	}
}

func (ptr *QMacToolBar) DestroyQMacToolBar() {
	defer qt.Recovering("QMacToolBar::~QMacToolBar")

	if ptr.Pointer() != nil {
		C.QMacToolBar_DestroyQMacToolBar(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
