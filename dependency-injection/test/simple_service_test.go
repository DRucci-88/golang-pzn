package test

import (
	"fmt"
	"restful-api/simple"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleService(t *testing.T) {
	simpleService, err := simple.IntializeService(true)
	fmt.Println(err)
	fmt.Println(simpleService)

	assert.Nil(t, simpleService)
	assert.NotNil(t, err)
}
func TestSimpleServiceSuccess(t *testing.T) {
	simpleService, err := simple.IntializeService(false)
	fmt.Println(err)
	fmt.Println(simpleService)

	assert.Nil(t, err)
	assert.NotNil(t, simpleService)
}
