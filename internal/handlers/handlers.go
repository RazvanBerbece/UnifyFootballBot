package handlers

import (
	messageEventHandlers "github.com/RazvanBerbece/UnifyFootballBot/internal/handlers/messageEvent"
	reactionAddHandlers "github.com/RazvanBerbece/UnifyFootballBot/internal/handlers/reactionAddEvent"
	reactionRemoveHandlers "github.com/RazvanBerbece/UnifyFootballBot/internal/handlers/reactionRemoveEvent"
	readyEventHandlers "github.com/RazvanBerbece/UnifyFootballBot/internal/handlers/readyEvent"
)

func GetHandlersAsList() []interface{} {
	// Add new handler methods here
	return []interface{}{
		// <---- On Ready ---->
		readyEventHandlers.Ready,
		// <---- On Message Created ---->
		messageEventHandlers.Ping, messageEventHandlers.Pong,
		// <---- On Reaction Added ---->
		reactionAddHandlers.MessageReactionAddTeamAssign,
		// <---- On Reaction Removed ---->
		reactionRemoveHandlers.MessageReactionRemoveTeamUnassign,
		// <---- On New Join ---->
		// <---- On Slash Command ---->
	}
}
