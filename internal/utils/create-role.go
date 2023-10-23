package utils

import (
	"github.com/bwmarrin/discordgo"
)

type RoleConfig struct {
	Params *discordgo.RoleParams
}

// func CreateRole(session *discordgo.Session, guildId string, role RoleConfig) error {

// 	// Create the role in the guild
// 	role, err := session.GuildRoleCreate(guildId, roleParams)
// 	if err != nil {
// 		return err
// 	}

// 	// Role successfully created
// 	logger.LogRoleCreation(role)
// 	return nil

// }
