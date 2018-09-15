package main

import (
	"os"
	"os/signal"
	"syscall"
	"fmt"
	"github.com/bwmarrin/discordgo"
)

func main() {
	token := os.Getenv("DISCORD_TOKEN")	

	discord, err := discordgo.New("Bot " + token)

	if err != nil {
		fmt.Println("Error creating discord session")
		return
	}

	discord.AddHandler(messageCreate)

	err = discord.Open()
	if err != nil {
		fmt.Println("Error opening connection,", err)
		return
	}

	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)

	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<- sc

	discord.Close()
}

func messageCreate(session *discordgo.Session, msg *discordgo.MessageCreate) {
	if msg.Author.ID == session.State.User.ID {
		return
	}

	if msg.Content == "ping" {
		session.ChannelMessageSend(msg.ChannelID, "Pong!")
	}

	if msg.Content == "pong" {
		session.ChannelMessageSend(msg.ChannelID, "Ping!")
	}
}