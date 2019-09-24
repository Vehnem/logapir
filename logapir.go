package main

import (
    "flag"
    "fmt"
	"os"
	"time"
)

func main() {

	telegramCmd := flag.NewFlagSet("telegram", flag.ExitOnError)
	slackCmd:= flag.NewFlagSet("slack", flag.ExitOnError)

	telegramToken := telegramCmd.String("token","","telgram api token")

    if len(os.Args) < 2 {
        fmt.Println("expected 'telegram' or 'slack' subcommands")
        os.Exit(1)
    }

    switch os.Args[1] {

    case "telegram":
		telegramCmd.Parse(os.Args[2:])
		fmt.Println(formatMessage(*telegramToken, "" , ""))
	case "slack":
		slackCmd.Parse(os.Args[2:])
    default:
        fmt.Println("expected 'telegram' or 'slack' subcommands")
        os.Exit(1)
    }
}

func sendTelegramAPIMessage() {
/* 	curl -X POST \
	-H 'Content-Type: application/json' \
	-d '{"chat_id": "${CHATID}", "text": "This is a test from curl", "disable_notification": true}' \
	https://api.telegram.org/bot${TOKEN}/sendMessage */
}

func sendSlackAPIMessage() {

}

func formatMessage(level string, from string, msg string) string {
	return fmt.Sprintf("%s | %s | %s | %s", time.Now().Format("01-02-2006 15:04:05.000000000"), level, from, msg)
}