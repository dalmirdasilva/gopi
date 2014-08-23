package board

import "os"

type Peripheral struct {
  Address int64
  memFile *os.File
  Memory []byte
}

func NewPeripheral(address int64) Peripheral {
  p := Peripheral{Address: address}
  return p
}

func (p *Peripheral) Open() error {
  return p.Map()
}

func (p *Peripheral) Map() error {
  board := GetInstance()
  return board.MapPeripheral(p)
}

func (p *Peripheral) Close() error {
  board := GetInstance()
  return board.UnmapPeripheral(p)
}

