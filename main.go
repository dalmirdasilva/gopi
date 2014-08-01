package main

import (
  "fmt"
  "github.com/dalmirdasilva/gopi/controller"
)

func main() {
  c := controller.SystemInformation{}
  c.CpuInfo()
  fmt.Println("Hello!!!")

}
