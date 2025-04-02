package matomic

import "sync/atomic"

// Uint32提供原子操作的uint32类型
type Uint32 struct {
	value uint32
}

// 确保Uint32实现了NumericalAtomicValue接口
var _ AtomicValueNumerical[uint32] = (*Uint32)(nil)

func (u *Uint32) Store(val uint32) {
	atomic.StoreUint32(&u.value, val)
}

func (u *Uint32) Load() uint32 {
	return atomic.LoadUint32(&u.value)
}

func (u *Uint32) Add(delta uint32) uint32 {
	return atomic.AddUint32(&u.value, delta)
}

func (u *Uint32) Swap(new uint32) uint32 {
	return atomic.SwapUint32(&u.value, new)
}

func (u *Uint32) CompareAndSwap(old, new uint32) bool {
	return atomic.CompareAndSwapUint32(&u.value, old, new)
}
