package goroutine_test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	group := &sync.WaitGroup{}
	pool := sync.Pool{
		New: func() any {
			return "New"
		},
	}

	pool.Put("Le")
	pool.Put("Rucco")
	pool.Put("Dewa")

	for i := 0; i < 10; i++ {
		go func() {
			defer group.Done()
			group.Add(1)
			data := pool.Get() // Ambil datanya
			fmt.Println(data)
			time.Sleep(1 * time.Second)
			pool.Put(data) // Balikin lagi datanya
		}()
	}

	group.Wait()
	fmt.Println("Completed 1")
	// time.Sleep(5 * time.Second)
	fmt.Println("Completed 2")
}
