package concurrent

import (
	"sync"
	"sync/atomic"
	"testing"
)

func TestStoreAndLoad(t *testing.T) {
	StoreAndLoad()
}

type manager struct {
	sync.RWMutex
	agents int
}

// 20000000	       107 ns/op	       0 B/op	       0 allocs/op
func BenchmarkManagerLock(b *testing.B) {
	m := &manager{}

	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			m.Lock()
			m.agents = 100
			m.Unlock()
		}
	})
}

// 30000000	        35.1 ns/op	       0 B/op	       0 allocs/op
func BenchmarkManagerRLock(b *testing.B) {
	m := manager{agents: 100}

	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			m.RLock()
			_ = m.agents
			m.RUnlock()
		}
	})
}

// 50000000	        52.6 ns/op	      32 B/op	       1 allocs/op
func BenchmarkManagerAtomicValueStore(b *testing.B) {
	var managerVal atomic.Value
	m := manager{agents: 100}

	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			managerVal.Store(m)
		}
	})
}

// 2000000000	         1.24 ns/op	       0 B/op	       0 allocs/op
func BenchmarkManagerAtomicValueLoad(b *testing.B) {
	var managerVal atomic.Value
	managerVal.Store(&manager{agents: 100})

	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = managerVal.Load().(*manager)
		}
	})
}
