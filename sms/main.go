package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"

	"golang.org/x/crypto/ssh/terminal"
)

var message, number string

func init() {
	if !terminal.IsTerminal(0) {
		pipe, _ := ioutil.ReadAll(os.Stdin)
		message = string(pipe)
		number = os.Getenv("SMS_DEFAULT_NUMBER")
	} else {
		args := os.Args[1:]
		switch len(args) {
		case 0:
			if message == "" {
				message = os.Getenv("SMS_DEFAULT_MESSAGE")
			}
			number = os.Getenv("SMS_DEFAULT_NUMBER")
		case 1:
			message = args[0]
			number = os.Getenv("SMS_DEFAULT_NUMBER")
		case 2:
			message = args[0]
			number = args[1]
		default:
			fmt.Println("Unexpected arguments: ", args[2:])
			showUsage()
		}

		for _, arg := range args {
			if arg == "-h" || arg == "--help" {
				showUsage()
			}
		}
	}
}

func main() {
	resp, err := http.PostForm("http://textbelt.com/text",
		url.Values{"number": {number}, "message": {message}})
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != 200 {
		fmt.Println(resp.Status)
	}
}

func showUsage() {
	fmt.Println("usage: sms [message] [number]")
	os.Exit(0)
}
