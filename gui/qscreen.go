package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QScreen struct {
	core.QObject
}

type QScreen_ITF interface {
	core.QObject_ITF
	QScreen_PTR() *QScreen
}

func PointerFromQScreen(ptr QScreen_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScreen_PTR().Pointer()
	}
	return nil
}

func NewQScreenFromPointer(ptr unsafe.Pointer) *QScreen {
	var n = new(QScreen)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QScreen_") {
		n.SetObjectName("QScreen_" + qt.Identifier())
	}
	return n
}

func (ptr *QScreen) QScreen_PTR() *QScreen {
	return ptr
}

func (ptr *QScreen) Depth() int {
	defer qt.Recovering("QScreen::depth")

	if ptr.Pointer() != nil {
		return int(C.QScreen_Depth(ptr.Pointer()))
	}
	return 0
}

func (ptr *QScreen) DevicePixelRatio() float64 {
	defer qt.Recovering("QScreen::devicePixelRatio")

	if ptr.Pointer() != nil {
		return float64(C.QScreen_DevicePixelRatio(ptr.Pointer()))
	}
	return 0
}

func (ptr *QScreen) LogicalDotsPerInch() float64 {
	defer qt.Recovering("QScreen::logicalDotsPerInch")

	if ptr.Pointer() != nil {
		return float64(C.QScreen_LogicalDotsPerInch(ptr.Pointer()))
	}
	return 0
}

func (ptr *QScreen) LogicalDotsPerInchX() float64 {
	defer qt.Recovering("QScreen::logicalDotsPerInchX")

	if ptr.Pointer() != nil {
		return float64(C.QScreen_LogicalDotsPerInchX(ptr.Pointer()))
	}
	return 0
}

func (ptr *QScreen) LogicalDotsPerInchY() float64 {
	defer qt.Recovering("QScreen::logicalDotsPerInchY")

	if ptr.Pointer() != nil {
		return float64(C.QScreen_LogicalDotsPerInchY(ptr.Pointer()))
	}
	return 0
}

func (ptr *QScreen) Name() string {
	defer qt.Recovering("QScreen::name")

	if ptr.Pointer() != nil {
		return C.GoString(C.QScreen_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QScreen) NativeOrientation() core.Qt__ScreenOrientation {
	defer qt.Recovering("QScreen::nativeOrientation")

	if ptr.Pointer() != nil {
		return core.Qt__ScreenOrientation(C.QScreen_NativeOrientation(ptr.Pointer()))
	}
	return 0
}

func (ptr *QScreen) Orientation() core.Qt__ScreenOrientation {
	defer qt.Recovering("QScreen::orientation")

	if ptr.Pointer() != nil {
		return core.Qt__ScreenOrientation(C.QScreen_Orientation(ptr.Pointer()))
	}
	return 0
}

func (ptr *QScreen) PhysicalDotsPerInch() float64 {
	defer qt.Recovering("QScreen::physicalDotsPerInch")

	if ptr.Pointer() != nil {
		return float64(C.QScreen_PhysicalDotsPerInch(ptr.Pointer()))
	}
	return 0
}

func (ptr *QScreen) PhysicalDotsPerInchX() float64 {
	defer qt.Recovering("QScreen::physicalDotsPerInchX")

	if ptr.Pointer() != nil {
		return float64(C.QScreen_PhysicalDotsPerInchX(ptr.Pointer()))
	}
	return 0
}

func (ptr *QScreen) PhysicalDotsPerInchY() float64 {
	defer qt.Recovering("QScreen::physicalDotsPerInchY")

	if ptr.Pointer() != nil {
		return float64(C.QScreen_PhysicalDotsPerInchY(ptr.Pointer()))
	}
	return 0
}

func (ptr *QScreen) PrimaryOrientation() core.Qt__ScreenOrientation {
	defer qt.Recovering("QScreen::primaryOrientation")

	if ptr.Pointer() != nil {
		return core.Qt__ScreenOrientation(C.QScreen_PrimaryOrientation(ptr.Pointer()))
	}
	return 0
}

func (ptr *QScreen) RefreshRate() float64 {
	defer qt.Recovering("QScreen::refreshRate")

	if ptr.Pointer() != nil {
		return float64(C.QScreen_RefreshRate(ptr.Pointer()))
	}
	return 0
}

func (ptr *QScreen) AngleBetween(a core.Qt__ScreenOrientation, b core.Qt__ScreenOrientation) int {
	defer qt.Recovering("QScreen::angleBetween")

	if ptr.Pointer() != nil {
		return int(C.QScreen_AngleBetween(ptr.Pointer(), C.int(a), C.int(b)))
	}
	return 0
}

func (ptr *QScreen) IsLandscape(o core.Qt__ScreenOrientation) bool {
	defer qt.Recovering("QScreen::isLandscape")

	if ptr.Pointer() != nil {
		return C.QScreen_IsLandscape(ptr.Pointer(), C.int(o)) != 0
	}
	return false
}

func (ptr *QScreen) IsPortrait(o core.Qt__ScreenOrientation) bool {
	defer qt.Recovering("QScreen::isPortrait")

	if ptr.Pointer() != nil {
		return C.QScreen_IsPortrait(ptr.Pointer(), C.int(o)) != 0
	}
	return false
}

func (ptr *QScreen) ConnectOrientationChanged(f func(orientation core.Qt__ScreenOrientation)) {
	defer qt.Recovering("connect QScreen::orientationChanged")

	if ptr.Pointer() != nil {
		C.QScreen_ConnectOrientationChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "orientationChanged", f)
	}
}

func (ptr *QScreen) DisconnectOrientationChanged() {
	defer qt.Recovering("disconnect QScreen::orientationChanged")

	if ptr.Pointer() != nil {
		C.QScreen_DisconnectOrientationChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "orientationChanged")
	}
}

