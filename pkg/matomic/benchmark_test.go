package matomic

import (
	"sync/atomic"
	"testing"
)

// 测试泛型原子操作 Load
func BenchmarkNumericLoad(b *testing.B) {
	var n Numeric[int64]
	n.Store(int64(100))

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = n.Load()
		}
	})
}

// 测试原生 atomic Load
func BenchmarkNativeLoad(b *testing.B) {
	var v int64 = 100

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = atomic.LoadInt64(&v)
		}
	})
}

// 测试泛型原子操作 Store
func BenchmarkNumericStore(b *testing.B) {
	var n Numeric[int64]

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			n.Store(int64(100))
		}
	})
}

// 测试原生 atomic Store
func BenchmarkNativeStore(b *testing.B) {
	var v int64

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			atomic.StoreInt64(&v, int64(100))
		}
	})
}

// 测试泛型原子操作 Add
func BenchmarkNumericAdd(b *testing.B) {
	var n Numeric[int64]

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			n.Add(1)
		}
	})
}

// 测试原生 atomic Add
func BenchmarkNativeAdd(b *testing.B) {
	var v int64

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			atomic.AddInt64(&v, 1)
		}
	})
}

// 测试泛型原子操作 CompareAndSwap
func BenchmarkNumericCAS(b *testing.B) {
	var n Numeric[int64]
	n.Store(0)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		counter := int64(0)
		for pb.Next() {
			n.CompareAndSwap(counter, counter+1)
			counter++
		}
	})
}

// 测试原生 atomic CompareAndSwap
func BenchmarkNativeCAS(b *testing.B) {
	var v int64

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		counter := int64(0)
		for pb.Next() {
			atomic.CompareAndSwapInt64(&v, counter, counter+1)
			counter++
		}
	})
}

// 测试不同类型的泛型原子操作
func BenchmarkNumericInt32(b *testing.B) {
	var n Numeric[int32]

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			n.Add(1)
		}
	})
}

func BenchmarkNumericInt64(b *testing.B) {
	var n Numeric[int64]

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			n.Add(1)
		}
	})
}

func BenchmarkNumericUint32(b *testing.B) {
	var n Numeric[uint32]

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			n.Add(1)
		}
	})
}

func BenchmarkNumericUint64(b *testing.B) {
	var n Numeric[uint64]

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			n.Add(1)
		}
	})
}

// 测试Bool类型操作
func BenchmarkBoolToggle(b *testing.B) {
	var boolVal Bool

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			boolVal.Toggle()
		}
	})
}

func BenchmarkBoolCompareAndSwap(b *testing.B) {
	var boolVal Bool

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			boolVal.CompareAndSwap(false, true)
		}
	})
}
