package controller

import (
  "os"
  "log"
  "bufio"
  "fmt"
  "strings"
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
    var line string = scanner.Text()
    var parts []string = strings.Split(line, ":")
    fmt.Println(parts)
  }
  if err := scanner.Err(); err != nil {
    log.Fatal(err)
  }
}
