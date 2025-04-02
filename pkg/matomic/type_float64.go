package matomic

import (
	"math"
	"sync/atomic"
)

// Float64提供原子操作的float64类型
type Float64 struct {
	value uint64
}

// 确保Float64实现了NumericalAtomicValue接口
var _ AtomicValueNumerical[float64] = (*Float64)(nil)

func (f *Float64) Store(val float64) {
	atomic.StoreUint64(&f.value, math.Float64bits(val))
}

func (f *Float64) Load() float64 {
	return math.Float64frombits(atomic.LoadUint64(&f.value))
}

func (f *Float64) CompareAndSwap(old, new float64) bool {
	return atomic.CompareAndSwapUint64(
		&f.value,
		math.Float64bits(old),
		math.Float64bits(new),
	)
}

func (f *Float64) Swap(new float64) float64 {
	oldBits := atomic.SwapUint64(&f.value, math.Float64bits(new))
	return math.Float64frombits(oldBits)
}

// Add以原子方式添加delta并返回新值
// 注意：这不是单一的原子操作，而是Load+Store的组合
func (f *Float64) Add(delta float64) float64 {
	for {
		oldVal := f.Load()
		newVal := oldVal + delta
		if f.CompareAndSwap(oldVal, newVal) {
			return newVal
		}
	}
}
