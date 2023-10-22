package teamassign

import (
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"

	"github.com/RazvanBerbece/UnifyFootballBot/internal/globals"
	"github.com/RazvanBerbece/UnifyFootballBot/internal/logger"

	favouriteTeamsRepository "github.com/RazvanBerbece/UnifyFootballBot/internal/data/favourite-teams"
)

func MessageReactionAddTeamAssign(s *discordgo.Session, event *discordgo.MessageReactionAdd) {

	// Only use this handler function in the team assignment channel
	if event.ChannelID != globals.TeamAssignChannelId {
		return
	}

	userId := event.MessageReaction.UserID

	// If bot added reaction, simply return
	if userId == os.Getenv("BOT_USER_ID") {
		return
	}

	logger.LogHandlerCall("MessageReactionAddTeamAssign", "")

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
			// Store favourite team name for given user to DB
			repo := favouriteTeamsRepository.NewFavouriteTeamsRepository()
			_, err := repo.InsertFavouriteTeam(userId, teamName)
			if err != nil {
				fmt.Errorf("Could not insert new favourite team entry into DB : %v", err)
			}
		}
	}

}

func UserHasFavouritedTeam(userId string) bool {
	repo := favouriteTeamsRepository.NewFavouriteTeamsRepository()
	teams, err := repo.GetFavouriteTeams(userId)
	if err != nil {
		fmt.Errorf("Could not retrieve favourite team from DB for user with id %s : %v", userId, err)
	}
	if len(teams) > 0 {
		return true
	}
	return false
}
