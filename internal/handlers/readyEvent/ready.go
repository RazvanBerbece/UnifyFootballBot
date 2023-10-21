package readyevent

import (
	"os"

	leagues "github.com/RazvanBerbece/UnifyFootballBot/internal/handlers/readyEvent/resources"
	"github.com/RazvanBerbece/UnifyFootballBot/internal/logger"
	"github.com/bwmarrin/discordgo"
)

// Called once the Discord servers confirm a succesful connection.
func Ready(s *discordgo.Session, event *discordgo.Ready) {

	logger.LogHandlerCall("Ready", "")

	// Send static messages to required channels
	// team-assign, etc.
	sendTeamAssignMessages(s, os.Getenv("TEAM_ASSIGN_CHANNEL_ID"))

}

// Sends the team assign messages and the reaction choices to the designed channel.
func sendTeamAssignMessages(session *discordgo.Session, channelId string) {

	// Get all available leagues and reactions to post in the channel
	leagueMessages := leagues.GetLeaguesAsList()

	// Step 1 - Send the messages
	for _, message := range leagueMessages {
		messageToSend := message.LeagueName + "\nReact to this message to get your roles!"
		session.ChannelMessageSend(channelId, messageToSend)
		logger.LogSentMessage("Ready", messageToSend)
	}

	// Step 2 - React with the team logos
	// TODO

}
