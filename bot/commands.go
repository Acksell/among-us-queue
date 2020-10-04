package bot

import "strings"

// CommandPrefix is the prefix to get the bot's attention in discord.
var CommandPrefix = "!amonq"

// CommandType is the command that follows after the CommandPrefix.
type CommandType int

// Enum for CommandTypes
const (
	HelpCmd CommandType = iota
	QueueCmd
	LeaveCmd
	ViewCmd
	NullCmd
)

// CommandTypeStringMapping translates a string to the corresponding CommandType.
var CommandTypeStringMapping = map[string]CommandType{
	"help":  HelpCmd,
	"queue": QueueCmd,
	"leave": LeaveCmd,
	"view":  ViewCmd,
	"":      NullCmd,
}

// GetCommandType returns the corresponding CommandType given a string. Is provided by CommandTypeStringMapping.
func GetCommandType(arg string) CommandType {
	for str, cmd := range CommandTypeStringMapping {
		// allow single letter commands
		if len(arg) == 1 && cmd != NullCmd {
			if str[0] == arg[0] {
				return cmd
			}
		} else {
			if strings.ToLower(arg) == str {
				return cmd
			}
		}
	}
	return NullCmd
}
