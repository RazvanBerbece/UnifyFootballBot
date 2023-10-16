package handlers

import (
	messageEventHandlers "github.com/RazvanBerbece/UnifyFootballBot/internal/handlers/messageEvent"
)

func GetHandlersAsList() []interface{} {
	// Add new handler methods here
	return []interface{}{
		messageEventHandlers.Ping, messageEventHandlers.Pong,
	}
}
