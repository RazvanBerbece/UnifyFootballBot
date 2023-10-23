package utils

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func CreateGuildEmoji(s *discordgo.Session, guildID string, emojiName string, base64Image string) (*discordgo.Emoji, error) {

	base64Image = createDataURI(base64Image)

	emoji, err := s.GuildEmojiCreate(guildID, &discordgo.EmojiParams{
		Name:  emojiName,
		Image: base64Image,
	})
	if err != nil {
		return nil, err
	}

	return emoji, nil

}

func createDataURI(base64Image string) string {
	mimeType := "image/png"
	dataURI := fmt.Sprintf("data:%s;base64,%s", mimeType, base64Image)
	return dataURI
}
