/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	urlPath string
)

// PingCmd represents the ping command
var PingCmd = &cobra.Command{
	Use:   "ping",
	Short: "A ping command that make something",
	Long: `Command for use ping, like something how ping-pong game`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(urlPath)

		fmt.Println("ping called")
	},
}

func init() {
	PingCmd.Flags().StringVarP(&urlPath, "url", "u", "", "The url path") //var, name, short, value, help
}
