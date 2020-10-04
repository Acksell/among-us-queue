package bot

import "github.com/bwmarrin/discordgo"

// QueuedUsers is a set of discord users wanting to play.
type QueuedUsers map[discordgo.User]bool

// Add a user to the set of queued users. Not safe for concurrent use.
func (q QueuedUsers) Add(user discordgo.User) {
	q[user] = true
}

// Remove a user from the set of queued users. Not safe for concurrent use.
func (q QueuedUsers) Remove(user discordgo.User) {
	delete(q, user)
}

// Size returns the number of queued users.
func (q QueuedUsers) Size() int {
	return len(q)
}
