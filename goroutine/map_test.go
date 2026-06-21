package goroutine_test

import (
	"fmt"
	"sync"
	"testing"
)

func AddToMap(
	group *sync.WaitGroup,
	data *sync.Map,
	value int,
) {
	defer group.Done()
	group.Add(1)

	data.Store(value, value)
}

func TestMap(t *testing.T) {
	data := &sync.Map{}
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go AddToMap(group, data, i)
	}

	group.Wait()
	fmt.Println("Completed Store")

	if value, ok := data.Load(1); ok {
		fmt.Println("Value ", value)
	}

	data.Range(func(key, value any) bool {
		fmt.Println(key, ":", value)
		return true
	})
}
