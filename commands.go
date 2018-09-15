package main

import (
	"os"
	"os/exec"
	"strings"
	"fmt"
	"path/filepath"
	"github.com/bwmarrin/discordgo"
)

const outputPath = "image_script/images/output.png"

// Test assures the client that the bot is up and running
func Test(session *discordgo.Session, msg *discordgo.MessageCreate) {
	session.ChannelMessageSend(msg.ChannelID, "System is up")
}

// Fallback will be called if a messgae references the bot
// but does not match any of the existing commands
func Fallback(session *discordgo.Session, msg *discordgo.MessageCreate) {
	errMsg := "Command does not exist or was improperly used"
	session.ChannelMessageSend(msg.ChannelID, errMsg)
}

// Collage runs the process to create a collage and send it back to the user
func Collage(session *discordgo.Session, msg *discordgo.MessageCreate) {
	msgText := strings.ToLower(msg.Content)
	url := msg.Attachments[0].URL

	filename, err := DownloadImage(url)

	if err != nil {
		fmt.Println("Error: %v", err)
		return
	}

	runCollageScript(filename, getKeyword(msgText))

	sendMessageFile(session, msg.ChannelID, filename)
}

func runCollageScript(filename string, keyword string) {
	inputPath := filepath.Join("image_script/images", filename)

	cmd := exec.Command("python3", "image_script/collage.py", inputPath, keyword)

	cmd.Run()
}

func getKeyword(msgText string) string {
	words := strings.Split(msgText, " ")
	return words[2]
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