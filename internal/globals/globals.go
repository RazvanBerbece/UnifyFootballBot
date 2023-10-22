package globals

import "os"

// =============== ENVIRONMENT VARIABLES ===============

var Environment = os.Getenv("ENVIRONMENT") // staging / production

var DiscordBotToken = os.Getenv("DISCORD_BOT_TOKEN")
var BotUserId = os.Getenv("BOT_USER_ID")

var MySqlRootPassword = os.Getenv("MYSQL_ROOT_PASSWORD")
var MySqlDatabaseName = os.Getenv("MYSQL_DATABASE")
var MySqlUserName = os.Getenv("MYSQL_USER")
var MySqlPassword = os.Getenv("MYSQL_PASSWORD")
var MySqlConnectionString = os.Getenv("UNIFYFOOTBALL_DB_CONNSTRING") // in format `root_username:root_password@tcp(host:port)/db_name`

var TeamAssignChannelId = os.Getenv("TEAM_ASSIGN_CHANNEL_ID")
