package main

import "fmt"

func main() {
	m := map[string]int{
		"zero":  0,
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
	}
	for name, value := range m {
		fmt.Println(name, "is english for number", value)
	}
}
