package readyevent

import (
	"fmt"
	"log"

	"github.com/RazvanBerbece/UnifyFootballBot/internal/globals"
	"github.com/RazvanBerbece/UnifyFootballBot/internal/logger"
	"github.com/RazvanBerbece/UnifyFootballBot/internal/utils"
	"github.com/bwmarrin/discordgo"

	apiFootballClient "github.com/RazvanBerbece/UnifyFootballBot/internal/apis/api-football"
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

	// Get all available leagues and teams
	// and reaction image data to post in the channel
	leagues := apiFootballClient.GetLeaguesForCountry("Romania", 1)
	for index, league := range leagues {
		teams := apiFootballClient.GetTeamsForLeague(league.Id, 2023, league.CountryName)
		leagues[index].Teams = teams
	}

	reactionIdsByMessageId := make(map[string][]string) // keep track of the reactions under each league message
	fmt.Println("Creating Guild reactions with teams from the given leagues...")
	for _, league := range leagues {

		// Step 1 - Send the league specific messages
		messageToSend := fmt.Sprintf("## %s", league.Name) +
			"\nReact to this message to pick your favourite teams from this league!"
		msg, err := session.ChannelMessageSend(channelId, messageToSend)
		if err != nil {
			log.Fatal("An error occured while sending a message. Err =", err)
			return
		}
		logger.LogSentMessage("Ready", msg.Content)

		// Step 2 - Create available team logo reactions in the Discord server
		for _, team := range league.Teams {
			logger.LogCreateReaction("Ready", "Reaction "+team.DisplayName+" creating...")
			emoji, err := utils.CreateGuildEmoji(session, globals.GuildId, team.Name, team.LogoBase64)
			if err != nil {
				log.Fatalf("An error occured while creating emojis: %v", err)
				return
			}
			reactionIdsByMessageId[msg.ID] = append(reactionIdsByMessageId[msg.ID], fmt.Sprintf("%s:%s", emoji.Name, emoji.ID))
			logger.LogCreateReaction("Ready", emoji.Name)
		}
	}

	// Step 3 - React with the team logos
	// For each message sent above
	fmt.Println("Reacting to league messages...")
	for k, v := range reactionIdsByMessageId {
		messageId := k
		reactionIds := v
		// Add each available reaction
		for _, emojiId := range reactionIds {
			err := session.MessageReactionAdd(channelId, messageId, emojiId)
			if err != nil {
				fmt.Printf("Error adding reaction with ID %s to message. Err = %s", emojiId, err)
			} else {
				logger.LogAddReaction("Ready", emojiId)
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
