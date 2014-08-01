package controller

import (
  "os"
  "log"
  "bufio"
  "fmt"
)

const CPU_INFO_FILE_PATH = "/proc/cpuinfo"

type SystemInformation struct {

}

func (si SystemInformation) CpuInfo() string {
  parseCpuInfoFile()
  return ""
}

func parseCpuInfoFile() {
  file, err := os.Open(CPU_INFO_FILE_PATH)
  if err != nil {
    log.Fatal(err)
  }
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    fmt.Println(scanner.Text())
    fmt.Println("----------------------------")
  }
  if err := scanner.Err(); err != nil {
    log.Fatal(err)
  }
}
