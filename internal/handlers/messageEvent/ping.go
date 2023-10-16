package messageevent

import (
	"github.com/RazvanBerbece/UnifyFootballBot/internal/logger"
	"github.com/bwmarrin/discordgo"
)

func Ping(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	// If the message is "ping" reply with "Pong!"
	if m.Content == "ping" {
		logger.LogHandlerCall("Ping", "../../logs", "handler_calls.log")
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}

}
