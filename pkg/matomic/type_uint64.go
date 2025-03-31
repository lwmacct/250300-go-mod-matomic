package matomic

import "sync/atomic"

// Uint64提供原子操作的uint64类型
type Uint64 struct {
	value uint64
}

// 确保Uint64实现了NumericalAtomicValue接口
var _ AtomicValueNumerical[uint64] = (*Uint64)(nil)

func (u *Uint64) Store(val uint64) {
	atomic.StoreUint64(&u.value, val)
}

func (u *Uint64) Load() uint64 {
	return atomic.LoadUint64(&u.value)
}

func (u *Uint64) Add(delta uint64) uint64 {
	return atomic.AddUint64(&u.value, delta)
}

func (u *Uint64) Swap(new uint64) uint64 {
	return atomic.SwapUint64(&u.value, new)
}

func (u *Uint64) CompareAndSwap(old, new uint64) bool {
	return atomic.CompareAndSwapUint64(&u.value, old, new)
}
