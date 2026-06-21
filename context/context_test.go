package context_test

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestContext(t *testing.T) {
	background := context.Background()
	fmt.Println(background)

	todo := context.TODO()
	fmt.Println(todo)
}

func TestContextWithValue(t *testing.T) {
	ctxA := context.Background()

	ctxB := context.WithValue(ctxA, "b", "B")
	ctxC := context.WithValue(ctxA, "c", "C")

	ctxD := context.WithValue(ctxB, "d", "D")
	ctxE := context.WithValue(ctxB, "e", "E")

	ctxF := context.WithValue(ctxC, "f", "F")
	ctxG := context.WithValue(ctxF, "g", "G")

	fmt.Println(ctxA)
	fmt.Println(ctxB)
	fmt.Println(ctxC)
	fmt.Println(ctxD)
	fmt.Println(ctxE)
	fmt.Println(ctxF)
	fmt.Println(ctxG)

	fmt.Println("\nContext Get Value:")
	fmt.Println(ctxF.Value("f")) // Dapat - milik sendiri
	fmt.Println(ctxF.Value("c")) // Dapat - milik parent
	fmt.Println(ctxF.Value("b")) // Tidak - beda parent
	fmt.Println(ctxF.Value("g")) // Tidak - punya child
	fmt.Println(ctxA.Value("b")) // Tidak - beda context
}

// / Contoh Goroutine Leak (1)
func CreateCounterLeak() chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)
		counter := 0
		for {
			destination <- counter
			counter++
		}
	}()
	return destination
}

// / Contoh Goroutine Leak (1)
func TestContextWithCancelLeak(t *testing.T) {
	fmt.Println("Total Goroutine ", runtime.NumGoroutine()) // 2

	destination := CreateCounterLeak()
	for n := range destination {
		fmt.Println("Counter", n)
		if n == 10 {
			break
		}
	}

	fmt.Println("Total Goroutine ", runtime.NumGoroutine()) // 3

	/**
	Karena infinity loop didalam function CreateCounterLeak nya,
	tidak ada yang memberhentikan.

	dan akan menjadi Goroutine Zombie RAWR
	*/
}

func CreateCounterWithCancel(ctx context.Context) chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)
		counter := 0
		for {
			select {
			case <-ctx.Done():
				return
			case destination <- counter:
				counter++
			}
		}
	}()
	return destination
}

func TestContextWithCancel(t *testing.T) {
	fmt.Println("Total Goroutine ", runtime.NumGoroutine()) // 2

	// parent := context.Background()
	ctx, cancel := context.WithCancel(t.Context())

	destination := CreateCounterWithCancel(ctx)

	fmt.Println("Total Goroutine ", runtime.NumGoroutine()) // 3
	for n := range destination {
		fmt.Println("Counter", n)
		if n == 10 {
			break
		}
	}
	cancel()                    // Mengirim sinyal Cancel ke Context
	time.Sleep(3 * time.Second) // Kasih waktu jeda agar si channel ke close

	fmt.Println("Total Goroutine ", runtime.NumGoroutine()) // 2

}

func CreateCounterWithTimeout(ctx context.Context) chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)
		counter := 0
		for {
			select {
			case <-ctx.Done():
				return
			case destination <- counter:
				counter++
				time.Sleep(1 * time.Second)
			}
		}
	}()
	return destination
}

func TestContextWithTimeout(t *testing.T) {
	fmt.Println(runtime.NumGoroutine())
	parent := context.Background()
	ctx, cancel := context.WithTimeout(parent, 5*time.Second)
	defer cancel()

	destination := CreateCounterWithTimeout(ctx)
	for n := range destination {
		fmt.Println("Counter ", n)
	}

	time.Sleep(3 * time.Second) // Kasih waktu jeda agar si channel ke close

	fmt.Println(runtime.NumGoroutine())
}

func TestContextWithDeadline(t *testing.T) {
	fmt.Println(runtime.NumGoroutine())
	parent := context.Background()
	ctx, cancel := context.WithDeadline(parent, time.Now().Add(5*time.Second))
	defer cancel()

	destination := CreateCounterWithTimeout(ctx)
	for n := range destination {
		fmt.Println("Counter ", n)
	}

	time.Sleep(3 * time.Second) // Kasih waktu jeda agar si channel ke close

	fmt.Println(runtime.NumGoroutine())
}
