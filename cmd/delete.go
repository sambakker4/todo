package cmd

import (
	"encoding/csv"
	"fmt"
	"os"
	"slices"
	"strconv"

	"github.com/spf13/cobra"
)

func Delete(id int) {
	file, err := os.OpenFile(CSVFilename, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("csv file failed to open")
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	csvData, err := reader.ReadAll()
	if err != nil {
		fmt.Printf("error reading csv data: %s\n", err.Error())
		return
	}

	row, err := findRowByID(csvData, id)
	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
		return
	}

	if row == len(csvData) - 1 {
		csvData = csvData[:row]	
	} else {
		csvData = slices.Concat(csvData[:row], csvData[row+1:])
	}
	updateIDs(csvData)
	file.Close()

	file, err = os.Create(CSVFilename)	
	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	err = writer.WriteAll(csvData)
	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
		return
	}

	writer.Flush()
}

func updateIDs(csvData [][]string) {
	for i, row := range csvData {
		if i == 0 {
			continue
		}

		row[0] = strconv.Itoa(i)
	}
}

var deleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"delete", "d"},
	Short:   "deletes a task",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		arg, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Printf("%v is not an integer\n", arg)
			return
		}
		Delete(arg)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
