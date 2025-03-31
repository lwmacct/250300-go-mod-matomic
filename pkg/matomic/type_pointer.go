package matomic

import "sync/atomic"

// Pointer提供原子操作的泛型指针类型
type Pointer[T any] struct {
	ptr atomic.Pointer[T]
}

// 确保Pointer实现了PointerAtomicValue接口
var _ AtomicValuePointer[int] = (*Pointer[int])(nil)

// Store以原子方式存储指针
func (p *Pointer[T]) Store(val *T) {
	p.ptr.Store(val)
}

// Load以原子方式加载指针
func (p *Pointer[T]) Load() *T {
	return p.ptr.Load()
}

// Swap以原子方式交换指针并返回旧指针
func (p *Pointer[T]) Swap(new *T) *T {
	return p.ptr.Swap(new)
}

// CompareAndSwap比较并交换指针
func (p *Pointer[T]) CompareAndSwap(old, new *T) bool {
	return p.ptr.CompareAndSwap(old, new)
}
