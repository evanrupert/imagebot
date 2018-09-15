package main

import (
	"github.com/bwmarrin/discordgo"
	"strings"
)

const botName = "imagebot"

// MessageIsTestRequest returns true if the message is a valid test command
func MessageIsTestRequest(msg *discordgo.MessageCreate) bool {
	msgText := strings.ToLower(msg.Content)
	return strings.HasSuffix(msgText, "test")
}

// MessageIsCollageRequest returns true if the message is a valid collage command
func MessageIsCollageRequest(msg *discordgo.MessageCreate) bool {
	msgText := strings.ToLower(msg.Content)
	words := strings.Split(msgText, " ")

	return len(words) == 3 &&
				 words[0] == botName &&
				 words[1] == "collage" &&
				 len(msg.Attachments) == 1
}

// IsValidBotCommand returns true if the given msg is a valid imagebot command
func IsValidBotCommand(session *discordgo.Session, msg *discordgo.MessageCreate) bool {
	msgText := strings.ToLower(msg.Content)

	return msg.Author.ID != session.State.User.ID &&
		strings.HasPrefix(msgText, botName)
}
