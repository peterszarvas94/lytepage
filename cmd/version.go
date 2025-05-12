package cmd

import (
	"fmt"

	"github.com/peterszarvas94/lytepage/pkg/version"
	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "lytepage version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("lytepage")
		fmt.Println(version.Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
