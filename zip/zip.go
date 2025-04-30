package main

import (
	"fmt"
	"os"

	"github.com/peterszarvas94/lytepage/utils"
)

func main() {
	err := utils.ZipFolder("cmd/embed", "cmd/embed.zip", []string{"node_modules", "go.work", "go.work.sum", "go.mod", "go.sum"})
	if err != nil {
		fmt.Println("Error zipping:", err.Error())
		os.Exit(1)
	}
}
