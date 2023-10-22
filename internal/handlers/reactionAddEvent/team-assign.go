package teamassign

import (
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"

	favouriteTeamsRepository "github.com/RazvanBerbece/UnifyFootballBot/internal/data/favourite-teams"
)

func MessageReactionAddTeamAssign(s *discordgo.Session, event *discordgo.MessageReactionAdd) {

	userId := event.MessageReaction.UserID

	// Only execute if user hasn't been assigned a favourite team
	if userHasFavouritedTeam(userId) {
		// TODO: And revert reaction + message to user about conditions to assign teams ?
		return
	}

	// Fetch message history for the team-assign channel
	maxMsgLimit := 5
	messages, err := s.ChannelMessages(os.Getenv("TEAM_ASSIGN_CHANNEL_ID"), maxMsgLimit, "", "", "")
	if err != nil {
		fmt.Println("Error fetching messages. Err = ", err)
		return
	}

	// If the current reaction is for a message in the team-assign channel
	for _, msg := range messages {
		if msg.ID == event.MessageReaction.MessageID {
			// Then gather the user and team data
			reaction := event.MessageReaction.Emoji
			teamName := reaction.Name
			fmt.Printf("User with ID %s reacted to message with emoji %s\n", userId, teamName)
			// Store favourite team name for given user to DB
			repo := favouriteTeamsRepository.NewFavouriteTeamsRepository()
			_, err := repo.InsertFavouriteTeam(userId, teamName)
			if err != nil {
				fmt.Errorf("Could not insert new favourite team entry into DB : %v", err)
			}
		}
	}

}

func userHasFavouritedTeam(userId string) bool {
	repo := favouriteTeamsRepository.NewFavouriteTeamsRepository()
	team, err := repo.GetFavouriteTeam(userId)
	if err != nil {
		fmt.Errorf("Could not retrieve favourite team from DB for user with id %s : %v", userId, err)
	}
	if team != "" {
		return true
	}
	return false
}
