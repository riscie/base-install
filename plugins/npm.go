/*
NpmPlugin checks for installed `npm` packages by checking if the binary is in $PATH.
NpmPlugin installs the `npm` package (without sudo)

Example Config file:

    [
        { "plugin": "npm", "check": "ng", "installPackage": "@angular/cli" }
    ]
 */
package plugins

import (
	"bytes"
	"io"
	"log"
	"os"
	"os/exec"
)

type Npm struct{}

func (p Npm) New() Plugin { return Npm{} }

// Check checks if `task.CheckValue` is installed by looking up the binary
func (p Npm) Check(task Task) (installed bool, err error) {
	_, lerr := exec.LookPath(task.CheckValue)

	if lerr != nil {
		installed = false
	} else {
		installed = true
	}

	return installed, err
}

// Install installs the `task.InstallPackage` globally via npm (without sudo)
func (p Npm) Install(task Task) error {
	installCmd := exec.Command("npm", "install", "-g", task.InstallPackage)

	var stdBuffer bytes.Buffer
	mw := io.MultiWriter(os.Stdout, &stdBuffer)

	installCmd.Stdout = mw
	installCmd.Stderr = mw

	err := installCmd.Run()
	log.Println(stdBuffer.String())
	return err
}

func (p Npm) Name() string {
	return "npm"
}