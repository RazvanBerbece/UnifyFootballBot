package discordgobase

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

type DiscordGoBase struct {
	botSession *discordgo.Session
}

// Initiates the instance's botSession with a fully configured discordgo session (auth, handlers, intents).
func (b *DiscordGoBase) ConfigureBase(withToken string, handlers []interface{}) {

	// Create session
	session, err := discordgo.New("Bot " + withToken)
	if err != nil {
		log.Fatal("Could not create a Discord Bot session. Err =", err)
	}

	// Register custom handlers as callbacks for various events
	for _, handler := range handlers {
		session.AddHandler(handler)
	}

	// Register intents
	session.Identify.Intents = discordgo.IntentsGuildMessages

	b.botSession = session
}
