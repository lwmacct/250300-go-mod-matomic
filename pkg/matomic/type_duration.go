package matomic

import (
	"sync/atomic"
	"time"
)

// Duration提供原子操作的time.Duration类型
type Duration struct {
	value int64
}

// 确保Duration实现了NumericalAtomicValue接口
var _ AtomicValueNumerical[time.Duration] = (*Duration)(nil)

// Store以原子方式存储持续时间
func (d *Duration) Store(val time.Duration) {
	atomic.StoreInt64(&d.value, int64(val))
}

// Load以原子方式加载持续时间
func (d *Duration) Load() time.Duration {
	return time.Duration(atomic.LoadInt64(&d.value))
}

// Add以原子方式添加时长并返回新值
func (d *Duration) Add(delta time.Duration) time.Duration {
	return time.Duration(atomic.AddInt64(&d.value, int64(delta)))
}

// Swap以原子方式交换时长并返回旧值
func (d *Duration) Swap(new time.Duration) time.Duration {
	return time.Duration(atomic.SwapInt64(&d.value, int64(new)))
}

// CompareAndSwap比较并交换时长
func (d *Duration) CompareAndSwap(old, new time.Duration) bool {
	return atomic.CompareAndSwapInt64(&d.value, int64(old), int64(new))
}
