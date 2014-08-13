package board

import (
  "os"
  "syscall"
  "log"
)

const (
  IO_BASE = 0x20000000
  PAGE_SIZE = 4096
  BLOCK_SIZE = 4096

)

type Bcm2835 struct {
}

// You need sudo to perform this operation
func (bmc *Bcm2835) MapPeripheral(p *Peripheral) error {
  file, err := os.OpenFile("/dev/mem", os.O_RDWR | os.O_SYNC, 0400)
  if err != nil {
    log.Fatalln("Cannot open /dev/mem: " + err.Error())
    return err
  }
  p.memFile = file
  mapped, err := syscall.Mmap(int(file.Fd()), IO_BASE + p.address, BLOCK_SIZE, syscall.PROT_READ | syscall.PROT_WRITE, syscall.MAP_SHARED)
  if err != nil {
    log.Fatalln("Cannot map: " + err.Error())
    return err
  }
  p.Memory = mapped
  return nil
}

// You need sudo to perform this operation
func (bmc *Bcm2835) UnmapPeripheral(p *Peripheral) error {
  err := syscall.Munmap(p.Memory)
  if err != nil {
    return err
  }
  return p.memFile.Close()
}
