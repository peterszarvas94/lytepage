package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/peterszarvas94/lytepage/pkg/utils"
	"github.com/peterszarvas94/lytepage/pkg/version"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:     "init [folder]? [--name name]?",
	Aliases: []string{"i"},
	Short:   "Initialize a new project at the given folder (default is pwd)",
	Args:    cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {
		var folder string
		if len(args) > 0 {
			folder = args[0]
		}

		if folder == "" || folder == "." {
			folder = "./"
		}

		fmt.Printf("Target dir is: %s\n", folder)

		targetDirFullPath, err := filepath.Abs(folder)
		utils.CheckError(err, "Can not get target directory full path")

		fmt.Printf("Target dir full path is: %s\n", targetDirFullPath)

		// get project name

		nameFlag, err := cmd.Flags().GetString("name")
		utils.CheckError(err, "Can not parse flag \"name\"")

		var projectName string
		if nameFlag != "" {
			projectName = nameFlag
		} else if folder == "./" {
			pwd, err := os.Getwd()
			utils.CheckError(err, "Can not get pwd")

			projectName = filepath.Base(pwd)
		} else {
			projectName = filepath.Base(folder)
		}

		fmt.Printf("Project name: %s\n", projectName)

		tmp, err := os.MkdirTemp("", "lytepage-template")
		utils.CheckError(err, "Error creating temp dir")

		fmt.Printf("Temp dir created: %s\n", tmp)

		// clone repo

		err = utils.Cmd("git", "clone", "https://github.com/peterszarvas94/lytepage.git", tmp)
		utils.CheckError(err, "Can not clone repo")

		err = os.Chdir(tmp)
		utils.CheckError(err, "Can change directory to tmp")

		// checkout version

		_, err = utils.CmdWithOutput("git", "checkout", version.Version)
		utils.CheckError(err, "Can checkout version")

		fmt.Printf("Checked out version: %s\n", version.Version)

		err = utils.CopyDir(filepath.Join(tmp, "scaffhold"), targetDirFullPath)
		utils.CheckError(err, "Can not copy dir")

		// rename

		err = utils.ReplaceAllString(targetDirFullPath, "scaffhold", projectName)
		utils.CheckError(err, "Error replacing strings")

		fmt.Printf("Renamed %s to %s\n", "scaffhold", projectName)

		// templ

		err = utils.Cmd("go", "install", "github.com/a-h/templ/cmd/templ@v0.3.865")
		utils.CheckError(err, "Error installing templ cli")

		// cd

		err = os.Chdir(targetDirFullPath)
		utils.CheckError(err, "Error changing directory")

		// git

		err = utils.Cmd("git", "init")
		utils.CheckError(err, "Error initializing git")

		// tidy

		err = utils.Cmd("go", "mod", "tidy")
		utils.CheckError(err, "Error tidying")

		// vendor

		err = utils.Cmd("go", "mod", "vendor")
		utils.CheckError(err, "Error vendoring")

		// generate

		err = utils.Cmd("templ", "generate")
		utils.CheckError(err, "Error generating with templ")
	},
}

func init() {
	initCmd.Flags().StringP("name", "n", "", "Project name, defaults to folder name")
	rootCmd.AddCommand(initCmd)
}
