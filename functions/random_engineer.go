package functions

import (
	"fmt"
	"math/rand"

	slack "github.com/nlopes/slack"
)

//RandomEngineer is randomizing an engineer from the team
func RandomEngineer(api *slack.Client, channel string) {
	channeldetails, err := api.GetChannelInfo(channel)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	username := ""
	breakout := 30
	for {
		nominateduser := channeldetails.Members[rand.Intn(len(channeldetails.Members))]

		userdetails, err := api.GetUserInfo(nominateduser)
		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}

		if breakout < 1 {
			fmt.Printf("Could not find a non-bot or app user in 30 loops\n")
			break
		}
		breakout--

		username = userdetails.Profile.RealNameNormalized

		if !userdetails.IsAppUser && !userdetails.IsBot {
			break
		}

	}

	channelID, timestamp, err := api.PostMessage(channel, slack.MsgOptionText(fmt.Sprintf("_Nominating someone from the channel_\nThe winner is:\n ---- *%s* ----\n", username), false))
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	msgRef := slack.NewRefToMessage(channelID, timestamp)

	if err = api.AddReaction("crown", msgRef); err != nil {
		fmt.Printf("Error adding reaction: %s\n", err)
		return
	}
}
