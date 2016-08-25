package main

import "fmt"
import "sync"

// Can't declare a "const" array, so this is a "var"
var nephews = [...]string{"Huey", "Dewey", "Louie"}

func main() {
	var wg sync.WaitGroup

	for _, nephew := range nephews {
		wg.Add(1)
		go greet(nephew, wg)
	}

	// Wait for all greetings to complete.
	wg.Wait()
}

func greet(name string, wg sync.WaitGroup) {
	fmt.Println("Hello", name)
	wg.Done()
}
