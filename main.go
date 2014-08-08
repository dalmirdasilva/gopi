package main

import (
  "fmt"
  "github.com/dalmirdasilva/gorpi/controller"
)

func main() {
  c := controller.SystemInformation{}
  c.CpuInfo()
  fmt.Println("Hello!!!")

}
