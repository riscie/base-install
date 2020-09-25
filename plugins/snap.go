/*
SnapPlugin checks for installed `snap` packages by checking if the binary is in $PATH
SnapPlugin installs the `snap` package using `snap`

Example Config file:

    [
        { "plugin": "snap", "check": "helm", "installPackage": "doctl", "installOption": "--classic" }
    ]
 */
package plugins

import (
	"os"
	"os/exec"
)

type Snap struct{}

func (p Snap) New() Plugin { return Snap{} }

// Check checks if `task.CheckValue` is installed by checking if the binary is in $PATH 
func (p Snap) Check(task Task) (installed bool, err error) {
	_, lerr := exec.LookPath(task.CheckValue)
	if lerr != nil {
		installed = false
	} else {
		installed = true
	}

	return installed, err
}

// Install installs the `task.InstallPackage` via `snap` with the (optional) `task.InstallOption` flag 
func (p Snap) Install(task Task) error {
	installCmd := exec.Command("sudo", "snap", "install", task.InstallPackage, task.InstallOption)
	installCmd.Stdout = os.Stdout
	return installCmd.Run()
}

func (p Snap) Name() string {
	return "snap"
}