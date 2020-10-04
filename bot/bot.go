package bot

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

// GuildQueues provides the QueuedUsers for a given guild ID.
var GuildQueues = make(map[string]QueuedUsers)

// MakeAndListen creates a new discord bot session and listens for messages.
func MakeAndListen(token string) {
	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)

	// In this example, we only care about receiving message events.
	dg.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages)

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	switch parsed := Parse(m.Content); parsed.Command {
	case QueueCmd:
		QueueHandler(s, m, parsed.Args)
	case LeaveCmd:
		LeaveHandler(s, m, parsed.Args)
	case ViewCmd:
		ViewHandler(s, m, parsed.Args)
	}
}

// QueueHandler handles a message from a user that want's to queue up.
func QueueHandler(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	// if Guild does not have a queue, add one.
	if _, ok := GuildQueues[m.GuildID]; !ok {
		GuildQueues[m.GuildID] = QueuedUsers{}
	}
	GuildQueues[m.GuildID].Add(*m.Author) // Add user to queue.
	s.ChannelMessageSend(m.ChannelID, "Ok, you are now queued.")
}

// LeaveHandler handles a message from a user that want's to leave the queue.
func LeaveHandler(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	GuildQueues[m.GuildID].Remove(*m.Author)
	s.ChannelMessageSend(m.ChannelID, "Ok, you left the queue.")
}

// ViewHandler handles a message from a user that want's to leave the queue.
func ViewHandler(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	var usernames []string
	for user := range GuildQueues[m.GuildID] {
		usernames = append(usernames, user.Username)
	}
	if len(usernames) > 0 {
		formatted := fmt.Sprintf("Here is the list of currently queued users\n```%s```", strings.Join(usernames, "\n"))
		s.ChannelMessageSend(m.ChannelID, formatted)
	} else {
		s.ChannelMessageSend(m.ChannelID, "Noone is queued!")
	}
}
