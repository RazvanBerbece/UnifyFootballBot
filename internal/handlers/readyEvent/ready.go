package readyevent

import (
	"fmt"
	"log"

	"github.com/RazvanBerbece/UnifyFootballBot/internal/globals"
	leagues "github.com/RazvanBerbece/UnifyFootballBot/internal/handlers/readyEvent/resources"
	"github.com/RazvanBerbece/UnifyFootballBot/internal/logger"
	"github.com/bwmarrin/discordgo"
)

// Called once the Discord servers confirm a succesful connection.
func Ready(s *discordgo.Session, event *discordgo.Ready) {

	logger.LogHandlerCall("Ready", "")

	// Send static messages to required channels
	// team-assign, etc.
	sendTeamAssignMessages(s, globals.TeamAssignChannelId)

}

// Sends the team assign messages and the reaction choices to the designed channel.
func sendTeamAssignMessages(session *discordgo.Session, channelId string) {

	// Only send the message if it hasn't been sent already. Override this behaviour in staging.
	hasMessage, err := channelHasTeamAssignmentMessage(session, channelId)
	if err != nil {
		fmt.Println("Error fetching team assignment messages: ", err)
		return
	}
	if hasMessage == 1 {
		return
	}

	// Get all available leagues and reactions to post in the channel
	leagueMessages := leagues.GetLeaguesAsList()

	// Step 1 - Send the messages
	reactionIdsByMessageId := make(map[string][]string)
	for _, message := range leagueMessages {
		messageToSend := message.LeagueName + "\nReact to this message to get your roles!"
		msg, err := session.ChannelMessageSend(channelId, messageToSend)
		if err != nil {
			log.Fatal("An error occured while sending a message. Err =", err)
		}
		logger.LogSentMessage("Ready", msg.Content)
		reactionIdsByMessageId[msg.ID] = message.ReactionStrings
	}

	// Step 2 - React with the team logos
	// For each message sent above
	for k, v := range reactionIdsByMessageId {
		messageId := k
		reactionIds := v
		// Add each available reaction
		for _, emojiId := range reactionIds {
			teamLogoFormatted := fmt.Sprintf("%s", emojiId)
			err := session.MessageReactionAdd(channelId, messageId, teamLogoFormatted)
			if err != nil {
				fmt.Printf("Error adding reaction with ID %s to message. Err = %s", emojiId, err)
			} else {
				logger.LogAddReaction("Ready", messageId, teamLogoFormatted)
			}
		}
	}

}

func channelHasTeamAssignmentMessage(session *discordgo.Session, channelId string) (int, error) {

	// Fetch message history for the specified channel
	messages, err := session.ChannelMessages(channelId, 1, "", "", "")
	if err != nil {
		fmt.Println("Error fetching messages: ", err)
		return -1, err
	}

	if len(messages) > 0 {
		return 1, nil
	} else {
		return 0, nil
	}

}
