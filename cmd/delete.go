package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

func Delete(id int) {

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
