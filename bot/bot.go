package bot

import (
	"fmt"
	"log"

	"github.com/Kelv1nG/hibana_countdown/hibana"
	"github.com/Kelv1nG/hibana_countdown/spotify"
	"github.com/bwmarrin/discordgo"
)

const (
	HibanaCountdown         = "hibana-countdown"
	HibanaSong              = "hibana-song"
	HibanaRegisterChannel   = "hibana-register-channel"
	HibanaUnregisterChannel = "hibana-unregister-channel"
)

func SendTimeRemain(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.ApplicationCommandData().Name == HibanaCountdown {

		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: hibana.CalculateTimeRemain(),
			},
		})
		if err != nil {
			log.Fatalf("Failed to respond to interaction: %v", err)
		}
	}
}

func RandomHibanaSong(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.ApplicationCommandData().Name == HibanaSong {
		url, err := spotify.RandomHibanaSong()
		if err != nil {
			log.Printf("Failed to respond to interaction: %v", err)
			return
		}

		err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Get a random hibana song...",
			},
		})
		if err != nil {
			log.Printf("Failed to respond to interaction: %v", err)
			return
		}

		_, err = s.ChannelMessageSend(i.ChannelID, url)
		if err != nil {
			log.Printf("Failed to send play command: %v", err)
		}
	}
}

func RegisterChannelForCountdown(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.ApplicationCommandData().Name == HibanaRegisterChannel {
		hibana.RegisterChannel(i.ChannelID)

		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "registering channel to be notified in hibana countdown",
			},
		})
		if err != nil {
			log.Printf("Failed to respond to interaction: %v", err)
			return
		}

		message := fmt.Sprintf("Registered channel for countdown %s", i.ChannelID)
		_, err = s.ChannelMessageSend(i.ChannelID, message)
		if err != nil {
			log.Printf("Failed %v", err)
		}
	}
}

func UnregisterChannelForCountdown(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.ApplicationCommandData().Name == HibanaUnregisterChannel {
		hibana.UnregisterChannel(i.ChannelID)

		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "removing channel to be notified in hibana countdown",
			},
		})
		if err != nil {
			log.Printf("Failed to respond to interaction: %v", err)
			return
		}
		message := fmt.Sprintf("removed channel for countdown %s", i.ChannelID)
		_, err = s.ChannelMessageSend(i.ChannelID, message)
		if err != nil {
			log.Printf("Failed %v", err)
		}
	}
}
