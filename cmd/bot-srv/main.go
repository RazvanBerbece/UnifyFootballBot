package main

import (
	discordgobase "github.com/RazvanBerbece/UnifyFootballBot/internal/discordgo-base"
	"github.com/RazvanBerbece/UnifyFootballBot/internal/setup"

	"github.com/RazvanBerbece/UnifyFootballBot/internal/handlers"
)

func main() {

	// Retrieve the configured runtime environment
	env := setup.Setup(".env")

	// Configure handler functions
	handlers := handlers.GetHandlersAsList()

	// Configure DiscordGo base
	botBase := discordgobase.DiscordGoBase{}
	botBase.ConfigureBase(env.DiscordBotToken, handlers)

	// Connect to Discord servers
	botBase.Connect()

	// Close connection
	botBase.Close()

}
