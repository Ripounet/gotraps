package main

import (
  "fmt"
  "math"
  "os"
)

func main() {
  sols := solveQuadratic(-1.0, 2.0, 3.0)
  fmt.Println("Solutions =", sols)
}

// Solves ax²+bx+c=0
func solveQuadratic(a, b, c float64) (solutions []float64) {
  defer log("Found", len(solutions), "solutions")

  discriminant := b*b - 4*a*c
  switch {
  case discriminant < 0.0:
    // No solution
    solutions = []float64{}
    return
  case discriminant == 0.0:
    // 1 solution
    sol1 := -b / (2.0 * a)
    solutions = []float64{sol1}
    return
  case discriminant > 0.0:
    // 2 solutions
    sq := math.Sqrt(discriminant)
    sol1 := (-b - sq) / (2.0 * a)
    sol2 := (-b + sq) / (2.0 * a)
    solutions = []float64{sol1, sol2}
    return
  default:
    panic("Unexpected")
  }
}

func log(values ...interface{}) {
  fmt.Fprintln(os.Stderr, values...)
}
