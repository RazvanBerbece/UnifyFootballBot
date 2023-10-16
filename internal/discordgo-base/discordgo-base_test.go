package discordgobase

import (
	"log"
	"testing"

	"github.com/RazvanBerbece/UnifyFootballBot/internal/setup"
)

func dummyHandler() {
	log.Println("Acting as a handler method.")
}

func TestConfigureBase(t *testing.T) {

	// Arrange
	env := setup.Setup("../../.env")
	handlers := []interface{}{dummyHandler}
	botBase := DiscordGoBase{}

	// Act
	botBase.ConfigureBase(env.DiscordBotToken, handlers)

	// Assert
	if botBase.botSession == nil {
		t.Error("discordgo session should be initialised after calling ConfigureBase.")
	}
}
