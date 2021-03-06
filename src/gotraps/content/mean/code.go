package main

import "fmt"
import "math/rand"

// The actual mean may not be an integer.
// This calculates the TRUNCATED VALUE (rounded towards zero) of the actual mean of a and b.
func integerMean(a, b int) int {
	return (a + b) / 2
}

func main() {
	lower := 0
	between := 0
	greater := 0
	noneOfTheAbove := 0

	for i := 0; i < 100; i++ {
		// Non-negative random integers
		a, b := rand.Int(), rand.Int()

		// Let a be the smaller
		if a > b {
			a, b = b, a
		}

		m := integerMean(a, b)

		if m < a {
			lower++
		} else if a <= m && m <= b {
			between++
		} else if b < m {
			greater++
		} else {
			noneOfTheAbove++
		}
	}
	fmt.Println("lower =", lower)
	fmt.Println("between =", between)
	fmt.Println("greater =", greater)
	fmt.Println("noneOfTheAbove =", noneOfTheAbove)
}
