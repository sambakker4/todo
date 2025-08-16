package cmd

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

func Add(task string) {
	file, err := os.OpenFile("tasks.csv", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("tasks.csv failed to open")
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	csvData, err := reader.ReadAll()
	if err != nil {
		fmt.Println("error reading tasks.csv")
		return
	}

	var id int64 = 1
	if len(csvData) > 1 {
		lastId, err := strconv.ParseInt(csvData[len(csvData)-1][0], 10, 64)
		if err == nil {
			id = lastId + 1
		}
	}
	timeStamp := time.Now()
	isComplete := "false"

	writer := csv.NewWriter(file)
	err = writer.Write([]string{strconv.FormatInt(id, 10), task, timeStamp.Format(time.RFC3339), isComplete})
	if err != nil {
		fmt.Printf("error writing to tasks.csv: %s", err.Error())
		return
	}

	writer.Flush()
}

var addCmd = &cobra.Command{
	Use:     "add",
	Aliases: []string{"add"},
	Short:   "adds a task",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		Add(args[0])
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
