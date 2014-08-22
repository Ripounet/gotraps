package main

import "fmt"
import "time"

var ch chan int = make(chan int)

// Generate all multiples of 17
func generate() {
  for i := 17; i < 5000; i += 17 {
    ch <- i
    time.Sleep(1 * time.Millisecond)
  }
  close(ch)
}

func main() {
  timeout := time.After(800 * time.Millisecond)
  go generate()
  found := 0
  for {
    select {
    case i, ok := <-ch:
      if ok {
        if i%38 == 0 {
          fmt.Println(i, "is a multiple of 17 and 38")
          found++
          if found == 3 {
            // Stop as soon as third answer is found
            break
          }
        }
      } else {
        // ch was closed
        break
      }
    case <-timeout:
      // Stop if taking too long
      fmt.Println("timed out")
      break
    }
  }
  fmt.Println("The end")
}
