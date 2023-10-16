package handlers

import (
	"github.com/RazvanBerbece/UnifyFootballBot/internal/logger"
	"github.com/bwmarrin/discordgo"
)

func pong(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	// If the message is "pong" reply with "Ping!"
	if m.Content == "pong" {
		logger.LogHandlerCall("pong")
		s.ChannelMessageSend(m.ChannelID, "Ping!")
	}
}
