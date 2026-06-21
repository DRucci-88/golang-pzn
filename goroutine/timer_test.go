package goroutine_test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	timer := time.NewTimer(5 * time.Second)
	now := time.Now()
	fmt.Println(now)

	time := <-timer.C
	fmt.Println(time)
	fmt.Println(time.Sub(now))
}

func TestTimerAfter(t *testing.T) {
	channel := time.After(5 * time.Second)
	now := time.Now()
	fmt.Println(now)

	time := <-channel
	fmt.Println(time)
	fmt.Println(time.Sub(now))
}

func TestTimerAfterFunc(t *testing.T) {
	wg := &sync.WaitGroup{}
	wg.Add(1)

	now := time.Now()

	time.AfterFunc(3*time.Second, func() {
		fmt.Println("Execute after 3 second")
		wg.Done()
	})

	wg.Wait()
	fmt.Println(time.Since(now))
}