//export callbackQScreenOrientationChanged
func callbackQScreenOrientationChanged(ptrName *C.char, orientation C.int) {
	defer qt.Recovering("callback QScreen::orientationChanged")

	var signal = qt.GetSignal(C.GoString(ptrName), "orientationChanged")
	if signal != nil {
		signal.(func(core.Qt__ScreenOrientation))(core.Qt__ScreenOrientation(orientation))
	}

}

func (ptr *QScreen) OrientationUpdateMask() core.Qt__ScreenOrientation {
	defer qt.Recovering("QScreen::orientationUpdateMask")

	if ptr.Pointer() != nil {
		return core.Qt__ScreenOrientation(C.QScreen_OrientationUpdateMask(ptr.Pointer()))
	}
	return 0
}

func (ptr *QScreen) ConnectPrimaryOrientationChanged(f func(orientation core.Qt__ScreenOrientation)) {
	defer qt.Recovering("connect QScreen::primaryOrientationChanged")

	if ptr.Pointer() != nil {
		C.QScreen_ConnectPrimaryOrientationChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "primaryOrientationChanged", f)
	}
}

func (ptr *QScreen) DisconnectPrimaryOrientationChanged() {
	defer qt.Recovering("disconnect QScreen::primaryOrientationChanged")

	if ptr.Pointer() != nil {
		C.QScreen_DisconnectPrimaryOrientationChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "primaryOrientationChanged")
	}
}

//export callbackQScreenPrimaryOrientationChanged
func callbackQScreenPrimaryOrientationChanged(ptrName *C.char, orientation C.int) {
	defer qt.Recovering("callback QScreen::primaryOrientationChanged")

	var signal = qt.GetSignal(C.GoString(ptrName), "primaryOrientationChanged")
	if signal != nil {
		signal.(func(core.Qt__ScreenOrientation))(core.Qt__ScreenOrientation(orientation))
	}

}

func (ptr *QScreen) SetOrientationUpdateMask(mask core.Qt__ScreenOrientation) {
	defer qt.Recovering("QScreen::setOrientationUpdateMask")

	if ptr.Pointer() != nil {
		C.QScreen_SetOrientationUpdateMask(ptr.Pointer(), C.int(mask))
	}
}

func (ptr *QScreen) DestroyQScreen() {
	defer qt.Recovering("QScreen::~QScreen")

	if ptr.Pointer() != nil {
		C.QScreen_DestroyQScreen(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
