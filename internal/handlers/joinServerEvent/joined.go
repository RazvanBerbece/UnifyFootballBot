package joinserverevent

import (
	"fmt"

	"github.com/RazvanBerbece/UnifyFootballBot/internal/logger"
	"github.com/bwmarrin/discordgo"
)

// Called once a new member joins the Discord server. Manages storing the user details to the DB
// and sending the welcome text. (WIP)
func NewMember(s *discordgo.Session, event *discordgo.GuildMemberAdd) {

	// event.GuildID contains the ID of the guild (server) the event occurred in.
	// event.User contains information about the new member who joined the server.
	logger.LogHandlerCall(fmt.Sprintf("NewMember (%s joined %s)", event.User.Username, event.GuildID), "")

	_, err := s.Guild(event.GuildID)
	if err != nil {
		fmt.Println("Error retrieving guild information: ", err)
		return
	}

	// Store new user
	// TODO

}
