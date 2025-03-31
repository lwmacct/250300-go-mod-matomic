package matomic

import (
	"sync/atomic"
)

// String提供原子操作的string类型
type String struct {
	ptr atomic.Pointer[string]
}

// 确保String实现了AtomicValue接口
var _ AtomicValue[string] = (*String)(nil)

// Store以原子方式存储字符串值
func (s *String) Store(val string) {
	s.ptr.Store(&val)
}

// Load以原子方式加载字符串值
func (s *String) Load() string {
	ptr := s.ptr.Load()
	if ptr == nil {
		return ""
	}
	return *ptr
}

// Swap以原子方式交换字符串值并返回旧值
func (s *String) Swap(new string) string {
	oldPtr := s.ptr.Swap(&new)
	if oldPtr == nil {
		return ""
	}
	return *oldPtr
}

// CompareAndSwap比较并交换字符串值
func (s *String) CompareAndSwap(old, new string) bool {
	var oldPtr *string
	if old != "" {
		oldPtr = &old
	}

	// 如果当前值为nil，且old也是空字符串，则尝试设置为new
	if s.ptr.Load() == nil && old == "" {
		// 新建一个指针指向new字符串
		newVal := new
		success := s.ptr.CompareAndSwap(nil, &newVal)
		return success
	}

	newVal := new
	return s.ptr.CompareAndSwap(oldPtr, &newVal)
}
