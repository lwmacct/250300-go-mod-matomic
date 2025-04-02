package matomic

import "sync/atomic"

// 原子布尔值，用于高效布尔值操作
type Bool struct {
	value int32
}

// 确保Bool实现了BooleanAtomicValue接口
var _ AtomicValueBoolean = (*Bool)(nil)

func (b *Bool) Set(newValue bool) {
	if newValue {
		atomic.StoreInt32(&b.value, 1)
	} else {
		atomic.StoreInt32(&b.value, 0)
	}
}

// Store是Set的别名，用于满足AtomicValue接口
func (b *Bool) Store(newValue bool) {
	b.Set(newValue)
}

func (b *Bool) Load() bool {
	return atomic.LoadInt32(&b.value) != 0
}

// CompareAndSwap 比较并交换
func (b *Bool) CompareAndSwap(old, new bool) bool {
	var oldInt, newInt int32
	if old {
		oldInt = 1
	}
	if new {
		newInt = 1
	}
	return atomic.CompareAndSwapInt32(&b.value, oldInt, newInt)
}

// Swap将值设置为新值并返回旧值
func (b *Bool) Swap(newValue bool) bool {
	var newInt int32 = 0
	if newValue {
		newInt = 1
	}
	oldInt := atomic.SwapInt32(&b.value, newInt)
	return oldInt != 0
}

// Toggle切换布尔值的状态并返回新值
func (b *Bool) Toggle() bool {
	for {
		oldValue := atomic.LoadInt32(&b.value)
		newValue := int32(1 - oldValue)
		if atomic.CompareAndSwapInt32(&b.value, oldValue, newValue) {
			return newValue != 0
		}
	}
}
