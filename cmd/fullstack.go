/*
Copyright © 2023 JON ROGERS <LONEPIE@GMAIL.COM>
*/
package cmd

import (
	"fmt"

	"github.com/lonepie/goboard/api"
	"github.com/spf13/cobra"
)

// fullstackCmd represents the fullstack command
var fullstackCmd = &cobra.Command{
	Use:   "fullstack",
	Short: "Runs both the clipboard monitor and the web server",
	Long:  `Runs both the clipboard monitor, the REST API server and the React frontend.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("fullstack called")
		// fmt.Printf("dbPath: %v\nstaticFilesPath: %v\nserverPort: %v\nargs: %v\ncmd: %v", dbPath, staticFilesPath, serverPort, args, cmd)
		go api.StartAPI(dbPath, staticFilesPath, serverPort)
		// webCmd.Run(cmd, args)
		monitorCmd.Run(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(fullstackCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fullstackCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fullstackCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
