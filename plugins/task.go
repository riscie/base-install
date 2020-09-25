package plugins

type Task struct {
	Plugin         string    `json:"plugin"`
	CheckType      CheckType `json:"checkType"`
	CheckValue     string    `json:"check"`
	InstallPackage string    `json:"installPackage"`
	InstallOption  string    `json:"installOption"`
	Commands       []string  `json:"commands"`
}

type CheckType string

const (
	Binary    CheckType = "bin"
	Directory CheckType = "dir"
	Yum       CheckType = "yum"
)
