package core

//#include "core.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QTemporaryDir struct {
	ptr unsafe.Pointer
}

type QTemporaryDir_ITF interface {
	QTemporaryDir_PTR() *QTemporaryDir
}

func (p *QTemporaryDir) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QTemporaryDir) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQTemporaryDir(ptr QTemporaryDir_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTemporaryDir_PTR().Pointer()
	}
	return nil
}

func NewQTemporaryDirFromPointer(ptr unsafe.Pointer) *QTemporaryDir {
	var n = new(QTemporaryDir)
	n.SetPointer(ptr)
	return n
}

func (ptr *QTemporaryDir) QTemporaryDir_PTR() *QTemporaryDir {
	return ptr
}

func NewQTemporaryDir() *QTemporaryDir {
	defer qt.Recovering("QTemporaryDir::QTemporaryDir")

	return NewQTemporaryDirFromPointer(C.QTemporaryDir_NewQTemporaryDir())
}

func NewQTemporaryDir2(templatePath string) *QTemporaryDir {
	defer qt.Recovering("QTemporaryDir::QTemporaryDir")

	return NewQTemporaryDirFromPointer(C.QTemporaryDir_NewQTemporaryDir2(C.CString(templatePath)))
}

func (ptr *QTemporaryDir) AutoRemove() bool {
	defer qt.Recovering("QTemporaryDir::autoRemove")

	if ptr.Pointer() != nil {
		return C.QTemporaryDir_AutoRemove(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTemporaryDir) IsValid() bool {
	defer qt.Recovering("QTemporaryDir::isValid")

	if ptr.Pointer() != nil {
		return C.QTemporaryDir_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTemporaryDir) Path() string {
	defer qt.Recovering("QTemporaryDir::path")

	if ptr.Pointer() != nil {
		return C.GoString(C.QTemporaryDir_Path(ptr.Pointer()))
	}
	return ""
}

func (ptr *QTemporaryDir) SetAutoRemove(b bool) {
	defer qt.Recovering("QTemporaryDir::setAutoRemove")

	if ptr.Pointer() != nil {
		C.QTemporaryDir_SetAutoRemove(ptr.Pointer(), C.int(qt.GoBoolToInt(b)))
	}
}

func (ptr *QTemporaryDir) DestroyQTemporaryDir() {
	defer qt.Recovering("QTemporaryDir::~QTemporaryDir")

	if ptr.Pointer() != nil {
		C.QTemporaryDir_DestroyQTemporaryDir(ptr.Pointer())
	}
}
