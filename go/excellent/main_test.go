package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEvenOrMain(t *testing.T) {
	result := EvenOrOdd(10)
	assert.Equal(t, Even, result)
}

func TestGenerateFizzBuzz(t *testing.T) {
	num := 100
	result := GenerateFizzBuzz(num)

	fizzBuzzNum := 0
	for _, i := range result {
		if i == FizzBuzz {
			fizzBuzzNum++
		}
	}
	assert.Equal(t, num/15, fizzBuzzNum)
}
