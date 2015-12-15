package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QGlyphRun struct {
	ptr unsafe.Pointer
}

type QGlyphRun_ITF interface {
	QGlyphRun_PTR() *QGlyphRun
}

func (p *QGlyphRun) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QGlyphRun) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQGlyphRun(ptr QGlyphRun_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGlyphRun_PTR().Pointer()
	}
	return nil
}

func NewQGlyphRunFromPointer(ptr unsafe.Pointer) *QGlyphRun {
	var n = new(QGlyphRun)
	n.SetPointer(ptr)
	return n
}

func (ptr *QGlyphRun) QGlyphRun_PTR() *QGlyphRun {
	return ptr
}

//QGlyphRun::GlyphRunFlag
type QGlyphRun__GlyphRunFlag int64

const (
	QGlyphRun__Overline      = QGlyphRun__GlyphRunFlag(0x01)
	QGlyphRun__Underline     = QGlyphRun__GlyphRunFlag(0x02)
	QGlyphRun__StrikeOut     = QGlyphRun__GlyphRunFlag(0x04)
	QGlyphRun__RightToLeft   = QGlyphRun__GlyphRunFlag(0x08)
	QGlyphRun__SplitLigature = QGlyphRun__GlyphRunFlag(0x10)
)

func NewQGlyphRun() *QGlyphRun {
	defer qt.Recovering("QGlyphRun::QGlyphRun")

	return NewQGlyphRunFromPointer(C.QGlyphRun_NewQGlyphRun())
}

func NewQGlyphRun2(other QGlyphRun_ITF) *QGlyphRun {
	defer qt.Recovering("QGlyphRun::QGlyphRun")

	return NewQGlyphRunFromPointer(C.QGlyphRun_NewQGlyphRun2(PointerFromQGlyphRun(other)))
}

func (ptr *QGlyphRun) Clear() {
	defer qt.Recovering("QGlyphRun::clear")

	if ptr.Pointer() != nil {
		C.QGlyphRun_Clear(ptr.Pointer())
	}
}

func (ptr *QGlyphRun) Flags() QGlyphRun__GlyphRunFlag {
	defer qt.Recovering("QGlyphRun::flags")

	if ptr.Pointer() != nil {
		return QGlyphRun__GlyphRunFlag(C.QGlyphRun_Flags(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGlyphRun) IsEmpty() bool {
	defer qt.Recovering("QGlyphRun::isEmpty")

	if ptr.Pointer() != nil {
		return C.QGlyphRun_IsEmpty(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QGlyphRun) IsRightToLeft() bool {
	defer qt.Recovering("QGlyphRun::isRightToLeft")

	if ptr.Pointer() != nil {
		return C.QGlyphRun_IsRightToLeft(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QGlyphRun) Overline() bool {
	defer qt.Recovering("QGlyphRun::overline")

	if ptr.Pointer() != nil {
		return C.QGlyphRun_Overline(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QGlyphRun) SetBoundingRect(boundingRect core.QRectF_ITF) {
	defer qt.Recovering("QGlyphRun::setBoundingRect")

	if ptr.Pointer() != nil {
		C.QGlyphRun_SetBoundingRect(ptr.Pointer(), core.PointerFromQRectF(boundingRect))
	}
}

func (ptr *QGlyphRun) SetFlag(flag QGlyphRun__GlyphRunFlag, enabled bool) {
	defer qt.Recovering("QGlyphRun::setFlag")

	if ptr.Pointer() != nil {
		C.QGlyphRun_SetFlag(ptr.Pointer(), C.int(flag), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QGlyphRun) SetFlags(flags QGlyphRun__GlyphRunFlag) {
	defer qt.Recovering("QGlyphRun::setFlags")

	if ptr.Pointer() != nil {
		C.QGlyphRun_SetFlags(ptr.Pointer(), C.int(flags))
	}
}

func (ptr *QGlyphRun) SetOverline(overline bool) {
	defer qt.Recovering("QGlyphRun::setOverline")

	if ptr.Pointer() != nil {
		C.QGlyphRun_SetOverline(ptr.Pointer(), C.int(qt.GoBoolToInt(overline)))
	}
}

func (ptr *QGlyphRun) SetRawFont(rawFont QRawFont_ITF) {
	defer qt.Recovering("QGlyphRun::setRawFont")

	if ptr.Pointer() != nil {
		C.QGlyphRun_SetRawFont(ptr.Pointer(), PointerFromQRawFont(rawFont))
	}
}

func (ptr *QGlyphRun) SetRightToLeft(rightToLeft bool) {
	defer qt.Recovering("QGlyphRun::setRightToLeft")

	if ptr.Pointer() != nil {
		C.QGlyphRun_SetRightToLeft(ptr.Pointer(), C.int(qt.GoBoolToInt(rightToLeft)))
	}
}

func (ptr *QGlyphRun) SetStrikeOut(strikeOut bool) {
	defer qt.Recovering("QGlyphRun::setStrikeOut")

	if ptr.Pointer() != nil {
		C.QGlyphRun_SetStrikeOut(ptr.Pointer(), C.int(qt.GoBoolToInt(strikeOut)))
	}
}

func (ptr *QGlyphRun) SetUnderline(underline bool) {
	defer qt.Recovering("QGlyphRun::setUnderline")

	if ptr.Pointer() != nil {
		C.QGlyphRun_SetUnderline(ptr.Pointer(), C.int(qt.GoBoolToInt(underline)))
	}
}

func (ptr *QGlyphRun) StrikeOut() bool {
	defer qt.Recovering("QGlyphRun::strikeOut")

	if ptr.Pointer() != nil {
		return C.QGlyphRun_StrikeOut(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QGlyphRun) Swap(other QGlyphRun_ITF) {
	defer qt.Recovering("QGlyphRun::swap")

	if ptr.Pointer() != nil {
		C.QGlyphRun_Swap(ptr.Pointer(), PointerFromQGlyphRun(other))
	}
}

func (ptr *QGlyphRun) Underline() bool {
	defer qt.Recovering("QGlyphRun::underline")

	if ptr.Pointer() != nil {
		return C.QGlyphRun_Underline(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QGlyphRun) DestroyQGlyphRun() {
	defer qt.Recovering("QGlyphRun::~QGlyphRun")

	if ptr.Pointer() != nil {
		C.QGlyphRun_DestroyQGlyphRun(ptr.Pointer())
	}
}
