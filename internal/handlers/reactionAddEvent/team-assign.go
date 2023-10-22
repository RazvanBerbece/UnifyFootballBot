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

	// Only execute if user hasn't been assigned a favourite team
	if UserHasFavouritedTeam(userId) {
		// Revert current reaction to enforce 1 favourite team per user
		revertFavouriteTeamAssignment(s, event.MessageReaction.MessageID, event.MessageReaction.Emoji.Name, userId)
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
	team, err := repo.GetFavouriteTeam(userId)
	if err != nil {
		fmt.Errorf("Could not retrieve favourite team from DB for user with id %s : %v", userId, err)
	}
	if team != "" {
		return true
	}
	return false
}

func revertFavouriteTeamAssignment(session *discordgo.Session, reactionMessageId string, teamName string, userId string) {

	// Get the message object from the channel
	msg, err := session.ChannelMessage(os.Getenv("TEAM_ASSIGN_CHANNEL_ID"), reactionMessageId)
	if err != nil {
		fmt.Println("Error retrieving message: ", err)
		return
	}
	// Loop through the reactions on the message to find the specific user reaction to remove
	for _, reaction := range msg.Reactions {
		if reaction.Emoji.Name == teamName {
			// Remove the reaction
			emojiId := fmt.Sprintf("%s:%s", reaction.Emoji.Name, reaction.Emoji.ID)
			err := session.MessageReactionRemove(
				os.Getenv("TEAM_ASSIGN_CHANNEL_ID"),
				reactionMessageId,
				emojiId,
				userId,
			)
			if err != nil {
				fmt.Println("Error removing reaction: ", err)
			}
		}
	}

	// + message to user about conditions to assign teams ?
}
