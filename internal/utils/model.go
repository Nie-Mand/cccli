package utils

import tea "github.com/charmbracelet/bubbletea"

type Model interface {
	tea.Model

	Summary() string
}
