package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"

	"github.com/magbeat/base-install/plugins"
)

func main() {
	var basePath string
	if len(os.Args) > 1 {
		basePath = os.Args[1]
	} else {
		currentUser, err := user.Current()
		if err != nil {
			log.Fatal(err)
		}
		basePath = currentUser.HomeDir
	}

	configDir := basePath + "/install/config"

	tasks := parseTasks(configDir)

	processTasks(tasks)
}

func parseTasks(configDir string) []plugins.Task {
	configFiles, err := ioutil.ReadDir(configDir)
	if err != nil {
		log.Fatal(err)
	}

	var tasks []plugins.Task
	for _, configFile := range configFiles {
		jsonFile, err := os.Open(configDir + "/" + configFile.Name())
		fmt.Println("Loading config file ", jsonFile.Name())
		if err != nil {
			fmt.Println("Error while loading config file ", jsonFile.Name())
			log.Fatal(err)
		}

		byteValue, err := ioutil.ReadAll(jsonFile)
		var tmpTasks []plugins.Task
		err = json.Unmarshal(byteValue, &tmpTasks)
		if err != nil {
			fmt.Println("Error while unmarshalling config file ", jsonFile.Name())
			log.Fatal(err)
		}
		tasks = append(tasks, tmpTasks...)
	}
	return tasks
}
func processTasks(tasks []plugins.Task) {
	ps := map[string]plugins.Plugin{
		plugins.Dnf{}.Name():     plugins.Dnf{},
		plugins.Snap{}.Name():    plugins.Snap{},
		plugins.Flatpak{}.Name(): plugins.Flatpak{},
		plugins.Custom{}.Name():  plugins.Custom{},
		plugins.Npm{}.Name():     plugins.Npm{},
		plugins.Yay{}.Name():     plugins.Yay{},
		plugins.Pacman{}.Name():  plugins.Pacman{},
	}

	for _, task := range tasks {
		fmt.Printf("[%s] Checking %s: ", task.Plugin, task.CheckValue)
		installed := false
		var err error

		if plugin, ok := ps[task.Plugin]; ok {
			fmt.Println("HERE")
			plugin = plugin.New()
			fmt.Println(plugin.Name())
			installed, err = plugin.Check(task)

			if err != nil {
				log.Printf("Error while checking %s with plugin %s", task.CheckValue, task.Plugin)
				log.Printf(err.Error())
			}

			if installed {
				fmt.Println(" ... installed")
			} else {
				fmt.Println(" ... installing")
				err = plugin.Install(task)
				if err != nil {
					log.Fatalf("Error while installing %s with plugin %s\n%s", task.CheckValue, task.Plugin, err.Error())
				}
				fmt.Println("... done")
			}
		} else {
			fmt.Println("nope")
		}
	}
}
