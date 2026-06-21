package goroutine_test

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestTicker(t *testing.T) {
	ticker := time.NewTicker(time.Second)
	done := make(chan struct{})

	go func() {
		time.Sleep(5 * time.Second)
		ticker.Stop()
		close(done)
	}()

	for {
		select {
		case now := <-ticker.C:
			fmt.Println(now)

		case <-done:
			fmt.Println("Ticker stopped")
			return
		}
	}
}

func TestTickerWithContext(t *testing.T) {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)
	defer cancel()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case now := <-ticker.C:
			fmt.Println(now)

		case <-ctx.Done():
			fmt.Println("Finished")
			return
		}
	}
}
