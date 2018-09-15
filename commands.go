package main

import (
	"os"
	"os/exec"
	"strings"
	"fmt"
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

	filename, err := downloadMessageAttachment(msg)

	if err != nil {
		fmt.Println(err)
		return
	}

	runCollageScript(filename, getKeyword(msgText))

	sendMessageFile(session, msg.ChannelID, outputPath)
}

// Minecraft runes the collage process but with only minecraft blocks
func Minecraft(session *discordgo.Session, msg *discordgo.MessageCreate) {
	filename, err := downloadMessageAttachment(msg)

	if err != nil {
		fmt.Println(err)
		return
	}

	runMinecraftScript(filename)

	sendMessageFile(session, msg.ChannelID, outputPath)
}

func downloadMessageAttachment(msg *discordgo.MessageCreate) (string, error) {
	url := msg.Attachments[0].URL

	return DownloadImage(url)
}

func runCollageScript(filename string, keyword string) {
	cmd := exec.Command("python3", "image_script/collage.py", filename, keyword)

	err := cmd.Run()

	if err != nil {
		fmt.Println(err)
	}
}

func runMinecraftScript(filename string) {
	cmd := exec.Command("python3", "image_script/minecraft.py", filename)

	err := cmd.Run()

	if err != nil {
		fmt.Println(err)
	}
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