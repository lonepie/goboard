/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"strconv"

	"github.com/atotto/clipboard"
	"github.com/lonepie/goboard/internal/clipboardmonitor"
	"github.com/spf13/cobra"
)

// cpCmd represents the cp command
var cpCmd = &cobra.Command{
	Use:   "cp <id>",
	Short: "Copy entry ID to clipboard.",
	Long:  `Copy entry ID to clipboard. Use ls command to get entry ID.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("cp called", args)
		id, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatalln("Error: argument 'id' must be an integer")
		}
		db, err := clipboardmonitor.NewClipboardDB(dbPath)
		if err != nil {
			log.Fatalln("Error:", err)
		}
		entry, err := db.GetEntry(id)
		if err != nil {
			log.Fatalln("Error:", err)
		}
		clipboard.WriteAll(entry.Data)
		log.Println("Wrote entry to clipboard:", entry.RowID)
	},
}

func init() {
	rootCmd.AddCommand(cpCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cpCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cpCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
