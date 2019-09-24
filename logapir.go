package main

import (
	"bytes"
	// "encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {

	telegramCmd := flag.NewFlagSet("telegram", flag.ExitOnError)
	slackCmd := flag.NewFlagSet("slack", flag.ExitOnError)

	telegramToken := telegramCmd.String("token", "", "telgram api token")
	telegramID := telegramCmd.String("id", "", "telgram chat id")

	if len(os.Args) < 2 {
		fmt.Println("expected 'telegram' or 'slack' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {

	case "telegram":
		telegramCmd.Parse(os.Args[2:])
		sendTelegramAPIMessage(*telegramToken, *telegramID, telegramCmd.Args()[0])

	case "slack":
		slackCmd.Parse(os.Args[2:])
	default:
		fmt.Println("expected 'telegram' or 'slack' subcommands")
		os.Exit(1)
	}
}

func sendTelegramAPIMessage(token string, id string, msg string) {
	/*
		curl -X POST \
			-H 'Content-Type: application/json' \
			-d '{ "chat_id": "${ID}", "text": "test", "disable_notification": "true"}' \
			https://api.telegram.org/bot${TOKEN}/sendMessage
	*/

	url := "https://api.telegram.org/bot" + token + "/sendMessage"
	reqBody := "{ \"chat_id\": \"" + id + "\", \"text\": \"" + msg + "\", \"disable_notification\": \"true\"}"

	res, _ := http.Post(url, "application/json", bytes.NewBuffer([]byte(reqBody)))
	resBody, _ := ioutil.ReadAll(res.Body)

	fmt.Println(string(resBody))
}

func sendSlackAPIMessage() {

}

func formatMessage(level string, from string, msg string) string {
	return fmt.Sprintf("%s | %s | %s | %s", time.Now().Format("01-02-2006 15:04:05.000000000"), level, from, msg)
}
