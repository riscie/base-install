/*
YayPlugin checks for installed packages by checking if the package was installed via yay

Example Config file:

    [
        { "plugin": "yay", "check": "firefox", "installPackage": "firefox" }
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


type Yay struct {
	InstalledPackages string
}

func (p Yay) New() Plugin {
		var buf bytes.Buffer

		listCmd := exec.Command("yay", "-Q")
		listCmd.Stdout = &buf

		err := listCmd.Run()
		if err != nil {
			log.Fatal("Could not read installed packages list")
		}
		installedPackages := string(buf.Bytes())

	return Yay{
		InstalledPackages: installedPackages,
	}
}

// Check checks if `task.CheckValue` is installed by looking at the installed Yay packages
func (p Yay) Check(task Task) (installed bool, err error) {
	installed = strings.Contains(p.InstalledPackages, task.CheckValue)
	return installed, err
}

// Install installs the `task.InstallPackage` via `yay`
func (p Yay) Install(task Task) error {
	installCmd := exec.Command("yay", "-S", task.InstallPackage, "--noconfirm")
	fmt.Println(installCmd.Args)
	return installCmd.Run()
}

func (p Yay) Name() string {
	return "yay"
}
