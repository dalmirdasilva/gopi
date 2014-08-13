package gpio

type PinNumber int

type Pin struct {
  Number PinNumber
}

func NewPin(number PinNumber) Pin {
  return Pin{number}
}

func (p *Pin) State() bool {
  return p.Read()
}

func (p *Pin) SetState(state bool) error {
  return p.Write(state)
}

func (p *Pin) Set() error {
  return p.SetState(HIGH)
}

func (p *Pin) Clear() error {
  return p.SetState(LOW)
}

func (p *Pin) Mode() error {
  return nil
}

func (p *Pin) SetMode(mode bool) error {
  return nil
}

func (p *Pin) Write(value bool) error {
  return nil
}

func (p *Pin) Read() bool {
  return false
}

