/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/atotto/clipboard"
	fuzzyfinder "github.com/ktr0731/go-fuzzyfinder"
	"github.com/lonepie/goboard/internal/clipboardmonitor"
	"github.com/spf13/cobra"
)

// fzfCmd represents the fzf command
var fzfCmd = &cobra.Command{
	Use:   "fzf",
	Short: "Fuzzy-find clipboard history",
	Long:  `Use fuzzyfinder to search clipboard history and select an item to copy to clipboard.`,
	Run: func(cmd *cobra.Command, args []string) {
		// log.Println("fzf called")
		fzf()
	},
}

func init() {
	rootCmd.AddCommand(fzfCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fzfCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fzfCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func fzf() {
	db, err := clipboardmonitor.NewClipboardDB(dbPath)
	if err != nil {
		log.Println("Error: ", err)
	}
	entries, _ := db.ReadEntries()
	index, err := fuzzyfinder.Find(entries, func(i int) string {
		return fmt.Sprintf("[%v] %s", entries[i].RowID, strings.TrimSpace(entries[i].Data))
	})
	if err != nil {
		log.Printf("Error selecting clipboard entry: %v\n", err)
		return
	}
	if index >= 0 && index < len(entries) {
		log.Println("Selected item:", index)
		clipboard.WriteAll(entries[index].Data)
		// entries[index].WriteToClipboard()
	}
}
