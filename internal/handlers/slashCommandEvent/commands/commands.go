package commands

import (
	"github.com/bwmarrin/discordgo"
)

var SlashCommands = []*discordgo.ApplicationCommand{
	{
		Name:        "ping_unify",
		Description: "Basic ping slash interaction",
	},
}

var SlashCommandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
	"ping_unify": HandleSlashPingUnify,
	// Add handlers for other slash commands
}
