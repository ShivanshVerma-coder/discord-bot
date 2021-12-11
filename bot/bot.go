package bot

import (
	"log"

	"github.com/ShivanshVerma-coder/discord-bot/config"
	"github.com/bwmarrin/discordgo"
)

var BotID string
var goBot *discordgo.Session
var err error

func Start() error {
	//create discord bot session
	goBot, err = discordgo.New("Bot " + config.Token) //creates new bot session

	if err != nil {
		log.Fatal(err)
		return err
	}

	user, err := goBot.User("@me")

	if err != nil {
		log.Fatal(err)
		return err
	}
	BotID = user.ID

	goBot.AddHandler(messageHandler)

	err = goBot.Open() //Starts the session

	if err != nil {
		log.Fatal(err)
		return err
	}

	log.Println("Bot is now running")

	// defer goBot.Close() //closes session
	return nil
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	//check if message is from bot
	if m.Author.ID == BotID {
		return
	}

	if m.Content != "" {
		FuncMap(s, m)
	}
}
