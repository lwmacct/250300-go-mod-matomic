package matomic

import "sync/atomic"

// Int64提供原子操作的int64类型
type Int64 struct {
	value int64
}

// 确保Int64实现了NumericalAtomicValue接口
var _ AtomicValueNumerical[int64] = (*Int64)(nil)

func (i *Int64) Store(val int64) {
	atomic.StoreInt64(&i.value, val)
}

func (i *Int64) Load() int64 {
	return atomic.LoadInt64(&i.value)
}

func (i *Int64) Add(delta int64) int64 {
	return atomic.AddInt64(&i.value, delta)
}

func (i *Int64) Swap(new int64) int64 {
	return atomic.SwapInt64(&i.value, new)
}

func (i *Int64) CompareAndSwap(old, new int64) bool {
	return atomic.CompareAndSwapInt64(&i.value, old, new)
}
