package positioning

//#include "positioning.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QGeoCircle struct {
	QGeoShape
}

type QGeoCircle_ITF interface {
	QGeoShape_ITF
	QGeoCircle_PTR() *QGeoCircle
}

func PointerFromQGeoCircle(ptr QGeoCircle_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGeoCircle_PTR().Pointer()
	}
	return nil
}

func NewQGeoCircleFromPointer(ptr unsafe.Pointer) *QGeoCircle {
	var n = new(QGeoCircle)
	n.SetPointer(ptr)
	return n
}

func (ptr *QGeoCircle) QGeoCircle_PTR() *QGeoCircle {
	return ptr
}

func NewQGeoCircle() *QGeoCircle {
	defer qt.Recovering("QGeoCircle::QGeoCircle")

	return NewQGeoCircleFromPointer(C.QGeoCircle_NewQGeoCircle())
}

func NewQGeoCircle3(other QGeoCircle_ITF) *QGeoCircle {
	defer qt.Recovering("QGeoCircle::QGeoCircle")

	return NewQGeoCircleFromPointer(C.QGeoCircle_NewQGeoCircle3(PointerFromQGeoCircle(other)))
}

func NewQGeoCircle2(center QGeoCoordinate_ITF, radius float64) *QGeoCircle {
	defer qt.Recovering("QGeoCircle::QGeoCircle")

	return NewQGeoCircleFromPointer(C.QGeoCircle_NewQGeoCircle2(PointerFromQGeoCoordinate(center), C.double(radius)))
}

func NewQGeoCircle4(other QGeoShape_ITF) *QGeoCircle {
	defer qt.Recovering("QGeoCircle::QGeoCircle")

	return NewQGeoCircleFromPointer(C.QGeoCircle_NewQGeoCircle4(PointerFromQGeoShape(other)))
}

func (ptr *QGeoCircle) Radius() float64 {
	defer qt.Recovering("QGeoCircle::radius")

	if ptr.Pointer() != nil {
		return float64(C.QGeoCircle_Radius(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGeoCircle) SetCenter(center QGeoCoordinate_ITF) {
	defer qt.Recovering("QGeoCircle::setCenter")

	if ptr.Pointer() != nil {
		C.QGeoCircle_SetCenter(ptr.Pointer(), PointerFromQGeoCoordinate(center))
	}
}

func (ptr *QGeoCircle) SetRadius(radius float64) {
	defer qt.Recovering("QGeoCircle::setRadius")

	if ptr.Pointer() != nil {
		C.QGeoCircle_SetRadius(ptr.Pointer(), C.double(radius))
	}
}

func (ptr *QGeoCircle) ToString() string {
	defer qt.Recovering("QGeoCircle::toString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QGeoCircle_ToString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QGeoCircle) DestroyQGeoCircle() {
	defer qt.Recovering("QGeoCircle::~QGeoCircle")

	if ptr.Pointer() != nil {
		C.QGeoCircle_DestroyQGeoCircle(ptr.Pointer())
	}
}
