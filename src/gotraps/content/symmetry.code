package main

import "fmt"

type Point struct{ x, y int }

func main() {
	s := []Point{
		{1, 2},
		{3, 4},
		{5, 6},
	}
	// Diagonal symmetry: just swap x and y
	for _, p := range s {
		p.x, p.y = p.y, p.x
	}
	fmt.Println(s)
}
