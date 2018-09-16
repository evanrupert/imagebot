package main

import (
  "fmt"
  "github.com/bwmarrin/discordgo"
  "strings"
)

const botID = "490372581833179166"
const testCmd = "test"
const collageCmd = "collage"
const minecraftCmd = "minecraft"
const weatherCmd = "weather"

// MessageIsTestRequest returns true if the message is a valid test command
func MessageIsTestRequest(msg *discordgo.MessageCreate) bool {
  msgText := strings.ToLower(msg.Content)
  return strings.HasSuffix(msgText, testCmd)
}

// MessageIsCollageRequest returns true if the message is a valid collage command
func MessageIsCollageRequest(msg *discordgo.MessageCreate) bool {
  msgText := strings.ToLower(msg.Content)
  words := strings.Fields(msgText)

  return len(words) == 3 &&
         hasCommand(msgText, collageCmd) &&
         len(msg.Attachments) == 1
}

// MessageIsHelpRequest returns true if the message is a valid help command
func MessageIsHelpRequest(msg *discordgo.MessageCreate) bool {
  return hasCommand(msg.Content, "help")
}

// MessageIsMinecraftRequest returns true if the message is a valid minecraft command
func MessageIsMinecraftRequest(msg *discordgo.MessageCreate) bool {
  return hasCommand(msg.Content, minecraftCmd) &&
         len(msg.Attachments) == 1
}

// MessageIsWeatherTodayRequest returns true if valid command for weather today 
func MessageIsWeatherTodayRequest(msg *discordgo.MessageCreate) bool {
  return hasWeatherCommand(msg.Content, "today")
}

// MessageIsWeatherTomorrowRequest returns true if valid command for weather tomorrow
func MessageIsWeatherTomorrowRequest(msg *discordgo.MessageCreate) bool {
  return hasWeatherCommand(msg.Content, "tomorrow")
}

// MessageIsWeatherWeekRequest returs true if command is a valid weather week command
func MessageIsWeatherWeekRequest(msg *discordgo.MessageCreate) bool {
  return hasWeatherCommand(msg.Content, "week")
}
// IsValidBotCommand returns true if the given msg is a valid imagebot command
func IsValidBotCommand(session *discordgo.Session, msg *discordgo.MessageCreate) bool {
  msgText := strings.ToLower(msg.Content)

  return msg.Author.ID != session.State.User.ID &&
    strings.HasPrefix(msgText, getBotTag())
}

func hasCommand(str string, cmd string) bool {
  lowerStr := strings.ToLower(str)
  words := strings.Fields(lowerStr)

  return words[0] == getBotTag() &&
         words[1] == cmd
}

func hasWeatherCommand(str string, cmd string) bool {
  lowerStr := strings.ToLower(str)
  words := strings.Fields(lowerStr)

  return hasCommand(str, "weather") && words[2] == cmd
}

func getBotTag() string {
  return fmt.Sprintf("<@%s>", botID)
}