package main

import "fmt"

type Point struct{ x, y int }

// Diagonal symmetry: just swap x and y
func (p Point) reflect(){
	p.x, p.y = p.y, p.x
}

func main() {
	s := []Point{
		{1, 2},
		{3, 4},
		{5, 6},
	}
	s[0].reflect()
	s[1].reflect()
		
	fmt.Println(s)
}
