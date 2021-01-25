package main

import (
  "fmt"
)

func add(pt *int) {
  (*pt)++
  fmt.Println(*pt)
}

func sub(pt *int) {
  (*pt)--
  fmt.Println(*pt)
}
func main() {
  i := 0
  go add(&i)
  go sub(&i)

  i++
  fmt.Println()

}