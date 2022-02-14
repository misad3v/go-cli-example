/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/misad3v/go-cli-example.git/web"
	"github.com/spf13/cobra"
)

// =================================
var port string

// bmakeCmd represents the bmake command
var bmakeCmd = &cobra.Command{
	Use:   "bmake",
	Short: "Create a new AIOps Broker resource",
	Long:  "This is a feature that makes the AIOps Broker extensible and easy to use",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("bmake called")
		server := web.Server{Port: port}
		server.Route()
	},
}

// =================================

func init() {
	rootCmd.AddCommand(bmakeCmd)
	bmakeCmd.Flags().StringVarP(&port, "port", "p", ":4040", "Port to HTTP Server")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// bmakeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// bmakeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
