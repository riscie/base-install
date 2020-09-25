package plugins

type Plugin interface {
	Check(task Task) (installed bool, err error)
	Install(task Task) error
	Name() string
	New() Plugin
}