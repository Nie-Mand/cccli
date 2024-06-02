package committype_model

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (m CommitTypeModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "up":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down":
			if m.cursor+1 != len(m.choices) {
				m.cursor++
			}

		case "enter":
			m.selected = m.choices[m.cursor]
			return m, nil

		case "backspace":
			if len(m.query) > 0 {
				m.query = m.query[:len(m.query)-1]
				m.cursor = 0
			}

			return m, nil

		default:
			m.query += msg.String()
			m.cursor = 0
			return m, nil
		}
	}

	return m, nil
}
