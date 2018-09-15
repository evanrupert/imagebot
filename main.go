package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	token := os.Getenv("DISCORD_TOKEN")

	discord, err := discordgo.New("Bot " + token)

	if err != nil {
		fmt.Println("Error creating discord session")
		return
	}

	discord.AddHandler(messageHandler)

	err = discord.Open()
	if err != nil {
		fmt.Println("Error opening connection,", err)
		return
	}

	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)

	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	discord.Close()
}

func messageHandler(session *discordgo.Session, msg *discordgo.MessageCreate) {
	if !IsValidBotCommand(session, msg) {
		return
	} else if MessageIsTestRequest(msg) {
		go Test(session, msg)
	} else if MessageIsCollageRequest(msg) {
		go Collage(session, msg)
	} else {
		go Fallback(session, msg)
	}
}
