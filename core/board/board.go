package board

type Board interface {
  MapPeripheral(*Peripheral) error
  UnmapPeripheral(*Peripheral) error
}

var boardContext sync.Once
var board *Board = nil

func GetInstance() *Board {
  boardContext.Do(func() {
    board = &Bcm2835{}
  })
  return board
}
