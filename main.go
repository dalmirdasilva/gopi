package main

import (
  "github.com/dalmirdasilva/gorpi/core/board"
  "fmt"
)

func main() {
  p := board.Peripheral{0x200000}
  err := p.Open()
  if err != nil {
    fmt.Println("Cannot map.")
    return
  }
  fmt.Println(p.Memory)
}
