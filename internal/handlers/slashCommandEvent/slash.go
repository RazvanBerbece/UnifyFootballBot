package slashCommandEvent

import (
	"github.com/bwmarrin/discordgo"

	slashCommands "github.com/RazvanBerbece/UnifyFootballBot/internal/handlers/slashCommandEvent/commands"
)

func HandleSlashCommands(s *discordgo.Session, event *discordgo.InteractionCreate) {
	// Check if the interaction type is a command
	if event.Type == discordgo.InteractionApplicationCommand {
		command := event.ApplicationCommandData()

		// Handle different slash commands
		switch command.Name {
		case "ping":
			// Handle /ping command
			s.InteractionRespond(event.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: slashCommands.HandlePingSlashCommand(),
				},
			})
			// TODO more slash commands
		}
	}
}
