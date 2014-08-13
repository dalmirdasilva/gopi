package board

import "os"

type Peripheral struct {
  MemFile *os.File
  Address int64
  Memory []byte
}
