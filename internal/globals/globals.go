package globals

import (
	"os"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"

	footballApiModels "github.com/RazvanBerbece/UnifyFootballBot/internal/apis/api-football/models"
)

// =============== ENVIRONMENT VARIABLES ===============
var err = godotenv.Load(".env") // Load

var Environment = os.Getenv("ENVIRONMENT") // staging / production

var DiscordBotToken = os.Getenv("DISCORD_BOT_TOKEN")
var BotUserId = os.Getenv("BOT_USER_ID")
var AppId = os.Getenv("APP_ID")

var MySqlRootPassword = os.Getenv("MYSQL_ROOT_PASSWORD")
var MySqlDatabaseName = os.Getenv("MYSQL_DATABASE")
var MySqlUserName = os.Getenv("MYSQL_USER")
var MySqlPassword = os.Getenv("MYSQL_PASSWORD")
var MySqlConnectionString = os.Getenv("UNIFYFOOTBALL_DB_CONNSTRING") // in format `root_username:root_password@tcp(host:port)/db_name`

var GuildId = os.Getenv("GUILD_ID") // the id of the server the bot is in
var TeamAssignChannelId = os.Getenv("TEAM_ASSIGN_CHANNEL_ID")

var RapidApiFootballKey = os.Getenv("RAPIDAPI_FOOTBALL_KEY")
var RapidApiHost = os.Getenv("RAPIDAPI_HOST")

// These environment variables are used to dictate the data required and used by the app

// A list of country names. This dictates what leagues and teams are watched by the server,
// and what roles and reactions are created as part of the team assignment feature.
var LeagueCountriesString = os.Getenv("LEAGUE_COUNTRIES")
var LeagueCountries = strings.Split(LeagueCountriesString, ",") // as a list
var LeaguesPerCountryString = os.Getenv("LEAGUES_PER_COUNTRY")
var LeaguesPerCountryNumber, _ = strconv.Atoi(LeaguesPerCountryString)

// =============== RUNTIME VARS ===============
var RegisteredCommands []*discordgo.ApplicationCommand
var AvailableLeagues []footballApiModels.League
