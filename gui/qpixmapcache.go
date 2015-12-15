package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QPixmapCache struct {
	ptr unsafe.Pointer
}

type QPixmapCache_ITF interface {
	QPixmapCache_PTR() *QPixmapCache
}

func (p *QPixmapCache) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QPixmapCache) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQPixmapCache(ptr QPixmapCache_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPixmapCache_PTR().Pointer()
	}
	return nil
}

func NewQPixmapCacheFromPointer(ptr unsafe.Pointer) *QPixmapCache {
	var n = new(QPixmapCache)
	n.SetPointer(ptr)
	return n
}

func (ptr *QPixmapCache) QPixmapCache_PTR() *QPixmapCache {
	return ptr
}

func QPixmapCache_CacheLimit() int {
	defer qt.Recovering("QPixmapCache::cacheLimit")

	return int(C.QPixmapCache_QPixmapCache_CacheLimit())
}

func QPixmapCache_Clear() {
	defer qt.Recovering("QPixmapCache::clear")

	C.QPixmapCache_QPixmapCache_Clear()
}

func QPixmapCache_SetCacheLimit(n int) {
	defer qt.Recovering("QPixmapCache::setCacheLimit")

	C.QPixmapCache_QPixmapCache_SetCacheLimit(C.int(n))
}
