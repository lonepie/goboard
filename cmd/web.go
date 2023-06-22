/*
Copyright © 2023 JON ROGERS <LONEPIE@GMAIL.COM>
*/
package cmd

import (
	"github.com/lonepie/goboard/api"
	"github.com/spf13/cobra"
)

// webCmd represents the web command
var webCmd = &cobra.Command{
	Use:   "web",
	Short: "Start web API server",
	Long:  `Start web API server listening on port 3000 by default.`,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("web called")
		api.StartAPI(dbPath, staticFilesPath, serverPort)
	},
}

var staticFilesPath string
var serverPort int

func init() {
	rootCmd.AddCommand(webCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// webCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// webCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	webCmd.Flags().StringVar(&staticFilesPath, "staticfiles", "", "Path to static files to serve")
	webCmd.Flags().IntVarP(&serverPort, "port", "p", 3000, "Port for webserver to listen on")
}
