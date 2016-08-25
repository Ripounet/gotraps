package main

import "fmt"

func main() {
	twoNephews := []string{"Huey", "Dewey"}
	threeNephews := append(twoNephews , "Louie")

	for i := range threeNephews {
		nephew := threeNephews[i]
		go func(){
			fmt.Println("Hello", nephew)
		}()
	}	
}