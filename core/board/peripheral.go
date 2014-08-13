package board

import "os"

type Peripheral struct {
  memFile *os.File
  address int64
  Memory []byte
}

func NewPeripheral(address int64) Peripheral {
  p := Peripheral{}
  p.address = address
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
