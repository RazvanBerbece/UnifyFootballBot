package main

import (
	discordgobase "github.com/RazvanBerbece/UnifyFootballBot/internal/discordgo-base"

	"github.com/RazvanBerbece/UnifyFootballBot/internal/globals"
	"github.com/RazvanBerbece/UnifyFootballBot/internal/handlers"
)

func main() {

	// Configure handler functions
	handlers := handlers.GetHandlersAsList()

	// Configure DiscordGo base
	botBase := discordgobase.DiscordGoBase{}
	botBase.ConfigureBase(globals.DiscordBotToken, handlers)

	// Connect to Discord servers
	botBase.Connect()

	// Close connection
	botBase.Close()

}
