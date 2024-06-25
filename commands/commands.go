package commands

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

var Projects = make(map[string]map[string]Task) // projects[projectName] = tasks[taskName] = Task

type Task struct {
	Description string
	Deadline    string
	Priority    string
}

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	content := strings.ToLower(m.Content)
	if strings.HasPrefix(content, "!createproject") {
		createProject(s, m)
	} else if strings.HasPrefix(content, "!createtask") {
		createTask(s, m)
	}
}

func createProject(s *discordgo.Session, m *discordgo.MessageCreate) {
	projectName := strings.TrimSpace(strings.TrimPrefix(m.Content, "!createproject"))
	if projectName == "" {
		s.ChannelMessageSend(m.ChannelID, "Please provide a project name.")
		return
	}

	if _, exists := Projects[projectName]; exists {
		s.ChannelMessageSend(m.ChannelID, "Project already exists.")
		return
	}

	Projects[projectName] = make(map[string]Task)
	s.ChannelMessageSend(m.ChannelID, "Project '"+projectName+"' created successfully.")
}

func createTask(s *discordgo.Session, m *discordgo.MessageCreate) {
	parts := strings.SplitN(strings.TrimSpace(strings.TrimPrefix(m.Content, "!createtask")), " ", 4)
	if len(parts) < 4 {
		s.ChannelMessageSend(m.ChannelID, "Please provide all task details: project name, task description, deadline, and priority.")
		return
	}

	projectName, taskName, deadline, priority := parts[0], parts[1], parts[2], parts[3]
	if _, exists := Projects[projectName]; !exists {
		s.ChannelMessageSend(m.ChannelID, "Project does not exist.")
		return
	}

	Projects[projectName][taskName] = Task{Description: taskName, Deadline: deadline, Priority: priority}
	s.ChannelMessageSend(m.ChannelID, "Task '"+taskName+"' created in project '"+projectName+"' successfully.")
}

