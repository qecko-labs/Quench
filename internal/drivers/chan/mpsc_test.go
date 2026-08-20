package ch

import (
	"sync"
	"sync/atomic"
	"testing"
)

func TestMPSCConcurrentDequeue(t *testing.T) {
	q := NewMPSC(256)
	const n = 1000
	for i := 0; i < n; i++ {
		if !q.Enqueue(i) {
			t.Fatal("enqueue failed")
		}
	}
	var got atomic.Int32
	var wg sync.WaitGroup
	workers := 8
	wg.Add(workers)
	for w := 0; w < workers; w++ {
		go func() {
			defer wg.Done()
			for {
				v, ok := q.Dequeue()
				if !ok {
					break
				}
				got.Add(1)
				if v.(int) < 0 || v.(int) >= n {
					t.Errorf("unexpected value %v", v)
				}
			}
		}()
	}
	wg.Wait()
	if got.Load() != n {
		t.Fatalf("expected %d dequeued, got %d", n, got.Load())
	}
}
