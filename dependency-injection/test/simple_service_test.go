package test

import (
	"fmt"
	"restful-api/simple"
	"testing"
)

func TestSimpleService(t *testing.T) {
	simpleService, err := simple.IntializeService()
	fmt.Println(err)
	fmt.Println(simpleService)
}
