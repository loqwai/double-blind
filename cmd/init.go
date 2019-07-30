package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var initCommand = &cobra.Command{
	Use:   "init",
	Short: "Initialize a config file for running a study",
	Run:   runInit,
	Args:  cobra.MaximumNArgs(1),
}

func init() {
	rootCmd.AddCommand(initCommand)
}

func runInit(cmd *cobra.Command, args []string) {
	filename := getFilename(args)
	file, err := os.Create(filename)
	if os.IsExist(err) {
		log.Fatalf("File at %v already exists, refusing to overwrite", filename)
	}
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println("file", file)
}

func getFilename(args []string) string {
	if len(args) > 0 {
		return args[0]
	}
	return "study.json"
}
