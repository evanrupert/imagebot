package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strings"
)

const BotName = "imagebot"

func MessageIsTestRequest(msg *discordgo.MessageCreate) bool {
	msgText := strings.ToLower(msg.Content)
	return strings.HasSuffix(msgText, "test")
}

func MessageIsCollageRequest(msg *discordgo.MessageCreate) bool {
	return strings.HasPrefix(msg.Content, fmt.Sprintf("%s %s", BotName, "collage")) &&
		len(msg.Attachments) == 1
}

// IsValidBotCommand returns true if the given msg is a valid imagebot command 
func IsValidBotCommand(session *discordgo.Session, msg *discordgo.MessageCreate) bool {
	msgText := strings.ToLower(msg.Content)

	return msg.Author.ID != session.State.User.ID &&
		strings.HasPrefix(msgText, BotName)
}
