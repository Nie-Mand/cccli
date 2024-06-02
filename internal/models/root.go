package models

import (
	root_model "github.com/Nie-Mand/cccli/internal/models/root"
	tea "github.com/charmbracelet/bubbletea"
)

func Init() tea.Model {
	return root_model.InitRootModel()
}
