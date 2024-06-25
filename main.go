package main

import (
    "log"
    "os"
    "os/signal"
    "syscall"

    "discordlifebot/config"
    "discordlifebot/commands"
    "discordlifebot/reminders"

    "github.com/bwmarrin/discordgo"
)

func main() {
    token, channelID := config.LoadConfig()

    dg, err := discordgo.New("Bot " + token)
    if err != nil {
        log.Fatalf("Error creating discord session : %v", err)
    }

    dg.AddHandler(commands.MessageCreate)

    err = dg.Open()
    if err != nil {
        log.Fatalf("Error opening connection : %v", err)
    }

    log.Println("Bot is now running")

    // Start reminders
    go reminders.StartDailyReminders(dg, channelID)

    // Wait for a signal to gracefully shut down the bot
    stop := make(chan os.Signal, 1)
    signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
    <-stop

    log.Println("Shutting down the bot")
    dg.Close()
}
