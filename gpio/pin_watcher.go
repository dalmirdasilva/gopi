package gpio

type Watcher interface {
  Notify(p Pin, e Event)
}

type Event int

type eventWatcherWrapper struct {
  e Event
  w Watcher
}

const (
  FALLING Event = iota
  RISING
  CHANGE
)

var watchers = map[PinNumber][]eventWatcherWrapper{}

func AddWatcher(p Pin, w Watcher, e Event) {
  ew := eventWatcherWrapper{e, w}
  watchers[p.Number] = append(watchers[p.Number], ew)

}

