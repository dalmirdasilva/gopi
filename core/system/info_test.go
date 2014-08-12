package system_test

import (
  "github.com/dalmirdasilva/gorpi/core/system"
  "testing"
  "os"
)

var instance system.Info = system.InfoInstance()

// Not for testing, its for use the fake cpu info file
func TestAndConfigureCpuInfo(t *testing.T) {
  instance.SetCpuInfoFilePath(os.Getenv("GOPATH") + "src/github.com/dalmirdasilva/gorpi/cpuinfo.fake")
}

func TestProcessor(t *testing.T) {
  processor, _ := instance.Processor()
  if (processor != "0") {
    t.Errorf("Processor did not match. Got: %s, expected: 0", processor)
  }
}

func TestCpuFeatures(t *testing.T) {
  cpuFeatures, _ := instance.CpuFeatures()
  if (cpuFeatures[0] != "swp") {
    t.Errorf("CpuFeatures did not match. Got: %s, expected: swp", cpuFeatures[0])
  }
}

func TestRevision(t *testing.T) {
  board := instance.BoardModel()
  if (board != system.ModelBRev1) {
    t.Errorf("Revision did not match. Got: %s, expected: %s", board, system.ModelBRev1)
  }
}
