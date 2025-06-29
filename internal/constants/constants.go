package constants

import "slices"

var (
	QUIT   = []string{"q", "quit", "ctrl+c", "ctrl+d"}
	UP     = []string{"up", "k", "ctrl+p"}
	ESCAPE = []string{"esc", "escape"}
	DOWN   = []string{"down", "j", "ctrl+n"}
)

func IsCommand(input string, commands []string) bool {
	return slices.Contains(commands, input)
}
