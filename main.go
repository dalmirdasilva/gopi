package main

import (
  "fmt"
  "github.com/dalmirdasilva/gorpi/core/system"
)

func main() {
  sysInfo := system.InfoInstance()
  fmt.Println(sysInfo.Revision())
}
