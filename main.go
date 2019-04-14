package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"

	random_engineer "./functions"
	"github.com/nlopes/slack"
)

func main() {
	api := slack.New(
		"",
		slack.OptionDebug(true),
		slack.OptionLog(log.New(os.Stdout, "slack-bot: ", log.Lshortfile|log.LstdFlags)),
	)

	// turn on the batch_presence_aware option
	rtm := api.NewRTM(slack.RTMOptionConnParams(url.Values{
		"batch_presence_aware": {"1"},
	}))
	go rtm.ManageConnection()

	for msg := range rtm.IncomingEvents {
		switch ev := msg.Data.(type) {

		case *slack.MessageEvent:
			if strings.HasPrefix(ev.Msg.Text, "bot nominate") || strings.HasPrefix(ev.Msg.Text, ".nominate") || strings.Contains(ev.Msg.Text, fmt.Sprintf("<@%s> nominate", rtm.GetInfo().User.ID)) {
				random_engineer.RandomEngineer(api, ev.Msg.Channel)
			}

		case *slack.RTMError:
			fmt.Printf("Error: %s\n", ev.Error())

		case *slack.InvalidAuthEvent:
			fmt.Printf("Invalid credentials")
			return

		default:
			// Ignore other events..
		}
	}
}
