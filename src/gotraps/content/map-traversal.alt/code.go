package main

import "fmt"

func main() {
	m := make(map[string]int, 16)
	m["zero"] = 0
	m["one"] = 1
	m["two"] = 2
	m["three"] = 3
	m["four"] = 4
	m["five"] = 5

	for name, value := range m {
		fmt.Println(name, "is english for number", value)
	}
}
