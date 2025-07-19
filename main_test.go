package main

import (
	"sync"
	"testing"
)

const ops = 1000

// ---------------- Mutex ----------------

func BenchmarkMutex(b *testing.B) {
	var mu sync.Mutex
	data := make(map[int]int)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		wg := sync.WaitGroup{}
		for j := 0; j < ops; j++ {
			wg.Add(1)
			go func(k int) {
				defer wg.Done()
				if k%2 == 0 {
					mu.Lock()
					data[k] = k
					mu.Unlock()
				} else {
					mu.Lock()
					_ = data[k]
					mu.Unlock()
				}
			}(j)
		}
		wg.Wait()
	}
}

// ---------------- RWMutex ----------------

func BenchmarkRWMutex(b *testing.B) {
	var mu sync.RWMutex
	data := make(map[int]int)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		wg := sync.WaitGroup{}
		for j := 0; j < ops; j++ {
			wg.Add(1)
			go func(k int) {
				defer wg.Done()
				if k%2 == 0 {
					mu.Lock()
					data[k] = k
					mu.Unlock()
				} else {
					mu.RLock()
					_ = data[k]
					mu.RUnlock()
				}
			}(j)
		}
		wg.Wait()
	}
}
