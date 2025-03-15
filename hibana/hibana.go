package hibana

import (
	"fmt"
	"time"
)

var hibanaConcert = time.Date(2025, time.May, 8, 0, 0, 0, 0, time.UTC)

// discord channels to be notified
var RegisteredChannels notifiedChannels

func CalculateTimeRemain() string {
	now := time.Now().UTC()
	duration := hibanaConcert.Sub(now)

	daysRemaining := int(duration.Hours() / 24)
	months := daysRemaining / 30
	days := daysRemaining % 30

	hours := int(duration.Hours()) % 24
	minutes := int(duration.Minutes()) % 60

	return fmt.Sprintf("Hibana in: %d months %d days, %d hours %d minutes", months, days, hours, minutes)
}

func RegisterChannel(channelID string) {
	for _, channel := range RegisteredChannels.Channels {
		if channel.ID == channelID {
			return
		}
	}
	RegisteredChannels.Channels = append(RegisteredChannels.Channels, channel{ID: channelID})
	RegisteredChannels.writeToFile("hibana_channels.json")
}

func UnregisterChannel(channelID string) {
	RegisteredChannels.removeChannel(channel{ID: channelID})
	RegisteredChannels.writeToFile("hibana_channels.json")
}

func init() {
	err := RegisteredChannels.readFromFile("hibana_channels.json")
	if err != nil {
		fmt.Printf("Error reading channels file: %v\n", err)
		RegisteredChannels.Channels = []channel{}
	}
	fmt.Printf("Initialized with %d channels\n", len(RegisteredChannels.Channels))
}
