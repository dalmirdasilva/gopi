package util

import "os/exec"
import "strings"

func Execute(cmd string, args ...string) ([]string, error) {
  var result []string
  output, err := exec.Command(cmd, args...).Output()
  if (err == nil) {
    result = strings.Split(string(output), "\n")
  }
  return result, err
}
