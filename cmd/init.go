package cmd

import (
	"fmt"

	"github.com/peterszarvas94/lt2/utils"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init [folder?]",
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
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
