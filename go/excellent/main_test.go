package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEvenOrMain(t *testing.T) {
	result := EvenOrOdd(10)
	assert.Equal(t, result, Even)
}
