package setup

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type RuntimeEnvironment struct {
	DiscordBotToken string
	AppId           string
	Environment     string // Staging / production flag
}

// Returns a struct containing the necessary environment variables used to build and configure the DiscordGo base.
func Setup(filepathToEnv string) RuntimeEnvironment {

	err := godotenv.Load(filepathToEnv)
	if err != nil {
		log.Fatal("Error occured while loading .env file. Exiting.")
	}

	env := RuntimeEnvironment{
		DiscordBotToken: os.Getenv("DISCORD_BOT_TOKEN"),
		AppId:           os.Getenv("APP_ID"),
		Environment:     os.Getenv("ENVIRONMENT"),
	}

	return env

}
