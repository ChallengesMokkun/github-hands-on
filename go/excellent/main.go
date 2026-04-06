package main

const (
	Even = "even"
	Odd  = "odd"
)

func EvenOrOdd(num int) string {
	if num%2 == 0 {
		return Even
	}
	return Odd
}
