package gpio

import (
  "fmt"
  "github.com/dalmirdasilva/gorpi/core/board"
  "sync"
)

const (
  ADDRESS = 0x200000
  INPUT = false
  OUTPUT = true
  LOW = false
  HIGH = true
)

const (
  BCM2835_FSEL_MASK int = 0x000007 // Function select bits mask
  BCM2835_GPFSEL0 = 0x000000 // GPIO Function Select 0
  BCM2835_GPFSEL1 = 0x000004 // GPIO Function Select 1
  BCM2835_GPFSEL2 = 0x000008 // GPIO Function Select 2
  BCM2835_GPFSEL3 = 0x00000c // GPIO Function Select 3
  BCM2835_GPFSEL4 = 0x000010 // GPIO Function Select 4
  BCM2835_GPFSEL5 = 0x000014 // GPIO Function Select 5
  BCM2835_GPSET0 = 0x00001c // GPIO Pin Output Set 0
  BCM2835_GPSET1 = 0x000020 // GPIO Pin Output Set 1
  BCM2835_GPCLR0 = 0x000028 // GPIO Pin Output Clear 0
  BCM2835_GPCLR1 = 0x00002c // GPIO Pin Output Clear 1
  BCM2835_GPLEV0 = 0x000034 // GPIO Pin Level 0
  BCM2835_GPLEV1 = 0x000038 // GPIO Pin Level 1
  BCM2835_GPEDS0 = 0x000040 // GPIO Pin Event Detect Status 0
  BCM2835_GPEDS1 = 0x000044 // GPIO Pin Event Detect Status 1
  BCM2835_GPREN0 = 0x00004c // GPIO Pin Rising Edge Detect Enable 0
  BCM2835_GPREN1 = 0x000050 // GPIO Pin Rising Edge Detect Enable 1
  BCM2835_GPFEN0 = 0x000048 // GPIO Pin Falling Edge Detect Enable 0
  BCM2835_GPFEN1 = 0x00005c // GPIO Pin Falling Edge Detect Enable 1
  BCM2835_GPHEN0 = 0x000064 // GPIO Pin High Detect Enable 0
  BCM2835_GPHEN1 = 0x000068 // GPIO Pin High Detect Enable 1
  BCM2835_GPLEN0 = 0x000070 // GPIO Pin Low Detect Enable 0
  BCM2835_GPLEN1 = 0x000074 // GPIO Pin Low Detect Enable 1
  BCM2835_GPAREN0 = 0x00007c // GPIO Pin Async. Rising Edge Detect 0
  BCM2835_GPAREN1 = 0x000080 // GPIO Pin Async. Rising Edge Detect 1
  BCM2835_GPAFEN0 = 0x000088 // GPIO Pin Async. Falling Edge Detect 0
  BCM2835_GPAFEN1 = 0x00008c // GPIO Pin Async. Falling Edge Detect 1
  BCM2835_GPPUD = 0x000094 // GPIO Pin Pull-up/down Enable
  BCM2835_GPPUDCLK0 = 0x000098 // GPIO Pin Pull-up/down Enable Clock 0
  BCM2835_GPPUDCLK1 = 0x00009c // GPIO Pin Pull-up/down Enable Clock 1
)

type Gpio struct {
  peripheral board.Peripheral
}

var gpioContext sync.Once
var gpio Gpio

func GetInstance() Gpio {
  gpioContext.Do(func () {
    p := board.NewPeripheral(ADDRESS)
    err := p.Open()
    if err != nil {
      fmt.Println(err)
    }
    gpio = Gpio{p}
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
  pinNumber := int(pin.Number)
  address := BCM2835_GPFSEL0 / 4 + (pinNumber / 10)
  shift := (pinNumber % 10) * 3
  mask := uint(BCM2835_FSEL_MASK << uint(shift))
  var value uint
  if mode == OUTPUT {
    value = 0x000001
  } else {
    value = 0x000000
  }
  g.configureBits(address, value, mask);
  return nil
}

func (g *Gpio) DigitalRead(pin Pin) error {
  return nil
}

func (g *Gpio) SetPin(pin Pin) error {
  return g.DigitalWrite(pin, HIGH)
}

func (g *Gpio) ClearPin(pin Pin) error {
  return g.DigitalWrite(pin, LOW)
}

func (g *Gpio) DigitalWrite(pin Pin, value bool) error {
  pinNumber := int(pin.Number)
  var position int
  if value == HIGH {
    position = BCM2835_GPSET0
  } else {
    position = BCM2835_GPCLR0
  }
  var address int = int(position) / 4 + pinNumber / 32
  shift := pinNumber % 32
  g.safeWrite(address, uint(0x000001) << uint(shift))
  return nil
}

func (g *Gpio) configureBits(address int, value uint, mask uint) {
  v := g.safeRead(address)
  v = (v & ^mask) | (value & mask)
  g.safeWrite(address, v);
}

func (g *Gpio) safeRead(address int) uint {
  var result uint = 0
  b0 := g.peripheral.Memory[address + 0]
  b1 := g.peripheral.Memory[address + 1]
  b2 := g.peripheral.Memory[address + 2]
  b3 := g.peripheral.Memory[address + 3]
  result = uint(b3)
  result <<= 8
  result |= uint(b2)
  result <<= 8
  result |= uint(b1)
  result <<= 8
  result |= uint(b0)
  return result
}

func (g *Gpio) safeWrite(address int, value uint) {
  g.peripheral.Memory[address + 0] = byte(value & 0xff)
  g.peripheral.Memory[address + 1] = byte((value >> 8) & 0xff)
  g.peripheral.Memory[address + 2] = byte((value >> 16) & 0xff)
  g.peripheral.Memory[address + 3] = byte((value >> 24) & 0xff)
}
