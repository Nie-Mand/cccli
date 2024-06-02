package commitmessage_model

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (m CommitMessageModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "backspace":
			if len(m.message) > 0 {
				m.message = m.message[:len(m.message)-1]
			}

			return m, nil

		case "enter":
			return m, nil

		default:
			m.message += msg.String()
			return m, nil
		}
	}

	return m, nil
}
