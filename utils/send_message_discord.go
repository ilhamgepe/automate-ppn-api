package utils

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"os"
	"time"
)

func SendMessageToDiscord(message string) {
	fmt.Println("###  SEND MESSAGE TO DISCORD  ###")
	discord, err := discordgo.New(fmt.Sprintf("Bot %s", os.Getenv("DISCORD_TOKEN")))
	PanicIfError(err)
	err = discord.Open()
	PanicIfError(err)
	_, err = discord.ChannelMessageSend(os.Getenv("DISCORD_CHANNEL_ID"), message)
	PanicIfError(err)
	fmt.Printf("### FINISHED SEND MESSAGE TO DISCORD %v ###\n", fmt.Sprintf("%v-%v-%v-%v-%v", time.Now().Year(), time.Now().Month(), time.Now().Day(), time.Now().Hour(), time.Now().Minute()))
}
