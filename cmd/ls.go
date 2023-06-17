/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/lonepie/goboard/internal/clipboardmonitor"
	"github.com/spf13/cobra"
)

// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List clipboard history entries",
	Long:  "List clipboard history entries with ID number to pass to cp command.",
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("ls called")
		ls()
	},
}

func ls() {
	db, err := clipboardmonitor.NewClipboardDB("clipboard.db")
	if err != nil {
		log.Println("Error: ", err)
	}
	entries, _ := db.ReadEntries()
	for _, entry := range entries {
		fmt.Println(entry.RowID, strings.TrimSpace(entry.Data))
	}
}

func init() {
	rootCmd.AddCommand(lsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// lsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// lsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
