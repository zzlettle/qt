package sql

//#include "sql.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QSqlField struct {
	ptr unsafe.Pointer
}

type QSqlField_ITF interface {
	QSqlField_PTR() *QSqlField
}

func (p *QSqlField) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QSqlField) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQSqlField(ptr QSqlField_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSqlField_PTR().Pointer()
	}
	return nil
}

func NewQSqlFieldFromPointer(ptr unsafe.Pointer) *QSqlField {
	var n = new(QSqlField)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSqlField) QSqlField_PTR() *QSqlField {
	return ptr
}

//QSqlField::RequiredStatus
type QSqlField__RequiredStatus int64

const (
	QSqlField__Unknown  = QSqlField__RequiredStatus(-1)
	QSqlField__Optional = QSqlField__RequiredStatus(0)
	QSqlField__Required = QSqlField__RequiredStatus(1)
)

func NewQSqlField2(other QSqlField_ITF) *QSqlField {
	defer qt.Recovering("QSqlField::QSqlField")

	return NewQSqlFieldFromPointer(C.QSqlField_NewQSqlField2(PointerFromQSqlField(other)))
}

func (ptr *QSqlField) Clear() {
	defer qt.Recovering("QSqlField::clear")

	if ptr.Pointer() != nil {
		C.QSqlField_Clear(ptr.Pointer())
	}
}

func (ptr *QSqlField) DefaultValue() *core.QVariant {
	defer qt.Recovering("QSqlField::defaultValue")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QSqlField_DefaultValue(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSqlField) IsAutoValue() bool {
	defer qt.Recovering("QSqlField::isAutoValue")

	if ptr.Pointer() != nil {
		return C.QSqlField_IsAutoValue(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlField) IsGenerated() bool {
	defer qt.Recovering("QSqlField::isGenerated")

	if ptr.Pointer() != nil {
		return C.QSqlField_IsGenerated(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlField) IsNull() bool {
	defer qt.Recovering("QSqlField::isNull")

	if ptr.Pointer() != nil {
		return C.QSqlField_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlField) IsReadOnly() bool {
	defer qt.Recovering("QSqlField::isReadOnly")

	if ptr.Pointer() != nil {
		return C.QSqlField_IsReadOnly(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlField) IsValid() bool {
	defer qt.Recovering("QSqlField::isValid")

	if ptr.Pointer() != nil {
		return C.QSqlField_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlField) Length() int {
	defer qt.Recovering("QSqlField::length")

	if ptr.Pointer() != nil {
		return int(C.QSqlField_Length(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSqlField) Name() string {
	defer qt.Recovering("QSqlField::name")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlField_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSqlField) Precision() int {
	defer qt.Recovering("QSqlField::precision")

	if ptr.Pointer() != nil {
		return int(C.QSqlField_Precision(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSqlField) RequiredStatus() QSqlField__RequiredStatus {
	defer qt.Recovering("QSqlField::requiredStatus")

	if ptr.Pointer() != nil {
		return QSqlField__RequiredStatus(C.QSqlField_RequiredStatus(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSqlField) SetAutoValue(autoVal bool) {
	defer qt.Recovering("QSqlField::setAutoValue")

	if ptr.Pointer() != nil {
		C.QSqlField_SetAutoValue(ptr.Pointer(), C.int(qt.GoBoolToInt(autoVal)))
	}
}

func (ptr *QSqlField) SetDefaultValue(value core.QVariant_ITF) {
	defer qt.Recovering("QSqlField::setDefaultValue")

	if ptr.Pointer() != nil {
		C.QSqlField_SetDefaultValue(ptr.Pointer(), core.PointerFromQVariant(value))
	}
}

func (ptr *QSqlField) SetGenerated(gen bool) {
	defer qt.Recovering("QSqlField::setGenerated")

	if ptr.Pointer() != nil {
		C.QSqlField_SetGenerated(ptr.Pointer(), C.int(qt.GoBoolToInt(gen)))
	}
}

func (ptr *QSqlField) SetLength(fieldLength int) {
	defer qt.Recovering("QSqlField::setLength")

	if ptr.Pointer() != nil {
		C.QSqlField_SetLength(ptr.Pointer(), C.int(fieldLength))
	}
}

func (ptr *QSqlField) SetName(name string) {
	defer qt.Recovering("QSqlField::setName")

	if ptr.Pointer() != nil {
		C.QSqlField_SetName(ptr.Pointer(), C.CString(name))
	}
}

func (ptr *QSqlField) SetPrecision(precision int) {
	defer qt.Recovering("QSqlField::setPrecision")

	if ptr.Pointer() != nil {
		C.QSqlField_SetPrecision(ptr.Pointer(), C.int(precision))
	}
}

func (ptr *QSqlField) SetReadOnly(readOnly bool) {
	defer qt.Recovering("QSqlField::setReadOnly")

	if ptr.Pointer() != nil {
		C.QSqlField_SetReadOnly(ptr.Pointer(), C.int(qt.GoBoolToInt(readOnly)))
	}
}

func (ptr *QSqlField) SetRequired(required bool) {
	defer qt.Recovering("QSqlField::setRequired")

	if ptr.Pointer() != nil {
		C.QSqlField_SetRequired(ptr.Pointer(), C.int(qt.GoBoolToInt(required)))
	}
}

func (ptr *QSqlField) SetRequiredStatus(required QSqlField__RequiredStatus) {
	defer qt.Recovering("QSqlField::setRequiredStatus")

	if ptr.Pointer() != nil {
		C.QSqlField_SetRequiredStatus(ptr.Pointer(), C.int(required))
	}
}

func (ptr *QSqlField) SetValue(value core.QVariant_ITF) {
	defer qt.Recovering("QSqlField::setValue")

	if ptr.Pointer() != nil {
		C.QSqlField_SetValue(ptr.Pointer(), core.PointerFromQVariant(value))
	}
}

func (ptr *QSqlField) Value() *core.QVariant {
	defer qt.Recovering("QSqlField::value")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QSqlField_Value(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSqlField) DestroyQSqlField() {
	defer qt.Recovering("QSqlField::~QSqlField")

	if ptr.Pointer() != nil {
		C.QSqlField_DestroyQSqlField(ptr.Pointer())
	}
}
