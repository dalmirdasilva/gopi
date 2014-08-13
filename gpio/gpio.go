package gpio

import "github.com/dalmirdasilva/gorpi/core/board"

const (
  ADDRESS = 0x200000
  INPUT = false
  OUTPUT = true
  LOW = false
  HIGH = true
)

type Gpio struct {
  peripheral board.Peripheral
}

var gpioContext sync.Once
var gpio *Gpio = nil

func GetInstance() *Gpio {
  gpioContext.Do(func () {
    p := board.NewPeripheral(ADDRESS)
    p.Open()
    gpio = &Gpio{p}
  })
  return gpio
}

func (g *Gpio) Close() error {
  return g.peripheral.Close()
}

func (g *Gpio) NewPin(number int) Pin {
  return NewPin(PinNumber(number))
}

func (g *Gpio) PinMode(pin Pin, mode bool) error {
  return nil
}

func (g *Gpio) DigitalWrite(pin Pin, value bool) error {
  return nil
}

func (g *Gpio) DigitalRead(pin Pin) error {
  return nil
}

func (g *Gpio) SetPin(pin Pin) error {
  return nil
}

func (g *Gpio) ClearPin(pin Pin) error {
  return nil
}
