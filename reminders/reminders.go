package reminders

import (
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
	"discordlifebot/commands"
)

func StartDailyReminders(s *discordgo.Session, channelID string) {
	for {
		now := time.Now()
		next := now.Add(time.Hour * 24)
		next = time.Date(next.Year(), next.Month(), next.Day(), 9, 0, 0, 0, next.Location())

		duration := next.Sub(now)
		time.Sleep(duration)

		sendDailyReminders(s, channelID)
	}
}

func sendDailyReminders(s *discordgo.Session, channelID string) {
	for projectName, tasks := range commands.Projects {
		for taskName, task := range tasks {
			// Check if task is due today or if it's overdue
			deadline, err := time.Parse("2006-01-02", task.Deadline)
			if err != nil {
				log.Printf("Error parsing deadline for task '%s': %v", taskName, err)
				continue
			}

			if deadline.Before(time.Now()) || deadline.Equal(time.Now()) {
				// Send reminder for overdue tasks
				s.ChannelMessageSend(channelID, "Task '"+taskName+"' in project '"+projectName+"' is due today or overdue!")
			}
		}
	}
}

