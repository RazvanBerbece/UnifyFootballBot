package teamassign

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"

	"github.com/RazvanBerbece/UnifyFootballBot/internal/globals"
	"github.com/RazvanBerbece/UnifyFootballBot/internal/logger"
	"github.com/RazvanBerbece/UnifyFootballBot/internal/utils"

	favouriteTeamsRepository "github.com/RazvanBerbece/UnifyFootballBot/internal/data/favourite-teams"
)

func MessageReactionAddTeamAssign(s *discordgo.Session, event *discordgo.MessageReactionAdd) {

	// Only use this handler function in the team assignment channel
	if event.ChannelID != globals.TeamAssignChannelId {
		return
	}

	userId := event.MessageReaction.UserID

	// If bot added reaction, simply return
	if userId == globals.BotUserId {
		return
	}

	logger.LogHandlerCall("MessageReactionAddTeamAssign", "")

	// Fetch message history for the team-assign channel
	maxMsgLimit := 5
	messages, err := s.ChannelMessages(globals.TeamAssignChannelId, maxMsgLimit, "", "", "")
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
				fmt.Printf("Could not insert new favourite team entry into DB : %v", err)
			}
			// Send DM to user to confirm transaction
			displayTeamName := strings.Replace(teamName, "_", " ", -1)
			transactionMsg := fmt.Sprintf("Confirmation of selecting %s as one of your favourite teams.", displayTeamName)
			errDm := utils.SendDmToUserFromMsgReactionAdd(s, event, userId, transactionMsg)
			if errDm != nil {
				fmt.Println("Error sending DM: ", errDm)
			}
		}
	}

}

func UserHasFavouritedTeam(userId string) bool {
	repo := favouriteTeamsRepository.NewFavouriteTeamsRepository()
	teams, err := repo.GetFavouriteTeams(userId)
	if err != nil {
		fmt.Printf("Could not retrieve favourite team from DB for user with id %s : %v", userId, err)
	}
	if len(teams) > 0 {
		return true
	}
	return false
}
