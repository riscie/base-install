/*
SnapPlugin checks for installed `snap` packages by checking if the binary is in $PATH
SnapPlugin installs the `snap` package using `snap`

Example Config file:

    [
        { "plugin": "snap", "check": "helm", "installPackage": "doctl", "installOption": "--classic" }
    ]
 */
package snap

import (
	"github.com/magbeat/base-install/plugins"
	"os"
	"os/exec"
)

type Plugin struct{}

func NewSnapPlugin() Plugin { return Plugin{} }

// Check checks if `task.CheckValue` is installed by checking if the binary is in $PATH 
func (p Plugin) Check(task plugins.Task) (installed bool, err error) {
	_, lerr := exec.LookPath(task.CheckValue)
	if lerr != nil {
		installed = false
	} else {
		installed = true
	}

	return installed, err
}

// Install installs the `task.InstallPackage` via `snap` with the (optional) `task.InstallOption` flag 
func (p Plugin) Install(task plugins.Task) error {
	installCmd := exec.Command("sudo", "snap", "install", task.InstallPackage, task.InstallOption)
	installCmd.Stdout = os.Stdout
	return installCmd.Run()
}

func (p Plugin) Name() string {
	return "snap"
}