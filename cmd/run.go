package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/loqwai/double-blind/study"
	"github.com/spf13/cobra"
)

var runCommand = &cobra.Command{
	Use:   "run",
	Short: "Run a study",
	Run:   runRun,
	Args:  cobra.MaximumNArgs(1),
}

func init() {
	rootCmd.AddCommand(runCommand)
}

func runRun(cmd *cobra.Command, args []string) {
	filename := getFilename(args)
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}

	study := study.Study{}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&study)
	if err != nil {
		log.Fatalln(err)
	}

	results, err := study.Run()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(results)
}
