package main

import (
  "fmt"
  "github.com/dalmirdasilva/gorpi/gpio"
)

func main() {
  g := gpio.GetInstance()
  pin21 := g.NewPin(21)
  pin21.Set()
  pin22 := g.NewPin(22)
  pin22.Clear()
  pin23 := g.NewPin(23)
  pin23.Set()

  fmt.Println("done")
}
