/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// NetCmd represents the net command
var NetCmd = &cobra.Command{
	Use:   "net",
	Short: "This is a net command",
	Long: `A command that make something XD`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("net called")
	},
}

func init() {
	
}
