package board

import "sync"

type Board interface {
  MapPeripheral(p *Peripheral) error
  UnmapPeripheral(p *Peripheral) error
}

var boardContext sync.Once
var board Board

func GetInstance() Board {
  boardContext.Do(func() {
    board = Bcm2835{}
  })
  return board
}
