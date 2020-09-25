/*
DnfPlugin checks for installed `dnf` packages by checking if the package was installed via yum / dnf 
DnfPlugin installs the `dnf` package using `dnf`

Example Config file:

    [
        { "plugin": "dnf", "check": "thunderbird", "installPackage": "thunderbird" }
    ]
 */
package plugins

import (
	"bytes"
	"log"
	"os"
	"os/exec"
	"strings"
)

var installedPackages string

type Dnf struct {
	InstalledPackages string
}

func (p Dnf) New() Plugin {
	if len(installedPackages) == 0 {
		var buf bytes.Buffer

		listCmd := exec.Command("yum", "list", "installed")
		listCmd.Stdout = &buf

		err := listCmd.Run()
		if err != nil {
			log.Fatal("Could not read installed packages list")
		}
		installedPackages = string(buf.Bytes())
	}
	return Dnf{
		InstalledPackages: installedPackages,
	}
}

// Check checks if `task.CheckValue` is installed by looking at the installed yum packages
func (p Dnf) Check(task Task) (installed bool, err error) {
	installed = strings.Contains(p.InstalledPackages, task.CheckValue)
	return installed, err
}

// Install installs the `task.InstallPackage` via `dnf` 
func (p Dnf) Install(task Task) error {
	installCmd := exec.Command("sudo", "dnf", "install", "-y", task.InstallPackage)
	installCmd.Stdout = os.Stdout
	return installCmd.Run()
}

func (p Dnf) Name() string {
	return "dnf"
}