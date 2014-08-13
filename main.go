package main

import (
  "github.com/dalmirdasilva/gorpi/core/board"
  "fmt"
)

func main() {
  p := board.Peripheral{}
  p.Address = 0x200000
  b := board.BoardInstance()
  mapped := b.MapPeripheral(&p)
  if mapped != nil {
    fmt.Println("Cannot map.")
  }
  fmt.Println(p.Memory)
}
