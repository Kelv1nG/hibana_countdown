package bot

import (
	"log"

	"github.com/Kelv1nG/hibana_countdown/config"
	"github.com/bwmarrin/discordgo"
)

func CreateSession() *discordgo.Session {
	if config.AppConfig.Token == "" {
		log.Fatal("Bot token is required. Use -token=<YOUR_TOKEN>")
	}

	sess, err := discordgo.New("Bot " + config.AppConfig.Token)
	if err != nil {
		log.Fatalf("Failed to create Discord session: %v", err)
	}
	AddBotCommands(sess)
	return sess
}

func AddBotCommands(sess *discordgo.Session) {
	sess.AddHandler(SendTimeRemain)
	sess.AddHandler(RandomHibanaSong)
	sess.AddHandler(RegisterChannelForCountdown)
	sess.AddHandler(UnregisterChannelForCountdown)

	_, err := sess.ApplicationCommandCreate(config.AppConfig.AppID, "", &discordgo.ApplicationCommand{
		Name:        HibanaCountdown,
		Description: "Calculates remaining time before Ados Hibana concert",
	})
	if err != nil {
		log.Fatalf("Cannot create slash command: %v", err)
	}

	_, err = sess.ApplicationCommandCreate(config.AppConfig.AppID, "", &discordgo.ApplicationCommand{
		Name:        HibanaSong,
		Description: "Get a random song from the Hibana playlist",
	})
	if err != nil {
		log.Fatalf("Cannot create slash command: %v", err)
	}

	_, err = sess.ApplicationCommandCreate(config.AppConfig.AppID, "", &discordgo.ApplicationCommand{
		Name:        HibanaRegisterChannel,
		Description: "Register a channel to broadcast countdown for hibana",
	})
	if err != nil {
		log.Fatalf("Cannot create slash command: %v", err)
	}

	_, err = sess.ApplicationCommandCreate(config.AppConfig.AppID, "", &discordgo.ApplicationCommand{
		Name:        HibanaUnregisterChannel,
		Description: "Remove a channel to broadcast countdown for hibana",
	})
	if err != nil {
		log.Fatalf("Cannot create slash command: %v", err)
	}
}
