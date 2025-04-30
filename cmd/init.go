package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/peterszarvas94/lytepage/utils"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init [folder]? [--name name]?",
	Short: "Initialize a new project at the given folder (default is pwd)",
	Args:  cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {
		var targetDir string
		if len(args) > 0 {
			targetDir = args[0]
		}

		if targetDir == "" || targetDir == "." {
			targetDir = "./"
		}
		err := utils.UnzipFromEmbed(embedZip, targetDir)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		err = os.Chdir(targetDir)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		// go mod init
		command := exec.Command("go", "mod", "init", "scaffhold")
		command.Dir = "."
		err = command.Run()
		if err != nil {
			fmt.Println("Error initializing:", err.Error())
			os.Exit(1)
		}

		// rename
		name, err := cmd.Flags().GetString("name")
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		targetDirName := targetDir
		if targetDirName == "./" {
			pwd, err := os.Getwd()
			if err != nil {
				fmt.Println(err.Error())
			}
			targetDirName = filepath.Base(pwd)
		}

		if name == "" {
			name = targetDirName
		}

		err = utils.ReplaceAllString(".", "scaffhold", name)
		if err != nil {
			fmt.Printf("Error replacing %s with %s in dir %s: %s\n", "scaffhold", name, targetDir, err.Error())
			os.Exit(1)
		}

		// install deps
		dependencies := []string{
			"github.com/a-h/templ@v0.3.857",
			"github.com/peterszarvas94/lytepage@v0.1.0",
		}

		for _, dep := range dependencies {
			command := exec.Command("go", "get", "-u", dep)
			command.Dir = "."
			err = command.Run()
			if err != nil {
				fmt.Printf("Error installing %s: %v\n", dep, err.Error())
				os.Exit(1)
			}
		}

		// tidy
		command = exec.Command("go", "mod", "tidy")
		command.Dir = "."
		err = command.Run()
		if err != nil {
			fmt.Printf("Error tidying: %v\n", err.Error())
			os.Exit(1)
		}

		// generate
		command = exec.Command("templ", "generate")
		command.Dir = "."
		err = command.Run()
		if err != nil {
			fmt.Printf("Error generating: %v\n", err.Error())
			os.Exit(1)
		}
	},
}

func init() {
	initCmd.Flags().StringP("name", "n", "", "Project name, defaults to folder name")
	rootCmd.AddCommand(initCmd)
}
