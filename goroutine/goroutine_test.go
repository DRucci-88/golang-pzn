package goroutine_test

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld() {
	fmt.Println("Hello World")
}

func TestCreateGoroutine(t *testing.T) {
	go RunHelloWorld() // Running async
	fmt.Println("Ups")

	time.Sleep(1 * time.Second)
}

func DisplayNumber(number int) {
	fmt.Println("Display ", number)
}

func TestManyGoroutine(t *testing.T) {
	for i := 0; i < 500_000; i++ {
		go DisplayNumber(i)
	}

	time.Sleep(5 * time.Second)
}
