package script

//#include "script.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QScriptValue struct {
	ptr unsafe.Pointer
}

type QScriptValue_ITF interface {
	QScriptValue_PTR() *QScriptValue
}

func (p *QScriptValue) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QScriptValue) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQScriptValue(ptr QScriptValue_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScriptValue_PTR().Pointer()
	}
	return nil
}

func NewQScriptValueFromPointer(ptr unsafe.Pointer) *QScriptValue {
	var n = new(QScriptValue)
	n.SetPointer(ptr)
	return n
}

func (ptr *QScriptValue) QScriptValue_PTR() *QScriptValue {
	return ptr
}

//QScriptValue::PropertyFlag
type QScriptValue__PropertyFlag int64

const (
	QScriptValue__ReadOnly          = QScriptValue__PropertyFlag(0x00000001)
	QScriptValue__Undeletable       = QScriptValue__PropertyFlag(0x00000002)
	QScriptValue__SkipInEnumeration = QScriptValue__PropertyFlag(0x00000004)
	QScriptValue__PropertyGetter    = QScriptValue__PropertyFlag(0x00000008)
	QScriptValue__PropertySetter    = QScriptValue__PropertyFlag(0x00000010)
	QScriptValue__QObjectMember     = QScriptValue__PropertyFlag(0x00000020)
	QScriptValue__KeepExistingFlags = QScriptValue__PropertyFlag(0x00000800)
	QScriptValue__UserRange         = QScriptValue__PropertyFlag(0xff000000)
)

//QScriptValue::ResolveFlag
type QScriptValue__ResolveFlag int64

const (
	QScriptValue__ResolveLocal     = QScriptValue__ResolveFlag(0x00)
	QScriptValue__ResolvePrototype = QScriptValue__ResolveFlag(0x01)
	QScriptValue__ResolveScope     = QScriptValue__ResolveFlag(0x02)
	QScriptValue__ResolveFull      = QScriptValue__ResolveFlag(QScriptValue__ResolvePrototype | QScriptValue__ResolveScope)
)

//QScriptValue::SpecialValue
type QScriptValue__SpecialValue int64

const (
	QScriptValue__NullValue      = QScriptValue__SpecialValue(0)
	QScriptValue__UndefinedValue = QScriptValue__SpecialValue(1)
)

func NewQScriptValue() *QScriptValue {
	defer qt.Recovering("QScriptValue::QScriptValue")

	return NewQScriptValueFromPointer(C.QScriptValue_NewQScriptValue())
}

func NewQScriptValue10(value QScriptValue__SpecialValue) *QScriptValue {
	defer qt.Recovering("QScriptValue::QScriptValue")

	return NewQScriptValueFromPointer(C.QScriptValue_NewQScriptValue10(C.int(value)))
}

func NewQScriptValue11(value bool) *QScriptValue {
	defer qt.Recovering("QScriptValue::QScriptValue")

	return NewQScriptValueFromPointer(C.QScriptValue_NewQScriptValue11(C.int(qt.GoBoolToInt(value))))
}

func NewQScriptValue16(value core.QLatin1String_ITF) *QScriptValue {
	defer qt.Recovering("QScriptValue::QScriptValue")

	return NewQScriptValueFromPointer(C.QScriptValue_NewQScriptValue16(core.PointerFromQLatin1String(value)))
}

func NewQScriptValue2(other QScriptValue_ITF) *QScriptValue {
	defer qt.Recovering("QScriptValue::QScriptValue")

	return NewQScriptValueFromPointer(C.QScriptValue_NewQScriptValue2(PointerFromQScriptValue(other)))
}

func NewQScriptValue15(value string) *QScriptValue {
	defer qt.Recovering("QScriptValue::QScriptValue")

	return NewQScriptValueFromPointer(C.QScriptValue_NewQScriptValue15(C.CString(value)))
}

func NewQScriptValue17(value string) *QScriptValue {
	defer qt.Recovering("QScriptValue::QScriptValue")

	return NewQScriptValueFromPointer(C.QScriptValue_NewQScriptValue17(C.CString(value)))
}

func NewQScriptValue12(value int) *QScriptValue {
	defer qt.Recovering("QScriptValue::QScriptValue")

	return NewQScriptValueFromPointer(C.QScriptValue_NewQScriptValue12(C.int(value)))
}

func (ptr *QScriptValue) Call2(thisObject QScriptValue_ITF, arguments QScriptValue_ITF) *QScriptValue {
	defer qt.Recovering("QScriptValue::call")

	if ptr.Pointer() != nil {
		return NewQScriptValueFromPointer(C.QScriptValue_Call2(ptr.Pointer(), PointerFromQScriptValue(thisObject), PointerFromQScriptValue(arguments)))
	}
	return nil
}

func (ptr *QScriptValue) Construct2(arguments QScriptValue_ITF) *QScriptValue {
	defer qt.Recovering("QScriptValue::construct")

	if ptr.Pointer() != nil {
		return NewQScriptValueFromPointer(C.QScriptValue_Construct2(ptr.Pointer(), PointerFromQScriptValue(arguments)))
	}
	return nil
}

