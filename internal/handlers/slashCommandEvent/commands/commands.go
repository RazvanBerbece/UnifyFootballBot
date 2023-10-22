package commands

import (
	commandHandlers "github.com/RazvanBerbece/UnifyFootballBot/internal/handlers/slashCommandEvent/commands/slash-handlers"
	"github.com/bwmarrin/discordgo"
)

var SlashCommands = []*discordgo.ApplicationCommand{
	{
		Name:        "ping_unify",
		Description: "Basic ping slash interaction for the Unify Bot",
	},
	{
		Name:        "my_teams",
		Description: "Replies with a list of the user's favourited teams from the `team-assignment` channel",
	},
}

var SlashCommandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
	"ping_unify": commandHandlers.HandleSlashPingUnify,
	"my_teams":   commandHandlers.HandleSlashMyTeams,
}
