package tools

import (
	"fmt"

	"github.com/muesli/termenv"
)

var (
	term = termenv.EnvColorProfile()
)

// bubbletea tools
func Checkbox(label string, checked bool) string {
	if checked {
		return ColorFg("[x] "+label, "212")
	}
	return fmt.Sprintf("[ ] %s", label)
}

func ColorFg(val, color string) string {
	return termenv.String(val).Foreground(term.Color(color)).String()
}

// Return a function that will colorize the foreground of a given string.
func MakeFgStyle(color string) func(string) string {
	return termenv.Style{}.Foreground(term.Color(color)).Styled
}