func (ptr *QScriptValue) Data() *QScriptValue {
	defer qt.Recovering("QScriptValue::data")

	if ptr.Pointer() != nil {
		return NewQScriptValueFromPointer(C.QScriptValue_Data(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScriptValue) Engine() *QScriptEngine {
	defer qt.Recovering("QScriptValue::engine")

	if ptr.Pointer() != nil {
		return NewQScriptEngineFromPointer(C.QScriptValue_Engine(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScriptValue) Equals(other QScriptValue_ITF) bool {
	defer qt.Recovering("QScriptValue::equals")

	if ptr.Pointer() != nil {
		return C.QScriptValue_Equals(ptr.Pointer(), PointerFromQScriptValue(other)) != 0
	}
	return false
}

func (ptr *QScriptValue) InstanceOf(other QScriptValue_ITF) bool {
	defer qt.Recovering("QScriptValue::instanceOf")

	if ptr.Pointer() != nil {
		return C.QScriptValue_InstanceOf(ptr.Pointer(), PointerFromQScriptValue(other)) != 0
	}
	return false
}

func (ptr *QScriptValue) IsArray() bool {
	defer qt.Recovering("QScriptValue::isArray")

	if ptr.Pointer() != nil {
		return C.QScriptValue_IsArray(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QScriptValue) IsBool() bool {
	defer qt.Recovering("QScriptValue::isBool")

	if ptr.Pointer() != nil {
		return C.QScriptValue_IsBool(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QScriptValue) IsDate() bool {
	defer qt.Recovering("QScriptValue::isDate")

	if ptr.Pointer() != nil {
		return C.QScriptValue_IsDate(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QScriptValue) IsError() bool {
	defer qt.Recovering("QScriptValue::isError")

	if ptr.Pointer() != nil {
		return C.QScriptValue_IsError(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QScriptValue) IsFunction() bool {
	defer qt.Recovering("QScriptValue::isFunction")

	if ptr.Pointer() != nil {
		return C.QScriptValue_IsFunction(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QScriptValue) IsNull() bool {
	defer qt.Recovering("QScriptValue::isNull")

	if ptr.Pointer() != nil {
		return C.QScriptValue_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QScriptValue) IsNumber() bool {
	defer qt.Recovering("QScriptValue::isNumber")

	if ptr.Pointer() != nil {
		return C.QScriptValue_IsNumber(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QScriptValue) IsObject() bool {
	defer qt.Recovering("QScriptValue::isObject")

	if ptr.Pointer() != nil {
		return C.QScriptValue_IsObject(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QScriptValue) IsQMetaObject() bool {
	defer qt.Recovering("QScriptValue::isQMetaObject")

	if ptr.Pointer() != nil {
		return C.QScriptValue_IsQMetaObject(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QScriptValue) IsQObject() bool {
	defer qt.Recovering("QScriptValue::isQObject")

	if ptr.Pointer() != nil {
		return C.QScriptValue_IsQObject(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QScriptValue) IsRegExp() bool {
	defer qt.Recovering("QScriptValue::isRegExp")

	if ptr.Pointer() != nil {
		return C.QScriptValue_IsRegExp(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QScriptValue) IsString() bool {
	defer qt.Recovering("QScriptValue::isString")

	if ptr.Pointer() != nil {
		return C.QScriptValue_IsString(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QScriptValue) IsUndefined() bool {
	defer qt.Recovering("QScriptValue::isUndefined")

	if ptr.Pointer() != nil {
		return C.QScriptValue_IsUndefined(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QScriptValue) IsValid() bool {
	defer qt.Recovering("QScriptValue::isValid")

	if ptr.Pointer() != nil {
		return C.QScriptValue_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QScriptValue) IsVariant() bool {
	defer qt.Recovering("QScriptValue::isVariant")

	if ptr.Pointer() != nil {
		return C.QScriptValue_IsVariant(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QScriptValue) LessThan(other QScriptValue_ITF) bool {
	defer qt.Recovering("QScriptValue::lessThan")

	if ptr.Pointer() != nil {
		return C.QScriptValue_LessThan(ptr.Pointer(), PointerFromQScriptValue(other)) != 0
	}
	return false
}

func (ptr *QScriptValue) Property2(name QScriptString_ITF, mode QScriptValue__ResolveFlag) *QScriptValue {
	defer qt.Recovering("QScriptValue::property")

	if ptr.Pointer() != nil {
		return NewQScriptValueFromPointer(C.QScriptValue_Property2(ptr.Pointer(), PointerFromQScriptString(name), C.int(mode)))
	}
	return nil
}

func (ptr *QScriptValue) Property(name string, mode QScriptValue__ResolveFlag) *QScriptValue {
	defer qt.Recovering("QScriptValue::property")

	if ptr.Pointer() != nil {
		return NewQScriptValueFromPointer(C.QScriptValue_Property(ptr.Pointer(), C.CString(name), C.int(mode)))
	}
	return nil
}

func (ptr *QScriptValue) PropertyFlags2(name QScriptString_ITF, mode QScriptValue__ResolveFlag) QScriptValue__PropertyFlag {
	defer qt.Recovering("QScriptValue::propertyFlags")

	if ptr.Pointer() != nil {
		return QScriptValue__PropertyFlag(C.QScriptValue_PropertyFlags2(ptr.Pointer(), PointerFromQScriptString(name), C.int(mode)))
	}
	return 0
}

func (ptr *QScriptValue) PropertyFlags(name string, mode QScriptValue__ResolveFlag) QScriptValue__PropertyFlag {
	defer qt.Recovering("QScriptValue::propertyFlags")

	if ptr.Pointer() != nil {
		return QScriptValue__PropertyFlag(C.QScriptValue_PropertyFlags(ptr.Pointer(), C.CString(name), C.int(mode)))
	}
	return 0
}

func (ptr *QScriptValue) Prototype() *QScriptValue {
	defer qt.Recovering("QScriptValue::prototype")

	if ptr.Pointer() != nil {
		return NewQScriptValueFromPointer(C.QScriptValue_Prototype(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScriptValue) ScriptClass() *QScriptClass {
	defer qt.Recovering("QScriptValue::scriptClass")

	if ptr.Pointer() != nil {
		return NewQScriptClassFromPointer(C.QScriptValue_ScriptClass(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScriptValue) SetData(data QScriptValue_ITF) {
	defer qt.Recovering("QScriptValue::setData")

	if ptr.Pointer() != nil {
		C.QScriptValue_SetData(ptr.Pointer(), PointerFromQScriptValue(data))
	}
}

func (ptr *QScriptValue) SetProperty2(name QScriptString_ITF, value QScriptValue_ITF, flags QScriptValue__PropertyFlag) {
	defer qt.Recovering("QScriptValue::setProperty")

	if ptr.Pointer() != nil {
		C.QScriptValue_SetProperty2(ptr.Pointer(), PointerFromQScriptString(name), PointerFromQScriptValue(value), C.int(flags))
	}
}

func (ptr *QScriptValue) SetProperty(name string, value QScriptValue_ITF, flags QScriptValue__PropertyFlag) {
	defer qt.Recovering("QScriptValue::setProperty")

	if ptr.Pointer() != nil {
		C.QScriptValue_SetProperty(ptr.Pointer(), C.CString(name), PointerFromQScriptValue(value), C.int(flags))
	}
}

func (ptr *QScriptValue) SetPrototype(prototype QScriptValue_ITF) {
	defer qt.Recovering("QScriptValue::setPrototype")

	if ptr.Pointer() != nil {
		C.QScriptValue_SetPrototype(ptr.Pointer(), PointerFromQScriptValue(prototype))
	}
}

func (ptr *QScriptValue) SetScriptClass(scriptClass QScriptClass_ITF) {
	defer qt.Recovering("QScriptValue::setScriptClass")

	if ptr.Pointer() != nil {
		C.QScriptValue_SetScriptClass(ptr.Pointer(), PointerFromQScriptClass(scriptClass))
	}
}

func (ptr *QScriptValue) StrictlyEquals(other QScriptValue_ITF) bool {
	defer qt.Recovering("QScriptValue::strictlyEquals")

	if ptr.Pointer() != nil {
		return C.QScriptValue_StrictlyEquals(ptr.Pointer(), PointerFromQScriptValue(other)) != 0
	}
	return false
}

func (ptr *QScriptValue) ToBool() bool {
	defer qt.Recovering("QScriptValue::toBool")

	if ptr.Pointer() != nil {
		return C.QScriptValue_ToBool(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QScriptValue) ToDateTime() *core.QDateTime {
	defer qt.Recovering("QScriptValue::toDateTime")

	if ptr.Pointer() != nil {
		return core.NewQDateTimeFromPointer(C.QScriptValue_ToDateTime(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScriptValue) ToQMetaObject() *core.QMetaObject {
	defer qt.Recovering("QScriptValue::toQMetaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QScriptValue_ToQMetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScriptValue) ToQObject() *core.QObject {
	defer qt.Recovering("QScriptValue::toQObject")

	if ptr.Pointer() != nil {
		return core.NewQObjectFromPointer(C.QScriptValue_ToQObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScriptValue) ToRegExp() *core.QRegExp {
	defer qt.Recovering("QScriptValue::toRegExp")

	if ptr.Pointer() != nil {
		return core.NewQRegExpFromPointer(C.QScriptValue_ToRegExp(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScriptValue) ToString() string {
	defer qt.Recovering("QScriptValue::toString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QScriptValue_ToString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QScriptValue) ToVariant() *core.QVariant {
	defer qt.Recovering("QScriptValue::toVariant")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QScriptValue_ToVariant(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScriptValue) DestroyQScriptValue() {
	defer qt.Recovering("QScriptValue::~QScriptValue")

	if ptr.Pointer() != nil {
		C.QScriptValue_DestroyQScriptValue(ptr.Pointer())
	}
}
