package board

import (
  "os"
  "syscall"
  "fmt"
)

const (
  IO_BASE = 0x20000000
  PAGE_SIZE = 4096
  BLOCK_SIZE = 4096
)

type Bcm2835 struct {
}

func (bmc Bcm2835) MapPeripheral(p *Peripheral) error {
  file, err := os.OpenFile("/dev/mem", os.O_RDWR | os.O_SYNC, 0400)
  if err != nil {
    fmt.Println("Cannot open /dev/mem: " + err.Error())
    return err
  }
  p.MemFile = file
  mapped, err := syscall.Mmap(int(file.Fd()), IO_BASE + p.Address, BLOCK_SIZE, syscall.PROT_READ | syscall.PROT_WRITE, syscall.MAP_SHARED)
  if err != nil {
    fmt.Println("Cannot map: " + err.Error())
    return err
  }
  p.Memory = mapped
  return nil
}

func (bmc Bcm2835) UnmapPeripheral(p *Peripheral) {
  syscall.Munmap(p.Memory)
  p.MemFile.Close()
}
