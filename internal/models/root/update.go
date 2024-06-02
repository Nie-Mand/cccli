package root_model

import (
	"github.com/Nie-Mand/cccli/internal/commands"
	"github.com/Nie-Mand/cccli/internal/core/domain"
	commitmessage_model "github.com/Nie-Mand/cccli/internal/models/commit_message"
	committype_model "github.com/Nie-Mand/cccli/internal/models/commit_type"
	"github.com/Nie-Mand/cccli/internal/utils"
	tea "github.com/charmbracelet/bubbletea"
)

func (m RootModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "enter":
			updatedModel, cmd := m.phases[m.cursor].Update(msg)
			m.phases[m.cursor] = updatedModel.(utils.Model)

			if m.cursor >= len(m.phases)-1 {

				commit := domain.Commit{
					CommitType:    m.phases[0].(committype_model.CommitTypeModel).Value(),
					CommitMessage: m.phases[1].(commitmessage_model.CommitMessageModel).Value(),
				}

				return m, commands.CommitCommand(commit)
			}

			m.cursor++
			return m, cmd

		default:
			updatedModel, cmd := m.phases[m.cursor].Update(msg)
			m.phases[m.cursor] = updatedModel.(utils.Model)
			return m, cmd
		}
	}
	return m, nil
}
