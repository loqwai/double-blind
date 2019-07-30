package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var initCommand = &cobra.Command{
	Use:   "init",
	Short: "Initialize a config file for running a study",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("init called")
	},
}

func init() {
	rootCmd.AddCommand(initCommand)
}
