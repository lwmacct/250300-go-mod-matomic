package matomic

import (
	"sync/atomic"
	"time"
	"unsafe"
)

// Numeric 泛型原子操作结构体，支持 int32, int64, uint32, uint64 类型
type Numeric[T int32 | int64 | uint32 | uint64] struct {
	value T
}

// 确保Numeric实现了NumericalAtomicValue接口
var _ AtomicValueNumerical[int32] = (*Numeric[int32])(nil)
var _ AtomicValueNumerical[int64] = (*Numeric[int64])(nil)
var _ AtomicValueNumerical[uint32] = (*Numeric[uint32])(nil)
var _ AtomicValueNumerical[uint64] = (*Numeric[uint64])(nil)

// Store 原子性地存储值
func (n *Numeric[T]) Store(val T) {
	switch any(val).(type) {
	case int32:
		atomic.StoreInt32((*int32)(unsafe.Pointer(&n.value)), int32(val))
	case int64:
		atomic.StoreInt64((*int64)(unsafe.Pointer(&n.value)), int64(val))
	case uint32:
		atomic.StoreUint32((*uint32)(unsafe.Pointer(&n.value)), uint32(val))
	case uint64:
		atomic.StoreUint64((*uint64)(unsafe.Pointer(&n.value)), uint64(val))
	}
}

// Load 原子性地加载值
func (n *Numeric[T]) Load() T {
	switch any(n.value).(type) {
	case int32:
		return T(atomic.LoadInt32((*int32)(unsafe.Pointer(&n.value))))
	case int64:
		return T(atomic.LoadInt64((*int64)(unsafe.Pointer(&n.value))))
	case uint32:
		return T(atomic.LoadUint32((*uint32)(unsafe.Pointer(&n.value))))
	case uint64:
		return T(atomic.LoadUint64((*uint64)(unsafe.Pointer(&n.value))))
	default:
		return n.value
	}
}

// Add 原子性地添加值并返回新值
func (n *Numeric[T]) Add(delta T) T {
	switch any(n.value).(type) {
	case int32:
		return T(atomic.AddInt32((*int32)(unsafe.Pointer(&n.value)), int32(delta)))
	case int64:
		return T(atomic.AddInt64((*int64)(unsafe.Pointer(&n.value)), int64(delta)))
	case uint32:
		return T(atomic.AddUint32((*uint32)(unsafe.Pointer(&n.value)), uint32(delta)))
	case uint64:
		return T(atomic.AddUint64((*uint64)(unsafe.Pointer(&n.value)), uint64(delta)))
	default:
		return n.value
	}
}

// Swap 原子性地交换值并返回旧值
func (n *Numeric[T]) Swap(new T) T {
	switch any(new).(type) {
	case int32:
		return T(atomic.SwapInt32((*int32)(unsafe.Pointer(&n.value)), int32(new)))
	case int64:
		return T(atomic.SwapInt64((*int64)(unsafe.Pointer(&n.value)), int64(new)))
	case uint32:
		return T(atomic.SwapUint32((*uint32)(unsafe.Pointer(&n.value)), uint32(new)))
	case uint64:
		return T(atomic.SwapUint64((*uint64)(unsafe.Pointer(&n.value)), uint64(new)))
	default:
		return n.value
	}
}

// CompareAndSwap 原子性地比较并交换值
func (n *Numeric[T]) CompareAndSwap(old, new T) bool {
	switch any(old).(type) {
	case int32:
		return atomic.CompareAndSwapInt32((*int32)(unsafe.Pointer(&n.value)), int32(old), int32(new))
	case int64:
		return atomic.CompareAndSwapInt64((*int64)(unsafe.Pointer(&n.value)), int64(old), int64(new))
	case uint32:
		return atomic.CompareAndSwapUint32((*uint32)(unsafe.Pointer(&n.value)), uint32(old), uint32(new))
	case uint64:
		return atomic.CompareAndSwapUint64((*uint64)(unsafe.Pointer(&n.value)), uint64(old), uint64(new))
	default:
		return false
	}
}

// NewInt64 创建一个新的原子int64类型，可选择初始值
func NewInt64(initialValue ...int64) *Int64 {
	i := &Int64{}
	if len(initialValue) > 0 {
		i.Store(initialValue[0])
	}
	return i
}

// NewInt32 创建一个新的原子int32类型，可选择初始值
func NewInt32(initialValue ...int32) *Int32 {
	i := &Int32{}
	if len(initialValue) > 0 {
		i.Store(initialValue[0])
	}
	return i
}

// NewUint64 创建一个新的原子uint64类型，可选择初始值
func NewUint64(initialValue ...uint64) *Uint64 {
	u := &Uint64{}
	if len(initialValue) > 0 {
		u.Store(initialValue[0])
	}
	return u
}

// NewUint32 创建一个新的原子uint32类型，可选择初始值
func NewUint32(initialValue ...uint32) *Uint32 {
	u := &Uint32{}
	if len(initialValue) > 0 {
		u.Store(initialValue[0])
	}
	return u
}

// NewFloat64 创建一个新的原子float64类型，可选择初始值
func NewFloat64(initialValue ...float64) *Float64 {
	f := &Float64{}
	if len(initialValue) > 0 {
		f.Store(initialValue[0])
	}
	return f
}

// NewBool 创建一个新的原子bool类型，可选择初始值
func NewBool(initialValue ...bool) *Bool {
	b := &Bool{}
	if len(initialValue) > 0 {
		b.Store(initialValue[0])
	}
	return b
}

// NewString 创建一个新的原子string类型，可选择初始值
func NewString(initialValue ...string) *String {
	s := &String{}
	if len(initialValue) > 0 {
		s.Store(initialValue[0])
	}
	return s
}

// NewDuration 创建一个新的原子Duration类型，可选择初始值
func NewDuration(initialValue ...time.Duration) *Duration {
	d := &Duration{}
	if len(initialValue) > 0 {
		d.Store(initialValue[0])
	}
	return d
}

// NewPointer 创建一个新的原子Pointer类型，可选择初始值
func NewPointer[T any](initialValue ...*T) *Pointer[T] {
	p := &Pointer[T]{}
	if len(initialValue) > 0 {
		p.Store(initialValue[0])
	}
	return p
}

// NewNumeric 创建一个新的泛型原子数值类型，可选择初始值
func NewNumeric[T int32 | int64 | uint32 | uint64](initialValue ...T) *Numeric[T] {
	n := &Numeric[T]{}
	if len(initialValue) > 0 {
		n.Store(initialValue[0])
	}
	return n
}
