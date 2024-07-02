package views

import (
	"horsaen/afdl-cookie-loader/cookies"

	tea "github.com/charmbracelet/bubbletea"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return m, tea.Quit

		case "tab", "shift+tab", "enter", "up", "down":
			if m.platformSelected && !m.submitted {
				return m.HandleInputNavigation(msg.String())
			} else {
				return m.HandlePlatformNavigation(msg.String())
			}
		}
	}

	cmd := m.UpdateInputs(msg)

	return m, cmd
}

func (m model) HandlePlatformNavigation(key string) (tea.Model, tea.Cmd) {
	if key == "enter" {
		m.platformSelected = true
	}

	if key == "up" || key == "shift+tab" {
		m.platform--
		if m.platform < 0 {
			m.platform = 0
		}
	} else {
		m.platform++
		if m.platform > 1 {
			m.platform = 1
		}
	}

	return m, nil
}

func (m model) HandleInputNavigation(key string) (tea.Model, tea.Cmd) {
	if key == "enter" && m.focusIndex == len(m.inputs) {
		m.submitted = true

		switch m.platform {
		case 1:
			// Afreeca
			cookies.Afreeca(m.inputs[0].Value(), m.inputs[1].Value())
			return m, tea.Quit
		case 2:
			// Flex
			cookies.Flex(m.inputs[0].Value(), m.inputs[1].Value())
			// case 3:
			// Panda
		}

		return m, nil
	}

	if key == "up" || key == "shift+tab" {
		m.focusIndex--
	} else {
		m.focusIndex++
	}

	if m.focusIndex > len(m.inputs) {
		m.focusIndex = 0
	} else if m.focusIndex < 0 {
		m.focusIndex = len(m.inputs)
	}

	cmds := make([]tea.Cmd, len(m.inputs))
	for i := 0; i <= len(m.inputs)-1; i++ {
		if i == m.focusIndex {
			cmds[i] = m.inputs[i].Focus()
			m.inputs[i].PromptStyle = focusedStyle
			m.inputs[i].TextStyle = focusedStyle
			continue
		}
		m.inputs[i].Blur()
		m.inputs[i].PromptStyle = noStyle
		m.inputs[i].TextStyle = noStyle
	}

	return m, tea.Batch(cmds...)
}

func (m *model) UpdateInputs(msg tea.Msg) tea.Cmd {
	if m.platformSelected {

		cmds := make([]tea.Cmd, len(m.inputs))

		for i := range m.inputs {
			m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
		}

		return tea.Batch(cmds...)
	}

	return nil
}
