/*
FlatpakPlugin checks for installed `flatpak` packages by checking if the package was installed via `flatpak`
FlatpakPlugin installs the `flatpak` package using `flatpak`

Example Config file:

    [
        { "plugin": "flatpak", "check": "Slack", "installPackage": "com.slack.Slack", "installOption": "flathub" }
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

var flatpakPackages string

type Flatpak struct {
	InstalledPackages string
}

func (p Flatpak) New() Plugin {
	if len(flatpakPackages) == 0 {
		var buf bytes.Buffer

		listCmd := exec.Command("flatpak", "list")
		listCmd.Stdout = &buf

		err := listCmd.Run()
		if err != nil {
			log.Fatal("Could not read installed packages list")
		}
		flatpakPackages = string(buf.Bytes())

		addRemoteCmd := exec.Command("flatpak", "remote-add", "--if-not-exists", "flathub", "https://flathub.org/repo/flathub.flatpakrepo")
		err = addRemoteCmd.Run()
	}
	return Flatpak{
		InstalledPackages: flatpakPackages,
	}
}

// Check checks if `task.CheckValue` is installed by looking at the installed flatpak packages
func (p Flatpak) Check(task Task) (installed bool, err error) {
	installed = strings.Contains(p.InstalledPackages, task.CheckValue)
	return installed, err
}

// Install installs the `task.InstallPackage` via `flatpak` from the `task.InstallOption` repository
func (p Flatpak) Install(task Task) error {
	installCmd := exec.Command("sudo", "flatpak", "install", "-y", task.InstallOption, task.InstallPackage)
	installCmd.Stdout = os.Stdout
	return installCmd.Run()

}

func (p Flatpak) Name() string {
	return "flatpack"
}