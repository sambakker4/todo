package cmd

import (
	"encoding/csv"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
)
const minCellWidth = 0
const tabWidth = 2
const padding = 4
const padchar = ' '
const flags = 0

func List() {
	file, err := os.Open("tasks.csv")
	if err != nil {
		fmt.Printf("error open tasks.csv: %s", err.Error())
		return
	}

	reader := csv.NewReader(file)
	csvData, err := reader.ReadAll()
	if err != nil {
		fmt.Printf("error reading data from tasks.csv: %s", err.Error())
		return
	}

	writer := tabwriter.NewWriter(os.Stdout, minCellWidth, tabWidth, padding, padchar, flags)

	for _, row := range csvData {
		currentLine := ""
		for _, item := range row {
			currentLine += item + "\t"
		}
		currentLine = currentLine[:len(currentLine) - 1]
		currentLine += "\n"
		_, err = writer.Write([]byte(currentLine))
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}
	err = writer.Flush()
	if err != nil {
		fmt.Println(err.Error())
	}
}

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"list"},
	Short:   "lists all todos",
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		List()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
