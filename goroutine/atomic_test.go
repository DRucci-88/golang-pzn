package goroutine_test

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

func TestAtomic(t *testing.T) {
	wg := sync.WaitGroup{}
	var x int64 = 0

	for i := 0; i < 1000; i++ {

		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				// x++
				atomic.AddInt64(&x, 1)
			}
		}()
	}

	wg.Wait()
	fmt.Println("Counter ", x)
}
