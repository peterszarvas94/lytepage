package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/peterszarvas94/lt2/utils"
)

func main() {
	err := utils.ZipFolder("cmd/embed", "cmd/embed.zip", []string{"node_modules"})
	if err != nil {
		fmt.Println("Error zipping:", err.Error())
		os.Exit(1)
	}

	cmd := exec.Command("go", "install", "./...")
	cmd.Dir = "."
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error installing:", err.Error())
		os.Exit(1)
	}

	fmt.Println("Zipping and installation complete!")
}
