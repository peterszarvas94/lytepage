/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	_ "embed"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "lytepage",
	Short: "lytepage - static site generator",
	Run: func(cmd *cobra.Command, args []string) {
		versionFlag, err := cmd.Flags().GetBool("version")
		if err != nil {
			fmt.Printf("Error parsing flags: %v", err)
			os.Exit(1)
		}

		if versionFlag {
			versionCmd.Run(cmd, args)
			os.Exit(0)
		}

		fmt.Println("Welcome to lytepage!\nTo get started, run \"lytepage init my-app\", or \"lytepage --help\"")

		err = cmd.Usage()
		if err != nil {
			fmt.Printf("Error printing usage: %v", err)
			os.Exit(1)
		}
	},
}

func Execute() {
	rootCmd.Flags().BoolP("version", "v", false, "Print GOAT version")

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
