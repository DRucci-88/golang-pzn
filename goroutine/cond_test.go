package goroutine_test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var cond = sync.NewCond(&sync.Mutex{})
var group = &sync.WaitGroup{}

func WaitCondition(value int) {
	fmt.Println(value)
	cond.L.Lock()
	cond.Wait()
	fmt.Println("Done", value)
	cond.L.Unlock()
	group.Done()
}

func TestCond(t *testing.T) {
	for i := 0; i < 10; i++ {
		group.Add(1)
		go WaitCondition(i)
	}

	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(1 * time.Second)
			fmt.Println("Signal")
			cond.Signal()
		}
	}()

	group.Wait()
	fmt.Println("Completed")
}
