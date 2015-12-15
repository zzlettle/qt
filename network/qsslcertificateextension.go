package network

//#include "network.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QSslCertificateExtension struct {
	ptr unsafe.Pointer
}

type QSslCertificateExtension_ITF interface {
	QSslCertificateExtension_PTR() *QSslCertificateExtension
}

func (p *QSslCertificateExtension) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QSslCertificateExtension) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQSslCertificateExtension(ptr QSslCertificateExtension_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSslCertificateExtension_PTR().Pointer()
	}
	return nil
}

func NewQSslCertificateExtensionFromPointer(ptr unsafe.Pointer) *QSslCertificateExtension {
	var n = new(QSslCertificateExtension)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSslCertificateExtension) QSslCertificateExtension_PTR() *QSslCertificateExtension {
	return ptr
}

func NewQSslCertificateExtension() *QSslCertificateExtension {
	defer qt.Recovering("QSslCertificateExtension::QSslCertificateExtension")

	return NewQSslCertificateExtensionFromPointer(C.QSslCertificateExtension_NewQSslCertificateExtension())
}

func NewQSslCertificateExtension2(other QSslCertificateExtension_ITF) *QSslCertificateExtension {
	defer qt.Recovering("QSslCertificateExtension::QSslCertificateExtension")

	return NewQSslCertificateExtensionFromPointer(C.QSslCertificateExtension_NewQSslCertificateExtension2(PointerFromQSslCertificateExtension(other)))
}

func (ptr *QSslCertificateExtension) IsCritical() bool {
	defer qt.Recovering("QSslCertificateExtension::isCritical")

	if ptr.Pointer() != nil {
		return C.QSslCertificateExtension_IsCritical(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSslCertificateExtension) IsSupported() bool {
	defer qt.Recovering("QSslCertificateExtension::isSupported")

	if ptr.Pointer() != nil {
		return C.QSslCertificateExtension_IsSupported(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSslCertificateExtension) Name() string {
	defer qt.Recovering("QSslCertificateExtension::name")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSslCertificateExtension_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSslCertificateExtension) Oid() string {
	defer qt.Recovering("QSslCertificateExtension::oid")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSslCertificateExtension_Oid(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSslCertificateExtension) Swap(other QSslCertificateExtension_ITF) {
	defer qt.Recovering("QSslCertificateExtension::swap")

	if ptr.Pointer() != nil {
		C.QSslCertificateExtension_Swap(ptr.Pointer(), PointerFromQSslCertificateExtension(other))
	}
}

func (ptr *QSslCertificateExtension) Value() *core.QVariant {
	defer qt.Recovering("QSslCertificateExtension::value")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QSslCertificateExtension_Value(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSslCertificateExtension) DestroyQSslCertificateExtension() {
	defer qt.Recovering("QSslCertificateExtension::~QSslCertificateExtension")

	if ptr.Pointer() != nil {
		C.QSslCertificateExtension_DestroyQSslCertificateExtension(ptr.Pointer())
	}
}
