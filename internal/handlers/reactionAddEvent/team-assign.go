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
			// TODO STORE ETC.
			repo := favouriteTeamsRepository.NewFavouriteTeamsRepository()
			_, err := repo.InsertFavouriteTeam(userId, teamName)
			if err != nil {
				fmt.Errorf("Could not insert new favourite team entry into DB : %v", err)
			}
		}
	}

}

func userHasFavouritedTeam(userId string) bool {
	return false
}
