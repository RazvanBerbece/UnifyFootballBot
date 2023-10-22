package discordgobase

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"

	commands "github.com/RazvanBerbece/UnifyFootballBot/internal/handlers/slashCommandEvent"
)

type DiscordGoBase struct {
	botSession  *discordgo.Session
	isConnected bool
}

type DiscordClient interface {
	ConfigureBase(withToken string, handlers []interface{})
	Connect()
	Close()
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
	intents := discordgo.IntentsGuilds |
		discordgo.IntentsGuildMessages |
		discordgo.IntentsGuildMessageReactions |
		discordgo.PermissionManageMessages |
		discordgo.PermissionManageServer
	session.Identify.Intents = intents

	// Register slash commands
	err = commands.RegisterSlashCommands(session)
	if err != nil {
		log.Fatal("Error registering slash commands: ", err)
	}

	b.botSession = session

}

// Opens a persistent websocket connection to the Discord servers. Note that this method waits
// until CTRL-C or anther term signal is received.
func (b *DiscordGoBase) Connect() {

	err := b.botSession.Open()
	if err != nil {
		log.Fatal("Could not connect the bot to the Discord servers. Err =", err)
	}

	// wait here until CTRL-C or anther term signal is received
	fmt.Println("Discord bot session is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

}

// Closes the existing persistent websocket connection to the Discord servers.
func (b *DiscordGoBase) Close() {
	// Cleanup
	commands.CleanupSlashCommands(b.botSession)
	// Connection closing
	b.botSession.Close()
	b.isConnected = false
}
