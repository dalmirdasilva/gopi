package system

import (
  "github.com/dalmirdasilva/gorpi/core/system"
  "testing"
  "fmt"
)

const REVISION = system.ModelBRev1

var instance system.Info = system.InfoInstance()

func TestRevision(b *testing.T) {
  fmt.Println(instance.Revision())
}

func BenchmarkHello(b *testing.B) {
  for i := 0; i < b.N; i++ {
    fmt.Sprintf("hello")
  }
}
