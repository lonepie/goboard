/*
Copyright © 2023 JON ROGERS <LONEPIE@GMAIL.COM>
*/
package cmd

import (
	"log"
	"strings"

	"github.com/lonepie/goboard/internal/clipboardmonitor"
	"github.com/spf13/cobra"
)

// monitorCmd represents the monitor command
var monitorCmd = &cobra.Command{
	Use:   "monitor",
	Short: "Start monitoring clipboard history",
	Long:  `Monitor clipboard history and save in sqlite database.`,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("monitor called")
		StartMonitor()
	},
}

func init() {
	rootCmd.AddCommand(monitorCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// monitorCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// monitorCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func StartMonitor() {
	monitor, err := clipboardmonitor.NewClipboardMonitor(dbPath)
	if err != nil {
		log.Fatalln("Error:", err)
	}
	go monitor.MonitorClipboard()
	log.Println("Monitoring Clipboard...")
	for entry := range monitor.EntryChan {
		log.Println("New clip:", strings.TrimSpace(entry.Data))
	}
}
