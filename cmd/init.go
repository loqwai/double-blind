package cmd

import (
	"encoding/json"
	"log"
	"os"

	"github.com/loqwai/double-blind/study"
	"github.com/spf13/cobra"
)

var initCommand = &cobra.Command{
	Use:   "init",
	Short: "Initialize a config file for running a study",
	Run:   runInit,
	Args:  cobra.MaximumNArgs(1),
}

func init() {
	initCommand.Flags().StringP("name", "n", "Study", "Name of the study")
	rootCmd.AddCommand(initCommand)
}

func runInit(cmd *cobra.Command, args []string) {
	name, err := cmd.Flags().GetString("name")
	if err != nil {
		log.Fatalln(err.Error())
	}

	filename := getFilename(args)
	file, err := os.Create(filename)
	if os.IsExist(err) {
		log.Fatalf("File at %v already exists, refusing to overwrite", filename)
	}
	if err != nil {
		log.Fatalln(err.Error())
	}

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	encoder.Encode(&study.Study{
		Name: name,
		Groups: []study.Group{
			{Name: "Control Group", Command: "exit 0"},
			{Name: "Experiment Group A", Command: "exit 0"},
			{Name: "Experiment Group B", Command: "exit 0"},
		},
	})
}

func getFilename(args []string) string {
	if len(args) > 0 {
		return args[0]
	}
	return "study.json"
}
