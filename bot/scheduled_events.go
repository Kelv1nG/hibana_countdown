package bot

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Kelv1nG/hibana_countdown/hibana"
	"github.com/Kelv1nG/hibana_countdown/spotify"
	"github.com/bwmarrin/discordgo"
)

func ScheduleCountdown(sess *discordgo.Session) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	scheduledTimes := []struct {
		hour   int
		minute int
	}{
		{0, 0},
		{23, 59},
	}

	go func() {
		for {
			now := time.Now()
			currentHour := now.Hour()
			currentMinute := now.Minute()
			currentSecond := now.Second()

			// find the next scheduled time
			var nextTime time.Time
			minDuration := 24 * time.Hour // Start with max duration (1 day)

			for _, scheduled := range scheduledTimes {
				// calculate candidate time for today
				candidateTime := time.Date(now.Year(), now.Month(), now.Day(),
					scheduled.hour, scheduled.minute, 0, 0, now.Location())

				// if the time has already passed today, schedule for tomorrow
				if candidateTime.Before(now) {
					candidateTime = candidateTime.Add(24 * time.Hour)
				}

				// check if this is sooner than current next time
				duration := candidateTime.Sub(now)
				if duration < minDuration {
					minDuration = duration
					nextTime = candidateTime
				}
			}

			sleepDuration := nextTime.Sub(now)
			fmt.Printf("Current time: %02d:%02d:%02d\n", currentHour, currentMinute, currentSecond)
			fmt.Printf("Next message scheduled at: %s (in %s)\n",
				nextTime.Format("2006-01-02 15:04:05"),
				sleepDuration.String())

			select {
			case <-time.After(sleepDuration):
				currentTime := time.Now().Format("15:04:05")
				message := fmt.Sprintf("Scheduled announcement - Current time: %s", currentTime)
				fmt.Println(message)
				registeredChannels := hibana.RegisteredChannels
				fmt.Printf("Channel count: %v\n", len(registeredChannels.Channels))
				for _, channel := range registeredChannels.Channels {
					fmt.Printf("Sending message to channel %s\n", channel.ID)
					_, err := sess.ChannelMessageSend(channel.ID, hibana.CalculateTimeRemain())
					if err != nil {
						log.Printf("Error sending to channel %s: %v\n", channel.ID, err)
					}
					url, err := spotify.RandomHibanaSong()
					if err != nil {
						log.Print(err)
					}
					_, err = sess.ChannelMessageSend(channel.ID, url)
					if err != nil {
						log.Printf("Error sending to channel %s: %v\n", channel.ID, err)
					} else {
						fmt.Printf("Successfully sent message to channel %s\n", channel.ID)
					}
				}
			case <-stop:
				fmt.Println("Shutting down scheduler")
				return
			}
		}
	}()

	// wait for interrupt signal
	<-stop
	log.Println("Bot shutting down")
}
