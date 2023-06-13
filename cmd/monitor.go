/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"strings"

	"github.com/getlantern/systray"
	"github.com/lonepie/goboard/assets/icon"
	"github.com/lonepie/goboard/internal/clipboardmonitor"
	"github.com/spf13/cobra"
)

var bSystray bool

// monitorCmd represents the monitor command
var monitorCmd = &cobra.Command{
	Use:   "monitor",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("monitor called")
		if bSystray {
			systray.Run(func() {
				systray.SetIcon(icon.Data)
				systray.SetTitle("goboard")
				systray.SetTooltip("goboard")
				mQuit := systray.AddMenuItem("Quit", "Quit")
				go func() {
					<-mQuit.ClickedCh
					systray.Quit()
				}()
				startMonitor()
			}, func() {
			})
		} else {
			startMonitor()
		}
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
	monitorCmd.Flags().BoolVar(&bSystray, "systray", false, "Enable systray")
}

func startMonitor() {
	monitor, err := clipboardmonitor.NewClipboardMonitor(dbPath)
	if err != nil {
		log.Fatalln("Error:", err)
	}
	log.Println("Monitoring Clipboard...")
	for entry := range monitor.EntryChan {
		log.Println("New clip:", strings.TrimSpace(entry.Data))
	}
}
