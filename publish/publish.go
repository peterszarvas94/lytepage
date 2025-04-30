package main

import (
	"fmt"
	"os"

	"github.com/peterszarvas94/lytepage/constants"
	"github.com/peterszarvas94/lytepage/utils"
)

func main() {
	err := utils.ZipFolder("cmd/embed", "cmd/embed.zip", []string{"node_modules", "go.work", "go.work.sum", "go.mod", "go.sum"})
	if err != nil {
		fmt.Println("Error zipping:", err.Error())
		os.Exit(1)
	}

	if err := utils.Cmd("git", "diff", "--quiet"); err != nil {
		fmt.Println("Uncommitted changes found. Commit before tagging.")
		os.Exit(1)
	}

	if err := utils.Cmd("git", "rev-parse", "--verify", "refs/tags/"+constants.Version); err == nil {
		fmt.Printf("Tag %s already exists\n", constants.Version)
		os.Exit(1)
	}

	if err := utils.Cmd("git", "tag", constants.Version, "-m", constants.Version); err != nil {
		os.Exit(1)
	}

	if err := utils.Cmd("git", "push"); err != nil {
		os.Exit(1)
	}

	if err := utils.Cmd("git", "push", "--tags"); err != nil {
		os.Exit(1)
	}

	fmt.Printf("Version %s published successfully\n", constants.Version)
}
