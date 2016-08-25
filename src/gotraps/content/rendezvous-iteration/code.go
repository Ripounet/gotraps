package main

import "fmt"
import "sync"

func main() {
	twoNephews := []string{"Huey", "Dewey"}
	threeNephews := append(twoNephews, "Louie")

	var wg sync.WaitGroup

	for _, nephew := range threeNephews {
		wg.Add(1)
		go func() {
			fmt.Println("Hello", nephew)
			wg.Done()
		}()
	}

	// Wait for all greetings to complete.
	wg.Wait()
}
