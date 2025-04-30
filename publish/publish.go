package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/peterszarvas94/lytepage/constants"
	"github.com/peterszarvas94/lytepage/utils"
)

func main() {
	err := utils.ZipFolder("cmd/embed", "cmd/embed.zip", []string{"node_modules", "go.work", "go.work.sum", "go.mod", "go.sum"})
	if err != nil {
		fmt.Println("Error zipping:", err.Error())
		os.Exit(1)
	}

	command := exec.Command("git", "tag", constants.Version, "-m", constants.Version)
	err = command.Run()
	if err != nil {
		fmt.Println("Error initializing:", err.Error())
		os.Exit(1)
	}

	command = exec.Command("git", "push")
	err = command.Run()
	if err != nil {
		fmt.Println("Error initializing:", err.Error())
		os.Exit(1)
	}

	command = exec.Command("git", "push", "--tags")
	err = command.Run()
	if err != nil {
		fmt.Println("Error initializing:", err.Error())
		os.Exit(1)
	}
}
