package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/Kelv1nG/hibana_countdown/bot"
)

func main() {
	var err error
	sess := bot.CreateSession()

	err = sess.Open()
	if err != nil {
		log.Fatalf("Failed to open connection: %v", err)
	}
	bot.ScheduleCountdown(sess)

	log.Println("Bot is running. Press CTRL+C to exit.")

	// wait for ctrl + c to exit
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop

	log.Println("Shutting down bot...")
	sess.Close()
}
