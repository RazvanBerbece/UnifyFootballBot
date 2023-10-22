package teamassign

import (
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"

	"github.com/RazvanBerbece/UnifyFootballBot/internal/logger"

	"github.com/RazvanBerbece/UnifyFootballBot/internal/globals"

	favouriteTeamsRepository "github.com/RazvanBerbece/UnifyFootballBot/internal/data/favourite-teams"
	reactionAddHandlers "github.com/RazvanBerbece/UnifyFootballBot/internal/handlers/reactionAddEvent"
)

func MessageReactionRemoveTeamUnassign(s *discordgo.Session, event *discordgo.MessageReactionRemove) {

	if !globals.TeamUnassignHandlerEnabled {
		return
	}

	// Only use this handler function in the team assignment channel
	if event.ChannelID != globals.TeamAssignChannelId {
		return
	}

	logger.LogHandlerCall("MessageReactionRemoveTeamUnassign", "")

	userId := event.MessageReaction.UserID

	// Return immediately if user does not have a favourited team
	if !reactionAddHandlers.UserHasFavouritedTeam(userId) {
		return
	}

	// Fetch message history for the team-assign channel
	maxMsgLimit := 10
	messages, err := s.ChannelMessages(os.Getenv("TEAM_ASSIGN_CHANNEL_ID"), maxMsgLimit, "", "", "")
	if err != nil {
		fmt.Println("Error fetching messages. Err = ", err)
		return
	}

	// If the current reaction which was deleted is for a message in the team-assign channel
	for _, msg := range messages {
		if msg.ID == event.MessageReaction.MessageID {
			// Remove favourite team entry for given user from DB
			repo := favouriteTeamsRepository.NewFavouriteTeamsRepository()
			if err != nil {
				fmt.Errorf("Could not retrieve favourite team entry from DB : %v", err)
			}
			_, errDelete := repo.DeleteFavouriteTeam(userId, event.MessageReaction.Emoji.Name)
			if errDelete != nil {
				fmt.Errorf("Could not insert new favourite team entry into DB : %v", err)
			}
		}
	}

}
