/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"time"
	"net"
	"net/http"

	"github.com/spf13/cobra"
)

var (
	urlPath string
)

func ping(url string){
	fmt.Println("Request to: ", url)

	req, err := http.NewRequest("HEAD", "http://" + url, nil)

	if err != nil{
		fmt.Println(err)
	}

	response, err := client.Do(req)

	if err != nil{
		fmt.Println(err)
	}

	fmt.Println("Status code: ", response.StatusCode)
}

var client http.Client 

// PingCmd represents the ping command
var PingCmd = &cobra.Command{
	Use:   "ping",
	Short: "A ping command that make something",
	Long: `Command for use ping, like something how ping-pong game`,
	Run: func(cmd *cobra.Command, args []string) {

		client = http.Client{
			Transport: &http.Transport{
				DialContext: (&net.Dialer{
					Timeout:   2 * time.Second,
					KeepAlive: 30 * time.Second,
					DualStack: true,
				}).DialContext,
			},
		}

		ping(urlPath)
		
	},
}

func init() {
	PingCmd.Flags().StringVarP(&urlPath, "url", "u", "", "The url path") //var, name, short, value, help

	if err := PingCmd.MarkFlagRequired("url"); err != nil{
		fmt.Println(err)
	}
}
