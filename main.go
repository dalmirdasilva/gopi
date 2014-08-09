package main

import (
  "fmt"
  "github.com/dalmirdasilva/gorpi/core/system"
)

func main() {
  sysInfo := system.InfoInstance()
  fmt.Println(sysInfo.Temperature())
  fmt.Println(sysInfo.CpuInfo())
  fmt.Println(sysInfo.Memory())
}
