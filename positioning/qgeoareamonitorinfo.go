package positioning

//#include "positioning.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QGeoAreaMonitorInfo struct {
	ptr unsafe.Pointer
}

type QGeoAreaMonitorInfo_ITF interface {
	QGeoAreaMonitorInfo_PTR() *QGeoAreaMonitorInfo
}

func (p *QGeoAreaMonitorInfo) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QGeoAreaMonitorInfo) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQGeoAreaMonitorInfo(ptr QGeoAreaMonitorInfo_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGeoAreaMonitorInfo_PTR().Pointer()
	}
	return nil
}

func NewQGeoAreaMonitorInfoFromPointer(ptr unsafe.Pointer) *QGeoAreaMonitorInfo {
	var n = new(QGeoAreaMonitorInfo)
	n.SetPointer(ptr)
	return n
}

func (ptr *QGeoAreaMonitorInfo) QGeoAreaMonitorInfo_PTR() *QGeoAreaMonitorInfo {
	return ptr
}

func NewQGeoAreaMonitorInfo2(other QGeoAreaMonitorInfo_ITF) *QGeoAreaMonitorInfo {
	defer qt.Recovering("QGeoAreaMonitorInfo::QGeoAreaMonitorInfo")

	return NewQGeoAreaMonitorInfoFromPointer(C.QGeoAreaMonitorInfo_NewQGeoAreaMonitorInfo2(PointerFromQGeoAreaMonitorInfo(other)))
}

func NewQGeoAreaMonitorInfo(name string) *QGeoAreaMonitorInfo {
	defer qt.Recovering("QGeoAreaMonitorInfo::QGeoAreaMonitorInfo")

	return NewQGeoAreaMonitorInfoFromPointer(C.QGeoAreaMonitorInfo_NewQGeoAreaMonitorInfo(C.CString(name)))
}

func (ptr *QGeoAreaMonitorInfo) Expiration() *core.QDateTime {
	defer qt.Recovering("QGeoAreaMonitorInfo::expiration")

	if ptr.Pointer() != nil {
		return core.NewQDateTimeFromPointer(C.QGeoAreaMonitorInfo_Expiration(ptr.Pointer()))
	}
	return nil
}

func (ptr *QGeoAreaMonitorInfo) Identifier() string {
	defer qt.Recovering("QGeoAreaMonitorInfo::identifier")

	if ptr.Pointer() != nil {
		return C.GoString(C.QGeoAreaMonitorInfo_Identifier(ptr.Pointer()))
	}
	return ""
}

func (ptr *QGeoAreaMonitorInfo) IsPersistent() bool {
	defer qt.Recovering("QGeoAreaMonitorInfo::isPersistent")

	if ptr.Pointer() != nil {
		return C.QGeoAreaMonitorInfo_IsPersistent(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QGeoAreaMonitorInfo) IsValid() bool {
	defer qt.Recovering("QGeoAreaMonitorInfo::isValid")

	if ptr.Pointer() != nil {
		return C.QGeoAreaMonitorInfo_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QGeoAreaMonitorInfo) Name() string {
	defer qt.Recovering("QGeoAreaMonitorInfo::name")

	if ptr.Pointer() != nil {
		return C.GoString(C.QGeoAreaMonitorInfo_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QGeoAreaMonitorInfo) SetArea(newShape QGeoShape_ITF) {
	defer qt.Recovering("QGeoAreaMonitorInfo::setArea")

	if ptr.Pointer() != nil {
		C.QGeoAreaMonitorInfo_SetArea(ptr.Pointer(), PointerFromQGeoShape(newShape))
	}
}

func (ptr *QGeoAreaMonitorInfo) SetExpiration(expiry core.QDateTime_ITF) {
	defer qt.Recovering("QGeoAreaMonitorInfo::setExpiration")

	if ptr.Pointer() != nil {
		C.QGeoAreaMonitorInfo_SetExpiration(ptr.Pointer(), core.PointerFromQDateTime(expiry))
	}
}

func (ptr *QGeoAreaMonitorInfo) SetName(name string) {
	defer qt.Recovering("QGeoAreaMonitorInfo::setName")

	if ptr.Pointer() != nil {
		C.QGeoAreaMonitorInfo_SetName(ptr.Pointer(), C.CString(name))
	}
}

func (ptr *QGeoAreaMonitorInfo) SetPersistent(isPersistent bool) {
	defer qt.Recovering("QGeoAreaMonitorInfo::setPersistent")

	if ptr.Pointer() != nil {
		C.QGeoAreaMonitorInfo_SetPersistent(ptr.Pointer(), C.int(qt.GoBoolToInt(isPersistent)))
	}
}

func (ptr *QGeoAreaMonitorInfo) DestroyQGeoAreaMonitorInfo() {
	defer qt.Recovering("QGeoAreaMonitorInfo::~QGeoAreaMonitorInfo")

	if ptr.Pointer() != nil {
		C.QGeoAreaMonitorInfo_DestroyQGeoAreaMonitorInfo(ptr.Pointer())
	}
}
