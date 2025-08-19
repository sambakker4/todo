/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
)

const TimeFormat = time.RFC3339
const CSVFilename = "tasks.csv"

var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "A simple todo cli application",
	Long:  `todo is an simple CLI application that stores your things you want to do`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.todo.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func init() {
	_, err := os.Stat(CSVFilename)
	if err == nil {
		return
	}

	file, err := os.Create(CSVFilename)
	if err != nil {
		fmt.Printf("error creating tasks.csv on startup: %s\n", err.Error())
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	err = writer.Write([]string{"ID", "Description", "Created At", "Is Complete"})
	if err != nil {
		fmt.Printf("error creating tasks.csv on startup: %s\n", err.Error())
		return
	}
}
