package matomic

import "sync/atomic"

// Int32提供原子操作的int32类型
type Int32 struct {
	value int32
}

// 确保Int32实现了NumericalAtomicValue接口
var _ AtomicValueNumerical[int32] = (*Int32)(nil)

func (i *Int32) Store(val int32) {
	atomic.StoreInt32(&i.value, val)
}

func (i *Int32) Load() int32 {
	return atomic.LoadInt32(&i.value)
}

func (i *Int32) Add(delta int32) int32 {
	return atomic.AddInt32(&i.value, delta)
}

func (i *Int32) Swap(new int32) int32 {
	return atomic.SwapInt32(&i.value, new)
}

func (i *Int32) CompareAndSwap(old, new int32) bool {
	return atomic.CompareAndSwapInt32(&i.value, old, new)
}
