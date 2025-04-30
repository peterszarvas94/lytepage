package cmd

import (
	"fmt"

	"github.com/peterszarvas94/lytepage/constants"
	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "lytepage version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("lytepage")
		fmt.Println(constants.Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
