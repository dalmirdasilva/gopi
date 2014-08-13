package board

type Board interface {
  MapPeripheral(*Peripheral) error
  UnmapPeripheral(*Peripheral)
}

func BoardInstance() Board {
  return Bcm2835{}
}
