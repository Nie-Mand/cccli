package root_model

import tea "github.com/charmbracelet/bubbletea"

func (m RootModel) Init() tea.Cmd {
	return tea.SetWindowTitle("Conventional Commits CLI")
}
