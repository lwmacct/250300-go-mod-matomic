package matomic

// AtomicValue 定义了原子值操作的通用接口，适用于所有具有基本原子操作的类型
type AtomicValue[T any] interface {
	// Store 原子地存储值
	Store(value T)

	// Load 原子地加载值
	Load() T

	// Swap 原子地交换值并返回旧值
	Swap(new T) T

	// CompareAndSwap 原子地比较并交换值
	CompareAndSwap(old, new T) bool
}

// AtomicValueNumerical 定义了数值类型的原子值，扩展了AtomicValue接口并添加了Add方法
type AtomicValueNumerical[T any] interface {
	AtomicValue[T]

	// Add 原子地将delta添加到当前值并返回新值
	Add(delta T) T
}

// AtomicValueBoolean 定义了布尔类型的原子值，扩展了AtomicValue接口并添加了Toggle方法
type AtomicValueBoolean interface {
	AtomicValue[bool]

	// Toggle 切换布尔值并返回新值
	Toggle() bool
}

// AtomicValuePointer 定义了指针类型的原子值
type AtomicValuePointer[T any] interface {
	AtomicValue[*T]
}
