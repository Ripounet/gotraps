package main

import "fmt"

func main() {
	a := make([]int, 3, 4)
	a[0], a[1], a[2] = 0, 1, 2

	b := append(a, 66)
	b[0] = 6
	c := append(a, 77)
	c[0] = 7
	d := append(a, 88, 99)
	d[0] = 9

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}