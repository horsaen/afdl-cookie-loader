package views

import (
	"fmt"
	"horsaen/afdl-cookie-loader/tools"
	"strings"
)

func Platforms(m model) string {
	c := m.platform

	tpl := "Select Platform:\n\n"
	tpl += "%s\n\n"
	tpl += subtle("tab/shift+tab, up/down: select") + dot + subtle("enter: choose") + dot + subtle("q, esc: quit")

	Platforms := fmt.Sprintf(
		"%s\n%s\n",
		tools.Checkbox("Afreeca", c == 0),
		tools.Checkbox("Flex", c == 1),
		// tools.Checkbox("Panda", c == 2),
	)

	return fmt.Sprintf(tpl, Platforms)
}

func Inputs(m model) string {
	var b strings.Builder

	for i := range m.inputs {
		b.WriteString(m.inputs[i].View())
		if i < len(m.inputs)-1 {
			b.WriteRune('\n')
		}
	}

	button := &blurredButton
	if m.focusIndex == len(m.inputs) {
		button = &focusedButton
	}
	fmt.Fprintf(&b, "\n\n%s\n\n", *button)

	return b.String()
}
