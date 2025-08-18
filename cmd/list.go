package cmd

import (
	"encoding/csv"
	"fmt"
	"os"
	"text/tabwriter"
	"time"

	"github.com/mergestat/timediff"

	"github.com/spf13/cobra"
)

const minCellWidth = 0
const tabWidth = 2
const padding = 4
const padchar = ' '
const flags = 0

func List() {
	file, err := os.Open(CSVFilename)
	if err != nil {
		fmt.Printf("error opening csv file: %s\n", err.Error())
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	csvData, err := reader.ReadAll()
	if err != nil {
		fmt.Printf("error reading data from csv file: %s\n", err.Error())
		return
	}

	writer := tabwriter.NewWriter(os.Stdout, minCellWidth, tabWidth, padding, padchar, flags)

	for i, row := range csvData {
		currentLine := ""
		for j, item := range row {
			if item == "false" {
				currentLine += "\t"
				continue
			}
			if item == "true" {
				currentLine += "✓" + "\t"
				continue
			}
			if j == 2 && i != 0 {
				time, err := time.Parse(TimeFormat, item)
				if err != nil {
					fmt.Printf("error parsing time in csv %s", err.Error())
					return
				}
				currentLine += timediff.TimeDiff(time) + "\t"
				continue
			}
			currentLine += item + "\t"
		}
		currentLine = currentLine[:len(currentLine)-1]
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
