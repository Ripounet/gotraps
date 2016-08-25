package main

import "fmt"

func main() {
	str := "Café-society"

	for i, c := range str {
		fmt.Printf("Character #%d is %c \n", i, c)
	}
}
