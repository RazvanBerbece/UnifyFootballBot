package utils

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

// Oct 2023: Currently Go doesn't support any kind of function overloading and available workarounds are not that good,
// thus in order to account for different event message types when sending DMs in handlers, we will have to duplicate a lot of work
// (for now; this remains as a refactor task)

func SendDmToUserFromMsgReactionAdd(session *discordgo.Session, originalMessage *discordgo.MessageReactionAdd, userId string, content string) error {

	channel, err := session.UserChannelCreate(userId)
	if err != nil {
		fmt.Println("error creating channel: ", err)
		session.ChannelMessageSend(
			originalMessage.ChannelID,
			"Something went wrong while sending the DM",
		)
		return err
	}

	_, err = session.ChannelMessageSend(channel.ID, content)
	if err != nil {
		fmt.Println("error sending DM message: ", err)
		session.ChannelMessageSend(
			originalMessage.ChannelID,
			"Failed to send you a DM. "+
				"Did you disable DM in your privacy settings?",
		)
		return err
	}

	return nil

}

func SendDmToUserFromMsgReactionRemove(session *discordgo.Session, originalMessage *discordgo.MessageReactionRemove, userId string, content string) error {

	channel, err := session.UserChannelCreate(userId)
	if err != nil {
		fmt.Println("error creating channel: ", err)
		session.ChannelMessageSend(
			originalMessage.ChannelID,
			"Something went wrong while sending the DM",
		)
		return err
	}

	_, err = session.ChannelMessageSend(channel.ID, content)
	if err != nil {
		fmt.Println("error sending DM message: ", err)
		session.ChannelMessageSend(
			originalMessage.ChannelID,
			"Failed to send you a DM. "+
				"Did you disable DM in your privacy settings?",
		)
		return err
	}

	return nil

}
