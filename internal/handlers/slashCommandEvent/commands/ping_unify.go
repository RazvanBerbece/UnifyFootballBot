package commands

import (
	"github.com/RazvanBerbece/UnifyFootballBot/internal/logger"
	"github.com/bwmarrin/discordgo"
)

func HandleSlashPingUnify(s *discordgo.Session, i *discordgo.InteractionCreate) {
	logger.LogSlashCommand("ping_unify", i.Interaction.Member.User.Username)
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Slash Pong!",
		},
	})
}
