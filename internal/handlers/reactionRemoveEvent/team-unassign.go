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

	// Only use this handler function in the team assignment channel
	if event.ChannelID != globals.TeamAssignChannelId {
		return
	}

	logger.LogHandlerCall("MessageReactionRemoveTeamUnassign", "")

	userId := event.MessageReaction.UserID

	// Return immediately if user does not have a favourited team
	// or if the bot removed the reaction
	if !reactionAddHandlers.UserHasFavouritedTeam(userId) || (event.UserID == os.Getenv("BOT_USER_ID")) {
		return
	}

	// Fetch message history for the team-assign channel
	maxMsgLimit := 100
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
			// Oct 2022: It seems that calling MessageReactionRemove below triggers the associated handler for happy path team unassignment.
			// That is undersirable because it will then delete the favourite team from the DB for the given user when it shouldn't.
			// So I conditioned the deletion on the current team name that is in the DB.
			currentFavTeamName, err := repo.GetFavouriteTeam(userId)
			logger.LogHandlerCall(fmt.Sprintf("MessageReactionRemoveTeamUnassign - %s", currentFavTeamName), "")
			if err != nil {
				fmt.Errorf("Could not retrieve favourite team entry from DB : %v", err)
			}
			_, errDelete := repo.DeleteFavouriteTeam(userId, currentFavTeamName)
			if errDelete != nil {
				fmt.Errorf("Could not insert new favourite team entry into DB : %v", err)
			}
		}
	}

}
