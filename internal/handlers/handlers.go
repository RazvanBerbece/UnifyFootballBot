package handlers

import (
	messageEventHandlers "github.com/RazvanBerbece/UnifyFootballBot/internal/handlers/messageEvent"
	reactionAddHandlers "github.com/RazvanBerbece/UnifyFootballBot/internal/handlers/reactionAddEvent"
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
		// <---- On New Join ---->
		// <---- On Slash Command ---->
	}
}
