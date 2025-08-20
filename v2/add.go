// Package math provides some simple mathematical functions
package math

import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Integer | constraints.Float
}

// Add adds 2 numbers and return the sum
// That this is fun you can see here https://www.mathsisfun.com/numbers/addition.html
func Add[T Number](a T, b T) T {
	return a + b
}
