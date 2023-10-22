package slashHandlers

import (
	"fmt"

	"github.com/RazvanBerbece/UnifyFootballBot/internal/globals"
	"github.com/RazvanBerbece/UnifyFootballBot/internal/logger"
	"github.com/bwmarrin/discordgo"

	favouriteTeamsRepository "github.com/RazvanBerbece/UnifyFootballBot/internal/data/favourite-teams"
)

func HandleSlashMyTeams(s *discordgo.Session, i *discordgo.InteractionCreate) {

	userId := i.Interaction.Member.User.ID
	userName := i.Interaction.Member.User.Username
	response := ""

	logger.LogSlashCommand("my_teams", userName)

	// Get user's favourite teams
	repo := favouriteTeamsRepository.NewFavouriteTeamsRepository()
	teams, err := repo.GetFavouriteTeams(userId)
	if err != nil {
		fmt.Errorf("Could not retrieve favourite teams from DB for user with id %s : %v", userId, err)
	}

	if teams != nil {
		if len(teams) == 0 {
			// User has *no* favourited teams
			response =
				"**You don't have any favourited teams !**\n\n" +
					fmt.Sprintf("To choose a favourite team please go to the <#%s> channel and pick your teams by reacting to the messages there.", globals.TeamAssignChannelId)
		} else {
			// User has favourited teams
			response = "Your favourited teams are: "
			for index, team := range teams {
				if index == len(teams)-1 {
					// Useful for pretty formatting of the list.
					// this accounts for the last entry in the favourite list, which doesn't need a comma after
					response = response + fmt.Sprintf("%s", team.TeamName)
					break
				}
				response = response + fmt.Sprintf("%s, ", team.TeamName)
			}
		}
	} else {
		// In case GetFavouriteTeams fails
		response = "An error occured while retrieving your favourite teams from the database."
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: response,
		},
	})
}
