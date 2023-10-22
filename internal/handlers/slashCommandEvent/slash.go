package slashCommandEvent

import (
	"log"

	"github.com/RazvanBerbece/UnifyFootballBot/internal/globals"
	commands "github.com/RazvanBerbece/UnifyFootballBot/internal/handlers/slashCommandEvent/commands"
	"github.com/bwmarrin/discordgo"
)

func RegisterSlashCommands(s *discordgo.Session) error {

	globals.RegisteredCommands = make([]*discordgo.ApplicationCommand, len(commands.SlashCommands))
	for index, cmd := range commands.SlashCommands {
		_, err := s.ApplicationCommandCreate(globals.AppId, "", cmd)
		if err != nil {
			return err
		}
		globals.RegisteredCommands[index] = cmd
	}

	// Add slash command handlers
	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if handlerFunc, ok := commands.SlashCommandHandlers[i.ApplicationCommandData().Name]; ok {
			handlerFunc(s, i)
		}
	})

	return nil
}

func CleanupSlashCommands(s *discordgo.Session) {
	log.Println("Removing commands...")
	for _, cmd := range globals.RegisteredCommands {
		err := s.ApplicationCommandDelete(globals.AppId, "", cmd.ID)
		if err != nil {
			log.Panicf("Cannot delete '%v' command: %v", cmd.Name, err)
		}
	}
}
