/*
PacmanPlugin checks for installed packages by checking if the package was installed via pacman

Example Config file:

    [
        { "plugin": "pacman", "check": "firefox", "installPackage": "firefox" }
    ]
*/
package plugins

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

type Pacman struct {
	InstalledPackages string
}

func (p Pacman) New() Plugin {
		var buf bytes.Buffer

		listCmd := exec.Command("pacman", "-Q")
		listCmd.Stdout = &buf

		err := listCmd.Run()
		if err != nil {
			log.Fatal("Could not read installed packages list")
		}
		installedPackages := string(buf.Bytes())

	return Pacman{
		InstalledPackages: installedPackages,
	}
}

// Check checks if `task.CheckValue` is installed by looking at the installed Pacman packages
func (p Pacman) Check(task Task) (installed bool, err error) {
	installed = strings.Contains(p.InstalledPackages, task.CheckValue)
	return installed, err
}

// Install installs the `task.InstallPackage` via `pacman`
func (p Pacman) Install(task Task) error {
	installCmd := exec.Command("sudo", "pacman", "-S", task.InstallPackage, "--noconfirm")
	fmt.Println(installCmd.Args)
	return installCmd.Run()
}

func (p Pacman) Name() string {
	return "pacman"
}