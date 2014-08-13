package gpio

type Pin struct {
  Number int
}

func (p Pin) SetState(state bool) {
}

func (p Pin) State() bool {
  return true
}

