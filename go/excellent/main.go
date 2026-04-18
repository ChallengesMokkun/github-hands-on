package main

import (
	"math/rand"
)

const (
	Even     = "even"
	Odd      = "odd"
	FizzBuzz = "FizzBuzz"
	Fizz     = "Fizz"
	Buzz     = "Buzz"

	SuperLucky   = "大吉"
	MiddleLucky  = "中吉"
	LittleLucky  = "小吉"
	Lucky        = "吉"
	LuckyUnlucky = "吉凶"
	UnLucky      = "凶"
)

var Fortune = []string{
	SuperLucky,
	MiddleLucky,
	LittleLucky,
	Lucky,
	LuckyUnlucky,
	UnLucky,
}

func EvenOrOdd(num int) string {
	if num%2 == 0 {
		return Even
	}
	return Odd
}

func GenerateFizzBuzz(num int) []string {
	slice := make([]string, 0)
	for i := 1; i <= num; i++ {
		if i%3 == 0 && i%5 == 0 {
			slice = append(slice, FizzBuzz)
		} else if i%3 == 0 {
			slice = append(slice, Fizz)
		} else if i%5 == 0 {
			slice = append(slice, Buzz)
		}
	}
	return slice
}

func TellFortune() string {
	return Fortune[rand.Intn(len(Fortune))]
}
