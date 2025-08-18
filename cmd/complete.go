package cmd

import (
	"errors"
	"fmt"
	"os"
	"encoding/csv"
	"strconv"

	"github.com/spf13/cobra"
)

func Complete(id int) {
	file, err := os.OpenFile(CSVFilename, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("error opening csv file: %s\n", err.Error())
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	csvData, err := reader.ReadAll()
	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
		return
	}

	row, err := findRowByID(csvData, id)
	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
		return
	}

	csvData[row][3] = "true"
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

func findRowByID(csvData [][]string, id int) (int, error) {
	idString := strconv.Itoa(id)
	for i, row := range csvData {
		if len(row) != 4 {
			return 0, errors.New("row contains more or less rows than expected (4)")
		}

		if row[0] == idString {
			return i, nil
		}
	}

	return 0, errors.New("id not found")
}

var completeCmd = &cobra.Command{
	Use:     "complete",
	Aliases: []string{"complete"},
	Short:   "completes a task",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		arg, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Printf("%v is not an integer\n", args[0])
			return
		}
		Complete(arg)
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)
}
