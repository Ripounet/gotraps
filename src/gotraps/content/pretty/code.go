package main

import "fmt"

const PLUS = '+'
const IMG_UNIT = 'i'

func main() {
	real := 4
	img := 3
	complex := real + PLUS + img + IMG_UNIT
	fmt.Println(complex)
}
