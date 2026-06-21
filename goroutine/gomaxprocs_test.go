package goroutine_test

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestGetGomaxprocs(t *testing.T) {
	wg := &sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(3 * time.Second)
		}()
	}

	totalCpu := runtime.NumCPU()
	fmt.Println("Total CPU ", totalCpu)

	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Total Thread ", totalThread)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("Total Goroutine ", totalGoroutine)

	wg.Done()
	fmt.Println("Completed")
}

func TestChangeThreadNumber(t *testing.T) {
	wg := &sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(3 * time.Second)
		}()
	}

	totalCpu := runtime.NumCPU()
	fmt.Println("Total CPU ", totalCpu)

	runtime.GOMAXPROCS(20)
	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Total Thread ", totalThread)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("Total Goroutine ", totalGoroutine)

	wg.Done()
	fmt.Println("Completed")
}
