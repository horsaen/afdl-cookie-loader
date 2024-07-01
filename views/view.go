package views

import (
	"fmt"
	"horsaen/afdl-cookie-loader/tools"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	focusedStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	blurredStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	cursorStyle   = focusedStyle
	noStyle       = lipgloss.NewStyle()
	focusedButton = focusedStyle.Render("[ Submit ]")
	blurredButton = fmt.Sprintf("[ %s ]", blurredStyle.Render("Submit"))
	// keyword       = tools.MakeFgStyle("211")
	subtle = tools.MakeFgStyle("241")
	dot    = tools.ColorFg(" • ", "236")
)

type model struct {
	platform         int
	platformSelected bool
	focusIndex       int
	inputs           []textinput.Model
	submitted        bool
}

func InitialModel() model {
	m := model{
		platform:         0,
		platformSelected: false,
		inputs:           make([]textinput.Model, 2),
		submitted:        false,
	}

	var t textinput.Model
	for i := range m.inputs {
		t = textinput.New()
		t.Cursor.Style = cursorStyle
		t.CharLimit = 32

		switch i {
		case 0:
			t.Placeholder = "Username"
			t.Focus()
			t.PromptStyle = focusedStyle
			t.TextStyle = focusedStyle
		case 1:
			t.Placeholder = "Password"
			t.EchoMode = textinput.EchoPassword
			t.EchoCharacter = '•'
		}

		m.inputs[i] = t
	}

	return m
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) View() string {
	var s string

	if !m.platformSelected {
		s = Platforms(m)
	} else if !m.submitted {
		s = Inputs(m)
	}

	return s
}
