/*
NpmPlugin checks for installed  npm  packages by checking if the binary is available.

Example Config file:
```
[
    { "plugin": "npm", "check": "ng", "installPackage": "@angular/cli" }
]
```
 */
package plugins

import (
	"bytes"
	"io"
	"log"
	"os"
	"os/exec"
)

type NpmPlugin struct{}

func NewNpmPlugin() NpmPlugin { return NpmPlugin{} }

// Check checks if `task.CheckValue` is installed by looking up the binary
func (p NpmPlugin) Check(task Task) (installed bool, err error) {
	_, lerr := exec.LookPath(task.CheckValue)
	if lerr != nil {
		installed = false
	} else {
		installed = true
	}

	return installed, err
}

// Install installs the `task.InstallPackage` globally via npm (without sudo)
func (p NpmPlugin) Install(task Task) (success bool, err error) {
	success = false
	installCmd := exec.Command("npm", "install", "-g", task.InstallPackage)

	var stdBuffer bytes.Buffer
	mw := io.MultiWriter(os.Stdout, &stdBuffer)

	installCmd.Stdout = mw
	installCmd.Stderr = mw

	err = installCmd.Run()

	log.Println(stdBuffer.String())

	if err == nil {
		success = true
	}
	return success, err
}
