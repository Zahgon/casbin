package persist

type Watcher interface {
	SetUpdateCallback(func(string)) error

	Update() error

	Close()
}
