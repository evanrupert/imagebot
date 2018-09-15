package main

import (
	"os"
	"fmt"
	"github.com/bwmarrin/discordgo"
)

func Test(session *discordgo.Session, msg *discordgo.MessageCreate) {
	session.ChannelMessageSend(msg.ChannelID, "System is up")
}

func Collage(session *discordgo.Session, msg *discordgo.MessageCreate) {
	session.ChannelMessageSend(msg.ChannelID, "Sent proper image request")
}

func Fallback(session *discordgo.Session, msg *discordgo.MessageCreate) {
	errMsg := "Command does not exist or was improperly used"
	session.ChannelMessage(msg.ChannelID, errMsg)
}

func sendMessageFile(session *discordgo.Session, channelID, filename string) {
  file, err := os.Open(filename)

  if err != nil {
    session.ChannelMessageSend(channelID, fmt.Sprintf("Failed to open file, Error: %v", err))
    return
  }

  defer file.Close()

  session.ChannelFileSend(channelID, filename, file)
}