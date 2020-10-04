package bot

import (
	"strings"
)

// ParsedMessage contains parsed data from the underlying user message.
type ParsedMessage struct {
	Command CommandType
	Args    []string
}

// Parse returns a ParsedMessage given a user message.
func Parse(message string) *ParsedMessage {
	oldLen := len(message)
	message = strings.Replace(message, CommandPrefix+" ", "", 1)
	if len(message) == oldLen { //didn't have a space, allows typos like !auqqueue
		message = strings.Replace(message, CommandPrefix, "", 1)
	}

	args := strings.Split(message, " ")
	cmd := GetCommandType(args[0])
	parsed := &ParsedMessage{cmd, args}

	return parsed
}
