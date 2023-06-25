/*
Copyright © 2023 JON ROGERS <LONEPIE@GMAIL.COM>
*/
package cmd

import (
	"github.com/getlantern/systray"
	"github.com/lonepie/goboard/api"
	"github.com/lonepie/goboard/assets/icon"
	"github.com/spf13/cobra"
)

var staticFilesPath string
var serverPort int
var bSystray bool

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run the web frontend, REST API and Clipboard monitor service",
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("server called")
		// go api.StartAPI(dbPath, staticFilesPath, serverPort)
		// StartMonitor()
		initSystray()
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

	serverCmd.Flags().StringVar(&staticFilesPath, "staticfiles", "", "Path to static files to serve")
	serverCmd.Flags().IntVarP(&serverPort, "port", "p", 3000, "Port for webserver to listen on")
	serverCmd.Flags().BoolVar(&bSystray, "systray", true, "Enable/disable systray icon")
}

func initSystray() {
	if bSystray {
		systray.Run(func() {
			systray.SetIcon(icon.Data)
			systray.SetTitle("goBoard")
			systray.SetTooltip("goBoard")
			mQuit := systray.AddMenuItem("Quit", "Quit")
			go func() {
				<-mQuit.ClickedCh
				systray.Quit()
			}()
			go api.StartAPI(dbPath, staticFilesPath, serverPort)
			StartMonitor()
		}, func() {
		})
	} else {
		go api.StartAPI(dbPath, staticFilesPath, serverPort)
		StartMonitor()
	}
}
