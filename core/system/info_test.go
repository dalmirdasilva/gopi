package system_test

import (
  "github.com/dalmirdasilva/gorpi/core/system"
  "testing"
  "fmt"
)

const REVISION = system.ModelBRev1

var instance system.Info = system.InfoInstance()

func TestRevision(t *testing.T) {
  if (instance.Revision() != REVISION) {

  }
}

func BenchmarkHello(t *testing.B) {
  for i := 0; i < t.N; i++ {
    fmt.Sprintf("hello")
  }
}
