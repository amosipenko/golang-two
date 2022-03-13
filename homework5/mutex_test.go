package main

import (
	"math/rand"
	"sync"
	"testing"
	"time"
)

type Numbers struct {
	mp   map[float64]struct{}
	mx   sync.Mutex
	rwmx sync.RWMutex
}

func NewSet() *Numbers {
	return &Numbers{mp: map[float64]struct{}{}}
}

func (n *Numbers) AddUsingMutex(i float64) {
	n.mx.Lock()
	n.mp[i] = struct{}{}
	n.mx.Unlock()
}

func (n *Numbers) ReadUsingMutex(i float64) bool {
	n.mx.Lock()
	defer n.mx.Unlock()
	_, ok := n.mp[i]
	return ok
}

func (n *Numbers) AddUsingRWMutex(i float64) {
	n.rwmx.Lock()
	n.mp[i] = struct{}{}
	n.rwmx.Unlock()
}

func (n *Numbers) ReadUsingRWMutex(i float64) bool {
	n.rwmx.RLock()
	defer n.rwmx.RUnlock()
	_, ok := n.mp[i]
	return ok
}

var (
	numberOfWrites = []int{10, 50, 90}
	steps          = 100
)

func BenchmarkMutex_10(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	writes := 10
	reads := steps * (100 - writes) / 100
	var set = NewSet()
	b.Run("", func(b *testing.B) {
		b.SetParallelism(writes)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				set.AddUsingMutex(rand.Float64())
			}
		})
	})
	b.Run("", func(b *testing.B) {
		b.SetParallelism(reads)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				set.ReadUsingMutex(rand.Float64())
			}
		})
	})
}

func BenchmarkRWMutex_10(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	writes := 10
	reads := steps * (100 - writes) / 100
	var set = NewSet()
	b.Run("", func(b *testing.B) {
		b.SetParallelism(writes)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				set.AddUsingRWMutex(rand.Float64())
			}
		})
	})
	b.Run("", func(b *testing.B) {
		b.SetParallelism(reads)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				set.ReadUsingRWMutex(rand.Float64())
			}
		})
	})
}

func BenchmarkMutex_50(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	writes := 50
	reads := steps * (100 - writes) / 100
	var set = NewSet()
	b.Run("", func(b *testing.B) {
		b.SetParallelism(writes)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				set.AddUsingMutex(rand.Float64())
			}
		})
	})
	b.Run("", func(b *testing.B) {
		b.SetParallelism(reads)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				set.ReadUsingMutex(rand.Float64())
			}
		})
	})
}

func BenchmarkRWMutex_50(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	writes := 50
	reads := steps * (100 - writes) / 100
	var set = NewSet()
	b.Run("", func(b *testing.B) {
		b.SetParallelism(writes)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				set.AddUsingRWMutex(rand.Float64())
			}
		})
	})
	b.Run("", func(b *testing.B) {
		b.SetParallelism(reads)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				set.ReadUsingRWMutex(rand.Float64())
			}
		})
	})
}

func BenchmarkMutex_90(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	writes := 90
	reads := steps * (100 - writes) / 100
	var set = NewSet()
	b.Run("", func(b *testing.B) {
		b.SetParallelism(writes)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				set.AddUsingMutex(rand.Float64())
			}
		})
	})
	b.Run("", func(b *testing.B) {
		b.SetParallelism(reads)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				set.ReadUsingMutex(rand.Float64())
			}
		})
	})
}

func BenchmarkRWMutex_90(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	writes := 90
	reads := steps * (100 - writes) / 100
	var set = NewSet()
	b.Run("", func(b *testing.B) {
		b.SetParallelism(writes)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				set.AddUsingRWMutex(rand.Float64())
			}
		})
	})
	b.Run("", func(b *testing.B) {
		b.SetParallelism(reads)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				set.ReadUsingRWMutex(rand.Float64())
			}
		})
	})
}
