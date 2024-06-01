package models

import (
	committype_model "github.com/Nie-Mand/cccli/internal/models/commit_type"
	tea "github.com/charmbracelet/bubbletea"
)

func Init() tea.Model {
	return committype_model.InitCommitTypeModel()
}
